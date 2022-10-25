package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHealthStatusHealthy(t *testing.T) {
	server := setUpMockCollector(t, "127.0.0.1:13133", http.StatusOK)
	os.Args = []string{"-port=13133"}
	defer server.Close()
	port := "13133"
	got, err := executeHealthCheck("127.0.0.1", &port, "/")

	expectedErrorString := "STATUS: 200"

	assert.Contains(t, got, expectedErrorString,
		fmt.Sprintf("Unexpected log message. Got %s but should contain %s", got, expectedErrorString))

	assert.NoError(t, err)

}

func TestHealthStatusUnhealthy(t *testing.T) {
	server := setUpMockCollector(t, "127.0.0.1:13133", http.StatusInternalServerError)
	os.Args = []string{"-port=13133"}
	defer server.Close()
	port := "13133"
	got, err := executeHealthCheck("127.0.0.1", &port, "/")

	expectedErrorString := "STATUS: 500"

	assert.Contains(t, err.Error(), expectedErrorString,
		fmt.Sprintf("Unexpected log message. Got %s but should contain %s", got, expectedErrorString))

}
func TestHealthStatusServerDown(t *testing.T) {
	server := setUpMockCollector(t, "127.0.0.1:13132", http.StatusInternalServerError)
	os.Args = []string{"-port=13133"}
	defer server.Close()
	port := "13133"
	got, err := executeHealthCheck("127.0.0.1", &port, "/")

	expectedErrorString := "unable to retrieve health status"

	assert.Contains(t, err.Error(), expectedErrorString,
		fmt.Sprintf("Unexpected log message. Got %s but should contain %s", got, expectedErrorString))

}

func setUpMockCollector(t *testing.T, healthCheckDefaultEndpoint string, statusCode int) *httptest.Server {
	l, err := net.Listen("tcp", healthCheckDefaultEndpoint)
	require.NoError(t, err)
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(statusCode)
	}))
	server.Listener.Close()
	server.Listener = l
	server.Start()
	return server
}

func TestValidatePort(t *testing.T) {
	testCases := []struct {
		name          string
		port          string
		errorExpected bool
	}{
		{
			name:          "WrongString",
			port:          "StringPort",
			errorExpected: true,
		},
		{
			name:          "EmptyString",
			port:          "",
			errorExpected: true,
		},
		{
			name:          "WrongPort",
			port:          "65536",
			errorExpected: true,
		},
		{
			name:          "ValidPort",
			port:          "13133",
			errorExpected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validatePort(tc.port)
			if tc.errorExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
