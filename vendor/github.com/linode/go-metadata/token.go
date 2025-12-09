package metadata

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type TokenOption func(opts *tokenCreateOpts)

// TokenWithExpiry configures the expiry in seconds for a token.
// Default: 3600
func TokenWithExpiry(seconds int) TokenOption {
	return func(opts *tokenCreateOpts) {
		opts.ExpirySeconds = seconds
	}
}

type tokenCreateOpts struct {
	ExpirySeconds int
}

// GenerateToken generates a token to access the Metadata API.
func (c *Client) GenerateToken(ctx context.Context, opts ...TokenOption) (string, error) {
	// Handle create options
	createOpts := tokenCreateOpts{
		ExpirySeconds: 3600,
	}

	for _, opt := range opts {
		opt(&createOpts)
	}

	req := c.R(ctx).
		SetResult(&[]string{}).
		SetHeader("Metadata-Token-Expiry-Seconds", strconv.Itoa(createOpts.ExpirySeconds))

	resp, err := req.Put("token")
	if err != nil {
		return "", err
	}

	result := *resp.Result().(*[]string)

	if len(result) < 1 {
		return "", fmt.Errorf("no token returned from API")
	}

	return result[0], nil
}

// UseToken applies the given token to this Metadata client.
func (c *Client) UseToken(token string) *Client {
	c.resty.SetHeader("Metadata-Token", token)
	return c
}

// RefreshToken generates and applies a new token for this client.
func (c *Client) RefreshToken(ctx context.Context, opts ...TokenOption) (*Client, error) {
	creationTime := time.Now()

	token, err := c.GenerateToken(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to generate metadata token: %w", err)
	}

	c.UseToken(token)

	c.managedTokenExpiry = creationTime

	return c, nil
}
