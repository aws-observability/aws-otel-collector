package linodego

import (
	"context"
	"encoding/json"
	"time"

	"github.com/linode/linodego/internal/parseabletime"
)

// VolumeStatus indicates the status of the Volume
type VolumeStatus string

const (
	// VolumeCreating indicates the Volume is being created and is not yet available for use
	VolumeCreating VolumeStatus = "creating"

	// VolumeActive indicates the Volume is online and available for use
	VolumeActive VolumeStatus = "active"

	// VolumeResizing indicates the Volume is in the process of upgrading its current capacity
	VolumeResizing VolumeStatus = "resizing"

	// VolumeContactSupport indicates there is a problem with the Volume. A support ticket must be opened to resolve the issue
	VolumeContactSupport VolumeStatus = "contact_support"
)

// Volume represents a linode volume object
type Volume struct {
	ID             int          `json:"id"`
	Label          string       `json:"label"`
	Status         VolumeStatus `json:"status"`
	Region         string       `json:"region"`
	Size           int          `json:"size"`
	LinodeID       *int         `json:"linode_id"`
	FilesystemPath string       `json:"filesystem_path"`
	Tags           []string     `json:"tags"`
	HardwareType   string       `json:"hardware_type"`
	LinodeLabel    string       `json:"linode_label"`
	Created        *time.Time   `json:"-"`
	Updated        *time.Time   `json:"-"`

	// Note: Block Storage Disk Encryption is not currently available to all users.
	Encryption string `json:"encryption"`
}

// VolumeCreateOptions fields are those accepted by CreateVolume
type VolumeCreateOptions struct {
	Label    string `json:"label,omitempty"`
	Region   string `json:"region,omitempty"`
	LinodeID int    `json:"linode_id,omitempty"`
	ConfigID int    `json:"config_id,omitempty"`
	// The Volume's size, in GiB. Minimum size is 10GiB, maximum size is 10240GiB. A "0" value will result in the default size.
	Size int `json:"size,omitempty"`
	// An array of tags applied to this object. Tags are for organizational purposes only.
	Tags               []string `json:"tags"`
	PersistAcrossBoots *bool    `json:"persist_across_boots,omitempty"`
	Encryption         string   `json:"encryption,omitempty"`
}

// VolumeUpdateOptions fields are those accepted by UpdateVolume
type VolumeUpdateOptions struct {
	Label string    `json:"label,omitempty"`
	Tags  *[]string `json:"tags,omitempty"`
}

// VolumeAttachOptions fields are those accepted by AttachVolume
type VolumeAttachOptions struct {
	LinodeID           int   `json:"linode_id"`
	ConfigID           int   `json:"config_id,omitempty"`
	PersistAcrossBoots *bool `json:"persist_across_boots,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (v *Volume) UnmarshalJSON(b []byte) error {
	type Mask Volume

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(v),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	v.Created = (*time.Time)(p.Created)
	v.Updated = (*time.Time)(p.Updated)

	return nil
}

// GetUpdateOptions converts a Volume to VolumeUpdateOptions for use in UpdateVolume
func (v Volume) GetUpdateOptions() (updateOpts VolumeUpdateOptions) {
	updateOpts.Label = v.Label
	updateOpts.Tags = &v.Tags
	return
}

// GetCreateOptions converts a Volume to VolumeCreateOptions for use in CreateVolume
func (v Volume) GetCreateOptions() (createOpts VolumeCreateOptions) {
	createOpts.Label = v.Label
	createOpts.Tags = v.Tags
	createOpts.Region = v.Region
	createOpts.Size = v.Size
	if v.LinodeID != nil && *v.LinodeID > 0 {
		createOpts.LinodeID = *v.LinodeID
	}
	return
}

// ListVolumes lists Volumes
func (c *Client) ListVolumes(ctx context.Context, opts *ListOptions) ([]Volume, error) {
	return getPaginatedResults[Volume](ctx, c, "volumes", opts)
}

// GetVolume gets the template with the provided ID
func (c *Client) GetVolume(ctx context.Context, volumeID int) (*Volume, error) {
	e := formatAPIPath("volumes/%d", volumeID)
	return doGETRequest[Volume](ctx, c, e)
}

// AttachVolume attaches a volume to a Linode instance
func (c *Client) AttachVolume(ctx context.Context, volumeID int, opts *VolumeAttachOptions) (*Volume, error) {
	e := formatAPIPath("volumes/%d/attach", volumeID)
	return doPOSTRequest[Volume](ctx, c, e, opts)
}

// CreateVolume creates a Linode Volume
func (c *Client) CreateVolume(ctx context.Context, opts VolumeCreateOptions) (*Volume, error) {
	return doPOSTRequest[Volume](ctx, c, "volumes", opts)
}

// UpdateVolume updates the Volume with the specified id
func (c *Client) UpdateVolume(ctx context.Context, volumeID int, opts VolumeUpdateOptions) (*Volume, error) {
	e := formatAPIPath("volumes/%d", volumeID)
	return doPUTRequest[Volume](ctx, c, e, opts)
}

// CloneVolume clones a Linode volume
func (c *Client) CloneVolume(ctx context.Context, volumeID int, label string) (*Volume, error) {
	opts := map[string]any{
		"label": label,
	}

	e := formatAPIPath("volumes/%d/clone", volumeID)
	return doPOSTRequest[Volume](ctx, c, e, opts)
}

// DetachVolume detaches a Linode volume
func (c *Client) DetachVolume(ctx context.Context, volumeID int) error {
	e := formatAPIPath("volumes/%d/detach", volumeID)
	return doPOSTRequestNoRequestResponseBody(ctx, c, e)
}

// ResizeVolume resizes an instance to new Linode type
func (c *Client) ResizeVolume(ctx context.Context, volumeID int, size int) error {
	opts := map[string]int{
		"size": size,
	}

	e := formatAPIPath("volumes/%d/resize", volumeID)
	return doPOSTRequestNoResponseBody(ctx, c, e, opts)
}

// DeleteVolume deletes the Volume with the specified id
func (c *Client) DeleteVolume(ctx context.Context, volumeID int) error {
	e := formatAPIPath("volumes/%d", volumeID)
	return doDELETERequest(ctx, c, e)
}
