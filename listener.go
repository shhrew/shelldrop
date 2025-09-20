package main

import (
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
	connChan chan net.Conn
	listener net.Listener
}

func NewReverseShellListener(cfg ListenerConfig) *Listener {
	return &Listener{
		ListenerConfig: cfg,
		connChan:       make(chan net.Conn, 1), // Buffer of 1 since only 1 connection pending
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

		select {
		case r.connChan <- conn:
		default:
			log.Warn("Connection rejected: already have a pending connection")
			conn.Close()
		}
	}
}

func (r *Listener) Interact() error {
	conn := <-r.connChan
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

	log.Info("Reverse shell closed")
	return nil
}

func (r *Listener) Close() error {
	if r.listener != nil {
		return r.listener.Close()
	}
	return nil
}
