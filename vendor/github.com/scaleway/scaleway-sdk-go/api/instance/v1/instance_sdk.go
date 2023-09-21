// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package instance provides methods and message types of the instance v1 API.
package instance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ fmt.Stringer
	_ json.Unmarshaler
	_ url.URL
	_ net.IP
	_ http.Header
	_ bytes.Reader
	_ time.Time
	_ = strings.Join

	_ scw.ScalewayRequest
	_ marshaler.Duration
	_ scw.File
	_ = parameter.AddToQuery
	_ = namegenerator.GetRandomName
)

// API: instance API.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type Arch string

const (
	ArchX86_64 = Arch("x86_64")
	ArchArm    = Arch("arm")
)

func (enum Arch) String() string {
	if enum == "" {
		// return default value if empty
		return "x86_64"
	}
	return string(enum)
}

func (enum Arch) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Arch) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Arch(Arch(tmp).String())
	return nil
}

type BootType string

const (
	BootTypeLocal      = BootType("local")
	BootTypeBootscript = BootType("bootscript")
	BootTypeRescue     = BootType("rescue")
)

func (enum BootType) String() string {
	if enum == "" {
		// return default value if empty
		return "local"
	}
	return string(enum)
}

func (enum BootType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BootType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BootType(BootType(tmp).String())
	return nil
}

<<<<<<< HEAD
=======
type IPState string

const (
	IPStateUnknownState = IPState("unknown_state")
	IPStateDetached     = IPState("detached")
	IPStateAttached     = IPState("attached")
	IPStatePending      = IPState("pending")
	IPStateError        = IPState("error")
)

func (enum IPState) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_state"
	}
	return string(enum)
}

func (enum IPState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPState(IPState(tmp).String())
	return nil
}

type IPType string

const (
	IPTypeUnknownIptype = IPType("unknown_iptype")
	IPTypeNat           = IPType("nat")
	IPTypeRoutedIPv4    = IPType("routed_ipv4")
	IPTypeRoutedIPv6    = IPType("routed_ipv6")
)

func (enum IPType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_iptype"
	}
	return string(enum)
}

func (enum IPType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPType(IPType(tmp).String())
	return nil
}

>>>>>>> main
type ImageState string

const (
	ImageStateAvailable = ImageState("available")
	ImageStateCreating  = ImageState("creating")
	ImageStateError     = ImageState("error")
)

func (enum ImageState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum ImageState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ImageState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ImageState(ImageState(tmp).String())
	return nil
}

type ListServersRequestOrder string

const (
	ListServersRequestOrderCreationDateDesc     = ListServersRequestOrder("creation_date_desc")
	ListServersRequestOrderCreationDateAsc      = ListServersRequestOrder("creation_date_asc")
	ListServersRequestOrderModificationDateDesc = ListServersRequestOrder("modification_date_desc")
	ListServersRequestOrderModificationDateAsc  = ListServersRequestOrder("modification_date_asc")
)

func (enum ListServersRequestOrder) String() string {
	if enum == "" {
		// return default value if empty
		return "creation_date_desc"
	}
	return string(enum)
}

func (enum ListServersRequestOrder) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServersRequestOrder) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServersRequestOrder(ListServersRequestOrder(tmp).String())
	return nil
}

type PlacementGroupPolicyMode string

const (
	PlacementGroupPolicyModeOptional = PlacementGroupPolicyMode("optional")
	PlacementGroupPolicyModeEnforced = PlacementGroupPolicyMode("enforced")
)

func (enum PlacementGroupPolicyMode) String() string {
	if enum == "" {
		// return default value if empty
		return "optional"
	}
	return string(enum)
}

func (enum PlacementGroupPolicyMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlacementGroupPolicyMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlacementGroupPolicyMode(PlacementGroupPolicyMode(tmp).String())
	return nil
}

type PlacementGroupPolicyType string

const (
	PlacementGroupPolicyTypeMaxAvailability = PlacementGroupPolicyType("max_availability")
	PlacementGroupPolicyTypeLowLatency      = PlacementGroupPolicyType("low_latency")
)

func (enum PlacementGroupPolicyType) String() string {
	if enum == "" {
		// return default value if empty
		return "max_availability"
	}
	return string(enum)
}

func (enum PlacementGroupPolicyType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlacementGroupPolicyType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlacementGroupPolicyType(PlacementGroupPolicyType(tmp).String())
	return nil
}

type PrivateNICState string

const (
	PrivateNICStateAvailable    = PrivateNICState("available")
	PrivateNICStateSyncing      = PrivateNICState("syncing")
	PrivateNICStateSyncingError = PrivateNICState("syncing_error")
)

func (enum PrivateNICState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum PrivateNICState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PrivateNICState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PrivateNICState(PrivateNICState(tmp).String())
	return nil
}

type SecurityGroupPolicy string

const (
	SecurityGroupPolicyAccept = SecurityGroupPolicy("accept")
	SecurityGroupPolicyDrop   = SecurityGroupPolicy("drop")
)

func (enum SecurityGroupPolicy) String() string {
	if enum == "" {
		// return default value if empty
		return "accept"
	}
	return string(enum)
}

func (enum SecurityGroupPolicy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupPolicy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupPolicy(SecurityGroupPolicy(tmp).String())
	return nil
}

type SecurityGroupRuleAction string

const (
	SecurityGroupRuleActionAccept = SecurityGroupRuleAction("accept")
	SecurityGroupRuleActionDrop   = SecurityGroupRuleAction("drop")
)

func (enum SecurityGroupRuleAction) String() string {
	if enum == "" {
		// return default value if empty
		return "accept"
	}
	return string(enum)
}

func (enum SecurityGroupRuleAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupRuleAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupRuleAction(SecurityGroupRuleAction(tmp).String())
	return nil
}

type SecurityGroupRuleDirection string

const (
	SecurityGroupRuleDirectionInbound  = SecurityGroupRuleDirection("inbound")
	SecurityGroupRuleDirectionOutbound = SecurityGroupRuleDirection("outbound")
)

func (enum SecurityGroupRuleDirection) String() string {
	if enum == "" {
		// return default value if empty
		return "inbound"
	}
	return string(enum)
}

func (enum SecurityGroupRuleDirection) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupRuleDirection) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupRuleDirection(SecurityGroupRuleDirection(tmp).String())
	return nil
}

type SecurityGroupRuleProtocol string

const (
	SecurityGroupRuleProtocolTCP  = SecurityGroupRuleProtocol("TCP")
	SecurityGroupRuleProtocolUDP  = SecurityGroupRuleProtocol("UDP")
	SecurityGroupRuleProtocolICMP = SecurityGroupRuleProtocol("ICMP")
	SecurityGroupRuleProtocolANY  = SecurityGroupRuleProtocol("ANY")
)

func (enum SecurityGroupRuleProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return "TCP"
	}
	return string(enum)
}

func (enum SecurityGroupRuleProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupRuleProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupRuleProtocol(SecurityGroupRuleProtocol(tmp).String())
	return nil
}

type SecurityGroupState string

const (
	SecurityGroupStateAvailable    = SecurityGroupState("available")
	SecurityGroupStateSyncing      = SecurityGroupState("syncing")
	SecurityGroupStateSyncingError = SecurityGroupState("syncing_error")
)

func (enum SecurityGroupState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum SecurityGroupState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupState(SecurityGroupState(tmp).String())
	return nil
}

type ServerAction string

const (
	ServerActionPoweron     = ServerAction("poweron")
	ServerActionBackup      = ServerAction("backup")
	ServerActionStopInPlace = ServerAction("stop_in_place")
	ServerActionPoweroff    = ServerAction("poweroff")
	ServerActionTerminate   = ServerAction("terminate")
	ServerActionReboot      = ServerAction("reboot")
)

func (enum ServerAction) String() string {
	if enum == "" {
		// return default value if empty
		return "poweron"
	}
	return string(enum)
}

func (enum ServerAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerAction(ServerAction(tmp).String())
	return nil
}

<<<<<<< HEAD
=======
type ServerIPIPFamily string

const (
	ServerIPIPFamilyInet  = ServerIPIPFamily("inet")
	ServerIPIPFamilyInet6 = ServerIPIPFamily("inet6")
)

func (enum ServerIPIPFamily) String() string {
	if enum == "" {
		// return default value if empty
		return "inet"
	}
	return string(enum)
}

func (enum ServerIPIPFamily) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerIPIPFamily) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerIPIPFamily(ServerIPIPFamily(tmp).String())
	return nil
}

type ServerIPProvisioningMode string

const (
	ServerIPProvisioningModeManual = ServerIPProvisioningMode("manual")
	ServerIPProvisioningModeDHCP   = ServerIPProvisioningMode("dhcp")
	ServerIPProvisioningModeSlaac  = ServerIPProvisioningMode("slaac")
)

func (enum ServerIPProvisioningMode) String() string {
	if enum == "" {
		// return default value if empty
		return "manual"
	}
	return string(enum)
}

func (enum ServerIPProvisioningMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerIPProvisioningMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerIPProvisioningMode(ServerIPProvisioningMode(tmp).String())
	return nil
}

>>>>>>> main
type ServerState string

const (
	ServerStateRunning        = ServerState("running")
	ServerStateStopped        = ServerState("stopped")
	ServerStateStoppedInPlace = ServerState("stopped in place")
	ServerStateStarting       = ServerState("starting")
	ServerStateStopping       = ServerState("stopping")
	ServerStateLocked         = ServerState("locked")
)

func (enum ServerState) String() string {
	if enum == "" {
		// return default value if empty
		return "running"
	}
	return string(enum)
}

func (enum ServerState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerState(ServerState(tmp).String())
	return nil
}

type ServerTypesAvailability string

const (
	ServerTypesAvailabilityAvailable = ServerTypesAvailability("available")
	ServerTypesAvailabilityScarce    = ServerTypesAvailability("scarce")
	ServerTypesAvailabilityShortage  = ServerTypesAvailability("shortage")
)

func (enum ServerTypesAvailability) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum ServerTypesAvailability) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerTypesAvailability) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerTypesAvailability(ServerTypesAvailability(tmp).String())
	return nil
}

type SnapshotState string

const (
	SnapshotStateAvailable    = SnapshotState("available")
	SnapshotStateSnapshotting = SnapshotState("snapshotting")
	SnapshotStateError        = SnapshotState("error")
	SnapshotStateInvalidData  = SnapshotState("invalid_data")
	SnapshotStateImporting    = SnapshotState("importing")
	SnapshotStateExporting    = SnapshotState("exporting")
)

func (enum SnapshotState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum SnapshotState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SnapshotState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SnapshotState(SnapshotState(tmp).String())
	return nil
}

type SnapshotVolumeType string

const (
	SnapshotVolumeTypeUnknownVolumeType = SnapshotVolumeType("unknown_volume_type")
	SnapshotVolumeTypeLSSD              = SnapshotVolumeType("l_ssd")
	SnapshotVolumeTypeBSSD              = SnapshotVolumeType("b_ssd")
	SnapshotVolumeTypeUnified           = SnapshotVolumeType("unified")
)

func (enum SnapshotVolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_volume_type"
	}
	return string(enum)
}

func (enum SnapshotVolumeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SnapshotVolumeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SnapshotVolumeType(SnapshotVolumeType(tmp).String())
	return nil
}

type TaskStatus string

const (
	TaskStatusPending = TaskStatus("pending")
	TaskStatusStarted = TaskStatus("started")
	TaskStatusSuccess = TaskStatus("success")
	TaskStatusFailure = TaskStatus("failure")
	TaskStatusRetry   = TaskStatus("retry")
)

func (enum TaskStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "pending"
	}
	return string(enum)
}

func (enum TaskStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TaskStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TaskStatus(TaskStatus(tmp).String())
	return nil
}

type VolumeServerState string

const (
	VolumeServerStateAvailable    = VolumeServerState("available")
	VolumeServerStateSnapshotting = VolumeServerState("snapshotting")
	VolumeServerStateError        = VolumeServerState("error")
	VolumeServerStateFetching     = VolumeServerState("fetching")
	VolumeServerStateResizing     = VolumeServerState("resizing")
	VolumeServerStateSaving       = VolumeServerState("saving")
	VolumeServerStateHotsyncing   = VolumeServerState("hotsyncing")
)

func (enum VolumeServerState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum VolumeServerState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeServerState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeServerState(VolumeServerState(tmp).String())
	return nil
}

type VolumeServerVolumeType string

const (
	VolumeServerVolumeTypeLSSD = VolumeServerVolumeType("l_ssd")
	VolumeServerVolumeTypeBSSD = VolumeServerVolumeType("b_ssd")
)

func (enum VolumeServerVolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return "l_ssd"
	}
	return string(enum)
}

func (enum VolumeServerVolumeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeServerVolumeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeServerVolumeType(VolumeServerVolumeType(tmp).String())
	return nil
}

type VolumeState string

const (
	VolumeStateAvailable    = VolumeState("available")
	VolumeStateSnapshotting = VolumeState("snapshotting")
	VolumeStateError        = VolumeState("error")
	VolumeStateFetching     = VolumeState("fetching")
	VolumeStateResizing     = VolumeState("resizing")
	VolumeStateSaving       = VolumeState("saving")
	VolumeStateHotsyncing   = VolumeState("hotsyncing")
)

func (enum VolumeState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum VolumeState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeState(VolumeState(tmp).String())
	return nil
}

type VolumeVolumeType string

const (
	VolumeVolumeTypeLSSD    = VolumeVolumeType("l_ssd")
	VolumeVolumeTypeBSSD    = VolumeVolumeType("b_ssd")
	VolumeVolumeTypeUnified = VolumeVolumeType("unified")
)

func (enum VolumeVolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return "l_ssd"
	}
	return string(enum)
}

func (enum VolumeVolumeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeVolumeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeVolumeType(VolumeVolumeType(tmp).String())
	return nil
}

// Bootscript: bootscript.
type Bootscript struct {
<<<<<<< HEAD
	// Bootcmdargs: the bootscript arguments.
	Bootcmdargs string `json:"bootcmdargs"`
	// Default: dispmay if the bootscript is the default bootscript if no other boot option is configured.
	Default bool `json:"default"`
	// Dtb: provide information regarding a Device Tree Binary (dtb) for use with C1 servers.
	Dtb string `json:"dtb"`
	// ID: the bootscript ID.
	ID string `json:"id"`
	// Initrd: the initrd (initial ramdisk) configuration.
	Initrd string `json:"initrd"`
	// Kernel: the server kernel version.
	Kernel string `json:"kernel"`
	// Organization: the bootscript organization ID.
	Organization string `json:"organization"`
	// Project: the bootscript project ID.
	Project string `json:"project"`
	// Public: provide information if the bootscript is public.
	Public bool `json:"public"`
	// Title: the bootscript title.
	Title string `json:"title"`
	// Arch: the bootscript arch.
	// Default value: x86_64
	Arch Arch `json:"arch"`
	// Zone: the zone in which is the bootscript.
=======
	// Bootcmdargs: bootscript arguments.
	Bootcmdargs string `json:"bootcmdargs"`
	// Default: display if the bootscript is the default bootscript (if no other boot option is configured).
	Default bool `json:"default"`
	// Dtb: provide information regarding a Device Tree Binary (DTB) for use with C1 servers.
	Dtb string `json:"dtb"`
	// ID: bootscript ID.
	ID string `json:"id"`
	// Initrd: initrd (initial ramdisk) configuration.
	Initrd string `json:"initrd"`
	// Kernel: instance kernel version.
	Kernel string `json:"kernel"`
	// Organization: bootscript Organization ID.
	Organization string `json:"organization"`
	// Project: bootscript Project ID.
	Project string `json:"project"`
	// Public: provide information if the bootscript is public.
	Public bool `json:"public"`
	// Title: bootscript title.
	Title string `json:"title"`
	// Arch: bootscript architecture.
	// Default value: x86_64
	Arch Arch `json:"arch"`
	// Zone: zone in which the bootscript is located.
>>>>>>> main
	Zone scw.Zone `json:"zone"`
}

type CreateIPResponse struct {
	IP *IP `json:"ip"`
}

type CreateImageResponse struct {
	Image *Image `json:"image"`
}

type CreatePlacementGroupResponse struct {
	PlacementGroup *PlacementGroup `json:"placement_group"`
}

type CreatePrivateNICResponse struct {
	PrivateNic *PrivateNIC `json:"private_nic"`
}

type CreateSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security_group"`
}

type CreateSecurityGroupRuleResponse struct {
	Rule *SecurityGroupRule `json:"rule"`
}

type CreateServerResponse struct {
	Server *Server `json:"server"`
}

type CreateSnapshotResponse struct {
	Snapshot *Snapshot `json:"snapshot"`

	Task *Task `json:"task"`
}

type CreateVolumeResponse struct {
	Volume *Volume `json:"volume"`
}

type Dashboard struct {
	VolumesCount uint32 `json:"volumes_count"`

	RunningServersCount uint32 `json:"running_servers_count"`

	ServersByTypes map[string]uint32 `json:"servers_by_types"`

	ImagesCount uint32 `json:"images_count"`

	SnapshotsCount uint32 `json:"snapshots_count"`

	ServersCount uint32 `json:"servers_count"`

	IPsCount uint32 `json:"ips_count"`

	SecurityGroupsCount uint32 `json:"security_groups_count"`

	IPsUnused uint32 `json:"ips_unused"`

	VolumesLSSDCount uint32 `json:"volumes_l_ssd_count"`

	VolumesBSSDCount uint32 `json:"volumes_b_ssd_count"`

	VolumesLSSDTotalSize scw.Size `json:"volumes_l_ssd_total_size"`

	VolumesBSSDTotalSize scw.Size `json:"volumes_b_ssd_total_size"`

	PrivateNicsCount uint32 `json:"private_nics_count"`

	PlacementGroupsCount uint32 `json:"placement_groups_count"`
}

type ExportSnapshotResponse struct {
	Task *Task `json:"task"`
}

type GetBootscriptResponse struct {
	Bootscript *Bootscript `json:"bootscript"`
}

type GetDashboardResponse struct {
	Dashboard *Dashboard `json:"dashboard"`
}

type GetIPResponse struct {
	IP *IP `json:"ip"`
}

type GetImageResponse struct {
	Image *Image `json:"image"`
}

type GetPlacementGroupResponse struct {
	PlacementGroup *PlacementGroup `json:"placement_group"`
}

<<<<<<< HEAD
type GetPlacementGroupServersResponse struct {
=======
// GetPlacementGroupServersResponse: get placement group servers response.
type GetPlacementGroupServersResponse struct {
	// Servers: instances attached to the placement group.
>>>>>>> main
	Servers []*PlacementGroupServer `json:"servers"`
}

type GetPrivateNICResponse struct {
	PrivateNic *PrivateNIC `json:"private_nic"`
}

type GetSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security_group"`
}

type GetSecurityGroupRuleResponse struct {
	Rule *SecurityGroupRule `json:"rule"`
}

type GetServerResponse struct {
	Server *Server `json:"server"`
}

<<<<<<< HEAD
type GetServerTypesAvailabilityResponse struct {
	Servers map[string]*GetServerTypesAvailabilityResponseAvailability `json:"servers"`
}

type GetServerTypesAvailabilityResponseAvailability struct {
	// Availability:
	// Default value: available
=======
// GetServerTypesAvailabilityResponse: get server types availability response.
type GetServerTypesAvailabilityResponse struct {
	// Servers: map of server types.
	Servers map[string]*GetServerTypesAvailabilityResponseAvailability `json:"servers"`

	TotalCount uint32 `json:"total_count"`
}

type GetServerTypesAvailabilityResponseAvailability struct {
	// Availability: default value: available
>>>>>>> main
	Availability ServerTypesAvailability `json:"availability"`
}

type GetSnapshotResponse struct {
	Snapshot *Snapshot `json:"snapshot"`
}

type GetVolumeResponse struct {
	Volume *Volume `json:"volume"`
}

type IP struct {
	ID string `json:"id"`

	Address net.IP `json:"address"`

	Reverse *string `json:"reverse"`

	Server *ServerSummary `json:"server"`

	Organization string `json:"organization"`

	Tags []string `json:"tags"`

	Project string `json:"project"`
<<<<<<< HEAD
=======
	// Type: default value: unknown_iptype
	Type IPType `json:"type"`
	// State: default value: unknown_state
	State IPState `json:"state"`

	Prefix scw.IPNet `json:"prefix"`
>>>>>>> main

	Zone scw.Zone `json:"zone"`
}

type Image struct {
	ID string `json:"id"`

	Name string `json:"name"`
<<<<<<< HEAD
	// Arch:
	// Default value: x86_64
=======
	// Arch: default value: x86_64
>>>>>>> main
	Arch Arch `json:"arch"`

	CreationDate *time.Time `json:"creation_date"`

	ModificationDate *time.Time `json:"modification_date"`
	// Deprecated
	DefaultBootscript *Bootscript `json:"default_bootscript"`

	ExtraVolumes map[string]*Volume `json:"extra_volumes"`

	FromServer string `json:"from_server"`

	Organization string `json:"organization"`

	Public bool `json:"public"`

	RootVolume *VolumeSummary `json:"root_volume"`
<<<<<<< HEAD
	// State:
	// Default value: available
=======
	// State: default value: available
>>>>>>> main
	State ImageState `json:"state"`

	Project string `json:"project"`

	Tags []string `json:"tags"`

	Zone scw.Zone `json:"zone"`
}

// ListBootscriptsResponse: list bootscripts response.
type ListBootscriptsResponse struct {
	// TotalCount: total number of bootscripts.
	TotalCount uint32 `json:"total_count"`
	// Bootscripts: list of bootscripts.
	Bootscripts []*Bootscript `json:"bootscripts"`
}

// ListIPsResponse: list ips response.
type ListIPsResponse struct {
	// TotalCount: total number of ips.
	TotalCount uint32 `json:"total_count"`
	// IPs: list of ips.
	IPs []*IP `json:"ips"`
}

// ListImagesResponse: list images response.
type ListImagesResponse struct {
	// TotalCount: total number of images.
	TotalCount uint32 `json:"total_count"`
	// Images: list of images.
	Images []*Image `json:"images"`
}

// ListPlacementGroupsResponse: list placement groups response.
type ListPlacementGroupsResponse struct {
	// TotalCount: total number of placement groups.
	TotalCount uint32 `json:"total_count"`
	// PlacementGroups: list of placement groups.
	PlacementGroups []*PlacementGroup `json:"placement_groups"`
}

type ListPrivateNICsResponse struct {
	PrivateNics []*PrivateNIC `json:"private_nics"`

	TotalCount uint64 `json:"total_count"`
}

// ListSecurityGroupRulesResponse: list security group rules response.
type ListSecurityGroupRulesResponse struct {
	// TotalCount: total number of security groups.
	TotalCount uint32 `json:"total_count"`
	// Rules: list of security rules.
	Rules []*SecurityGroupRule `json:"rules"`
}

// ListSecurityGroupsResponse: list security groups response.
type ListSecurityGroupsResponse struct {
	// TotalCount: total number of security groups.
	TotalCount uint32 `json:"total_count"`
	// SecurityGroups: list of security groups.
	SecurityGroups []*SecurityGroup `json:"security_groups"`
}

type ListServerActionsResponse struct {
	Actions []ServerAction `json:"actions"`
}

type ListServerUserDataResponse struct {
	UserData []string `json:"user_data"`
}

// ListServersResponse: list servers response.
type ListServersResponse struct {
<<<<<<< HEAD
	// TotalCount: total number of servers.
	TotalCount uint32 `json:"total_count"`
	// Servers: list of servers.
=======
	// TotalCount: total number of Instances.
	TotalCount uint32 `json:"total_count"`
	// Servers: list of Instances.
>>>>>>> main
	Servers []*Server `json:"servers"`
}

// ListServersTypesResponse: list servers types response.
type ListServersTypesResponse struct {
<<<<<<< HEAD
	// TotalCount: total number of server types.
	TotalCount uint32 `json:"total_count"`
	// Servers: list of server types.
=======
	// TotalCount: total number of Instance types.
	TotalCount uint32 `json:"total_count"`
	// Servers: list of Instance types.
>>>>>>> main
	Servers map[string]*ServerType `json:"servers"`
}

// ListSnapshotsResponse: list snapshots response.
type ListSnapshotsResponse struct {
	// TotalCount: total number of snapshots.
	TotalCount uint32 `json:"total_count"`
	// Snapshots: list of snapshots.
	Snapshots []*Snapshot `json:"snapshots"`
}

// ListVolumesResponse: list volumes response.
type ListVolumesResponse struct {
	// TotalCount: total number of volumes.
	TotalCount uint32 `json:"total_count"`
	// Volumes: list of volumes.
	Volumes []*Volume `json:"volumes"`
}

// ListVolumesTypesResponse: list volumes types response.
type ListVolumesTypesResponse struct {
	// TotalCount: total number of volume types.
	TotalCount uint32 `json:"total_count"`
	// Volumes: map of volume types.
	Volumes map[string]*VolumeType `json:"volumes"`
}

type NullableStringValue struct {
	Null bool `json:"null,omitempty"`

	Value string `json:"value,omitempty"`
}

// PlacementGroup: placement group.
type PlacementGroup struct {
<<<<<<< HEAD
	// ID: the placement group unique ID.
	ID string `json:"id"`
	// Name: the placement group name.
	Name string `json:"name"`
	// Organization: the placement group organization ID.
	Organization string `json:"organization"`
	// Project: the placement group project ID.
	Project string `json:"project"`
	// Tags: the placement group tags.
	Tags []string `json:"tags"`
	// PolicyMode: select the failling mode when the placement cannot be respected, either optional or enforced.
=======
	// ID: placement group unique ID.
	ID string `json:"id"`
	// Name: placement group name.
	Name string `json:"name"`
	// Organization: placement group Organization ID.
	Organization string `json:"organization"`
	// Project: placement group Project ID.
	Project string `json:"project"`
	// Tags: placement group tags.
	Tags []string `json:"tags"`
	// PolicyMode: select the failure mode when the placement cannot be respected, either optional or enforced.
>>>>>>> main
	// Default value: optional
	PolicyMode PlacementGroupPolicyMode `json:"policy_mode"`
	// PolicyType: select the behavior of the placement group, either low_latency (group) or max_availability (spread).
	// Default value: max_availability
	PolicyType PlacementGroupPolicyType `json:"policy_type"`
	// PolicyRespected: returns true if the policy is respected, false otherwise.
	PolicyRespected bool `json:"policy_respected"`
<<<<<<< HEAD
	// Zone: the zone in which is the placement group.
	Zone scw.Zone `json:"zone"`
}

type PlacementGroupServer struct {
	ID string `json:"id"`

	Name string `json:"name"`

=======
	// Zone: zone in which the placement group is located.
	Zone scw.Zone `json:"zone"`
}

// PlacementGroupServer: placement group server.
type PlacementGroupServer struct {
	// ID: instance UUID.
	ID string `json:"id"`
	// Name: instance name.
	Name string `json:"name"`
	// PolicyRespected: defines whether the placement group policy is respected (either 1 or 0).
>>>>>>> main
	PolicyRespected bool `json:"policy_respected"`
}

// PrivateNIC: private nic.
type PrivateNIC struct {
<<<<<<< HEAD
	// ID: the private NIC unique ID.
	ID string `json:"id,omitempty"`
	// ServerID: the server the private NIC is attached to.
	ServerID string `json:"server_id,omitempty"`
	// PrivateNetworkID: the private network where the private NIC is attached.
	PrivateNetworkID string `json:"private_network_id,omitempty"`
	// MacAddress: the private NIC MAC address.
	MacAddress string `json:"mac_address,omitempty"`
	// State: the private NIC state.
	// Default value: available
	State PrivateNICState `json:"state,omitempty"`
	// Tags: the private NIC tags.
=======
	// ID: private NIC unique ID.
	ID string `json:"id,omitempty"`
	// ServerID: instance to which the private NIC is attached.
	ServerID string `json:"server_id,omitempty"`
	// PrivateNetworkID: private Network the private NIC is attached to.
	PrivateNetworkID string `json:"private_network_id,omitempty"`
	// MacAddress: private NIC MAC address.
	MacAddress string `json:"mac_address,omitempty"`
	// State: private NIC state.
	// Default value: available
	State PrivateNICState `json:"state,omitempty"`
	// Tags: private NIC tags.
>>>>>>> main
	Tags []string `json:"tags,omitempty"`
}

// SecurityGroup: security group.
type SecurityGroup struct {
<<<<<<< HEAD
	// ID: the security groups' unique ID.
	ID string `json:"id"`
	// Name: the security groups name.
	Name string `json:"name"`
	// Description: the security groups description.
	Description string `json:"description"`
	// EnableDefaultSecurity: true if SMTP is blocked on IPv4 and IPv6.
	EnableDefaultSecurity bool `json:"enable_default_security"`
	// InboundDefaultPolicy: the default inbound policy.
	// Default value: accept
	InboundDefaultPolicy SecurityGroupPolicy `json:"inbound_default_policy"`
	// OutboundDefaultPolicy: the default outbound policy.
	// Default value: accept
	OutboundDefaultPolicy SecurityGroupPolicy `json:"outbound_default_policy"`
	// Organization: the security groups organization ID.
	Organization string `json:"organization"`
	// Project: the security group project ID.
	Project string `json:"project"`
	// Tags: the security group tags.
	Tags []string `json:"tags"`
	// Deprecated: OrganizationDefault: true if it is your default security group for this organization ID.
	OrganizationDefault *bool `json:"organization_default,omitempty"`
	// ProjectDefault: true if it is your default security group for this project ID.
	ProjectDefault bool `json:"project_default"`
	// CreationDate: the security group creation date.
	CreationDate *time.Time `json:"creation_date"`
	// ModificationDate: the security group modification date.
	ModificationDate *time.Time `json:"modification_date"`
	// Servers: list of servers attached to this security group.
	Servers []*ServerSummary `json:"servers"`
	// Stateful: true if the security group is stateful.
=======
	// ID: security group unique ID.
	ID string `json:"id"`
	// Name: security group name.
	Name string `json:"name"`
	// Description: security group description.
	Description string `json:"description"`
	// EnableDefaultSecurity: true if SMTP is blocked on IPv4 and IPv6. This feature is read only, please open a support ticket if you need to make it configurable.
	EnableDefaultSecurity bool `json:"enable_default_security"`
	// InboundDefaultPolicy: default inbound policy.
	// Default value: accept
	InboundDefaultPolicy SecurityGroupPolicy `json:"inbound_default_policy"`
	// OutboundDefaultPolicy: default outbound policy.
	// Default value: accept
	OutboundDefaultPolicy SecurityGroupPolicy `json:"outbound_default_policy"`
	// Organization: security group Organization ID.
	Organization string `json:"organization"`
	// Project: security group Project ID.
	Project string `json:"project"`
	// Tags: security group tags.
	Tags []string `json:"tags"`
	// Deprecated: OrganizationDefault: true if it is your default security group for this Organization ID.
	OrganizationDefault *bool `json:"organization_default,omitempty"`
	// ProjectDefault: true if it is your default security group for this Project ID.
	ProjectDefault bool `json:"project_default"`
	// CreationDate: security group creation date.
	CreationDate *time.Time `json:"creation_date"`
	// ModificationDate: security group modification date.
	ModificationDate *time.Time `json:"modification_date"`
	// Servers: list of Instances attached to this security group.
	Servers []*ServerSummary `json:"servers"`
	// Stateful: defines whether the security group is stateful.
>>>>>>> main
	Stateful bool `json:"stateful"`
	// State: security group state.
	// Default value: available
	State SecurityGroupState `json:"state"`
<<<<<<< HEAD
	// Zone: the zone in which is the security group.
=======
	// Zone: zone in which the security group is located.
>>>>>>> main
	Zone scw.Zone `json:"zone"`
}

type SecurityGroupRule struct {
	ID string `json:"id"`
<<<<<<< HEAD
	// Protocol:
	// Default value: TCP
	Protocol SecurityGroupRuleProtocol `json:"protocol"`
	// Direction:
	// Default value: inbound
	Direction SecurityGroupRuleDirection `json:"direction"`
	// Action:
	// Default value: accept
=======
	// Protocol: default value: TCP
	Protocol SecurityGroupRuleProtocol `json:"protocol"`
	// Direction: default value: inbound
	Direction SecurityGroupRuleDirection `json:"direction"`
	// Action: default value: accept
>>>>>>> main
	Action SecurityGroupRuleAction `json:"action"`

	IPRange scw.IPNet `json:"ip_range"`

	DestPortFrom *uint32 `json:"dest_port_from"`

	DestPortTo *uint32 `json:"dest_port_to"`

	Position uint32 `json:"position"`

	Editable bool `json:"editable"`

	Zone scw.Zone `json:"zone"`
}

type SecurityGroupSummary struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

type SecurityGroupTemplate struct {
	ID string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}

// Server: server.
type Server struct {
<<<<<<< HEAD
	// ID: the server unique ID.
	ID string `json:"id"`
	// Name: the server name.
	Name string `json:"name"`
	// Organization: the server organization ID.
	Organization string `json:"organization"`
	// Project: the server project ID.
	Project string `json:"project"`
	// AllowedActions: provide as list of allowed actions on the server.
	AllowedActions []ServerAction `json:"allowed_actions"`
	// Tags: the server associated tags.
	Tags []string `json:"tags"`
	// CommercialType: the server commercial type (eg. GP1-M).
	CommercialType string `json:"commercial_type"`
	// CreationDate: the server creation date.
	CreationDate *time.Time `json:"creation_date"`
	// DynamicIPRequired: true if a dynamic IP is required.
	DynamicIPRequired bool `json:"dynamic_ip_required"`
	// EnableIPv6: true if IPv6 is enabled.
	EnableIPv6 bool `json:"enable_ipv6"`
	// Hostname: the server host name.
	Hostname string `json:"hostname"`
	// Image: provide information on the server image.
	Image *Image `json:"image"`
	// Protected: the server protection option is activated.
	Protected bool `json:"protected"`
	// PrivateIP: the server private IP address.
	PrivateIP *string `json:"private_ip"`
	// PublicIP: information about the public IP.
	PublicIP *ServerIP `json:"public_ip"`
	// ModificationDate: the server modification date.
	ModificationDate *time.Time `json:"modification_date"`
	// State: the server state.
	// Default value: running
	State ServerState `json:"state"`
	// Location: the server location.
	Location *ServerLocation `json:"location"`
	// IPv6: the server IPv6 address.
	IPv6 *ServerIPv6 `json:"ipv6"`
	// Deprecated: Bootscript: the server bootscript.
	Bootscript *Bootscript `json:"bootscript,omitempty"`
	// BootType: the server boot type.
	// Default value: local
	BootType BootType `json:"boot_type"`
	// Volumes: the server volumes.
	Volumes map[string]*VolumeServer `json:"volumes"`
	// SecurityGroup: the server security group.
	SecurityGroup *SecurityGroupSummary `json:"security_group"`
	// Maintenances: the server planned maintenances.
	Maintenances []*ServerMaintenance `json:"maintenances"`
	// StateDetail: the server state_detail.
	StateDetail string `json:"state_detail"`
	// Arch: the server arch.
	// Default value: x86_64
	Arch Arch `json:"arch"`
	// PlacementGroup: the server placement group.
	PlacementGroup *PlacementGroup `json:"placement_group"`
	// PrivateNics: the server private NICs.
	PrivateNics []*PrivateNIC `json:"private_nics"`
	// Zone: the zone in which is the server.
=======
	// ID: instance unique ID.
	ID string `json:"id"`
	// Name: instance name.
	Name string `json:"name"`
	// Organization: instance Organization ID.
	Organization string `json:"organization"`
	// Project: instance Project ID.
	Project string `json:"project"`
	// AllowedActions: list of allowed actions on the Instance.
	AllowedActions []ServerAction `json:"allowed_actions"`
	// Tags: tags associated with the Instance.
	Tags []string `json:"tags"`
	// CommercialType: instance commercial type (eg. GP1-M).
	CommercialType string `json:"commercial_type"`
	// CreationDate: instance creation date.
	CreationDate *time.Time `json:"creation_date"`
	// DynamicIPRequired: true if a dynamic IPv4 is required.
	DynamicIPRequired bool `json:"dynamic_ip_required"`
	// RoutedIPEnabled: true to configure the instance so it uses the new routed IP mode.
	RoutedIPEnabled bool `json:"routed_ip_enabled"`
	// EnableIPv6: true if IPv6 is enabled.
	EnableIPv6 bool `json:"enable_ipv6"`
	// Hostname: instance host name.
	Hostname string `json:"hostname"`
	// Image: information about the Instance image.
	Image *Image `json:"image"`
	// Protected: defines whether the Instance protection option is activated.
	Protected bool `json:"protected"`
	// PrivateIP: private IP address of the Instance.
	PrivateIP *string `json:"private_ip"`
	// PublicIP: information about the public IP.
	PublicIP *ServerIP `json:"public_ip"`
	// PublicIPs: information about all the public IPs attached to the server.
	PublicIPs []*ServerIP `json:"public_ips"`
	// MacAddress: the server's MAC address.
	MacAddress string `json:"mac_address"`
	// ModificationDate: instance modification date.
	ModificationDate *time.Time `json:"modification_date"`
	// State: instance state.
	// Default value: running
	State ServerState `json:"state"`
	// Location: instance location.
	Location *ServerLocation `json:"location"`
	// IPv6: instance IPv6 address.
	IPv6 *ServerIPv6 `json:"ipv6"`
	// Deprecated: Bootscript: instance bootscript.
	Bootscript *Bootscript `json:"bootscript,omitempty"`
	// BootType: instance boot type.
	// Default value: local
	BootType BootType `json:"boot_type"`
	// Volumes: instance volumes.
	Volumes map[string]*VolumeServer `json:"volumes"`
	// SecurityGroup: instance security group.
	SecurityGroup *SecurityGroupSummary `json:"security_group"`
	// Maintenances: instance planned maintenance.
	Maintenances []*ServerMaintenance `json:"maintenances"`
	// StateDetail: detailed information about the Instance state.
	StateDetail string `json:"state_detail"`
	// Arch: instance architecture.
	// Default value: x86_64
	Arch Arch `json:"arch"`
	// PlacementGroup: instance placement group.
	PlacementGroup *PlacementGroup `json:"placement_group"`
	// PrivateNics: instance private NICs.
	PrivateNics []*PrivateNIC `json:"private_nics"`
	// Zone: zone in which the Instance is located.
>>>>>>> main
	Zone scw.Zone `json:"zone"`
}

// ServerActionRequestVolumeBackupTemplate: server action request. volume backup template.
type ServerActionRequestVolumeBackupTemplate struct {
<<<<<<< HEAD
	// VolumeType: overrides the volume_type of the snapshot for this volume.
=======
	// VolumeType: snapshot's volume type.
	// Overrides the `volume_type` of the snapshot for this volume.
>>>>>>> main
	// If omitted, the volume type of the original volume will be used.
	// Default value: unknown_volume_type
	VolumeType SnapshotVolumeType `json:"volume_type,omitempty"`
}

type ServerActionResponse struct {
	Task *Task `json:"task"`
}

// ServerIP: server. ip.
type ServerIP struct {
<<<<<<< HEAD
	// ID: the unique ID of the IP address.
	ID string `json:"id"`
	// Address: the server public IPv4 IP-Address.
	Address net.IP `json:"address"`
	// Dynamic: true if the IP address is dynamic.
	Dynamic bool `json:"dynamic"`
=======
	// ID: unique ID of the IP address.
	ID string `json:"id,omitempty"`
	// Address: instance's public IP-Address.
	Address net.IP `json:"address,omitempty"`
	// Gateway: gateway's IP address.
	Gateway net.IP `json:"gateway,omitempty"`
	// Netmask: cIDR netmask.
	Netmask string `json:"netmask,omitempty"`
	// Family: IP address family (inet or inet6).
	// Default value: inet
	Family ServerIPIPFamily `json:"family,omitempty"`
	// Dynamic: true if the IP address is dynamic.
	Dynamic bool `json:"dynamic,omitempty"`
	// ProvisioningMode: information about this address provisioning mode.
	// Default value: manual
	ProvisioningMode ServerIPProvisioningMode `json:"provisioning_mode,omitempty"`
>>>>>>> main
}

// ServerIPv6: server. ipv6.
type ServerIPv6 struct {
<<<<<<< HEAD
	// Address: the server IPv6 IP-Address.
	Address net.IP `json:"address"`
	// Gateway: the IPv6 IP-addresses gateway.
	Gateway net.IP `json:"gateway"`
	// Netmask: the IPv6 IP-addresses CIDR netmask.
=======
	// Address: instance IPv6 IP-Address.
	Address net.IP `json:"address"`
	// Gateway: iPv6 IP-addresses gateway.
	Gateway net.IP `json:"gateway"`
	// Netmask: iPv6 IP-addresses CIDR netmask.
>>>>>>> main
	Netmask string `json:"netmask"`
}

type ServerLocation struct {
	ClusterID string `json:"cluster_id"`

	HypervisorID string `json:"hypervisor_id"`

	NodeID string `json:"node_id"`

	PlatformID string `json:"platform_id"`

	ZoneID string `json:"zone_id"`
}

type ServerMaintenance struct {
	Reason string `json:"reason"`
}

type ServerSummary struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

// ServerType: server type.
type ServerType struct {
	// Deprecated: MonthlyPrice: estimated monthly price, for a 30 days month, in Euro.
	MonthlyPrice *float32 `json:"monthly_price,omitempty"`
	// HourlyPrice: hourly price in Euro.
	HourlyPrice float32 `json:"hourly_price"`
<<<<<<< HEAD
	// AltNames: alternative instance name if any.
=======
	// AltNames: alternative Instance name, if any.
>>>>>>> main
	AltNames []string `json:"alt_names"`
	// PerVolumeConstraint: additional volume constraints.
	PerVolumeConstraint *ServerTypeVolumeConstraintsByType `json:"per_volume_constraint"`
	// VolumesConstraint: initial volume constraints.
	VolumesConstraint *ServerTypeVolumeConstraintSizes `json:"volumes_constraint"`
	// Ncpus: number of CPU.
	Ncpus uint32 `json:"ncpus"`
	// Gpu: number of GPU.
	Gpu *uint64 `json:"gpu"`
	// RAM: available RAM in bytes.
	RAM uint64 `json:"ram"`
	// Arch: CPU architecture.
	// Default value: x86_64
	Arch Arch `json:"arch"`
<<<<<<< HEAD
	// Baremetal: true if it is a baremetal instance.
	Baremetal bool `json:"baremetal"`
	// Network: network available for the instance.
=======
	// Baremetal: true if it is a baremetal Instance.
	Baremetal bool `json:"baremetal"`
	// Network: network available for the Instance.
>>>>>>> main
	Network *ServerTypeNetwork `json:"network"`
	// Capabilities: capabilities.
	Capabilities *ServerTypeCapabilities `json:"capabilities"`
}

// ServerTypeCapabilities: server type. capabilities.
type ServerTypeCapabilities struct {
<<<<<<< HEAD
	// BlockStorage: true if server supports block storage.
=======
	// BlockStorage: defines whether the Instance supports block storage.
>>>>>>> main
	BlockStorage *bool `json:"block_storage"`
	// BootTypes: list of supported boot types.
	BootTypes []BootType `json:"boot_types"`
}

// ServerTypeNetwork: server type. network.
type ServerTypeNetwork struct {
	// Interfaces: list of available network interfaces.
	Interfaces []*ServerTypeNetworkInterface `json:"interfaces"`
	// SumInternalBandwidth: total maximum internal bandwidth in bits per seconds.
	SumInternalBandwidth *uint64 `json:"sum_internal_bandwidth"`
	// SumInternetBandwidth: total maximum internet bandwidth in bits per seconds.
	SumInternetBandwidth *uint64 `json:"sum_internet_bandwidth"`
	// IPv6Support: true if IPv6 is enabled.
	IPv6Support bool `json:"ipv6_support"`
}

// ServerTypeNetworkInterface: server type. network. interface.
type ServerTypeNetworkInterface struct {
	// InternalBandwidth: maximum internal bandwidth in bits per seconds.
	InternalBandwidth *uint64 `json:"internal_bandwidth"`
	// InternetBandwidth: maximum internet bandwidth in bits per seconds.
	InternetBandwidth *uint64 `json:"internet_bandwidth"`
}

// ServerTypeVolumeConstraintSizes: server type. volume constraint sizes.
type ServerTypeVolumeConstraintSizes struct {
	// MinSize: minimum volume size in bytes.
	MinSize scw.Size `json:"min_size"`
	// MaxSize: maximum volume size in bytes.
	MaxSize scw.Size `json:"max_size"`
}

// ServerTypeVolumeConstraintsByType: server type. volume constraints by type.
type ServerTypeVolumeConstraintsByType struct {
	// LSSD: local SSD volumes.
	LSSD *ServerTypeVolumeConstraintSizes `json:"l_ssd"`
}

type SetPlacementGroupResponse struct {
	PlacementGroup *PlacementGroup `json:"placement_group"`
}

<<<<<<< HEAD
type SetPlacementGroupServersResponse struct {
=======
// SetPlacementGroupServersResponse: set placement group servers response.
type SetPlacementGroupServersResponse struct {
	// Servers: instances attached to the placement group.
>>>>>>> main
	Servers []*PlacementGroupServer `json:"servers"`
}

// SetSecurityGroupRulesRequestRule: set security group rules request. rule.
type SetSecurityGroupRulesRequestRule struct {
	// ID: UUID of the security rule to update. If no value is provided, a new rule will be created.
	ID *string `json:"id"`
	// Action: action to apply when the rule matches a packet.
	// Default value: accept
	Action SecurityGroupRuleAction `json:"action"`
	// Protocol: protocol family this rule applies to.
	// Default value: TCP
	Protocol SecurityGroupRuleProtocol `json:"protocol"`
	// Direction: direction the rule applies to.
	// Default value: inbound
	Direction SecurityGroupRuleDirection `json:"direction"`
<<<<<<< HEAD
	// IPRange: the range of IP address this rules applies to.
=======
	// IPRange: range of IP addresses these rules apply to.
>>>>>>> main
	IPRange scw.IPNet `json:"ip_range"`
	// DestPortFrom: beginning of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY.
	DestPortFrom *uint32 `json:"dest_port_from"`
	// DestPortTo: end of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY, or if it is equal to dest_port_from.
	DestPortTo *uint32 `json:"dest_port_to"`
	// Position: position of this rule in the security group rules list. If several rules are passed with the same position, the resulting order is undefined.
	Position uint32 `json:"position"`
	// Editable: indicates if this rule is editable. Rules with the value false will be ignored.
	Editable *bool `json:"editable"`
	// Zone: zone of the rule. This field is ignored.
	Zone *scw.Zone `json:"zone"`
}

type SetSecurityGroupRulesResponse struct {
	Rules []*SecurityGroupRule `json:"rules"`
}

// Snapshot: snapshot.
type Snapshot struct {
<<<<<<< HEAD
	// ID: the snapshot ID.
	ID string `json:"id"`
	// Name: the snapshot name.
	Name string `json:"name"`
	// Organization: the snapshot organization ID.
	Organization string `json:"organization"`
	// Project: the snapshot project ID.
	Project string `json:"project"`
	// Tags: the snapshot tags.
	Tags []string `json:"tags"`
	// VolumeType: the snapshot volume type.
	// Default value: l_ssd
	VolumeType VolumeVolumeType `json:"volume_type"`
	// Size: the snapshot size.
	Size scw.Size `json:"size"`
	// State: the snapshot state.
	// Default value: available
	State SnapshotState `json:"state"`
	// BaseVolume: the volume on which the snapshot is based on.
	BaseVolume *SnapshotBaseVolume `json:"base_volume"`
	// CreationDate: the snapshot creation date.
	CreationDate *time.Time `json:"creation_date"`
	// ModificationDate: the snapshot modification date.
	ModificationDate *time.Time `json:"modification_date"`
	// Zone: the snapshot zone.
	Zone scw.Zone `json:"zone"`
	// ErrorReason: the reason for the failed snapshot import.
=======
	// ID: snapshot ID.
	ID string `json:"id"`
	// Name: snapshot name.
	Name string `json:"name"`
	// Organization: snapshot Organization ID.
	Organization string `json:"organization"`
	// Project: snapshot Project ID.
	Project string `json:"project"`
	// Tags: snapshot tags.
	Tags []string `json:"tags"`
	// VolumeType: snapshot volume type.
	// Default value: l_ssd
	VolumeType VolumeVolumeType `json:"volume_type"`
	// Size: snapshot size.
	Size scw.Size `json:"size"`
	// State: snapshot state.
	// Default value: available
	State SnapshotState `json:"state"`
	// BaseVolume: volume on which the snapshot is based on.
	BaseVolume *SnapshotBaseVolume `json:"base_volume"`
	// CreationDate: snapshot creation date.
	CreationDate *time.Time `json:"creation_date"`
	// ModificationDate: snapshot modification date.
	ModificationDate *time.Time `json:"modification_date"`
	// Zone: snapshot zone.
	Zone scw.Zone `json:"zone"`
	// ErrorReason: reason for the failed snapshot import.
>>>>>>> main
	ErrorReason *string `json:"error_reason"`
}

// SnapshotBaseVolume: snapshot. base volume.
type SnapshotBaseVolume struct {
<<<<<<< HEAD
	// ID: the volume ID on which the snapshot is based on.
	ID string `json:"id"`
	// Name: the volume name on which the snapshot is based on.
=======
	// ID: volume ID on which the snapshot is based.
	ID string `json:"id"`
	// Name: volume name on which the snapshot is based on.
>>>>>>> main
	Name string `json:"name"`
}

// Task: task.
type Task struct {
<<<<<<< HEAD
	// ID: the unique ID of the task.
	ID string `json:"id"`
	// Description: the description of the task.
	Description string `json:"description"`
	// Progress: the progress of the task in percent.
	Progress int32 `json:"progress"`
	// StartedAt: the task start date.
	StartedAt *time.Time `json:"started_at"`
	// TerminatedAt: the task end date.
	TerminatedAt *time.Time `json:"terminated_at"`
	// Status: the task status.
=======
	// ID: unique ID of the task.
	ID string `json:"id"`
	// Description: description of the task.
	Description string `json:"description"`
	// Progress: progress of the task in percent.
	Progress int32 `json:"progress"`
	// StartedAt: task start date.
	StartedAt *time.Time `json:"started_at"`
	// TerminatedAt: task end date.
	TerminatedAt *time.Time `json:"terminated_at"`
	// Status: task status.
>>>>>>> main
	// Default value: pending
	Status TaskStatus `json:"status"`

	HrefFrom string `json:"href_from"`

	HrefResult string `json:"href_result"`
<<<<<<< HEAD
	// Zone: the zone in which is the task.
=======
	// Zone: zone in which the task is excecuted.
>>>>>>> main
	Zone scw.Zone `json:"zone"`
}

type UpdateIPResponse struct {
	IP *IP `json:"ip"`
}

type UpdatePlacementGroupResponse struct {
	PlacementGroup *PlacementGroup `json:"placement_group"`
}

<<<<<<< HEAD
type UpdatePlacementGroupServersResponse struct {
=======
// UpdatePlacementGroupServersResponse: update placement group servers response.
type UpdatePlacementGroupServersResponse struct {
	// Servers: instances attached to the placement group.
>>>>>>> main
	Servers []*PlacementGroupServer `json:"servers"`
}

type UpdateServerResponse struct {
	Server *Server `json:"server"`
}

type UpdateVolumeResponse struct {
	Volume *Volume `json:"volume"`
}

// Volume: volume.
type Volume struct {
<<<<<<< HEAD
	// ID: the volume unique ID.
	ID string `json:"id"`
	// Name: the volume name.
	Name string `json:"name"`
	// Deprecated: ExportURI: show the volume NBD export URI.
	ExportURI *string `json:"export_uri"`
	// Size: the volume disk size.
	Size scw.Size `json:"size"`
	// VolumeType: the volume type.
	// Default value: l_ssd
	VolumeType VolumeVolumeType `json:"volume_type"`
	// CreationDate: the volume creation date.
	CreationDate *time.Time `json:"creation_date"`
	// ModificationDate: the volume modification date.
	ModificationDate *time.Time `json:"modification_date"`
	// Organization: the volume organization ID.
	Organization string `json:"organization"`
	// Project: the volume project ID.
	Project string `json:"project"`
	// Tags: the volume tags.
	Tags []string `json:"tags"`
	// Server: the server attached to the volume.
	Server *ServerSummary `json:"server"`
	// State: the volume state.
	// Default value: available
	State VolumeState `json:"state"`
	// Zone: the zone in which is the volume.
=======
	// ID: volume unique ID.
	ID string `json:"id"`
	// Name: volume name.
	Name string `json:"name"`
	// Deprecated: ExportURI: show the volume NBD export URI.
	ExportURI *string `json:"export_uri"`
	// Size: volume disk size.
	Size scw.Size `json:"size"`
	// VolumeType: volume type.
	// Default value: l_ssd
	VolumeType VolumeVolumeType `json:"volume_type"`
	// CreationDate: volume creation date.
	CreationDate *time.Time `json:"creation_date"`
	// ModificationDate: volume modification date.
	ModificationDate *time.Time `json:"modification_date"`
	// Organization: volume Organization ID.
	Organization string `json:"organization"`
	// Project: volume Project ID.
	Project string `json:"project"`
	// Tags: volume tags.
	Tags []string `json:"tags"`
	// Server: instance attached to the volume.
	Server *ServerSummary `json:"server"`
	// State: volume state.
	// Default value: available
	State VolumeState `json:"state"`
	// Zone: zone in which the volume is located.
>>>>>>> main
	Zone scw.Zone `json:"zone"`
}

type VolumeServer struct {
	ID string `json:"id"`

	Name string `json:"name"`

	ExportURI string `json:"export_uri"`

	Organization string `json:"organization"`

	Server *ServerSummary `json:"server"`

	Size scw.Size `json:"size"`
<<<<<<< HEAD
	// VolumeType:
	// Default value: l_ssd
=======
	// VolumeType: default value: l_ssd
>>>>>>> main
	VolumeType VolumeServerVolumeType `json:"volume_type"`

	CreationDate *time.Time `json:"creation_date"`

	ModificationDate *time.Time `json:"modification_date"`
<<<<<<< HEAD
	// State:
	// Default value: available
=======
	// State: default value: available
>>>>>>> main
	State VolumeServerState `json:"state"`

	Project string `json:"project"`

	Boot bool `json:"boot"`

	Zone scw.Zone `json:"zone"`
}

// VolumeServerTemplate: volume server template.
type VolumeServerTemplate struct {
	// ID: UUID of the volume.
<<<<<<< HEAD
	ID string `json:"id,omitempty"`
	// Boot: force the server to boot on this volume.
	// Default value: false
	Boot bool `json:"boot,omitempty"`
	// Name: name of the volume.
	Name string `json:"name,omitempty"`
	// Size: disk size of the volume, must be a multiple of 512.
	Size scw.Size `json:"size,omitempty"`
	// VolumeType: type of the volume.
	// Default value: l_ssd
	VolumeType VolumeVolumeType `json:"volume_type,omitempty"`
	// BaseSnapshot: the ID of the snapshot on which this volume will be based.
	BaseSnapshot string `json:"base_snapshot,omitempty"`
	// Organization: organization ID of the volume.
	Organization string `json:"organization,omitempty"`
	// Project: project ID of the volume.
	Project string `json:"project,omitempty"`
=======
	ID *string `json:"id,omitempty"`
	// Boot: force the Instance to boot on this volume.
	// Default value: false
	Boot *bool `json:"boot,omitempty"`
	// Name: name of the volume.
	Name *string `json:"name,omitempty"`
	// Size: disk size of the volume, must be a multiple of 512.
	Size *scw.Size `json:"size,omitempty"`
	// VolumeType: type of the volume.
	// Default value: l_ssd
	VolumeType VolumeVolumeType `json:"volume_type,omitempty"`
	// BaseSnapshot: ID of the snapshot on which this volume will be based.
	BaseSnapshot *string `json:"base_snapshot,omitempty"`
	// Organization: organization ID of the volume.
	Organization *string `json:"organization,omitempty"`
	// Project: project ID of the volume.
	Project *string `json:"project,omitempty"`
>>>>>>> main
}

type VolumeSummary struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Size scw.Size `json:"size"`
<<<<<<< HEAD
	// VolumeType:
	// Default value: l_ssd
=======
	// VolumeType: default value: l_ssd
>>>>>>> main
	VolumeType VolumeVolumeType `json:"volume_type"`
}

// VolumeTemplate: volume template.
type VolumeTemplate struct {
	// ID: UUID of the volume.
	ID string `json:"id,omitempty"`
	// Name: name of the volume.
	Name string `json:"name,omitempty"`
	// Size: disk size of the volume, must be a multiple of 512.
	Size scw.Size `json:"size,omitempty"`
	// VolumeType: type of the volume.
	// Default value: l_ssd
	VolumeType VolumeVolumeType `json:"volume_type,omitempty"`
	// Deprecated: Organization: organization ID of the volume.
	// Precisely one of Organization, Project must be set.
	Organization *string `json:"organization,omitempty"`
	// Project: project ID of the volume.
	// Precisely one of Organization, Project must be set.
	Project *string `json:"project,omitempty"`
}

type VolumeType struct {
	DisplayName string `json:"display_name"`

	Capabilities *VolumeTypeCapabilities `json:"capabilities"`

	Constraints *VolumeTypeConstraints `json:"constraints"`
}

type VolumeTypeCapabilities struct {
	Snapshot bool `json:"snapshot"`
}

type VolumeTypeConstraints struct {
	Min scw.Size `json:"min"`

	Max scw.Size `json:"max"`
}

// setImageResponse: set image response.
type setImageResponse struct {
	Image *Image `json:"image"`
}

// setSecurityGroupResponse: set security group response.
type setSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security_group"`
}

// setSecurityGroupRuleResponse: set security group rule response.
type setSecurityGroupRuleResponse struct {
	Rule *SecurityGroupRule `json:"rule"`
}

// setServerResponse: set server response.
type setServerResponse struct {
	Server *Server `json:"server"`
}

// setSnapshotResponse: set snapshot response.
type setSnapshotResponse struct {
	Snapshot *Snapshot `json:"snapshot"`
}

// Service API

// Zones list localities the api is available in
func (s *API) Zones() []scw.Zone {
<<<<<<< HEAD
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw1, scw.ZonePlWaw2}
=======
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2}
>>>>>>> main
}

type GetServerTypesAvailabilityRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD

	PerPage *uint32 `json:"-"`

	Page *int32 `json:"-"`
}

// GetServerTypesAvailability: get availability for all server types.
=======
	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	// Default value: 50
	PerPage *uint32 `json:"-"`
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`
}

// GetServerTypesAvailability: get availability.
// Get availability for all Instance types.
>>>>>>> main
func (s *API) GetServerTypesAvailability(req *GetServerTypesAvailabilityRequest, opts ...scw.RequestOption) (*GetServerTypesAvailabilityResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/products/servers/availability",
		Query:   query,
		Headers: http.Header{},
	}

	var resp GetServerTypesAvailabilityResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListServersTypesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PerPage *uint32 `json:"-"`

	Page *int32 `json:"-"`
}

<<<<<<< HEAD
// ListServersTypes: get server types technical details.
=======
// ListServersTypes: list Instance types.
// List available Instance types and their technical details.
>>>>>>> main
func (s *API) ListServersTypes(req *ListServersTypesRequest, opts ...scw.RequestOption) (*ListServersTypesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/products/servers",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListServersTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListVolumesTypesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PerPage *uint32 `json:"-"`

	Page *int32 `json:"-"`
}

<<<<<<< HEAD
// ListVolumesTypes: get volumes technical details.
=======
// ListVolumesTypes: list volume types.
// List all volume types and their technical details.
>>>>>>> main
func (s *API) ListVolumesTypes(req *ListVolumesTypesRequest, opts ...scw.RequestOption) (*ListVolumesTypesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/products/volumes",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListVolumesTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListServersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	// Default value: 50
	PerPage *uint32 `json:"-"`
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`
<<<<<<< HEAD
	// Organization: list only servers of this organization ID.
	Organization *string `json:"-"`
	// Project: list only servers of this project ID.
	Project *string `json:"-"`
	// Name: filter servers by name (for eg. "server1" will return "server100" and "server1" but not "foo").
	Name *string `json:"-"`
	// PrivateIP: list servers by private_ip.
	PrivateIP *net.IP `json:"-"`
	// WithoutIP: list servers that are not attached to a public IP.
	WithoutIP *bool `json:"-"`
	// CommercialType: list servers of this commercial type.
	CommercialType *string `json:"-"`
	// State: list servers in this state.
	// Default value: running
	State *ServerState `json:"-"`
	// Tags: list servers with these exact tags (to filter with several tags, use commas to separate them).
	Tags []string `json:"-"`
	// PrivateNetwork: list servers in this Private Network.
=======
	// Organization: list only Instances of this Organization ID.
	Organization *string `json:"-"`
	// Project: list only Instances of this Project ID.
	Project *string `json:"-"`
	// Name: filter Instances by name (eg. "server1" will return "server100" and "server1" but not "foo").
	Name *string `json:"-"`
	// PrivateIP: list Instances by private_ip.
	PrivateIP *net.IP `json:"-"`
	// WithoutIP: list Instances that are not attached to a public IP.
	WithoutIP *bool `json:"-"`
	// CommercialType: list Instances of this commercial type.
	CommercialType *string `json:"-"`
	// State: list Instances in this state.
	// Default value: running
	State *ServerState `json:"-"`
	// Tags: list Instances with these exact tags (to filter with several tags, use commas to separate them).
	Tags []string `json:"-"`
	// PrivateNetwork: list Instances in this Private Network.
>>>>>>> main
	PrivateNetwork *string `json:"-"`
	// Order: define the order of the returned servers.
	// Default value: creation_date_desc
	Order ListServersRequestOrder `json:"-"`
<<<<<<< HEAD
}

// ListServers: list all servers.
=======
	// PrivateNetworks: list Instances from the given Private Networks (use commas to separate them).
	PrivateNetworks []string `json:"-"`
}

// ListServers: list all Instances.
// List all Instances in a specified Availability Zone, e.g. `fr-par-1`.
>>>>>>> main
func (s *API) ListServers(req *ListServersRequest, opts ...scw.RequestOption) (*ListServersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "project", req.Project)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "private_ip", req.PrivateIP)
	parameter.AddToQuery(query, "without_ip", req.WithoutIP)
	parameter.AddToQuery(query, "commercial_type", req.CommercialType)
	parameter.AddToQuery(query, "state", req.State)
	if len(req.Tags) != 0 {
		parameter.AddToQuery(query, "tags", strings.Join(req.Tags, ","))
	}
	parameter.AddToQuery(query, "private_network", req.PrivateNetwork)
	parameter.AddToQuery(query, "order", req.Order)
<<<<<<< HEAD
=======
	if len(req.PrivateNetworks) != 0 {
		parameter.AddToQuery(query, "private_networks", strings.Join(req.PrivateNetworks, ","))
	}
>>>>>>> main

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// Name: the server name.
	Name string `json:"name,omitempty"`
	// DynamicIPRequired: define if a dynamic IP is required for the instance.
	DynamicIPRequired *bool `json:"dynamic_ip_required,omitempty"`
	// CommercialType: define the server commercial type (i.e. GP1-S).
	CommercialType string `json:"commercial_type,omitempty"`
	// Image: the server image ID or label.
	Image string `json:"image,omitempty"`
	// Volumes: the volumes attached to the server.
	Volumes map[string]*VolumeServerTemplate `json:"volumes,omitempty"`
	// EnableIPv6: true if IPv6 is enabled on the server.
	EnableIPv6 bool `json:"enable_ipv6,omitempty"`
	// PublicIP: the ID of the reserved IP to attach to the server.
	PublicIP *string `json:"public_ip,omitempty"`
	// BootType: the boot type to use.
	// Default value: local
	BootType *BootType `json:"boot_type,omitempty"`
	// Deprecated: Bootscript: the bootscript ID to use when `boot_type` is set to `bootscript`.
	Bootscript *string `json:"bootscript,omitempty"`
	// Deprecated: Organization: the server organization ID.
	// Precisely one of Organization, Project must be set.
	Organization *string `json:"organization,omitempty"`
	// Project: the server project ID.
	// Precisely one of Organization, Project must be set.
	Project *string `json:"project,omitempty"`
	// Tags: the server tags.
	Tags []string `json:"tags,omitempty"`
	// SecurityGroup: the security group ID.
	SecurityGroup *string `json:"security_group,omitempty"`
	// PlacementGroup: placement group ID if server must be part of a placement group.
	PlacementGroup *string `json:"placement_group,omitempty"`
}

// createServer: the `volumes` key is a dictionary composed of the volume position as key and the volume parameters as value.
// Depending of the volume parameters, you can achieve different behaviours :
//
// Create a volume from a snapshot of an image :
// Optional : `volume_type`, `size`, `boot`.
// If the `size` parameter is not set, the size of the volume will equal the size of the corresponding snapshot of the image.
//
// Attach an existing volume :
// Required : `id`, `name`.
// Optional : `boot`.
//
// Create an empty volume :
// Required : `name`, `volume_type`, `size`.
// Optional : `organization`, `project`, `boot`.
//
// Create a volume from a snapshot :
// Required : `base_snapshot`, `name`, `volume_type`.
// Optional : `organization`, `project`, `boot`.
=======
	// Name: instance name.
	Name string `json:"name,omitempty"`
	// DynamicIPRequired: define if a dynamic IPv4 is required for the Instance.
	DynamicIPRequired *bool `json:"dynamic_ip_required,omitempty"`
	// RoutedIPEnabled: if true, configure the Instance so it uses the new routed IP mode.
	RoutedIPEnabled *bool `json:"routed_ip_enabled,omitempty"`
	// CommercialType: define the Instance commercial type (i.e. GP1-S).
	CommercialType string `json:"commercial_type,omitempty"`
	// Image: instance image ID or label.
	Image string `json:"image,omitempty"`
	// Volumes: volumes attached to the server.
	Volumes map[string]*VolumeServerTemplate `json:"volumes,omitempty"`
	// EnableIPv6: true if IPv6 is enabled on the server.
	EnableIPv6 bool `json:"enable_ipv6,omitempty"`
	// PublicIP: ID of the reserved IP to attach to the Instance.
	PublicIP *string `json:"public_ip,omitempty"`
	// PublicIPs: a list of reserved IP IDs to attach to the Instance.
	PublicIPs []*string `json:"public_ips,omitempty"`
	// BootType: boot type to use.
	// Default value: local
	BootType *BootType `json:"boot_type,omitempty"`
	// Deprecated: Bootscript: bootscript ID to use when `boot_type` is set to `bootscript`.
	Bootscript *string `json:"bootscript,omitempty"`
	// Deprecated: Organization: instance Organization ID.
	// Precisely one of Organization, Project must be set.
	Organization *string `json:"organization,omitempty"`
	// Project: instance Project ID.
	// Precisely one of Organization, Project must be set.
	Project *string `json:"project,omitempty"`
	// Tags: instance tags.
	Tags []string `json:"tags,omitempty"`
	// SecurityGroup: security group ID.
	SecurityGroup *string `json:"security_group,omitempty"`
	// PlacementGroup: placement group ID if Instance must be part of a placement group.
	PlacementGroup *string `json:"placement_group,omitempty"`
}

// createServer: create an Instance.
// Create a new Instance of the specified commercial type in the specified zone. Pay attention to the volumes parameter, which takes an object which can be used in different ways to achieve different behaviors.
// Get more information in the [Technical Information](#technical-information) section of the introduction.
>>>>>>> main
func (s *API) createServer(req *CreateServerRequest, opts ...scw.RequestOption) (*CreateServerResponse, error) {
	var err error

	defaultProject, exist := s.client.GetDefaultProjectID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Project = &defaultProject
	}

	defaultOrganization, exist := s.client.GetDefaultOrganizationID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Organization = &defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("srv")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateServerResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`
}

<<<<<<< HEAD
// DeleteServer: delete a server with the given ID.
=======
// DeleteServer: delete an Instance.
// Delete the Instance with the specified ID.
>>>>>>> main
func (s *API) DeleteServer(req *DeleteServerRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// ServerID: UUID of the server you want to get.
	ServerID string `json:"-"`
}

// GetServer: get the details of a specified Server.
=======
	// ServerID: UUID of the Instance you want to get.
	ServerID string `json:"-"`
}

// GetServer: get an Instance.
// Get the details of a specified Instance.
>>>>>>> main
func (s *API) GetServer(req *GetServerRequest, opts ...scw.RequestOption) (*GetServerResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
		Headers: http.Header{},
	}

	var resp GetServerResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type setServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// ID: the server unique ID.
	ID string `json:"-"`
	// Name: the server name.
	Name string `json:"name"`
	// Organization: the server organization ID.
	Organization string `json:"organization"`
	// Project: the server project ID.
	Project string `json:"project"`
	// AllowedActions: provide as list of allowed actions on the server.
	AllowedActions []ServerAction `json:"allowed_actions"`
	// Tags: the server associated tags.
	Tags *[]string `json:"tags"`
	// CommercialType: the server commercial type (eg. GP1-M).
	CommercialType string `json:"commercial_type"`
	// CreationDate: the server creation date.
	CreationDate *time.Time `json:"creation_date"`
	// DynamicIPRequired: true if a dynamic IP is required.
	DynamicIPRequired bool `json:"dynamic_ip_required"`
	// EnableIPv6: true if IPv6 is enabled.
	EnableIPv6 bool `json:"enable_ipv6"`
	// Hostname: the server host name.
	Hostname string `json:"hostname"`
	// Image: provide information on the server image.
	Image *Image `json:"image"`
	// Protected: the server protection option is activated.
	Protected bool `json:"protected"`
	// PrivateIP: the server private IP address.
	PrivateIP *string `json:"private_ip"`
	// PublicIP: information about the public IP.
	PublicIP *ServerIP `json:"public_ip"`
	// ModificationDate: the server modification date.
	ModificationDate *time.Time `json:"modification_date"`
	// State: the server state.
	// Default value: running
	State ServerState `json:"state"`
	// Location: the server location.
	Location *ServerLocation `json:"location"`
	// IPv6: the server IPv6 address.
	IPv6 *ServerIPv6 `json:"ipv6"`
	// Deprecated: Bootscript: the server bootscript.
	Bootscript *Bootscript `json:"bootscript"`
	// BootType: the server boot type.
	// Default value: local
	BootType BootType `json:"boot_type"`
	// Volumes: the server volumes.
	Volumes map[string]*Volume `json:"volumes"`
	// SecurityGroup: the server security group.
	SecurityGroup *SecurityGroupSummary `json:"security_group"`
	// Maintenances: the server planned maintenances.
	Maintenances []*ServerMaintenance `json:"maintenances"`
	// StateDetail: the server state_detail.
	StateDetail string `json:"state_detail"`
	// Arch: the server arch.
	// Default value: x86_64
	Arch Arch `json:"arch"`
	// PlacementGroup: the server placement group.
	PlacementGroup *PlacementGroup `json:"placement_group"`
	// PrivateNics: the server private NICs.
=======
	// ID: instance unique ID.
	ID string `json:"-"`
	// Name: instance name.
	Name string `json:"name"`
	// Organization: instance Organization ID.
	Organization string `json:"organization"`
	// Project: instance Project ID.
	Project string `json:"project"`
	// AllowedActions: provide a list of allowed actions on the server.
	AllowedActions []ServerAction `json:"allowed_actions"`
	// Tags: tags associated with the Instance.
	Tags *[]string `json:"tags"`
	// CommercialType: instance commercial type (eg. GP1-M).
	CommercialType string `json:"commercial_type"`
	// CreationDate: instance creation date.
	CreationDate *time.Time `json:"creation_date"`
	// DynamicIPRequired: true if a dynamic IPv4 is required.
	DynamicIPRequired bool `json:"dynamic_ip_required"`
	// RoutedIPEnabled: true to configure the instance so it uses the new routed IP mode (once this is set to True you cannot set it back to False).
	RoutedIPEnabled *bool `json:"routed_ip_enabled"`
	// EnableIPv6: true if IPv6 is enabled.
	EnableIPv6 bool `json:"enable_ipv6"`
	// Hostname: instance host name.
	Hostname string `json:"hostname"`
	// Image: provide information on the Instance image.
	Image *Image `json:"image"`
	// Protected: instance protection option is activated.
	Protected bool `json:"protected"`
	// PrivateIP: instance private IP address.
	PrivateIP *string `json:"private_ip"`
	// PublicIP: information about the public IP.
	PublicIP *ServerIP `json:"public_ip"`
	// PublicIPs: information about all the public IPs attached to the server.
	PublicIPs []*ServerIP `json:"public_ips"`
	// ModificationDate: instance modification date.
	ModificationDate *time.Time `json:"modification_date"`
	// State: instance state.
	// Default value: running
	State ServerState `json:"state"`
	// Location: instance location.
	Location *ServerLocation `json:"location"`
	// IPv6: instance IPv6 address.
	IPv6 *ServerIPv6 `json:"ipv6"`
	// Deprecated: Bootscript: instance bootscript.
	Bootscript *Bootscript `json:"bootscript"`
	// BootType: instance boot type.
	// Default value: local
	BootType BootType `json:"boot_type"`
	// Volumes: instance volumes.
	Volumes map[string]*Volume `json:"volumes"`
	// SecurityGroup: instance security group.
	SecurityGroup *SecurityGroupSummary `json:"security_group"`
	// Maintenances: instance planned maintenances.
	Maintenances []*ServerMaintenance `json:"maintenances"`
	// StateDetail: instance state_detail.
	StateDetail string `json:"state_detail"`
	// Arch: instance architecture (refers to the CPU architecture used for the Instance, e.g. x86_64, arm64).
	// Default value: x86_64
	Arch Arch `json:"arch"`
	// PlacementGroup: instance placement group.
	PlacementGroup *PlacementGroup `json:"placement_group"`
	// PrivateNics: instance private NICs.
>>>>>>> main
	PrivateNics []*PrivateNIC `json:"private_nics"`
}

func (s *API) setServer(req *setServerRequest, opts ...scw.RequestOption) (*setServerResponse, error) {
	var err error

	if req.Project == "" {
		defaultProject, _ := s.client.GetDefaultProjectID()
		req.Project = defaultProject
	}

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setServerResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// ServerID: UUID of the server.
	ServerID string `json:"-"`
	// Name: name of the server.
	Name *string `json:"name,omitempty"`
	// BootType:
	// Default value: local
	BootType *BootType `json:"boot_type,omitempty"`
	// Tags: tags of the server.
=======
	// ServerID: UUID of the Instance.
	ServerID string `json:"-"`
	// Name: name of the Instance.
	Name *string `json:"name,omitempty"`
	// BootType: default value: local
	BootType *BootType `json:"boot_type,omitempty"`
	// Tags: tags of the Instance.
>>>>>>> main
	Tags *[]string `json:"tags,omitempty"`

	Volumes *map[string]*VolumeServerTemplate `json:"volumes,omitempty"`
	// Deprecated
	Bootscript *string `json:"bootscript,omitempty"`

	DynamicIPRequired *bool `json:"dynamic_ip_required,omitempty"`
<<<<<<< HEAD
=======
	// RoutedIPEnabled: true to configure the instance so it uses the new routed IP mode (once this is set to True you cannot set it back to False).
	RoutedIPEnabled *bool `json:"routed_ip_enabled,omitempty"`

	PublicIPs []*ServerIP `json:"public_ips,omitempty"`
>>>>>>> main

	EnableIPv6 *bool `json:"enable_ipv6,omitempty"`

	Protected *bool `json:"protected,omitempty"`

	SecurityGroup *SecurityGroupTemplate `json:"security_group,omitempty"`
<<<<<<< HEAD
	// PlacementGroup: placement group ID if server must be part of a placement group.
	PlacementGroup *NullableStringValue `json:"placement_group,omitempty"`
	// PrivateNics: the server private NICs.
	PrivateNics []*PrivateNIC `json:"private_nics,omitempty"`
}

// updateServer: update a server.
=======
	// PlacementGroup: placement group ID if Instance must be part of a placement group.
	PlacementGroup *NullableStringValue `json:"placement_group,omitempty"`
	// PrivateNics: instance private NICs.
	PrivateNics []*PrivateNIC `json:"private_nics,omitempty"`
	// CommercialType: set the commercial_type for this Instance.
	// Warning: This field has some restrictions:
	// - Cannot be changed if the Instance is not in `stopped` state.
	// - Cannot be changed if the Instance is in a placement group.
	// - Local storage requirements of the target commercial_types must be fulfilled (i.e. if an Instance has 80GB of local storage, it can be changed into a GP1-XS, which has a maximum of 150GB, but it cannot be changed into a DEV1-S, which has only 20GB).
	CommercialType *string `json:"commercial_type,omitempty"`
}

// updateServer: update an Instance.
// Update the Instance information, such as name, boot mode, or tags.
>>>>>>> main
func (s *API) updateServer(req *UpdateServerRequest, opts ...scw.RequestOption) (*UpdateServerResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateServerResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListServerActionsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`
}

<<<<<<< HEAD
// ListServerActions: list all actions that can currently be performed on a server.
=======
// ListServerActions: list Instance actions.
// List all actions (e.g. power on, power off, reboot) that can currently be performed on an Instance.
>>>>>>> main
func (s *API) ListServerActions(req *ListServerActionsRequest, opts ...scw.RequestOption) (*ListServerActionsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/action",
		Headers: http.Header{},
	}

	var resp ListServerActionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ServerActionRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// ServerID: UUID of the server.
	ServerID string `json:"-"`
	// Action: the action to perform on the server.
	// Default value: poweron
	Action ServerAction `json:"action"`
	// Name: the name of the backup you want to create.
	// This field should only be specified when performing a backup action.
	Name *string `json:"name,omitempty"`
	// Volumes: for each volume UUID, the snapshot parameters of the volume.
=======
	// ServerID: UUID of the Instance.
	ServerID string `json:"-"`
	// Action: action to perform on the Instance.
	// Default value: poweron
	Action ServerAction `json:"action"`
	// Name: name of the backup you want to create.
	// Name of the backup you want to create.
	// This field should only be specified when performing a backup action.
	Name *string `json:"name,omitempty"`
	// Volumes: for each volume UUID, the snapshot parameters of the volume.
	// For each volume UUID, the snapshot parameters of the volume.
>>>>>>> main
	// This field should only be specified when performing a backup action.
	Volumes map[string]*ServerActionRequestVolumeBackupTemplate `json:"volumes,omitempty"`
}

<<<<<<< HEAD
// ServerAction: perform power related actions on a server. Be wary that when terminating a server, all the attached volumes (local *and* block storage) are deleted. So, if you want to keep your local volumes, you must use the `archive` action instead of `terminate`. And if you want to keep block-storage volumes, **you must** detach it beforehand you issue the `terminate` call.  For more information, read the [Volumes](#volumes-7e8a39) documentation.
=======
// ServerAction: perform action.
// Perform an action on an Instance.
// Available actions are:
// * `poweron`: Start a stopped Instance.
// * `poweroff`: Fully stop the Instance and release the hypervisor slot.
// * `stop_in_place`: Stop the Instance, but keep the slot on the hypervisor.
// * `reboot`: Stop the instance and restart it.
// * `backup`:  Create an image with all the volumes of an Instance.
// * `terminate`: Delete the Instance along with all attached volumes.
// * `enable_routed_ip`: Migrate the Instance to the new network stack.
//
// Keep in mind that terminating an Instance will result in the deletion of all attached volumes, including local and block storage.
// If you want to preserve your local volumes, you should use the `archive` action instead of `terminate`. Similarly, if you want to keep your block storage volumes, you must first detach them before issuing the `terminate` command.
// For more information, read the [Volumes](#path-volumes-list-volumes) documentation.
>>>>>>> main
func (s *API) ServerAction(req *ServerActionRequest, opts ...scw.RequestOption) (*ServerActionResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/action",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ServerActionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListServerUserDataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// ServerID: UUID of the server.
	ServerID string `json:"-"`
}

// ListServerUserData: list all user data keys registered on a given server.
=======
	// ServerID: UUID of the Instance.
	ServerID string `json:"-"`
}

// ListServerUserData: list user data.
// List all user data keys registered on a specified Instance.
>>>>>>> main
func (s *API) ListServerUserData(req *ListServerUserDataRequest, opts ...scw.RequestOption) (*ListServerUserDataResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/user_data",
		Headers: http.Header{},
	}

	var resp ListServerUserDataResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteServerUserDataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// ServerID: UUID of the server.
=======
	// ServerID: UUID of the Instance.
>>>>>>> main
	ServerID string `json:"-"`
	// Key: key of the user data to delete.
	Key string `json:"-"`
}

<<<<<<< HEAD
// DeleteServerUserData: delete the given key from a server user data.
=======
// DeleteServerUserData: delete user data.
// Delete the specified key from an Instance's user data.
>>>>>>> main
func (s *API) DeleteServerUserData(req *DeleteServerUserDataRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	if fmt.Sprint(req.Key) == "" {
		return errors.New("field Key cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/user_data/" + fmt.Sprint(req.Key) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListImagesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Organization *string `json:"-"`

	PerPage *uint32 `json:"-"`

	Page *int32 `json:"-"`

	Name *string `json:"-"`

	Public *bool `json:"-"`

	Arch *string `json:"-"`

	Project *string `json:"-"`

	Tags *string `json:"-"`
}

<<<<<<< HEAD
// ListImages: list all images available in an account.
=======
// ListImages: list Instance images.
// List all existing Instance images.
>>>>>>> main
func (s *API) ListImages(req *ListImagesRequest, opts ...scw.RequestOption) (*ListImagesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "public", req.Public)
	parameter.AddToQuery(query, "arch", req.Arch)
	parameter.AddToQuery(query, "project", req.Project)
	parameter.AddToQuery(query, "tags", req.Tags)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListImagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetImageRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ImageID: UUID of the image you want to get.
	ImageID string `json:"-"`
}

<<<<<<< HEAD
// GetImage: get details of an image with the given ID.
=======
// GetImage: get an Instance image.
// Get details of an image with the specified ID.
>>>>>>> main
func (s *API) GetImage(req *GetImageRequest, opts ...scw.RequestOption) (*GetImageResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images/" + fmt.Sprint(req.ImageID) + "",
		Headers: http.Header{},
	}

	var resp GetImageResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateImageRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// Name: name of the image.
	Name string `json:"name,omitempty"`
	// RootVolume: UUID of the snapshot.
	RootVolume string `json:"root_volume,omitempty"`
	// Arch: architecture of the image.
	// Default value: x86_64
	Arch Arch `json:"arch"`
	// Deprecated: DefaultBootscript: default bootscript of the image.
	DefaultBootscript *string `json:"default_bootscript,omitempty"`
	// ExtraVolumes: additional volumes of the image.
	ExtraVolumes map[string]*VolumeTemplate `json:"extra_volumes,omitempty"`
	// Deprecated: Organization: organization ID of the image.
	// Precisely one of Organization, Project must be set.
	Organization *string `json:"organization,omitempty"`
	// Project: project ID of the image.
	// Precisely one of Organization, Project must be set.
	Project *string `json:"project,omitempty"`
<<<<<<< HEAD
	// Tags: the tags of the image.
=======
	// Tags: tags of the image.
>>>>>>> main
	Tags []string `json:"tags,omitempty"`
	// Public: true to create a public image.
	Public *bool `json:"public,omitempty"`
}

<<<<<<< HEAD
// CreateImage: create an instance image.
=======
// CreateImage: create an Instance image.
// Create an Instance image from the specified snapshot ID.
>>>>>>> main
func (s *API) CreateImage(req *CreateImageRequest, opts ...scw.RequestOption) (*CreateImageResponse, error) {
	var err error

	defaultProject, exist := s.client.GetDefaultProjectID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Project = &defaultProject
	}

	defaultOrganization, exist := s.client.GetDefaultOrganizationID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Organization = &defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("img")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateImageResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetImageRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ID string `json:"-"`

	Name string `json:"name"`
<<<<<<< HEAD
	// Arch:
	// Default value: x86_64
=======
	// Arch: default value: x86_64
>>>>>>> main
	Arch Arch `json:"arch"`

	CreationDate *time.Time `json:"creation_date"`

	ModificationDate *time.Time `json:"modification_date"`
	// Deprecated
	DefaultBootscript *Bootscript `json:"default_bootscript"`

	ExtraVolumes map[string]*Volume `json:"extra_volumes"`

	FromServer string `json:"from_server"`

	Organization string `json:"organization"`

	Public bool `json:"public"`

	RootVolume *VolumeSummary `json:"root_volume"`
<<<<<<< HEAD
	// State:
	// Default value: available
=======
	// State: default value: available
>>>>>>> main
	State ImageState `json:"state"`

	Project string `json:"project"`

	Tags *[]string `json:"tags"`
}

<<<<<<< HEAD
// setImage: replace all image properties with an image message.
=======
// setImage: update image.
// Replace all image properties with an image message.
>>>>>>> main
func (s *API) setImage(req *SetImageRequest, opts ...scw.RequestOption) (*setImageResponse, error) {
	var err error

	if req.Project == "" {
		defaultProject, _ := s.client.GetDefaultProjectID()
		req.Project = defaultProject
	}

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images/" + fmt.Sprint(req.ID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setImageResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteImageRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ImageID: UUID of the image you want to delete.
	ImageID string `json:"-"`
}

<<<<<<< HEAD
// DeleteImage: delete the image with the given ID.
=======
// DeleteImage: delete an Instance image.
// Delete the image with the specified ID.
>>>>>>> main
func (s *API) DeleteImage(req *DeleteImageRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ImageID) == "" {
		return errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images/" + fmt.Sprint(req.ImageID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListSnapshotsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Organization *string `json:"-"`

	PerPage *uint32 `json:"-"`

	Page *int32 `json:"-"`

	Name *string `json:"-"`

	Project *string `json:"-"`

	Tags *string `json:"-"`
}

// ListSnapshots: list snapshots.
<<<<<<< HEAD
=======
// List all snapshots of an Organization in a specified Availability Zone.
>>>>>>> main
func (s *API) ListSnapshots(req *ListSnapshotsRequest, opts ...scw.RequestOption) (*ListSnapshotsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "project", req.Project)
	parameter.AddToQuery(query, "tags", req.Tags)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListSnapshotsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateSnapshotRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// Name: name of the snapshot.
	Name string `json:"name,omitempty"`
	// VolumeID: UUID of the volume.
	VolumeID *string `json:"volume_id,omitempty"`
<<<<<<< HEAD
	// Tags: the tags of the snapshot.
=======
	// Tags: tags of the snapshot.
>>>>>>> main
	Tags []string `json:"tags,omitempty"`
	// Deprecated: Organization: organization ID of the snapshot.
	// Precisely one of Organization, Project must be set.
	Organization *string `json:"organization,omitempty"`
	// Project: project ID of the snapshot.
	// Precisely one of Organization, Project must be set.
	Project *string `json:"project,omitempty"`
<<<<<<< HEAD
	// VolumeType: overrides the volume_type of the snapshot.
=======
	// VolumeType: volume type of the snapshot.
	// Overrides the volume_type of the snapshot.
>>>>>>> main
	// If omitted, the volume type of the original volume will be used.
	// Default value: unknown_volume_type
	VolumeType SnapshotVolumeType `json:"volume_type"`
	// Bucket: bucket name for snapshot imports.
	Bucket *string `json:"bucket,omitempty"`
	// Key: object key for snapshot imports.
	Key *string `json:"key,omitempty"`
	// Size: imported snapshot size, must be a multiple of 512.
	Size *scw.Size `json:"size,omitempty"`
}

<<<<<<< HEAD
// CreateSnapshot: create a snapshot from a given volume or from a QCOW2 file.
=======
// CreateSnapshot: create a snapshot from a specified volume or from a QCOW2 file.
// Create a snapshot from a specified volume or from a QCOW2 file in a specified Availability Zone.
>>>>>>> main
func (s *API) CreateSnapshot(req *CreateSnapshotRequest, opts ...scw.RequestOption) (*CreateSnapshotResponse, error) {
	var err error

	defaultProject, exist := s.client.GetDefaultProjectID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Project = &defaultProject
	}

	defaultOrganization, exist := s.client.GetDefaultOrganizationID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Organization = &defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("snp")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateSnapshotResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetSnapshotRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// SnapshotID: UUID of the snapshot you want to get.
	SnapshotID string `json:"-"`
}

<<<<<<< HEAD
// GetSnapshot: get details of a snapshot with the given ID.
=======
// GetSnapshot: get a snapshot.
// Get details of a snapshot with the specified ID.
>>>>>>> main
func (s *API) GetSnapshot(req *GetSnapshotRequest, opts ...scw.RequestOption) (*GetSnapshotResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
		Headers: http.Header{},
	}

	var resp GetSnapshotResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type setSnapshotRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SnapshotID string `json:"-"`

	ID string `json:"id"`

	Name string `json:"name"`

	Organization string `json:"organization"`
<<<<<<< HEAD
	// VolumeType:
	// Default value: l_ssd
	VolumeType VolumeVolumeType `json:"volume_type"`

	Size scw.Size `json:"size"`
	// State:
	// Default value: available
=======
	// VolumeType: default value: l_ssd
	VolumeType VolumeVolumeType `json:"volume_type"`

	Size scw.Size `json:"size"`
	// State: default value: available
>>>>>>> main
	State SnapshotState `json:"state"`

	BaseVolume *SnapshotBaseVolume `json:"base_volume"`

	CreationDate *time.Time `json:"creation_date"`

	ModificationDate *time.Time `json:"modification_date"`

	Project string `json:"project"`

	Tags *[]string `json:"tags"`
}

<<<<<<< HEAD
// setSnapshot: replace all snapshot properties with a snapshot message.
=======
// setSnapshot: update snapshot.
// Replace all snapshot properties with a snapshot message.
>>>>>>> main
func (s *API) setSnapshot(req *setSnapshotRequest, opts ...scw.RequestOption) (*setSnapshotResponse, error) {
	var err error

	if req.Project == "" {
		defaultProject, _ := s.client.GetDefaultProjectID()
		req.Project = defaultProject
	}

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setSnapshotResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteSnapshotRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// SnapshotID: UUID of the snapshot you want to delete.
	SnapshotID string `json:"-"`
}

<<<<<<< HEAD
// DeleteSnapshot: delete the snapshot with the given ID.
=======
// DeleteSnapshot: delete a snapshot.
// Delete the snapshot with the specified ID.
>>>>>>> main
func (s *API) DeleteSnapshot(req *DeleteSnapshotRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ExportSnapshotRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// SnapshotID: the snapshot ID.
=======
	// SnapshotID: snapshot ID.
>>>>>>> main
	SnapshotID string `json:"-"`
	// Bucket: s3 bucket name.
	Bucket string `json:"bucket,omitempty"`
	// Key: s3 object key.
	Key string `json:"key,omitempty"`
}

<<<<<<< HEAD
// ExportSnapshot: export a snapshot to a given S3 bucket in the same region.
=======
// ExportSnapshot: export a snapshot.
// Export a snapshot to a specified S3 bucket in the same region.
>>>>>>> main
func (s *API) ExportSnapshot(req *ExportSnapshotRequest, opts ...scw.RequestOption) (*ExportSnapshotResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "/export",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ExportSnapshotResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListVolumesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// VolumeType: filter by volume type.
	// Default value: l_ssd
	VolumeType *VolumeVolumeType `json:"-"`
	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	// Default value: 50
	PerPage *uint32 `json:"-"`
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`
<<<<<<< HEAD
	// Organization: filter volume by organization ID.
	Organization *string `json:"-"`
	// Project: filter volume by project ID.
=======
	// Organization: filter volume by Organization ID.
	Organization *string `json:"-"`
	// Project: filter volume by Project ID.
>>>>>>> main
	Project *string `json:"-"`
	// Tags: filter volumes with these exact tags (to filter with several tags, use commas to separate them).
	Tags []string `json:"-"`
	// Name: filter volume by name (for eg. "vol" will return "myvolume" but not "data").
	Name *string `json:"-"`
}

// ListVolumes: list volumes.
<<<<<<< HEAD
=======
// List volumes in the specified Availability Zone. You can filter the output by volume type.
>>>>>>> main
func (s *API) ListVolumes(req *ListVolumesRequest, opts ...scw.RequestOption) (*ListVolumesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "volume_type", req.VolumeType)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "project", req.Project)
	if len(req.Tags) != 0 {
		parameter.AddToQuery(query, "tags", strings.Join(req.Tags, ","))
	}
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListVolumesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateVolumeRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// Name: the volume name.
	Name string `json:"name,omitempty"`
	// Deprecated: Organization: the volume organization ID.
	// Precisely one of Organization, Project must be set.
	Organization *string `json:"organization,omitempty"`
	// Project: the volume project ID.
	// Precisely one of Organization, Project must be set.
	Project *string `json:"project,omitempty"`
	// Tags: the volume tags.
	Tags []string `json:"tags,omitempty"`
	// VolumeType: the volume type.
	// Default value: l_ssd
	VolumeType VolumeVolumeType `json:"volume_type"`
	// Size: the volume disk size, must be a multiple of 512.
	// Precisely one of BaseSnapshot, BaseVolume, Size must be set.
	Size *scw.Size `json:"size,omitempty"`
	// BaseVolume: the ID of the volume on which this volume will be based.
	// Precisely one of BaseSnapshot, BaseVolume, Size must be set.
	BaseVolume *string `json:"base_volume,omitempty"`
	// BaseSnapshot: the ID of the snapshot on which this volume will be based.
=======
	// Name: volume name.
	Name string `json:"name,omitempty"`
	// Deprecated: Organization: volume Organization ID.
	// Precisely one of Organization, Project must be set.
	Organization *string `json:"organization,omitempty"`
	// Project: volume Project ID.
	// Precisely one of Organization, Project must be set.
	Project *string `json:"project,omitempty"`
	// Tags: volume tags.
	Tags []string `json:"tags,omitempty"`
	// VolumeType: volume type.
	// Default value: l_ssd
	VolumeType VolumeVolumeType `json:"volume_type"`
	// Size: volume disk size, must be a multiple of 512.
	// Precisely one of BaseSnapshot, BaseVolume, Size must be set.
	Size *scw.Size `json:"size,omitempty"`
	// BaseVolume: ID of the volume on which this volume will be based.
	// Precisely one of BaseSnapshot, BaseVolume, Size must be set.
	BaseVolume *string `json:"base_volume,omitempty"`
	// BaseSnapshot: ID of the snapshot on which this volume will be based.
>>>>>>> main
	// Precisely one of BaseSnapshot, BaseVolume, Size must be set.
	BaseSnapshot *string `json:"base_snapshot,omitempty"`
}

// CreateVolume: create a volume.
<<<<<<< HEAD
=======
// Create a volume of a specified type in an Availability Zone.
>>>>>>> main
func (s *API) CreateVolume(req *CreateVolumeRequest, opts ...scw.RequestOption) (*CreateVolumeResponse, error) {
	var err error

	defaultProject, exist := s.client.GetDefaultProjectID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Project = &defaultProject
	}

	defaultOrganization, exist := s.client.GetDefaultOrganizationID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Organization = &defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("vol")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateVolumeResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetVolumeRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// VolumeID: UUID of the volume you want to get.
	VolumeID string `json:"-"`
}

<<<<<<< HEAD
// GetVolume: get details of a volume with the given ID.
=======
// GetVolume: get a volume.
// Get details of a volume with the specified ID.
>>>>>>> main
func (s *API) GetVolume(req *GetVolumeRequest, opts ...scw.RequestOption) (*GetVolumeResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return nil, errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
		Headers: http.Header{},
	}

	var resp GetVolumeResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateVolumeRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// VolumeID: UUID of the volume.
	VolumeID string `json:"-"`
<<<<<<< HEAD
	// Name: the volume name.
	Name *string `json:"name,omitempty"`
	// Tags: the tags of the volume.
	Tags *[]string `json:"tags,omitempty"`
	// Size: the volume disk size, must be a multiple of 512.
	Size *scw.Size `json:"size,omitempty"`
}

// UpdateVolume: replace name and/or size properties of given ID volume with the given value(s). Any volume name can be changed while, for now, only `b_ssd` volume growing is supported.
=======
	// Name: volume name.
	Name *string `json:"name,omitempty"`
	// Tags: tags of the volume.
	Tags *[]string `json:"tags,omitempty"`
	// Size: volume disk size, must be a multiple of 512.
	Size *scw.Size `json:"size,omitempty"`
}

// UpdateVolume: update a volume.
// Replace the name and/or size properties of a volume specified by its ID, with the specified value(s). Any volume name can be changed, however only `b_ssd` volumes can currently be increased in size.
>>>>>>> main
func (s *API) UpdateVolume(req *UpdateVolumeRequest, opts ...scw.RequestOption) (*UpdateVolumeResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return nil, errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateVolumeResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteVolumeRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// VolumeID: UUID of the volume you want to delete.
	VolumeID string `json:"-"`
}

<<<<<<< HEAD
// DeleteVolume: delete the volume with the given ID.
=======
// DeleteVolume: delete a volume.
// Delete the volume with the specified ID.
>>>>>>> main
func (s *API) DeleteVolume(req *DeleteVolumeRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListSecurityGroupsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// Name: name of the security group.
	Name *string `json:"-"`
<<<<<<< HEAD
	// Organization: the security group organization ID.
	Organization *string `json:"-"`
	// Project: the security group project ID.
=======
	// Organization: security group Organization ID.
	Organization *string `json:"-"`
	// Project: security group Project ID.
>>>>>>> main
	Project *string `json:"-"`
	// Tags: list security groups with these exact tags (to filter with several tags, use commas to separate them).
	Tags []string `json:"-"`
	// ProjectDefault: filter security groups with this value for project_default.
	ProjectDefault *bool `json:"-"`
	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	// Default value: 50
	PerPage *uint32 `json:"-"`
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`
}

<<<<<<< HEAD
// ListSecurityGroups: list all security groups available in an account.
=======
// ListSecurityGroups: list security groups.
// List all existing security groups.
>>>>>>> main
func (s *API) ListSecurityGroups(req *ListSecurityGroupsRequest, opts ...scw.RequestOption) (*ListSecurityGroupsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "project", req.Project)
	if len(req.Tags) != 0 {
		parameter.AddToQuery(query, "tags", strings.Join(req.Tags, ","))
	}
	parameter.AddToQuery(query, "project_default", req.ProjectDefault)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListSecurityGroupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateSecurityGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// Name: name of the security group.
	Name string `json:"name,omitempty"`
	// Description: description of the security group.
	Description string `json:"description,omitempty"`
	// Deprecated: Organization: organization ID the security group belongs to.
	// Precisely one of Organization, Project must be set.
	Organization *string `json:"organization,omitempty"`
	// Project: project ID the security group belong to.
	// Precisely one of Organization, Project must be set.
	Project *string `json:"project,omitempty"`
<<<<<<< HEAD
	// Tags: the tags of the security group.
	Tags []string `json:"tags,omitempty"`
	// Deprecated: OrganizationDefault: whether this security group becomes the default security group for new instances.
	// Default value: false
	// Precisely one of OrganizationDefault, ProjectDefault must be set.
	OrganizationDefault *bool `json:"organization_default,omitempty"`
	// ProjectDefault: whether this security group becomes the default security group for new instances.
=======
	// Tags: tags of the security group.
	Tags []string `json:"tags,omitempty"`
	// Deprecated: OrganizationDefault: defines whether this security group becomes the default security group for new Instances.
	// Default value: false
	// Precisely one of OrganizationDefault, ProjectDefault must be set.
	OrganizationDefault *bool `json:"organization_default,omitempty"`
	// ProjectDefault: whether this security group becomes the default security group for new Instances.
>>>>>>> main
	// Default value: false
	// Precisely one of OrganizationDefault, ProjectDefault must be set.
	ProjectDefault *bool `json:"project_default,omitempty"`
	// Stateful: whether the security group is stateful or not.
	// Default value: false
	Stateful bool `json:"stateful,omitempty"`
	// InboundDefaultPolicy: default policy for inbound rules.
	// Default value: accept
	InboundDefaultPolicy SecurityGroupPolicy `json:"inbound_default_policy"`
	// OutboundDefaultPolicy: default policy for outbound rules.
	// Default value: accept
	OutboundDefaultPolicy SecurityGroupPolicy `json:"outbound_default_policy"`
<<<<<<< HEAD
	// EnableDefaultSecurity: true to block SMTP on IPv4 and IPv6.
=======
	// EnableDefaultSecurity: true to block SMTP on IPv4 and IPv6. This feature is read only, please open a support ticket if you need to make it configurable.
>>>>>>> main
	EnableDefaultSecurity *bool `json:"enable_default_security,omitempty"`
}

// CreateSecurityGroup: create a security group.
<<<<<<< HEAD
=======
// Create a security group with a specified name and description.
>>>>>>> main
func (s *API) CreateSecurityGroup(req *CreateSecurityGroupRequest, opts ...scw.RequestOption) (*CreateSecurityGroupResponse, error) {
	var err error

	defaultProject, exist := s.client.GetDefaultProjectID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Project = &defaultProject
	}

	defaultOrganization, exist := s.client.GetDefaultOrganizationID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Organization = &defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("sg")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateSecurityGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetSecurityGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// SecurityGroupID: UUID of the security group you want to get.
	SecurityGroupID string `json:"-"`
}

<<<<<<< HEAD
// GetSecurityGroup: get the details of a Security Group with the given ID.
=======
// GetSecurityGroup: get a security group.
// Get the details of a security group with the specified ID.
>>>>>>> main
func (s *API) GetSecurityGroup(req *GetSecurityGroupRequest, opts ...scw.RequestOption) (*GetSecurityGroupResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "",
		Headers: http.Header{},
	}

	var resp GetSecurityGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteSecurityGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// SecurityGroupID: UUID of the security group you want to delete.
	SecurityGroupID string `json:"-"`
}

// DeleteSecurityGroup: delete a security group.
<<<<<<< HEAD
=======
// Delete a security group with the specified ID.
>>>>>>> main
func (s *API) DeleteSecurityGroup(req *DeleteSecurityGroupRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type setSecurityGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// ID: the ID of the security group (will be ignored).
	ID string `json:"-"`
	// Name: the name of the security group.
	Name string `json:"name"`
	// Tags: the tags of the security group.
	Tags *[]string `json:"tags"`
	// CreationDate: the creation date of the security group (will be ignored).
	CreationDate *time.Time `json:"creation_date"`
	// ModificationDate: the modification date of the security group (will be ignored).
	ModificationDate *time.Time `json:"modification_date"`
	// Description: the description of the security group.
	Description string `json:"description"`
	// EnableDefaultSecurity: true to block SMTP on IPv4 and IPv6.
	EnableDefaultSecurity bool `json:"enable_default_security"`
	// InboundDefaultPolicy: the default inbound policy.
	// Default value: accept
	InboundDefaultPolicy SecurityGroupPolicy `json:"inbound_default_policy"`
	// OutboundDefaultPolicy: the default outbound policy.
	// Default value: accept
	OutboundDefaultPolicy SecurityGroupPolicy `json:"outbound_default_policy"`
	// Organization: the security groups organization ID.
	Organization string `json:"organization"`
	// Project: the security group project ID.
	Project string `json:"project"`
	// Deprecated: OrganizationDefault: please use project_default instead.
	OrganizationDefault *bool `json:"organization_default"`
	// ProjectDefault: true use this security group for future instances created in this project.
	ProjectDefault bool `json:"project_default"`
	// Servers: the servers attached to this security group.
=======
	// ID: ID of the security group (will be ignored).
	ID string `json:"-"`
	// Name: name of the security group.
	Name string `json:"name"`
	// Tags: tags of the security group.
	Tags *[]string `json:"tags"`
	// CreationDate: creation date of the security group (will be ignored).
	CreationDate *time.Time `json:"creation_date"`
	// ModificationDate: modification date of the security group (will be ignored).
	ModificationDate *time.Time `json:"modification_date"`
	// Description: description of the security group.
	Description string `json:"description"`
	// EnableDefaultSecurity: true to block SMTP on IPv4 and IPv6. This feature is read only, please open a support ticket if you need to make it configurable.
	EnableDefaultSecurity bool `json:"enable_default_security"`
	// InboundDefaultPolicy: default inbound policy.
	// Default value: accept
	InboundDefaultPolicy SecurityGroupPolicy `json:"inbound_default_policy"`
	// OutboundDefaultPolicy: default outbound policy.
	// Default value: accept
	OutboundDefaultPolicy SecurityGroupPolicy `json:"outbound_default_policy"`
	// Organization: security groups Organization ID.
	Organization string `json:"organization"`
	// Project: security group Project ID.
	Project string `json:"project"`
	// Deprecated: OrganizationDefault: please use project_default instead.
	OrganizationDefault *bool `json:"organization_default"`
	// ProjectDefault: true use this security group for future Instances created in this project.
	ProjectDefault bool `json:"project_default"`
	// Servers: instances attached to this security group.
>>>>>>> main
	Servers []*ServerSummary `json:"servers"`
	// Stateful: true to set the security group as stateful.
	Stateful bool `json:"stateful"`
}

<<<<<<< HEAD
// setSecurityGroup: replace all security group properties with a security group message.
=======
// setSecurityGroup: update a security group.
// Replace all security group properties with a security group message.
>>>>>>> main
func (s *API) setSecurityGroup(req *setSecurityGroupRequest, opts ...scw.RequestOption) (*setSecurityGroupResponse, error) {
	var err error

	if req.Project == "" {
		defaultProject, _ := s.client.GetDefaultProjectID()
		req.Project = defaultProject
	}

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.ID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setSecurityGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListDefaultSecurityGroupRulesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

<<<<<<< HEAD
// ListDefaultSecurityGroupRules: lists the default rules applied to all the security groups.
=======
// ListDefaultSecurityGroupRules: get default rules.
// Lists the default rules applied to all the security groups.
>>>>>>> main
func (s *API) ListDefaultSecurityGroupRules(req *ListDefaultSecurityGroupRulesRequest, opts ...scw.RequestOption) (*ListSecurityGroupRulesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/default/rules",
		Headers: http.Header{},
	}

	var resp ListSecurityGroupRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListSecurityGroupRulesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// SecurityGroupID: UUID of the security group.
	SecurityGroupID string `json:"-"`
	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	// Default value: 50
	PerPage *uint32 `json:"-"`
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`
}

// ListSecurityGroupRules: list rules.
<<<<<<< HEAD
=======
// List the rules of the a specified security group ID.
>>>>>>> main
func (s *API) ListSecurityGroupRules(req *ListSecurityGroupRulesRequest, opts ...scw.RequestOption) (*ListSecurityGroupRulesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListSecurityGroupRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateSecurityGroupRuleRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// SecurityGroupID: UUID of the security group.
	SecurityGroupID string `json:"-"`
<<<<<<< HEAD
	// Protocol:
	// Default value: TCP
	Protocol SecurityGroupRuleProtocol `json:"protocol"`
	// Direction:
	// Default value: inbound
	Direction SecurityGroupRuleDirection `json:"direction"`
	// Action:
	// Default value: accept
	Action SecurityGroupRuleAction `json:"action"`

	IPRange scw.IPNet `json:"ip_range,omitempty"`
	// DestPortFrom: the beginning of the range of ports to apply this rule to (inclusive).
	DestPortFrom *uint32 `json:"dest_port_from,omitempty"`
	// DestPortTo: the end of the range of ports to apply this rule to (inclusive).
	DestPortTo *uint32 `json:"dest_port_to,omitempty"`
	// Position: the position of this rule in the security group rules list.
=======
	// Protocol: default value: TCP
	Protocol SecurityGroupRuleProtocol `json:"protocol"`
	// Direction: default value: inbound
	Direction SecurityGroupRuleDirection `json:"direction"`
	// Action: default value: accept
	Action SecurityGroupRuleAction `json:"action"`

	IPRange scw.IPNet `json:"ip_range,omitempty"`
	// DestPortFrom: beginning of the range of ports to apply this rule to (inclusive).
	DestPortFrom *uint32 `json:"dest_port_from,omitempty"`
	// DestPortTo: end of the range of ports to apply this rule to (inclusive).
	DestPortTo *uint32 `json:"dest_port_to,omitempty"`
	// Position: position of this rule in the security group rules list.
>>>>>>> main
	Position uint32 `json:"position,omitempty"`
	// Editable: indicates if this rule is editable (will be ignored).
	Editable bool `json:"editable,omitempty"`
}

// CreateSecurityGroupRule: create rule.
<<<<<<< HEAD
=======
// Create a rule in the specified security group ID.
>>>>>>> main
func (s *API) CreateSecurityGroupRule(req *CreateSecurityGroupRuleRequest, opts ...scw.RequestOption) (*CreateSecurityGroupRuleResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateSecurityGroupRuleResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetSecurityGroupRulesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// SecurityGroupID: UUID of the security group to update the rules on.
	SecurityGroupID string `json:"-"`
	// Rules: list of rules to update in the security group.
	Rules []*SetSecurityGroupRulesRequestRule `json:"rules"`
}

<<<<<<< HEAD
// SetSecurityGroupRules: replaces the rules of the security group with the rules provided. This endpoint supports the update of existing rules, creation of new rules and deletion of existing rules when they are not passed in the request.
=======
// SetSecurityGroupRules: update all the rules of a security group.
// Replaces the existing rules of the security group with the rules provided. This endpoint supports the update of existing rules, creation of new rules and deletion of existing rules when they are not passed in the request.
>>>>>>> main
func (s *API) SetSecurityGroupRules(req *SetSecurityGroupRulesRequest, opts ...scw.RequestOption) (*SetSecurityGroupRulesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetSecurityGroupRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteSecurityGroupRuleRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"-"`

	SecurityGroupRuleID string `json:"-"`
}

<<<<<<< HEAD
// DeleteSecurityGroupRule: delete a security group rule with the given ID.
=======
// DeleteSecurityGroupRule: delete rule.
// Delete a security group rule with the specified ID.
>>>>>>> main
func (s *API) DeleteSecurityGroupRule(req *DeleteSecurityGroupRuleRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return errors.New("field SecurityGroupID cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupRuleID) == "" {
		return errors.New("field SecurityGroupRuleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules/" + fmt.Sprint(req.SecurityGroupRuleID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetSecurityGroupRuleRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"-"`

	SecurityGroupRuleID string `json:"-"`
}

<<<<<<< HEAD
// GetSecurityGroupRule: get details of a security group rule with the given ID.
=======
// GetSecurityGroupRule: get rule.
// Get details of a security group rule with the specified ID.
>>>>>>> main
func (s *API) GetSecurityGroupRule(req *GetSecurityGroupRuleRequest, opts ...scw.RequestOption) (*GetSecurityGroupRuleResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupRuleID) == "" {
		return nil, errors.New("field SecurityGroupRuleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules/" + fmt.Sprint(req.SecurityGroupRuleID) + "",
		Headers: http.Header{},
	}

	var resp GetSecurityGroupRuleResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type setSecurityGroupRuleRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"-"`

	SecurityGroupRuleID string `json:"-"`

	ID string `json:"id"`
<<<<<<< HEAD
	// Protocol:
	// Default value: TCP
	Protocol SecurityGroupRuleProtocol `json:"protocol"`
	// Direction:
	// Default value: inbound
	Direction SecurityGroupRuleDirection `json:"direction"`
	// Action:
	// Default value: accept
=======
	// Protocol: default value: TCP
	Protocol SecurityGroupRuleProtocol `json:"protocol"`
	// Direction: default value: inbound
	Direction SecurityGroupRuleDirection `json:"direction"`
	// Action: default value: accept
>>>>>>> main
	Action SecurityGroupRuleAction `json:"action"`

	IPRange scw.IPNet `json:"ip_range"`

	DestPortFrom *uint32 `json:"dest_port_from"`

	DestPortTo *uint32 `json:"dest_port_to"`

	Position uint32 `json:"position"`

	Editable bool `json:"editable"`
}

// setSecurityGroupRule: update security group rule.
<<<<<<< HEAD
=======
// Update the rule of a specified security group ID.
>>>>>>> main
func (s *API) setSecurityGroupRule(req *setSecurityGroupRuleRequest, opts ...scw.RequestOption) (*setSecurityGroupRuleResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupRuleID) == "" {
		return nil, errors.New("field SecurityGroupRuleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules/" + fmt.Sprint(req.SecurityGroupRuleID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setSecurityGroupRuleResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListPlacementGroupsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	// Default value: 50
	PerPage *uint32 `json:"-"`
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`
<<<<<<< HEAD
	// Organization: list only placement groups of this organization ID.
	Organization *string `json:"-"`
	// Project: list only placement groups of this project ID.
=======
	// Organization: list only placement groups of this Organization ID.
	Organization *string `json:"-"`
	// Project: list only placement groups of this Project ID.
>>>>>>> main
	Project *string `json:"-"`
	// Tags: list placement groups with these exact tags (to filter with several tags, use commas to separate them).
	Tags []string `json:"-"`
	// Name: filter placement groups by name (for eg. "cluster1" will return "cluster100" and "cluster1" but not "foo").
	Name *string `json:"-"`
}

<<<<<<< HEAD
// ListPlacementGroups: list all placement groups.
=======
// ListPlacementGroups: list placement groups.
// List all placement groups in a specified Availability Zone.
>>>>>>> main
func (s *API) ListPlacementGroups(req *ListPlacementGroupsRequest, opts ...scw.RequestOption) (*ListPlacementGroupsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "project", req.Project)
	if len(req.Tags) != 0 {
		parameter.AddToQuery(query, "tags", strings.Join(req.Tags, ","))
	}
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListPlacementGroupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreatePlacementGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// Name: name of the placement group.
	Name string `json:"name,omitempty"`
	// Deprecated: Organization: organization ID of the placement group.
	// Precisely one of Organization, Project must be set.
	Organization *string `json:"organization,omitempty"`
	// Project: project ID of the placement group.
	// Precisely one of Organization, Project must be set.
	Project *string `json:"project,omitempty"`
<<<<<<< HEAD
	// Tags: the tags of the placement group.
	Tags []string `json:"tags,omitempty"`
	// PolicyMode: the operating mode of the placement group.
	// Default value: optional
	PolicyMode PlacementGroupPolicyMode `json:"policy_mode"`
	// PolicyType: the policy type of the placement group.
=======
	// Tags: tags of the placement group.
	Tags []string `json:"tags,omitempty"`
	// PolicyMode: operating mode of the placement group.
	// Default value: optional
	PolicyMode PlacementGroupPolicyMode `json:"policy_mode"`
	// PolicyType: policy type of the placement group.
>>>>>>> main
	// Default value: max_availability
	PolicyType PlacementGroupPolicyType `json:"policy_type"`
}

<<<<<<< HEAD
// CreatePlacementGroup: create a new placement group.
=======
// CreatePlacementGroup: create a placement group.
// Create a new placement group in a specified Availability Zone.
>>>>>>> main
func (s *API) CreatePlacementGroup(req *CreatePlacementGroupRequest, opts ...scw.RequestOption) (*CreatePlacementGroupResponse, error) {
	var err error

	defaultProject, exist := s.client.GetDefaultProjectID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Project = &defaultProject
	}

	defaultOrganization, exist := s.client.GetDefaultOrganizationID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Organization = &defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("pg")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreatePlacementGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetPlacementGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// PlacementGroupID: UUID of the placement group you want to get.
	PlacementGroupID string `json:"-"`
}

<<<<<<< HEAD
// GetPlacementGroup: get the given placement group.
=======
// GetPlacementGroup: get a placement group.
// Get the specified placement group.
>>>>>>> main
func (s *API) GetPlacementGroup(req *GetPlacementGroupRequest, opts ...scw.RequestOption) (*GetPlacementGroupResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "",
		Headers: http.Header{},
	}

	var resp GetPlacementGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetPlacementGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PlacementGroupID string `json:"-"`

	Name string `json:"name"`

	Organization string `json:"organization"`
<<<<<<< HEAD
	// PolicyMode:
	// Default value: optional
	PolicyMode PlacementGroupPolicyMode `json:"policy_mode"`
	// PolicyType:
	// Default value: max_availability
=======
	// PolicyMode: default value: optional
	PolicyMode PlacementGroupPolicyMode `json:"policy_mode"`
	// PolicyType: default value: max_availability
>>>>>>> main
	PolicyType PlacementGroupPolicyType `json:"policy_type"`

	Project string `json:"project"`

	Tags *[]string `json:"tags"`
}

<<<<<<< HEAD
// SetPlacementGroup: set all parameters of the given placement group.
=======
// SetPlacementGroup: set placement group.
// Set all parameters of the specified placement group.
>>>>>>> main
func (s *API) SetPlacementGroup(req *SetPlacementGroupRequest, opts ...scw.RequestOption) (*SetPlacementGroupResponse, error) {
	var err error

	if req.Project == "" {
		defaultProject, _ := s.client.GetDefaultProjectID()
		req.Project = defaultProject
	}

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetPlacementGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdatePlacementGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// PlacementGroupID: UUID of the placement group.
	PlacementGroupID string `json:"-"`
	// Name: name of the placement group.
	Name *string `json:"name,omitempty"`
<<<<<<< HEAD
	// Tags: the tags of the placement group.
	Tags *[]string `json:"tags,omitempty"`
	// PolicyMode: the operating mode of the placement group.
	// Default value: optional
	PolicyMode *PlacementGroupPolicyMode `json:"policy_mode,omitempty"`
	// PolicyType: the policy type of the placement group.
=======
	// Tags: tags of the placement group.
	Tags *[]string `json:"tags,omitempty"`
	// PolicyMode: operating mode of the placement group.
	// Default value: optional
	PolicyMode *PlacementGroupPolicyMode `json:"policy_mode,omitempty"`
	// PolicyType: policy type of the placement group.
>>>>>>> main
	// Default value: max_availability
	PolicyType *PlacementGroupPolicyType `json:"policy_type,omitempty"`
}

<<<<<<< HEAD
// UpdatePlacementGroup: update one or more parameter of the given placement group.
=======
// UpdatePlacementGroup: update a placement group.
// Update one or more parameter of the specified placement group.
>>>>>>> main
func (s *API) UpdatePlacementGroup(req *UpdatePlacementGroupRequest, opts ...scw.RequestOption) (*UpdatePlacementGroupResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdatePlacementGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeletePlacementGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// PlacementGroupID: UUID of the placement group you want to delete.
	PlacementGroupID string `json:"-"`
}

<<<<<<< HEAD
// DeletePlacementGroup: delete the given placement group.
=======
// DeletePlacementGroup: delete the specified placement group.
>>>>>>> main
func (s *API) DeletePlacementGroup(req *DeletePlacementGroupRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetPlacementGroupServersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD

	PlacementGroupID string `json:"-"`
}

// GetPlacementGroupServers: get all servers belonging to the given placement group.
=======
	// PlacementGroupID: UUID of the placement group you want to get.
	PlacementGroupID string `json:"-"`
}

// GetPlacementGroupServers: get placement group servers.
// Get all Instances belonging to the specified placement group.
>>>>>>> main
func (s *API) GetPlacementGroupServers(req *GetPlacementGroupServersRequest, opts ...scw.RequestOption) (*GetPlacementGroupServersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "/servers",
		Headers: http.Header{},
	}

	var resp GetPlacementGroupServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetPlacementGroupServersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD

	PlacementGroupID string `json:"-"`

	Servers []string `json:"servers"`
}

// SetPlacementGroupServers: set all servers belonging to the given placement group.
=======
	// PlacementGroupID: UUID of the placement group you want to set.
	PlacementGroupID string `json:"-"`
	// Servers: an array of the Instances' UUIDs you want to configure.
	Servers []string `json:"servers"`
}

// SetPlacementGroupServers: set placement group servers.
// Set all Instances belonging to the specified placement group.
>>>>>>> main
func (s *API) SetPlacementGroupServers(req *SetPlacementGroupServersRequest, opts ...scw.RequestOption) (*SetPlacementGroupServersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "/servers",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetPlacementGroupServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdatePlacementGroupServersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// PlacementGroupID: UUID of the placement group.
	PlacementGroupID string `json:"-"`

	Servers []string `json:"servers,omitempty"`
}

// UpdatePlacementGroupServers: update all servers belonging to the given placement group.
=======
	// PlacementGroupID: UUID of the placement group you want to update.
	PlacementGroupID string `json:"-"`
	// Servers: an array of the Instances' UUIDs you want to configure.
	Servers []string `json:"servers,omitempty"`
}

// UpdatePlacementGroupServers: update placement group servers.
// Update all Instances belonging to the specified placement group.
>>>>>>> main
func (s *API) UpdatePlacementGroupServers(req *UpdatePlacementGroupServersRequest, opts ...scw.RequestOption) (*UpdatePlacementGroupServersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "/servers",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdatePlacementGroupServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListIPsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// Project: the project ID the IPs are reserved in.
	Project *string `json:"-"`
	// Organization: the organization ID the IPs are reserved in.
=======
	// Project: project ID in which the IPs are reserved.
	Project *string `json:"-"`
	// Organization: organization ID in which the IPs are reserved.
>>>>>>> main
	Organization *string `json:"-"`
	// Tags: filter IPs with these exact tags (to filter with several tags, use commas to separate them).
	Tags []string `json:"-"`
	// Name: filter on the IP address (Works as a LIKE operation on the IP address).
	Name *string `json:"-"`
	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	// Default value: 50
	PerPage *uint32 `json:"-"`
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`
}

// ListIPs: list all flexible IPs.
<<<<<<< HEAD
=======
// List all flexible IPs in a specified zone.
>>>>>>> main
func (s *API) ListIPs(req *ListIPsRequest, opts ...scw.RequestOption) (*ListIPsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project", req.Project)
	parameter.AddToQuery(query, "organization", req.Organization)
	if len(req.Tags) != 0 {
		parameter.AddToQuery(query, "tags", strings.Join(req.Tags, ","))
	}
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// Deprecated: Organization: the organization ID the IP is reserved in.
	// Precisely one of Organization, Project must be set.
	Organization *string `json:"organization,omitempty"`
	// Project: the project ID the IP is reserved in.
	// Precisely one of Organization, Project must be set.
	Project *string `json:"project,omitempty"`
	// Tags: the tags of the IP.
	Tags []string `json:"tags,omitempty"`
	// Server: UUID of the server you want to attach the IP to.
	Server *string `json:"server,omitempty"`
}

// CreateIP: reserve a flexible IP.
=======
	// Deprecated: Organization: organization ID in which the IP is reserved.
	// Precisely one of Organization, Project must be set.
	Organization *string `json:"organization,omitempty"`
	// Project: project ID in which the IP is reserved.
	// Precisely one of Organization, Project must be set.
	Project *string `json:"project,omitempty"`
	// Tags: tags of the IP.
	Tags []string `json:"tags,omitempty"`
	// Server: UUID of the Instance you want to attach the IP to.
	Server *string `json:"server,omitempty"`
	// Type: IP type to reserve (either 'nat', 'routed_ipv4' or 'routed_ipv6').
	// Default value: unknown_iptype
	Type IPType `json:"type"`
}

// CreateIP: reserve a flexible IP.
// Reserve a flexible IP and attach it to the specified Instance.
>>>>>>> main
func (s *API) CreateIP(req *CreateIPRequest, opts ...scw.RequestOption) (*CreateIPResponse, error) {
	var err error

	defaultProject, exist := s.client.GetDefaultProjectID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Project = &defaultProject
	}

	defaultOrganization, exist := s.client.GetDefaultOrganizationID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Organization = &defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// IP: the IP ID or address to get.
	IP string `json:"-"`
}

// GetIP: get details of an IP with the given ID or address.
=======
	// IP: IP ID or address to get.
	IP string `json:"-"`
}

// GetIP: get a flexible IP.
// Get details of an IP with the specified ID or address.
>>>>>>> main
func (s *API) GetIP(req *GetIPRequest, opts ...scw.RequestOption) (*GetIPResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IP) == "" {
		return nil, errors.New("field IP cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IP) + "",
		Headers: http.Header{},
	}

	var resp GetIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// IP: IP ID or IP address.
	IP string `json:"-"`
	// Reverse: reverse domain name.
	Reverse *NullableStringValue `json:"reverse,omitempty"`
<<<<<<< HEAD
=======
	// Type: convert a 'nat' IP to a 'routed_ipv4'.
	// Default value: unknown_iptype
	Type IPType `json:"type"`
>>>>>>> main
	// Tags: an array of keywords you want to tag this IP with.
	Tags *[]string `json:"tags,omitempty"`

	Server *NullableStringValue `json:"server,omitempty"`
}

// UpdateIP: update a flexible IP.
<<<<<<< HEAD
=======
// Update a flexible IP in the specified zone with the specified ID.
>>>>>>> main
func (s *API) UpdateIP(req *UpdateIPRequest, opts ...scw.RequestOption) (*UpdateIPResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IP) == "" {
		return nil, errors.New("field IP cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IP) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// IP: the ID or the address of the IP to delete.
	IP string `json:"-"`
}

// DeleteIP: delete the IP with the given ID.
=======
	// IP: ID or address of the IP to delete.
	IP string `json:"-"`
}

// DeleteIP: delete a flexible IP.
// Delete the IP with the specified ID.
>>>>>>> main
func (s *API) DeleteIP(req *DeleteIPRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IP) == "" {
		return errors.New("field IP cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IP) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListPrivateNICsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// ServerID: the server the private NIC is attached to.
	ServerID string `json:"-"`
	// Tags: the private NIC tags.
=======
	// ServerID: instance to which the private NIC is attached.
	ServerID string `json:"-"`
	// Tags: private NIC tags.
>>>>>>> main
	Tags []string `json:"-"`
	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	// Default value: 50
	PerPage *uint32 `json:"-"`
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`
}

<<<<<<< HEAD
// ListPrivateNICs: list all private NICs of a given server.
=======
// ListPrivateNICs: list all private NICs.
// List all private NICs of a specified Instance.
>>>>>>> main
func (s *API) ListPrivateNICs(req *ListPrivateNICsRequest, opts ...scw.RequestOption) (*ListPrivateNICsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	if len(req.Tags) != 0 {
		parameter.AddToQuery(query, "tags", strings.Join(req.Tags, ","))
	}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private_nics",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListPrivateNICsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreatePrivateNICRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// ServerID: UUID of the server the private NIC will be attached to.
	ServerID string `json:"-"`
	// PrivateNetworkID: UUID of the private network where the private NIC will be attached.
	PrivateNetworkID string `json:"private_network_id,omitempty"`
	// Tags: the private NIC tags.
	Tags []string `json:"tags,omitempty"`
}

// CreatePrivateNIC: create a private NIC connecting a server to a private network.
=======
	// ServerID: UUID of the Instance the private NIC will be attached to.
	ServerID string `json:"-"`
	// PrivateNetworkID: UUID of the private network where the private NIC will be attached.
	PrivateNetworkID string `json:"private_network_id,omitempty"`
	// Tags: private NIC tags.
	Tags []string `json:"tags,omitempty"`
	// IPIDs: ip_ids defined from IPAM.
	IPIDs []string `json:"ip_ids,omitempty"`
}

// CreatePrivateNIC: create a private NIC connecting an Instance to a Private Network.
>>>>>>> main
func (s *API) CreatePrivateNIC(req *CreatePrivateNICRequest, opts ...scw.RequestOption) (*CreatePrivateNICResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private_nics",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreatePrivateNICResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetPrivateNICRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// ServerID: the server the private NIC is attached to.
	ServerID string `json:"-"`
	// PrivateNicID: the private NIC unique ID.
	PrivateNicID string `json:"-"`
}

// GetPrivateNIC: get private NIC properties.
=======
	// ServerID: instance to which the private NIC is attached.
	ServerID string `json:"-"`
	// PrivateNicID: private NIC unique ID.
	PrivateNicID string `json:"-"`
}

// GetPrivateNIC: get a private NIC.
// Get private NIC properties.
>>>>>>> main
func (s *API) GetPrivateNIC(req *GetPrivateNICRequest, opts ...scw.RequestOption) (*GetPrivateNICResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNicID) == "" {
		return nil, errors.New("field PrivateNicID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private_nics/" + fmt.Sprint(req.PrivateNicID) + "",
		Headers: http.Header{},
	}

	var resp GetPrivateNICResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdatePrivateNICRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// ServerID: UUID of the server the private NIC will be attached to.
	ServerID string `json:"-"`
	// PrivateNicID: the private NIC unique ID.
=======
	// ServerID: UUID of the Instance the private NIC will be attached to.
	ServerID string `json:"-"`
	// PrivateNicID: private NIC unique ID.
>>>>>>> main
	PrivateNicID string `json:"-"`
	// Tags: tags used to select private NIC/s.
	Tags *[]string `json:"tags,omitempty"`
}

<<<<<<< HEAD
// UpdatePrivateNIC: update one or more parameter/s to a given private NIC.
=======
// UpdatePrivateNIC: update a private NIC.
// Update one or more parameter(s) of a specified private NIC.
>>>>>>> main
func (s *API) UpdatePrivateNIC(req *UpdatePrivateNICRequest, opts ...scw.RequestOption) (*PrivateNIC, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNicID) == "" {
		return nil, errors.New("field PrivateNicID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private_nics/" + fmt.Sprint(req.PrivateNicID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrivateNIC

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeletePrivateNICRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
<<<<<<< HEAD
	// ServerID: the server the private NIC is attached to.
	ServerID string `json:"-"`
	// PrivateNicID: the private NIC unique ID.
=======
	// ServerID: instance to which the private NIC is attached.
	ServerID string `json:"-"`
	// PrivateNicID: private NIC unique ID.
>>>>>>> main
	PrivateNicID string `json:"-"`
}

// DeletePrivateNIC: delete a private NIC.
func (s *API) DeletePrivateNIC(req *DeletePrivateNICRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNicID) == "" {
		return errors.New("field PrivateNicID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private_nics/" + fmt.Sprint(req.PrivateNicID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListBootscriptsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Arch *string `json:"-"`

	Title *string `json:"-"`

	Default *bool `json:"-"`

	Public *bool `json:"-"`

	PerPage *uint32 `json:"-"`

	Page *int32 `json:"-"`
}

// Deprecated: ListBootscripts: list bootscripts.
func (s *API) ListBootscripts(req *ListBootscriptsRequest, opts ...scw.RequestOption) (*ListBootscriptsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "arch", req.Arch)
	parameter.AddToQuery(query, "title", req.Title)
	parameter.AddToQuery(query, "default", req.Default)
	parameter.AddToQuery(query, "public", req.Public)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/bootscripts",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListBootscriptsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetBootscriptRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	BootscriptID string `json:"-"`
}

<<<<<<< HEAD
// Deprecated: GetBootscript: get details of a bootscript with the given ID.
=======
// Deprecated: GetBootscript: get bootscripts.
// Get details of a bootscript with the specified ID.
>>>>>>> main
func (s *API) GetBootscript(req *GetBootscriptRequest, opts ...scw.RequestOption) (*GetBootscriptResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.BootscriptID) == "" {
		return nil, errors.New("field BootscriptID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/bootscripts/" + fmt.Sprint(req.BootscriptID) + "",
		Headers: http.Header{},
	}

	var resp GetBootscriptResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetDashboardRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Organization *string `json:"-"`

	Project *string `json:"-"`
}

func (s *API) GetDashboard(req *GetDashboardRequest, opts ...scw.RequestOption) (*GetDashboardResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "project", req.Project)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/dashboard",
		Query:   query,
		Headers: http.Header{},
	}

	var resp GetDashboardResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListServersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint32(len(results.Servers))
	return uint32(len(results.Servers)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListImagesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListImagesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListImagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Images = append(r.Images, results.Images...)
	r.TotalCount += uint32(len(results.Images))
	return uint32(len(results.Images)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSnapshotsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSnapshotsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSnapshotsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Snapshots = append(r.Snapshots, results.Snapshots...)
	r.TotalCount += uint32(len(results.Snapshots))
	return uint32(len(results.Snapshots)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVolumesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVolumesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListVolumesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Volumes = append(r.Volumes, results.Volumes...)
	r.TotalCount += uint32(len(results.Volumes))
	return uint32(len(results.Volumes)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecurityGroupsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecurityGroupsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSecurityGroupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SecurityGroups = append(r.SecurityGroups, results.SecurityGroups...)
	r.TotalCount += uint32(len(results.SecurityGroups))
	return uint32(len(results.SecurityGroups)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecurityGroupRulesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecurityGroupRulesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSecurityGroupRulesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Rules = append(r.Rules, results.Rules...)
	r.TotalCount += uint32(len(results.Rules))
	return uint32(len(results.Rules)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPlacementGroupsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPlacementGroupsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPlacementGroupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PlacementGroups = append(r.PlacementGroups, results.PlacementGroups...)
	r.TotalCount += uint32(len(results.PlacementGroups))
	return uint32(len(results.PlacementGroups)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.IPs = append(r.IPs, results.IPs...)
	r.TotalCount += uint32(len(results.IPs))
	return uint32(len(results.IPs)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPrivateNICsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPrivateNICsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPrivateNICsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PrivateNics = append(r.PrivateNics, results.PrivateNics...)
	r.TotalCount += uint64(len(results.PrivateNics))
	return uint64(len(results.PrivateNics)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListBootscriptsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListBootscriptsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListBootscriptsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Bootscripts = append(r.Bootscripts, results.Bootscripts...)
	r.TotalCount += uint32(len(results.Bootscripts))
	return uint32(len(results.Bootscripts)), nil
}
