rigger-host
===========

Rigger CI host service

[![GoDoc](https://godoc.org/github.com/rigger-dot-io/rigger-host?status.svg)](https://godoc.org/github.com/rigger-dot-io/rigger-host)

## Usage ##

To start a rigger host run:

```
$ sudo rigger [options]
```

To see all available options:

```
$ sudo rigger -h
Usage of rigger:
  -c, --config="/etc/rigger.conf"        Load configuration from file
  -d, --daemon=false                     Enable daemon mode
  -l, --logfile="/var/log/rigger.log"    Path to rigger log file
  -p, --port=9876                        RPC Port
  --pidfile="/var/run/rigger.pid"        Path to use for PID file
  -v, --version=false                    Print version information and quit
```

## Development ##

To get started install [godeps](https://github.com/tools/godep) first.

### Building ###

To compile rigger just call:

```
$ make all
```

It will create a `bin` folder in the project root and will also put a copy into your `GOPATH/bin` for convenience.

### Testing ###

You've got two options to run the test suite:

```
$ make test
```

Or if you'd like to see a coverage report:

```
$ make cov
```

*Note: coverage report requires [gocov](https://github.com/axw/gocov) and [gocov-html](https://github.com/matm/gocov-html).*

## Status ##

**Rigger is in active development and currently unsuitable for production use.**

Users are encouraged to experiment with Rigger but should assume there are stability, security, and performance weaknesses throughout the project.

Please report bugs as issues on the appropriate repository.

## Contributing ##

We welcome and encourage community contributions to Rigger.

Since the project is still unstable, there are specific priorities for development. Pull requests that do not address these priorities will not be accepted until Rigger is production ready.

There are many ways to help besides contributing code:

- Fix bugs or file issues
- Improve the documentation
- Contribute financially to support core development (please contact hello@rigger.io)
