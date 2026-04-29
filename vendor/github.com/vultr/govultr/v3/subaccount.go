package govultr

import (
	"context"
	"net/http"
)

// SubAccountService is the interface to interact with sub-accounts endpoint on the Vultr API
// Link : https://www.vultr.com/api/#tag/subaccount
type SubAccountService interface {
	List(ctx context.Context, options *ListOptions) ([]SubAccount, *Meta, *http.Response, error)
	Create(ctx context.Context, saReq *SubAccountReq) (*SubAccount, *http.Response, error)
}

// SubAccountServiceHandler handles interaction with the account methods for the Vultr API
type SubAccountServiceHandler struct {
	client *Client
}

type subAccountsBase struct {
	SubAccounts []SubAccount `json:"subaccounts"`
	Meta        *Meta        `json:"meta"`
}

type subAccountBase struct {
	SubAccount *SubAccount `json:"subaccount"`
}

// SubAccount represents a Vultr sub-account
type SubAccount struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"subaccount_name"`
	OtherID        string `json:"subaccount_id"`
	Activated      bool   `json:"activated"`
	Balance        int    `json:"balance"`
	PendingCharges int    `json:"pending_charges"`
}

// SubAccountReq is the sub-account struct for create calls
type SubAccountReq struct {
	Email   string `json:"email"`
	Name    string `json:"subaccount_name,omitempty"`
	OtherID string `json:"subaccount_id,omitempty"`
}

// List all sub-accounts
func (s *SubAccountServiceHandler) List(ctx context.Context, options *ListOptions) ([]SubAccount, *Meta, *http.Response, error) {
	uri := "/v2/subaccounts"
	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	sas := new(subAccountsBase)
	resp, err := s.client.DoWithContext(ctx, req, sas)
	if err != nil {
		return nil, nil, resp, err
	}

	return sas.SubAccounts, sas.Meta, resp, nil
}

// Create a sub-account
func (s *SubAccountServiceHandler) Create(ctx context.Context, saReq *SubAccountReq) (*SubAccount, *http.Response, error) {
	uri := "/v2/subaccounts"
	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, saReq)
	if err != nil {
		return nil, nil, err
	}

	sa := new(subAccountBase)
	resp, err := s.client.DoWithContext(ctx, req, sa)
	if err != nil {
		return nil, resp, err
	}

	return sa.SubAccount, resp, nil
}
