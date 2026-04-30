package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// SSHKeyService is the interface to interact with the SSH Key endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/ssh
type SSHKeyService interface { //nolint:dupl
	Create(ctx context.Context, sshKeyReq *SSHKeyReq) (*SSHKey, *http.Response, error)
	Get(ctx context.Context, sshKeyID string) (*SSHKey, *http.Response, error)
	Update(ctx context.Context, sshKeyID string, sshKeyReq *SSHKeyReq) error
	Delete(ctx context.Context, sshKeyID string) error
	List(ctx context.Context, options *ListOptions) ([]SSHKey, *Meta, *http.Response, error)
}

// SSHKeyServiceHandler handles interaction with the SSH Key methods for the Vultr API
type SSHKeyServiceHandler struct {
	client *Client
}

// SSHKey represents an SSH Key on Vultr
type SSHKey struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	SSHKey      string `json:"ssh_key"`
	DateCreated string `json:"date_created"`
}

// SSHKeyReq is the ssh key struct for create and update calls
type SSHKeyReq struct {
	Name   string `json:"name,omitempty"`
	SSHKey string `json:"ssh_key,omitempty"`
}

type sshKeysBase struct {
	SSHKeys []SSHKey `json:"ssh_keys"`
	Meta    *Meta    `json:"meta"`
}

type sshKeyBase struct {
	SSHKey *SSHKey `json:"ssh_key"`
}

// Create a ssh key
func (s *SSHKeyServiceHandler) Create(ctx context.Context, sshKeyReq *SSHKeyReq) (*SSHKey, *http.Response, error) {
	uri := "/v2/ssh-keys"

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, sshKeyReq)
	if err != nil {
		return nil, nil, err
	}

	key := new(sshKeyBase)
	resp, err := s.client.DoWithContext(ctx, req, key)
	if err != nil {
		return nil, resp, err
	}

	return key.SSHKey, resp, nil
}

// Get a specific ssh key.
func (s *SSHKeyServiceHandler) Get(ctx context.Context, sshKeyID string) (*SSHKey, *http.Response, error) {
	uri := fmt.Sprintf("/v2/ssh-keys/%s", sshKeyID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	sshKey := new(sshKeyBase)
	resp, err := s.client.DoWithContext(ctx, req, sshKey)
	if err != nil {
		return nil, resp, err
	}

	return sshKey.SSHKey, resp, nil
}

// Update will update the given SSH Key. Empty strings will be ignored.
func (s *SSHKeyServiceHandler) Update(ctx context.Context, sshKeyID string, sshKeyReq *SSHKeyReq) error {
	uri := fmt.Sprintf("/v2/ssh-keys/%s", sshKeyID)

	req, err := s.client.NewRequest(ctx, http.MethodPatch, uri, sshKeyReq)
	if err != nil {
		return err
	}

	_, err = s.client.DoWithContext(ctx, req, nil)
	return err
}

// Delete a specific ssh-key.
func (s *SSHKeyServiceHandler) Delete(ctx context.Context, sshKeyID string) error {
	uri := fmt.Sprintf("/v2/ssh-keys/%s", sshKeyID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	_, err = s.client.DoWithContext(ctx, req, nil)
	return err
}

// List all available SSH Keys.
func (s *SSHKeyServiceHandler) List(ctx context.Context, options *ListOptions) ([]SSHKey, *Meta, *http.Response, error) { //nolint:dupl
	uri := "/v2/ssh-keys"

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	sshKeys := new(sshKeysBase)
	resp, err := s.client.DoWithContext(ctx, req, sshKeys)
	if err != nil {
		return nil, nil, resp, err
	}

	return sshKeys.SSHKeys, sshKeys.Meta, resp, nil
}
