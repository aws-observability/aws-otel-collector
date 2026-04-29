package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const domainPath = "/v2/domains"

// DomainService is the interface to interact with the DNS endpoints on the Vultr API
// https://www.vultr.com/api/#tag/dns
type DomainService interface {
	Create(ctx context.Context, domainReq *DomainReq) (*Domain, *http.Response, error)
	Get(ctx context.Context, domain string) (*Domain, *http.Response, error)
	Update(ctx context.Context, domain, dnsSec string) error
	Delete(ctx context.Context, domain string) error
	List(ctx context.Context, options *ListOptions) ([]Domain, *Meta, *http.Response, error)

	GetSoa(ctx context.Context, domain string) (*Soa, *http.Response, error)
	UpdateSoa(ctx context.Context, domain string, soaReq *Soa) error

	GetDNSSec(ctx context.Context, domain string) ([]string, *http.Response, error)
}

// DomainServiceHandler handles interaction with the DNS methods for the Vultr API
type DomainServiceHandler struct {
	client *Client
}

// Domain represents a Domain entry on Vultr
type Domain struct {
	Domain      string `json:"domain,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
	DNSSec      string `json:"dns_sec,omitempty"`
}

// DomainReq is the struct to create a domain
// If IP is omitted then an empty DNS entry will be created. If supplied the domain will be pre populated with entries
type DomainReq struct {
	Domain string `json:"domain,omitempty"`
	IP     string `json:"ip,omitempty"`
	DNSSec string `json:"dns_sec,omitempty"`
}

type domainsBase struct {
	Domains []Domain `json:"domains"`
	Meta    *Meta    `json:"meta"`
}

type domainBase struct {
	Domain *Domain `json:"domain"`
}

// Soa information for the Domain
type Soa struct {
	NSPrimary string `json:"nsprimary,omitempty"`
	Email     string `json:"email,omitempty"`
}

type soaBase struct {
	DNSSoa *Soa `json:"dns_soa,omitempty"`
}

type dnsSecBase struct {
	DNSSec []string `json:"dns_sec,omitempty"`
}

// Create a domain entry
func (d *DomainServiceHandler) Create(ctx context.Context, domainReq *DomainReq) (*Domain, *http.Response, error) {
	req, err := d.client.NewRequest(ctx, http.MethodPost, domainPath, domainReq)
	if err != nil {
		return nil, nil, err
	}

	domain := new(domainBase)
	resp, err := d.client.DoWithContext(ctx, req, domain)
	if err != nil {
		return nil, resp, err
	}

	return domain.Domain, resp, nil
}

// Get a domain from your Vultr account.
func (d *DomainServiceHandler) Get(ctx context.Context, domain string) (*Domain, *http.Response, error) {
	req, err := d.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s", domainPath, domain), nil)
	if err != nil {
		return nil, nil, err
	}

	dBase := new(domainBase)
	resp, err := d.client.DoWithContext(ctx, req, dBase)
	if err != nil {
		return nil, resp, err
	}

	return dBase.Domain, resp, nil
}

// Update allows you to enable or disable DNS Sec on the domain.
// The two valid options for dnsSec are "enabled" or "disabled"
func (d *DomainServiceHandler) Update(ctx context.Context, domain, dnsSec string) error {
	body := &RequestBody{"dns_sec": dnsSec}
	req, err := d.client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("%s/%s", domainPath, domain), body)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// Delete a domain with all associated records.
func (d *DomainServiceHandler) Delete(ctx context.Context, domain string) error {
	req, err := d.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s", domainPath, domain), nil)
	if err != nil {
		return err
	}
	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// List gets all domains associated with the current Vultr account.
func (d *DomainServiceHandler) List(ctx context.Context, options *ListOptions) ([]Domain, *Meta, *http.Response, error) { //nolint:dupl
	req, err := d.client.NewRequest(ctx, http.MethodGet, domainPath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	domains := new(domainsBase)
	resp, err := d.client.DoWithContext(ctx, req, domains)
	if err != nil {
		return nil, nil, resp, err
	}

	return domains.Domains, domains.Meta, resp, nil
}

// GetSoa gets the SOA record information for a domain
func (d *DomainServiceHandler) GetSoa(ctx context.Context, domain string) (*Soa, *http.Response, error) {
	req, err := d.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/soa", domainPath, domain), nil)
	if err != nil {
		return nil, nil, err
	}

	soa := new(soaBase)
	resp, err := d.client.DoWithContext(ctx, req, soa)
	if err != nil {
		return nil, resp, err
	}

	return soa.DNSSoa, resp, nil
}

// UpdateSoa will update the SOA record information for a domain.
func (d *DomainServiceHandler) UpdateSoa(ctx context.Context, domain string, soaReq *Soa) error {
	req, err := d.client.NewRequest(ctx, http.MethodPatch, fmt.Sprintf("%s/%s/soa", domainPath, domain), soaReq)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// GetDNSSec gets the DNSSec keys for a domain (if enabled)
func (d *DomainServiceHandler) GetDNSSec(ctx context.Context, domain string) ([]string, *http.Response, error) {
	req, err := d.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/dnssec", domainPath, domain), nil)
	if err != nil {
		return nil, nil, err
	}

	dnsSec := new(dnsSecBase)
	resp, err := d.client.DoWithContext(ctx, req, dnsSec)
	if err != nil {
		return nil, resp, err
	}

	return dnsSec.DNSSec, resp, nil
}
