package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// ISOService is the interface to interact with the ISO endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/iso
type ISOService interface {
	Create(ctx context.Context, isoReq *ISOReq) (*ISO, *http.Response, error)
	Get(ctx context.Context, isoID string) (*ISO, *http.Response, error)
	Delete(ctx context.Context, isoID string) error
	List(ctx context.Context, options *ListOptions) ([]ISO, *Meta, *http.Response, error)
	ListPublic(ctx context.Context, options *ListOptions) ([]PublicISO, *Meta, *http.Response, error)
}

// ISOServiceHandler handles interaction with the ISO methods for the Vultr API
type ISOServiceHandler struct {
	client *Client
}

// ISO represents ISOs currently available on this account.
type ISO struct {
	ID          string `json:"id"`
	DateCreated string `json:"date_created"`
	FileName    string `json:"filename"`
	Size        int    `json:"size,omitempty"`
	MD5Sum      string `json:"md5sum,omitempty"`
	SHA512Sum   string `json:"sha512sum,omitempty"`
	Status      string `json:"status"`
}

// PublicISO represents public ISOs offered in the Vultr ISO library.
type PublicISO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MD5Sum      string `json:"md5sum,omitempty"`
}

// ISOReq is used for creating ISOs.
type ISOReq struct {
	URL string `json:"url"`
}

type isosBase struct {
	ISOs []ISO `json:"isos"`
	Meta *Meta `json:"meta"`
}

type isoBase struct {
	ISO *ISO `json:"iso"`
}

type publicIsosBase struct {
	PublicIsos []PublicISO `json:"public_isos"`
	Meta       *Meta       `json:"meta"`
}

// Create will create a new ISO image on your account
func (i *ISOServiceHandler) Create(ctx context.Context, isoReq *ISOReq) (*ISO, *http.Response, error) {
	uri := "/v2/iso"

	req, err := i.client.NewRequest(ctx, http.MethodPost, uri, isoReq)
	if err != nil {
		return nil, nil, err
	}

	iso := new(isoBase)
	resp, err := i.client.DoWithContext(ctx, req, iso)
	if err != nil {
		return nil, resp, err
	}

	return iso.ISO, resp, nil
}

// Get an ISO
func (i *ISOServiceHandler) Get(ctx context.Context, isoID string) (*ISO, *http.Response, error) {
	uri := fmt.Sprintf("/v2/iso/%s", isoID)

	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	iso := new(isoBase)
	resp, err := i.client.DoWithContext(ctx, req, iso)
	if err != nil {
		return nil, resp, err
	}
	return iso.ISO, resp, nil
}

// Delete will delete an ISO image from your account
func (i *ISOServiceHandler) Delete(ctx context.Context, isoID string) error {
	uri := fmt.Sprintf("/v2/iso/%s", isoID)

	req, err := i.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = i.client.DoWithContext(ctx, req, nil)
	return err
}

// List will list all ISOs currently available on your account
func (i *ISOServiceHandler) List(ctx context.Context, options *ListOptions) ([]ISO, *Meta, *http.Response, error) { //nolint:dupl
	uri := "/v2/iso"

	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	iso := new(isosBase)
	resp, err := i.client.DoWithContext(ctx, req, iso)
	if err != nil {
		return nil, nil, resp, err
	}

	return iso.ISOs, iso.Meta, resp, nil
}

// ListPublic will list public ISOs offered in the Vultr ISO library.
func (i *ISOServiceHandler) ListPublic(ctx context.Context, options *ListOptions) ([]PublicISO, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := "/v2/iso-public"

	req, err := i.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	iso := new(publicIsosBase)
	resp, err := i.client.DoWithContext(ctx, req, iso)
	if err != nil {
		return nil, nil, resp, err
	}

	return iso.PublicIsos, iso.Meta, resp, nil
}
