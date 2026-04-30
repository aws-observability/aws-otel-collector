package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// FirewallGroupService is the interface to interact with the firewall group endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/firewall
type FirewallGroupService interface { //nolint:dupl
	Create(ctx context.Context, fwGroupReq *FirewallGroupReq) (*FirewallGroup, *http.Response, error)
	Get(ctx context.Context, groupID string) (*FirewallGroup, *http.Response, error)
	Update(ctx context.Context, fwGroupID string, fwGroupReq *FirewallGroupReq) error
	Delete(ctx context.Context, fwGroupID string) error
	List(ctx context.Context, options *ListOptions) ([]FirewallGroup, *Meta, *http.Response, error)
}

// FireWallGroupServiceHandler handles interaction with the firewall group methods for the Vultr API
type FireWallGroupServiceHandler struct {
	client *Client
}

// FirewallGroup represents a Vultr firewall group
type FirewallGroup struct {
	ID            string `json:"id"`
	Description   string `json:"description"`
	DateCreated   string `json:"date_created"`
	DateModified  string `json:"date_modified"`
	InstanceCount int    `json:"instance_count"`
	RuleCount     int    `json:"rule_count"`
	MaxRuleCount  int    `json:"max_rule_count"`
}

// FirewallGroupReq struct is used to create and update a Firewall Group.
type FirewallGroupReq struct {
	Description string `json:"description"`
}

type firewallGroupsBase struct {
	FirewallGroups []FirewallGroup `json:"firewall_groups"`
	Meta           *Meta           `json:"meta"`
}

type firewallGroupBase struct {
	FirewallGroup *FirewallGroup `json:"firewall_group"`
}

// Create will create a new firewall group on your Vultr account
func (f *FireWallGroupServiceHandler) Create(ctx context.Context, fwGroupReq *FirewallGroupReq) (*FirewallGroup, *http.Response, error) {
	uri := "/v2/firewalls"

	req, err := f.client.NewRequest(ctx, http.MethodPost, uri, fwGroupReq)
	if err != nil {
		return nil, nil, err
	}

	firewall := new(firewallGroupBase)
	resp, err := f.client.DoWithContext(ctx, req, firewall)
	if err != nil {
		return nil, resp, err
	}

	return firewall.FirewallGroup, resp, nil
}

// Get will return a firewall group based on provided groupID from your Vultr account
func (f *FireWallGroupServiceHandler) Get(ctx context.Context, fwGroupID string) (*FirewallGroup, *http.Response, error) {
	uri := fmt.Sprintf("/v2/firewalls/%s", fwGroupID)

	req, err := f.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	firewall := new(firewallGroupBase)
	resp, err := f.client.DoWithContext(ctx, req, firewall)
	if err != nil {
		return nil, resp, err
	}

	return firewall.FirewallGroup, resp, nil
}

// Update will change the description of a firewall group
func (f *FireWallGroupServiceHandler) Update(ctx context.Context, fwGroupID string, fwGroupReq *FirewallGroupReq) error {
	uri := fmt.Sprintf("/v2/firewalls/%s", fwGroupID)

	req, err := f.client.NewRequest(ctx, http.MethodPut, uri, fwGroupReq)
	if err != nil {
		return err
	}

	_, err = f.client.DoWithContext(ctx, req, nil)
	return err
}

// Delete will delete a firewall group from your Vultr account
func (f *FireWallGroupServiceHandler) Delete(ctx context.Context, fwGroupID string) error {
	uri := fmt.Sprintf("/v2/firewalls/%s", fwGroupID)

	req, err := f.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	_, err = f.client.DoWithContext(ctx, req, nil)
	return err
}

// List will return a list of  all firewall groups on your Vultr account
func (f *FireWallGroupServiceHandler) List(ctx context.Context, options *ListOptions) ([]FirewallGroup, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := "/v2/firewalls"

	req, err := f.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	firewalls := new(firewallGroupsBase)
	resp, err := f.client.DoWithContext(ctx, req, firewalls)
	if err != nil {
		return nil, nil, resp, err
	}

	return firewalls.FirewallGroups, firewalls.Meta, resp, nil
}
