.PHONY: default
default: test

.PHONY: all
all: test

.PHONY: test
test:
	go test -v ./... -bench . -benchtime 1ms

.PHONY: generate
generate:
	cd pkg/tablegen && go install .
	cd pkg && go generate
