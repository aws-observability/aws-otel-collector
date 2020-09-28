-include .env

AOC_IMPORT_PATH=aws-observability.io/collector
VERSION := $(shell cat VERSION)
PROJECTNAME := $(shell basename "$(PWD)")

GIT_SHA=$(shell git rev-parse HEAD)
DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
# All source code excluding any third party code and excluding the testbed.
# This is the code that we want to run tests for and lint, staticcheck, etc.
ALL_SRC := $(shell find . -name '*.go' \
							-not -path './testbed/*' \
							-not -path '*/third_party/*' \
							-not -path './.circleci/scripts/reportgenerator/*' \
							-not -path './pkg/devexporter/*' \
							-type f | sort)
# ALL_PKGS is the list of all packages where ALL_SRC files reside.
ALL_PKGS := $(shell go list $(sort $(dir $(ALL_SRC))))

GOTEST_OPT?= -short -coverprofile coverage.txt -v -race -timeout 180s
GOTEST=go test

BUILD_INFO_IMPORT_PATH=$(AOC_IMPORT_PATH)/tools/version

GOBUILD=GO111MODULE=on CGO_ENABLED=0 installsuffix=cgo go build -trimpath

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-s -w -X $(BUILD_INFO_IMPORT_PATH).GitHash=$(GIT_SHA) -X $(BUILD_INFO_IMPORT_PATH).Version=$(VERSION) -X $(BUILD_INFO_IMPORT_PATH).Date=$(DATE)"
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
DOCKER_NAMESPACE=amazon
COMPONENT=awscollector
LINT=golangci-lint
STATIC_CHECK=staticcheck

.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./build/darwin/aoc_darwin_amd64 ./cmd/awscollector
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./build/linux/aoc_linux_x86_64 ./cmd/awscollector
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o ./build/linux/aoc_linux_aarch64 ./cmd/awscollector
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./build/windows/aoc_windows_amd64 ./cmd/awscollector

.PHONY: awscollector
awscollector:
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./bin/awscollector_$(GOOS)_$(GOARCH) ./cmd/awscollector
	GOOS=windows GOARCH=amd64 EXTENSION=.exe $(GOBUILD) $(LDFLAGS) -o ./bin/windows/aoc_windows_amd64.exe ./cmd/awscollector

.PHONY: package-rpm
package-rpm: build
	ARCH=x86_64 DEST=build/packages/linux/amd64 tools/packaging/linux/create_rpm.sh
	ARCH=aarch64 DEST=build/packages/linux/arm64 tools/packaging/linux/create_rpm.sh

.PHONY: package-deb
package-deb: build
	ARCH=amd64 TARGET_SUPPORTED_ARCH=x86_64 DEST=build/packages/debian/amd64 tools/packaging/debian/create_deb.sh
	ARCH=arm64 TARGET_SUPPORTED_ARCH=aarch64 DEST=build/packages/debian/arm64 tools/packaging/debian/create_deb.sh

.PHONY: docker-component # Not intended to be used directly
docker-component: check-component
	GOOS=linux $(MAKE) $(COMPONENT)
	cp ./bin/$(COMPONENT)_$(GOOS)_$(GOARCH) ./cmd/$(COMPONENT)/$(COMPONENT)
	cp ./config.yaml ./cmd/$(COMPONENT)/
	docker build -t $(DOCKER_NAMESPACE)/$(COMPONENT):$(VERSION) ./cmd/$(COMPONENT)/
	rm ./cmd/$(COMPONENT)/$(COMPONENT)
	rm ./cmd/$(COMPONENT)/config.yaml

.PHONY: check-component
check-component:
ifndef COMPONENT
	$(error COMPONENT variable was not defined)
endif

.PHONY: docker-build
docker-build: awscollector
	$(MAKE) docker-component

.PHONY: docker-push
docker-push:
	docker push $(DOCKER_NAMESPACE)/$(COMPONENT):$(VERSION)

.PHONY: docker-run
docker-run:
	docker run --rm -p 55680:55680 -p 55679:55679 -p 8889:8888 \
            -v "${PWD}/config.yaml":/otel-local-config.yaml \
            --name awscollector $(DOCKER_NAMESPACE)/$(COMPONENT):$(VERSION) \
            --config otel-local-config.yaml; \

.PHONY: docker-compose
docker-compose:
	cd examples; docker-compose up -d

.PHONY: docker-stop
docker-stop:
	docker stop $(shell docker ps -aq)

.PHONY: test
test:
	echo $(ALL_PKGS) | xargs -n 10 $(GOTEST) $(GOTEST_OPT)

.PHONY: lint-static-check
lint-static-check:
	@STATIC_CHECK_OUT=`$(STATIC_CHECK) $(ALL_PKGS) 2>&1`; \
		if [ "$$STATIC_CHECK_OUT" ]; then \
			echo "$(STATIC_CHECK) FAILED => static check errors:\n"; \
			echo "$$STATIC_CHECK_OUT\n"; \
			exit 1; \
		else \
			echo "Static check finished successfully"; \
		fi

.PHONY: lint
lint: lint-static-check
	$(LINT) run --timeout 5m

.PHONY: install-tools
install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	go install github.com/jstemmer/go-junit-report
	go install honnef.co/go/tools/cmd/staticcheck

.PHONY: clean
clean:
	rm -rf ./build
