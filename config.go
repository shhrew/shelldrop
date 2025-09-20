package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

type (
	Config struct {
		Payload  string
		Listener ListenerConfig
		Request  RequestConfig
	}

	ListenerConfig struct {
		Disabled bool
		Host     string
		Port     int
	}

	RequestConfig struct {
		Url string
	}
)

func ParseConfig() *Config {
	parser := argparse.NewParser("shelldrop", `Leverages a command injection vulnerability by finding and executing compatible reverse shell payloads.

[*] = Asterisked arguments can contain the SHELLDROP injection keyword`)

	url := parser.String("u", "url", &argparse.Options{Required: true, Help: "The target url [*]"})
	lhost := parser.String("l", "lhost", &argparse.Options{Required: true, Help: "The listen address"})
	lport := parser.Int("p", "lport", &argparse.Options{Required: true, Help: "The listen port"})
	payload := parser.String("P", "payload", &argparse.Options{Required: false, Help: "Optional payload to use"})
	noListener := parser.Flag("", "no-listener", &argparse.Options{Required: false, Help: "Disable the built-in listener"})

	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	// todo: validate at least one argument has SHELLDROP keyword

	return &Config{
		Payload: *payload,
		Listener: ListenerConfig{
			Disabled: *noListener,
			Host:     *lhost,
			Port:     *lport,
		},
		Request: RequestConfig{
			Url: *url,
		},
	}
}
