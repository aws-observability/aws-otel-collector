include ./Makefile.Common

VERSION := $(shell cat VERSION)
PROJECTNAME := $(shell basename "$(PWD)")

GIT_SHA=$(shell git rev-parse HEAD)
DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
# All source code excluding any third party code and excluding the testbed.
# This is the code that we want to run tests for and lint, staticcheck, etc.
ALL_SRC := $(shell find . -name '*.go' \
							-not -path './testbed/*' \
							-not -path '*/third_party/*' \
							-not -path './.github/*' \
							-not -path './bin/*' \
							-not -path './build/*' \
							-not -path './tools/workflow/linters/*' \
							-not -path './vendor/*' \
							-type f | sort)

# ALL_MODULES includes ./* dirs (excludes . dir)
ALL_MODULES := $(shell find . -type f -name "go.mod" -exec dirname {} \; | sort | egrep  '^./' )

ALL_SHELL_SCRIPTS := $(shell find . -type f -name "*.sh" -not -path './vendor/*' )
SHELLCHECK_OPTS := "-e SC1071"

BUILD_INFO_IMPORT_PATH=$(AOC_IMPORT_PATH)/tools/version

GOBUILD=GO111MODULE=on CGO_ENABLED=0 installsuffix=cgo go build -trimpath

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-s -w -X $(BUILD_INFO_IMPORT_PATH).GitHash=$(GIT_SHA) \
-X $(BUILD_INFO_IMPORT_PATH).Version=$(VERSION) -X $(BUILD_INFO_IMPORT_PATH).Date=$(DATE)"

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
DOCKER_NAMESPACE=amazon
COMPONENT=awscollector
TOOLS_MOD_DIR := $(abspath ./tools/workflow/linters)
TOOLS_BIN_DIR := $(abspath ./bin)
DBOTCONF = $(TOOLS_BIN_DIR)/dbotconf
# Append root module to all modules
GOMODULES = $(ALL_MODULES) $(PWD)

# Define a delegation target for each module
.PHONY: $(GOMODULES)
$(GOMODULES):
	@echo "Running target '$(TARGET)' in module '$@'"
	TOOL_BIN=$(TOOLS_BIN_DIR) $(MAKE) -C $@ $(TARGET)

# Triggers each module's delegation target
.PHONY: for-all-target
for-all-target: $(GOMODULES)

PATCHES := $(shell find ./patches -name *.patch)
apply-patches: $(PATCHES)
	$(foreach patch,$(PATCHES), patch --posix --forward -p1 < $(patch);)

.PHONY: apply-patches

all-modules:
	@echo $(ALL_MODULES) | tr ' ' '\n' | sort

all-pkgs:
	@echo $(ALL_PKGS) | tr ' ' '\n' | sort

all-srcs:
	@echo $(ALL_SRC) | tr ' ' '\n' | sort


DEPENDABOT_CONFIG = .github/dependabot.yml
.PHONY: dependabot-check
dependabot-check: install-dbotconf
	@$(DBOTCONF) verify $(DEPENDABOT_CONFIG) || echo "(run: make dependabot-generate)"

.PHONY: dependabot-generate
dependabot-generate: install-dbotconf
	@$(DBOTCONF) generate > $(DEPENDABOT_CONFIG); 

.PHONY: build
build: install-tools golint
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./build/darwin/amd64/aoc ./cmd/awscollector
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./build/linux/amd64/aoc ./cmd/awscollector
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o ./build/linux/arm64/aoc ./cmd/awscollector
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./build/windows/amd64/aoc ./cmd/awscollector
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./build/darwin/amd64/healthcheck ./cmd/healthcheck
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./build/linux/amd64/healthcheck ./cmd/healthcheck
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o ./build/linux/arm64/healthcheck ./cmd/healthcheck
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./build/windows/amd64/healthcheck ./cmd/healthcheck


.PHONY: amd64-build
amd64-build: install-tools golint
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./build/linux/amd64/aoc ./cmd/awscollector

.PHONY: arm64-build
arm64-build: install-tools golint
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o ./build/linux/arm64/aoc ./cmd/awscollector

.PHONY: windows-build
windows-build: install-tools golint
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./build/windows/amd64/aoc ./cmd/awscollector

# For building container image during development, no lint nor other platforms
.PHONY: amd64-build-only
amd64-build-only:
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./build/linux/amd64/aoc ./cmd/awscollector

.PHONY: awscollector
awscollector:
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./bin/awscollector_$(GOOS)_$(GOARCH) ./cmd/awscollector
	GOOS=windows GOARCH=amd64 EXTENSION=.exe $(GOBUILD) $(LDFLAGS) -o ./bin/windows/aoc_windows_amd64.exe ./cmd/awscollector
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./bin/healthcheck_$(GOOS)_$(GOARCH) ./cmd/healthcheck
	GOOS=windows GOARCH=amd64 EXTENSION=.exe $(GOBUILD) $(LDFLAGS) -o ./bin/windows/healthcheck_windows_amd64.exe ./cmd/healthcheck

.PHONY: package-rpm
package-rpm: build
	ARCH=x86_64 SOURCE_ARCH=amd64 DEST=build/packages/linux/amd64 tools/packaging/linux/create_rpm.sh

.PHONY: package-deb
package-deb: build
	ARCH=amd64 DEST=build/packages/debian/amd64 tools/packaging/debian/create_deb.sh
	ARCH=arm64 DEST=build/packages/debian/arm64 tools/packaging/debian/create_deb.sh

.PHONY: docker-build
docker-build: amd64-build amd64-build-healthcheck
	docker buildx build --platform linux/amd64 --build-arg BUILDMODE=copy --load -t $(DOCKER_NAMESPACE)/$(COMPONENT):$(VERSION) -f ./cmd/$(COMPONENT)/Dockerfile .

.PHONY: amd64-build-healthcheck
amd64-build-healthcheck: install-tools golint
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./build/linux/amd64/healthcheck ./cmd/healthcheck

.PHONY: docker-build-arm
docker-build-arm: arm64-build arm64-build-healthcheck
	docker buildx build --platform linux/arm64 --build-arg BUILDMODE=copy --load -t $(DOCKER_NAMESPACE)/$(COMPONENT):$(VERSION) -f ./cmd/$(COMPONENT)/Dockerfile .

.PHONY: arm64-build-healthcheck
arm64-build-healthcheck: install-tools golint
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o ./build/linux/arm64/healthcheck ./cmd/healthcheck

.PHONY: windows-build-healthcheck
windows-build-healthcheck: install-tools golint
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./build/windows/amd64/healthcheck ./cmd/healthcheck

.PHONY: docker-push
docker-push:
	docker push $(DOCKER_NAMESPACE)/$(COMPONENT):$(VERSION)

.PHONY: docker-run
docker-run:
	docker run --rm -p 4317:4317 -p 55679:55679 -p 8889:8888 \
            -v "${PWD}/config.yaml":/otel-local-config.yaml \
            --name awscollector $(DOCKER_NAMESPACE)/$(COMPONENT):$(VERSION) \
            --config otel-local-config.yaml; \

.PHONY: docker-compose
docker-compose:
	cd examples/docker; docker-compose up -d

.PHONY: docker-stop
docker-stop:
	docker stop $(shell docker ps -aq)

.PHONY: gotest
gotest:
	@$(MAKE) for-all-target TARGET="test"

.PHONY: lint-sh
lint-sh:
	SHELLCHECK_OPTS=$(SHELLCHECK_OPTS) shellcheck  ${ALL_SHELL_SCRIPTS}

.PHONY: test-all
test-all: gotest lint-sh

.PHONY: gofmt
gofmt:
	@$(MAKE) for-all-target TARGET="fmt"


.PHONY: fmt-sh
fmt-sh: $(SHFMT)
	$(SHFMT) -w -d -i 4 .

.PHONY: lint-static-check
lint-static-check:
	@STATIC_CHECK_OUT=`$(TOOLS_BIN_DIR)/staticcheck $(ALL_PKGS) 2>&1`; \
		if [ "$$STATIC_CHECK_OUT" ]; then \
			echo "$(STATIC_CHECK) FAILED => static check errors:\n"; \
			echo "$$STATIC_CHECK_OUT\n"; \
			exit 1; \
		else \
			echo "Static check finished successfully"; \
		fi

.PHONY: golint
golint: lint-static-check
	@$(MAKE) for-all-target TARGET="lint"

.PHONY: gomod-tidy
gomod-tidy:
	@$(MAKE) for-all-target TARGET="mod-tidy"

.PHONY: gomod-vendor
gomod-vendor:
	go mod vendor

.PHONY: install-tools
install-tools:
	cd $(TOOLS_MOD_DIR) && GOBIN=$(TOOLS_BIN_DIR) go install golang.org/x/tools/cmd/goimports
	cd $(TOOLS_MOD_DIR) && GOBIN=$(TOOLS_BIN_DIR) go install honnef.co/go/tools/cmd/staticcheck
	cd $(TOOLS_MOD_DIR) && GOBIN=$(TOOLS_BIN_DIR) go install github.com/golangci/golangci-lint/cmd/golangci-lint
	cd $(TOOLS_MOD_DIR) && GOBIN=$(TOOLS_BIN_DIR) go install mvdan.cc/sh/v3/cmd/shfmt
	cd $(TOOLS_MOD_DIR) && GOBIN=$(TOOLS_BIN_DIR) go install go.opentelemetry.io/build-tools/dbotconf

.PHONY: install-dbotconf
install-dbotconf:
	if [ ! -f "$(DBOTCONF)" ]; then \
		cd $(TOOLS_MOD_DIR) go install go.opentelemetry.io/build-tools/dbotconf; \
	fi

.PHONY: clean
clean:
	rm -rf ./build
