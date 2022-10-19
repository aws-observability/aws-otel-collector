package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	errInvalidPort = errors.New("bad config: invalid port")
)

// validatePort checks if the port configuration is valid
func validatePort(port string) error {

	if portInt, err := strconv.Atoi(port); err == nil {
		if portInt < 1 || portInt > 65535 {
			return errInvalidPort
		}
	} else {
		return errInvalidPort
	}
	return nil
}

func main() {
	host := "127.0.0.1"      // default host
	usedPort := "13133"      // default port
	path := "/health/status" //default path
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	port := generateCmd.String("port", usedPort, "Specify collector health-check port")

	if len(os.Args) > 1 {
		err := generateCmd.Parse(os.Args[1:])
		if err != nil {
			log.Panic(err)
		}
	}

	validationErr := validatePort(*port)
	if validationErr != nil {
		log.Panic(validationErr)
	}
	
	resp, err := http.Get(fmt.Sprint("http://", host, ":", *port, path))
	if err != nil {
		log.Fatalf("Unable to retrieve health status: %s", err.Error())
	} else {
		if resp != nil && resp.StatusCode == 200 {
			log.Printf("STATUS: %d", resp.StatusCode)
			os.Exit(0)
		} else {
			log.Fatalf("STATUS: %d", resp.StatusCode)
		}
	}
}
