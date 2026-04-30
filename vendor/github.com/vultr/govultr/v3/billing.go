package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// BillingService is the interface to interact with the billing endpoint on the Vultr API
// Link : https://www.vultr.com/api/#tag/billing
type BillingService interface {
	ListHistory(ctx context.Context, options *ListOptions) ([]History, *Meta, *http.Response, error)
	ListInvoices(ctx context.Context, options *ListOptions) ([]Invoice, *Meta, *http.Response, error)
	GetInvoice(ctx context.Context, invoiceID string) (*Invoice, *http.Response, error)
	ListInvoiceItems(ctx context.Context, invoiceID int, options *ListOptions) ([]InvoiceItem, *Meta, *http.Response, error)
	ListPendingCharges(ctx context.Context, options *ListOptions) ([]InvoiceItem, *http.Response, error)
}

// BillingServiceHandler handles interaction with the billing methods for the Vultr API
type BillingServiceHandler struct {
	client *Client
}

// History represents a billing history item on an account
type History struct {
	ID          int     `json:"id"`
	Date        string  `json:"date"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Amount      float32 `json:"amount"`
	Balance     float32 `json:"balance"`
}

// Invoice represents an invoice on an account
type Invoice struct {
	ID          int     `json:"id"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Amount      float32 `json:"amount"`
	Balance     float32 `json:"balance"`
}

// InvoiceItem represents an item on an accounts invoice
type InvoiceItem struct {
	Description string  `json:"description"`
	Product     string  `json:"product"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	Units       int     `json:"units"`
	UnitType    string  `json:"unit_type"`
	UnitPrice   float32 `json:"unit_price"`
	Total       float32 `json:"total"`
}

type billingHistoryBase struct {
	History []History `json:"billing_history"`
	Meta    *Meta     `json:"meta"`
}

type invoicesBase struct {
	Invoice []Invoice `json:"billing_invoices"`
	Meta    *Meta     `json:"meta"`
}

type invoiceBase struct {
	Invoice *Invoice `json:"billing_invoice"`
}

type invoiceItemsBase struct {
	InvoiceItems []InvoiceItem `json:"invoice_items"`
	Meta         *Meta         `json:"meta"`
}

type pendingChargesBase struct {
	PendingCharges []InvoiceItem `json:"pending_charges"`
}

// ListHistory retrieves a list of all billing history on the current account
func (b *BillingServiceHandler) ListHistory(ctx context.Context, options *ListOptions) ([]History, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := "/v2/billing/history"
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	invoices := new(billingHistoryBase)
	resp, err := b.client.DoWithContext(ctx, req, invoices)
	if err != nil {
		return nil, nil, resp, err
	}

	return invoices.History, invoices.Meta, resp, nil
}

// ListInvoices retrieves a list of all billing invoices on the current account
func (b *BillingServiceHandler) ListInvoices(ctx context.Context, options *ListOptions) ([]Invoice, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := "/v2/billing/invoices"
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	invoices := new(invoicesBase)
	resp, err := b.client.DoWithContext(ctx, req, invoices)
	if err != nil {
		return nil, nil, resp, err
	}

	return invoices.Invoice, invoices.Meta, resp, nil
}

// GetInvoice retrieves an invoice that matches the given invoiceID
func (b *BillingServiceHandler) GetInvoice(ctx context.Context, invoiceID string) (*Invoice, *http.Response, error) {
	uri := fmt.Sprintf("/v2/billing/invoices/%s", invoiceID)
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, nil, err
	}

	invoice := new(invoiceBase)
	resp, err := b.client.DoWithContext(ctx, req, invoice)
	if err != nil {
		return nil, resp, err
	}

	return invoice.Invoice, resp, nil
}

// ListInvoiceItems retrieves items in an invoice that matches the given invoiceID
func (b *BillingServiceHandler) ListInvoiceItems(ctx context.Context, invoiceID int, options *ListOptions) ([]InvoiceItem, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("/v2/billing/invoices/%d/items", invoiceID)
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	invoice := new(invoiceItemsBase)
	resp, err := b.client.DoWithContext(ctx, req, invoice)
	if err != nil {
		return nil, nil, resp, err
	}

	return invoice.InvoiceItems, invoice.Meta, resp, nil
}

// ListPendingCharges retrieves a list of all pending charges on the current account
func (b *BillingServiceHandler) ListPendingCharges(ctx context.Context, options *ListOptions) ([]InvoiceItem, *http.Response, error) {
	uri := "/v2/billing/pending-charges"
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	invoice := new(pendingChargesBase)
	resp, err := b.client.DoWithContext(ctx, req, invoice)
	if err != nil {
		return nil, resp, err
	}

	return invoice.PendingCharges, resp, nil
}
