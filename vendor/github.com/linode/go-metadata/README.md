# Linode Metadata Service Client for Go

[![Release](https://img.shields.io/github/v/release/linode/go-metadata)](https://github.com/linode/go-metadata/releases/latest)
[![GoDoc](https://godoc.org/github.com/linode/go-metadata?status.svg)](https://godoc.org/github.com/linode/go-metadata)

This package allows Go projects to easily interact with the [Linode Metadata Service](https://www.linode.com/docs/products/compute/compute-instances/guides/metadata/?tabs=linode-api).

## Getting Started

### Prerequisites 

- Go >= 1.20
- A running [Linode Instance](https://www.linode.com/docs/api/linode-instances/)

### Installation

```bash
go get github.com/linode/go-metadata
```

### Basic Example

The following sample shows a simple Go project that initializes a new metadata client and retrieves various information
about the current Linode.

```go
package main

import (
	"context"
	"fmt"
	"log"

	metadata "github.com/linode/go-metadata"
)

func main() {
	// Create a new client
	client, err := metadata.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve metadata about the current instance from the metadata API
	instanceInfo, err := client.GetInstance(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Instance Label:", instanceInfo.Label)
}
```

### Without Token Management

By default, metadata API tokens are automatically generated and refreshed without any user intervention.
If you would like to manage API tokens yourself, this functionality can be disabled:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	metadata "github.com/linode/go-metadata"
)

func main() {
	// Get a token from the environment
	token := os.Getenv("LINODE_METADATA_TOKEN")

	// Create a new client
	client, err := metadata.NewClient(
		context.Background(), 
		metadata.ClientWithoutManagedToken(), 
		metadata.ClientWithToken(token),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve metadata about the current instance from the metadata API
	instanceInfo, err := client.GetInstance(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Instance Label:", instanceInfo.Label)
}
```

For more examples, visit the [examples directory](./examples).

## Documentation

See [godoc](https://pkg.go.dev/github.com/linode/go-metadata) for a complete documentation reference.

## Testing

Before running tests on this project, please ensure you have a 
[Linode Personal Access Token](https://www.linode.com/docs/products/tools/api/guides/manage-api-tokens/)
exported under the `LINODE_TOKEN` environment variable.

### End-to-End Testing Using Ansible

This project contains an Ansible playbook to automatically deploy the necessary infrastructure 
and run end-to-end tests on it.

To install the dependencies for this playbook, ensure you have Python 3 installed and run the following:

```bash
make test-deps
```

After all dependencies have been installed, you can run the end-to-end test suite by running the following:

```bash
make e2e
```

If your local SSH public key is stored in a location other than `~/.ssh/id_rsa.pub`, 
you may need to override the `TEST_PUBKEY` argument:

```bash
make TEST_PUBKEY=/path/to/my/pubkey e2e
```

**NOTE: To speed up subsequent test runs, the infrastructure provisioned for testing will persist after the test run is complete. 
This infrastructure is safe to manually remove.**

### Manual End-to-End Testing

End-to-end tests can also be manually run using the `make e2e-local` target.
This test suite is expected to run from within a Linode instance and will likely 
fail in other environments.

## Contribution Guidelines

Want to improve metadata-go? Please start [here](CONTRIBUTING.md).

## License

This software is Copyright Akamai Technologies, Inc. and is released under the [Apache 2.0 license](./LICENSE).
