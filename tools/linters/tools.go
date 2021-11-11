package tools

// Based on https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md

import (
	_ "golang.org/x/tools/cmd/goimports"
	_ "honnef.co/go/tools/cmd/staticcheck"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "go.opentelemetry.io/build-tools/multimod"
)
