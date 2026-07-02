// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmUnifiedHostAttributes Attributes of a unified host, combining data from agent and agentless sources.
type CsmUnifiedHostAttributes struct {
	// The ID of the cloud account that the host belongs to. Present only when the host was discovered through agentless scanning.
	AccountId datadog.NullableString `json:"account_id,omitempty"`
	// Whether CSM Vulnerabilities is enabled for containers through the Datadog Agent. `true` if enabled; `false` if disabled.
	AgentCsmVmContainersEnabled datadog.NullableBool `json:"agent_csm_vm_containers_enabled,omitempty"`
	// Whether CSM Vulnerabilities is enabled for hosts through the Datadog Agent. `true` if enabled; `false` if disabled.
	AgentCsmVmHostsEnabled datadog.NullableBool `json:"agent_csm_vm_hosts_enabled,omitempty"`
	// Whether CSM Threats is enabled for this host through the Datadog Agent. `true` if enabled; `false` if disabled.
	AgentCwsEnabled datadog.NullableBool `json:"agent_cws_enabled,omitempty"`
	// Whether CSM Misconfigurations is enabled for this host through the Datadog Agent. `true` if enabled; `false` if disabled.
	AgentPostureManagement datadog.NullableBool `json:"agent_posture_management,omitempty"`
	// The version of the Datadog Agent running on this host.
	AgentVersion datadog.NullableString `json:"agent_version,omitempty"`
	// Whether CSM Misconfigurations is enabled for this host via agentless scanning. `true` if enabled; `false` if disabled.
	AgentlessPostureManagement datadog.NullableBool `json:"agentless_posture_management,omitempty"`
	// Whether CSM Vulnerabilities is enabled for this host via agentless scanning. `true` if enabled; `false` if disabled.
	AgentlessVulnerabilityScanning datadog.NullableBool `json:"agentless_vulnerability_scanning,omitempty"`
	// The cloud provider of a host resource.
	CloudProvider *CsmCloudProvider `json:"cloud_provider,omitempty"`
	// The name of the Kubernetes cluster the host belongs to, if applicable.
	ClusterName datadog.NullableString `json:"cluster_name,omitempty"`
	// The Datadog Agent key associated with this host. Present only for agent-sourced hosts.
	DatadogAgentKey datadog.NullableString `json:"datadog_agent_key,omitempty"`
	// The list of environment tags associated with this host.
	Env datadog.NullableList[string] `json:"env,omitempty"`
	// The internal Datadog host identifier. Present only for agent-sourced hosts.
	HostId datadog.NullableInt64 `json:"host_id,omitempty"`
	// The tool used to install the Datadog Agent on this host.
	InstallMethodTool datadog.NullableString `json:"install_method_tool,omitempty"`
	// The operating system of the host. Present only for agent-sourced hosts.
	Os datadog.NullableString `json:"os,omitempty"`
	// The type of cloud resource for an agentless host.
	ResourceType *CsmAgentlessHostResourceType `json:"resource_type,omitempty"`
	// The source of a unified host entry, indicating whether it was discovered via agent, agentless scanning, or both.
	Source CsmUnifiedHostSource `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCsmUnifiedHostAttributes instantiates a new CsmUnifiedHostAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCsmUnifiedHostAttributes(source CsmUnifiedHostSource) *CsmUnifiedHostAttributes {
	this := CsmUnifiedHostAttributes{}
	this.Source = source
	return &this
}

// NewCsmUnifiedHostAttributesWithDefaults instantiates a new CsmUnifiedHostAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCsmUnifiedHostAttributesWithDefaults() *CsmUnifiedHostAttributes {
	this := CsmUnifiedHostAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmUnifiedHostAttributes) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmUnifiedHostAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// HasAccountId returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasAccountId() bool {
	return o != nil && o.AccountId.IsSet()
}

// SetAccountId gets a reference to the given datadog.NullableString and assigns it to the AccountId field.
func (o *CsmUnifiedHostAttributes) SetAccountId(v string) {
	o.AccountId.Set(&v)
}

// SetAccountIdNil sets the value for AccountId to be an explicit nil.
func (o *CsmUnifiedHostAttributes) SetAccountIdNil() {
	o.AccountId.Set(nil)
}

// UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil.
func (o *CsmUnifiedHostAttributes) UnsetAccountId() {
	o.AccountId.Unset()
}

// GetAgentCsmVmContainersEnabled returns the AgentCsmVmContainersEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmUnifiedHostAttributes) GetAgentCsmVmContainersEnabled() bool {
	if o == nil || o.AgentCsmVmContainersEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AgentCsmVmContainersEnabled.Get()
}

// GetAgentCsmVmContainersEnabledOk returns a tuple with the AgentCsmVmContainersEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmUnifiedHostAttributes) GetAgentCsmVmContainersEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentCsmVmContainersEnabled.Get(), o.AgentCsmVmContainersEnabled.IsSet()
}

// HasAgentCsmVmContainersEnabled returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasAgentCsmVmContainersEnabled() bool {
	return o != nil && o.AgentCsmVmContainersEnabled.IsSet()
}

// SetAgentCsmVmContainersEnabled gets a reference to the given datadog.NullableBool and assigns it to the AgentCsmVmContainersEnabled field.
func (o *CsmUnifiedHostAttributes) SetAgentCsmVmContainersEnabled(v bool) {
	o.AgentCsmVmContainersEnabled.Set(&v)
}

// SetAgentCsmVmContainersEnabledNil sets the value for AgentCsmVmContainersEnabled to be an explicit nil.
func (o *CsmUnifiedHostAttributes) SetAgentCsmVmContainersEnabledNil() {
	o.AgentCsmVmContainersEnabled.Set(nil)
}

// UnsetAgentCsmVmContainersEnabled ensures that no value is present for AgentCsmVmContainersEnabled, not even an explicit nil.
func (o *CsmUnifiedHostAttributes) UnsetAgentCsmVmContainersEnabled() {
	o.AgentCsmVmContainersEnabled.Unset()
}

// GetAgentCsmVmHostsEnabled returns the AgentCsmVmHostsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmUnifiedHostAttributes) GetAgentCsmVmHostsEnabled() bool {
	if o == nil || o.AgentCsmVmHostsEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AgentCsmVmHostsEnabled.Get()
}

// GetAgentCsmVmHostsEnabledOk returns a tuple with the AgentCsmVmHostsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmUnifiedHostAttributes) GetAgentCsmVmHostsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentCsmVmHostsEnabled.Get(), o.AgentCsmVmHostsEnabled.IsSet()
}

// HasAgentCsmVmHostsEnabled returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasAgentCsmVmHostsEnabled() bool {
	return o != nil && o.AgentCsmVmHostsEnabled.IsSet()
}

// SetAgentCsmVmHostsEnabled gets a reference to the given datadog.NullableBool and assigns it to the AgentCsmVmHostsEnabled field.
func (o *CsmUnifiedHostAttributes) SetAgentCsmVmHostsEnabled(v bool) {
	o.AgentCsmVmHostsEnabled.Set(&v)
}

// SetAgentCsmVmHostsEnabledNil sets the value for AgentCsmVmHostsEnabled to be an explicit nil.
func (o *CsmUnifiedHostAttributes) SetAgentCsmVmHostsEnabledNil() {
	o.AgentCsmVmHostsEnabled.Set(nil)
}

// UnsetAgentCsmVmHostsEnabled ensures that no value is present for AgentCsmVmHostsEnabled, not even an explicit nil.
func (o *CsmUnifiedHostAttributes) UnsetAgentCsmVmHostsEnabled() {
	o.AgentCsmVmHostsEnabled.Unset()
}

// GetAgentCwsEnabled returns the AgentCwsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmUnifiedHostAttributes) GetAgentCwsEnabled() bool {
	if o == nil || o.AgentCwsEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AgentCwsEnabled.Get()
}

// GetAgentCwsEnabledOk returns a tuple with the AgentCwsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmUnifiedHostAttributes) GetAgentCwsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentCwsEnabled.Get(), o.AgentCwsEnabled.IsSet()
}

// HasAgentCwsEnabled returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasAgentCwsEnabled() bool {
	return o != nil && o.AgentCwsEnabled.IsSet()
}

// SetAgentCwsEnabled gets a reference to the given datadog.NullableBool and assigns it to the AgentCwsEnabled field.
func (o *CsmUnifiedHostAttributes) SetAgentCwsEnabled(v bool) {
	o.AgentCwsEnabled.Set(&v)
}

// SetAgentCwsEnabledNil sets the value for AgentCwsEnabled to be an explicit nil.
func (o *CsmUnifiedHostAttributes) SetAgentCwsEnabledNil() {
	o.AgentCwsEnabled.Set(nil)
}

// UnsetAgentCwsEnabled ensures that no value is present for AgentCwsEnabled, not even an explicit nil.
func (o *CsmUnifiedHostAttributes) UnsetAgentCwsEnabled() {
	o.AgentCwsEnabled.Unset()
}

// GetAgentPostureManagement returns the AgentPostureManagement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmUnifiedHostAttributes) GetAgentPostureManagement() bool {
	if o == nil || o.AgentPostureManagement.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AgentPostureManagement.Get()
}

// GetAgentPostureManagementOk returns a tuple with the AgentPostureManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmUnifiedHostAttributes) GetAgentPostureManagementOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentPostureManagement.Get(), o.AgentPostureManagement.IsSet()
}

// HasAgentPostureManagement returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasAgentPostureManagement() bool {
	return o != nil && o.AgentPostureManagement.IsSet()
}

// SetAgentPostureManagement gets a reference to the given datadog.NullableBool and assigns it to the AgentPostureManagement field.
func (o *CsmUnifiedHostAttributes) SetAgentPostureManagement(v bool) {
	o.AgentPostureManagement.Set(&v)
}

// SetAgentPostureManagementNil sets the value for AgentPostureManagement to be an explicit nil.
func (o *CsmUnifiedHostAttributes) SetAgentPostureManagementNil() {
	o.AgentPostureManagement.Set(nil)
}

// UnsetAgentPostureManagement ensures that no value is present for AgentPostureManagement, not even an explicit nil.
func (o *CsmUnifiedHostAttributes) UnsetAgentPostureManagement() {
	o.AgentPostureManagement.Unset()
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmUnifiedHostAttributes) GetAgentVersion() string {
	if o == nil || o.AgentVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.AgentVersion.Get()
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmUnifiedHostAttributes) GetAgentVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentVersion.Get(), o.AgentVersion.IsSet()
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasAgentVersion() bool {
	return o != nil && o.AgentVersion.IsSet()
}

// SetAgentVersion gets a reference to the given datadog.NullableString and assigns it to the AgentVersion field.
func (o *CsmUnifiedHostAttributes) SetAgentVersion(v string) {
	o.AgentVersion.Set(&v)
}

// SetAgentVersionNil sets the value for AgentVersion to be an explicit nil.
func (o *CsmUnifiedHostAttributes) SetAgentVersionNil() {
	o.AgentVersion.Set(nil)
}

// UnsetAgentVersion ensures that no value is present for AgentVersion, not even an explicit nil.
func (o *CsmUnifiedHostAttributes) UnsetAgentVersion() {
	o.AgentVersion.Unset()
}

// GetAgentlessPostureManagement returns the AgentlessPostureManagement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmUnifiedHostAttributes) GetAgentlessPostureManagement() bool {
	if o == nil || o.AgentlessPostureManagement.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AgentlessPostureManagement.Get()
}

// GetAgentlessPostureManagementOk returns a tuple with the AgentlessPostureManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmUnifiedHostAttributes) GetAgentlessPostureManagementOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentlessPostureManagement.Get(), o.AgentlessPostureManagement.IsSet()
}

// HasAgentlessPostureManagement returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasAgentlessPostureManagement() bool {
	return o != nil && o.AgentlessPostureManagement.IsSet()
}

// SetAgentlessPostureManagement gets a reference to the given datadog.NullableBool and assigns it to the AgentlessPostureManagement field.
func (o *CsmUnifiedHostAttributes) SetAgentlessPostureManagement(v bool) {
	o.AgentlessPostureManagement.Set(&v)
}

// SetAgentlessPostureManagementNil sets the value for AgentlessPostureManagement to be an explicit nil.
func (o *CsmUnifiedHostAttributes) SetAgentlessPostureManagementNil() {
	o.AgentlessPostureManagement.Set(nil)
}

// UnsetAgentlessPostureManagement ensures that no value is present for AgentlessPostureManagement, not even an explicit nil.
func (o *CsmUnifiedHostAttributes) UnsetAgentlessPostureManagement() {
	o.AgentlessPostureManagement.Unset()
}

// GetAgentlessVulnerabilityScanning returns the AgentlessVulnerabilityScanning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmUnifiedHostAttributes) GetAgentlessVulnerabilityScanning() bool {
	if o == nil || o.AgentlessVulnerabilityScanning.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AgentlessVulnerabilityScanning.Get()
}

// GetAgentlessVulnerabilityScanningOk returns a tuple with the AgentlessVulnerabilityScanning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmUnifiedHostAttributes) GetAgentlessVulnerabilityScanningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentlessVulnerabilityScanning.Get(), o.AgentlessVulnerabilityScanning.IsSet()
}

// HasAgentlessVulnerabilityScanning returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasAgentlessVulnerabilityScanning() bool {
	return o != nil && o.AgentlessVulnerabilityScanning.IsSet()
}

// SetAgentlessVulnerabilityScanning gets a reference to the given datadog.NullableBool and assigns it to the AgentlessVulnerabilityScanning field.
func (o *CsmUnifiedHostAttributes) SetAgentlessVulnerabilityScanning(v bool) {
	o.AgentlessVulnerabilityScanning.Set(&v)
}

// SetAgentlessVulnerabilityScanningNil sets the value for AgentlessVulnerabilityScanning to be an explicit nil.
func (o *CsmUnifiedHostAttributes) SetAgentlessVulnerabilityScanningNil() {
	o.AgentlessVulnerabilityScanning.Set(nil)
}

// UnsetAgentlessVulnerabilityScanning ensures that no value is present for AgentlessVulnerabilityScanning, not even an explicit nil.
func (o *CsmUnifiedHostAttributes) UnsetAgentlessVulnerabilityScanning() {
	o.AgentlessVulnerabilityScanning.Unset()
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *CsmUnifiedHostAttributes) GetCloudProvider() CsmCloudProvider {
	if o == nil || o.CloudProvider == nil {
		var ret CsmCloudProvider
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmUnifiedHostAttributes) GetCloudProviderOk() (*CsmCloudProvider, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasCloudProvider() bool {
	return o != nil && o.CloudProvider != nil
}

// SetCloudProvider gets a reference to the given CsmCloudProvider and assigns it to the CloudProvider field.
func (o *CsmUnifiedHostAttributes) SetCloudProvider(v CsmCloudProvider) {
	o.CloudProvider = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmUnifiedHostAttributes) GetClusterName() string {
	if o == nil || o.ClusterName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterName.Get()
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmUnifiedHostAttributes) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterName.Get(), o.ClusterName.IsSet()
}

// HasClusterName returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasClusterName() bool {
	return o != nil && o.ClusterName.IsSet()
}

// SetClusterName gets a reference to the given datadog.NullableString and assigns it to the ClusterName field.
func (o *CsmUnifiedHostAttributes) SetClusterName(v string) {
	o.ClusterName.Set(&v)
}

// SetClusterNameNil sets the value for ClusterName to be an explicit nil.
func (o *CsmUnifiedHostAttributes) SetClusterNameNil() {
	o.ClusterName.Set(nil)
}

// UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil.
func (o *CsmUnifiedHostAttributes) UnsetClusterName() {
	o.ClusterName.Unset()
}

// GetDatadogAgentKey returns the DatadogAgentKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmUnifiedHostAttributes) GetDatadogAgentKey() string {
	if o == nil || o.DatadogAgentKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.DatadogAgentKey.Get()
}

// GetDatadogAgentKeyOk returns a tuple with the DatadogAgentKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmUnifiedHostAttributes) GetDatadogAgentKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatadogAgentKey.Get(), o.DatadogAgentKey.IsSet()
}

// HasDatadogAgentKey returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasDatadogAgentKey() bool {
	return o != nil && o.DatadogAgentKey.IsSet()
}

// SetDatadogAgentKey gets a reference to the given datadog.NullableString and assigns it to the DatadogAgentKey field.
func (o *CsmUnifiedHostAttributes) SetDatadogAgentKey(v string) {
	o.DatadogAgentKey.Set(&v)
}

// SetDatadogAgentKeyNil sets the value for DatadogAgentKey to be an explicit nil.
func (o *CsmUnifiedHostAttributes) SetDatadogAgentKeyNil() {
	o.DatadogAgentKey.Set(nil)
}

// UnsetDatadogAgentKey ensures that no value is present for DatadogAgentKey, not even an explicit nil.
func (o *CsmUnifiedHostAttributes) UnsetDatadogAgentKey() {
	o.DatadogAgentKey.Unset()
}

// GetEnv returns the Env field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmUnifiedHostAttributes) GetEnv() []string {
	if o == nil || o.Env.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Env.Get()
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmUnifiedHostAttributes) GetEnvOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Env.Get(), o.Env.IsSet()
}

// HasEnv returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasEnv() bool {
	return o != nil && o.Env.IsSet()
}

// SetEnv gets a reference to the given datadog.NullableList[string] and assigns it to the Env field.
func (o *CsmUnifiedHostAttributes) SetEnv(v []string) {
	o.Env.Set(&v)
}

// SetEnvNil sets the value for Env to be an explicit nil.
func (o *CsmUnifiedHostAttributes) SetEnvNil() {
	o.Env.Set(nil)
}

// UnsetEnv ensures that no value is present for Env, not even an explicit nil.
func (o *CsmUnifiedHostAttributes) UnsetEnv() {
	o.Env.Unset()
}

// GetHostId returns the HostId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmUnifiedHostAttributes) GetHostId() int64 {
	if o == nil || o.HostId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.HostId.Get()
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmUnifiedHostAttributes) GetHostIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HostId.Get(), o.HostId.IsSet()
}

// HasHostId returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasHostId() bool {
	return o != nil && o.HostId.IsSet()
}

// SetHostId gets a reference to the given datadog.NullableInt64 and assigns it to the HostId field.
func (o *CsmUnifiedHostAttributes) SetHostId(v int64) {
	o.HostId.Set(&v)
}

// SetHostIdNil sets the value for HostId to be an explicit nil.
func (o *CsmUnifiedHostAttributes) SetHostIdNil() {
	o.HostId.Set(nil)
}

// UnsetHostId ensures that no value is present for HostId, not even an explicit nil.
func (o *CsmUnifiedHostAttributes) UnsetHostId() {
	o.HostId.Unset()
}

// GetInstallMethodTool returns the InstallMethodTool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmUnifiedHostAttributes) GetInstallMethodTool() string {
	if o == nil || o.InstallMethodTool.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstallMethodTool.Get()
}

// GetInstallMethodToolOk returns a tuple with the InstallMethodTool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmUnifiedHostAttributes) GetInstallMethodToolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstallMethodTool.Get(), o.InstallMethodTool.IsSet()
}

// HasInstallMethodTool returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasInstallMethodTool() bool {
	return o != nil && o.InstallMethodTool.IsSet()
}

// SetInstallMethodTool gets a reference to the given datadog.NullableString and assigns it to the InstallMethodTool field.
func (o *CsmUnifiedHostAttributes) SetInstallMethodTool(v string) {
	o.InstallMethodTool.Set(&v)
}

// SetInstallMethodToolNil sets the value for InstallMethodTool to be an explicit nil.
func (o *CsmUnifiedHostAttributes) SetInstallMethodToolNil() {
	o.InstallMethodTool.Set(nil)
}

// UnsetInstallMethodTool ensures that no value is present for InstallMethodTool, not even an explicit nil.
func (o *CsmUnifiedHostAttributes) UnsetInstallMethodTool() {
	o.InstallMethodTool.Unset()
}

// GetOs returns the Os field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmUnifiedHostAttributes) GetOs() string {
	if o == nil || o.Os.Get() == nil {
		var ret string
		return ret
	}
	return *o.Os.Get()
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmUnifiedHostAttributes) GetOsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Os.Get(), o.Os.IsSet()
}

// HasOs returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasOs() bool {
	return o != nil && o.Os.IsSet()
}

// SetOs gets a reference to the given datadog.NullableString and assigns it to the Os field.
func (o *CsmUnifiedHostAttributes) SetOs(v string) {
	o.Os.Set(&v)
}

// SetOsNil sets the value for Os to be an explicit nil.
func (o *CsmUnifiedHostAttributes) SetOsNil() {
	o.Os.Set(nil)
}

// UnsetOs ensures that no value is present for Os, not even an explicit nil.
func (o *CsmUnifiedHostAttributes) UnsetOs() {
	o.Os.Unset()
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *CsmUnifiedHostAttributes) GetResourceType() CsmAgentlessHostResourceType {
	if o == nil || o.ResourceType == nil {
		var ret CsmAgentlessHostResourceType
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmUnifiedHostAttributes) GetResourceTypeOk() (*CsmAgentlessHostResourceType, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *CsmUnifiedHostAttributes) HasResourceType() bool {
	return o != nil && o.ResourceType != nil
}

// SetResourceType gets a reference to the given CsmAgentlessHostResourceType and assigns it to the ResourceType field.
func (o *CsmUnifiedHostAttributes) SetResourceType(v CsmAgentlessHostResourceType) {
	o.ResourceType = &v
}

// GetSource returns the Source field value.
func (o *CsmUnifiedHostAttributes) GetSource() CsmUnifiedHostSource {
	if o == nil {
		var ret CsmUnifiedHostSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CsmUnifiedHostAttributes) GetSourceOk() (*CsmUnifiedHostSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *CsmUnifiedHostAttributes) SetSource(v CsmUnifiedHostSource) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CsmUnifiedHostAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId.IsSet() {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if o.AgentCsmVmContainersEnabled.IsSet() {
		toSerialize["agent_csm_vm_containers_enabled"] = o.AgentCsmVmContainersEnabled.Get()
	}
	if o.AgentCsmVmHostsEnabled.IsSet() {
		toSerialize["agent_csm_vm_hosts_enabled"] = o.AgentCsmVmHostsEnabled.Get()
	}
	if o.AgentCwsEnabled.IsSet() {
		toSerialize["agent_cws_enabled"] = o.AgentCwsEnabled.Get()
	}
	if o.AgentPostureManagement.IsSet() {
		toSerialize["agent_posture_management"] = o.AgentPostureManagement.Get()
	}
	if o.AgentVersion.IsSet() {
		toSerialize["agent_version"] = o.AgentVersion.Get()
	}
	if o.AgentlessPostureManagement.IsSet() {
		toSerialize["agentless_posture_management"] = o.AgentlessPostureManagement.Get()
	}
	if o.AgentlessVulnerabilityScanning.IsSet() {
		toSerialize["agentless_vulnerability_scanning"] = o.AgentlessVulnerabilityScanning.Get()
	}
	if o.CloudProvider != nil {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if o.ClusterName.IsSet() {
		toSerialize["cluster_name"] = o.ClusterName.Get()
	}
	if o.DatadogAgentKey.IsSet() {
		toSerialize["datadog_agent_key"] = o.DatadogAgentKey.Get()
	}
	if o.Env.IsSet() {
		toSerialize["env"] = o.Env.Get()
	}
	if o.HostId.IsSet() {
		toSerialize["host_id"] = o.HostId.Get()
	}
	if o.InstallMethodTool.IsSet() {
		toSerialize["install_method_tool"] = o.InstallMethodTool.Get()
	}
	if o.Os.IsSet() {
		toSerialize["os"] = o.Os.Get()
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CsmUnifiedHostAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId                      datadog.NullableString        `json:"account_id,omitempty"`
		AgentCsmVmContainersEnabled    datadog.NullableBool          `json:"agent_csm_vm_containers_enabled,omitempty"`
		AgentCsmVmHostsEnabled         datadog.NullableBool          `json:"agent_csm_vm_hosts_enabled,omitempty"`
		AgentCwsEnabled                datadog.NullableBool          `json:"agent_cws_enabled,omitempty"`
		AgentPostureManagement         datadog.NullableBool          `json:"agent_posture_management,omitempty"`
		AgentVersion                   datadog.NullableString        `json:"agent_version,omitempty"`
		AgentlessPostureManagement     datadog.NullableBool          `json:"agentless_posture_management,omitempty"`
		AgentlessVulnerabilityScanning datadog.NullableBool          `json:"agentless_vulnerability_scanning,omitempty"`
		CloudProvider                  *CsmCloudProvider             `json:"cloud_provider,omitempty"`
		ClusterName                    datadog.NullableString        `json:"cluster_name,omitempty"`
		DatadogAgentKey                datadog.NullableString        `json:"datadog_agent_key,omitempty"`
		Env                            datadog.NullableList[string]  `json:"env,omitempty"`
		HostId                         datadog.NullableInt64         `json:"host_id,omitempty"`
		InstallMethodTool              datadog.NullableString        `json:"install_method_tool,omitempty"`
		Os                             datadog.NullableString        `json:"os,omitempty"`
		ResourceType                   *CsmAgentlessHostResourceType `json:"resource_type,omitempty"`
		Source                         *CsmUnifiedHostSource         `json:"source"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "agent_csm_vm_containers_enabled", "agent_csm_vm_hosts_enabled", "agent_cws_enabled", "agent_posture_management", "agent_version", "agentless_posture_management", "agentless_vulnerability_scanning", "cloud_provider", "cluster_name", "datadog_agent_key", "env", "host_id", "install_method_tool", "os", "resource_type", "source"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountId = all.AccountId
	o.AgentCsmVmContainersEnabled = all.AgentCsmVmContainersEnabled
	o.AgentCsmVmHostsEnabled = all.AgentCsmVmHostsEnabled
	o.AgentCwsEnabled = all.AgentCwsEnabled
	o.AgentPostureManagement = all.AgentPostureManagement
	o.AgentVersion = all.AgentVersion
	o.AgentlessPostureManagement = all.AgentlessPostureManagement
	o.AgentlessVulnerabilityScanning = all.AgentlessVulnerabilityScanning
	if all.CloudProvider != nil && !all.CloudProvider.IsValid() {
		hasInvalidField = true
	} else {
		o.CloudProvider = all.CloudProvider
	}
	o.ClusterName = all.ClusterName
	o.DatadogAgentKey = all.DatadogAgentKey
	o.Env = all.Env
	o.HostId = all.HostId
	o.InstallMethodTool = all.InstallMethodTool
	o.Os = all.Os
	if all.ResourceType != nil && !all.ResourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ResourceType = all.ResourceType
	}
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
