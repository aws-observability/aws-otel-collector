export GOBIN ?= $(shell pwd)/bin

<<<<<<< HEAD
GOLINT = $(GOBIN)/golint
STATICCHECK = $(GOBIN)/staticcheck
=======
REVIVE = $(GOBIN)/revive
STATICCHECK = $(GOBIN)/staticcheck
GOVULNCHECK = $(GOBIN)/govulncheck
>>>>>>> main
BENCH_FLAGS ?= -cpuprofile=cpu.pprof -memprofile=mem.pprof -benchmem

# Directories containing independent Go modules.
#
# We track coverage only for the main module.
<<<<<<< HEAD
MODULE_DIRS = . ./benchmarks ./zapgrpc/internal/test
=======
MODULE_DIRS = . ./exp ./benchmarks ./zapgrpc/internal/test
>>>>>>> main

# Many Go tools take file globs or directories as arguments instead of packages.
GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

.PHONY: all
all: lint test

.PHONY: lint
<<<<<<< HEAD
lint: $(GOLINT) $(STATICCHECK)
=======
lint: $(REVIVE) $(STATICCHECK)
>>>>>>> main
	@rm -rf lint.log
	@echo "Checking formatting..."
	@gofmt -d -s $(GO_FILES) 2>&1 | tee lint.log
	@echo "Checking vet..."
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go vet ./... 2>&1) &&) true | tee -a lint.log
	@echo "Checking lint..."
<<<<<<< HEAD
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && $(GOLINT) ./... 2>&1) &&) true | tee -a lint.log
=======
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && \
		$(REVIVE) -set_exit_status ./... 2>&1) &&) true | tee -a lint.log
>>>>>>> main
	@echo "Checking staticcheck..."
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && $(STATICCHECK) ./... 2>&1) &&) true | tee -a lint.log
	@echo "Checking for unresolved FIXMEs..."
	@git grep -i fixme | grep -v -e Makefile | tee -a lint.log
	@echo "Checking for license headers..."
	@./checklicense.sh | tee -a lint.log
	@[ ! -s lint.log ]
	@echo "Checking 'go mod tidy'..."
	@make tidy
	@if ! git diff --quiet; then \
		echo "'go mod tidy' resulted in changes or working tree is dirty:"; \
		git --no-pager diff; \
	fi

<<<<<<< HEAD
$(GOLINT):
	cd tools && go install golang.org/x/lint/golint
=======
$(REVIVE):
	cd tools && go install github.com/mgechev/revive

$(GOVULNCHECK):
	cd tools && go install golang.org/x/vuln/cmd/govulncheck
>>>>>>> main

$(STATICCHECK):
	cd tools && go install honnef.co/go/tools/cmd/staticcheck

.PHONY: test
test:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go test -race ./...) &&) true

.PHONY: cover
cover:
	go test -race -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html

.PHONY: bench
BENCH ?= .
bench:
	@$(foreach dir,$(MODULE_DIRS), ( \
		cd $(dir) && \
		go list ./... | xargs -n1 go test -bench=$(BENCH) -run="^$$" $(BENCH_FLAGS) \
	) &&) true

.PHONY: updatereadme
updatereadme:
	rm -f README.md
	cat .readme.tmpl | go run internal/readme/readme.go > README.md

.PHONY: tidy
tidy:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go mod tidy) &&) true
<<<<<<< HEAD
=======

.PHONY: vulncheck
vulncheck: $(GOVULNCHECK)
	$(GOVULNCHECK) ./...
>>>>>>> main
