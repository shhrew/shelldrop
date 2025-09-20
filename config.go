package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

type (
	Config struct {
		Listener ListenerConfig
		Request  RequestConfig
	}

	ListenerConfig struct {
		Host string
		Port int
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

	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	return &Config{
		Listener: ListenerConfig{
			Host: *lhost,
			Port: *lport,
		},
		Request: RequestConfig{
			Url: *url,
		},
	}
}
