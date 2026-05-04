package govultr

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

// OSService is the interface to interact with the operating system endpoint on the Vultr API
// Link : https://www.vultr.com/api/#tag/os
type OSService interface {
	List(ctx context.Context, options *ListOptions) ([]OS, *Meta, *http.Response, error)
}

// OSServiceHandler handles interaction with the operating system methods for the Vultr API
type OSServiceHandler struct {
	client *Client
}

// OS represents a Vultr operating system
type OS struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Arch   string `json:"arch"`
	Family string `json:"family"`
}

type osBase struct {
	OS   []OS  `json:"os"`
	Meta *Meta `json:"meta"`
}

var _ OSService = &OSServiceHandler{}

// List retrieves a list of available operating systems.
func (o *OSServiceHandler) List(ctx context.Context, options *ListOptions) ([]OS, *Meta, *http.Response, error) { //nolint:dupl
	uri := "/v2/os"
	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}
	req.URL.RawQuery = newValues.Encode()

	os := new(osBase)
	resp, err := o.client.DoWithContext(ctx, req, os)
	if err != nil {
		return nil, nil, resp, err
	}

	return os.OS, os.Meta, resp, nil
}
