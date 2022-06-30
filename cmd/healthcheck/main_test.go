package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestHealthStatusHealthy(t *testing.T) {
	if os.Getenv("FLAG") == "1" {
		healthCheckDefaultUrl := "127.0.0.1:13133"
		l, _ := net.Listen("tcp", healthCheckDefaultUrl)
		server := httptest.NewUnstartedServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))
		server.Listener.Close()
		server.Listener = l
		server.Start()
		os.Args = []string{"-port=13133"}

		defer server.Close()
		main()
		return
	} // Run the test in a subprocess
	cmd := exec.Command(os.Args[0], "-test.run=TestHealthStatusHealthy") //nolint:gosec
	cmd.Env = append(os.Environ(), "FLAG=1")

	stdout, _ := cmd.StderrPipe()
	er := cmd.Start()
	if er != nil {
		t.Fatal("Unable to start the testing sub-process.")
	}

	gotBytes, _ := ioutil.ReadAll(stdout)
	got := string(gotBytes)
	expectedErrorString := "STATUS: 200"

	if !strings.Contains(got[:len(got)-1], expectedErrorString) {
		t.Fatalf("Unexpected log message. Got %s but should contain %s", got[:len(got)-1], expectedErrorString)
	}
	err := cmd.Wait()
	e, _ := err.(*exec.ExitError)
	assert.Nil(t, e)

}

func TestHealthStatusUnhealthy(t *testing.T) {
	if os.Getenv("FLAG") == "1" {
		healthCheckDefaultUrl := "127.0.0.1:13133"
		l, _ := net.Listen("tcp", healthCheckDefaultUrl)
		server := httptest.NewUnstartedServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		}))
		server.Listener.Close()
		server.Listener = l
		server.Start()
		os.Args = []string{"-port=13133"}

		defer server.Close()
		main()
		return
	} // Run the test in a subprocess
	cmd := exec.Command(os.Args[0], "-test.run=TestHealthStatusUnhealthy") //nolint:gosec
	cmd.Env = append(os.Environ(), "FLAG=1")

	stdout, _ := cmd.StderrPipe()
	er := cmd.Start()
	if er != nil {
		t.Fatal("Unable to start the testing sub-process.")
	}

	gotBytes, _ := ioutil.ReadAll(stdout)
	got := string(gotBytes)
	expectedErrorString := "STATUS: 500"

	if !strings.Contains(got[:len(got)-1], expectedErrorString) {
		t.Fatalf("Unexpected log message. Got %s but should contain %s", got[:len(got)-1], expectedErrorString)
	}
	err := cmd.Wait()
	expectedErrorStatus := "exit status 1"
	e, ok := err.(*exec.ExitError)
	assert.Equal(t, true, ok)
	assert.Equal(t, expectedErrorStatus, e.Error())

}
func TestHealthStatusServerDown(t *testing.T) {
	if os.Getenv("FLAG") == "1" {
		healthCheckDefaultUrl := "127.0.0.1:13132"
		l, _ := net.Listen("tcp", healthCheckDefaultUrl)
		server := httptest.NewUnstartedServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		}))
		server.Listener.Close()
		server.Listener = l
		server.Start()
		os.Args = []string{"-port=13133"}

		defer server.Close()
		main()
		return
	} // Run the test in a subprocess
	cmd := exec.Command(os.Args[0], "-test.run=TestHealthStatusServerDown") //nolint:gosec
	cmd.Env = append(os.Environ(), "FLAG=1")

	stdout, _ := cmd.StderrPipe()
	er := cmd.Start()
	if er != nil {
		t.Fatal("Unable to start the testing sub-process.")
	}

	gotBytes, _ := ioutil.ReadAll(stdout)
	got := string(gotBytes)
	expectedErrorString := "Unable to retrieve health status"

	if !strings.Contains(got[:len(got)-1], expectedErrorString) {
		t.Fatalf("Unexpected log message. Got %s but should contain %s", got[:len(got)-1], expectedErrorString)
	}
	err := cmd.Wait()
	expectedErrorStatus := "exit status 1"
	e, ok := err.(*exec.ExitError)
	assert.Equal(t, true, ok)
	assert.Equal(t, expectedErrorStatus, e.Error())

}
