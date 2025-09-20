package payloads

import (
	_ "embed"
	"net/url"
	"shelldrop/log"
	"sort"
	"strconv"
	"strings"
)

var payloads = map[string]string{
	"bash_tcp_1": bashTcp1,
	"bash_tcp_2": bashTcp2,
	"bash_tcp_3": bashTcp3,
	"php_1":      php1,
	"php_2":      php2,
	"php_3":      php3,
	"php_4":      php4,
	"php_5":      php5,
	"php_6":      php6,
}

func GetNames() []string {
	keys := make([]string, 0, len(payloads))
	for key := range payloads {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func Get(name string, lhost string, lport int) string {
	payload, ok := payloads[name]
	if !ok {
		log.Fatalf("Invalid payload specified: %s", name)
	}

	payload = strings.Replace(payload, "SHELLDROP_HOST", lhost, -1)
	payload = strings.Replace(payload, "SHELLDROP_PORT", strconv.Itoa(lport), -1)

	return payload
}

func GetUrlEncoded(name string, lhost string, lport int) string {
	return url.QueryEscape(Get(name, lhost, lport))
}

//go:embed bash_tcp_1.txt
var bashTcp1 string

//go:embed bash_tcp_2.txt
var bashTcp2 string

//go:embed bash_tcp_3.txt
var bashTcp3 string

//go:embed php_1.txt
var php1 string

//go:embed php_2.txt
var php2 string

//go:embed php_3.txt
var php3 string

//go:embed php_4.txt
var php4 string

//go:embed php_5.txt
var php5 string

//go:embed php_6.txt
var php6 string
