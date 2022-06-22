package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	usedPort := "13133"          // default port
	usedHost := "127.0.0.1"      // default host
	usedPath := "/health/status" //default path
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	host := generateCmd.String("host", usedHost, "Specify collector health-check host")
	port := generateCmd.String("port", usedPort, "Specify collector health-check port")
	path := generateCmd.String("path", usedPath, "Specify collector health-check path")

	if len(os.Args) > 1 {
		err := generateCmd.Parse(os.Args[1:])
		if err != nil {
			log.Panic(err)
		}
	}

	resp, err := http.Get(fmt.Sprint("http://", *host, ":", *port, *path))

	if err != nil {
		log.Printf("ERROR")
		os.Exit(1)
	} else {
		log.Printf(`STATUS: ${res.statusCode}`)
		if resp != nil && resp.StatusCode == 200 {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

}
