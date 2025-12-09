# Function to execute a command.
# Accepts command to execute as first parameter.
define exec-command
$(1)

endef

# Find all .proto files.
TEF_PROTO_FILES := $(wildcard ./proto/*.proto)

.PHONY: default
default: test

.PHONY: all
all: genproto test

.PHONY: genproto
genproto:
	$(foreach file,$(TEF_PROTO_FILES),$(call exec-command,protoc -I/usr/local/include -I ./proto $(file) --go_out=./ --go-grpc_out=./ ))

.PHONY: test
test:
	go test -v ./...
