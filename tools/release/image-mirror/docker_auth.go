package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/docker/docker/api/types/registry"
)

func getDockerCredentials(authToken string) (string, error) {
	credentials, err := base64.StdEncoding.DecodeString(authToken)
	if err != nil {
		return "", err
	}
	parts := strings.SplitN(string(credentials), ":", 2)
	if len(parts) != 2 {
		return "", errors.New("unable to split authentication token into username/password")
	}
	authConfig := registry.AuthConfig{
		Username: parts[0],
		Password: parts[1],
	}
	encoded, err := json.Marshal(authConfig)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(encoded), nil
}
