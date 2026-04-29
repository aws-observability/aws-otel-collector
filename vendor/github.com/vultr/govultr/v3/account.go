package govultr

import (
	"context"
	"net/http"
)

// AccountService is the interface to interact with Accounts endpoint on the Vultr API
// Link : https://www.vultr.com/api/#tag/account
type AccountService interface {
	Get(ctx context.Context) (*Account, *http.Response, error)
	GetBandwidth(ctx context.Context) (*AccountBandwidth, *http.Response, error)
}

// AccountServiceHandler handles interaction with the account methods for the Vultr API
type AccountServiceHandler struct {
	client *Client
}

type accountBase struct {
	Account *Account `json:"account"`
}

// Account represents a Vultr account
type Account struct {
	Balance           float32  `json:"balance"`
	PendingCharges    float32  `json:"pending_charges"`
	LastPaymentDate   string   `json:"last_payment_date"`
	LastPaymentAmount float32  `json:"last_payment_amount"`
	Name              string   `json:"name"`
	Email             string   `json:"email"`
	ACL               []string `json:"acls"`
}

type accountBandwidthBase struct {
	Bandwidth *AccountBandwidth `json:"bandwidth"`
}

// Bandwidth represents a Vultr account bandwidth
type AccountBandwidth struct {
	PreviousMonth         AccountBandwidthPeriod `json:"previous_month"`
	CurrentMonthToDate    AccountBandwidthPeriod `json:"current_month_to_date"`
	CurrentMonthProjected AccountBandwidthPeriod `json:"current_month_projected"`
}

// AccountBandwidthPeriod represents a Vultr account bandwidth period
type AccountBandwidthPeriod struct {
	TimestampStart            string  `json:"timestamp_start"`
	TimestampEnd              string  `json:"timestamp_end"`
	GBIn                      int     `json:"gb_in"`
	GBOut                     int     `json:"gb_out"`
	TotalInstanceHours        int     `json:"total_instance_hours"`
	TotalInstanceCount        int     `json:"total_instance_count"`
	InstanceBandwidthCredits  int     `json:"instance_bandwidth_credits"`
	FreeBandwidthCredits      int     `json:"free_bandwidth_credits"`
	PurchasedBandwidthCredits int     `json:"purchased_bandwidth_credits"`
	Overage                   float32 `json:"overage"`
	OverageUnitCost           float32 `json:"overage_unit_cost"`
	OverageCost               float32 `json:"overage_cost"`
}

// Get Vultr account info
func (a *AccountServiceHandler) Get(ctx context.Context) (*Account, *http.Response, error) {
	uri := "/v2/account"
	req, err := a.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	account := new(accountBase)
	resp, err := a.client.DoWithContext(ctx, req, account)
	if err != nil {
		return nil, resp, err
	}

	return account.Account, resp, nil
}

// Get Vultr account bandwidth info
func (a *AccountServiceHandler) GetBandwidth(ctx context.Context) (*AccountBandwidth, *http.Response, error) {
	uri := "/v2/account/bandwidth"
	req, err := a.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	accountBandwidth := new(accountBandwidthBase)
	resp, err := a.client.DoWithContext(ctx, req, accountBandwidth)
	if err != nil {
		return nil, resp, err
	}

	return accountBandwidth.Bandwidth, resp, nil
}
