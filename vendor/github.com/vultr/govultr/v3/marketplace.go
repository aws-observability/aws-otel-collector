package govultr

import (
	"context"
	"fmt"
	"net/http"
)

const marketplacePath = "/v2/marketplace"

// MarketplaceService is the interface to interact with the Marketplace endpoints on the Vultr API
// Link: https://www.vultr.com/api/#tag/marketplace
type MarketplaceService interface {
	ListAppVariables(ctx context.Context, imageID string) ([]MarketplaceAppVariable, *http.Response, error)
}

// MarketplaceServiceHandler handles interaction with the server methods for the Vultr API
type MarketplaceServiceHandler struct {
	client *Client
}

// MarketplaceAppVariable represents a user-supplied variable for a Marketplace app
type MarketplaceAppVariable struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    *bool  `json:"required"`
}

// marketplaceAppVariablesBase holds the API response for retrieving a list of user-supplied variables for a Marketplace app
type marketplaceAppVariablesBase struct {
	MarketplaceAppVariables []MarketplaceAppVariable `json:"variables"`
}

// ListAppVariables retrieves all user-supplied variables for a Marketplace app
func (d *MarketplaceServiceHandler) ListAppVariables(ctx context.Context, imageID string) ([]MarketplaceAppVariable, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/apps/%s/variables", marketplacePath, imageID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	marketplaceAppVariables := new(marketplaceAppVariablesBase)
	resp, err := d.client.DoWithContext(ctx, req, marketplaceAppVariables)
	if err != nil {
		return nil, nil, err
	}

	return marketplaceAppVariables.MarketplaceAppVariables, resp, nil
}
