// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	tokenUrlEndpoint  = "%s/api/v2/delegated-token"
	authorizationType = "Delegated"

	contentTypeHeader   = "Content-Type"
	contextLengthHeader = "Content-Length"
	applicationJson     = "application/json"
)

func GetDelegatedToken(ctx context.Context, orgUUID, delegatedAuthProof string) (*DelegatedTokenCredentials, error) {
	url, err := GetDelegatedTokenUrl(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get URL for delegated token: %w", err)
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, err
	}
	req.Header.Set(contentTypeHeader, applicationJson)
	req.Header.Set(authorizationHeader, fmt.Sprintf("%s %s", authorizationType, delegatedAuthProof))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	tokenBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get token: %s", resp.Status)
	}

	creds, err := ParseDelegatedTokenResponse(tokenBytes, orgUUID, delegatedAuthProof)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token response: %w", err)
	}
	return creds, nil
}

func ParseDelegatedTokenResponse(tokenBytes []byte, orgUUID, delegatedAuthProof string) (*DelegatedTokenCredentials, error) {
	// Parse the response to get the token
	var tokenResponse map[string]interface{}
	err := json.Unmarshal(tokenBytes, &tokenResponse)
	if err != nil {
		return nil, err
	}

	// Get attributes from the response
	dataResponse, ok := tokenResponse["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to get data from response: %v", tokenResponse)
	}

	// Get the attributes from the data
	attributes, ok := dataResponse["attributes"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to get attributes from response: %v", tokenResponse)
	}

	// Get the access token from the response
	token, ok := attributes["access_token"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to get token from response: %v", tokenResponse)
	}

	// Get the expiration time from the response
	// Default to 15 minutes if the expiration time is not set
	expirationTime := time.Now().Add(time.Duration(15) * time.Minute)
	expirationStr, ok := attributes["expires"].(string)
	if ok {
		expirationInt, err := strconv.ParseInt(expirationStr, 10, 64)
		if err == nil {
			// Convert the expiration time to a time.Time object
			expirationTime = time.Unix(expirationInt, 0)
		}
	}

	return &DelegatedTokenCredentials{
		OrgUUID:        orgUUID,
		DelegatedToken: token,
		DelegatedProof: delegatedAuthProof,
		Expiration:     expirationTime,
	}, nil
}

func GetDelegatedTokenUrl(ctx context.Context) (string, error) {
	config := NewConfiguration()
	url, err := config.ServerURLWithContext(ctx, "")
	if err != nil {
		return "", err
	}
	url = fmt.Sprintf(tokenUrlEndpoint, url)
	return url, nil
}
