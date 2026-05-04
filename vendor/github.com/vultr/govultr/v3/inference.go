package govultr

import (
	"context"
	"fmt"
	"net/http"
)

const inferencePath = "/v2/inference"

// InferenceService is the interface to interact with the Inference endpoints on the Vultr API
// Link: https://www.vultr.com/api/#tag/serverless-inference
type InferenceService interface {
	List(ctx context.Context) ([]Inference, *http.Response, error)
	Create(ctx context.Context, inferenceReq *InferenceCreateUpdateReq) (*Inference, *http.Response, error)
	Get(ctx context.Context, inferenceID string) (*Inference, *http.Response, error)
	Update(ctx context.Context, inferenceID string, inferenceReq *InferenceCreateUpdateReq) (*Inference, *http.Response, error)
	Delete(ctx context.Context, inferenceID string) error
	GetUsage(ctx context.Context, inferenceID string) (*InferenceUsage, *http.Response, error)
}

// InferenceServiceHandler handles interaction with the server methods for the Vultr API
type InferenceServiceHandler struct {
	client *Client
}

// Inference represents a Serverless Inference subscription
type Inference struct {
	ID          string `json:"id"`
	DateCreated string `json:"date_created"`
	Label       string `json:"label"`
	APIKey      string `json:"api_key"`
}

// inferenceSubsBase holds the entire List API response
type inferenceSubsBase struct {
	Subscriptions []Inference `json:"subscriptions"`
}

// inferenceSubBase holds the entire Get API response
type inferenceSubBase struct {
	Subscription *Inference `json:"subscription"`
}

// InferenceCreateUpdateReq struct used to create and update a Serverless Inference subscription
type InferenceCreateUpdateReq struct {
	Label string `json:"label,omitempty"`
}

// InferenceChatUsage represents chat/embeddings usage details for a Serverless Inference subscription
type InferenceChatUsage struct {
	CurrentTokens    int `json:"current_tokens"`
	MonthlyAllotment int `json:"monthly_allotment"`
	Overage          int `json:"overage"`
}

// InferenceAudioUsage represents audio generation details for a Serverless Inference subscription
type InferenceAudioUsage struct {
	TTSCharacters   int `json:"tts_characters"`
	TTSSMCharacters int `json:"tts_sm_characters"`
}

// InferenceUsage represents chat/embeddings and audio usage for a Serverless Inference subscription
type InferenceUsage struct {
	Chat  InferenceChatUsage  `json:"chat"`
	Audio InferenceAudioUsage `json:"audio"`
}

// inferenceUsageBase represents a migration status object API response for a Serverless Inference subscription
type inferenceUsageBase struct {
	Usage *InferenceUsage `json:"usage"`
}

// List retrieves all Serverless Inference subscriptions on your account
func (d *InferenceServiceHandler) List(ctx context.Context) ([]Inference, *http.Response, error) {
	req, err := d.client.NewRequest(ctx, http.MethodGet, inferencePath, nil)
	if err != nil {
		return nil, nil, err
	}

	inferenceSubs := new(inferenceSubsBase)
	resp, err := d.client.DoWithContext(ctx, req, inferenceSubs)
	if err != nil {
		return nil, nil, err
	}

	return inferenceSubs.Subscriptions, resp, nil
}

// Create will create the Serverless Inference subscription with the given parameters
func (d *InferenceServiceHandler) Create(ctx context.Context, inferenceReq *InferenceCreateUpdateReq) (*Inference, *http.Response, error) { //nolint:lll
	req, err := d.client.NewRequest(ctx, http.MethodPost, inferencePath, inferenceReq)
	if err != nil {
		return nil, nil, err
	}

	inferenceSub := new(inferenceSubBase)
	resp, err := d.client.DoWithContext(ctx, req, inferenceSub)
	if err != nil {
		return nil, nil, err
	}

	return inferenceSub.Subscription, resp, nil
}

// Get will get the Serverless Inference subscription with the given inferenceID
func (d *InferenceServiceHandler) Get(ctx context.Context, inferenceID string) (*Inference, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", inferencePath, inferenceID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	inferenceSub := new(inferenceSubBase)
	resp, err := d.client.DoWithContext(ctx, req, inferenceSub)
	if err != nil {
		return nil, nil, err
	}

	return inferenceSub.Subscription, resp, nil
}

// Update will update the Serverless Inference subscription with the given parameters
func (d *InferenceServiceHandler) Update(ctx context.Context, inferenceID string, inferenceReq *InferenceCreateUpdateReq) (*Inference, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s", inferencePath, inferenceID)

	req, err := d.client.NewRequest(ctx, http.MethodPatch, uri, inferenceReq)
	if err != nil {
		return nil, nil, err
	}

	inferenceSub := new(inferenceSubBase)
	resp, err := d.client.DoWithContext(ctx, req, inferenceSub)
	if err != nil {
		return nil, nil, err
	}

	return inferenceSub.Subscription, resp, nil
}

// Delete a Serverless Inference subscription. All data will be permanently lost
func (d *InferenceServiceHandler) Delete(ctx context.Context, inferenceID string) error {
	uri := fmt.Sprintf("%s/%s", inferencePath, inferenceID)

	req, err := d.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// GetUsage retrieves chat/embeddings and audio generation usage information for your Serverless Inference subscription
func (d *InferenceServiceHandler) GetUsage(ctx context.Context, inferenceID string) (*InferenceUsage, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/usage", inferencePath, inferenceID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	inferenceUsage := new(inferenceUsageBase)
	resp, err := d.client.DoWithContext(ctx, req, inferenceUsage)
	if err != nil {
		return nil, nil, err
	}

	return inferenceUsage.Usage, resp, nil
}
