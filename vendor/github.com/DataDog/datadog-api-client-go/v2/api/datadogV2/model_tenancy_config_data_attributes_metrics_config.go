// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TenancyConfigDataAttributesMetricsConfig Metrics collection configuration for an OCI tenancy, indicating which compartments and services are included or excluded.
type TenancyConfigDataAttributesMetricsConfig struct {
	// List of compartment tag filters scoping metrics collection to specific compartments.
	CompartmentTagFilters []string `json:"compartment_tag_filters,omitempty"`
	// Whether metrics collection is enabled for the tenancy.
	Enabled *bool `json:"enabled,omitempty"`
	// List of OCI service names excluded from metrics collection.
	ExcludedServices []string `json:"excluded_services,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTenancyConfigDataAttributesMetricsConfig instantiates a new TenancyConfigDataAttributesMetricsConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTenancyConfigDataAttributesMetricsConfig() *TenancyConfigDataAttributesMetricsConfig {
	this := TenancyConfigDataAttributesMetricsConfig{}
	return &this
}

// NewTenancyConfigDataAttributesMetricsConfigWithDefaults instantiates a new TenancyConfigDataAttributesMetricsConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTenancyConfigDataAttributesMetricsConfigWithDefaults() *TenancyConfigDataAttributesMetricsConfig {
	this := TenancyConfigDataAttributesMetricsConfig{}
	return &this
}

// GetCompartmentTagFilters returns the CompartmentTagFilters field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributesMetricsConfig) GetCompartmentTagFilters() []string {
	if o == nil || o.CompartmentTagFilters == nil {
		var ret []string
		return ret
	}
	return o.CompartmentTagFilters
}

// GetCompartmentTagFiltersOk returns a tuple with the CompartmentTagFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributesMetricsConfig) GetCompartmentTagFiltersOk() (*[]string, bool) {
	if o == nil || o.CompartmentTagFilters == nil {
		return nil, false
	}
	return &o.CompartmentTagFilters, true
}

// HasCompartmentTagFilters returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributesMetricsConfig) HasCompartmentTagFilters() bool {
	return o != nil && o.CompartmentTagFilters != nil
}

// SetCompartmentTagFilters gets a reference to the given []string and assigns it to the CompartmentTagFilters field.
func (o *TenancyConfigDataAttributesMetricsConfig) SetCompartmentTagFilters(v []string) {
	o.CompartmentTagFilters = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributesMetricsConfig) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributesMetricsConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributesMetricsConfig) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TenancyConfigDataAttributesMetricsConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExcludedServices returns the ExcludedServices field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributesMetricsConfig) GetExcludedServices() []string {
	if o == nil || o.ExcludedServices == nil {
		var ret []string
		return ret
	}
	return o.ExcludedServices
}

// GetExcludedServicesOk returns a tuple with the ExcludedServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributesMetricsConfig) GetExcludedServicesOk() (*[]string, bool) {
	if o == nil || o.ExcludedServices == nil {
		return nil, false
	}
	return &o.ExcludedServices, true
}

// HasExcludedServices returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributesMetricsConfig) HasExcludedServices() bool {
	return o != nil && o.ExcludedServices != nil
}

// SetExcludedServices gets a reference to the given []string and assigns it to the ExcludedServices field.
func (o *TenancyConfigDataAttributesMetricsConfig) SetExcludedServices(v []string) {
	o.ExcludedServices = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TenancyConfigDataAttributesMetricsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompartmentTagFilters != nil {
		toSerialize["compartment_tag_filters"] = o.CompartmentTagFilters
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ExcludedServices != nil {
		toSerialize["excluded_services"] = o.ExcludedServices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TenancyConfigDataAttributesMetricsConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompartmentTagFilters []string `json:"compartment_tag_filters,omitempty"`
		Enabled               *bool    `json:"enabled,omitempty"`
		ExcludedServices      []string `json:"excluded_services,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compartment_tag_filters", "enabled", "excluded_services"})
	} else {
		return err
	}
	o.CompartmentTagFilters = all.CompartmentTagFilters
	o.Enabled = all.Enabled
	o.ExcludedServices = all.ExcludedServices

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
