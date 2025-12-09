/*
go-metadata allows Go projects to easily interact with the Linode Metadata Service.

Basic example:
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

For more examples see: https://github.com/linode/go-metadata/tree/ref/watcher-refactor/examples

To learn more about the Linode Metadata Service, see the official guide: https://www.linode.com/docs/products/compute/compute-instances/guides/metadata/?tabs=linode-api
*/

package metadata
