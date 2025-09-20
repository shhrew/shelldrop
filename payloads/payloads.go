package payloads

import (
	_ "embed"
	"shelldrop/log"
)

var All = map[string]string{
	"bash_tcp_1": bashTcp1,
	"bash_tcp_2": bashTcp2,
	"bash_tcp_3": bashTcp3,
}

func Get(name string) string {
	payload, ok := All[name]
	if !ok {
		log.Fatalf("Invalid payload specified: %s", name)
	}

	return payload
}

//go:embed bash_tcp_1.txt
var bashTcp1 string

//go:embed bash_tcp_2.txt
var bashTcp2 string

//go:embed bash_tcp_3.txt
var bashTcp3 string
