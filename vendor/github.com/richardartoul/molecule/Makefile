gen-proto:
	rm -rf ./src/proto/gen
	# Requires to install protoc and google.golang.org/protobuf/cmd/protoc-gen-go manually.
	protoc --proto_path=./src/proto --go_out=src/proto src/proto/simple.proto

test:
	go test ./...