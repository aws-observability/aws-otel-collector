package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// SnapshotService is the interface to interact with Snapshot endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/snapshot
type SnapshotService interface {
	Create(ctx context.Context, snapshotReq *SnapshotReq) (*Snapshot, *http.Response, error)
	CreateFromURL(ctx context.Context, snapshotURLReq *SnapshotURLReq) (*Snapshot, *http.Response, error)
	Get(ctx context.Context, snapshotID string) (*Snapshot, *http.Response, error)
	Delete(ctx context.Context, snapshotID string) error
	List(ctx context.Context, options *ListOptions) ([]Snapshot, *Meta, *http.Response, error)
}

// SnapshotServiceHandler handles interaction with the snapshot methods for the Vultr API
type SnapshotServiceHandler struct {
	client *Client
}

// Snapshot represents a Vultr snapshot
type Snapshot struct {
	ID             string `json:"id"`
	DateCreated    string `json:"date_created"`
	Description    string `json:"description"`
	Size           int    `json:"size"`
	CompressedSize int    `json:"compressed_size"`
	Status         string `json:"status"`
	OsID           int    `json:"os_id"`
	AppID          int    `json:"app_id"`
}

// SnapshotReq struct is used to create snapshots.
type SnapshotReq struct {
	InstanceID  string `json:"instance_id,omitempty"`
	Description string `json:"description,omitempty"`
}

// SnapshotURLReq struct is used to create snapshots from a URL.
type SnapshotURLReq struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
	UEFI        *bool  `json:"uefi"`
}

type snapshotsBase struct {
	Snapshots []Snapshot `json:"snapshots"`
	Meta      *Meta      `json:"meta"`
}

type snapshotBase struct {
	Snapshot *Snapshot `json:"snapshot"`
}

// Create makes a snapshot of a provided server
func (s *SnapshotServiceHandler) Create(ctx context.Context, snapshotReq *SnapshotReq) (*Snapshot, *http.Response, error) {
	uri := "/v2/snapshots"

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, snapshotReq)
	if err != nil {
		return nil, nil, err
	}

	snapshot := new(snapshotBase)
	resp, err := s.client.DoWithContext(ctx, req, snapshot)
	if err != nil {
		return nil, resp, err
	}

	return snapshot.Snapshot, resp, nil
}

// CreateFromURL will create a snapshot based on an image iso from a URL you provide
func (s *SnapshotServiceHandler) CreateFromURL(ctx context.Context, snapshotURLReq *SnapshotURLReq) (*Snapshot, *http.Response, error) {
	uri := "/v2/snapshots/create-from-url"

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, snapshotURLReq)
	if err != nil {
		return nil, nil, err
	}

	snapshot := new(snapshotBase)
	resp, err := s.client.DoWithContext(ctx, req, snapshot)
	if err != nil {
		return nil, resp, err
	}

	return snapshot.Snapshot, resp, nil
}

// Get a specific snapshot
func (s *SnapshotServiceHandler) Get(ctx context.Context, snapshotID string) (*Snapshot, *http.Response, error) {
	uri := fmt.Sprintf("/v2/snapshots/%s", snapshotID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	snapshot := new(snapshotBase)
	resp, err := s.client.DoWithContext(ctx, req, snapshot)
	if err != nil {
		return nil, resp, err
	}

	return snapshot.Snapshot, resp, nil
}

// Delete a snapshot.
func (s *SnapshotServiceHandler) Delete(ctx context.Context, snapshotID string) error {
	uri := fmt.Sprintf("/v2/snapshots/%s", snapshotID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = s.client.DoWithContext(ctx, req, nil)
	return err
}

// List all available snapshots.
func (s *SnapshotServiceHandler) List(ctx context.Context, options *ListOptions) ([]Snapshot, *Meta, *http.Response, error) { //nolint:dupl
	uri := "/v2/snapshots"

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}
	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	snapshots := new(snapshotsBase)
	resp, err := s.client.DoWithContext(ctx, req, snapshots)
	if err != nil {
		return nil, nil, resp, err
	}

	return snapshots.Snapshots, snapshots.Meta, resp, nil
}
