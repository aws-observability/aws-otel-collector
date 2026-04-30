package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const bmPath = "/v2/bare-metals"

// BareMetalServerService is the interface to interact with the Bare Metal endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/baremetal
type BareMetalServerService interface {
	Create(ctx context.Context, bmCreate *BareMetalCreate) (*BareMetalServer, *http.Response, error)
	Get(ctx context.Context, serverID string) (*BareMetalServer, *http.Response, error)
	Update(ctx context.Context, serverID string, bmReq *BareMetalUpdate) (*BareMetalServer, *http.Response, error)
	Delete(ctx context.Context, serverID string) error
	List(ctx context.Context, options *ListOptions) ([]BareMetalServer, *Meta, *http.Response, error)

	GetBandwidth(ctx context.Context, serverID string) (*Bandwidth, *http.Response, error)
	GetUserData(ctx context.Context, serverID string) (*UserData, *http.Response, error)
	GetVNCUrl(ctx context.Context, serverID string) (*VNCUrl, *http.Response, error)

	ListIPv4s(ctx context.Context, serverID string, options *ListOptions) ([]IPv4, *Meta, *http.Response, error)
	ListIPv6s(ctx context.Context, serverID string, options *ListOptions) ([]IPv6, *Meta, *http.Response, error)

	Halt(ctx context.Context, serverID string) error
	Reboot(ctx context.Context, serverID string) error
	Start(ctx context.Context, serverID string) error
	Reinstall(ctx context.Context, serverID string) (*BareMetalServer, *http.Response, error)

	MassStart(ctx context.Context, serverList []string) error
	MassHalt(ctx context.Context, serverList []string) error
	MassReboot(ctx context.Context, serverList []string) error

	GetUpgrades(ctx context.Context, serverID string) (*Upgrades, *http.Response, error)

	ListVPCInfo(ctx context.Context, serverID string) ([]VPCInfo, *http.Response, error)
	AttachVPC(ctx context.Context, serverID, vpcID string) error
	DetachVPC(ctx context.Context, serverID, vpcID string) error

	// Deprecated: VPC2 is no longer supported
	ListVPC2Info(ctx context.Context, serverID string) ([]VPC2Info, *http.Response, error)

	// Deprecated: VPC2 is no longer supported
	AttachVPC2(ctx context.Context, serverID string, vpc2Req *AttachVPC2Req) error

	// Deprecated: VPC2 is no longer supported
	DetachVPC2(ctx context.Context, serverID, vpcID string) error
}

// BareMetalServerServiceHandler handles interaction with the Bare Metal methods for the Vultr API
type BareMetalServerServiceHandler struct {
	client *Client
}

// BareMetalServer represents a Bare Metal server on Vultr
type BareMetalServer struct {
	ID              string   `json:"id"`
	Os              string   `json:"os"`
	RAM             string   `json:"ram"`
	Disk            string   `json:"disk"`
	MainIP          string   `json:"main_ip"`
	CPUCount        int      `json:"cpu_count"`
	Region          string   `json:"region"`
	DefaultPassword string   `json:"default_password"`
	DateCreated     string   `json:"date_created"`
	Status          string   `json:"status"`
	NetmaskV4       string   `json:"netmask_v4"`
	GatewayV4       string   `json:"gateway_v4"`
	Plan            string   `json:"plan"`
	V6Network       string   `json:"v6_network"`
	V6MainIP        string   `json:"v6_main_ip"`
	V6NetworkSize   int      `json:"v6_network_size"`
	MacAddress      int      `json:"mac_address"`
	Label           string   `json:"label"`
	OsID            int      `json:"os_id"`
	AppID           int      `json:"app_id"`
	ImageID         string   `json:"image_id"`
	SnapshotID      string   `json:"snapshot_id"`
	Features        []string `json:"features"`
	Tags            []string `json:"tags"`
	UserScheme      string   `json:"user_scheme"`
}

// BareMetalCreate represents the optional parameters that can be set when creating a Bare Metal server
type BareMetalCreate struct {
	Region          string            `json:"region"`
	Plan            string            `json:"plan"`
	OsID            int               `json:"os_id,omitempty"`
	StartupScriptID string            `json:"script_id,omitempty"`
	SnapshotID      string            `json:"snapshot_id,omitempty"`
	EnableIPv6      *bool             `json:"enable_ipv6,omitempty"`
	Label           string            `json:"label,omitempty"`
	SSHKeyIDs       []string          `json:"sshkey_id,omitempty"`
	AppID           int               `json:"app_id,omitempty"`
	ImageID         string            `json:"image_id,omitempty"`
	UserData        string            `json:"user_data,omitempty"`
	ActivationEmail *bool             `json:"activation_email,omitempty"`
	Hostname        string            `json:"hostname,omitempty"`
	MdiskMode       string            `json:"mdisk_mode,omitempty"`
	ReservedIPv4    string            `json:"reserved_ipv4,omitempty"`
	PersistentPxe   *bool             `json:"persistent_pxe,omitempty"`
	IPXEChainURL    string            `json:"ipxe_chain_url,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
	UserScheme      string            `json:"user_scheme,omitempty"`
	AppVariables    map[string]string `json:"app_variables,omitempty"`

	// Deprecated: VPC2 is no longer supported
	AttachVPC2 []string `json:"attach_vpc2,omitempty"`

	// Deprecated: VPC2 is no longer supported
	DetachVPC2 []string `json:"detach_vpc2,omitempty"`

	// Deprecated: VPC2 is no longer supported
	EnableVPC2 *bool `json:"enable_vpc2,omitempty"`
}

// BareMetalUpdate represents the optional parameters that can be set when updating a Bare Metal server
type BareMetalUpdate struct {
	OsID         int      `json:"os_id,omitempty"`
	EnableIPv6   *bool    `json:"enable_ipv6,omitempty"`
	Label        string   `json:"label,omitempty"`
	AppID        int      `json:"app_id,omitempty"`
	ImageID      string   `json:"image_id,omitempty"`
	UserData     string   `json:"user_data,omitempty"`
	IPXEChainURL string   `json:"ipxe_chain_url,omitempty"`
	MdiskMode    string   `json:"mdisk_mode,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	UserScheme   string   `json:"user_scheme,omitempty"`

	// Deprecated: VPC2 is no longer supported
	AttachVPC2 []string `json:"attach_vpc2,omitempty"`

	// Deprecated: VPC2 is no longer supported
	DetachVPC2 []string `json:"detach_vpc2,omitempty"`

	// Deprecated: VPC2 is no longer supported
	EnableVPC2 *bool `json:"enable_vpc2,omitempty"`
}

// BareMetalServerBandwidth represents bandwidth information for a Bare Metal server
type BareMetalServerBandwidth struct {
	IncomingBytes int `json:"incoming_bytes"`
	OutgoingBytes int `json:"outgoing_bytes"`
}

type bareMetalsBase struct {
	BareMetals []BareMetalServer `json:"bare_metals"`
	Meta       *Meta             `json:"meta"`
}

type bareMetalBase struct {
	BareMetal *BareMetalServer `json:"bare_metal"`
}

// BMBareMetalBase represents the base struct for a Bare Metal server
type BMBareMetalBase struct {
	BareMetalBandwidth map[string]BareMetalServerBandwidth `json:"bandwidth"`
}

type vncBase struct {
	VNCUrl *VNCUrl `json:"vnc"`
}

// VNCUrl contains the URL for a given Bare Metals VNC
type VNCUrl struct {
	URL string `json:"url"`
}

// Create a new Bare Metal server.
func (b *BareMetalServerServiceHandler) Create(ctx context.Context, bmCreate *BareMetalCreate) (*BareMetalServer, *http.Response, error) {
	req, err := b.client.NewRequest(ctx, http.MethodPost, bmPath, bmCreate)
	if err != nil {
		return nil, nil, err
	}

	bm := new(bareMetalBase)
	resp, err := b.client.DoWithContext(ctx, req, bm)
	if err != nil {
		return nil, resp, err
	}

	return bm.BareMetal, resp, nil
}

// Get information for a Bare Metal instance.
func (b *BareMetalServerServiceHandler) Get(ctx context.Context, serverID string) (*BareMetalServer, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", bmPath, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	bms := new(bareMetalBase)
	resp, err := b.client.DoWithContext(ctx, req, bms)
	if err != nil {
		return nil, resp, err
	}

	return bms.BareMetal, resp, nil
}

// Update a Bare Metal server
func (b *BareMetalServerServiceHandler) Update(ctx context.Context, serverID string, bmReq *BareMetalUpdate) (*BareMetalServer, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s", bmPath, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodPatch, uri, bmReq)
	if err != nil {
		return nil, nil, err
	}

	bms := new(bareMetalBase)
	resp, err := b.client.DoWithContext(ctx, req, bms)
	if err != nil {
		return nil, resp, err
	}

	return bms.BareMetal, resp, nil
}

// Delete a Bare Metal server.
func (b *BareMetalServerServiceHandler) Delete(ctx context.Context, serverID string) error {
	uri := fmt.Sprintf("%s/%s", bmPath, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return nil
	}
	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}

// List all Bare Metal instances in your account.
func (b *BareMetalServerServiceHandler) List(ctx context.Context, options *ListOptions) ([]BareMetalServer, *Meta, *http.Response, error) { //nolint:dupl,lll
	req, err := b.client.NewRequest(ctx, http.MethodGet, bmPath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	bms := new(bareMetalsBase)
	resp, err := b.client.DoWithContext(ctx, req, bms)
	if err != nil {
		return nil, nil, resp, err
	}

	return bms.BareMetals, bms.Meta, resp, nil
}

// GetBandwidth  used by a Bare Metal server.
func (b *BareMetalServerServiceHandler) GetBandwidth(ctx context.Context, serverID string) (*Bandwidth, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/bandwidth", bmPath, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	bms := new(Bandwidth)
	resp, err := b.client.DoWithContext(ctx, req, &bms)
	if err != nil {
		return nil, resp, err
	}

	return bms, resp, nil
}

// GetUserData for a Bare Metal server. The userdata returned will be in base64 encoding.
func (b *BareMetalServerServiceHandler) GetUserData(ctx context.Context, serverID string) (*UserData, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/user-data", bmPath, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	userData := new(userDataBase)
	resp, err := b.client.DoWithContext(ctx, req, userData)
	if err != nil {
		return nil, nil, err
	}

	return userData.UserData, resp, nil
}

// GetVNCUrl gets the vnc url for a given Bare Metal server.
func (b *BareMetalServerServiceHandler) GetVNCUrl(ctx context.Context, serverID string) (*VNCUrl, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/vnc", bmPath, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	vnc := new(vncBase)
	resp, err := b.client.DoWithContext(ctx, req, vnc)
	if err != nil {
		return nil, resp, err
	}

	return vnc.VNCUrl, resp, nil
}

// ListIPv4s information of a Bare Metal server.
// IP information is only available for Bare Metal servers in the "active" state.
func (b *BareMetalServerServiceHandler) ListIPv4s(ctx context.Context, serverID string, options *ListOptions) ([]IPv4, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/ipv4", bmPath, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	ipv4 := new(ipBase)
	resp, err := b.client.DoWithContext(ctx, req, ipv4)
	if err != nil {
		return nil, nil, resp, err
	}

	return ipv4.IPv4s, ipv4.Meta, resp, nil
}

// ListIPv6s information of a Bare Metal server.
// IP information is only available for Bare Metal servers in the "active" state.
// If the Bare Metal server does not have IPv6 enabled, then an empty array is returned.
func (b *BareMetalServerServiceHandler) ListIPv6s(ctx context.Context, serverID string, options *ListOptions) ([]IPv6, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/ipv6", bmPath, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	ipv6 := new(ipBase)
	resp, err := b.client.DoWithContext(ctx, req, ipv6)
	if err != nil {
		return nil, nil, resp, err
	}

	return ipv6.IPv6s, ipv6.Meta, resp, nil
}

// Halt a Bare Metal server.
// This is a hard power off, meaning that the power to the machine is severed.
// The data on the machine will not be modified, and you will still be billed for the machine.
func (b *BareMetalServerServiceHandler) Halt(ctx context.Context, serverID string) error {
	uri := fmt.Sprintf("%s/%s/halt", bmPath, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}

// Reboot a Bare Metal server. This is a hard reboot, which means that the server is powered off, then back on.
func (b *BareMetalServerServiceHandler) Reboot(ctx context.Context, serverID string) error {
	uri := fmt.Sprintf("%s/%s/reboot", bmPath, serverID)

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}

// Start a Bare Metal server.
func (b *BareMetalServerServiceHandler) Start(ctx context.Context, serverID string) error {
	uri := fmt.Sprintf("%s/%s/start", bmPath, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}

// Reinstall the operating system on a Bare Metal server.
// All data will be permanently lost, but the IP address will remain the same.
func (b *BareMetalServerServiceHandler) Reinstall(ctx context.Context, serverID string) (*BareMetalServer, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/reinstall", bmPath, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	bms := new(bareMetalBase)
	resp, err := b.client.DoWithContext(ctx, req, bms)
	if err != nil {
		return nil, resp, err
	}

	return bms.BareMetal, resp, nil
}

// MassStart will start a list of Bare Metal servers the machine is already running, it will be restarted.
func (b *BareMetalServerServiceHandler) MassStart(ctx context.Context, serverList []string) error {
	uri := fmt.Sprintf("%s/start", bmPath)

	reqBody := RequestBody{"baremetal_ids": serverList}
	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, reqBody)
	if err != nil {
		return err
	}
	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}

// MassHalt a list of Bare Metal servers.
func (b *BareMetalServerServiceHandler) MassHalt(ctx context.Context, serverList []string) error {
	uri := fmt.Sprintf("%s/halt", bmPath)

	reqBody := RequestBody{"baremetal_ids": serverList}
	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, reqBody)
	if err != nil {
		return err
	}
	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}

// MassReboot a list of Bare Metal servers.
func (b *BareMetalServerServiceHandler) MassReboot(ctx context.Context, serverList []string) error {
	uri := fmt.Sprintf("%s/reboot", bmPath)

	reqBody := RequestBody{"baremetal_ids": serverList}
	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, reqBody)
	if err != nil {
		return err
	}
	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}

// GetUpgrades that are available for a Bare Metal server.
func (b *BareMetalServerServiceHandler) GetUpgrades(ctx context.Context, serverID string) (*Upgrades, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/upgrades", bmPath, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	upgrades := new(upgradeBase)
	resp, err := b.client.DoWithContext(ctx, req, upgrades)
	if err != nil {
		return nil, resp, err
	}

	return upgrades.Upgrades, resp, nil
}

// ListVPCInfo will list all currently attached VPC IP information for the
// given bare metal server.
func (b *BareMetalServerServiceHandler) ListVPCInfo(ctx context.Context, serverID string) ([]VPCInfo, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/vpcs", bmPath, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	vpcs := new(vpcInfoBase)
	resp, err := b.client.DoWithContext(ctx, req, vpcs)
	if err != nil {
		return nil, resp, err
	}

	return vpcs.VPCs, resp, nil
}

// AttachVPC serves to attach a VPC to a bare metal server.
func (b *BareMetalServerServiceHandler) AttachVPC(ctx context.Context, serverID, vpcID string) error {
	uri := fmt.Sprintf("%s/%s/vpcs/attach", bmPath, serverID)
	body := RequestBody{"vpc_id": vpcID}

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}

// DetachVPC will detach a VPC from a bare metal server.
func (b *BareMetalServerServiceHandler) DetachVPC(ctx context.Context, serverID, vpcID string) error {
	uri := fmt.Sprintf("%s/%s/vpcs/detach", bmPath, serverID)
	body := RequestBody{"vpc_id": vpcID}

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}

// ListVPC2Info currently attached to a Bare Metal server.
//
// Deprecated: VPC2 is no longer supported
func (b *BareMetalServerServiceHandler) ListVPC2Info(ctx context.Context, serverID string) ([]VPC2Info, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/vpc2", bmPath, serverID)
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	vpcs := new(vpc2InfoBase)
	resp, err := b.client.DoWithContext(ctx, req, vpcs)
	if err != nil {
		return nil, resp, err
	}

	return vpcs.VPCs, resp, nil
}

// AttachVPC2 to a Bare Metal server.
//
// Deprecated: VPC2 is no longer supported
func (b *BareMetalServerServiceHandler) AttachVPC2(ctx context.Context, serverID string, vpc2Req *AttachVPC2Req) error {
	uri := fmt.Sprintf("%s/%s/vpc2/attach", bmPath, serverID)

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, vpc2Req)
	if err != nil {
		return err
	}

	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}

// DetachVPC2 from a Bare Metal server.
//
// Deprecated: VPC2 is no longer supported
func (b *BareMetalServerServiceHandler) DetachVPC2(ctx context.Context, serverID, vpcID string) error {
	uri := fmt.Sprintf("%s/%s/vpc2/detach", bmPath, serverID)
	body := RequestBody{"vpc_id": vpcID}

	req, err := b.client.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	_, err = b.client.DoWithContext(ctx, req, nil)
	return err
}
