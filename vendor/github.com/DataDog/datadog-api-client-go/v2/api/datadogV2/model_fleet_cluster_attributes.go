// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetClusterAttributes Attributes of a Kubernetes cluster in the fleet.
type FleetClusterAttributes struct {
	// Datadog Agent versions running in the cluster.
	AgentVersions []string `json:"agent_versions,omitempty"`
	// API key names used by agents in the cluster.
	ApiKeyNames []string `json:"api_key_names,omitempty"`
	// API key UUIDs used by agents in the cluster.
	ApiKeyUuids []string `json:"api_key_uuids,omitempty"`
	// Cloud providers hosting the cluster.
	CloudProviders []string `json:"cloud_providers,omitempty"`
	// The name of the Kubernetes cluster.
	ClusterName *string `json:"cluster_name,omitempty"`
	// Datadog products enabled in the cluster.
	EnabledProducts []string `json:"enabled_products,omitempty"`
	// Environments associated with the cluster.
	Envs []string `json:"envs,omitempty"`
	// Timestamp when the cluster was first seen.
	FirstSeenAt *int64 `json:"first_seen_at,omitempty"`
	// The tool used to install agents in the cluster.
	InstallMethodTool *string `json:"install_method_tool,omitempty"`
	// Total number of nodes in the cluster.
	NodeCount *int64 `json:"node_count,omitempty"`
	// Node counts grouped by status.
	NodeCountByStatus map[string]int64 `json:"node_count_by_status,omitempty"`
	// Operating systems of nodes in the cluster.
	OperatingSystems []string `json:"operating_systems,omitempty"`
	// OpenTelemetry collector distributions in the cluster.
	OtelCollectorDistributions []string `json:"otel_collector_distributions,omitempty"`
	// OpenTelemetry collector versions in the cluster.
	OtelCollectorVersions []string `json:"otel_collector_versions,omitempty"`
	// Pod counts grouped by state.
	PodCountByState map[string]int64 `json:"pod_count_by_state,omitempty"`
	// Services running in the cluster.
	Services []string `json:"services,omitempty"`
	// Teams associated with the cluster.
	Teams []string `json:"teams,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetClusterAttributes instantiates a new FleetClusterAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetClusterAttributes() *FleetClusterAttributes {
	this := FleetClusterAttributes{}
	return &this
}

// NewFleetClusterAttributesWithDefaults instantiates a new FleetClusterAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetClusterAttributesWithDefaults() *FleetClusterAttributes {
	this := FleetClusterAttributes{}
	return &this
}

// GetAgentVersions returns the AgentVersions field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetAgentVersions() []string {
	if o == nil || o.AgentVersions == nil {
		var ret []string
		return ret
	}
	return o.AgentVersions
}

// GetAgentVersionsOk returns a tuple with the AgentVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetAgentVersionsOk() (*[]string, bool) {
	if o == nil || o.AgentVersions == nil {
		return nil, false
	}
	return &o.AgentVersions, true
}

// HasAgentVersions returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasAgentVersions() bool {
	return o != nil && o.AgentVersions != nil
}

// SetAgentVersions gets a reference to the given []string and assigns it to the AgentVersions field.
func (o *FleetClusterAttributes) SetAgentVersions(v []string) {
	o.AgentVersions = v
}

// GetApiKeyNames returns the ApiKeyNames field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetApiKeyNames() []string {
	if o == nil || o.ApiKeyNames == nil {
		var ret []string
		return ret
	}
	return o.ApiKeyNames
}

// GetApiKeyNamesOk returns a tuple with the ApiKeyNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetApiKeyNamesOk() (*[]string, bool) {
	if o == nil || o.ApiKeyNames == nil {
		return nil, false
	}
	return &o.ApiKeyNames, true
}

// HasApiKeyNames returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasApiKeyNames() bool {
	return o != nil && o.ApiKeyNames != nil
}

// SetApiKeyNames gets a reference to the given []string and assigns it to the ApiKeyNames field.
func (o *FleetClusterAttributes) SetApiKeyNames(v []string) {
	o.ApiKeyNames = v
}

// GetApiKeyUuids returns the ApiKeyUuids field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetApiKeyUuids() []string {
	if o == nil || o.ApiKeyUuids == nil {
		var ret []string
		return ret
	}
	return o.ApiKeyUuids
}

// GetApiKeyUuidsOk returns a tuple with the ApiKeyUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetApiKeyUuidsOk() (*[]string, bool) {
	if o == nil || o.ApiKeyUuids == nil {
		return nil, false
	}
	return &o.ApiKeyUuids, true
}

// HasApiKeyUuids returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasApiKeyUuids() bool {
	return o != nil && o.ApiKeyUuids != nil
}

// SetApiKeyUuids gets a reference to the given []string and assigns it to the ApiKeyUuids field.
func (o *FleetClusterAttributes) SetApiKeyUuids(v []string) {
	o.ApiKeyUuids = v
}

// GetCloudProviders returns the CloudProviders field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetCloudProviders() []string {
	if o == nil || o.CloudProviders == nil {
		var ret []string
		return ret
	}
	return o.CloudProviders
}

// GetCloudProvidersOk returns a tuple with the CloudProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetCloudProvidersOk() (*[]string, bool) {
	if o == nil || o.CloudProviders == nil {
		return nil, false
	}
	return &o.CloudProviders, true
}

// HasCloudProviders returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasCloudProviders() bool {
	return o != nil && o.CloudProviders != nil
}

// SetCloudProviders gets a reference to the given []string and assigns it to the CloudProviders field.
func (o *FleetClusterAttributes) SetCloudProviders(v []string) {
	o.CloudProviders = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *FleetClusterAttributes) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetEnabledProducts returns the EnabledProducts field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetEnabledProducts() []string {
	if o == nil || o.EnabledProducts == nil {
		var ret []string
		return ret
	}
	return o.EnabledProducts
}

// GetEnabledProductsOk returns a tuple with the EnabledProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetEnabledProductsOk() (*[]string, bool) {
	if o == nil || o.EnabledProducts == nil {
		return nil, false
	}
	return &o.EnabledProducts, true
}

// HasEnabledProducts returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasEnabledProducts() bool {
	return o != nil && o.EnabledProducts != nil
}

// SetEnabledProducts gets a reference to the given []string and assigns it to the EnabledProducts field.
func (o *FleetClusterAttributes) SetEnabledProducts(v []string) {
	o.EnabledProducts = v
}

// GetEnvs returns the Envs field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetEnvs() []string {
	if o == nil || o.Envs == nil {
		var ret []string
		return ret
	}
	return o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetEnvsOk() (*[]string, bool) {
	if o == nil || o.Envs == nil {
		return nil, false
	}
	return &o.Envs, true
}

// HasEnvs returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasEnvs() bool {
	return o != nil && o.Envs != nil
}

// SetEnvs gets a reference to the given []string and assigns it to the Envs field.
func (o *FleetClusterAttributes) SetEnvs(v []string) {
	o.Envs = v
}

// GetFirstSeenAt returns the FirstSeenAt field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetFirstSeenAt() int64 {
	if o == nil || o.FirstSeenAt == nil {
		var ret int64
		return ret
	}
	return *o.FirstSeenAt
}

// GetFirstSeenAtOk returns a tuple with the FirstSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetFirstSeenAtOk() (*int64, bool) {
	if o == nil || o.FirstSeenAt == nil {
		return nil, false
	}
	return o.FirstSeenAt, true
}

// HasFirstSeenAt returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasFirstSeenAt() bool {
	return o != nil && o.FirstSeenAt != nil
}

// SetFirstSeenAt gets a reference to the given int64 and assigns it to the FirstSeenAt field.
func (o *FleetClusterAttributes) SetFirstSeenAt(v int64) {
	o.FirstSeenAt = &v
}

// GetInstallMethodTool returns the InstallMethodTool field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetInstallMethodTool() string {
	if o == nil || o.InstallMethodTool == nil {
		var ret string
		return ret
	}
	return *o.InstallMethodTool
}

// GetInstallMethodToolOk returns a tuple with the InstallMethodTool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetInstallMethodToolOk() (*string, bool) {
	if o == nil || o.InstallMethodTool == nil {
		return nil, false
	}
	return o.InstallMethodTool, true
}

// HasInstallMethodTool returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasInstallMethodTool() bool {
	return o != nil && o.InstallMethodTool != nil
}

// SetInstallMethodTool gets a reference to the given string and assigns it to the InstallMethodTool field.
func (o *FleetClusterAttributes) SetInstallMethodTool(v string) {
	o.InstallMethodTool = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetNodeCount() int64 {
	if o == nil || o.NodeCount == nil {
		var ret int64
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetNodeCountOk() (*int64, bool) {
	if o == nil || o.NodeCount == nil {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasNodeCount() bool {
	return o != nil && o.NodeCount != nil
}

// SetNodeCount gets a reference to the given int64 and assigns it to the NodeCount field.
func (o *FleetClusterAttributes) SetNodeCount(v int64) {
	o.NodeCount = &v
}

// GetNodeCountByStatus returns the NodeCountByStatus field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetNodeCountByStatus() map[string]int64 {
	if o == nil || o.NodeCountByStatus == nil {
		var ret map[string]int64
		return ret
	}
	return o.NodeCountByStatus
}

// GetNodeCountByStatusOk returns a tuple with the NodeCountByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetNodeCountByStatusOk() (*map[string]int64, bool) {
	if o == nil || o.NodeCountByStatus == nil {
		return nil, false
	}
	return &o.NodeCountByStatus, true
}

// HasNodeCountByStatus returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasNodeCountByStatus() bool {
	return o != nil && o.NodeCountByStatus != nil
}

// SetNodeCountByStatus gets a reference to the given map[string]int64 and assigns it to the NodeCountByStatus field.
func (o *FleetClusterAttributes) SetNodeCountByStatus(v map[string]int64) {
	o.NodeCountByStatus = v
}

// GetOperatingSystems returns the OperatingSystems field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetOperatingSystems() []string {
	if o == nil || o.OperatingSystems == nil {
		var ret []string
		return ret
	}
	return o.OperatingSystems
}

// GetOperatingSystemsOk returns a tuple with the OperatingSystems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetOperatingSystemsOk() (*[]string, bool) {
	if o == nil || o.OperatingSystems == nil {
		return nil, false
	}
	return &o.OperatingSystems, true
}

// HasOperatingSystems returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasOperatingSystems() bool {
	return o != nil && o.OperatingSystems != nil
}

// SetOperatingSystems gets a reference to the given []string and assigns it to the OperatingSystems field.
func (o *FleetClusterAttributes) SetOperatingSystems(v []string) {
	o.OperatingSystems = v
}

// GetOtelCollectorDistributions returns the OtelCollectorDistributions field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetOtelCollectorDistributions() []string {
	if o == nil || o.OtelCollectorDistributions == nil {
		var ret []string
		return ret
	}
	return o.OtelCollectorDistributions
}

// GetOtelCollectorDistributionsOk returns a tuple with the OtelCollectorDistributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetOtelCollectorDistributionsOk() (*[]string, bool) {
	if o == nil || o.OtelCollectorDistributions == nil {
		return nil, false
	}
	return &o.OtelCollectorDistributions, true
}

// HasOtelCollectorDistributions returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasOtelCollectorDistributions() bool {
	return o != nil && o.OtelCollectorDistributions != nil
}

// SetOtelCollectorDistributions gets a reference to the given []string and assigns it to the OtelCollectorDistributions field.
func (o *FleetClusterAttributes) SetOtelCollectorDistributions(v []string) {
	o.OtelCollectorDistributions = v
}

// GetOtelCollectorVersions returns the OtelCollectorVersions field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetOtelCollectorVersions() []string {
	if o == nil || o.OtelCollectorVersions == nil {
		var ret []string
		return ret
	}
	return o.OtelCollectorVersions
}

// GetOtelCollectorVersionsOk returns a tuple with the OtelCollectorVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetOtelCollectorVersionsOk() (*[]string, bool) {
	if o == nil || o.OtelCollectorVersions == nil {
		return nil, false
	}
	return &o.OtelCollectorVersions, true
}

// HasOtelCollectorVersions returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasOtelCollectorVersions() bool {
	return o != nil && o.OtelCollectorVersions != nil
}

// SetOtelCollectorVersions gets a reference to the given []string and assigns it to the OtelCollectorVersions field.
func (o *FleetClusterAttributes) SetOtelCollectorVersions(v []string) {
	o.OtelCollectorVersions = v
}

// GetPodCountByState returns the PodCountByState field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetPodCountByState() map[string]int64 {
	if o == nil || o.PodCountByState == nil {
		var ret map[string]int64
		return ret
	}
	return o.PodCountByState
}

// GetPodCountByStateOk returns a tuple with the PodCountByState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetPodCountByStateOk() (*map[string]int64, bool) {
	if o == nil || o.PodCountByState == nil {
		return nil, false
	}
	return &o.PodCountByState, true
}

// HasPodCountByState returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasPodCountByState() bool {
	return o != nil && o.PodCountByState != nil
}

// SetPodCountByState gets a reference to the given map[string]int64 and assigns it to the PodCountByState field.
func (o *FleetClusterAttributes) SetPodCountByState(v map[string]int64) {
	o.PodCountByState = v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasServices() bool {
	return o != nil && o.Services != nil
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *FleetClusterAttributes) SetServices(v []string) {
	o.Services = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *FleetClusterAttributes) GetTeams() []string {
	if o == nil || o.Teams == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClusterAttributes) GetTeamsOk() (*[]string, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return &o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *FleetClusterAttributes) HasTeams() bool {
	return o != nil && o.Teams != nil
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *FleetClusterAttributes) SetTeams(v []string) {
	o.Teams = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetClusterAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AgentVersions != nil {
		toSerialize["agent_versions"] = o.AgentVersions
	}
	if o.ApiKeyNames != nil {
		toSerialize["api_key_names"] = o.ApiKeyNames
	}
	if o.ApiKeyUuids != nil {
		toSerialize["api_key_uuids"] = o.ApiKeyUuids
	}
	if o.CloudProviders != nil {
		toSerialize["cloud_providers"] = o.CloudProviders
	}
	if o.ClusterName != nil {
		toSerialize["cluster_name"] = o.ClusterName
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
	if o.InstallMethodTool != nil {
		toSerialize["install_method_tool"] = o.InstallMethodTool
	}
	if o.NodeCount != nil {
		toSerialize["node_count"] = o.NodeCount
	}
	if o.NodeCountByStatus != nil {
		toSerialize["node_count_by_status"] = o.NodeCountByStatus
	}
	if o.OperatingSystems != nil {
		toSerialize["operating_systems"] = o.OperatingSystems
	}
	if o.OtelCollectorDistributions != nil {
		toSerialize["otel_collector_distributions"] = o.OtelCollectorDistributions
	}
	if o.OtelCollectorVersions != nil {
		toSerialize["otel_collector_versions"] = o.OtelCollectorVersions
	}
	if o.PodCountByState != nil {
		toSerialize["pod_count_by_state"] = o.PodCountByState
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetClusterAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentVersions              []string         `json:"agent_versions,omitempty"`
		ApiKeyNames                []string         `json:"api_key_names,omitempty"`
		ApiKeyUuids                []string         `json:"api_key_uuids,omitempty"`
		CloudProviders             []string         `json:"cloud_providers,omitempty"`
		ClusterName                *string          `json:"cluster_name,omitempty"`
		EnabledProducts            []string         `json:"enabled_products,omitempty"`
		Envs                       []string         `json:"envs,omitempty"`
		FirstSeenAt                *int64           `json:"first_seen_at,omitempty"`
		InstallMethodTool          *string          `json:"install_method_tool,omitempty"`
		NodeCount                  *int64           `json:"node_count,omitempty"`
		NodeCountByStatus          map[string]int64 `json:"node_count_by_status,omitempty"`
		OperatingSystems           []string         `json:"operating_systems,omitempty"`
		OtelCollectorDistributions []string         `json:"otel_collector_distributions,omitempty"`
		OtelCollectorVersions      []string         `json:"otel_collector_versions,omitempty"`
		PodCountByState            map[string]int64 `json:"pod_count_by_state,omitempty"`
		Services                   []string         `json:"services,omitempty"`
		Teams                      []string         `json:"teams,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_versions", "api_key_names", "api_key_uuids", "cloud_providers", "cluster_name", "enabled_products", "envs", "first_seen_at", "install_method_tool", "node_count", "node_count_by_status", "operating_systems", "otel_collector_distributions", "otel_collector_versions", "pod_count_by_state", "services", "teams"})
	} else {
		return err
	}
	o.AgentVersions = all.AgentVersions
	o.ApiKeyNames = all.ApiKeyNames
	o.ApiKeyUuids = all.ApiKeyUuids
	o.CloudProviders = all.CloudProviders
	o.ClusterName = all.ClusterName
	o.EnabledProducts = all.EnabledProducts
	o.Envs = all.Envs
	o.FirstSeenAt = all.FirstSeenAt
	o.InstallMethodTool = all.InstallMethodTool
	o.NodeCount = all.NodeCount
	o.NodeCountByStatus = all.NodeCountByStatus
	o.OperatingSystems = all.OperatingSystems
	o.OtelCollectorDistributions = all.OtelCollectorDistributions
	o.OtelCollectorVersions = all.OtelCollectorVersions
	o.PodCountByState = all.PodCountByState
	o.Services = all.Services
	o.Teams = all.Teams

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
