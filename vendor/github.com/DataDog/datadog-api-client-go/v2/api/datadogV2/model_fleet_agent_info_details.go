// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAgentInfoDetails Detailed information about a Datadog Agent.
type FleetAgentInfoDetails struct {
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
	Env []string `json:"env,omitempty"`
	// Timestamp when the agent was first seen.
	FirstSeenAt *int64 `json:"first_seen_at,omitempty"`
	// The hostname of the agent.
	Hostname *string `json:"hostname,omitempty"`
	// Alternative hostname list for the agent.
	HostnameAliases []string `json:"hostname_aliases,omitempty"`
	// The version of the installer used.
	InstallMethodInstallerVersion *string `json:"install_method_installer_version,omitempty"`
	// The tool used to install the agent.
	InstallMethodTool *string `json:"install_method_tool,omitempty"`
	// IP addresses of the agent.
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Whether single-step instrumentation is enabled.
	IsSingleStepInstrumentationEnabled *bool `json:"is_single_step_instrumentation_enabled,omitempty"`
	// Timestamp of the last agent restart.
	LastRestartAt *int64 `json:"last_restart_at,omitempty"`
	// The operating system.
	Os *string `json:"os,omitempty"`
	// The operating system version.
	OsVersion *string `json:"os_version,omitempty"`
	// OpenTelemetry collector version (if applicable).
	OtelCollectorVersion *string `json:"otel_collector_version,omitempty"`
	// List of OpenTelemetry collector versions (if applicable).
	OtelCollectorVersions []string `json:"otel_collector_versions,omitempty"`
	// OpenTelemetry collectors associated with the agent (if applicable).
	OtelCollectors []map[string]interface{} `json:"otel_collectors,omitempty"`
	// Kubernetes pod name (if applicable).
	PodName *string `json:"pod_name,omitempty"`
	// The Python version used by the agent.
	PythonVersion *string `json:"python_version,omitempty"`
	// Regions where the agent is running.
	Region []string `json:"region,omitempty"`
	// Remote agent management status.
	RemoteAgentManagement *string `json:"remote_agent_management,omitempty"`
	// Remote configuration status.
	RemoteConfigStatus *string `json:"remote_config_status,omitempty"`
	// Services running on the agent.
	Services []string `json:"services,omitempty"`
	// Tags associated with the agent.
	Tags []string `json:"tags,omitempty"`
	// Team associated with the agent.
	Team *string `json:"team,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetAgentInfoDetails instantiates a new FleetAgentInfoDetails object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetAgentInfoDetails() *FleetAgentInfoDetails {
	this := FleetAgentInfoDetails{}
	return &this
}

// NewFleetAgentInfoDetailsWithDefaults instantiates a new FleetAgentInfoDetails object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetAgentInfoDetailsWithDefaults() *FleetAgentInfoDetails {
	this := FleetAgentInfoDetails{}
	return &this
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetAgentVersion() string {
	if o == nil || o.AgentVersion == nil {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetAgentVersionOk() (*string, bool) {
	if o == nil || o.AgentVersion == nil {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasAgentVersion() bool {
	return o != nil && o.AgentVersion != nil
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *FleetAgentInfoDetails) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetApiKeyName returns the ApiKeyName field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetApiKeyName() string {
	if o == nil || o.ApiKeyName == nil {
		var ret string
		return ret
	}
	return *o.ApiKeyName
}

// GetApiKeyNameOk returns a tuple with the ApiKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetApiKeyNameOk() (*string, bool) {
	if o == nil || o.ApiKeyName == nil {
		return nil, false
	}
	return o.ApiKeyName, true
}

// HasApiKeyName returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasApiKeyName() bool {
	return o != nil && o.ApiKeyName != nil
}

// SetApiKeyName gets a reference to the given string and assigns it to the ApiKeyName field.
func (o *FleetAgentInfoDetails) SetApiKeyName(v string) {
	o.ApiKeyName = &v
}

// GetApiKeyUuid returns the ApiKeyUuid field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetApiKeyUuid() string {
	if o == nil || o.ApiKeyUuid == nil {
		var ret string
		return ret
	}
	return *o.ApiKeyUuid
}

// GetApiKeyUuidOk returns a tuple with the ApiKeyUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetApiKeyUuidOk() (*string, bool) {
	if o == nil || o.ApiKeyUuid == nil {
		return nil, false
	}
	return o.ApiKeyUuid, true
}

// HasApiKeyUuid returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasApiKeyUuid() bool {
	return o != nil && o.ApiKeyUuid != nil
}

// SetApiKeyUuid gets a reference to the given string and assigns it to the ApiKeyUuid field.
func (o *FleetAgentInfoDetails) SetApiKeyUuid(v string) {
	o.ApiKeyUuid = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasCloudProvider() bool {
	return o != nil && o.CloudProvider != nil
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *FleetAgentInfoDetails) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *FleetAgentInfoDetails) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetDatadogAgentKey returns the DatadogAgentKey field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetDatadogAgentKey() string {
	if o == nil || o.DatadogAgentKey == nil {
		var ret string
		return ret
	}
	return *o.DatadogAgentKey
}

// GetDatadogAgentKeyOk returns a tuple with the DatadogAgentKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetDatadogAgentKeyOk() (*string, bool) {
	if o == nil || o.DatadogAgentKey == nil {
		return nil, false
	}
	return o.DatadogAgentKey, true
}

// HasDatadogAgentKey returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasDatadogAgentKey() bool {
	return o != nil && o.DatadogAgentKey != nil
}

// SetDatadogAgentKey gets a reference to the given string and assigns it to the DatadogAgentKey field.
func (o *FleetAgentInfoDetails) SetDatadogAgentKey(v string) {
	o.DatadogAgentKey = &v
}

// GetEnabledProducts returns the EnabledProducts field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetEnabledProducts() []string {
	if o == nil || o.EnabledProducts == nil {
		var ret []string
		return ret
	}
	return o.EnabledProducts
}

// GetEnabledProductsOk returns a tuple with the EnabledProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetEnabledProductsOk() (*[]string, bool) {
	if o == nil || o.EnabledProducts == nil {
		return nil, false
	}
	return &o.EnabledProducts, true
}

// HasEnabledProducts returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasEnabledProducts() bool {
	return o != nil && o.EnabledProducts != nil
}

// SetEnabledProducts gets a reference to the given []string and assigns it to the EnabledProducts field.
func (o *FleetAgentInfoDetails) SetEnabledProducts(v []string) {
	o.EnabledProducts = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetEnv() []string {
	if o == nil || o.Env == nil {
		var ret []string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetEnvOk() (*[]string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return &o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given []string and assigns it to the Env field.
func (o *FleetAgentInfoDetails) SetEnv(v []string) {
	o.Env = v
}

// GetFirstSeenAt returns the FirstSeenAt field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetFirstSeenAt() int64 {
	if o == nil || o.FirstSeenAt == nil {
		var ret int64
		return ret
	}
	return *o.FirstSeenAt
}

// GetFirstSeenAtOk returns a tuple with the FirstSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetFirstSeenAtOk() (*int64, bool) {
	if o == nil || o.FirstSeenAt == nil {
		return nil, false
	}
	return o.FirstSeenAt, true
}

// HasFirstSeenAt returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasFirstSeenAt() bool {
	return o != nil && o.FirstSeenAt != nil
}

// SetFirstSeenAt gets a reference to the given int64 and assigns it to the FirstSeenAt field.
func (o *FleetAgentInfoDetails) SetFirstSeenAt(v int64) {
	o.FirstSeenAt = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *FleetAgentInfoDetails) SetHostname(v string) {
	o.Hostname = &v
}

// GetHostnameAliases returns the HostnameAliases field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetHostnameAliases() []string {
	if o == nil || o.HostnameAliases == nil {
		var ret []string
		return ret
	}
	return o.HostnameAliases
}

// GetHostnameAliasesOk returns a tuple with the HostnameAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetHostnameAliasesOk() (*[]string, bool) {
	if o == nil || o.HostnameAliases == nil {
		return nil, false
	}
	return &o.HostnameAliases, true
}

// HasHostnameAliases returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasHostnameAliases() bool {
	return o != nil && o.HostnameAliases != nil
}

// SetHostnameAliases gets a reference to the given []string and assigns it to the HostnameAliases field.
func (o *FleetAgentInfoDetails) SetHostnameAliases(v []string) {
	o.HostnameAliases = v
}

// GetInstallMethodInstallerVersion returns the InstallMethodInstallerVersion field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetInstallMethodInstallerVersion() string {
	if o == nil || o.InstallMethodInstallerVersion == nil {
		var ret string
		return ret
	}
	return *o.InstallMethodInstallerVersion
}

// GetInstallMethodInstallerVersionOk returns a tuple with the InstallMethodInstallerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetInstallMethodInstallerVersionOk() (*string, bool) {
	if o == nil || o.InstallMethodInstallerVersion == nil {
		return nil, false
	}
	return o.InstallMethodInstallerVersion, true
}

// HasInstallMethodInstallerVersion returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasInstallMethodInstallerVersion() bool {
	return o != nil && o.InstallMethodInstallerVersion != nil
}

// SetInstallMethodInstallerVersion gets a reference to the given string and assigns it to the InstallMethodInstallerVersion field.
func (o *FleetAgentInfoDetails) SetInstallMethodInstallerVersion(v string) {
	o.InstallMethodInstallerVersion = &v
}

// GetInstallMethodTool returns the InstallMethodTool field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetInstallMethodTool() string {
	if o == nil || o.InstallMethodTool == nil {
		var ret string
		return ret
	}
	return *o.InstallMethodTool
}

// GetInstallMethodToolOk returns a tuple with the InstallMethodTool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetInstallMethodToolOk() (*string, bool) {
	if o == nil || o.InstallMethodTool == nil {
		return nil, false
	}
	return o.InstallMethodTool, true
}

// HasInstallMethodTool returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasInstallMethodTool() bool {
	return o != nil && o.InstallMethodTool != nil
}

// SetInstallMethodTool gets a reference to the given string and assigns it to the InstallMethodTool field.
func (o *FleetAgentInfoDetails) SetInstallMethodTool(v string) {
	o.InstallMethodTool = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetIpAddresses() []string {
	if o == nil || o.IpAddresses == nil {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return &o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasIpAddresses() bool {
	return o != nil && o.IpAddresses != nil
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *FleetAgentInfoDetails) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetIsSingleStepInstrumentationEnabled returns the IsSingleStepInstrumentationEnabled field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetIsSingleStepInstrumentationEnabled() bool {
	if o == nil || o.IsSingleStepInstrumentationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsSingleStepInstrumentationEnabled
}

// GetIsSingleStepInstrumentationEnabledOk returns a tuple with the IsSingleStepInstrumentationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetIsSingleStepInstrumentationEnabledOk() (*bool, bool) {
	if o == nil || o.IsSingleStepInstrumentationEnabled == nil {
		return nil, false
	}
	return o.IsSingleStepInstrumentationEnabled, true
}

// HasIsSingleStepInstrumentationEnabled returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasIsSingleStepInstrumentationEnabled() bool {
	return o != nil && o.IsSingleStepInstrumentationEnabled != nil
}

// SetIsSingleStepInstrumentationEnabled gets a reference to the given bool and assigns it to the IsSingleStepInstrumentationEnabled field.
func (o *FleetAgentInfoDetails) SetIsSingleStepInstrumentationEnabled(v bool) {
	o.IsSingleStepInstrumentationEnabled = &v
}

// GetLastRestartAt returns the LastRestartAt field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetLastRestartAt() int64 {
	if o == nil || o.LastRestartAt == nil {
		var ret int64
		return ret
	}
	return *o.LastRestartAt
}

// GetLastRestartAtOk returns a tuple with the LastRestartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetLastRestartAtOk() (*int64, bool) {
	if o == nil || o.LastRestartAt == nil {
		return nil, false
	}
	return o.LastRestartAt, true
}

// HasLastRestartAt returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasLastRestartAt() bool {
	return o != nil && o.LastRestartAt != nil
}

// SetLastRestartAt gets a reference to the given int64 and assigns it to the LastRestartAt field.
func (o *FleetAgentInfoDetails) SetLastRestartAt(v int64) {
	o.LastRestartAt = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetOs() string {
	if o == nil || o.Os == nil {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetOsOk() (*string, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasOs() bool {
	return o != nil && o.Os != nil
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *FleetAgentInfoDetails) SetOs(v string) {
	o.Os = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasOsVersion() bool {
	return o != nil && o.OsVersion != nil
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *FleetAgentInfoDetails) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetOtelCollectorVersion returns the OtelCollectorVersion field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetOtelCollectorVersion() string {
	if o == nil || o.OtelCollectorVersion == nil {
		var ret string
		return ret
	}
	return *o.OtelCollectorVersion
}

// GetOtelCollectorVersionOk returns a tuple with the OtelCollectorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetOtelCollectorVersionOk() (*string, bool) {
	if o == nil || o.OtelCollectorVersion == nil {
		return nil, false
	}
	return o.OtelCollectorVersion, true
}

// HasOtelCollectorVersion returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasOtelCollectorVersion() bool {
	return o != nil && o.OtelCollectorVersion != nil
}

// SetOtelCollectorVersion gets a reference to the given string and assigns it to the OtelCollectorVersion field.
func (o *FleetAgentInfoDetails) SetOtelCollectorVersion(v string) {
	o.OtelCollectorVersion = &v
}

// GetOtelCollectorVersions returns the OtelCollectorVersions field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetOtelCollectorVersions() []string {
	if o == nil || o.OtelCollectorVersions == nil {
		var ret []string
		return ret
	}
	return o.OtelCollectorVersions
}

// GetOtelCollectorVersionsOk returns a tuple with the OtelCollectorVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetOtelCollectorVersionsOk() (*[]string, bool) {
	if o == nil || o.OtelCollectorVersions == nil {
		return nil, false
	}
	return &o.OtelCollectorVersions, true
}

// HasOtelCollectorVersions returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasOtelCollectorVersions() bool {
	return o != nil && o.OtelCollectorVersions != nil
}

// SetOtelCollectorVersions gets a reference to the given []string and assigns it to the OtelCollectorVersions field.
func (o *FleetAgentInfoDetails) SetOtelCollectorVersions(v []string) {
	o.OtelCollectorVersions = v
}

// GetOtelCollectors returns the OtelCollectors field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetOtelCollectors() []map[string]interface{} {
	if o == nil || o.OtelCollectors == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.OtelCollectors
}

// GetOtelCollectorsOk returns a tuple with the OtelCollectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetOtelCollectorsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.OtelCollectors == nil {
		return nil, false
	}
	return &o.OtelCollectors, true
}

// HasOtelCollectors returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasOtelCollectors() bool {
	return o != nil && o.OtelCollectors != nil
}

// SetOtelCollectors gets a reference to the given []map[string]interface{} and assigns it to the OtelCollectors field.
func (o *FleetAgentInfoDetails) SetOtelCollectors(v []map[string]interface{}) {
	o.OtelCollectors = v
}

// GetPodName returns the PodName field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetPodName() string {
	if o == nil || o.PodName == nil {
		var ret string
		return ret
	}
	return *o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetPodNameOk() (*string, bool) {
	if o == nil || o.PodName == nil {
		return nil, false
	}
	return o.PodName, true
}

// HasPodName returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasPodName() bool {
	return o != nil && o.PodName != nil
}

// SetPodName gets a reference to the given string and assigns it to the PodName field.
func (o *FleetAgentInfoDetails) SetPodName(v string) {
	o.PodName = &v
}

// GetPythonVersion returns the PythonVersion field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetPythonVersion() string {
	if o == nil || o.PythonVersion == nil {
		var ret string
		return ret
	}
	return *o.PythonVersion
}

// GetPythonVersionOk returns a tuple with the PythonVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetPythonVersionOk() (*string, bool) {
	if o == nil || o.PythonVersion == nil {
		return nil, false
	}
	return o.PythonVersion, true
}

// HasPythonVersion returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasPythonVersion() bool {
	return o != nil && o.PythonVersion != nil
}

// SetPythonVersion gets a reference to the given string and assigns it to the PythonVersion field.
func (o *FleetAgentInfoDetails) SetPythonVersion(v string) {
	o.PythonVersion = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetRegion() []string {
	if o == nil || o.Region == nil {
		var ret []string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetRegionOk() (*[]string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return &o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given []string and assigns it to the Region field.
func (o *FleetAgentInfoDetails) SetRegion(v []string) {
	o.Region = v
}

// GetRemoteAgentManagement returns the RemoteAgentManagement field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetRemoteAgentManagement() string {
	if o == nil || o.RemoteAgentManagement == nil {
		var ret string
		return ret
	}
	return *o.RemoteAgentManagement
}

// GetRemoteAgentManagementOk returns a tuple with the RemoteAgentManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetRemoteAgentManagementOk() (*string, bool) {
	if o == nil || o.RemoteAgentManagement == nil {
		return nil, false
	}
	return o.RemoteAgentManagement, true
}

// HasRemoteAgentManagement returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasRemoteAgentManagement() bool {
	return o != nil && o.RemoteAgentManagement != nil
}

// SetRemoteAgentManagement gets a reference to the given string and assigns it to the RemoteAgentManagement field.
func (o *FleetAgentInfoDetails) SetRemoteAgentManagement(v string) {
	o.RemoteAgentManagement = &v
}

// GetRemoteConfigStatus returns the RemoteConfigStatus field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetRemoteConfigStatus() string {
	if o == nil || o.RemoteConfigStatus == nil {
		var ret string
		return ret
	}
	return *o.RemoteConfigStatus
}

// GetRemoteConfigStatusOk returns a tuple with the RemoteConfigStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetRemoteConfigStatusOk() (*string, bool) {
	if o == nil || o.RemoteConfigStatus == nil {
		return nil, false
	}
	return o.RemoteConfigStatus, true
}

// HasRemoteConfigStatus returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasRemoteConfigStatus() bool {
	return o != nil && o.RemoteConfigStatus != nil
}

// SetRemoteConfigStatus gets a reference to the given string and assigns it to the RemoteConfigStatus field.
func (o *FleetAgentInfoDetails) SetRemoteConfigStatus(v string) {
	o.RemoteConfigStatus = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasServices() bool {
	return o != nil && o.Services != nil
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *FleetAgentInfoDetails) SetServices(v []string) {
	o.Services = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FleetAgentInfoDetails) SetTags(v []string) {
	o.Tags = v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *FleetAgentInfoDetails) GetTeam() string {
	if o == nil || o.Team == nil {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoDetails) GetTeamOk() (*string, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *FleetAgentInfoDetails) HasTeam() bool {
	return o != nil && o.Team != nil
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *FleetAgentInfoDetails) SetTeam(v string) {
	o.Team = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetAgentInfoDetails) MarshalJSON() ([]byte, error) {
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
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.FirstSeenAt != nil {
		toSerialize["first_seen_at"] = o.FirstSeenAt
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.HostnameAliases != nil {
		toSerialize["hostname_aliases"] = o.HostnameAliases
	}
	if o.InstallMethodInstallerVersion != nil {
		toSerialize["install_method_installer_version"] = o.InstallMethodInstallerVersion
	}
	if o.InstallMethodTool != nil {
		toSerialize["install_method_tool"] = o.InstallMethodTool
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
	if o.OsVersion != nil {
		toSerialize["os_version"] = o.OsVersion
	}
	if o.OtelCollectorVersion != nil {
		toSerialize["otel_collector_version"] = o.OtelCollectorVersion
	}
	if o.OtelCollectorVersions != nil {
		toSerialize["otel_collector_versions"] = o.OtelCollectorVersions
	}
	if o.OtelCollectors != nil {
		toSerialize["otel_collectors"] = o.OtelCollectors
	}
	if o.PodName != nil {
		toSerialize["pod_name"] = o.PodName
	}
	if o.PythonVersion != nil {
		toSerialize["python_version"] = o.PythonVersion
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
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
func (o *FleetAgentInfoDetails) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentVersion                       *string                  `json:"agent_version,omitempty"`
		ApiKeyName                         *string                  `json:"api_key_name,omitempty"`
		ApiKeyUuid                         *string                  `json:"api_key_uuid,omitempty"`
		CloudProvider                      *string                  `json:"cloud_provider,omitempty"`
		ClusterName                        *string                  `json:"cluster_name,omitempty"`
		DatadogAgentKey                    *string                  `json:"datadog_agent_key,omitempty"`
		EnabledProducts                    []string                 `json:"enabled_products,omitempty"`
		Env                                []string                 `json:"env,omitempty"`
		FirstSeenAt                        *int64                   `json:"first_seen_at,omitempty"`
		Hostname                           *string                  `json:"hostname,omitempty"`
		HostnameAliases                    []string                 `json:"hostname_aliases,omitempty"`
		InstallMethodInstallerVersion      *string                  `json:"install_method_installer_version,omitempty"`
		InstallMethodTool                  *string                  `json:"install_method_tool,omitempty"`
		IpAddresses                        []string                 `json:"ip_addresses,omitempty"`
		IsSingleStepInstrumentationEnabled *bool                    `json:"is_single_step_instrumentation_enabled,omitempty"`
		LastRestartAt                      *int64                   `json:"last_restart_at,omitempty"`
		Os                                 *string                  `json:"os,omitempty"`
		OsVersion                          *string                  `json:"os_version,omitempty"`
		OtelCollectorVersion               *string                  `json:"otel_collector_version,omitempty"`
		OtelCollectorVersions              []string                 `json:"otel_collector_versions,omitempty"`
		OtelCollectors                     []map[string]interface{} `json:"otel_collectors,omitempty"`
		PodName                            *string                  `json:"pod_name,omitempty"`
		PythonVersion                      *string                  `json:"python_version,omitempty"`
		Region                             []string                 `json:"region,omitempty"`
		RemoteAgentManagement              *string                  `json:"remote_agent_management,omitempty"`
		RemoteConfigStatus                 *string                  `json:"remote_config_status,omitempty"`
		Services                           []string                 `json:"services,omitempty"`
		Tags                               []string                 `json:"tags,omitempty"`
		Team                               *string                  `json:"team,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_version", "api_key_name", "api_key_uuid", "cloud_provider", "cluster_name", "datadog_agent_key", "enabled_products", "env", "first_seen_at", "hostname", "hostname_aliases", "install_method_installer_version", "install_method_tool", "ip_addresses", "is_single_step_instrumentation_enabled", "last_restart_at", "os", "os_version", "otel_collector_version", "otel_collector_versions", "otel_collectors", "pod_name", "python_version", "region", "remote_agent_management", "remote_config_status", "services", "tags", "team"})
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
	o.Env = all.Env
	o.FirstSeenAt = all.FirstSeenAt
	o.Hostname = all.Hostname
	o.HostnameAliases = all.HostnameAliases
	o.InstallMethodInstallerVersion = all.InstallMethodInstallerVersion
	o.InstallMethodTool = all.InstallMethodTool
	o.IpAddresses = all.IpAddresses
	o.IsSingleStepInstrumentationEnabled = all.IsSingleStepInstrumentationEnabled
	o.LastRestartAt = all.LastRestartAt
	o.Os = all.Os
	o.OsVersion = all.OsVersion
	o.OtelCollectorVersion = all.OtelCollectorVersion
	o.OtelCollectorVersions = all.OtelCollectorVersions
	o.OtelCollectors = all.OtelCollectors
	o.PodName = all.PodName
	o.PythonVersion = all.PythonVersion
	o.Region = all.Region
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
