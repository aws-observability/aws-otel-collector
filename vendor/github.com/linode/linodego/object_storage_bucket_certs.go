package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type ObjectStorageBucketCert struct {
	SSL bool `json:"ssl"`
}

type ObjectStorageBucketCertUploadOptions struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

// UploadObjectStorageBucketCert uploads a TLS/SSL Cert to be used with an Object Storage Bucket.
func (c *Client) UploadObjectStorageBucketCert(ctx context.Context, clusterOrRegionID, bucket string, opts ObjectStorageBucketCertUploadOptions) (*ObjectStorageBucketCert, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	clusterOrRegionID = url.PathEscape(clusterOrRegionID)
	bucket = url.PathEscape(bucket)
	e := fmt.Sprintf("object-storage/buckets/%s/%s/ssl", clusterOrRegionID, bucket)
	req := c.R(ctx).SetResult(&ObjectStorageBucketCert{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Post(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*ObjectStorageBucketCert), nil
}

// GetObjectStorageBucketCert gets an ObjectStorageBucketCert
func (c *Client) GetObjectStorageBucketCert(ctx context.Context, clusterOrRegionID, bucket string) (*ObjectStorageBucketCert, error) {
	clusterOrRegionID = url.PathEscape(clusterOrRegionID)
	bucket = url.PathEscape(bucket)
	e := fmt.Sprintf("object-storage/buckets/%s/%s/ssl", clusterOrRegionID, bucket)
	req := c.R(ctx).SetResult(&ObjectStorageBucketCert{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*ObjectStorageBucketCert), nil
}

// DeleteObjectStorageBucketCert deletes an ObjectStorageBucketCert
func (c *Client) DeleteObjectStorageBucketCert(ctx context.Context, clusterOrRegionID, bucket string) error {
	clusterOrRegionID = url.PathEscape(clusterOrRegionID)
	bucket = url.PathEscape(bucket)
	e := fmt.Sprintf("object-storage/buckets/%s/%s/ssl", clusterOrRegionID, bucket)
	_, err := coupleAPIErrors(c.R(ctx).Delete(e))
	return err
}
