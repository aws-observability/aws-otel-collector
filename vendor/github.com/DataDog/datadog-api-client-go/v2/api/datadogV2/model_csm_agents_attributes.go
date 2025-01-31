// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmAgentsAttributes A CSM Agent returned by the API.
type CsmAgentsAttributes struct {
	// Version of the Datadog Agent.
	AgentVersion *string `json:"agent_version,omitempty"`
	// AWS Fargate details.
	AwsFargate *string `json:"aws_fargate,omitempty"`
	// List of cluster names associated with the Agent.
	ClusterName []string `json:"cluster_name,omitempty"`
	// Unique identifier for the Datadog Agent.
	DatadogAgent *string `json:"datadog_agent,omitempty"`
	// ARN of the ECS Fargate task.
	EcsFargateTaskArn *string `json:"ecs_fargate_task_arn,omitempty"`
	// List of environments associated with the Agent.
	Envs datadog.NullableList[string] `json:"envs,omitempty"`
	// ID of the host.
	HostId *int64 `json:"host_id,omitempty"`
	// Name of the host.
	Hostname *string `json:"hostname,omitempty"`
	// Version of the installer used for installing the Datadog Agent.
	InstallMethodInstallerVersion *string `json:"install_method_installer_version,omitempty"`
	// Tool used for installing the Datadog Agent.
	InstallMethodTool *string `json:"install_method_tool,omitempty"`
	// Indicates if CSM VM Containers is enabled.
	IsCsmVmContainersEnabled datadog.NullableBool `json:"is_csm_vm_containers_enabled,omitempty"`
	// Indicates if CSM VM Hosts is enabled.
	IsCsmVmHostsEnabled datadog.NullableBool `json:"is_csm_vm_hosts_enabled,omitempty"`
	// Indicates if CSPM is enabled.
	IsCspmEnabled datadog.NullableBool `json:"is_cspm_enabled,omitempty"`
	// Indicates if CWS is enabled.
	IsCwsEnabled datadog.NullableBool `json:"is_cws_enabled,omitempty"`
	// Indicates if CWS Remote Configuration is enabled.
	IsCwsRemoteConfigurationEnabled datadog.NullableBool `json:"is_cws_remote_configuration_enabled,omitempty"`
	// Indicates if Remote Configuration is enabled.
	IsRemoteConfigurationEnabled datadog.NullableBool `json:"is_remote_configuration_enabled,omitempty"`
	// Operating system of the host.
	Os *string `json:"os,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCsmAgentsAttributes instantiates a new CsmAgentsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCsmAgentsAttributes() *CsmAgentsAttributes {
	this := CsmAgentsAttributes{}
	return &this
}

// NewCsmAgentsAttributesWithDefaults instantiates a new CsmAgentsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCsmAgentsAttributesWithDefaults() *CsmAgentsAttributes {
	this := CsmAgentsAttributes{}
	return &this
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *CsmAgentsAttributes) GetAgentVersion() string {
	if o == nil || o.AgentVersion == nil {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmAgentsAttributes) GetAgentVersionOk() (*string, bool) {
	if o == nil || o.AgentVersion == nil {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasAgentVersion() bool {
	return o != nil && o.AgentVersion != nil
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *CsmAgentsAttributes) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetAwsFargate returns the AwsFargate field value if set, zero value otherwise.
func (o *CsmAgentsAttributes) GetAwsFargate() string {
	if o == nil || o.AwsFargate == nil {
		var ret string
		return ret
	}
	return *o.AwsFargate
}

// GetAwsFargateOk returns a tuple with the AwsFargate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmAgentsAttributes) GetAwsFargateOk() (*string, bool) {
	if o == nil || o.AwsFargate == nil {
		return nil, false
	}
	return o.AwsFargate, true
}

// HasAwsFargate returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasAwsFargate() bool {
	return o != nil && o.AwsFargate != nil
}

// SetAwsFargate gets a reference to the given string and assigns it to the AwsFargate field.
func (o *CsmAgentsAttributes) SetAwsFargate(v string) {
	o.AwsFargate = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *CsmAgentsAttributes) GetClusterName() []string {
	if o == nil || o.ClusterName == nil {
		var ret []string
		return ret
	}
	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmAgentsAttributes) GetClusterNameOk() (*[]string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given []string and assigns it to the ClusterName field.
func (o *CsmAgentsAttributes) SetClusterName(v []string) {
	o.ClusterName = v
}

// GetDatadogAgent returns the DatadogAgent field value if set, zero value otherwise.
func (o *CsmAgentsAttributes) GetDatadogAgent() string {
	if o == nil || o.DatadogAgent == nil {
		var ret string
		return ret
	}
	return *o.DatadogAgent
}

// GetDatadogAgentOk returns a tuple with the DatadogAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmAgentsAttributes) GetDatadogAgentOk() (*string, bool) {
	if o == nil || o.DatadogAgent == nil {
		return nil, false
	}
	return o.DatadogAgent, true
}

// HasDatadogAgent returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasDatadogAgent() bool {
	return o != nil && o.DatadogAgent != nil
}

// SetDatadogAgent gets a reference to the given string and assigns it to the DatadogAgent field.
func (o *CsmAgentsAttributes) SetDatadogAgent(v string) {
	o.DatadogAgent = &v
}

// GetEcsFargateTaskArn returns the EcsFargateTaskArn field value if set, zero value otherwise.
func (o *CsmAgentsAttributes) GetEcsFargateTaskArn() string {
	if o == nil || o.EcsFargateTaskArn == nil {
		var ret string
		return ret
	}
	return *o.EcsFargateTaskArn
}

// GetEcsFargateTaskArnOk returns a tuple with the EcsFargateTaskArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmAgentsAttributes) GetEcsFargateTaskArnOk() (*string, bool) {
	if o == nil || o.EcsFargateTaskArn == nil {
		return nil, false
	}
	return o.EcsFargateTaskArn, true
}

// HasEcsFargateTaskArn returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasEcsFargateTaskArn() bool {
	return o != nil && o.EcsFargateTaskArn != nil
}

// SetEcsFargateTaskArn gets a reference to the given string and assigns it to the EcsFargateTaskArn field.
func (o *CsmAgentsAttributes) SetEcsFargateTaskArn(v string) {
	o.EcsFargateTaskArn = &v
}

// GetEnvs returns the Envs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmAgentsAttributes) GetEnvs() []string {
	if o == nil || o.Envs.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Envs.Get()
}

// GetEnvsOk returns a tuple with the Envs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmAgentsAttributes) GetEnvsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Envs.Get(), o.Envs.IsSet()
}

// HasEnvs returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasEnvs() bool {
	return o != nil && o.Envs.IsSet()
}

// SetEnvs gets a reference to the given datadog.NullableList[string] and assigns it to the Envs field.
func (o *CsmAgentsAttributes) SetEnvs(v []string) {
	o.Envs.Set(&v)
}

// SetEnvsNil sets the value for Envs to be an explicit nil.
func (o *CsmAgentsAttributes) SetEnvsNil() {
	o.Envs.Set(nil)
}

// UnsetEnvs ensures that no value is present for Envs, not even an explicit nil.
func (o *CsmAgentsAttributes) UnsetEnvs() {
	o.Envs.Unset()
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *CsmAgentsAttributes) GetHostId() int64 {
	if o == nil || o.HostId == nil {
		var ret int64
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmAgentsAttributes) GetHostIdOk() (*int64, bool) {
	if o == nil || o.HostId == nil {
		return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasHostId() bool {
	return o != nil && o.HostId != nil
}

// SetHostId gets a reference to the given int64 and assigns it to the HostId field.
func (o *CsmAgentsAttributes) SetHostId(v int64) {
	o.HostId = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *CsmAgentsAttributes) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmAgentsAttributes) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *CsmAgentsAttributes) SetHostname(v string) {
	o.Hostname = &v
}

// GetInstallMethodInstallerVersion returns the InstallMethodInstallerVersion field value if set, zero value otherwise.
func (o *CsmAgentsAttributes) GetInstallMethodInstallerVersion() string {
	if o == nil || o.InstallMethodInstallerVersion == nil {
		var ret string
		return ret
	}
	return *o.InstallMethodInstallerVersion
}

// GetInstallMethodInstallerVersionOk returns a tuple with the InstallMethodInstallerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmAgentsAttributes) GetInstallMethodInstallerVersionOk() (*string, bool) {
	if o == nil || o.InstallMethodInstallerVersion == nil {
		return nil, false
	}
	return o.InstallMethodInstallerVersion, true
}

// HasInstallMethodInstallerVersion returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasInstallMethodInstallerVersion() bool {
	return o != nil && o.InstallMethodInstallerVersion != nil
}

// SetInstallMethodInstallerVersion gets a reference to the given string and assigns it to the InstallMethodInstallerVersion field.
func (o *CsmAgentsAttributes) SetInstallMethodInstallerVersion(v string) {
	o.InstallMethodInstallerVersion = &v
}

// GetInstallMethodTool returns the InstallMethodTool field value if set, zero value otherwise.
func (o *CsmAgentsAttributes) GetInstallMethodTool() string {
	if o == nil || o.InstallMethodTool == nil {
		var ret string
		return ret
	}
	return *o.InstallMethodTool
}

// GetInstallMethodToolOk returns a tuple with the InstallMethodTool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmAgentsAttributes) GetInstallMethodToolOk() (*string, bool) {
	if o == nil || o.InstallMethodTool == nil {
		return nil, false
	}
	return o.InstallMethodTool, true
}

// HasInstallMethodTool returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasInstallMethodTool() bool {
	return o != nil && o.InstallMethodTool != nil
}

// SetInstallMethodTool gets a reference to the given string and assigns it to the InstallMethodTool field.
func (o *CsmAgentsAttributes) SetInstallMethodTool(v string) {
	o.InstallMethodTool = &v
}

// GetIsCsmVmContainersEnabled returns the IsCsmVmContainersEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmAgentsAttributes) GetIsCsmVmContainersEnabled() bool {
	if o == nil || o.IsCsmVmContainersEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsCsmVmContainersEnabled.Get()
}

// GetIsCsmVmContainersEnabledOk returns a tuple with the IsCsmVmContainersEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmAgentsAttributes) GetIsCsmVmContainersEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCsmVmContainersEnabled.Get(), o.IsCsmVmContainersEnabled.IsSet()
}

// HasIsCsmVmContainersEnabled returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasIsCsmVmContainersEnabled() bool {
	return o != nil && o.IsCsmVmContainersEnabled.IsSet()
}

// SetIsCsmVmContainersEnabled gets a reference to the given datadog.NullableBool and assigns it to the IsCsmVmContainersEnabled field.
func (o *CsmAgentsAttributes) SetIsCsmVmContainersEnabled(v bool) {
	o.IsCsmVmContainersEnabled.Set(&v)
}

// SetIsCsmVmContainersEnabledNil sets the value for IsCsmVmContainersEnabled to be an explicit nil.
func (o *CsmAgentsAttributes) SetIsCsmVmContainersEnabledNil() {
	o.IsCsmVmContainersEnabled.Set(nil)
}

// UnsetIsCsmVmContainersEnabled ensures that no value is present for IsCsmVmContainersEnabled, not even an explicit nil.
func (o *CsmAgentsAttributes) UnsetIsCsmVmContainersEnabled() {
	o.IsCsmVmContainersEnabled.Unset()
}

// GetIsCsmVmHostsEnabled returns the IsCsmVmHostsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmAgentsAttributes) GetIsCsmVmHostsEnabled() bool {
	if o == nil || o.IsCsmVmHostsEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsCsmVmHostsEnabled.Get()
}

// GetIsCsmVmHostsEnabledOk returns a tuple with the IsCsmVmHostsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmAgentsAttributes) GetIsCsmVmHostsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCsmVmHostsEnabled.Get(), o.IsCsmVmHostsEnabled.IsSet()
}

// HasIsCsmVmHostsEnabled returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasIsCsmVmHostsEnabled() bool {
	return o != nil && o.IsCsmVmHostsEnabled.IsSet()
}

// SetIsCsmVmHostsEnabled gets a reference to the given datadog.NullableBool and assigns it to the IsCsmVmHostsEnabled field.
func (o *CsmAgentsAttributes) SetIsCsmVmHostsEnabled(v bool) {
	o.IsCsmVmHostsEnabled.Set(&v)
}

// SetIsCsmVmHostsEnabledNil sets the value for IsCsmVmHostsEnabled to be an explicit nil.
func (o *CsmAgentsAttributes) SetIsCsmVmHostsEnabledNil() {
	o.IsCsmVmHostsEnabled.Set(nil)
}

// UnsetIsCsmVmHostsEnabled ensures that no value is present for IsCsmVmHostsEnabled, not even an explicit nil.
func (o *CsmAgentsAttributes) UnsetIsCsmVmHostsEnabled() {
	o.IsCsmVmHostsEnabled.Unset()
}

// GetIsCspmEnabled returns the IsCspmEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmAgentsAttributes) GetIsCspmEnabled() bool {
	if o == nil || o.IsCspmEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsCspmEnabled.Get()
}

// GetIsCspmEnabledOk returns a tuple with the IsCspmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmAgentsAttributes) GetIsCspmEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCspmEnabled.Get(), o.IsCspmEnabled.IsSet()
}

// HasIsCspmEnabled returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasIsCspmEnabled() bool {
	return o != nil && o.IsCspmEnabled.IsSet()
}

// SetIsCspmEnabled gets a reference to the given datadog.NullableBool and assigns it to the IsCspmEnabled field.
func (o *CsmAgentsAttributes) SetIsCspmEnabled(v bool) {
	o.IsCspmEnabled.Set(&v)
}

// SetIsCspmEnabledNil sets the value for IsCspmEnabled to be an explicit nil.
func (o *CsmAgentsAttributes) SetIsCspmEnabledNil() {
	o.IsCspmEnabled.Set(nil)
}

// UnsetIsCspmEnabled ensures that no value is present for IsCspmEnabled, not even an explicit nil.
func (o *CsmAgentsAttributes) UnsetIsCspmEnabled() {
	o.IsCspmEnabled.Unset()
}

// GetIsCwsEnabled returns the IsCwsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmAgentsAttributes) GetIsCwsEnabled() bool {
	if o == nil || o.IsCwsEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsCwsEnabled.Get()
}

// GetIsCwsEnabledOk returns a tuple with the IsCwsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmAgentsAttributes) GetIsCwsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCwsEnabled.Get(), o.IsCwsEnabled.IsSet()
}

// HasIsCwsEnabled returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasIsCwsEnabled() bool {
	return o != nil && o.IsCwsEnabled.IsSet()
}

// SetIsCwsEnabled gets a reference to the given datadog.NullableBool and assigns it to the IsCwsEnabled field.
func (o *CsmAgentsAttributes) SetIsCwsEnabled(v bool) {
	o.IsCwsEnabled.Set(&v)
}

// SetIsCwsEnabledNil sets the value for IsCwsEnabled to be an explicit nil.
func (o *CsmAgentsAttributes) SetIsCwsEnabledNil() {
	o.IsCwsEnabled.Set(nil)
}

// UnsetIsCwsEnabled ensures that no value is present for IsCwsEnabled, not even an explicit nil.
func (o *CsmAgentsAttributes) UnsetIsCwsEnabled() {
	o.IsCwsEnabled.Unset()
}

// GetIsCwsRemoteConfigurationEnabled returns the IsCwsRemoteConfigurationEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmAgentsAttributes) GetIsCwsRemoteConfigurationEnabled() bool {
	if o == nil || o.IsCwsRemoteConfigurationEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsCwsRemoteConfigurationEnabled.Get()
}

// GetIsCwsRemoteConfigurationEnabledOk returns a tuple with the IsCwsRemoteConfigurationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmAgentsAttributes) GetIsCwsRemoteConfigurationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCwsRemoteConfigurationEnabled.Get(), o.IsCwsRemoteConfigurationEnabled.IsSet()
}

// HasIsCwsRemoteConfigurationEnabled returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasIsCwsRemoteConfigurationEnabled() bool {
	return o != nil && o.IsCwsRemoteConfigurationEnabled.IsSet()
}

// SetIsCwsRemoteConfigurationEnabled gets a reference to the given datadog.NullableBool and assigns it to the IsCwsRemoteConfigurationEnabled field.
func (o *CsmAgentsAttributes) SetIsCwsRemoteConfigurationEnabled(v bool) {
	o.IsCwsRemoteConfigurationEnabled.Set(&v)
}

// SetIsCwsRemoteConfigurationEnabledNil sets the value for IsCwsRemoteConfigurationEnabled to be an explicit nil.
func (o *CsmAgentsAttributes) SetIsCwsRemoteConfigurationEnabledNil() {
	o.IsCwsRemoteConfigurationEnabled.Set(nil)
}

// UnsetIsCwsRemoteConfigurationEnabled ensures that no value is present for IsCwsRemoteConfigurationEnabled, not even an explicit nil.
func (o *CsmAgentsAttributes) UnsetIsCwsRemoteConfigurationEnabled() {
	o.IsCwsRemoteConfigurationEnabled.Unset()
}

// GetIsRemoteConfigurationEnabled returns the IsRemoteConfigurationEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsmAgentsAttributes) GetIsRemoteConfigurationEnabled() bool {
	if o == nil || o.IsRemoteConfigurationEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsRemoteConfigurationEnabled.Get()
}

// GetIsRemoteConfigurationEnabledOk returns a tuple with the IsRemoteConfigurationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CsmAgentsAttributes) GetIsRemoteConfigurationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRemoteConfigurationEnabled.Get(), o.IsRemoteConfigurationEnabled.IsSet()
}

// HasIsRemoteConfigurationEnabled returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasIsRemoteConfigurationEnabled() bool {
	return o != nil && o.IsRemoteConfigurationEnabled.IsSet()
}

// SetIsRemoteConfigurationEnabled gets a reference to the given datadog.NullableBool and assigns it to the IsRemoteConfigurationEnabled field.
func (o *CsmAgentsAttributes) SetIsRemoteConfigurationEnabled(v bool) {
	o.IsRemoteConfigurationEnabled.Set(&v)
}

// SetIsRemoteConfigurationEnabledNil sets the value for IsRemoteConfigurationEnabled to be an explicit nil.
func (o *CsmAgentsAttributes) SetIsRemoteConfigurationEnabledNil() {
	o.IsRemoteConfigurationEnabled.Set(nil)
}

// UnsetIsRemoteConfigurationEnabled ensures that no value is present for IsRemoteConfigurationEnabled, not even an explicit nil.
func (o *CsmAgentsAttributes) UnsetIsRemoteConfigurationEnabled() {
	o.IsRemoteConfigurationEnabled.Unset()
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *CsmAgentsAttributes) GetOs() string {
	if o == nil || o.Os == nil {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmAgentsAttributes) GetOsOk() (*string, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *CsmAgentsAttributes) HasOs() bool {
	return o != nil && o.Os != nil
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *CsmAgentsAttributes) SetOs(v string) {
	o.Os = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CsmAgentsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AgentVersion != nil {
		toSerialize["agent_version"] = o.AgentVersion
	}
	if o.AwsFargate != nil {
		toSerialize["aws_fargate"] = o.AwsFargate
	}
	if o.ClusterName != nil {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if o.DatadogAgent != nil {
		toSerialize["datadog_agent"] = o.DatadogAgent
	}
	if o.EcsFargateTaskArn != nil {
		toSerialize["ecs_fargate_task_arn"] = o.EcsFargateTaskArn
	}
	if o.Envs.IsSet() {
		toSerialize["envs"] = o.Envs.Get()
	}
	if o.HostId != nil {
		toSerialize["host_id"] = o.HostId
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.InstallMethodInstallerVersion != nil {
		toSerialize["install_method_installer_version"] = o.InstallMethodInstallerVersion
	}
	if o.InstallMethodTool != nil {
		toSerialize["install_method_tool"] = o.InstallMethodTool
	}
	if o.IsCsmVmContainersEnabled.IsSet() {
		toSerialize["is_csm_vm_containers_enabled"] = o.IsCsmVmContainersEnabled.Get()
	}
	if o.IsCsmVmHostsEnabled.IsSet() {
		toSerialize["is_csm_vm_hosts_enabled"] = o.IsCsmVmHostsEnabled.Get()
	}
	if o.IsCspmEnabled.IsSet() {
		toSerialize["is_cspm_enabled"] = o.IsCspmEnabled.Get()
	}
	if o.IsCwsEnabled.IsSet() {
		toSerialize["is_cws_enabled"] = o.IsCwsEnabled.Get()
	}
	if o.IsCwsRemoteConfigurationEnabled.IsSet() {
		toSerialize["is_cws_remote_configuration_enabled"] = o.IsCwsRemoteConfigurationEnabled.Get()
	}
	if o.IsRemoteConfigurationEnabled.IsSet() {
		toSerialize["is_remote_configuration_enabled"] = o.IsRemoteConfigurationEnabled.Get()
	}
	if o.Os != nil {
		toSerialize["os"] = o.Os
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CsmAgentsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentVersion                    *string                      `json:"agent_version,omitempty"`
		AwsFargate                      *string                      `json:"aws_fargate,omitempty"`
		ClusterName                     []string                     `json:"cluster_name,omitempty"`
		DatadogAgent                    *string                      `json:"datadog_agent,omitempty"`
		EcsFargateTaskArn               *string                      `json:"ecs_fargate_task_arn,omitempty"`
		Envs                            datadog.NullableList[string] `json:"envs,omitempty"`
		HostId                          *int64                       `json:"host_id,omitempty"`
		Hostname                        *string                      `json:"hostname,omitempty"`
		InstallMethodInstallerVersion   *string                      `json:"install_method_installer_version,omitempty"`
		InstallMethodTool               *string                      `json:"install_method_tool,omitempty"`
		IsCsmVmContainersEnabled        datadog.NullableBool         `json:"is_csm_vm_containers_enabled,omitempty"`
		IsCsmVmHostsEnabled             datadog.NullableBool         `json:"is_csm_vm_hosts_enabled,omitempty"`
		IsCspmEnabled                   datadog.NullableBool         `json:"is_cspm_enabled,omitempty"`
		IsCwsEnabled                    datadog.NullableBool         `json:"is_cws_enabled,omitempty"`
		IsCwsRemoteConfigurationEnabled datadog.NullableBool         `json:"is_cws_remote_configuration_enabled,omitempty"`
		IsRemoteConfigurationEnabled    datadog.NullableBool         `json:"is_remote_configuration_enabled,omitempty"`
		Os                              *string                      `json:"os,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_version", "aws_fargate", "cluster_name", "datadog_agent", "ecs_fargate_task_arn", "envs", "host_id", "hostname", "install_method_installer_version", "install_method_tool", "is_csm_vm_containers_enabled", "is_csm_vm_hosts_enabled", "is_cspm_enabled", "is_cws_enabled", "is_cws_remote_configuration_enabled", "is_remote_configuration_enabled", "os"})
	} else {
		return err
	}
	o.AgentVersion = all.AgentVersion
	o.AwsFargate = all.AwsFargate
	o.ClusterName = all.ClusterName
	o.DatadogAgent = all.DatadogAgent
	o.EcsFargateTaskArn = all.EcsFargateTaskArn
	o.Envs = all.Envs
	o.HostId = all.HostId
	o.Hostname = all.Hostname
	o.InstallMethodInstallerVersion = all.InstallMethodInstallerVersion
	o.InstallMethodTool = all.InstallMethodTool
	o.IsCsmVmContainersEnabled = all.IsCsmVmContainersEnabled
	o.IsCsmVmHostsEnabled = all.IsCsmVmHostsEnabled
	o.IsCspmEnabled = all.IsCspmEnabled
	o.IsCwsEnabled = all.IsCwsEnabled
	o.IsCwsRemoteConfigurationEnabled = all.IsCwsRemoteConfigurationEnabled
	o.IsRemoteConfigurationEnabled = all.IsRemoteConfigurationEnabled
	o.Os = all.Os

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
