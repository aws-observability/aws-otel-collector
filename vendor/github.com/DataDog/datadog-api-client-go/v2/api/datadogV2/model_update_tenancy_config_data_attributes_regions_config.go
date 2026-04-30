// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateTenancyConfigDataAttributesRegionsConfig Region configuration for updating an OCI tenancy, specifying which regions are available, enabled, or disabled for data collection.
type UpdateTenancyConfigDataAttributesRegionsConfig struct {
	// List of OCI regions available for data collection in the tenancy.
	Available []string `json:"available,omitempty"`
	// List of OCI regions explicitly disabled for data collection.
	Disabled []string `json:"disabled,omitempty"`
	// List of OCI regions enabled for data collection.
	Enabled []string `json:"enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateTenancyConfigDataAttributesRegionsConfig instantiates a new UpdateTenancyConfigDataAttributesRegionsConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateTenancyConfigDataAttributesRegionsConfig() *UpdateTenancyConfigDataAttributesRegionsConfig {
	this := UpdateTenancyConfigDataAttributesRegionsConfig{}
	return &this
}

// NewUpdateTenancyConfigDataAttributesRegionsConfigWithDefaults instantiates a new UpdateTenancyConfigDataAttributesRegionsConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateTenancyConfigDataAttributesRegionsConfigWithDefaults() *UpdateTenancyConfigDataAttributesRegionsConfig {
	this := UpdateTenancyConfigDataAttributesRegionsConfig{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *UpdateTenancyConfigDataAttributesRegionsConfig) GetAvailable() []string {
	if o == nil || o.Available == nil {
		var ret []string
		return ret
	}
	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigDataAttributesRegionsConfig) GetAvailableOk() (*[]string, bool) {
	if o == nil || o.Available == nil {
		return nil, false
	}
	return &o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *UpdateTenancyConfigDataAttributesRegionsConfig) HasAvailable() bool {
	return o != nil && o.Available != nil
}

// SetAvailable gets a reference to the given []string and assigns it to the Available field.
func (o *UpdateTenancyConfigDataAttributesRegionsConfig) SetAvailable(v []string) {
	o.Available = v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *UpdateTenancyConfigDataAttributesRegionsConfig) GetDisabled() []string {
	if o == nil || o.Disabled == nil {
		var ret []string
		return ret
	}
	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigDataAttributesRegionsConfig) GetDisabledOk() (*[]string, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *UpdateTenancyConfigDataAttributesRegionsConfig) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given []string and assigns it to the Disabled field.
func (o *UpdateTenancyConfigDataAttributesRegionsConfig) SetDisabled(v []string) {
	o.Disabled = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateTenancyConfigDataAttributesRegionsConfig) GetEnabled() []string {
	if o == nil || o.Enabled == nil {
		var ret []string
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigDataAttributesRegionsConfig) GetEnabledOk() (*[]string, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateTenancyConfigDataAttributesRegionsConfig) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given []string and assigns it to the Enabled field.
func (o *UpdateTenancyConfigDataAttributesRegionsConfig) SetEnabled(v []string) {
	o.Enabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateTenancyConfigDataAttributesRegionsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Available != nil {
		toSerialize["available"] = o.Available
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateTenancyConfigDataAttributesRegionsConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Available []string `json:"available,omitempty"`
		Disabled  []string `json:"disabled,omitempty"`
		Enabled   []string `json:"enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"available", "disabled", "enabled"})
	} else {
		return err
	}
	o.Available = all.Available
	o.Disabled = all.Disabled
	o.Enabled = all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
