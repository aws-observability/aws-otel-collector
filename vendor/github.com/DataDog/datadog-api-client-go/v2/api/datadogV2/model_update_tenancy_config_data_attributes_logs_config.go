// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateTenancyConfigDataAttributesLogsConfig Log collection configuration for updating an OCI tenancy, controlling which compartments and services have log collection enabled.
type UpdateTenancyConfigDataAttributesLogsConfig struct {
	// List of compartment tag filters to scope log collection to specific compartments.
	CompartmentTagFilters []string `json:"compartment_tag_filters,omitempty"`
	// Whether log collection is enabled for the tenancy.
	Enabled *bool `json:"enabled,omitempty"`
	// List of OCI service names for which log collection is enabled.
	EnabledServices []string `json:"enabled_services,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateTenancyConfigDataAttributesLogsConfig instantiates a new UpdateTenancyConfigDataAttributesLogsConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateTenancyConfigDataAttributesLogsConfig() *UpdateTenancyConfigDataAttributesLogsConfig {
	this := UpdateTenancyConfigDataAttributesLogsConfig{}
	return &this
}

// NewUpdateTenancyConfigDataAttributesLogsConfigWithDefaults instantiates a new UpdateTenancyConfigDataAttributesLogsConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateTenancyConfigDataAttributesLogsConfigWithDefaults() *UpdateTenancyConfigDataAttributesLogsConfig {
	this := UpdateTenancyConfigDataAttributesLogsConfig{}
	return &this
}

// GetCompartmentTagFilters returns the CompartmentTagFilters field value if set, zero value otherwise.
func (o *UpdateTenancyConfigDataAttributesLogsConfig) GetCompartmentTagFilters() []string {
	if o == nil || o.CompartmentTagFilters == nil {
		var ret []string
		return ret
	}
	return o.CompartmentTagFilters
}

// GetCompartmentTagFiltersOk returns a tuple with the CompartmentTagFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigDataAttributesLogsConfig) GetCompartmentTagFiltersOk() (*[]string, bool) {
	if o == nil || o.CompartmentTagFilters == nil {
		return nil, false
	}
	return &o.CompartmentTagFilters, true
}

// HasCompartmentTagFilters returns a boolean if a field has been set.
func (o *UpdateTenancyConfigDataAttributesLogsConfig) HasCompartmentTagFilters() bool {
	return o != nil && o.CompartmentTagFilters != nil
}

// SetCompartmentTagFilters gets a reference to the given []string and assigns it to the CompartmentTagFilters field.
func (o *UpdateTenancyConfigDataAttributesLogsConfig) SetCompartmentTagFilters(v []string) {
	o.CompartmentTagFilters = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateTenancyConfigDataAttributesLogsConfig) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigDataAttributesLogsConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateTenancyConfigDataAttributesLogsConfig) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateTenancyConfigDataAttributesLogsConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEnabledServices returns the EnabledServices field value if set, zero value otherwise.
func (o *UpdateTenancyConfigDataAttributesLogsConfig) GetEnabledServices() []string {
	if o == nil || o.EnabledServices == nil {
		var ret []string
		return ret
	}
	return o.EnabledServices
}

// GetEnabledServicesOk returns a tuple with the EnabledServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigDataAttributesLogsConfig) GetEnabledServicesOk() (*[]string, bool) {
	if o == nil || o.EnabledServices == nil {
		return nil, false
	}
	return &o.EnabledServices, true
}

// HasEnabledServices returns a boolean if a field has been set.
func (o *UpdateTenancyConfigDataAttributesLogsConfig) HasEnabledServices() bool {
	return o != nil && o.EnabledServices != nil
}

// SetEnabledServices gets a reference to the given []string and assigns it to the EnabledServices field.
func (o *UpdateTenancyConfigDataAttributesLogsConfig) SetEnabledServices(v []string) {
	o.EnabledServices = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateTenancyConfigDataAttributesLogsConfig) MarshalJSON() ([]byte, error) {
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
	if o.EnabledServices != nil {
		toSerialize["enabled_services"] = o.EnabledServices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateTenancyConfigDataAttributesLogsConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompartmentTagFilters []string `json:"compartment_tag_filters,omitempty"`
		Enabled               *bool    `json:"enabled,omitempty"`
		EnabledServices       []string `json:"enabled_services,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compartment_tag_filters", "enabled", "enabled_services"})
	} else {
		return err
	}
	o.CompartmentTagFilters = all.CompartmentTagFilters
	o.Enabled = all.Enabled
	o.EnabledServices = all.EnabledServices

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
