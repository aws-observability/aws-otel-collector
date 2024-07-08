package linodego

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
)

// This is an enumeration of Capabilities Linode offers that can be referenced
// through the user-facing parts of the application.
// Defined as strings rather than a custom type to avoid breaking change.
// Can be changed in the potential v2 version.
const (
	Linodes                string = "Linodes"
	NodeBalancers          string = "NodeBalancers"
	BlockStorage           string = "Block Storage"
	ObjectStorage          string = "Object Storage"
	ObjectStorageRegions   string = "Object Storage Access Key Regions"
	LKE                    string = "Kubernetes"
	LkeHaControlPlanes     string = "LKE HA Control Planes"
	CloudFirewall          string = "Cloud Firewall"
	GPU                    string = "GPU Linodes"
	Vlans                  string = "Vlans"
	VPCs                   string = "VPCs"
	VPCsExtra              string = "VPCs Extra"
	MachineImages          string = "Machine Images"
	BareMetal              string = "Bare Metal"
	DBAAS                  string = "Managed Databases"
	BlockStorageMigrations string = "Block Storage Migrations"
	Metadata               string = "Metadata"
	PremiumPlans           string = "Premium Plans"
	EdgePlans              string = "Edge Plans"
	LKEControlPlaneACL     string = "LKE Network Access Control List (IP ACL)"
	ACLB                   string = "Akamai Cloud Load Balancer"
	SupportTicketSeverity  string = "Support Ticket Severity"
	Backups                string = "Backups"
	PlacementGroup         string = "Placement Group"
	DiskEncryption         string = "Disk Encryption"
)

// Region-related endpoints have a custom expiry time as the
// `status` field may update for database outages.
var cacheExpiryTime = time.Minute

// Region represents a linode region object
type Region struct {
	ID      string `json:"id"`
	Country string `json:"country"`

	// A List of enums from the above constants
	Capabilities []string `json:"capabilities"`

	Status    string          `json:"status"`
	Resolvers RegionResolvers `json:"resolvers"`
	Label     string          `json:"label"`
	SiteType  string          `json:"site_type"`
}

// RegionResolvers contains the DNS resolvers of a region
type RegionResolvers struct {
	IPv4 string `json:"ipv4"`
	IPv6 string `json:"ipv6"`
}

// RegionsPagedResponse represents a linode API response for listing
type RegionsPagedResponse struct {
	*PageOptions
	Data []Region `json:"data"`
}

// endpoint gets the endpoint URL for Region
func (RegionsPagedResponse) endpoint(_ ...any) string {
	return "regions"
}

func (resp *RegionsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(RegionsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*RegionsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListRegions lists Regions. This endpoint is cached by default.
func (c *Client) ListRegions(ctx context.Context, opts *ListOptions) ([]Region, error) {
	response := RegionsPagedResponse{}

	endpoint, err := generateListCacheURL(response.endpoint(), opts)
	if err != nil {
		return nil, err
	}

	if result := c.getCachedResponse(endpoint); result != nil {
		return result.([]Region), nil
	}

	err = c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}

	c.addCachedResponse(endpoint, response.Data, &cacheExpiryTime)

	return response.Data, nil
}

// GetRegion gets the template with the provided ID. This endpoint is cached by default.
func (c *Client) GetRegion(ctx context.Context, regionID string) (*Region, error) {
	e := fmt.Sprintf("regions/%s", url.PathEscape(regionID))

	if result := c.getCachedResponse(e); result != nil {
		result := result.(Region)
		return &result, nil
	}

	req := c.R(ctx).SetResult(&Region{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}

	c.addCachedResponse(e, r.Result(), &cacheExpiryTime)

	return r.Result().(*Region), nil
}
