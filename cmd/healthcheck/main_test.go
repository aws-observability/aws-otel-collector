package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"
)

func TestHealthStatusHealthy(t *testing.T) {
	if os.Getenv("FLAG") == "1" {
		server := setUpMockCollector("127.0.0.1:13133", http.StatusOK)
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

	assert.Contains(t, got[:len(got)-1], expectedErrorString,
		fmt.Sprintf("Unexpected log message. Got %s but should contain %s", got, expectedErrorString))

	err := cmd.Wait()
	e, _ := err.(*exec.ExitError)
	assert.Nil(t, e)

}

func TestHealthStatusUnhealthy(t *testing.T) {
	if os.Getenv("FLAG") == "1" {
		server := setUpMockCollector("127.0.0.1:13133", http.StatusInternalServerError)
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

	assert.Contains(t, got[:len(got)-1], expectedErrorString,
		fmt.Sprintf("Unexpected log message. Got %s but should contain %s", got, expectedErrorString))

	err := cmd.Wait()
	expectedErrorStatus := "exit status 1"
	e, ok := err.(*exec.ExitError)
	assert.Equal(t, true, ok)
	assert.Equal(t, expectedErrorStatus, e.Error())

}
func TestHealthStatusServerDown(t *testing.T) {
	if os.Getenv("FLAG") == "1" {
		server := setUpMockCollector("127.0.0.1:13132", http.StatusInternalServerError)
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

	assert.Contains(t, got[:len(got)-1], expectedErrorString,
		fmt.Sprintf("Unexpected log message. Got %s but should contain %s", got, expectedErrorString))

	err := cmd.Wait()
	expectedErrorStatus := "exit status 1"
	e, ok := err.(*exec.ExitError)
	assert.Equal(t, true, ok)
	assert.Equal(t, expectedErrorStatus, e.Error())

}

func TestValidatePortWithWrongString(t *testing.T) {
	expectedErrorString := "invalid port"
	actual := validatePort("wrongPort")

	assert.Contains(t, actual.Error(), expectedErrorString,
		fmt.Sprintf("Unexpected log message. Got %s but should contain %s", actual, expectedErrorString))
}

func TestValidatePortWithWrongPort(t *testing.T) {
	expectedErrorString := "port outside of range"
	actual := validatePort("65536")

	assert.Contains(t, actual.Error(), expectedErrorString,
		fmt.Sprintf("Unexpected log message. Got %s but should contain %s", actual, expectedErrorString))
}

func TestValidatePortWithValidPort(t *testing.T) {
	actual := validatePort("65535")

	assert.Nil(t, actual)
}

func setUpMockCollector(healthCheckDefaultEndpoint string, statusCode int) *httptest.Server {
	l, _ := net.Listen("tcp", healthCheckDefaultEndpoint)
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(statusCode)
	}))
	server.Listener.Close()
	server.Listener = l
	server.Start()
	return server
}
