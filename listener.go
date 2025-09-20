package main

import (
	"context"
	"errors"
	"io"
	"net"
	"os"
	"shelldrop/log"
	"strconv"
)

type Listener struct {
	ListenerConfig
	listener    net.Listener
	connections chan net.Conn
	cancel      context.CancelFunc
}

func NewListener(cfg ListenerConfig, cancel context.CancelFunc) *Listener {
	return &Listener{
		ListenerConfig: cfg,
		connections:    make(chan net.Conn, 1),
		cancel:         cancel,
	}
}

func (r *Listener) Start() {
	if r.ListenerConfig.Disabled {
		log.Warn("Built-in listener is disabled, ensure your own is running on the specified host and port")
		return
	}

	listener, err := net.Listen("tcp", r.ListenerConfig.Host+":"+strconv.Itoa(r.ListenerConfig.Port))
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}
	r.listener = listener

	log.Infof("Reverse shell listener started on %s:%d", r.ListenerConfig.Host, r.ListenerConfig.Port)

	go r.acceptConnections()
}

func (r *Listener) acceptConnections() {
	for {
		conn, err := r.listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}

			log.Errorf("Failed to accept connection: %v", err)
			return
		}

		log.Successf("Connection received from %s", conn.RemoteAddr().String())
		r.cancel()

		select {
		case r.connections <- conn:
		default:
			log.Warn("Connection rejected: already have a pending connection")
			conn.Close()
		}
	}
}

func (r *Listener) Interact() {
	conn := <-r.connections
	defer conn.Close()

	log.Infof("Dropping shell to %s", conn.RemoteAddr().String())

	done := make(chan struct{}, 2)

	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	go func() {
		io.Copy(conn, os.Stdin)
		done <- struct{}{}
	}()

	<-done
}

func (r *Listener) Close() error {
	if r.listener != nil {
		return r.listener.Close()
	}
	return nil
}
