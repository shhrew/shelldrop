package main

import (
	"context"
	"errors"
	"fmt"
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

func (r *Listener) Start() error {
	listener, err := net.Listen("tcp", r.ListenerConfig.Host+":"+strconv.Itoa(r.ListenerConfig.Port))
	if err != nil {
		return fmt.Errorf("failed to start listener: %v", err)
	}
	r.listener = listener

	log.Infof("Reverse shell listener started on %s:%d", r.ListenerConfig.Host, r.ListenerConfig.Port)
	log.Info("Ready for connections...")

	go r.acceptConnections()

	return nil
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

		log.Infof("Connection received from %s", conn.RemoteAddr().String())
		r.cancel()

		select {
		case r.connections <- conn:
		default:
			log.Warn("Connection rejected: already have a pending connection")
			conn.Close()
		}
	}
}

func (r *Listener) Interact() error {
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

	return nil
}

func (r *Listener) Close() error {
	if r.listener != nil {
		return r.listener.Close()
	}
	return nil
}
