# go-fqdn

Go package to provide reasonable robust access to fully qualified hostname. It
first tries to looks up your hostname in hosts file. If that fails, it falls
back to doing lookup via dns.

Basically it tries to mirror how standard linux `hostname -f` works. For that
reason, your hosts file should be configured properly, please refer to hosts(5)
for that.

It also has no 3rd party dependencies.

## Usage

This package uses go modules, so just writing code that uses it should be
enough. For example of usage you can check out
[the example](examples/example_test.go).

Documentation can be found
[here](https://pkg.go.dev/github.com/Showmax/go-fqdn?tab=doc).

## Supported go versions

Current and current - 1 versions of go are supported.

## Known issues

On macos, when **not** using cgo (`CGO_ENABLED=0`), getting the fqdn hostname
might not work. Depends on rest of your setup and how `/etc/resolv.conf` looks
like. Since that file is not used much (at least based on documentation) by
macos programs, it is possible it is not in correct state.
