// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package ipam provides methods and message types of the ipam v1 API.
package ipam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"text/template"
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/marshaler"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/parameter"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ fmt.Stringer
	_ json.Unmarshaler
	_ url.URL
	_ net.IP
	_ http.Header
	_ bytes.Reader
	_ time.Time
	_ = strings.Join

	_ scw.ScalewayRequest
	_ marshaler.Duration
	_ scw.File
	_ = parameter.AddToQuery
	_ = namegenerator.GetRandomName
)

type ListIPsRequestOrderBy string

const (
	ListIPsRequestOrderByCreatedAtDesc  = ListIPsRequestOrderBy("created_at_desc")
	ListIPsRequestOrderByCreatedAtAsc   = ListIPsRequestOrderBy("created_at_asc")
	ListIPsRequestOrderByUpdatedAtDesc  = ListIPsRequestOrderBy("updated_at_desc")
	ListIPsRequestOrderByUpdatedAtAsc   = ListIPsRequestOrderBy("updated_at_asc")
	ListIPsRequestOrderByAttachedAtDesc = ListIPsRequestOrderBy("attached_at_desc")
	ListIPsRequestOrderByAttachedAtAsc  = ListIPsRequestOrderBy("attached_at_asc")
	ListIPsRequestOrderByIPAddressDesc  = ListIPsRequestOrderBy("ip_address_desc")
	ListIPsRequestOrderByIPAddressAsc   = ListIPsRequestOrderBy("ip_address_asc")
	ListIPsRequestOrderByMacAddressDesc = ListIPsRequestOrderBy("mac_address_desc")
	ListIPsRequestOrderByMacAddressAsc  = ListIPsRequestOrderBy("mac_address_asc")
)

func (enum ListIPsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListIPsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListIPsRequestOrderBy) Values() []ListIPsRequestOrderBy {
	return []ListIPsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
		"updated_at_desc",
		"updated_at_asc",
		"attached_at_desc",
		"attached_at_asc",
		"ip_address_desc",
		"ip_address_asc",
		"mac_address_desc",
		"mac_address_asc",
	}
}

func (enum ListIPsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListIPsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListIPsRequestOrderBy(ListIPsRequestOrderBy(tmp).String())
	return nil
}

type ResourceType string

const (
	ResourceTypeUnknownType            = ResourceType("unknown_type")
	ResourceTypeCustom                 = ResourceType("custom")
	ResourceTypeInstanceServer         = ResourceType("instance_server")
	ResourceTypeInstanceIP             = ResourceType("instance_ip")
	ResourceTypeInstancePrivateNic     = ResourceType("instance_private_nic")
	ResourceTypeLBServer               = ResourceType("lb_server")
	ResourceTypeFipIP                  = ResourceType("fip_ip")
	ResourceTypeVpcGateway             = ResourceType("vpc_gateway")
	ResourceTypeVpcGatewayNetwork      = ResourceType("vpc_gateway_network")
	ResourceTypeK8sNode                = ResourceType("k8s_node")
	ResourceTypeK8sCluster             = ResourceType("k8s_cluster")
	ResourceTypeRdbInstance            = ResourceType("rdb_instance")
	ResourceTypeRedisCluster           = ResourceType("redis_cluster")
	ResourceTypeBaremetalServer        = ResourceType("baremetal_server")
	ResourceTypeBaremetalPrivateNic    = ResourceType("baremetal_private_nic")
	ResourceTypeLlmDeployment          = ResourceType("llm_deployment")
	ResourceTypeMgdbInstance           = ResourceType("mgdb_instance")
	ResourceTypeAppleSiliconServer     = ResourceType("apple_silicon_server")
	ResourceTypeAppleSiliconPrivateNic = ResourceType("apple_silicon_private_nic")
	ResourceTypeServerlessContainer    = ResourceType("serverless_container")
	ResourceTypeServerlessFunction     = ResourceType("serverless_function")
	ResourceTypeVpnGateway             = ResourceType("vpn_gateway")
	ResourceTypeDdlDatalab             = ResourceType("ddl_datalab")
	ResourceTypeKafkaCluster           = ResourceType("kafka_cluster")
	ResourceTypeBgpEndpoint            = ResourceType("bgp_endpoint")
	ResourceTypeScblSedbCluster        = ResourceType("scbl_sedb_cluster")
	ResourceTypeDtwhDeployment         = ResourceType("dtwh_deployment")
	ResourceTypeSedbCluster            = ResourceType("sedb_cluster")
	ResourceTypeMsgqCluster            = ResourceType("msgq_cluster")
	ResourceTypeEdgeVpcEndpoint        = ResourceType("edge_vpc_endpoint")
	ResourceTypeDvizCluster            = ResourceType("dviz_cluster")
)

func (enum ResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ResourceTypeUnknownType)
	}
	return string(enum)
}

func (enum ResourceType) Values() []ResourceType {
	return []ResourceType{
		"unknown_type",
		"custom",
		"instance_server",
		"instance_ip",
		"instance_private_nic",
		"lb_server",
		"fip_ip",
		"vpc_gateway",
		"vpc_gateway_network",
		"k8s_node",
		"k8s_cluster",
		"rdb_instance",
		"redis_cluster",
		"baremetal_server",
		"baremetal_private_nic",
		"llm_deployment",
		"mgdb_instance",
		"apple_silicon_server",
		"apple_silicon_private_nic",
		"serverless_container",
		"serverless_function",
		"vpn_gateway",
		"ddl_datalab",
		"kafka_cluster",
		"bgp_endpoint",
		"scbl_sedb_cluster",
		"dtwh_deployment",
		"sedb_cluster",
		"msgq_cluster",
		"edge_vpc_endpoint",
		"dviz_cluster",
	}
}

func (enum ResourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ResourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ResourceType(ResourceType(tmp).String())
	return nil
}

// Resource: resource.
type Resource struct {
	// Type: type of resource the IP is attached to.
	// Default value: unknown_type
	Type ResourceType `json:"type"`

	// ID: ID of the resource the IP is attached to.
	ID string `json:"id"`

	// MacAddress: mAC of the resource the IP is attached to.
	MacAddress *string `json:"mac_address"`

	// Name: when the IP is in a Private Network, then a DNS record is available to resolve the resource name to this IP.
	Name *string `json:"name"`
}

// Reverse: reverse.
type Reverse struct {
	// Hostname: reverse domain name.
	Hostname string `json:"hostname"`

	// Address: IP corresponding to the hostname.
	Address *net.IP `json:"address"`
}

// Source: source.
type Source struct {
	// Zonal: this source is global.
	// Precisely one of Zonal, PrivateNetworkID, SubnetID, VpcID, Regional must be set.
	Zonal *string `json:"zonal,omitempty"`

	// PrivateNetworkID: this source is specific.
	// Precisely one of Zonal, PrivateNetworkID, SubnetID, VpcID, Regional must be set.
	PrivateNetworkID *string `json:"private_network_id,omitempty"`

	// SubnetID: this source is specific.
	// Precisely one of Zonal, PrivateNetworkID, SubnetID, VpcID, Regional must be set.
	SubnetID *string `json:"subnet_id,omitempty"`

	// Precisely one of Zonal, PrivateNetworkID, SubnetID, VpcID, Regional must be set.
	VpcID *string `json:"vpc_id,omitempty"`

	// Regional: this source is global.
	// Precisely one of Zonal, PrivateNetworkID, SubnetID, VpcID, Regional must be set.
	Regional *bool `json:"regional,omitempty"`
}

// CustomResource: custom resource.
type CustomResource struct {
	// MacAddress: mAC address of the custom resource.
	MacAddress string `json:"mac_address"`

	// Name: when the resource is in a Private Network, a DNS record is available to resolve the resource name.
	Name *string `json:"name"`
}

// IP: ip.
type IP struct {
	// ID: IP ID.
	ID string `json:"id"`

	// Address: iPv4 or IPv6 address in CIDR notation.
	Address scw.IPNet `json:"address"`

	// ProjectID: scaleway Project the IP belongs to.
	ProjectID string `json:"project_id"`

	// IsIPv6: defines whether the IP is an IPv6 (false = IPv4).
	IsIPv6 bool `json:"is_ipv6"`

	// CreatedAt: date the IP was reserved.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the IP was last modified.
	UpdatedAt *time.Time `json:"updated_at"`

	// Source: source pool where the IP was reserved in.
	Source *Source `json:"source"`

	// Resource: resource which the IP is attached to.
	Resource *Resource `json:"resource"`

	// Tags: tags for the IP.
	Tags []string `json:"tags"`

	// Reverses: array of reverses associated with the IP.
	Reverses []*Reverse `json:"reverses"`

	// Region: region of the IP.
	Region scw.Region `json:"region"`

	// Zone: zone of the IP, if zonal.
	Zone *scw.Zone `json:"zone"`

	// This field is automatically generated, do not edit it
	Srn string `json:"srn,omitempty"`
}

func (m *IP) getSRNTemplates() []string {
	return []string{
		"srn://ipam.{{ notempty .Platform }}/zones/{{ notempty .Zone }}/ips/{{ notempty .ID }}",
		"srn://ipam.{{ notempty .Platform }}/regions/{{ notempty .Region }}/ips/{{ notempty .ID }}",
	}
}

func (m *IP) setSRN(platform string) {
	if m.Srn != "" {
		// if the field is set server-side, trust the server
		return
	}
	data := struct {
		IP
		Platform string
	}{
		IP:       *m,
		Platform: platform,
	}

	notEmpty := func(a any) (string, error) {
		s := fmt.Sprint(a)
		if s == "" {
			return "", errors.New("value is empty")
		}
		return s, nil
	}
	for _, templ := range m.getSRNTemplates() {
		t, err := template.New("srn").Funcs(template.FuncMap{"notempty": notEmpty}).Parse(templ)
		if err != nil {
			continue
		}
		var out bytes.Buffer
		if err := t.Execute(&out, data); err != nil {
			continue
		}
		m.Srn = out.String()
		// first pattern wins
		break
	}
}

// AttachIPRequest: attach ip request.
type AttachIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// IPID: IP ID.
	IPID string `json:"-"`

	// Resource: custom resource to be attached to the IP.
	Resource *CustomResource `json:"resource"`
}

// BookIPRequest: book ip request.
type BookIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: when creating an IP in a Private Network, the Project must match the Private Network's Project.
	ProjectID string `json:"project_id"`

	// Source: source in which to reserve the IP. Not all sources are available for reservation.
	Source *Source `json:"source"`

	// IsIPv6: request an IPv6 instead of an IPv4.
	IsIPv6 bool `json:"is_ipv6"`

	// Address: the requested address should not include the subnet mask (/suffix). Note that only the Private Network source allows you to pick a specific IP. If the requested IP is already reserved, then the call will fail.
	Address *net.IP `json:"address,omitempty"`

	// Tags: tags for the IP.
	Tags []string `json:"tags"`

	// Resource: custom resource to attach to the IP being reserved. An example of a custom resource is a virtual machine hosted on an Elastic Metal server. Do not use this for attaching IP addresses to standard Scaleway resources, as it will fail - instead, see the relevant product API for an equivalent method.
	Resource *CustomResource `json:"resource,omitempty"`
}

// DetachIPRequest: detach ip request.
type DetachIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// IPID: IP ID.
	IPID string `json:"-"`

	// Resource: custom resource currently attached to the IP.
	Resource *CustomResource `json:"resource"`
}

// GetIPRequest: get ip request.
type GetIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// IPID: IP ID.
	IPID string `json:"-"`
}

// ListIPsRequest: list i ps request.
type ListIPsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: sort order of the returned IPs.
	// Default value: created_at_desc
	OrderBy ListIPsRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: maximum number of IPs to return per page.
	PageSize *uint32 `json:"-"`

	// ProjectID: project ID to filter for. Only IPs belonging to this Project will be returned.
	ProjectID *string `json:"-"`

	// Zonal: zone to filter for. Only IPs that are zonal, and in this zone, will be returned.
	// Precisely one of Zonal, Regional, PrivateNetworkID, SubnetID, SourceVpcID must be set.
	Zonal *string `json:"zonal,omitempty"`

	// Regional: filter on regional IPs only.
	// Precisely one of Zonal, Regional, PrivateNetworkID, SubnetID, SourceVpcID must be set.
	Regional *bool `json:"regional,omitempty"`

	// PrivateNetworkID: only IPs that are private, and in this Private Network, will be returned.
	// Precisely one of Zonal, Regional, PrivateNetworkID, SubnetID, SourceVpcID must be set.
	PrivateNetworkID *string `json:"private_network_id,omitempty"`

	// SubnetID: only IPs inside this exact subnet will be returned.
	// Precisely one of Zonal, Regional, PrivateNetworkID, SubnetID, SourceVpcID must be set.
	SubnetID *string `json:"subnet_id,omitempty"`

	// VpcID: only IPs owned by resources in this VPC will be returned.
	VpcID *string `json:"-"`

	// Attached: defines whether to filter only for IPs which are attached to a resource.
	Attached *bool `json:"-"`

	// ResourceName: attached resource name to filter for, only IPs attached to a resource with this string within their name will be returned.
	ResourceName *string `json:"-"`

	// ResourceID: resource ID to filter for. Only IPs attached to this resource will be returned.
	ResourceID *string `json:"-"`

	// ResourceIDs: resource IDs to filter for. Only IPs attached to at least one of these resources will be returned.
	ResourceIDs []string `json:"-"`

	// ResourceType: resource type to filter for. Only IPs attached to this type of resource will be returned.
	// Default value: unknown_type
	ResourceType ResourceType `json:"-"`

	// ResourceTypes: resource types to filter for. Only IPs attached to these types of resources will be returned.
	ResourceTypes []ResourceType `json:"-"`

	// MacAddress: mAC address to filter for. Only IPs attached to a resource with this MAC address will be returned.
	MacAddress *string `json:"-"`

	// Tags: tags to filter for, only IPs with one or more matching tags will be returned.
	Tags []string `json:"-"`

	// OrganizationID: organization ID to filter for. Only IPs belonging to this Organization will be returned.
	OrganizationID *string `json:"-"`

	// IsIPv6: defines whether to filter only for IPv4s or IPv6s.
	IsIPv6 *bool `json:"-"`

	// IPIDs: IP IDs to filter for. Only IPs with these UUIDs will be returned.
	IPIDs []string `json:"-"`

	// Precisely one of Zonal, Regional, PrivateNetworkID, SubnetID, SourceVpcID must be set.
	SourceVpcID *string `json:"source_vpc_id,omitempty"`
}

// ListIPsResponse: list i ps response.
type ListIPsResponse struct {
	TotalCount uint64 `json:"total_count"`

	IPs []*IP `json:"ips"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.IPs = append(r.IPs, results.IPs...)
	r.TotalCount += uint64(len(results.IPs))
	return uint64(len(results.IPs)), nil
}

// MoveIPRequest: move ip request.
type MoveIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// IPID: IP ID.
	IPID string `json:"-"`

	// FromResource: custom resource currently attached to the IP.
	FromResource *CustomResource `json:"from_resource"`

	// ToResource: custom resource to be attached to the IP.
	ToResource *CustomResource `json:"to_resource,omitempty"`
}

// ReleaseIPRequest: release ip request.
type ReleaseIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// IPID: IP ID.
	IPID string `json:"-"`
}

// ReleaseIPSetRequest: release ip set request.
type ReleaseIPSetRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IPIDs []string `json:"ip_ids"`
}

// UpdateIPRequest: update ip request.
type UpdateIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// IPID: IP ID.
	IPID string `json:"-"`

	// Tags: tags for the IP.
	Tags *[]string `json:"tags,omitempty"`

	// Reverses: array of reverse domain names associated with an IP in the subnet of the current IP.
	Reverses []*Reverse `json:"reverses"`
}

// This API allows you to manage your Scaleway IP addresses with our IP Address Management tool.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionItMil, scw.RegionNlAms, scw.RegionPlWaw}
}

// BookIP: Reserve a new IP from the specified source. Currently IPs can only be reserved from a Private Network.
func (s *API) BookIP(req *BookIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam/v1/regions/" + fmt.Sprint(req.Region) + "/ips",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	// platform := s.client.GetPlatform()
	platform := "scw.eu"
	resp.setSRN(platform)
	return &resp, nil
}

// ReleaseIP: Release an IP not currently attached to a resource, and returns it to the available IP pool.
func (s *API) ReleaseIP(req *ReleaseIPRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/ipam/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ReleaseIPSet:
func (s *API) ReleaseIPSet(req *ReleaseIPSetRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam/v1/regions/" + fmt.Sprint(req.Region) + "/ip-sets/release",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetIP: Retrieve details of an existing IP, specified by its IP ID.
func (s *API) GetIP(req *GetIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipam/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "",
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	// platform := s.client.GetPlatform()
	platform := "scw.eu"
	resp.setSRN(platform)
	return &resp, nil
}

// UpdateIP: Update parameters including tags of the specified IP.
func (s *API) UpdateIP(req *UpdateIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/ipam/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	// platform := s.client.GetPlatform()
	platform := "scw.eu"
	resp.setSRN(platform)
	return &resp, nil
}

// ListIPs: List existing IPs in the specified region using various filters. For example, you can filter for IPs within a specified Private Network, or for public IPs within a specified Project. By default, the IPs returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.
func (s *API) ListIPs(req *ListIPsRequest, opts ...scw.RequestOption) (*ListIPsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "vpc_id", req.VpcID)
	parameter.AddToQuery(query, "attached", req.Attached)
	parameter.AddToQuery(query, "resource_name", req.ResourceName)
	parameter.AddToQuery(query, "resource_id", req.ResourceID)
	parameter.AddToQuery(query, "resource_ids", req.ResourceIDs)
	parameter.AddToQuery(query, "resource_type", req.ResourceType)
	parameter.AddToQuery(query, "resource_types", req.ResourceTypes)
	parameter.AddToQuery(query, "mac_address", req.MacAddress)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "is_ipv6", req.IsIPv6)
	parameter.AddToQuery(query, "ip_ids", req.IPIDs)
	parameter.AddToQuery(query, "zonal", req.Zonal)
	parameter.AddToQuery(query, "regional", req.Regional)
	parameter.AddToQuery(query, "private_network_id", req.PrivateNetworkID)
	parameter.AddToQuery(query, "subnet_id", req.SubnetID)
	parameter.AddToQuery(query, "source_vpc_id", req.SourceVpcID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipam/v1/regions/" + fmt.Sprint(req.Region) + "/ips",
		Query:  query,
	}

	var resp ListIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	// platform := s.client.GetPlatform()
	platform := "scw.eu"
	for _, el := range resp.IPs {
		el.setSRN(platform)
	}
	return &resp, nil
}

// AttachIP: Attach an existing reserved private IP from a Private Network subnet to a custom, named resource via its MAC address. An example of a custom resource is a virtual machine hosted on an Elastic Metal server. Do not use this method for attaching IP addresses to standard Scaleway resources as it will fail - see the relevant product API for an equivalent method.
func (s *API) AttachIP(req *AttachIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "/attach",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	// platform := s.client.GetPlatform()
	platform := "scw.eu"
	resp.setSRN(platform)
	return &resp, nil
}

// DetachIP: Detach a private IP from a custom resource. An example of a custom resource is a virtual machine hosted on an Elastic Metal server. Do not use this method for detaching IP addresses from standard Scaleway resources (e.g. Instances, Load Balancers) as it will fail - see the relevant product API for an equivalent method.
func (s *API) DetachIP(req *DetachIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "/detach",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	// platform := s.client.GetPlatform()
	platform := "scw.eu"
	resp.setSRN(platform)
	return &resp, nil
}

// MoveIP: Move an existing reserved private IP from one custom resource (e.g. a virtual machine hosted on an Elastic Metal server) to another custom resource. This will detach it from the first resource, and attach it to the second. Do not use this method for moving IP addresses between standard Scaleway resources (e.g. Instances, Load Balancers) as it will fail - see the relevant product API for an equivalent method.
func (s *API) MoveIP(req *MoveIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "/move",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	// platform := s.client.GetPlatform()
	platform := "scw.eu"
	resp.setSRN(platform)
	return &resp, nil
}
