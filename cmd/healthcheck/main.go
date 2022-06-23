package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	errNoHostProvided = errors.New("bad config: endpoint must be specified")
	errInvalidPath    = errors.New("bad config: path must start with /")
	errInvalidPort    = errors.New("bad config: invalid port")
)

// Validate checks if the extension configuration is valid
func Validate(host string, port string, path string) error {
	if host == "" {
		return errNoHostProvided
	}
	if portInt, err := strconv.Atoi(port); err == nil {
		if portInt < 1 || portInt > 65536 {
			return errInvalidPort
		}
	} else {
		return errInvalidPath
	}
	if !strings.HasPrefix(path, "/") {
		return errInvalidPath
	}
	return nil
}

func main() {
	usedHost := "127.0.0.1"      // default host
	usedPort := "13133"          // default port
	usedPath := "/health/status" //default path
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	host := generateCmd.String("host", usedHost, "Specify collector health-check host")
	port := generateCmd.String("port", usedPort, "Specify collector health-check port")
	path := generateCmd.String("path", usedPath, "Specify collector health-check path")

	validationErr := Validate(*host, *port, *path)
	if validationErr != nil {
		log.Panic(validationErr)
	}
	if len(os.Args) > 1 {
		err := generateCmd.Parse(os.Args[1:])
		if err != nil {
			log.Panic(err)
		}
	}

	resp, err := http.Get(fmt.Sprint("http://", *host, ":", *port, *path))
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
