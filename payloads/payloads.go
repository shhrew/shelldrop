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
	"awk_1":        awk1,
	"bash_tcp_1":   bashTcp1,
	"bash_tcp_2":   bashTcp2,
	"bash_tcp_3":   bashTcp3,
	"bash_udp_1":   bashUdp1,
	"go_1":         go1,
	"lua_1":        lua1,
	"lua_2":        lua2,
	"ncat_1":       ncat1,
	"ncat_2":       ncat2,
	"netcat_1":     netcat1,
	"netcat_2":     netcat2,
	"netcat_3":     netcat3,
	"netcat_4":     netcat4,
	"netcat_5":     netcat5,
	"perl_1":       perl1,
	"perl_2":       perl2,
	"perl_win_1":   perlWin1,
	"php_1":        php1,
	"php_2":        php2,
	"php_3":        php3,
	"php_4":        php4,
	"php_5":        php5,
	"php_6":        php6,
	"php_7":        php7,
	"powershell_1": powershell1,
	"powershell_2": powershell2,
	"python_1":     python1,
	"python_2":     python2,
	"python_3":     python3,
	"python_4":     python4,
	"python_5":     python5,
	"python_6":     python6,
	"python_7":     python7,
	"python_8":     python8,
	"python_9":     python9,
	"python_10":    python10,
	"python_11":    python11,
	"python_12":    python12,
	"python_13":    python13,
	"python_win_1": pythonWin1,
	"python_win_2": pythonWin2,
	"ruby_1":       ruby1,
	"ruby_2":       ruby2,
	"ruby_win_1":   rubyWin1,
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

//go:embed awk_1.txt
var awk1 string

//go:embed bash_tcp_1.txt
var bashTcp1 string

//go:embed bash_tcp_2.txt
var bashTcp2 string

//go:embed bash_tcp_3.txt
var bashTcp3 string

//go:embed bash_udp_1.txt
var bashUdp1 string

//go:embed go_1.txt
var go1 string

//go:embed lua_1.txt
var lua1 string

//go:embed lua_2.txt
var lua2 string

//go:embed ncat_1.txt
var ncat1 string

//go:embed ncat_2.txt
var ncat2 string

//go:embed netcat_1.txt
var netcat1 string

//go:embed netcat_2.txt
var netcat2 string

//go:embed netcat_3.txt
var netcat3 string

//go:embed netcat_4.txt
var netcat4 string

//go:embed netcat_5.txt
var netcat5 string

//go:embed perl_1.txt
var perl1 string

//go:embed perl_2.txt
var perl2 string

//go:embed perl_win_1.txt
var perlWin1 string

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

//go:embed php_7.txt
var php7 string

//go:embed powershell_1.txt
var powershell1 string

//go:embed powershell_2.txt
var powershell2 string

//go:embed python_1.txt
var python1 string

//go:embed python_2.txt
var python2 string

//go:embed python_3.txt
var python3 string

//go:embed python_4.txt
var python4 string

//go:embed python_5.txt
var python5 string

//go:embed python_6.txt
var python6 string

//go:embed python_7.txt
var python7 string

//go:embed python_8.txt
var python8 string

//go:embed python_9.txt
var python9 string

//go:embed python_10.txt
var python10 string

//go:embed python_11.txt
var python11 string

//go:embed python_12.txt
var python12 string

//go:embed python_13.txt
var python13 string

//go:embed python_win_1.txt
var pythonWin1 string

//go:embed python_win_2.txt
var pythonWin2 string

//go:embed ruby_1.txt
var ruby1 string

//go:embed ruby_2.txt
var ruby2 string

//go:embed ruby_win_1.txt
var rubyWin1 string
