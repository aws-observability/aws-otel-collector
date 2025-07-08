package linodego

import (
	"context"
	"time"
)

// This is an enumeration of Capabilities Linode offers that can be referenced
// through the user-facing parts of the application.
// Defined as strings rather than a custom type to avoid breaking change.
// Can be changed in the potential v2 version.
const (
	CapabilityACLB                          string = "Akamai Cloud Load Balancer"
	CapabilityBackups                       string = "Backups"
	CapabilityBareMetal                     string = "Bare Metal"
	CapabilityBlockStorage                  string = "Block Storage"
	CapabilityBlockStorageEncryption        string = "Block Storage Encryption"
	CapabilityBlockStorageMigrations        string = "Block Storage Migrations"
	CapabilityCloudFirewall                 string = "Cloud Firewall"
	CapabilityDBAAS                         string = "Managed Databases"
	CapabilityDiskEncryption                string = "Disk Encryption"
	CapabilityEdgePlans                     string = "Edge Plans"
	CapabilityGPU                           string = "GPU Linodes"
	CapabilityKubernetesEnterprise          string = "Kubernetes Enterprise"
	CapabilityLKE                           string = "Kubernetes"
	CapabilityLKEControlPlaneACL            string = "LKE Network Access Control List (IP ACL)"
	CapabilityLinodes                       string = "Linodes"
	CapabilityLkeHaControlPlanes            string = "LKE HA Control Planes"
	CapabilityMachineImages                 string = "Machine Images"
	CapabilityMetadata                      string = "Metadata"
	CapabilityNodeBalancers                 string = "NodeBalancers"
	CapabilityObjectStorage                 string = "Object Storage"
	CapabilityObjectStorageAccessKeyRegions string = "Object Storage Access Key Regions"
	CapabilityObjectStorageEndpointTypes    string = "Object Storage Endpoint Types"
	CapabilityPlacementGroup                string = "Placement Group"
	CapabilityPremiumPlans                  string = "Premium Plans"
	CapabilityQuadraT1UVPU                  string = "NETINT Quadra T1U"
	CapabilitySupportTicketSeverity         string = "Support Ticket Severity"
	CapabilityVPCs                          string = "VPCs"
	CapabilityVPCsExtra                     string = "VPCs Extra"
	CapabilityVlans                         string = "Vlans"

	// Deprecated: CapabilityObjectStorageRegions constant has been
	// renamed to `CapabilityObjectStorageAccessKeyRegions`.
	CapabilityObjectStorageRegions string = CapabilityObjectStorageAccessKeyRegions
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

	Status   string `json:"status"`
	Label    string `json:"label"`
	SiteType string `json:"site_type"`

	Resolvers            RegionResolvers             `json:"resolvers"`
	PlacementGroupLimits *RegionPlacementGroupLimits `json:"placement_group_limits"`
}

// RegionResolvers contains the DNS resolvers of a region
type RegionResolvers struct {
	IPv4 string `json:"ipv4"`
	IPv6 string `json:"ipv6"`
}

// RegionPlacementGroupLimits contains information about the
// placement group limits for the current user in the current region.
type RegionPlacementGroupLimits struct {
	MaximumPGsPerCustomer int `json:"maximum_pgs_per_customer"`
	MaximumLinodesPerPG   int `json:"maximum_linodes_per_pg"`
}

// ListRegions lists Regions. This endpoint is cached by default.
func (c *Client) ListRegions(ctx context.Context, opts *ListOptions) ([]Region, error) {
	endpoint, err := generateListCacheURL("regions", opts)
	if err != nil {
		return nil, err
	}

	if result := c.getCachedResponse(endpoint); result != nil {
		return result.([]Region), nil
	}

	response, err := getPaginatedResults[Region](ctx, c, "regions", opts)
	if err != nil {
		return nil, err
	}

	c.addCachedResponse(endpoint, response, &cacheExpiryTime)

	return response, nil
}

// GetRegion gets the template with the provided ID. This endpoint is cached by default.
func (c *Client) GetRegion(ctx context.Context, regionID string) (*Region, error) {
	e := formatAPIPath("regions/%s", regionID)

	if result := c.getCachedResponse(e); result != nil {
		result := result.(Region)
		return &result, nil
	}

	response, err := doGETRequest[Region](ctx, c, e)
	if err != nil {
		return nil, err
	}

	c.addCachedResponse(e, response, &cacheExpiryTime)

	return response, nil
}
