package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const lbPath = "/v2/load-balancers"

// LoadBalancerService is the interface to interact with the server endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/load-balancer
type LoadBalancerService interface {
	Create(ctx context.Context, createReq *LoadBalancerReq) (*LoadBalancer, *http.Response, error)
	Get(ctx context.Context, lbID string) (*LoadBalancer, *http.Response, error)
	Update(ctx context.Context, lbID string, updateReq *LoadBalancerReq) error
	Delete(ctx context.Context, lbID string) error
	DeleteSSL(ctx context.Context, lbID string) error
	DeleteAutoSSL(ctx context.Context, lbID string) error
	List(ctx context.Context, options *ListOptions) ([]LoadBalancer, *Meta, *http.Response, error)
	CreateForwardingRule(ctx context.Context, lbID string, rule *ForwardingRule) (*ForwardingRule, *http.Response, error)
	GetForwardingRule(ctx context.Context, lbID string, ruleID string) (*ForwardingRule, *http.Response, error)
	DeleteForwardingRule(ctx context.Context, lbID string, RuleID string) error
	ListForwardingRules(ctx context.Context, lbID string, options *ListOptions) ([]ForwardingRule, *Meta, *http.Response, error)
	ListFirewallRules(ctx context.Context, lbID string, options *ListOptions) ([]LBFirewallRule, *Meta, *http.Response, error)
	GetFirewallRule(ctx context.Context, lbID string, ruleID string) (*LBFirewallRule, *http.Response, error)
}

// LoadBalancerHandler handles interaction with the server methods for the Vultr API
type LoadBalancerHandler struct {
	client *Client
}

// LoadBalancer represent the structure of a load balancer
type LoadBalancer struct {
	ID              string           `json:"id,omitempty"`
	DateCreated     string           `json:"date_created,omitempty"`
	Region          string           `json:"region,omitempty"`
	Label           string           `json:"label,omitempty"`
	Status          string           `json:"status,omitempty"`
	IPV4            string           `json:"ipv4,omitempty"`
	IPV6            string           `json:"ipv6,omitempty"`
	Instances       []string         `json:"instances,omitempty"`
	Nodes           int              `json:"nodes,omitempty"`
	HealthCheck     *HealthCheck     `json:"health_check,omitempty"`
	GenericInfo     *GenericInfo     `json:"generic_info,omitempty"`
	SSLInfo         *bool            `json:"has_ssl,omitempty"`
	AutoSSL         *AutoSSL         `json:"auto_ssl,omitempty"`
	HTTP2           *bool            `json:"http2,omitempty"`
	HTTP3           *bool            `json:"http3,omitempty"`
	ForwardingRules []ForwardingRule `json:"forwarding_rules,omitempty"`
	FirewallRules   []LBFirewallRule `json:"firewall_rules,omitempty"`
	GlobalRegions   []string         `json:"global_regions,omitempty"`
}

// LoadBalancerReq gives options for creating or updating a load balancer
type LoadBalancerReq struct {
	Region             string           `json:"region,omitempty"`
	Label              string           `json:"label,omitempty"`
	Instances          []string         `json:"instances,omitempty"`
	Nodes              int              `json:"nodes,omitempty"`
	HealthCheck        *HealthCheck     `json:"health_check,omitempty"`
	StickySessions     *StickySessions  `json:"sticky_session,omitempty"`
	ForwardingRules    []ForwardingRule `json:"forwarding_rules,omitempty"`
	SSL                *SSL             `json:"ssl,omitempty"`
	AutoSSL            *AutoSSL         `json:"auto_ssl,omitempty"`
	SSLRedirect        *bool            `json:"ssl_redirect,omitempty"`
	HTTP2              *bool            `json:"http2,omitempty"`
	HTTP3              *bool            `json:"http3,omitempty"`
	ProxyProtocol      *bool            `json:"proxy_protocol,omitempty"`
	BalancingAlgorithm string           `json:"balancing_algorithm,omitempty"`
	FirewallRules      []LBFirewallRule `json:"firewall_rules,omitempty"`
	Timeout            int              `json:"timeout,omitempty"`
	VPC                *string          `json:"vpc,omitempty"`
	GlobalRegions      []string         `json:"global_regions,omitempty"`
}

// InstanceList represents instances that are attached to your load balancer
type InstanceList struct {
	InstanceList []string
}

// HealthCheck represents your health check configuration for your load balancer.
type HealthCheck struct {
	Protocol           string `json:"protocol,omitempty"`
	Port               int    `json:"port,omitempty"`
	Path               string `json:"path,omitempty"`
	CheckInterval      int    `json:"check_interval,omitempty"`
	ResponseTimeout    int    `json:"response_timeout,omitempty"`
	UnhealthyThreshold int    `json:"unhealthy_threshold,omitempty"`
	HealthyThreshold   int    `json:"healthy_threshold,omitempty"`
}

// GenericInfo represents generic configuration of your load balancer
type GenericInfo struct {
	BalancingAlgorithm string          `json:"balancing_algorithm,omitempty"`
	Timeout            int             `json:"timeout,omitempty"`
	SSLRedirect        *bool           `json:"ssl_redirect,omitempty"`
	StickySessions     *StickySessions `json:"sticky_sessions,omitempty"`
	ProxyProtocol      *bool           `json:"proxy_protocol,omitempty"`
	VPC                string          `json:"vpc,omitempty"`
}

// StickySessions represents cookie for your load balancer
type StickySessions struct {
	CookieName string `json:"cookie_name,omitempty"`
}

// ForwardingRules represent a list of forwarding rules
type ForwardingRules struct {
	ForwardRuleList []ForwardingRule `json:"forwarding_rules,omitempty"`
}

// ForwardingRule represent a single forwarding rule
type ForwardingRule struct {
	RuleID           string `json:"id,omitempty"`
	FrontendProtocol string `json:"frontend_protocol,omitempty"`
	FrontendPort     int    `json:"frontend_port,omitempty"`
	BackendProtocol  string `json:"backend_protocol,omitempty"`
	BackendPort      int    `json:"backend_port,omitempty"`
}

// LBFirewallRule represent a single firewall rule
type LBFirewallRule struct {
	RuleID string `json:"id,omitempty"`
	Port   int    `json:"port,omitempty"`
	IPType string `json:"ip_type,omitempty"`
	Source string `json:"source,omitempty"`
}

// SSL represents valid SSL config
type SSL struct {
	PrivateKey     string `json:"private_key,omitempty"`
	Certificate    string `json:"certificate,omitempty"`
	Chain          string `json:"chain,omitempty"`
	PrivateKeyB64  string `json:"private_key_b64,omitempty"`
	CertificateB64 string `json:"certificate_b64,omitempty"`
	ChainB64       string `json:"chain_b64,omitempty"`
}

// AutoSSL represents valid AutoSSL config
type AutoSSL struct {
	DomainZone string `json:"domain_zone"`
	DomainSub  string `json:"domain_sub,omitempty"`
	Domain     string `json:"domain,omitempty"`
}

type lbsBase struct {
	LoadBalancers []LoadBalancer `json:"load_balancers"`
	Meta          *Meta          `json:"meta"`
}

type lbBase struct {
	LoadBalancer *LoadBalancer `json:"load_balancer"`
}

type lbRulesBase struct {
	ForwardingRules []ForwardingRule `json:"forwarding_rules"`
	Meta            *Meta            `json:"meta"`
}

type lbRuleBase struct {
	ForwardingRule *ForwardingRule `json:"forwarding_rule"`
}

type lbFWRulesBase struct {
	FirewallRules []LBFirewallRule `json:"firewall_rules"`
	Meta          *Meta            `json:"meta"`
}

type lbFWRuleBase struct {
	FirewallRule *LBFirewallRule `json:"firewall_rule"`
}

// Create a load balancer
func (l *LoadBalancerHandler) Create(ctx context.Context, createReq *LoadBalancerReq) (*LoadBalancer, *http.Response, error) {
	req, err := l.client.NewRequest(ctx, http.MethodPost, lbPath, createReq)
	if err != nil {
		return nil, nil, err
	}

	var lb = new(lbBase)
	resp, err := l.client.DoWithContext(ctx, req, &lb)
	if err != nil {
		return nil, resp, err
	}

	return lb.LoadBalancer, resp, nil
}

// Get a load balancer
func (l *LoadBalancerHandler) Get(ctx context.Context, lbID string) (*LoadBalancer, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", lbPath, lbID)
	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	var lb = new(lbBase)
	resp, err := l.client.DoWithContext(ctx, req, lb)
	if err != nil {
		return nil, resp, err
	}

	return lb.LoadBalancer, resp, nil
}

// Update updates your your load balancer
func (l *LoadBalancerHandler) Update(ctx context.Context, lbID string, updateReq *LoadBalancerReq) error {
	uri := fmt.Sprintf("%s/%s", lbPath, lbID)
	req, err := l.client.NewRequest(ctx, http.MethodPatch, uri, updateReq)
	if err != nil {
		return err
	}

	_, err = l.client.DoWithContext(ctx, req, nil)
	return err
}

// Delete a load balancer subscription.
func (l *LoadBalancerHandler) Delete(ctx context.Context, lbID string) error {
	uri := fmt.Sprintf("%s/%s", lbPath, lbID)
	req, err := l.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	_, err = l.client.DoWithContext(ctx, req, nil)
	return err
}

// List all load balancer subscriptions on the current account.
func (l *LoadBalancerHandler) List(ctx context.Context, options *ListOptions) ([]LoadBalancer, *Meta, *http.Response, error) { //nolint:dupl
	req, err := l.client.NewRequest(ctx, http.MethodGet, lbPath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	lbs := new(lbsBase)
	resp, err := l.client.DoWithContext(ctx, req, &lbs)
	if err != nil {
		return nil, nil, resp, err
	}

	return lbs.LoadBalancers, lbs.Meta, resp, nil
}

// CreateForwardingRule will create a new forwarding rule for your load balancer subscription.
// Note the RuleID will be returned in the ForwardingRule struct
func (l *LoadBalancerHandler) CreateForwardingRule(ctx context.Context, lbID string, rule *ForwardingRule) (*ForwardingRule, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/forwarding-rules", lbPath, lbID)
	req, err := l.client.NewRequest(ctx, http.MethodPost, uri, rule)
	if err != nil {
		return nil, nil, err
	}

	fwRule := new(lbRuleBase)
	resp, err := l.client.DoWithContext(ctx, req, fwRule)
	if err != nil {
		return nil, resp, err
	}

	return fwRule.ForwardingRule, resp, nil
}

// GetForwardingRule will get a forwarding rule from your load balancer subscription.
func (l *LoadBalancerHandler) GetForwardingRule(ctx context.Context, lbID, ruleID string) (*ForwardingRule, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/forwarding-rules/%s", lbPath, lbID, ruleID)
	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	fwRule := new(lbRuleBase)
	resp, err := l.client.DoWithContext(ctx, req, fwRule)
	if err != nil {
		return nil, resp, err
	}

	return fwRule.ForwardingRule, resp, nil
}

// ListForwardingRules lists all forwarding rules for a load balancer subscription
func (l *LoadBalancerHandler) ListForwardingRules(ctx context.Context, lbID string, options *ListOptions) ([]ForwardingRule, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/forwarding-rules", lbPath, lbID)
	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	fwRules := new(lbRulesBase)
	resp, err := l.client.DoWithContext(ctx, req, &fwRules)
	if err != nil {
		return nil, nil, resp, err
	}

	return fwRules.ForwardingRules, fwRules.Meta, resp, nil
}

// DeleteForwardingRule removes a forwarding rule from a load balancer subscription
func (l *LoadBalancerHandler) DeleteForwardingRule(ctx context.Context, lbID, ruleID string) error {
	uri := fmt.Sprintf("%s/%s/forwarding-rules/%s", lbPath, lbID, ruleID)
	req, err := l.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = l.client.DoWithContext(ctx, req, nil)
	return err
}

// GetFirewallRule will get a firewall rule from your load balancer subscription.
func (l *LoadBalancerHandler) GetFirewallRule(ctx context.Context, lbID, ruleID string) (*LBFirewallRule, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/firewall-rules/%s", lbPath, lbID, ruleID)
	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	fwRule := new(lbFWRuleBase)
	resp, err := l.client.DoWithContext(ctx, req, fwRule)
	if err != nil {
		return nil, resp, err
	}

	return fwRule.FirewallRule, resp, nil
}

// ListFirewallRules lists all firewall rules for a load balancer subscription
func (l *LoadBalancerHandler) ListFirewallRules(ctx context.Context, lbID string, options *ListOptions) ([]LBFirewallRule, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/firewall-rules", lbPath, lbID)
	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	fwRules := new(lbFWRulesBase)
	resp, err := l.client.DoWithContext(ctx, req, &fwRules)
	if err != nil {
		return nil, nil, resp, err
	}

	return fwRules.FirewallRules, fwRules.Meta, resp, nil
}

// DeleteSSL removes the SSL configuration from a load balancer subscription.
func (l *LoadBalancerHandler) DeleteSSL(ctx context.Context, lbID string) error {
	uri := fmt.Sprintf("%s/%s/ssl", lbPath, lbID)
	req, err := l.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = l.client.DoWithContext(ctx, req, nil)
	return err
}

// DeleteAutoSSL removes the AutoSSL configuration from a load balancer subscription.
func (l *LoadBalancerHandler) DeleteAutoSSL(ctx context.Context, lbID string) error {
	uri := fmt.Sprintf("%s/%s/auto_ssl", lbPath, lbID)
	req, err := l.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = l.client.DoWithContext(ctx, req, nil)
	return err
}
