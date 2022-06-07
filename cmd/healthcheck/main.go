package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello")
	_, err := http.Get("http://127.0.0.1:13133/health/status")
	if err != nil {
		os.Exit(1)
	}

}
