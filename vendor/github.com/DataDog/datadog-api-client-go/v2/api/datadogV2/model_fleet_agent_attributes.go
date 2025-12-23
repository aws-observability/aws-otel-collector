// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAgentAttributes Attributes of a Datadog Agent in the list view.
type FleetAgentAttributes struct {
	// The Datadog Agent version.
	AgentVersion *string `json:"agent_version,omitempty"`
	// The API key name (if available and not redacted).
	ApiKeyName *string `json:"api_key_name,omitempty"`
	// The API key UUID.
	ApiKeyUuid *string `json:"api_key_uuid,omitempty"`
	// The cloud provider where the agent is running.
	CloudProvider *string `json:"cloud_provider,omitempty"`
	// Kubernetes cluster name (if applicable).
	ClusterName *string `json:"cluster_name,omitempty"`
	// The unique agent key identifier.
	DatadogAgentKey *string `json:"datadog_agent_key,omitempty"`
	// Datadog products enabled on the agent.
	EnabledProducts []string `json:"enabled_products,omitempty"`
	// Environments the agent is reporting from.
	Envs []string `json:"envs,omitempty"`
	// Timestamp when the agent was first seen.
	FirstSeenAt *int64 `json:"first_seen_at,omitempty"`
	// The hostname of the agent.
	Hostname *string `json:"hostname,omitempty"`
	// IP addresses of the agent.
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Whether single-step instrumentation is enabled.
	IsSingleStepInstrumentationEnabled *bool `json:"is_single_step_instrumentation_enabled,omitempty"`
	// Timestamp of the last agent restart.
	LastRestartAt *int64 `json:"last_restart_at,omitempty"`
	// The operating system.
	Os *string `json:"os,omitempty"`
	// OpenTelemetry collector version (if applicable).
	OtelCollectorVersion *string `json:"otel_collector_version,omitempty"`
	// List of OpenTelemetry collector versions (if applicable).
	OtelCollectorVersions []string `json:"otel_collector_versions,omitempty"`
	// Kubernetes pod name (if applicable).
	PodName *string `json:"pod_name,omitempty"`
	// Remote agent management status.
	RemoteAgentManagement *string `json:"remote_agent_management,omitempty"`
	// Remote configuration status.
	RemoteConfigStatus *string `json:"remote_config_status,omitempty"`
	// Services running on the agent.
	Services []string `json:"services,omitempty"`
	// Tags associated with the agent.
	Tags []FleetAgentAttributesTagsItems `json:"tags,omitempty"`
	// Team associated with the agent.
	Team *string `json:"team,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetAgentAttributes instantiates a new FleetAgentAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetAgentAttributes() *FleetAgentAttributes {
	this := FleetAgentAttributes{}
	return &this
}

// NewFleetAgentAttributesWithDefaults instantiates a new FleetAgentAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetAgentAttributesWithDefaults() *FleetAgentAttributes {
	this := FleetAgentAttributes{}
	return &this
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetAgentVersion() string {
	if o == nil || o.AgentVersion == nil {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetAgentVersionOk() (*string, bool) {
	if o == nil || o.AgentVersion == nil {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasAgentVersion() bool {
	return o != nil && o.AgentVersion != nil
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *FleetAgentAttributes) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetApiKeyName returns the ApiKeyName field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetApiKeyName() string {
	if o == nil || o.ApiKeyName == nil {
		var ret string
		return ret
	}
	return *o.ApiKeyName
}

// GetApiKeyNameOk returns a tuple with the ApiKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetApiKeyNameOk() (*string, bool) {
	if o == nil || o.ApiKeyName == nil {
		return nil, false
	}
	return o.ApiKeyName, true
}

// HasApiKeyName returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasApiKeyName() bool {
	return o != nil && o.ApiKeyName != nil
}

// SetApiKeyName gets a reference to the given string and assigns it to the ApiKeyName field.
func (o *FleetAgentAttributes) SetApiKeyName(v string) {
	o.ApiKeyName = &v
}

// GetApiKeyUuid returns the ApiKeyUuid field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetApiKeyUuid() string {
	if o == nil || o.ApiKeyUuid == nil {
		var ret string
		return ret
	}
	return *o.ApiKeyUuid
}

// GetApiKeyUuidOk returns a tuple with the ApiKeyUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetApiKeyUuidOk() (*string, bool) {
	if o == nil || o.ApiKeyUuid == nil {
		return nil, false
	}
	return o.ApiKeyUuid, true
}

// HasApiKeyUuid returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasApiKeyUuid() bool {
	return o != nil && o.ApiKeyUuid != nil
}

// SetApiKeyUuid gets a reference to the given string and assigns it to the ApiKeyUuid field.
func (o *FleetAgentAttributes) SetApiKeyUuid(v string) {
	o.ApiKeyUuid = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasCloudProvider() bool {
	return o != nil && o.CloudProvider != nil
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *FleetAgentAttributes) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *FleetAgentAttributes) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetDatadogAgentKey returns the DatadogAgentKey field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetDatadogAgentKey() string {
	if o == nil || o.DatadogAgentKey == nil {
		var ret string
		return ret
	}
	return *o.DatadogAgentKey
}

// GetDatadogAgentKeyOk returns a tuple with the DatadogAgentKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetDatadogAgentKeyOk() (*string, bool) {
	if o == nil || o.DatadogAgentKey == nil {
		return nil, false
	}
	return o.DatadogAgentKey, true
}

// HasDatadogAgentKey returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasDatadogAgentKey() bool {
	return o != nil && o.DatadogAgentKey != nil
}

// SetDatadogAgentKey gets a reference to the given string and assigns it to the DatadogAgentKey field.
func (o *FleetAgentAttributes) SetDatadogAgentKey(v string) {
	o.DatadogAgentKey = &v
}

// GetEnabledProducts returns the EnabledProducts field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetEnabledProducts() []string {
	if o == nil || o.EnabledProducts == nil {
		var ret []string
		return ret
	}
	return o.EnabledProducts
}

// GetEnabledProductsOk returns a tuple with the EnabledProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetEnabledProductsOk() (*[]string, bool) {
	if o == nil || o.EnabledProducts == nil {
		return nil, false
	}
	return &o.EnabledProducts, true
}

// HasEnabledProducts returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasEnabledProducts() bool {
	return o != nil && o.EnabledProducts != nil
}

// SetEnabledProducts gets a reference to the given []string and assigns it to the EnabledProducts field.
func (o *FleetAgentAttributes) SetEnabledProducts(v []string) {
	o.EnabledProducts = v
}

// GetEnvs returns the Envs field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetEnvs() []string {
	if o == nil || o.Envs == nil {
		var ret []string
		return ret
	}
	return o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetEnvsOk() (*[]string, bool) {
	if o == nil || o.Envs == nil {
		return nil, false
	}
	return &o.Envs, true
}

// HasEnvs returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasEnvs() bool {
	return o != nil && o.Envs != nil
}

// SetEnvs gets a reference to the given []string and assigns it to the Envs field.
func (o *FleetAgentAttributes) SetEnvs(v []string) {
	o.Envs = v
}

// GetFirstSeenAt returns the FirstSeenAt field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetFirstSeenAt() int64 {
	if o == nil || o.FirstSeenAt == nil {
		var ret int64
		return ret
	}
	return *o.FirstSeenAt
}

// GetFirstSeenAtOk returns a tuple with the FirstSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetFirstSeenAtOk() (*int64, bool) {
	if o == nil || o.FirstSeenAt == nil {
		return nil, false
	}
	return o.FirstSeenAt, true
}

// HasFirstSeenAt returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasFirstSeenAt() bool {
	return o != nil && o.FirstSeenAt != nil
}

// SetFirstSeenAt gets a reference to the given int64 and assigns it to the FirstSeenAt field.
func (o *FleetAgentAttributes) SetFirstSeenAt(v int64) {
	o.FirstSeenAt = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *FleetAgentAttributes) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetIpAddresses() []string {
	if o == nil || o.IpAddresses == nil {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return &o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasIpAddresses() bool {
	return o != nil && o.IpAddresses != nil
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *FleetAgentAttributes) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetIsSingleStepInstrumentationEnabled returns the IsSingleStepInstrumentationEnabled field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetIsSingleStepInstrumentationEnabled() bool {
	if o == nil || o.IsSingleStepInstrumentationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsSingleStepInstrumentationEnabled
}

// GetIsSingleStepInstrumentationEnabledOk returns a tuple with the IsSingleStepInstrumentationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetIsSingleStepInstrumentationEnabledOk() (*bool, bool) {
	if o == nil || o.IsSingleStepInstrumentationEnabled == nil {
		return nil, false
	}
	return o.IsSingleStepInstrumentationEnabled, true
}

// HasIsSingleStepInstrumentationEnabled returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasIsSingleStepInstrumentationEnabled() bool {
	return o != nil && o.IsSingleStepInstrumentationEnabled != nil
}

// SetIsSingleStepInstrumentationEnabled gets a reference to the given bool and assigns it to the IsSingleStepInstrumentationEnabled field.
func (o *FleetAgentAttributes) SetIsSingleStepInstrumentationEnabled(v bool) {
	o.IsSingleStepInstrumentationEnabled = &v
}

// GetLastRestartAt returns the LastRestartAt field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetLastRestartAt() int64 {
	if o == nil || o.LastRestartAt == nil {
		var ret int64
		return ret
	}
	return *o.LastRestartAt
}

// GetLastRestartAtOk returns a tuple with the LastRestartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetLastRestartAtOk() (*int64, bool) {
	if o == nil || o.LastRestartAt == nil {
		return nil, false
	}
	return o.LastRestartAt, true
}

// HasLastRestartAt returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasLastRestartAt() bool {
	return o != nil && o.LastRestartAt != nil
}

// SetLastRestartAt gets a reference to the given int64 and assigns it to the LastRestartAt field.
func (o *FleetAgentAttributes) SetLastRestartAt(v int64) {
	o.LastRestartAt = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetOs() string {
	if o == nil || o.Os == nil {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetOsOk() (*string, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasOs() bool {
	return o != nil && o.Os != nil
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *FleetAgentAttributes) SetOs(v string) {
	o.Os = &v
}

// GetOtelCollectorVersion returns the OtelCollectorVersion field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetOtelCollectorVersion() string {
	if o == nil || o.OtelCollectorVersion == nil {
		var ret string
		return ret
	}
	return *o.OtelCollectorVersion
}

// GetOtelCollectorVersionOk returns a tuple with the OtelCollectorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetOtelCollectorVersionOk() (*string, bool) {
	if o == nil || o.OtelCollectorVersion == nil {
		return nil, false
	}
	return o.OtelCollectorVersion, true
}

// HasOtelCollectorVersion returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasOtelCollectorVersion() bool {
	return o != nil && o.OtelCollectorVersion != nil
}

// SetOtelCollectorVersion gets a reference to the given string and assigns it to the OtelCollectorVersion field.
func (o *FleetAgentAttributes) SetOtelCollectorVersion(v string) {
	o.OtelCollectorVersion = &v
}

// GetOtelCollectorVersions returns the OtelCollectorVersions field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetOtelCollectorVersions() []string {
	if o == nil || o.OtelCollectorVersions == nil {
		var ret []string
		return ret
	}
	return o.OtelCollectorVersions
}

// GetOtelCollectorVersionsOk returns a tuple with the OtelCollectorVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetOtelCollectorVersionsOk() (*[]string, bool) {
	if o == nil || o.OtelCollectorVersions == nil {
		return nil, false
	}
	return &o.OtelCollectorVersions, true
}

// HasOtelCollectorVersions returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasOtelCollectorVersions() bool {
	return o != nil && o.OtelCollectorVersions != nil
}

// SetOtelCollectorVersions gets a reference to the given []string and assigns it to the OtelCollectorVersions field.
func (o *FleetAgentAttributes) SetOtelCollectorVersions(v []string) {
	o.OtelCollectorVersions = v
}

// GetPodName returns the PodName field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetPodName() string {
	if o == nil || o.PodName == nil {
		var ret string
		return ret
	}
	return *o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetPodNameOk() (*string, bool) {
	if o == nil || o.PodName == nil {
		return nil, false
	}
	return o.PodName, true
}

// HasPodName returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasPodName() bool {
	return o != nil && o.PodName != nil
}

// SetPodName gets a reference to the given string and assigns it to the PodName field.
func (o *FleetAgentAttributes) SetPodName(v string) {
	o.PodName = &v
}

// GetRemoteAgentManagement returns the RemoteAgentManagement field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetRemoteAgentManagement() string {
	if o == nil || o.RemoteAgentManagement == nil {
		var ret string
		return ret
	}
	return *o.RemoteAgentManagement
}

// GetRemoteAgentManagementOk returns a tuple with the RemoteAgentManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetRemoteAgentManagementOk() (*string, bool) {
	if o == nil || o.RemoteAgentManagement == nil {
		return nil, false
	}
	return o.RemoteAgentManagement, true
}

// HasRemoteAgentManagement returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasRemoteAgentManagement() bool {
	return o != nil && o.RemoteAgentManagement != nil
}

// SetRemoteAgentManagement gets a reference to the given string and assigns it to the RemoteAgentManagement field.
func (o *FleetAgentAttributes) SetRemoteAgentManagement(v string) {
	o.RemoteAgentManagement = &v
}

// GetRemoteConfigStatus returns the RemoteConfigStatus field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetRemoteConfigStatus() string {
	if o == nil || o.RemoteConfigStatus == nil {
		var ret string
		return ret
	}
	return *o.RemoteConfigStatus
}

// GetRemoteConfigStatusOk returns a tuple with the RemoteConfigStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetRemoteConfigStatusOk() (*string, bool) {
	if o == nil || o.RemoteConfigStatus == nil {
		return nil, false
	}
	return o.RemoteConfigStatus, true
}

// HasRemoteConfigStatus returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasRemoteConfigStatus() bool {
	return o != nil && o.RemoteConfigStatus != nil
}

// SetRemoteConfigStatus gets a reference to the given string and assigns it to the RemoteConfigStatus field.
func (o *FleetAgentAttributes) SetRemoteConfigStatus(v string) {
	o.RemoteConfigStatus = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasServices() bool {
	return o != nil && o.Services != nil
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *FleetAgentAttributes) SetServices(v []string) {
	o.Services = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetTags() []FleetAgentAttributesTagsItems {
	if o == nil || o.Tags == nil {
		var ret []FleetAgentAttributesTagsItems
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetTagsOk() (*[]FleetAgentAttributesTagsItems, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []FleetAgentAttributesTagsItems and assigns it to the Tags field.
func (o *FleetAgentAttributes) SetTags(v []FleetAgentAttributesTagsItems) {
	o.Tags = v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *FleetAgentAttributes) GetTeam() string {
	if o == nil || o.Team == nil {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentAttributes) GetTeamOk() (*string, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *FleetAgentAttributes) HasTeam() bool {
	return o != nil && o.Team != nil
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *FleetAgentAttributes) SetTeam(v string) {
	o.Team = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetAgentAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AgentVersion != nil {
		toSerialize["agent_version"] = o.AgentVersion
	}
	if o.ApiKeyName != nil {
		toSerialize["api_key_name"] = o.ApiKeyName
	}
	if o.ApiKeyUuid != nil {
		toSerialize["api_key_uuid"] = o.ApiKeyUuid
	}
	if o.CloudProvider != nil {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if o.ClusterName != nil {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if o.DatadogAgentKey != nil {
		toSerialize["datadog_agent_key"] = o.DatadogAgentKey
	}
	if o.EnabledProducts != nil {
		toSerialize["enabled_products"] = o.EnabledProducts
	}
	if o.Envs != nil {
		toSerialize["envs"] = o.Envs
	}
	if o.FirstSeenAt != nil {
		toSerialize["first_seen_at"] = o.FirstSeenAt
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.IpAddresses != nil {
		toSerialize["ip_addresses"] = o.IpAddresses
	}
	if o.IsSingleStepInstrumentationEnabled != nil {
		toSerialize["is_single_step_instrumentation_enabled"] = o.IsSingleStepInstrumentationEnabled
	}
	if o.LastRestartAt != nil {
		toSerialize["last_restart_at"] = o.LastRestartAt
	}
	if o.Os != nil {
		toSerialize["os"] = o.Os
	}
	if o.OtelCollectorVersion != nil {
		toSerialize["otel_collector_version"] = o.OtelCollectorVersion
	}
	if o.OtelCollectorVersions != nil {
		toSerialize["otel_collector_versions"] = o.OtelCollectorVersions
	}
	if o.PodName != nil {
		toSerialize["pod_name"] = o.PodName
	}
	if o.RemoteAgentManagement != nil {
		toSerialize["remote_agent_management"] = o.RemoteAgentManagement
	}
	if o.RemoteConfigStatus != nil {
		toSerialize["remote_config_status"] = o.RemoteConfigStatus
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Team != nil {
		toSerialize["team"] = o.Team
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetAgentAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentVersion                       *string                         `json:"agent_version,omitempty"`
		ApiKeyName                         *string                         `json:"api_key_name,omitempty"`
		ApiKeyUuid                         *string                         `json:"api_key_uuid,omitempty"`
		CloudProvider                      *string                         `json:"cloud_provider,omitempty"`
		ClusterName                        *string                         `json:"cluster_name,omitempty"`
		DatadogAgentKey                    *string                         `json:"datadog_agent_key,omitempty"`
		EnabledProducts                    []string                        `json:"enabled_products,omitempty"`
		Envs                               []string                        `json:"envs,omitempty"`
		FirstSeenAt                        *int64                          `json:"first_seen_at,omitempty"`
		Hostname                           *string                         `json:"hostname,omitempty"`
		IpAddresses                        []string                        `json:"ip_addresses,omitempty"`
		IsSingleStepInstrumentationEnabled *bool                           `json:"is_single_step_instrumentation_enabled,omitempty"`
		LastRestartAt                      *int64                          `json:"last_restart_at,omitempty"`
		Os                                 *string                         `json:"os,omitempty"`
		OtelCollectorVersion               *string                         `json:"otel_collector_version,omitempty"`
		OtelCollectorVersions              []string                        `json:"otel_collector_versions,omitempty"`
		PodName                            *string                         `json:"pod_name,omitempty"`
		RemoteAgentManagement              *string                         `json:"remote_agent_management,omitempty"`
		RemoteConfigStatus                 *string                         `json:"remote_config_status,omitempty"`
		Services                           []string                        `json:"services,omitempty"`
		Tags                               []FleetAgentAttributesTagsItems `json:"tags,omitempty"`
		Team                               *string                         `json:"team,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_version", "api_key_name", "api_key_uuid", "cloud_provider", "cluster_name", "datadog_agent_key", "enabled_products", "envs", "first_seen_at", "hostname", "ip_addresses", "is_single_step_instrumentation_enabled", "last_restart_at", "os", "otel_collector_version", "otel_collector_versions", "pod_name", "remote_agent_management", "remote_config_status", "services", "tags", "team"})
	} else {
		return err
	}
	o.AgentVersion = all.AgentVersion
	o.ApiKeyName = all.ApiKeyName
	o.ApiKeyUuid = all.ApiKeyUuid
	o.CloudProvider = all.CloudProvider
	o.ClusterName = all.ClusterName
	o.DatadogAgentKey = all.DatadogAgentKey
	o.EnabledProducts = all.EnabledProducts
	o.Envs = all.Envs
	o.FirstSeenAt = all.FirstSeenAt
	o.Hostname = all.Hostname
	o.IpAddresses = all.IpAddresses
	o.IsSingleStepInstrumentationEnabled = all.IsSingleStepInstrumentationEnabled
	o.LastRestartAt = all.LastRestartAt
	o.Os = all.Os
	o.OtelCollectorVersion = all.OtelCollectorVersion
	o.OtelCollectorVersions = all.OtelCollectorVersions
	o.PodName = all.PodName
	o.RemoteAgentManagement = all.RemoteAgentManagement
	o.RemoteConfigStatus = all.RemoteConfigStatus
	o.Services = all.Services
	o.Tags = all.Tags
	o.Team = all.Team

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
