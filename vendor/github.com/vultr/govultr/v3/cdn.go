package govultr

import (
	"context"
	"fmt"
	"net/http"
)

const cdnPath string = "/v2/cdns"

var (
	cdnPullPath string = fmt.Sprintf("%s/pull-zones", cdnPath)
	cdnPushPath string = fmt.Sprintf("%s/push-zones", cdnPath)
)

// CDNService is the interface to interact with the CDN endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/CDNs
type CDNService interface {
	ListPullZones(ctx context.Context) ([]CDNZone, *Meta, *http.Response, error)
	GetPullZone(ctx context.Context, zoneID string) (*CDNZone, *http.Response, error)
	CreatePullZone(ctx context.Context, zoneReq *CDNZoneReq) (*CDNZone, *http.Response, error)
	UpdatePullZone(ctx context.Context, zoneID string, zoneReq *CDNZoneReq) (*CDNZone, *http.Response, error)
	DeletePullZone(ctx context.Context, zoneID string) error
	PurgePullZone(ctx context.Context, zoneID string) error

	ListPushZones(ctx context.Context) ([]CDNZone, *Meta, *http.Response, error)
	GetPushZone(ctx context.Context, zoneID string) (*CDNZone, *http.Response, error)
	CreatePushZone(ctx context.Context, zoneReq *CDNZoneReq) (*CDNZone, *http.Response, error)
	UpdatePushZone(ctx context.Context, zoneID string, zoneReq *CDNZoneReq) (*CDNZone, *http.Response, error)
	DeletePushZone(ctx context.Context, zoneID string) error

	ListPushZoneFiles(ctx context.Context, zoneID string) (*CDNZoneFileData, *http.Response, error)
	GetPushZoneFile(ctx context.Context, zoneID, fileName string) (*CDNZoneFile, *http.Response, error)
	DeletePushZoneFile(ctx context.Context, zoneID, fileName string) error
	CreatePushZoneFileEndpoint(ctx context.Context, zoneID string, endpointReq *CDNZoneEndpointReq) (*CDNZoneEndpoint, *http.Response, error) //nolint:lll
}

// CDNZone represents the CDN push/pull zone data
type CDNZone struct {
	ID            string   `json:"id"`
	DateCreated   string   `json:"date_created"`
	Status        string   `json:"status"`
	Label         string   `json:"label"`
	OriginScheme  string   `json:"origin_scheme"`
	OriginDomain  string   `json:"origin_domain"`
	VanityDomain  string   `json:"vanity_domain"`
	SSLCert       string   `json:"ssl_cert"`
	SSLCertKey    string   `json:"ssl_cert_key"`
	CDNURL        string   `json:"cdn_url"`
	CacheSize     int      `json:"cache_size"`
	Requests      int      `json:"requests"`
	BytesIn       int      `json:"in_bytes"`
	BytesOut      int      `json:"out_bytes"`
	PacketsPerSec int      `json:"packets_per_sec"`
	DatePurged    string   `json:"last_purge"`
	CORS          bool     `json:"cors"`
	GZIP          bool     `json:"gzip"`
	BlockAI       bool     `json:"block_ai"`
	BlockBadBots  bool     `json:"block_bad_bots"`
	Regions       []string `json:"regions"`
}

// CDNZoneReq is the data used to create a push/pull zone
type CDNZoneReq struct {
	Label        string   `json:"label"`
	OriginScheme string   `json:"origin_scheme,omitempty"`
	OriginDomain string   `json:"origin_domain,omitempty"`
	VanityDomain string   `json:"vanity_domain,omitempty"`
	SSLCert      string   `json:"ssl_cert,omitempty"`
	SSLCertKey   string   `json:"ssl_cert_key,omitempty"`
	CORS         *bool    `json:"cors,omitempty"`
	GZIP         *bool    `json:"gzip,omitempty"`
	BlockAI      *bool    `json:"block_ai,omitempty"`
	BlockBadBots *bool    `json:"block_bad_bots,omitempty"`
	Regions      []string `json:"regions,omitempty"`
}

// CDNZoneFileData represents the push zone files and metadata
type CDNZoneFileData struct {
	Files []CDNZoneFile `json:"files"`
	Count int           `json:"count"`
	Size  int           `json:"total_size"`
}

// CDNZoneFile is the data for a push zone file
type CDNZoneFile struct {
	Name         string `json:"name"`
	Size         int    `json:"size"`
	DateModified string `json:"last_modified"`
}

// CDNZoneEndpointReq is the data used to create a push zone upload endpoint
type CDNZoneEndpointReq struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

// CDNZoneEndpoint is the data for a push zone file endpoint
type CDNZoneEndpoint struct {
	URL    string                `json:"url"`
	Inputs CDNZoneEndpointInputs `json:"inputs"`
}

// CDNZoneEndpointInputs is the data for push zone file endpoint inputs
type CDNZoneEndpointInputs struct {
	ACL        string `json:"acl"`
	Key        string `json:"key"`
	Policy     string `json:"policy"`
	Credential string `json:"x-amz-credential"`
	Algorithm  string `json:"x-amz-algorithm"`
	Signature  string `json:"x-amz-signature"`
}

type cdnPullZonesBase struct {
	PullZones []CDNZone `json:"pull_zones"`
	Meta      *Meta     `json:"meta"`
}

type cdnPullZoneBase struct {
	PullZone CDNZone `json:"pull_zone"`
}

type cdnPushZonesBase struct {
	PushZones []CDNZone `json:"push_zones"`
	Meta      *Meta     `json:"meta"`
}

type cdnPushZoneBase struct {
	PushZone CDNZone `json:"push_zone"`
}

type cdnPushZoneEndpointBase struct {
	Endpoint CDNZoneEndpoint `json:"upload_endpoint"`
}

// CDNServiceHandler handles interaction with the CDN methods for the Vultr API
type CDNServiceHandler struct {
	client *Client
}

// ListPullZones will retrieve a CDN pull zone
func (c *CDNServiceHandler) ListPullZones(ctx context.Context) ([]CDNZone, *Meta, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, cdnPullPath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	zones := new(cdnPullZonesBase)
	resp, err := c.client.DoWithContext(ctx, req, &zones)
	if err != nil {
		return nil, nil, resp, err
	}

	return zones.PullZones, zones.Meta, resp, nil
}

// GetPullZone will retrieve data for a CDN pull zone
func (c *CDNServiceHandler) GetPullZone(ctx context.Context, zoneID string) (*CDNZone, *http.Response, error) {
	cdnPullGetPath := fmt.Sprintf("%s/%s", cdnPullPath, zoneID)
	req, err := c.client.NewRequest(ctx, http.MethodGet, cdnPullGetPath, nil)
	if err != nil {
		return nil, nil, err
	}

	zone := new(cdnPullZoneBase)
	resp, err := c.client.DoWithContext(ctx, req, &zone)
	if err != nil {
		return nil, resp, err
	}

	return &zone.PullZone, resp, nil
}

// CreatePullZone will create a new CDN pull zone
func (c *CDNServiceHandler) CreatePullZone(ctx context.Context, zoneReq *CDNZoneReq) (*CDNZone, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, cdnPullPath, zoneReq)
	if err != nil {
		return nil, nil, err
	}

	zone := new(cdnPullZoneBase)
	resp, err := c.client.DoWithContext(ctx, req, &zone)
	if err != nil {
		return nil, resp, err
	}

	return &zone.PullZone, resp, nil
}

// UpdatePullZone will update an existing CDN pull zone
func (c *CDNServiceHandler) UpdatePullZone(ctx context.Context, zoneID string, zoneReq *CDNZoneReq) (*CDNZone, *http.Response, error) { //nolint:dupl,lll
	cndPullUpdatePath := fmt.Sprintf("%s/%s", cdnPullPath, zoneID)
	req, err := c.client.NewRequest(ctx, http.MethodPut, cndPullUpdatePath, zoneReq)
	if err != nil {
		return nil, nil, err
	}

	zone := new(cdnPullZoneBase)
	resp, err := c.client.DoWithContext(ctx, req, &zone)
	if err != nil {
		return nil, resp, err
	}

	return &zone.PullZone, resp, nil
}

// DeletePullZone will delete a CDN pull zone
func (c *CDNServiceHandler) DeletePullZone(ctx context.Context, zoneID string) error {
	cdnPullDeletePath := fmt.Sprintf("%s/%s", cdnPullPath, zoneID)
	req, err := c.client.NewRequest(ctx, http.MethodDelete, cdnPullDeletePath, nil)
	if err != nil {
		return err
	}

	_, err = c.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// PurgePullZone will clear the cache on a CDN pull zone
func (c *CDNServiceHandler) PurgePullZone(ctx context.Context, zoneID string) error {
	cdnPullPurgePath := fmt.Sprintf("%s/%s", cdnPullPath, zoneID)
	req, err := c.client.NewRequest(ctx, http.MethodGet, cdnPullPurgePath, nil)
	if err != nil {
		return err
	}

	_, err = c.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// ListPushZones will retrieve a CDN push zone
func (c *CDNServiceHandler) ListPushZones(ctx context.Context) ([]CDNZone, *Meta, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, cdnPushPath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	zones := new(cdnPushZonesBase)
	resp, err := c.client.DoWithContext(ctx, req, &zones)
	if err != nil {
		return nil, nil, resp, err
	}

	return zones.PushZones, zones.Meta, resp, nil
}

// GetPushZone will retrieve data for a CDN push zone
func (c *CDNServiceHandler) GetPushZone(ctx context.Context, zoneID string) (*CDNZone, *http.Response, error) {
	cdnPushGetPath := fmt.Sprintf("%s/%s", cdnPushPath, zoneID)
	req, err := c.client.NewRequest(ctx, http.MethodGet, cdnPushGetPath, nil)
	if err != nil {
		return nil, nil, err
	}

	zone := new(cdnPushZoneBase)
	resp, err := c.client.DoWithContext(ctx, req, &zone)
	if err != nil {
		return nil, resp, err
	}

	return &zone.PushZone, resp, nil
}

// CreatePushZone will create a new CDN push zone
func (c *CDNServiceHandler) CreatePushZone(ctx context.Context, zoneReq *CDNZoneReq) (*CDNZone, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, cdnPushPath, zoneReq)
	if err != nil {
		return nil, nil, err
	}

	zone := new(cdnPushZoneBase)
	resp, err := c.client.DoWithContext(ctx, req, &zone)
	if err != nil {
		return nil, resp, err
	}

	return &zone.PushZone, resp, nil
}

// UpdatePushZone will update an existing CDN push zone
func (c *CDNServiceHandler) UpdatePushZone(ctx context.Context, zoneID string, zoneReq *CDNZoneReq) (*CDNZone, *http.Response, error) { //nolint:dupl,lll
	cndPushUpdatePath := fmt.Sprintf("%s/%s", cdnPushPath, zoneID)
	req, err := c.client.NewRequest(ctx, http.MethodPut, cndPushUpdatePath, zoneReq)
	if err != nil {
		return nil, nil, err
	}

	zone := new(cdnPushZoneBase)
	resp, err := c.client.DoWithContext(ctx, req, &zone)
	if err != nil {
		return nil, resp, err
	}

	return &zone.PushZone, resp, nil
}

// DeletePushZone will delete a CDN push zone
func (c *CDNServiceHandler) DeletePushZone(ctx context.Context, zoneID string) error {
	cdnPushDeletePath := fmt.Sprintf("%s/%s", cdnPushPath, zoneID)
	req, err := c.client.NewRequest(ctx, http.MethodDelete, cdnPushDeletePath, nil)
	if err != nil {
		return err
	}

	_, err = c.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// CreatePushZoneFileEndpoint will create a new CDN push zone file upload
// endpoint
func (c *CDNServiceHandler) CreatePushZoneFileEndpoint(ctx context.Context, zoneID string, zoneEndpointReq *CDNZoneEndpointReq) (*CDNZoneEndpoint, *http.Response, error) { //nolint:lll,dupl
	cdnPushEnpointPath := fmt.Sprintf("%s/%s/files", cdnPushPath, zoneID)
	req, err := c.client.NewRequest(ctx, http.MethodPost, cdnPushEnpointPath, zoneEndpointReq)
	if err != nil {
		return nil, nil, err
	}

	endpoint := new(cdnPushZoneEndpointBase)
	resp, err := c.client.DoWithContext(ctx, req, &endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &endpoint.Endpoint, resp, nil
}

// ListPushZoneFiles will retrieve all CDN push zone file data that have been
// uploaded
func (c *CDNServiceHandler) ListPushZoneFiles(ctx context.Context, zoneID string) (*CDNZoneFileData, *http.Response, error) {
	cdnPushFilesPath := fmt.Sprintf("%s/%s/files", cdnPushPath, zoneID)
	req, err := c.client.NewRequest(ctx, http.MethodGet, cdnPushFilesPath, nil)
	if err != nil {
		return nil, nil, err
	}

	fd := new(CDNZoneFileData)
	resp, err := c.client.DoWithContext(ctx, req, &fd)
	if err != nil {
		return nil, resp, err
	}

	return fd, resp, nil
}

// GetPushZoneFile will retrieve data on a file in a CDN push zone
func (c *CDNServiceHandler) GetPushZoneFile(ctx context.Context, zoneID, fileName string) (*CDNZoneFile, *http.Response, error) {
	cdnFileGetPath := fmt.Sprintf("%s/%s/files/%s", cdnPushPath, zoneID, fileName)
	req, err := c.client.NewRequest(ctx, http.MethodGet, cdnFileGetPath, nil)
	if err != nil {
		return nil, nil, err
	}

	file := new(CDNZoneFile)
	resp, err := c.client.DoWithContext(ctx, req, &file)
	if err != nil {
		return nil, resp, err
	}

	return file, resp, nil
}

// DeletePushZoneFile delete a file in a CDN push zone
func (c *CDNServiceHandler) DeletePushZoneFile(ctx context.Context, zoneID, fileName string) error {
	cdnFileDelPath := fmt.Sprintf("%s/%s/files/%s", cdnPushPath, zoneID, fileName)
	if _, err := c.client.NewRequest(ctx, http.MethodDelete, cdnFileDelPath, nil); err != nil {
		return err
	}

	return nil
}
