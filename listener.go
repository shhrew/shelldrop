package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"shelldrop/log"
	"strconv"
)

type Listener struct {
	ListenerConfig
}

func NewReverseShellListener(cfg ListenerConfig) *Listener {
	return &Listener{
		ListenerConfig: cfg,
	}
}

func (r *Listener) Start() error {
	listener, err := net.Listen("tcp", r.ListenerConfig.Host+":"+strconv.Itoa(r.ListenerConfig.Port))
	if err != nil {
		return fmt.Errorf("failed to start listener: %v", err)
	}
	defer listener.Close()

	log.Infof("Reverse shell listener started on %s:%d", r.ListenerConfig.Host, r.ListenerConfig.Port)
	log.Info("Waiting for connections...")

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Failed to accept connection: %v", err)
	}
	defer conn.Close()

	log.Infof("Connection received from %s", conn.RemoteAddr().String())

	// Use channels to detect when either direction closes
	done := make(chan struct{}, 2)

	// Connection -> Stdout
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	// Stdin -> Connection
	go func() {
		io.Copy(conn, os.Stdin)
		done <- struct{}{}
	}()

	<-done

	log.Info("Session ended")
	return nil
}
