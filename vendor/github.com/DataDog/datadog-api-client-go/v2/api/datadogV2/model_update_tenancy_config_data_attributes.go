// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateTenancyConfigDataAttributes Attributes for updating an existing OCI tenancy integration configuration, including optional credentials, region settings, and collection options.
type UpdateTenancyConfigDataAttributes struct {
	// OCI API signing key credentials used to update the Datadog integration's authentication with the OCI tenancy.
	AuthCredentials *UpdateTenancyConfigDataAttributesAuthCredentials `json:"auth_credentials,omitempty"`
	// Whether cost data collection from OCI is enabled for the tenancy.
	CostCollectionEnabled *bool `json:"cost_collection_enabled,omitempty"`
	// The home region of the OCI tenancy (for example, us-ashburn-1).
	HomeRegion *string `json:"home_region,omitempty"`
	// Log collection configuration for updating an OCI tenancy, controlling which compartments and services have log collection enabled.
	LogsConfig *UpdateTenancyConfigDataAttributesLogsConfig `json:"logs_config,omitempty"`
	// Metrics collection configuration for updating an OCI tenancy, controlling which compartments and services are included or excluded.
	MetricsConfig *UpdateTenancyConfigDataAttributesMetricsConfig `json:"metrics_config,omitempty"`
	// Region configuration for updating an OCI tenancy, specifying which regions are available, enabled, or disabled for data collection.
	RegionsConfig *UpdateTenancyConfigDataAttributesRegionsConfig `json:"regions_config,omitempty"`
	// Whether resource collection from OCI is enabled for the tenancy.
	ResourceCollectionEnabled *bool `json:"resource_collection_enabled,omitempty"`
	// The OCID of the OCI user used by the Datadog integration for authentication.
	UserOcid *string `json:"user_ocid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateTenancyConfigDataAttributes instantiates a new UpdateTenancyConfigDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateTenancyConfigDataAttributes() *UpdateTenancyConfigDataAttributes {
	this := UpdateTenancyConfigDataAttributes{}
	return &this
}

// NewUpdateTenancyConfigDataAttributesWithDefaults instantiates a new UpdateTenancyConfigDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateTenancyConfigDataAttributesWithDefaults() *UpdateTenancyConfigDataAttributes {
	this := UpdateTenancyConfigDataAttributes{}
	return &this
}

// GetAuthCredentials returns the AuthCredentials field value if set, zero value otherwise.
func (o *UpdateTenancyConfigDataAttributes) GetAuthCredentials() UpdateTenancyConfigDataAttributesAuthCredentials {
	if o == nil || o.AuthCredentials == nil {
		var ret UpdateTenancyConfigDataAttributesAuthCredentials
		return ret
	}
	return *o.AuthCredentials
}

// GetAuthCredentialsOk returns a tuple with the AuthCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigDataAttributes) GetAuthCredentialsOk() (*UpdateTenancyConfigDataAttributesAuthCredentials, bool) {
	if o == nil || o.AuthCredentials == nil {
		return nil, false
	}
	return o.AuthCredentials, true
}

// HasAuthCredentials returns a boolean if a field has been set.
func (o *UpdateTenancyConfigDataAttributes) HasAuthCredentials() bool {
	return o != nil && o.AuthCredentials != nil
}

// SetAuthCredentials gets a reference to the given UpdateTenancyConfigDataAttributesAuthCredentials and assigns it to the AuthCredentials field.
func (o *UpdateTenancyConfigDataAttributes) SetAuthCredentials(v UpdateTenancyConfigDataAttributesAuthCredentials) {
	o.AuthCredentials = &v
}

// GetCostCollectionEnabled returns the CostCollectionEnabled field value if set, zero value otherwise.
func (o *UpdateTenancyConfigDataAttributes) GetCostCollectionEnabled() bool {
	if o == nil || o.CostCollectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CostCollectionEnabled
}

// GetCostCollectionEnabledOk returns a tuple with the CostCollectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigDataAttributes) GetCostCollectionEnabledOk() (*bool, bool) {
	if o == nil || o.CostCollectionEnabled == nil {
		return nil, false
	}
	return o.CostCollectionEnabled, true
}

// HasCostCollectionEnabled returns a boolean if a field has been set.
func (o *UpdateTenancyConfigDataAttributes) HasCostCollectionEnabled() bool {
	return o != nil && o.CostCollectionEnabled != nil
}

// SetCostCollectionEnabled gets a reference to the given bool and assigns it to the CostCollectionEnabled field.
func (o *UpdateTenancyConfigDataAttributes) SetCostCollectionEnabled(v bool) {
	o.CostCollectionEnabled = &v
}

// GetHomeRegion returns the HomeRegion field value if set, zero value otherwise.
func (o *UpdateTenancyConfigDataAttributes) GetHomeRegion() string {
	if o == nil || o.HomeRegion == nil {
		var ret string
		return ret
	}
	return *o.HomeRegion
}

// GetHomeRegionOk returns a tuple with the HomeRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigDataAttributes) GetHomeRegionOk() (*string, bool) {
	if o == nil || o.HomeRegion == nil {
		return nil, false
	}
	return o.HomeRegion, true
}

// HasHomeRegion returns a boolean if a field has been set.
func (o *UpdateTenancyConfigDataAttributes) HasHomeRegion() bool {
	return o != nil && o.HomeRegion != nil
}

// SetHomeRegion gets a reference to the given string and assigns it to the HomeRegion field.
func (o *UpdateTenancyConfigDataAttributes) SetHomeRegion(v string) {
	o.HomeRegion = &v
}

// GetLogsConfig returns the LogsConfig field value if set, zero value otherwise.
func (o *UpdateTenancyConfigDataAttributes) GetLogsConfig() UpdateTenancyConfigDataAttributesLogsConfig {
	if o == nil || o.LogsConfig == nil {
		var ret UpdateTenancyConfigDataAttributesLogsConfig
		return ret
	}
	return *o.LogsConfig
}

// GetLogsConfigOk returns a tuple with the LogsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigDataAttributes) GetLogsConfigOk() (*UpdateTenancyConfigDataAttributesLogsConfig, bool) {
	if o == nil || o.LogsConfig == nil {
		return nil, false
	}
	return o.LogsConfig, true
}

// HasLogsConfig returns a boolean if a field has been set.
func (o *UpdateTenancyConfigDataAttributes) HasLogsConfig() bool {
	return o != nil && o.LogsConfig != nil
}

// SetLogsConfig gets a reference to the given UpdateTenancyConfigDataAttributesLogsConfig and assigns it to the LogsConfig field.
func (o *UpdateTenancyConfigDataAttributes) SetLogsConfig(v UpdateTenancyConfigDataAttributesLogsConfig) {
	o.LogsConfig = &v
}

// GetMetricsConfig returns the MetricsConfig field value if set, zero value otherwise.
func (o *UpdateTenancyConfigDataAttributes) GetMetricsConfig() UpdateTenancyConfigDataAttributesMetricsConfig {
	if o == nil || o.MetricsConfig == nil {
		var ret UpdateTenancyConfigDataAttributesMetricsConfig
		return ret
	}
	return *o.MetricsConfig
}

// GetMetricsConfigOk returns a tuple with the MetricsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigDataAttributes) GetMetricsConfigOk() (*UpdateTenancyConfigDataAttributesMetricsConfig, bool) {
	if o == nil || o.MetricsConfig == nil {
		return nil, false
	}
	return o.MetricsConfig, true
}

// HasMetricsConfig returns a boolean if a field has been set.
func (o *UpdateTenancyConfigDataAttributes) HasMetricsConfig() bool {
	return o != nil && o.MetricsConfig != nil
}

// SetMetricsConfig gets a reference to the given UpdateTenancyConfigDataAttributesMetricsConfig and assigns it to the MetricsConfig field.
func (o *UpdateTenancyConfigDataAttributes) SetMetricsConfig(v UpdateTenancyConfigDataAttributesMetricsConfig) {
	o.MetricsConfig = &v
}

// GetRegionsConfig returns the RegionsConfig field value if set, zero value otherwise.
func (o *UpdateTenancyConfigDataAttributes) GetRegionsConfig() UpdateTenancyConfigDataAttributesRegionsConfig {
	if o == nil || o.RegionsConfig == nil {
		var ret UpdateTenancyConfigDataAttributesRegionsConfig
		return ret
	}
	return *o.RegionsConfig
}

// GetRegionsConfigOk returns a tuple with the RegionsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigDataAttributes) GetRegionsConfigOk() (*UpdateTenancyConfigDataAttributesRegionsConfig, bool) {
	if o == nil || o.RegionsConfig == nil {
		return nil, false
	}
	return o.RegionsConfig, true
}

// HasRegionsConfig returns a boolean if a field has been set.
func (o *UpdateTenancyConfigDataAttributes) HasRegionsConfig() bool {
	return o != nil && o.RegionsConfig != nil
}

// SetRegionsConfig gets a reference to the given UpdateTenancyConfigDataAttributesRegionsConfig and assigns it to the RegionsConfig field.
func (o *UpdateTenancyConfigDataAttributes) SetRegionsConfig(v UpdateTenancyConfigDataAttributesRegionsConfig) {
	o.RegionsConfig = &v
}

// GetResourceCollectionEnabled returns the ResourceCollectionEnabled field value if set, zero value otherwise.
func (o *UpdateTenancyConfigDataAttributes) GetResourceCollectionEnabled() bool {
	if o == nil || o.ResourceCollectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ResourceCollectionEnabled
}

// GetResourceCollectionEnabledOk returns a tuple with the ResourceCollectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigDataAttributes) GetResourceCollectionEnabledOk() (*bool, bool) {
	if o == nil || o.ResourceCollectionEnabled == nil {
		return nil, false
	}
	return o.ResourceCollectionEnabled, true
}

// HasResourceCollectionEnabled returns a boolean if a field has been set.
func (o *UpdateTenancyConfigDataAttributes) HasResourceCollectionEnabled() bool {
	return o != nil && o.ResourceCollectionEnabled != nil
}

// SetResourceCollectionEnabled gets a reference to the given bool and assigns it to the ResourceCollectionEnabled field.
func (o *UpdateTenancyConfigDataAttributes) SetResourceCollectionEnabled(v bool) {
	o.ResourceCollectionEnabled = &v
}

// GetUserOcid returns the UserOcid field value if set, zero value otherwise.
func (o *UpdateTenancyConfigDataAttributes) GetUserOcid() string {
	if o == nil || o.UserOcid == nil {
		var ret string
		return ret
	}
	return *o.UserOcid
}

// GetUserOcidOk returns a tuple with the UserOcid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigDataAttributes) GetUserOcidOk() (*string, bool) {
	if o == nil || o.UserOcid == nil {
		return nil, false
	}
	return o.UserOcid, true
}

// HasUserOcid returns a boolean if a field has been set.
func (o *UpdateTenancyConfigDataAttributes) HasUserOcid() bool {
	return o != nil && o.UserOcid != nil
}

// SetUserOcid gets a reference to the given string and assigns it to the UserOcid field.
func (o *UpdateTenancyConfigDataAttributes) SetUserOcid(v string) {
	o.UserOcid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateTenancyConfigDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AuthCredentials != nil {
		toSerialize["auth_credentials"] = o.AuthCredentials
	}
	if o.CostCollectionEnabled != nil {
		toSerialize["cost_collection_enabled"] = o.CostCollectionEnabled
	}
	if o.HomeRegion != nil {
		toSerialize["home_region"] = o.HomeRegion
	}
	if o.LogsConfig != nil {
		toSerialize["logs_config"] = o.LogsConfig
	}
	if o.MetricsConfig != nil {
		toSerialize["metrics_config"] = o.MetricsConfig
	}
	if o.RegionsConfig != nil {
		toSerialize["regions_config"] = o.RegionsConfig
	}
	if o.ResourceCollectionEnabled != nil {
		toSerialize["resource_collection_enabled"] = o.ResourceCollectionEnabled
	}
	if o.UserOcid != nil {
		toSerialize["user_ocid"] = o.UserOcid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateTenancyConfigDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthCredentials           *UpdateTenancyConfigDataAttributesAuthCredentials `json:"auth_credentials,omitempty"`
		CostCollectionEnabled     *bool                                             `json:"cost_collection_enabled,omitempty"`
		HomeRegion                *string                                           `json:"home_region,omitempty"`
		LogsConfig                *UpdateTenancyConfigDataAttributesLogsConfig      `json:"logs_config,omitempty"`
		MetricsConfig             *UpdateTenancyConfigDataAttributesMetricsConfig   `json:"metrics_config,omitempty"`
		RegionsConfig             *UpdateTenancyConfigDataAttributesRegionsConfig   `json:"regions_config,omitempty"`
		ResourceCollectionEnabled *bool                                             `json:"resource_collection_enabled,omitempty"`
		UserOcid                  *string                                           `json:"user_ocid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_credentials", "cost_collection_enabled", "home_region", "logs_config", "metrics_config", "regions_config", "resource_collection_enabled", "user_ocid"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AuthCredentials != nil && all.AuthCredentials.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AuthCredentials = all.AuthCredentials
	o.CostCollectionEnabled = all.CostCollectionEnabled
	o.HomeRegion = all.HomeRegion
	if all.LogsConfig != nil && all.LogsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogsConfig = all.LogsConfig
	if all.MetricsConfig != nil && all.MetricsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MetricsConfig = all.MetricsConfig
	if all.RegionsConfig != nil && all.RegionsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RegionsConfig = all.RegionsConfig
	o.ResourceCollectionEnabled = all.ResourceCollectionEnabled
	o.UserOcid = all.UserOcid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
