module linters

go 1.16

require (
	github.com/golangci/golangci-lint v1.39.0
	golang.org/x/tools v0.1.0
	honnef.co/go/tools v0.1.3
)

// somehow the v1.0.0 version on go mod proxy is inconsistent with using direct ..
replace github.com/tomarrell/wrapcheck v1.0.0 => github.com/tomarrell/wrapcheck v1.1.0
