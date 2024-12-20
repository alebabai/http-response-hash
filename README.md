# http-response-hash

> A simple tool which makes http requests and prints the address of it along with the hash of the response

[![build](https://img.shields.io/github/actions/workflow/status/alebabai/http-response-hash/ci.yml)](https://github.com/alebabai/http-response-hash/actions?query=workflow%3ACI)
[![version](https://img.shields.io/github/go-mod/go-version/alebabai/http-response-hash)](https://go.dev/)
[![report](https://goreportcard.com/badge/github.com/alebabai/http-response-hash)](https://goreportcard.com/report/github.com/alebabai/http-response-hash)
[![coverage](https://img.shields.io/codecov/c/github/alebabai/http-response-hash)](https://codecov.io/github/alebabai/http-response-hash)
[![tag](https://img.shields.io/github/tag/alebabai/http-response-hash.svg)](https://github.com/alebabai/http-response-hash/tags)
[![reference](https://pkg.go.dev/badge/github.com/alebabai/http-response-hash.svg)](https://pkg.go.dev/github.com/alebabai/http-response-hash)

## Getting started

```bash
make install
```

```bash
$GOPATH/bin/hasher http://google.com
```

## Development

### Local

Build application artifacts:

```bash
make build
```

Run tests:

```bash
make test
```

Install application:

```bash
make install
```

## Usage

### Flags

- `-parallel` (default value is `10`) to limit the number of parallel requests

### Notes

Hasher supports addresses with and without a schema (`http` will be used by default).
Addresses must be space-separated, just like regular command-line arguments.

### Examples

```bash
go run ./cmd/hasher google.com twitter.com
```

```bash
$GOPATH/bin/hasher https://google.com
```

```bash
$GOPATH/bin/hasher -parallel 3 https://google.com facebook.com https://yahoo.com yandex.com twitter.com baroquemusiclibrary.com
```
