package main

import (
	"fmt"
	"os"
	"shelldrop/log"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/fatih/color"
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
	parser := argparse.NewParser("shelldrop", fmt.Sprintf(`A command injection tool that automatically tests for working reverse shell payloads.

[*] = Asterisked arguments can contain the SHELLDROP injection keyword`))

	url := parser.String("u", "url", &argparse.Options{Required: true, Help: "The target url [*]"})
	lhost := parser.String("l", "lhost", &argparse.Options{Required: true, Help: "The listen address"})
	lport := parser.Int("p", "lport", &argparse.Options{Required: true, Help: "The listen port"})
	payload := parser.String("P", "payload", &argparse.Options{Required: false, Help: "Optional payload to use"})
	noListener := parser.Flag("", "no-listener", &argparse.Options{Required: false, Help: "Disable the built-in listener"})
	noColor := parser.Flag("", "no-color", &argparse.Options{Required: false, Help: "Disable color output"})

	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if *noColor {
		color.NoColor = true
	}

	cfg := &Config{
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

	if !cfg.hasShellDropKeyword() {
		log.Fatal("No SHELLDROP injection keyword found in configuration")
	}

	return cfg
}

func (c *Config) hasShellDropKeyword() bool {
	if strings.Contains(c.Request.Url, ShellDropKeyword) {
		return true
	}

	return false
}
