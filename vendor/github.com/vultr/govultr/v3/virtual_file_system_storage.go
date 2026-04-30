package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const virtualFileSystemStoragePath = "/v2/vfs"

// VirtualFileSystemStorageService is the interface to interact with virtual
// file system endpoint on the Vultr API
// Link : https://www.vultr.com/api/#tag/vfs
type VirtualFileSystemStorageService interface {
	Create(ctx context.Context, vfsReq *VirtualFileSystemStorageReq) (*VirtualFileSystemStorage, *http.Response, error)
	Get(ctx context.Context, vfsID string) (*VirtualFileSystemStorage, *http.Response, error)
	Update(ctx context.Context, vfsID string, vfsUpdateReq *VirtualFileSystemStorageUpdateReq) (*VirtualFileSystemStorage, *http.Response, error) //nolint:lll
	Delete(ctx context.Context, vfsID string) error
	List(ctx context.Context, options *ListOptions) ([]VirtualFileSystemStorage, *Meta, *http.Response, error)

	AttachmentList(ctx context.Context, vfsID string) ([]VirtualFileSystemStorageAttachment, *http.Response, error)
	AttachmentGet(ctx context.Context, vfsID, targetID string) (*VirtualFileSystemStorageAttachment, *http.Response, error)
	Attach(ctx context.Context, vfsID, targetID string) (*VirtualFileSystemStorageAttachment, *http.Response, error)
	Detach(ctx context.Context, vfsID, targetID string) error
}

// VirtualFileSystemStorageServiceHandler handles interaction with the virtual
// file system methods for the Vultr API.
type VirtualFileSystemStorageServiceHandler struct {
	client *Client
}

// VirtualFileSystemStorage represents a virtual file system storage.
type VirtualFileSystemStorage struct {
	ID          string                          `json:"id"`
	Region      string                          `json:"region"`
	DateCreated string                          `json:"date_created"`
	Status      string                          `json:"status"`
	Label       string                          `json:"label"`
	Tags        []string                        `json:"tags"`
	DiskType    string                          `json:"disk_type"`
	StorageSize VirtualFileSystemStorageSize    `json:"storage_size"`
	StorageUsed VirtualFileSystemStorageSize    `json:"storage_used"`
	Billing     VirtualFileSystemStorageBilling `json:"billing"`
}

// VirtualFileSystemStorageSize represents the on disk size of a virtual file
// system storage.
type VirtualFileSystemStorageSize struct {
	SizeBytes int `json:"bytes,omitempty"`
	SizeGB    int `json:"gb"`
}

// VirtualFileSystemStorageBilling represents the billing data of a virtual
// file system storage.
type VirtualFileSystemStorageBilling struct {
	Charges float32 `json:"charges"`
	Monthly float32 `json:"monthly"`
}

type virtualFileSystemStoragesBase struct {
	VFSs []VirtualFileSystemStorage `json:"vfs"`
	Meta *Meta                      `json:"meta"`
}

// VirtualFileSystemStorageReq struct represents the request used when creating
// a virtual file system storage.
type VirtualFileSystemStorageReq struct {
	Region      string                       `json:"region"`
	Label       string                       `json:"label"`
	StorageSize VirtualFileSystemStorageSize `json:"storage_size"`
	DiskType    string                       `json:"disk_type,omitempty"`
	Tags        []string                     `json:"tags,omitempty"`
}

// VirtualFileSystemStorageUpdateReq struct represents the request used when
// updating virtual file system storage.
type VirtualFileSystemStorageUpdateReq struct {
	Label       string                       `json:"label,omitempty"`
	StorageSize VirtualFileSystemStorageSize `json:"storage_size"`
}

// VirtualFileSystemStorageAttachment represents an attachment for a virtual
// file system.
type VirtualFileSystemStorageAttachment struct {
	ID       string `json:"vfs_id"`
	State    string `json:"state"`
	TargetID string `json:"target_id"`
	MountTag int    `json:"mount_tag"`
}

type virtualFileSystemStorageAttachmentsBase struct {
	Attachments []VirtualFileSystemStorageAttachment `json:"attachments"`
}

// Create sends a create request for a virtual file system storage.
func (f *VirtualFileSystemStorageServiceHandler) Create(ctx context.Context, vfsReq *VirtualFileSystemStorageReq) (*VirtualFileSystemStorage, *http.Response, error) { //nolint:lll
	req, err := f.client.NewRequest(ctx, http.MethodPost, virtualFileSystemStoragePath, vfsReq)
	if err != nil {
		return nil, nil, err
	}

	vfs := new(VirtualFileSystemStorage)
	resp, err := f.client.DoWithContext(ctx, req, vfs)
	if err != nil {
		return nil, resp, err
	}

	return vfs, resp, nil
}

// Get retrieves a single virtual file system storage.
func (f *VirtualFileSystemStorageServiceHandler) Get(ctx context.Context, vfsID string) (*VirtualFileSystemStorage, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", virtualFileSystemStoragePath, vfsID)

	req, err := f.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	vfs := new(VirtualFileSystemStorage)
	resp, err := f.client.DoWithContext(ctx, req, vfs)
	if err != nil {
		return nil, resp, err
	}

	return vfs, resp, nil
}

// Update sends a update request for a virtual file system storage.
func (f *VirtualFileSystemStorageServiceHandler) Update(ctx context.Context, vfsID string, vfsUpdateReq *VirtualFileSystemStorageUpdateReq) (*VirtualFileSystemStorage, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s", virtualFileSystemStoragePath, vfsID)

	req, err := f.client.NewRequest(ctx, http.MethodPut, uri, vfsUpdateReq)
	if err != nil {
		return nil, nil, err
	}

	vfs := new(VirtualFileSystemStorage)
	resp, err := f.client.DoWithContext(ctx, req, vfs)
	if err != nil {
		return nil, resp, err
	}
	return vfs, resp, err
}

// Delete sends a delete request for a virtual file system storage.
func (f *VirtualFileSystemStorageServiceHandler) Delete(ctx context.Context, vfsID string) error {
	uri := fmt.Sprintf("%s/%s", virtualFileSystemStoragePath, vfsID)

	req, err := f.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	_, err = f.client.DoWithContext(ctx, req, nil)
	return err
}

// List retrieves a list of all virtual file system storages.
func (f *VirtualFileSystemStorageServiceHandler) List(ctx context.Context, options *ListOptions) ([]VirtualFileSystemStorage, *Meta, *http.Response, error) { //nolint:dupl,lll
	req, err := f.client.NewRequest(ctx, http.MethodGet, virtualFileSystemStoragePath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	vfsStorages := new(virtualFileSystemStoragesBase)
	resp, err := f.client.DoWithContext(ctx, req, vfsStorages)
	if err != nil {
		return nil, nil, resp, err
	}

	return vfsStorages.VFSs, vfsStorages.Meta, resp, nil
}

// Attach attaches a virtual file system storage to another instance.
func (f *VirtualFileSystemStorageServiceHandler) Attach(ctx context.Context, vfsID, targetID string) (*VirtualFileSystemStorageAttachment, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/attachments/%s", virtualFileSystemStoragePath, vfsID, targetID)

	req, err := f.client.NewRequest(ctx, http.MethodPut, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	att := new(VirtualFileSystemStorageAttachment)
	resp, err := f.client.DoWithContext(ctx, req, att)
	if err != nil {
		return nil, resp, err
	}

	return att, resp, err
}

// AttachmentList retrieves a list the active attachments on a virtual file
// system storage.
func (f *VirtualFileSystemStorageServiceHandler) AttachmentList(ctx context.Context, vfsID string) ([]VirtualFileSystemStorageAttachment, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/attachments", virtualFileSystemStoragePath, vfsID)

	req, err := f.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	atts := new(virtualFileSystemStorageAttachmentsBase)
	resp, err := f.client.DoWithContext(ctx, req, atts)
	if err != nil {
		return nil, resp, err
	}

	return atts.Attachments, resp, err
}

// AttachmentGet retrieves the attachment of a virtual file system storage and
// its attached instance.
func (f *VirtualFileSystemStorageServiceHandler) AttachmentGet(ctx context.Context, vfsID, targetID string) (*VirtualFileSystemStorageAttachment, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/attachments/%s", virtualFileSystemStoragePath, vfsID, targetID)

	req, err := f.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	att := new(VirtualFileSystemStorageAttachment)
	resp, err := f.client.DoWithContext(ctx, req, att)
	if err != nil {
		return nil, resp, err
	}

	return att, resp, err
}

// Detach sends a delete request for an attachment of a virtual file system
// storage, detaching it from its instance.
func (f *VirtualFileSystemStorageServiceHandler) Detach(ctx context.Context, vfsID, targetID string) error {
	uri := fmt.Sprintf("%s/%s/attachments/%s", virtualFileSystemStoragePath, vfsID, targetID)

	req, err := f.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	if _, err := f.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
