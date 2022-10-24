package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// validatePort checks if the port configuration is valid
func validatePort(port string) error {

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("invalid port: %w", err)
	}
	if portInt < 1 || portInt > 65535 {
		return fmt.Errorf("port outside of range [1;65535]: %d", portInt)
	}
	return nil

}

func main() {
	const host = "127.0.0.1" // default host
	usedPort := "13133"      // default port
	const path = "/"         //default path
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	port := generateCmd.String("port", usedPort, "Specify collector health-check port")

	if len(os.Args) > 1 {
		err := generateCmd.Parse(os.Args[1:])
		if err != nil {
			log.Fatalf("%s", err)
		}
	}

	validationErr := validatePort(*port)
	if validationErr != nil {
		log.Fatalf("%s", validationErr)
	}

	resp, err := http.Get(fmt.Sprint("http://", host, ":", *port, path))
	if err != nil {
		log.Fatalf("Unable to retrieve health status: %s", err.Error())
	}
	if resp != nil && resp.StatusCode == 200 {
		log.Printf("STATUS: %d", resp.StatusCode)
	} else {
		log.Fatalf("STATUS: %d", resp.StatusCode)
	}

}
