//go:build tools
// +build tools

package tools

// Based on https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "go.opentelemetry.io/build-tools/dbotconf"
	_ "golang.org/x/tools/cmd/goimports"
	_ "honnef.co/go/tools/cmd/staticcheck"
	_ "mvdan.cc/sh/v3/cmd/shfmt"
)
