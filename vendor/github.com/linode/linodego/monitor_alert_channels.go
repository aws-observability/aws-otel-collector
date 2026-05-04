package linodego

import (
	"context"
	"encoding/json"
	"time"

	"github.com/linode/linodego/internal/parseabletime"
)

type AlertNotificationType string

const (
	EmailAlertNotification AlertNotificationType = "email"
)

type AlertChannelType string

const (
	SystemAlertChannel AlertChannelType = "system"
	UserAlertChannel   AlertChannelType = "user"
)

// AlertChannel represents a Monitor Channel object.
type AlertChannel struct {
	Alerts      AlertsInfo            `json:"alerts"`
	ChannelType AlertNotificationType `json:"channel_type"`
	Content     ChannelContent        `json:"content"`
	Details     ChannelDetails        `json:"details"`
	Created     *time.Time            `json:"-"`
	CreatedBy   string                `json:"created_by"`
	Updated     *time.Time            `json:"-"`
	UpdatedBy   string                `json:"updated_by"`
	ID          int                   `json:"id"`
	Label       string                `json:"label"`
	Type        AlertChannelType      `json:"type"`
}

// AlertsInfo represents alert information for a channel
type AlertsInfo struct {
	URL        string `json:"url"`
	Type       string `json:"type"`
	AlertCount int    `json:"alert_count"`
}

type EmailChannelContent struct {
	EmailAddresses []string `json:"email_addresses,omitempty"`
}

// ChannelContent represents the content block for an AlertChannel, which varies by channel type.
type ChannelContent struct {
	Email *EmailChannelContent `json:"email"`
	// Other channel types like 'webhook', 'slack' could be added her
}

// ChannelDetails represents the details block for an AlertChannel
type ChannelDetails struct {
	Email *EmailChannelDetails `json:"email"`
	// Other channel types could be added here
}

// EmailChannelDetails represents email-specific details for a channel
type EmailChannelDetails struct {
	Usernames     []string `json:"usernames"`
	RecipientType string   `json:"recipient_type"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (a *AlertChannel) UnmarshalJSON(b []byte) error {
	type Mask AlertChannel

	p := struct {
		*Mask

		Updated *parseabletime.ParseableTime `json:"updated"`
		Created *parseabletime.ParseableTime `json:"created"`
	}{
		Mask: (*Mask)(a),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	a.Updated = (*time.Time)(p.Updated)
	a.Created = (*time.Time)(p.Created)

	return nil
}

// ListAlertChannels gets a paginated list of Alert Channels.
func (c *Client) ListAlertChannels(ctx context.Context, opts *ListOptions) ([]AlertChannel, error) {
	endpoint := formatAPIPath("monitor/alert-channels")
	return getPaginatedResults[AlertChannel](ctx, c, endpoint, opts)
}
