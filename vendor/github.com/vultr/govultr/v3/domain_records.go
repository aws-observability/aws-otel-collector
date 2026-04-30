package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// DomainRecordService is the interface to interact with the DNS Records endpoints on the Vultr API
// Link: https://www.vultr.com/api/#tag/dns
type DomainRecordService interface {
	Create(ctx context.Context, domain string, domainRecordCreateReq *DomainRecordCreateReq) (*DomainRecord, *http.Response, error)
	Get(ctx context.Context, domain, recordID string) (*DomainRecord, *http.Response, error)
	Update(ctx context.Context, domain, recordID string, domainRecordUpdateReq *DomainRecordUpdateReq) error
	Delete(ctx context.Context, domain, recordID string) error
	List(ctx context.Context, domain string, options *ListOptions) ([]DomainRecord, *Meta, *http.Response, error)
}

// DomainRecordsServiceHandler handles interaction with the DNS Records methods for the Vultr API
type DomainRecordsServiceHandler struct {
	client *Client
}

// DomainRecord represents a DNS record on Vultr
type DomainRecord struct {
	ID       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Name     string `json:"name,omitempty"`
	Data     string `json:"data,omitempty"`
	Priority int    `json:"priority,omitempty"`
	TTL      int    `json:"ttl,omitempty"`
}

// DomainRecordCreateReq struct to use for create domain record calls.
type DomainRecordCreateReq struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Data     string `json:"data"`
	TTL      int    `json:"ttl,omitempty"`
	Priority *int   `json:"priority,omitempty"`
}

// DomainRecordUpdateReq struct to use for update domain record calls.
type DomainRecordUpdateReq struct {
	Name     *string `json:"name,omitempty"`
	Type     string  `json:"type,omitempty"`
	Data     string  `json:"data,omitempty"`
	TTL      int     `json:"ttl,omitempty"`
	Priority *int    `json:"priority,omitempty"`
}

type domainRecordsBase struct {
	Records []DomainRecord `json:"records,omitempty"`
	Meta    *Meta          `json:"meta,omitempty"`
}

type domainRecordBase struct {
	Record *DomainRecord `json:"record,omitempty"`
}

// Create will add a DNS record.
func (d *DomainRecordsServiceHandler) Create(ctx context.Context, domain string, domainRecordCreateReq *DomainRecordCreateReq) (*DomainRecord, *http.Response, error) { //nolint:lll
	req, err := d.client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("%s/%s/records", domainPath, domain), domainRecordCreateReq)
	if err != nil {
		return nil, nil, err
	}

	record := new(domainRecordBase)
	resp, err := d.client.DoWithContext(ctx, req, record)
	if err != nil {
		return nil, resp, err
	}

	return record.Record, resp, nil
}

// Get record from a domain
func (d *DomainRecordsServiceHandler) Get(ctx context.Context, domain, recordID string) (*DomainRecord, *http.Response, error) {
	req, err := d.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/records/%s", domainPath, domain, recordID), nil)
	if err != nil {
		return nil, nil, err
	}

	record := new(domainRecordBase)
	resp, err := d.client.DoWithContext(ctx, req, record)
	if err != nil {
		return nil, resp, err
	}

	return record.Record, resp, nil
}

// Update will update a Domain record
func (d *DomainRecordsServiceHandler) Update(ctx context.Context, domain, recordID string, domainRecordUpdateReq *DomainRecordUpdateReq) error { //nolint:lll
	uri := fmt.Sprintf("%s/%s/records/%s", domainPath, domain, recordID)

	req, err := d.client.NewRequest(ctx, http.MethodPatch, uri, domainRecordUpdateReq)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// Delete will delete a domain name and all associated records.
func (d *DomainRecordsServiceHandler) Delete(ctx context.Context, domain, recordID string) error {
	req, err := d.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s/records/%s", domainPath, domain, recordID), nil)
	if err != nil {
		return err
	}
	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// List will list all the records associated with a particular domain on Vultr.
func (d *DomainRecordsServiceHandler) List(ctx context.Context, domain string, options *ListOptions) ([]DomainRecord, *Meta, *http.Response, error) { //nolint:lll
	req, err := d.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/records", domainPath, domain), nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}
	req.URL.RawQuery = newValues.Encode()

	records := new(domainRecordsBase)
	resp, err := d.client.DoWithContext(ctx, req, records)
	if err != nil {
		return nil, nil, resp, err
	}

	return records.Records, records.Meta, resp, nil
}
