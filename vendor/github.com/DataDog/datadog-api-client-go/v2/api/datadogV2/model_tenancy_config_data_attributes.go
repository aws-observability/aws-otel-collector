// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TenancyConfigDataAttributes Attributes of an OCI tenancy integration configuration, including authentication details, region settings, and collection options.
type TenancyConfigDataAttributes struct {
	// The identifier of the billing plan associated with the OCI tenancy.
	BillingPlanId *int32 `json:"billing_plan_id,omitempty"`
	// Version number of the integration the tenancy is integrated with
	ConfigVersion *int64 `json:"config_version,omitempty"`
	// Whether cost data collection from OCI is enabled for the tenancy.
	CostCollectionEnabled *bool `json:"cost_collection_enabled,omitempty"`
	// The OCID of the OCI compartment used by the Datadog integration stack.
	DdCompartmentId *string `json:"dd_compartment_id,omitempty"`
	// The OCID of the OCI Resource Manager stack used by the Datadog integration.
	DdStackId *string `json:"dd_stack_id,omitempty"`
	// The home region of the OCI tenancy (for example, us-ashburn-1).
	HomeRegion *string `json:"home_region,omitempty"`
	// Log collection configuration for an OCI tenancy, indicating which compartments and services have log collection enabled.
	LogsConfig *TenancyConfigDataAttributesLogsConfig `json:"logs_config,omitempty"`
	// Metrics collection configuration for an OCI tenancy, indicating which compartments and services are included or excluded.
	MetricsConfig *TenancyConfigDataAttributesMetricsConfig `json:"metrics_config,omitempty"`
	// The name of the parent OCI tenancy, if applicable.
	ParentTenancyName *string `json:"parent_tenancy_name,omitempty"`
	// Region configuration for an OCI tenancy, indicating which regions are available, enabled, or disabled for data collection.
	RegionsConfig *TenancyConfigDataAttributesRegionsConfig `json:"regions_config,omitempty"`
	// Whether resource collection from OCI is enabled for the tenancy.
	ResourceCollectionEnabled *bool `json:"resource_collection_enabled,omitempty"`
	// The human-readable name of the OCI tenancy.
	TenancyName *string `json:"tenancy_name,omitempty"`
	// The OCID of the OCI user used by the Datadog integration for authentication.
	UserOcid *string `json:"user_ocid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTenancyConfigDataAttributes instantiates a new TenancyConfigDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTenancyConfigDataAttributes() *TenancyConfigDataAttributes {
	this := TenancyConfigDataAttributes{}
	return &this
}

// NewTenancyConfigDataAttributesWithDefaults instantiates a new TenancyConfigDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTenancyConfigDataAttributesWithDefaults() *TenancyConfigDataAttributes {
	this := TenancyConfigDataAttributes{}
	return &this
}

// GetBillingPlanId returns the BillingPlanId field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetBillingPlanId() int32 {
	if o == nil || o.BillingPlanId == nil {
		var ret int32
		return ret
	}
	return *o.BillingPlanId
}

// GetBillingPlanIdOk returns a tuple with the BillingPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetBillingPlanIdOk() (*int32, bool) {
	if o == nil || o.BillingPlanId == nil {
		return nil, false
	}
	return o.BillingPlanId, true
}

// HasBillingPlanId returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasBillingPlanId() bool {
	return o != nil && o.BillingPlanId != nil
}

// SetBillingPlanId gets a reference to the given int32 and assigns it to the BillingPlanId field.
func (o *TenancyConfigDataAttributes) SetBillingPlanId(v int32) {
	o.BillingPlanId = &v
}

// GetConfigVersion returns the ConfigVersion field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetConfigVersion() int64 {
	if o == nil || o.ConfigVersion == nil {
		var ret int64
		return ret
	}
	return *o.ConfigVersion
}

// GetConfigVersionOk returns a tuple with the ConfigVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetConfigVersionOk() (*int64, bool) {
	if o == nil || o.ConfigVersion == nil {
		return nil, false
	}
	return o.ConfigVersion, true
}

// HasConfigVersion returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasConfigVersion() bool {
	return o != nil && o.ConfigVersion != nil
}

// SetConfigVersion gets a reference to the given int64 and assigns it to the ConfigVersion field.
func (o *TenancyConfigDataAttributes) SetConfigVersion(v int64) {
	o.ConfigVersion = &v
}

// GetCostCollectionEnabled returns the CostCollectionEnabled field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetCostCollectionEnabled() bool {
	if o == nil || o.CostCollectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CostCollectionEnabled
}

// GetCostCollectionEnabledOk returns a tuple with the CostCollectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetCostCollectionEnabledOk() (*bool, bool) {
	if o == nil || o.CostCollectionEnabled == nil {
		return nil, false
	}
	return o.CostCollectionEnabled, true
}

// HasCostCollectionEnabled returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasCostCollectionEnabled() bool {
	return o != nil && o.CostCollectionEnabled != nil
}

// SetCostCollectionEnabled gets a reference to the given bool and assigns it to the CostCollectionEnabled field.
func (o *TenancyConfigDataAttributes) SetCostCollectionEnabled(v bool) {
	o.CostCollectionEnabled = &v
}

// GetDdCompartmentId returns the DdCompartmentId field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetDdCompartmentId() string {
	if o == nil || o.DdCompartmentId == nil {
		var ret string
		return ret
	}
	return *o.DdCompartmentId
}

// GetDdCompartmentIdOk returns a tuple with the DdCompartmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetDdCompartmentIdOk() (*string, bool) {
	if o == nil || o.DdCompartmentId == nil {
		return nil, false
	}
	return o.DdCompartmentId, true
}

// HasDdCompartmentId returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasDdCompartmentId() bool {
	return o != nil && o.DdCompartmentId != nil
}

// SetDdCompartmentId gets a reference to the given string and assigns it to the DdCompartmentId field.
func (o *TenancyConfigDataAttributes) SetDdCompartmentId(v string) {
	o.DdCompartmentId = &v
}

// GetDdStackId returns the DdStackId field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetDdStackId() string {
	if o == nil || o.DdStackId == nil {
		var ret string
		return ret
	}
	return *o.DdStackId
}

// GetDdStackIdOk returns a tuple with the DdStackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetDdStackIdOk() (*string, bool) {
	if o == nil || o.DdStackId == nil {
		return nil, false
	}
	return o.DdStackId, true
}

// HasDdStackId returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasDdStackId() bool {
	return o != nil && o.DdStackId != nil
}

// SetDdStackId gets a reference to the given string and assigns it to the DdStackId field.
func (o *TenancyConfigDataAttributes) SetDdStackId(v string) {
	o.DdStackId = &v
}

// GetHomeRegion returns the HomeRegion field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetHomeRegion() string {
	if o == nil || o.HomeRegion == nil {
		var ret string
		return ret
	}
	return *o.HomeRegion
}

// GetHomeRegionOk returns a tuple with the HomeRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetHomeRegionOk() (*string, bool) {
	if o == nil || o.HomeRegion == nil {
		return nil, false
	}
	return o.HomeRegion, true
}

// HasHomeRegion returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasHomeRegion() bool {
	return o != nil && o.HomeRegion != nil
}

// SetHomeRegion gets a reference to the given string and assigns it to the HomeRegion field.
func (o *TenancyConfigDataAttributes) SetHomeRegion(v string) {
	o.HomeRegion = &v
}

// GetLogsConfig returns the LogsConfig field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetLogsConfig() TenancyConfigDataAttributesLogsConfig {
	if o == nil || o.LogsConfig == nil {
		var ret TenancyConfigDataAttributesLogsConfig
		return ret
	}
	return *o.LogsConfig
}

// GetLogsConfigOk returns a tuple with the LogsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetLogsConfigOk() (*TenancyConfigDataAttributesLogsConfig, bool) {
	if o == nil || o.LogsConfig == nil {
		return nil, false
	}
	return o.LogsConfig, true
}

// HasLogsConfig returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasLogsConfig() bool {
	return o != nil && o.LogsConfig != nil
}

// SetLogsConfig gets a reference to the given TenancyConfigDataAttributesLogsConfig and assigns it to the LogsConfig field.
func (o *TenancyConfigDataAttributes) SetLogsConfig(v TenancyConfigDataAttributesLogsConfig) {
	o.LogsConfig = &v
}

// GetMetricsConfig returns the MetricsConfig field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetMetricsConfig() TenancyConfigDataAttributesMetricsConfig {
	if o == nil || o.MetricsConfig == nil {
		var ret TenancyConfigDataAttributesMetricsConfig
		return ret
	}
	return *o.MetricsConfig
}

// GetMetricsConfigOk returns a tuple with the MetricsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetMetricsConfigOk() (*TenancyConfigDataAttributesMetricsConfig, bool) {
	if o == nil || o.MetricsConfig == nil {
		return nil, false
	}
	return o.MetricsConfig, true
}

// HasMetricsConfig returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasMetricsConfig() bool {
	return o != nil && o.MetricsConfig != nil
}

// SetMetricsConfig gets a reference to the given TenancyConfigDataAttributesMetricsConfig and assigns it to the MetricsConfig field.
func (o *TenancyConfigDataAttributes) SetMetricsConfig(v TenancyConfigDataAttributesMetricsConfig) {
	o.MetricsConfig = &v
}

// GetParentTenancyName returns the ParentTenancyName field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetParentTenancyName() string {
	if o == nil || o.ParentTenancyName == nil {
		var ret string
		return ret
	}
	return *o.ParentTenancyName
}

// GetParentTenancyNameOk returns a tuple with the ParentTenancyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetParentTenancyNameOk() (*string, bool) {
	if o == nil || o.ParentTenancyName == nil {
		return nil, false
	}
	return o.ParentTenancyName, true
}

// HasParentTenancyName returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasParentTenancyName() bool {
	return o != nil && o.ParentTenancyName != nil
}

// SetParentTenancyName gets a reference to the given string and assigns it to the ParentTenancyName field.
func (o *TenancyConfigDataAttributes) SetParentTenancyName(v string) {
	o.ParentTenancyName = &v
}

// GetRegionsConfig returns the RegionsConfig field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetRegionsConfig() TenancyConfigDataAttributesRegionsConfig {
	if o == nil || o.RegionsConfig == nil {
		var ret TenancyConfigDataAttributesRegionsConfig
		return ret
	}
	return *o.RegionsConfig
}

// GetRegionsConfigOk returns a tuple with the RegionsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetRegionsConfigOk() (*TenancyConfigDataAttributesRegionsConfig, bool) {
	if o == nil || o.RegionsConfig == nil {
		return nil, false
	}
	return o.RegionsConfig, true
}

// HasRegionsConfig returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasRegionsConfig() bool {
	return o != nil && o.RegionsConfig != nil
}

// SetRegionsConfig gets a reference to the given TenancyConfigDataAttributesRegionsConfig and assigns it to the RegionsConfig field.
func (o *TenancyConfigDataAttributes) SetRegionsConfig(v TenancyConfigDataAttributesRegionsConfig) {
	o.RegionsConfig = &v
}

// GetResourceCollectionEnabled returns the ResourceCollectionEnabled field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetResourceCollectionEnabled() bool {
	if o == nil || o.ResourceCollectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ResourceCollectionEnabled
}

// GetResourceCollectionEnabledOk returns a tuple with the ResourceCollectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetResourceCollectionEnabledOk() (*bool, bool) {
	if o == nil || o.ResourceCollectionEnabled == nil {
		return nil, false
	}
	return o.ResourceCollectionEnabled, true
}

// HasResourceCollectionEnabled returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasResourceCollectionEnabled() bool {
	return o != nil && o.ResourceCollectionEnabled != nil
}

// SetResourceCollectionEnabled gets a reference to the given bool and assigns it to the ResourceCollectionEnabled field.
func (o *TenancyConfigDataAttributes) SetResourceCollectionEnabled(v bool) {
	o.ResourceCollectionEnabled = &v
}

// GetTenancyName returns the TenancyName field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetTenancyName() string {
	if o == nil || o.TenancyName == nil {
		var ret string
		return ret
	}
	return *o.TenancyName
}

// GetTenancyNameOk returns a tuple with the TenancyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetTenancyNameOk() (*string, bool) {
	if o == nil || o.TenancyName == nil {
		return nil, false
	}
	return o.TenancyName, true
}

// HasTenancyName returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasTenancyName() bool {
	return o != nil && o.TenancyName != nil
}

// SetTenancyName gets a reference to the given string and assigns it to the TenancyName field.
func (o *TenancyConfigDataAttributes) SetTenancyName(v string) {
	o.TenancyName = &v
}

// GetUserOcid returns the UserOcid field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetUserOcid() string {
	if o == nil || o.UserOcid == nil {
		var ret string
		return ret
	}
	return *o.UserOcid
}

// GetUserOcidOk returns a tuple with the UserOcid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetUserOcidOk() (*string, bool) {
	if o == nil || o.UserOcid == nil {
		return nil, false
	}
	return o.UserOcid, true
}

// HasUserOcid returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasUserOcid() bool {
	return o != nil && o.UserOcid != nil
}

// SetUserOcid gets a reference to the given string and assigns it to the UserOcid field.
func (o *TenancyConfigDataAttributes) SetUserOcid(v string) {
	o.UserOcid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TenancyConfigDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BillingPlanId != nil {
		toSerialize["billing_plan_id"] = o.BillingPlanId
	}
	if o.ConfigVersion != nil {
		toSerialize["config_version"] = o.ConfigVersion
	}
	if o.CostCollectionEnabled != nil {
		toSerialize["cost_collection_enabled"] = o.CostCollectionEnabled
	}
	if o.DdCompartmentId != nil {
		toSerialize["dd_compartment_id"] = o.DdCompartmentId
	}
	if o.DdStackId != nil {
		toSerialize["dd_stack_id"] = o.DdStackId
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
	if o.ParentTenancyName != nil {
		toSerialize["parent_tenancy_name"] = o.ParentTenancyName
	}
	if o.RegionsConfig != nil {
		toSerialize["regions_config"] = o.RegionsConfig
	}
	if o.ResourceCollectionEnabled != nil {
		toSerialize["resource_collection_enabled"] = o.ResourceCollectionEnabled
	}
	if o.TenancyName != nil {
		toSerialize["tenancy_name"] = o.TenancyName
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
func (o *TenancyConfigDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BillingPlanId             *int32                                    `json:"billing_plan_id,omitempty"`
		ConfigVersion             *int64                                    `json:"config_version,omitempty"`
		CostCollectionEnabled     *bool                                     `json:"cost_collection_enabled,omitempty"`
		DdCompartmentId           *string                                   `json:"dd_compartment_id,omitempty"`
		DdStackId                 *string                                   `json:"dd_stack_id,omitempty"`
		HomeRegion                *string                                   `json:"home_region,omitempty"`
		LogsConfig                *TenancyConfigDataAttributesLogsConfig    `json:"logs_config,omitempty"`
		MetricsConfig             *TenancyConfigDataAttributesMetricsConfig `json:"metrics_config,omitempty"`
		ParentTenancyName         *string                                   `json:"parent_tenancy_name,omitempty"`
		RegionsConfig             *TenancyConfigDataAttributesRegionsConfig `json:"regions_config,omitempty"`
		ResourceCollectionEnabled *bool                                     `json:"resource_collection_enabled,omitempty"`
		TenancyName               *string                                   `json:"tenancy_name,omitempty"`
		UserOcid                  *string                                   `json:"user_ocid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"billing_plan_id", "config_version", "cost_collection_enabled", "dd_compartment_id", "dd_stack_id", "home_region", "logs_config", "metrics_config", "parent_tenancy_name", "regions_config", "resource_collection_enabled", "tenancy_name", "user_ocid"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BillingPlanId = all.BillingPlanId
	o.ConfigVersion = all.ConfigVersion
	o.CostCollectionEnabled = all.CostCollectionEnabled
	o.DdCompartmentId = all.DdCompartmentId
	o.DdStackId = all.DdStackId
	o.HomeRegion = all.HomeRegion
	if all.LogsConfig != nil && all.LogsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogsConfig = all.LogsConfig
	if all.MetricsConfig != nil && all.MetricsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MetricsConfig = all.MetricsConfig
	o.ParentTenancyName = all.ParentTenancyName
	if all.RegionsConfig != nil && all.RegionsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RegionsConfig = all.RegionsConfig
	o.ResourceCollectionEnabled = all.ResourceCollectionEnabled
	o.TenancyName = all.TenancyName
	o.UserOcid = all.UserOcid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
