package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// RegionService is the interface to interact with Region endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/region
type RegionService interface {
	Availability(ctx context.Context, regionID string, planType string) (*PlanAvailability, *http.Response, error)
	List(ctx context.Context, options *ListOptions) ([]Region, *Meta, *http.Response, error)
}

var _ RegionService = &RegionServiceHandler{}

// RegionServiceHandler handles interaction with the region methods for the Vultr API
type RegionServiceHandler struct {
	client *Client
}

// Region represents a Vultr region
type Region struct {
	ID           string   `json:"id"`
	City         string   `json:"city"`
	Country      string   `json:"country"`
	Continent    string   `json:"continent,omitempty"`
	Options      []string `json:"options"`
	Connectivity []string `json:"connectivity"`
}

type regionBase struct {
	Regions []Region `json:"regions"`
	Meta    *Meta
}

// PlanAvailability contains all available plans.
type PlanAvailability struct {
	AvailablePlans []string `json:"available_plans"`
}

// List returns all available regions
func (r *RegionServiceHandler) List(ctx context.Context, options *ListOptions) ([]Region, *Meta, *http.Response, error) {
	uri := "/v2/regions"

	req, err := r.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	regions := new(regionBase)
	resp, err := r.client.DoWithContext(ctx, req, &regions)
	if err != nil {
		return nil, nil, resp, err
	}

	return regions.Regions, regions.Meta, resp, nil
}

// Availability retrieves a list of the plan IDs currently available for a given location.
func (r *RegionServiceHandler) Availability(ctx context.Context, regionID, planType string) (*PlanAvailability, *http.Response, error) {
	uri := fmt.Sprintf("/v2/regions/%s/availability", regionID)

	req, err := r.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, nil, err
	}

	// Optional planType filter
	if planType != "" {
		q := req.URL.Query()
		q.Add("type", planType)
		req.URL.RawQuery = q.Encode()
	}

	plans := new(PlanAvailability)
	resp, err := r.client.DoWithContext(ctx, req, plans)
	if err != nil {
		return nil, resp, err
	}

	return plans, resp, nil
}
