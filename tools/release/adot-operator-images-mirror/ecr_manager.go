package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/cenkalti/backoff"
)

const (
	operatorRepo  = "adot-operator"
	rbacProxyRepo = "mirror-kube-rbac-proxy"
)

type ecrManager struct {
	client       *ecrpublic.Client
	repositories map[string]bool
	registryAuth string
}

func (e *ecrManager) exists(name string) bool {
	if _, ok := e.repositories[name]; ok {
		return true
	}

	return false
}

func (e *ecrManager) ensure(ctx context.Context, name string) error {
	if e.exists(name) {
		return nil
	}

	return e.create(ctx, name)
}

func (e *ecrManager) create(ctx context.Context, name string) error {
	if name != operatorRepo && name != rbacProxyRepo {
		return fmt.Errorf("wrong repository name: %s", name)
	}

	_, err := e.client.CreateRepository(ctx, &ecrpublic.CreateRepositoryInput{
		RepositoryName: &name,
	})
	if err != nil {
		return err
	}

	e.repositories[name] = true
	return nil
}

func (e *ecrManager) buildCache(ctx context.Context, nextToken *string) error {
	if nextToken == nil {
		log.Println("Loading the list of ECR repositories")
	}

	resp, err := e.client.DescribeRepositories(ctx, &ecrpublic.DescribeRepositoriesInput{
		NextToken: nextToken,
	})
	if err != nil {
		return err
	}

	if e.repositories == nil {
		e.repositories = make(map[string]bool)
	}

	for _, repo := range resp.Repositories {
		e.repositories[*repo.RepositoryName] = true
	}

	// keep paging as long as there is a token for the next page
	if resp.NextToken != nil {
		err = e.buildCache(ctx, resp.NextToken)
		if err != nil {
			return err
		}
	}

	// no next token means we hit the last page
	if nextToken == nil {
		log.Println("Done loading ECR repositories")
	}

	return nil
}

func (e *ecrManager) buildCacheBackoff(ctx context.Context) backoff.Operation {
	return func() error {
		return e.buildCache(ctx, nil)
	}
}

func (e *ecrManager) getAuthToken(ctx context.Context) (string, error) {
	input := &ecrpublic.GetAuthorizationTokenInput{}
	output, err := e.client.GetAuthorizationToken(ctx, input)
	if err != nil {
		return "", err
	}

	return *output.AuthorizationData.AuthorizationToken, nil
}
