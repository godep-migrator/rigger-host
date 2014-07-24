rigger-host
===========

Rigger CI host service

[![GoDoc](https://godoc.org/github.com/rigger-dot-io/rigger-host?status.svg)](https://godoc.org/github.com/rigger-dot-io/rigger-host)

## Development ##

To get started install [godeps](https://github.com/tools/godep) first.

### Compiling ###

To compile the rigger binary just call:

```
$ make all
```

It will create a `bin` folder in the project root and it will also put a copy into your `GOPATH/bin` for convenience.

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
