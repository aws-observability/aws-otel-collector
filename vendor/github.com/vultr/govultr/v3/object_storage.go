package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// ObjectStorageService is the interface to interact with the object storage endpoints on the Vultr API.
// Link : https://www.vultr.com/api/#tag/s3
type ObjectStorageService interface {
	Create(ctx context.Context, objReq *ObjectStorageReq) (*ObjectStorage, *http.Response, error)
	Get(ctx context.Context, id string) (*ObjectStorage, *http.Response, error)
	Update(ctx context.Context, id string, objReq *ObjectStorageReq) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, options *ListOptions) ([]ObjectStorage, *Meta, *http.Response, error)

	ListCluster(ctx context.Context, options *ListOptions) ([]ObjectStorageCluster, *Meta, *http.Response, error)
	RegenerateKeys(ctx context.Context, id string) (*S3Keys, *http.Response, error)

	ListTiers(ctx context.Context) ([]ObjectStorageTier, *http.Response, error)
	ListClusterTiers(ctx context.Context, clusterID int) ([]ObjectStorageTier, *http.Response, error)
}

// ObjectStorageServiceHandler handles interaction between the object storage service and the Vultr API.
type ObjectStorageServiceHandler struct {
	client *Client
}

// ObjectStorage represents a Vultr Object Storage subscription.
type ObjectStorage struct {
	ID                   string `json:"id"`
	DateCreated          string `json:"date_created"`
	ObjectStoreClusterID int    `json:"cluster_id"`
	Region               string `json:"region"`
	Location             string `json:"location"`
	Label                string `json:"label"`
	Status               string `json:"status"`
	S3Keys
}

// ObjectStorageReq represents the parameters for creating and updating object
// storages
type ObjectStorageReq struct {
	ClusterID int    `json:"cluster_id,omitempty"`
	TierID    int    `json:"tier_id,omitempty"`
	Label     string `json:"label"`
}

// S3Keys define your api access to your cluster
type S3Keys struct {
	S3Hostname  string `json:"s3_hostname"`
	S3AccessKey string `json:"s3_access_key"`
	S3SecretKey string `json:"s3_secret_key"`
}

// ObjectStorageCluster represents a Vultr Object Storage cluster.
type ObjectStorageCluster struct {
	ID       int    `json:"id"`
	Region   string `json:"region"`
	Name     string `json:"name,omitempty"`
	Hostname string `json:"hostname"`
	Deploy   string `json:"deploy"`
}

// ObjectStorageTier represents the object storage tier data
type ObjectStorageTier struct {
	ID                int                    `json:"id"`
	Name              string                 `json:"sales_name"`
	Description       string                 `json:"sales_desc"`
	Price             float32                `json:"price"`
	PriceBandwidthGB  float32                `json:"bw_gb_price"`
	PriceDiskGB       float32                `json:"disk_gb_price"`
	RateLimitBytesSec int                    `json:"ratelimit_ops_bytes"`
	RateLimitOpsSec   int                    `json:"ratelimit_ops_secs"`
	Default           string                 `json:"is_default"`
	Locations         []ObjectStorageCluster `json:"locations,omitempty"`
	Slug              string                 `json:"slug"`
}

type objectStoragesBase struct {
	ObjectStorages []ObjectStorage `json:"object_storages"`
	Meta           *Meta           `json:"meta"`
}

type objectStorageBase struct {
	ObjectStorage *ObjectStorage `json:"object_storage"`
}

type objectStorageClustersBase struct {
	Clusters []ObjectStorageCluster `json:"clusters"`
	Meta     *Meta                  `json:"meta"`
}

type objectStorageTiersBase struct {
	Tiers []ObjectStorageTier `json:"tiers"`
}

type s3KeysBase struct {
	S3Credentials *S3Keys `json:"s3_credentials"`
}

// Create an object storage subscription
func (o *ObjectStorageServiceHandler) Create(ctx context.Context, objReq *ObjectStorageReq) (*ObjectStorage, *http.Response, error) {
	uri := "/v2/object-storage"

	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, objReq)
	if err != nil {
		return nil, nil, err
	}

	objectStorage := new(objectStorageBase)
	resp, err := o.client.DoWithContext(ctx, req, objectStorage)
	if err != nil {
		return nil, resp, err
	}

	return objectStorage.ObjectStorage, resp, nil
}

// Get returns a specified object storage by the provided ID
func (o *ObjectStorageServiceHandler) Get(ctx context.Context, id string) (*ObjectStorage, *http.Response, error) {
	uri := fmt.Sprintf("/v2/object-storage/%s", id)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	objectStorage := new(objectStorageBase)
	resp, err := o.client.DoWithContext(ctx, req, objectStorage)
	if err != nil {
		return nil, resp, err
	}

	return objectStorage.ObjectStorage, resp, nil
}

// Update a Object Storage Subscription.
func (o *ObjectStorageServiceHandler) Update(ctx context.Context, id string, objReq *ObjectStorageReq) error {
	uri := fmt.Sprintf("/v2/object-storage/%s", id)

	req, err := o.client.NewRequest(ctx, http.MethodPut, uri, objReq)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	return err
}

// Delete a object storage subscription.
func (o *ObjectStorageServiceHandler) Delete(ctx context.Context, id string) error {
	uri := fmt.Sprintf("/v2/object-storage/%s", id)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	return err
}

// List all object storage subscriptions on the current account. This includes both pending and active subscriptions.
func (o *ObjectStorageServiceHandler) List(ctx context.Context, options *ListOptions) ([]ObjectStorage, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := "/v2/object-storage"

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	objectStorage := new(objectStoragesBase)
	resp, err := o.client.DoWithContext(ctx, req, objectStorage)
	if err != nil {
		return nil, nil, resp, err
	}

	return objectStorage.ObjectStorages, objectStorage.Meta, resp, nil
}

// ListCluster returns back your object storage clusters.
// Clusters may be removed over time. The "deploy" field can be used to determine whether or not new deployments are allowed in the cluster.
func (o *ObjectStorageServiceHandler) ListCluster(ctx context.Context, options *ListOptions) ([]ObjectStorageCluster, *Meta, *http.Response, error) { //nolint:lll
	uri := "/v2/object-storage/clusters"
	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	clusters := new(objectStorageClustersBase)
	resp, err := o.client.DoWithContext(ctx, req, clusters)
	if err != nil {
		return nil, nil, resp, err
	}

	return clusters.Clusters, clusters.Meta, resp, nil
}

// RegenerateKeys of the S3 API Keys for an object storage subscription
func (o *ObjectStorageServiceHandler) RegenerateKeys(ctx context.Context, id string) (*S3Keys, *http.Response, error) {
	uri := fmt.Sprintf("/v2/object-storage/%s/regenerate-keys", id)

	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	s3Keys := new(s3KeysBase)
	resp, err := o.client.DoWithContext(ctx, req, s3Keys)
	if err != nil {
		return nil, resp, err
	}

	return s3Keys.S3Credentials, resp, nil
}

// ListTiers retrieves all tiers for object storage deployments
func (o *ObjectStorageServiceHandler) ListTiers(ctx context.Context) ([]ObjectStorageTier, *http.Response, error) {
	uri := "/v2/object-storage/tiers"
	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	tiers := new(objectStorageTiersBase)
	resp, err := o.client.DoWithContext(ctx, req, tiers)
	if err != nil {
		return nil, resp, err
	}

	return tiers.Tiers, resp, nil
}

// ListClusterTiers retrieves all tiers for object storage deployments on a specific cluster
func (o *ObjectStorageServiceHandler) ListClusterTiers(ctx context.Context, clusterID int) ([]ObjectStorageTier, *http.Response, error) {
	uri := fmt.Sprintf("/v2/object-storage/clusters/%d/tiers", clusterID)
	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	tiers := new(objectStorageTiersBase)
	resp, err := o.client.DoWithContext(ctx, req, tiers)
	if err != nil {
		return nil, resp, err
	}

	return tiers.Tiers, resp, nil
}
