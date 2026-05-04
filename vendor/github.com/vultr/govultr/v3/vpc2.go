package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const vpc2Path = "/v2/vpc2"

// VPC2Service is the interface to interact with the VPC 2.0 endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/vpc2
//
// Deprecated: VPC2 is no longer supported and functionality will cease in a
// future release
type VPC2Service interface {
	// Deprecated: VPC2 is no longer supported
	Create(ctx context.Context, createReq *VPC2Req) (*VPC2, *http.Response, error)
	// Deprecated: VPC2 is no longer supported
	Get(ctx context.Context, vpcID string) (*VPC2, *http.Response, error)
	// Deprecated: VPC2 is no longer supported
	Update(ctx context.Context, vpcID string, description string) error
	// Deprecated: VPC2 is no longer supported
	Delete(ctx context.Context, vpcID string) error
	// Deprecated: VPC2 is no longer supported
	List(ctx context.Context, options *ListOptions) ([]VPC2, *Meta, *http.Response, error)
	// Deprecated: VPC2 is no longer supported
	ListNodes(ctx context.Context, vpc2ID string, options *ListOptions) ([]VPC2Node, *Meta, *http.Response, error)
	// Deprecated: VPC2 is no longer supported
	Attach(ctx context.Context, vpcID string, attachReq *VPC2AttachDetachReq) error
	// Deprecated: VPC2 is no longer supported
	Detach(ctx context.Context, vpcID string, detachReq *VPC2AttachDetachReq) error
}

// VPC2ServiceHandler handles interaction with the VPC 2.0 methods for the Vultr API
type VPC2ServiceHandler struct {
	client *Client
}

// VPC2 represents a Vultr VPC 2.0
//
// Deprecated: VPC2 is no longer supported and functionality will cease in a
// future release
type VPC2 struct {
	ID           string `json:"id"`
	Region       string `json:"region"`
	Description  string `json:"description"`
	IPBlock      string `json:"ip_block"`
	PrefixLength int    `json:"prefix_length"`
	DateCreated  string `json:"date_created"`
}

// VPC2Node represents a node attached to a VPC 2.0 network
type VPC2Node struct {
	ID          string `json:"id"`
	IPAddress   string `json:"ip_address"`
	MACAddress  int    `json:"mac_address"`
	Description string `json:"description"`
	Type        string `json:"type"`
	NodeStatus  string `json:"node_status"`
}

// VPC2Req represents parameters to create or update a VPC 2.0 resource
type VPC2Req struct {
	Region       string `json:"region"`
	Description  string `json:"description"`
	IPType       string `json:"ip_type"`
	IPBlock      string `json:"ip_block"`
	PrefixLength int    `json:"prefix_length"`
}

// VPC2AttachDetachReq represents parameters to mass attach or detach nodes from VPC 2.0 networks
type VPC2AttachDetachReq struct {
	Nodes []string `json:"nodes"`
}

type vpcs2Base struct {
	VPCs []VPC2 `json:"vpcs"`
	Meta *Meta  `json:"meta"`
}

type vpc2Base struct {
	VPC *VPC2 `json:"vpc"`
}

type vpc2NodesBase struct {
	Nodes []VPC2Node `json:"nodes"`
	Meta  *Meta      `json:"meta"`
}

// Create creates a new VPC 2.0. A VPC 2.0 can only be used at the location for which it was created.
//
// Deprecated: VPC2 is no longer supported and functionality will cease in a
// future release
func (n *VPC2ServiceHandler) Create(ctx context.Context, createReq *VPC2Req) (*VPC2, *http.Response, error) {
	req, err := n.client.NewRequest(ctx, http.MethodPost, vpc2Path, createReq)
	if err != nil {
		return nil, nil, err
	}

	vpc := new(vpc2Base)
	resp, err := n.client.DoWithContext(ctx, req, vpc)
	if err != nil {
		return nil, resp, err
	}

	return vpc.VPC, resp, nil
}

// Get gets the VPC 2.0 of the requested ID
//
// Deprecated: VPC2 is no longer supported and functionality will cease in a
// future release
func (n *VPC2ServiceHandler) Get(ctx context.Context, vpcID string) (*VPC2, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", vpc2Path, vpcID)
	req, err := n.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	vpc := new(vpc2Base)
	resp, err := n.client.DoWithContext(ctx, req, vpc)
	if err != nil {
		return nil, resp, err
	}

	return vpc.VPC, resp, nil
}

// Update updates a VPC 2.0
//
// Deprecated: VPC2 is no longer supported and functionality will cease in a
// future release
func (n *VPC2ServiceHandler) Update(ctx context.Context, vpcID, description string) error {
	uri := fmt.Sprintf("%s/%s", vpc2Path, vpcID)

	vpcReq := RequestBody{"description": description}
	req, err := n.client.NewRequest(ctx, http.MethodPut, uri, vpcReq)
	if err != nil {
		return err
	}

	_, err = n.client.DoWithContext(ctx, req, nil)
	return err
}

// Delete deletes a VPC 2.0. Before deleting, a VPC 2.0 must be disabled from all instances
//
// Deprecated: VPC2 is no longer supported and functionality will cease in a
// future release
func (n *VPC2ServiceHandler) Delete(ctx context.Context, vpcID string) error {
	uri := fmt.Sprintf("%s/%s", vpc2Path, vpcID)
	req, err := n.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	_, err = n.client.DoWithContext(ctx, req, nil)
	return err
}

// List lists all VPCs 2.0 on the current account
//
// Deprecated: VPC2 is no longer supported and functionality will cease in a
// future release
func (n *VPC2ServiceHandler) List(ctx context.Context, options *ListOptions) ([]VPC2, *Meta, *http.Response, error) { //nolint:dupl
	req, err := n.client.NewRequest(ctx, http.MethodGet, vpc2Path, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	vpcs := new(vpcs2Base)
	resp, err := n.client.DoWithContext(ctx, req, vpcs)
	if err != nil {
		return nil, nil, resp, err
	}

	return vpcs.VPCs, vpcs.Meta, resp, nil
}

// ListNodes lists all nodes attached to a VPC 2.0 network
//
// Deprecated: VPC2 is no longer supported and functionality will cease in a
// future release
func (n *VPC2ServiceHandler) ListNodes(ctx context.Context, vpc2ID string, options *ListOptions) ([]VPC2Node, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/nodes", vpc2Path, vpc2ID)

	req, err := n.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	nodes := new(vpc2NodesBase)
	resp, err := n.client.DoWithContext(ctx, req, nodes)
	if err != nil {
		return nil, nil, resp, err
	}

	return nodes.Nodes, nodes.Meta, resp, nil
}

// Attach attaches nodes to a VPC 2.0 network
//
// Deprecated: VPC2 is no longer supported and functionality will cease in a
// future release
func (n *VPC2ServiceHandler) Attach(ctx context.Context, vpcID string, attachReq *VPC2AttachDetachReq) error {
	uri := fmt.Sprintf("%s/%s/nodes/attach", vpc2Path, vpcID)

	req, err := n.client.NewRequest(ctx, http.MethodPost, uri, attachReq)
	if err != nil {
		return err
	}

	_, err = n.client.DoWithContext(ctx, req, nil)
	return err
}

// Detach detaches nodes from a VPC 2.0 network
//
// Deprecated: VPC2 is no longer supported and functionality will cease in a
// future release
func (n *VPC2ServiceHandler) Detach(ctx context.Context, vpcID string, detachReq *VPC2AttachDetachReq) error {
	uri := fmt.Sprintf("%s/%s/nodes/detach", vpc2Path, vpcID)

	req, err := n.client.NewRequest(ctx, http.MethodPost, uri, detachReq)
	if err != nil {
		return err
	}

	_, err = n.client.DoWithContext(ctx, req, nil)
	return err
}
