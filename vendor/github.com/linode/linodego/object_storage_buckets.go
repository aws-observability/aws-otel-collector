package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/linode/linodego/internal/parseabletime"
)

// ObjectStorageBucket represents a ObjectStorage object
type ObjectStorageBucket struct {
	Label string `json:"label"`

	// Deprecated: The 'Cluster' field has been deprecated in favor of the 'Region' field.
	// For example, a Cluster value of `us-mia-1` will translate to a Region value of `us-mia`.
	//
	// This is necessary because there are now multiple Object Storage clusters to a region.
	//
	// NOTE: The 'Cluster' field will always return a value similar to `<REGION>-1` (e.g., `us-mia-1`)
	// for backward compatibility purposes.
	Cluster string `json:"cluster"`
	Region  string `json:"region"`

	Created  *time.Time `json:"-"`
	Hostname string     `json:"hostname"`
	Objects  int        `json:"objects"`
	Size     int        `json:"size"`
}

// ObjectStorageBucketAccess holds Object Storage access info
type ObjectStorageBucketAccess struct {
	ACL         ObjectStorageACL `json:"acl"`
	CorsEnabled bool             `json:"cors_enabled"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *ObjectStorageBucket) UnmarshalJSON(b []byte) error {
	type Mask ObjectStorageBucket

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)

	return nil
}

// ObjectStorageBucketCreateOptions fields are those accepted by CreateObjectStorageBucket
type ObjectStorageBucketCreateOptions struct {
	// Deprecated: The 'Cluster' field has been deprecated.
	//
	// Going forward, the 'Region' field will be the supported way to designate where an
	// Object Storage Bucket should be created. For example, a 'Cluster' value of `us-mia-1`
	// will translate to a Region value of `us-mia`.
	Cluster string `json:"cluster,omitempty"`
	Region  string `json:"region,omitempty"`

	Label string `json:"label"`

	ACL         ObjectStorageACL `json:"acl,omitempty"`
	CorsEnabled *bool            `json:"cors_enabled,omitempty"`
}

// ObjectStorageBucketUpdateAccessOptions fields are those accepted by UpdateObjectStorageBucketAccess
type ObjectStorageBucketUpdateAccessOptions struct {
	ACL         ObjectStorageACL `json:"acl,omitempty"`
	CorsEnabled *bool            `json:"cors_enabled,omitempty"`
}

// ObjectStorageACL options start with ACL and include all known ACL types
type ObjectStorageACL string

// ObjectStorageACL options represent the access control level of a bucket.
const (
	ACLPrivate           ObjectStorageACL = "private"
	ACLPublicRead        ObjectStorageACL = "public-read"
	ACLAuthenticatedRead ObjectStorageACL = "authenticated-read"
	ACLPublicReadWrite   ObjectStorageACL = "public-read-write"
)

// ObjectStorageBucketsPagedResponse represents a paginated ObjectStorageBucket API response
type ObjectStorageBucketsPagedResponse struct {
	*PageOptions
	Data []ObjectStorageBucket `json:"data"`
}

// endpoint gets the endpoint URL for ObjectStorageBucket
func (ObjectStorageBucketsPagedResponse) endpoint(args ...any) string {
	endpoint := "object-storage/buckets"
	if len(args) > 0 {
		endpoint = fmt.Sprintf(endpoint+"/%s", url.PathEscape(args[0].(string)))
	}
	return endpoint
}

func (resp *ObjectStorageBucketsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(ObjectStorageBucketsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*ObjectStorageBucketsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListObjectStorageBuckets lists ObjectStorageBuckets
func (c *Client) ListObjectStorageBuckets(ctx context.Context, opts *ListOptions) ([]ObjectStorageBucket, error) {
	response := ObjectStorageBucketsPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// ListObjectStorageBucketsInCluster lists all ObjectStorageBuckets of a cluster
func (c *Client) ListObjectStorageBucketsInCluster(ctx context.Context, opts *ListOptions, clusterOrRegionID string) ([]ObjectStorageBucket, error) {
	response := ObjectStorageBucketsPagedResponse{}
	err := c.listHelper(ctx, &response, opts, clusterOrRegionID)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetObjectStorageBucket gets the ObjectStorageBucket with the provided label
func (c *Client) GetObjectStorageBucket(ctx context.Context, clusterOrRegionID, label string) (*ObjectStorageBucket, error) {
	label = url.PathEscape(label)
	clusterOrRegionID = url.PathEscape(clusterOrRegionID)
	e := fmt.Sprintf("object-storage/buckets/%s/%s", clusterOrRegionID, label)
	req := c.R(ctx).SetResult(&ObjectStorageBucket{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*ObjectStorageBucket), nil
}

// CreateObjectStorageBucket creates an ObjectStorageBucket
func (c *Client) CreateObjectStorageBucket(ctx context.Context, opts ObjectStorageBucketCreateOptions) (*ObjectStorageBucket, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	e := "object-storage/buckets"
	req := c.R(ctx).SetResult(&ObjectStorageBucket{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Post(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*ObjectStorageBucket), nil
}

// GetObjectStorageBucketAccess gets the current access config for a bucket
func (c *Client) GetObjectStorageBucketAccess(ctx context.Context, clusterOrRegionID, label string) (*ObjectStorageBucketAccess, error) {
	label = url.PathEscape(label)
	clusterOrRegionID = url.PathEscape(clusterOrRegionID)
	e := fmt.Sprintf("object-storage/buckets/%s/%s/access", clusterOrRegionID, label)
	req := c.R(ctx).SetResult(&ObjectStorageBucketAccess{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*ObjectStorageBucketAccess), nil
}

// UpdateObjectStorageBucketAccess updates the access configuration for an ObjectStorageBucket
func (c *Client) UpdateObjectStorageBucketAccess(ctx context.Context, clusterOrRegionID, label string, opts ObjectStorageBucketUpdateAccessOptions) error {
	body, err := json.Marshal(opts)
	if err != nil {
		return err
	}

	label = url.PathEscape(label)
	clusterOrRegionID = url.PathEscape(clusterOrRegionID)
	e := fmt.Sprintf("object-storage/buckets/%s/%s/access", clusterOrRegionID, label)
	_, err = coupleAPIErrors(c.R(ctx).SetBody(string(body)).Post(e))
	if err != nil {
		return err
	}

	return nil
}

// DeleteObjectStorageBucket deletes the ObjectStorageBucket with the specified label
func (c *Client) DeleteObjectStorageBucket(ctx context.Context, clusterOrRegionID, label string) error {
	label = url.PathEscape(label)
	e := fmt.Sprintf("object-storage/buckets/%s/%s", clusterOrRegionID, label)
	_, err := coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}
