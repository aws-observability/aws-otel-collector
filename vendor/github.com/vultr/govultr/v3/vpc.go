package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const vpcPath = "/v2/vpcs"

// VPCService is the interface to interact with the VPC endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/vpcs
type VPCService interface {
	Create(ctx context.Context, createReq *VPCReq) (*VPC, *http.Response, error)
	Get(ctx context.Context, vpcID string) (*VPC, *http.Response, error)
	Update(ctx context.Context, vpcID string, description string) error
	Delete(ctx context.Context, vpcID string) error
	List(ctx context.Context, options *ListOptions) ([]VPC, *Meta, *http.Response, error)
	CreateNATGateway(ctx context.Context, vpcID string, createReq *NATGatewayReq) (*NATGateway, *http.Response, error)
	GetNATGateway(ctx context.Context, vpcID, gatewayID string) (*NATGateway, *http.Response, error)
	UpdateNATGateway(ctx context.Context, vpcID, gatewayID string, updateReq *NATGatewayReq) (*NATGateway, *http.Response, error)
	DeleteNATGateway(ctx context.Context, vpcID, gatewayID string) error
	ListNATGateways(ctx context.Context, vpcID string, options *ListOptions) ([]NATGateway, *Meta, *http.Response, error)
	CreateNATGatewayFirewallRule(ctx context.Context, vpcID, gatewayID string, createReq *NATGatewayFirewallRuleCreateReq) (*NATGatewayFirewallRule, *http.Response, error) //nolint:lll
	GetNATGatewayFirewallRule(ctx context.Context, vpcID, gatewayID, ruleID string) (*NATGatewayFirewallRule, *http.Response, error)
	UpdateNATGatewayFirewallRule(ctx context.Context, vpcID, gatewayID, ruleID string, updateReq *NATGatewayFirewallRuleUpdateReq) (*NATGatewayFirewallRule, *http.Response, error) //nolint:lll
	DeleteNATGatewayFirewallRule(ctx context.Context, vpcID, gatewayID, ruleID string) error
	ListNATGatewayFirewallRules(ctx context.Context, vpcID, gatewayID string, options *ListOptions) ([]NATGatewayFirewallRule, *Meta, *http.Response, error)                                    //nolint:lll
	CreateNATGatewayPortForwardingRule(ctx context.Context, vpcID, gatewayID string, createReq *NATGatewayPortForwardingRuleReq) (*NATGatewayPortForwardingRule, *http.Response, error)         //nolint:lll
	GetNATGatewayPortForwardingRule(ctx context.Context, vpcID, gatewayID, ruleID string) (*NATGatewayPortForwardingRule, *http.Response, error)                                                //nolint:lll
	UpdateNATGatewayPortForwardingRule(ctx context.Context, vpcID, gatewayID, ruleID string, updateReq *NATGatewayPortForwardingRuleReq) (*NATGatewayPortForwardingRule, *http.Response, error) //nolint:lll
	DeleteNATGatewayPortForwardingRule(ctx context.Context, vpcID, gatewayID, ruleID string) error
	ListNATGatewayPortForwardingRules(ctx context.Context, vpcID, gatewayID string, options *ListOptions) ([]NATGatewayPortForwardingRule, *Meta, *http.Response, error) //nolint:lll
}

// VPCServiceHandler handles interaction with the VPC methods for the Vultr API
type VPCServiceHandler struct {
	client *Client
}

// VPC represents a Vultr VPC
type VPC struct {
	ID           string `json:"id"`
	Region       string `json:"region"`
	Description  string `json:"description"`
	V4Subnet     string `json:"v4_subnet"`
	V4SubnetMask int    `json:"v4_subnet_mask"`
	DateCreated  string `json:"date_created"`
}

// VPCReq represents parameters to create or update a VPC resource
type VPCReq struct {
	Region       string `json:"region"`
	Description  string `json:"description"`
	V4Subnet     string `json:"v4_subnet"`
	V4SubnetMask int    `json:"v4_subnet_mask"`
}

type vpcsBase struct {
	VPCs []VPC `json:"vpcs"`
	Meta *Meta `json:"meta"`
}

type vpcBase struct {
	VPC *VPC `json:"vpc"`
}

// NATGateway represents a Vultr NAT Gateway
type NATGateway struct {
	ID          string            `json:"id"`
	VPCID       string            `json:"vpc_id"`
	DateCreated string            `json:"date_created"`
	Status      string            `json:"status"`
	Label       string            `json:"label"`
	Tag         string            `json:"tag"`
	PublicIPs   []string          `json:"public_ips"`
	PublicIPsV6 []string          `json:"public_ips_v6"`
	PrivateIPs  []string          `json:"private_ips"`
	Billing     NATGatewayBilling `json:"billing"`
}

// NATGatewayBilling represents a Vultr NAT Gateway
type NATGatewayBilling struct {
	Charges float32 `json:"charges"`
	Monthly float32 `json:"monthly"`
}

// NATGatewayReq represents parameters to create/update a NAT Gateway resource
type NATGatewayReq struct {
	Label string `json:"label,omitempty"`
	Tag   string `json:"tag,omitempty"`
}

type natGatewaysBase struct {
	NATGateways []NATGateway `json:"nat_gateways"`
	Meta        *Meta        `json:"meta"`
}

type natGatewayBase struct {
	NATGateway *NATGateway `json:"nat_gateway"`
}

// NATGatewayFirewallRule represents a firewall rule for a Vultr NAT Gateway
type NATGatewayFirewallRule struct {
	ID         string `json:"id"`
	Action     string `json:"action"`
	Protocol   string `json:"protocol"`
	Port       string `json:"port"`
	Subnet     string `json:"subnet"`
	SubnetSize int    `json:"subnet_size"`
	Notes      string `json:"notes"`
}

// NATGatewayFirewallRuleCreateReq represents parameters to create a NAT Gateway firewall rule resource
type NATGatewayFirewallRuleCreateReq struct {
	Protocol   string `json:"protocol,omitempty"`
	Port       string `json:"port,omitempty"`
	Subnet     string `json:"subnet,omitempty"`
	SubnetSize int    `json:"subnet_size,omitempty"`
	Notes      string `json:"notes,omitempty"`
}

// NATGatewayFirewallRuleUpdateReq represents parameters to update a NAT Gateway firewall rule resource
type NATGatewayFirewallRuleUpdateReq struct {
	Notes string `json:"notes,omitempty"`
}

type natGatewayFirewallRulesBase struct {
	FirewallRules []NATGatewayFirewallRule `json:"firewall_rules"`
	Meta          *Meta                    `json:"meta"`
}

type natGatewayFirewallRuleBase struct {
	FirewallRule *NATGatewayFirewallRule `json:"firewall_rule"`
}

// NATGatewayPortForwardingRule represents a port forwarding rule for a Vultr NAT Gateway
type NATGatewayPortForwardingRule struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Protocol     string `json:"protocol"`
	ExternalPort int    `json:"external_port"`
	InternalIP   string `json:"internal_ip"`
	InternalPort int    `json:"internal_port"`
	Enabled      *bool  `json:"enabled"`
	Description  string `json:"description"`
	DateCreated  string `json:"created_at"`
	DateUpdated  string `json:"updated_at"`
}

// NATGatewayPortForwardingRuleReq represents parameters to create or update a NAT Gateway port forwarding rule resource
type NATGatewayPortForwardingRuleReq struct {
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	InternalIP   string `json:"internal_ip,omitempty"`
	Protocol     string `json:"protocol,omitempty"`
	ExternalPort int    `json:"external_port,omitempty"`
	InternalPort int    `json:"internal_port,omitempty"`
	Enabled      *bool  `json:"enabled,omitempty"`
}

type natGatewayPortForwardingRulesBase struct {
	PortForwardingRules []NATGatewayPortForwardingRule `json:"port_forwarding_rules"`
	Meta                *Meta                          `json:"meta"`
}

type natGatewayPortForwardingRuleBase struct {
	PortForwardingRule *NATGatewayPortForwardingRule `json:"port_forwarding_rule"`
}

// Create creates a new VPC. A VPC can only be used at the location for which it was created.
func (n *VPCServiceHandler) Create(ctx context.Context, createReq *VPCReq) (*VPC, *http.Response, error) {
	req, err := n.client.NewRequest(ctx, http.MethodPost, vpcPath, createReq)
	if err != nil {
		return nil, nil, err
	}

	vpc := new(vpcBase)
	resp, err := n.client.DoWithContext(ctx, req, vpc)
	if err != nil {
		return nil, resp, err
	}

	return vpc.VPC, resp, nil
}

// Get gets the VPC of the requested ID
func (n *VPCServiceHandler) Get(ctx context.Context, vpcID string) (*VPC, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", vpcPath, vpcID)
	req, err := n.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	vpc := new(vpcBase)
	resp, err := n.client.DoWithContext(ctx, req, vpc)
	if err != nil {
		return nil, resp, err
	}

	return vpc.VPC, resp, nil
}

// Update updates a VPC
func (n *VPCServiceHandler) Update(ctx context.Context, vpcID, description string) error {
	uri := fmt.Sprintf("%s/%s", vpcPath, vpcID)

	vpcReq := RequestBody{"description": description}
	req, err := n.client.NewRequest(ctx, http.MethodPut, uri, vpcReq)
	if err != nil {
		return err
	}

	_, err = n.client.DoWithContext(ctx, req, nil)
	return err
}

// Delete deletes a VPC. Before deleting, a VPC must be disabled from all instances
func (n *VPCServiceHandler) Delete(ctx context.Context, vpcID string) error {
	uri := fmt.Sprintf("%s/%s", vpcPath, vpcID)
	req, err := n.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	_, err = n.client.DoWithContext(ctx, req, nil)
	return err
}

// List lists all VPCs on the current account
func (n *VPCServiceHandler) List(ctx context.Context, options *ListOptions) ([]VPC, *Meta, *http.Response, error) { //nolint:dupl
	req, err := n.client.NewRequest(ctx, http.MethodGet, vpcPath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	vpcs := new(vpcsBase)
	resp, err := n.client.DoWithContext(ctx, req, vpcs)
	if err != nil {
		return nil, nil, resp, err
	}

	return vpcs.VPCs, vpcs.Meta, resp, nil
}

// CreateNATGateway creates a new NAT Gateway under the given VPC network
func (n *VPCServiceHandler) CreateNATGateway(ctx context.Context, vpcID string, createReq *NATGatewayReq) (*NATGateway, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/nat-gateway", vpcPath, vpcID)
	req, err := n.client.NewRequest(ctx, http.MethodPost, uri, createReq)
	if err != nil {
		return nil, nil, err
	}

	natGateway := new(natGatewayBase)
	resp, err := n.client.DoWithContext(ctx, req, natGateway)
	if err != nil {
		return nil, resp, err
	}

	return natGateway.NATGateway, resp, nil
}

// GetNATGateway gets the NAT Gateway of the requested ID
func (n *VPCServiceHandler) GetNATGateway(ctx context.Context, vpcID, gatewayID string) (*NATGateway, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/nat-gateway/%s", vpcPath, vpcID, gatewayID)
	req, err := n.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	natGateway := new(natGatewayBase)
	resp, err := n.client.DoWithContext(ctx, req, natGateway)
	if err != nil {
		return nil, resp, err
	}

	return natGateway.NATGateway, resp, nil
}

// UpdateNATGateway updates a NAT Gateway
func (n *VPCServiceHandler) UpdateNATGateway(ctx context.Context, vpcID, gatewayID string, updateReq *NATGatewayReq) (*NATGateway, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/nat-gateway/%s", vpcPath, vpcID, gatewayID)
	req, err := n.client.NewRequest(ctx, http.MethodPut, uri, updateReq)
	if err != nil {
		return nil, nil, err
	}

	natGateway := new(natGatewayBase)
	resp, err := n.client.DoWithContext(ctx, req, natGateway)
	if err != nil {
		return nil, nil, err
	}

	return natGateway.NATGateway, resp, nil
}

// DeleteNATGateway deletes a NAT Gateway
func (n *VPCServiceHandler) DeleteNATGateway(ctx context.Context, vpcID, gatewayID string) error {
	uri := fmt.Sprintf("%s/%s/nat-gateway/%s", vpcPath, vpcID, gatewayID)
	req, err := n.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	_, err = n.client.DoWithContext(ctx, req, nil)
	return err
}

// ListNATGateways lists all NAT Gateways on the current VPC network
func (n *VPCServiceHandler) ListNATGateways(ctx context.Context, vpcID string, options *ListOptions) ([]NATGateway, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/nat-gateway", vpcPath, vpcID)
	req, err := n.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	natGateways := new(natGatewaysBase)
	resp, err := n.client.DoWithContext(ctx, req, natGateways)
	if err != nil {
		return nil, nil, resp, err
	}

	return natGateways.NATGateways, natGateways.Meta, resp, nil
}

// CreateNATGatewayFirewallRule creates a new firewall rule on the current NAT Gateway
func (n *VPCServiceHandler) CreateNATGatewayFirewallRule(ctx context.Context, vpcID, gatewayID string, createReq *NATGatewayFirewallRuleCreateReq) (*NATGatewayFirewallRule, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/nat-gateway/%s/global/firewall-rules", vpcPath, vpcID, gatewayID)
	req, err := n.client.NewRequest(ctx, http.MethodPost, uri, createReq)
	if err != nil {
		return nil, nil, err
	}

	natGatewayFirewallRule := new(natGatewayFirewallRuleBase)
	resp, err := n.client.DoWithContext(ctx, req, natGatewayFirewallRule)
	if err != nil {
		return nil, resp, err
	}

	return natGatewayFirewallRule.FirewallRule, resp, nil
}

// GetNATGatewayFirewallRule gets a firewall rule of the requested ID on the current NAT Gateway
func (n *VPCServiceHandler) GetNATGatewayFirewallRule(ctx context.Context, vpcID, gatewayID, ruleID string) (*NATGatewayFirewallRule, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/nat-gateway/%s/global/firewall-rules/%s", vpcPath, vpcID, gatewayID, ruleID)
	req, err := n.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	natGatewayFirewallRule := new(natGatewayFirewallRuleBase)
	resp, err := n.client.DoWithContext(ctx, req, natGatewayFirewallRule)
	if err != nil {
		return nil, resp, err
	}

	return natGatewayFirewallRule.FirewallRule, resp, nil
}

// UpdateNATGatewayFirewallRule updates a firewall rule on the current NAT Gateway
func (n *VPCServiceHandler) UpdateNATGatewayFirewallRule(ctx context.Context, vpcID, gatewayID, ruleID string, updateReq *NATGatewayFirewallRuleUpdateReq) (*NATGatewayFirewallRule, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/nat-gateway/%s/global/firewall-rules/%s", vpcPath, vpcID, gatewayID, ruleID)
	req, err := n.client.NewRequest(ctx, http.MethodPut, uri, updateReq)
	if err != nil {
		return nil, nil, err
	}

	natGatewayFirewallRule := new(natGatewayFirewallRuleBase)
	resp, err := n.client.DoWithContext(ctx, req, natGatewayFirewallRule)
	if err != nil {
		return nil, nil, err
	}

	return natGatewayFirewallRule.FirewallRule, resp, nil
}

// DeleteNATGatewayFirewallRule deletes a firewall rule from the current NAT Gateway
func (n *VPCServiceHandler) DeleteNATGatewayFirewallRule(ctx context.Context, vpcID, gatewayID, ruleID string) error {
	uri := fmt.Sprintf("%s/%s/nat-gateway/%s/global/firewall-rules/%s", vpcPath, vpcID, gatewayID, ruleID)
	req, err := n.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	_, err = n.client.DoWithContext(ctx, req, nil)
	return err
}

// ListNATGatewayFirewallRules lists all firewall rules on the current NAT Gateway
func (n *VPCServiceHandler) ListNATGatewayFirewallRules(ctx context.Context, vpcID, gatewayID string, options *ListOptions) ([]NATGatewayFirewallRule, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/nat-gateway/%s/global/firewall-rules", vpcPath, vpcID, gatewayID)
	req, err := n.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	natGatewayFirewallRules := new(natGatewayFirewallRulesBase)
	resp, err := n.client.DoWithContext(ctx, req, natGatewayFirewallRules)
	if err != nil {
		return nil, nil, resp, err
	}

	return natGatewayFirewallRules.FirewallRules, natGatewayFirewallRules.Meta, resp, nil
}

// CreateNATGatewayPortForwardingRule creates a new port forwarding rule on the current NAT Gateway
func (n *VPCServiceHandler) CreateNATGatewayPortForwardingRule(ctx context.Context, vpcID, gatewayID string, createReq *NATGatewayPortForwardingRuleReq) (*NATGatewayPortForwardingRule, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/nat-gateway/%s/global/port-forwarding-rules", vpcPath, vpcID, gatewayID)
	req, err := n.client.NewRequest(ctx, http.MethodPost, uri, createReq)
	if err != nil {
		return nil, nil, err
	}

	natGatewayPortForwardingRule := new(natGatewayPortForwardingRuleBase)
	resp, err := n.client.DoWithContext(ctx, req, natGatewayPortForwardingRule)
	if err != nil {
		return nil, resp, err
	}

	return natGatewayPortForwardingRule.PortForwardingRule, resp, nil
}

// GetNATGatewayPortForwardingRule gets a port forwarding rule of the requested ID on the current NAT Gateway
func (n *VPCServiceHandler) GetNATGatewayPortForwardingRule(ctx context.Context, vpcID, gatewayID, ruleID string) (*NATGatewayPortForwardingRule, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/nat-gateway/%s/global/port-forwarding-rules/%s", vpcPath, vpcID, gatewayID, ruleID)
	req, err := n.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	natGatewayPortForwardingRule := new(natGatewayPortForwardingRuleBase)
	resp, err := n.client.DoWithContext(ctx, req, natGatewayPortForwardingRule)
	if err != nil {
		return nil, resp, err
	}

	return natGatewayPortForwardingRule.PortForwardingRule, resp, nil
}

// UpdateNATGatewayPortForwardingRule updates a port forwarding rule on the current NAT Gateway
func (n *VPCServiceHandler) UpdateNATGatewayPortForwardingRule(ctx context.Context, vpcID, gatewayID, ruleID string, updateReq *NATGatewayPortForwardingRuleReq) (*NATGatewayPortForwardingRule, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/nat-gateway/%s/global/port-forwarding-rules/%s", vpcPath, vpcID, gatewayID, ruleID)
	req, err := n.client.NewRequest(ctx, http.MethodPut, uri, updateReq)
	if err != nil {
		return nil, nil, err
	}

	natGatewayPortForwardingRule := new(natGatewayPortForwardingRuleBase)
	resp, err := n.client.DoWithContext(ctx, req, natGatewayPortForwardingRule)
	if err != nil {
		return nil, nil, err
	}

	return natGatewayPortForwardingRule.PortForwardingRule, resp, nil
}

// DeleteNATGatewayPortForwardingRule deletes a port forwarding rule from the current NAT Gateway
func (n *VPCServiceHandler) DeleteNATGatewayPortForwardingRule(ctx context.Context, vpcID, gatewayID, ruleID string) error {
	uri := fmt.Sprintf("%s/%s/nat-gateway/%s/global/port-forwarding-rules/%s", vpcPath, vpcID, gatewayID, ruleID)
	req, err := n.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	_, err = n.client.DoWithContext(ctx, req, nil)
	return err
}

// ListNATGatewayPortForwardingRules lists all port forwarding rules on the current NAT Gateway
func (n *VPCServiceHandler) ListNATGatewayPortForwardingRules(ctx context.Context, vpcID, gatewayID string, options *ListOptions) ([]NATGatewayPortForwardingRule, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/nat-gateway/%s/global/port-forwarding-rules", vpcPath, vpcID, gatewayID)
	req, err := n.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	natGatewayPortForwardingRules := new(natGatewayPortForwardingRulesBase)
	resp, err := n.client.DoWithContext(ctx, req, natGatewayPortForwardingRules)
	if err != nil {
		return nil, nil, resp, err
	}

	return natGatewayPortForwardingRules.PortForwardingRules, natGatewayPortForwardingRules.Meta, resp, nil
}
