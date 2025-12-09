# metadata [![Documentation](https://godoc.org/github.com/digitalocean/go-metadata?status.svg)](https://godoc.org/github.com/digitalocean/go-metadata)

A Go client to interact with the [DigitalOcean Metadata API](https://developers.digitalocean.com/documentation/metadata/).

# Usage

```go
// Create a client
client := metadata.NewClient(opts)

// Request all the metadata about the current droplet
all, err := client.Metadata()
if err != nil {
    log.Fatal(err)
}

// Lookup what our IPv4 address is on our first public
// network interface.
publicIPv4Addr := all.Interfaces["public"][0].IPv4.IPAddress

fmt.Println(publicIPv4Addr)
```

# License

MIT license
