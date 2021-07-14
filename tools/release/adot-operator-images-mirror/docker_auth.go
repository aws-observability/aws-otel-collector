package main

import (
	"fmt"
	"log"

	docker "github.com/fsouza/go-dockerclient"
)

func getDockerCredentials(registry string) (*docker.AuthConfiguration, error) {
	authOptions, err := docker.NewAuthConfigurationsFromDockerCfg()
	if err != nil {
		log.Fatal(err)
	}

	creds, ok := authOptions.Configs[registry]
	if !ok {
		return nil, fmt.Errorf("No auth found for %s", registry)
	}

	return &creds, nil
}
