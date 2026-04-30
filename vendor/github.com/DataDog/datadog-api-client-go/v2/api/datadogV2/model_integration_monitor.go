// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationMonitor Monitor integration settings.
type IntegrationMonitor struct {
	// Whether auto-resolve is enabled.
	AutoResolveEnabled *bool `json:"auto_resolve_enabled,omitempty"`
	// Case type ID for monitor integration.
	CaseTypeId *string `json:"case_type_id,omitempty"`
	// Whether monitor integration is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Monitor handle.
	Handle *string `json:"handle,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationMonitor instantiates a new IntegrationMonitor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationMonitor() *IntegrationMonitor {
	this := IntegrationMonitor{}
	return &this
}

// NewIntegrationMonitorWithDefaults instantiates a new IntegrationMonitor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationMonitorWithDefaults() *IntegrationMonitor {
	this := IntegrationMonitor{}
	return &this
}

// GetAutoResolveEnabled returns the AutoResolveEnabled field value if set, zero value otherwise.
func (o *IntegrationMonitor) GetAutoResolveEnabled() bool {
	if o == nil || o.AutoResolveEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AutoResolveEnabled
}

// GetAutoResolveEnabledOk returns a tuple with the AutoResolveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationMonitor) GetAutoResolveEnabledOk() (*bool, bool) {
	if o == nil || o.AutoResolveEnabled == nil {
		return nil, false
	}
	return o.AutoResolveEnabled, true
}

// HasAutoResolveEnabled returns a boolean if a field has been set.
func (o *IntegrationMonitor) HasAutoResolveEnabled() bool {
	return o != nil && o.AutoResolveEnabled != nil
}

// SetAutoResolveEnabled gets a reference to the given bool and assigns it to the AutoResolveEnabled field.
func (o *IntegrationMonitor) SetAutoResolveEnabled(v bool) {
	o.AutoResolveEnabled = &v
}

// GetCaseTypeId returns the CaseTypeId field value if set, zero value otherwise.
func (o *IntegrationMonitor) GetCaseTypeId() string {
	if o == nil || o.CaseTypeId == nil {
		var ret string
		return ret
	}
	return *o.CaseTypeId
}

// GetCaseTypeIdOk returns a tuple with the CaseTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationMonitor) GetCaseTypeIdOk() (*string, bool) {
	if o == nil || o.CaseTypeId == nil {
		return nil, false
	}
	return o.CaseTypeId, true
}

// HasCaseTypeId returns a boolean if a field has been set.
func (o *IntegrationMonitor) HasCaseTypeId() bool {
	return o != nil && o.CaseTypeId != nil
}

// SetCaseTypeId gets a reference to the given string and assigns it to the CaseTypeId field.
func (o *IntegrationMonitor) SetCaseTypeId(v string) {
	o.CaseTypeId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IntegrationMonitor) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationMonitor) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IntegrationMonitor) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IntegrationMonitor) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *IntegrationMonitor) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationMonitor) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *IntegrationMonitor) HasHandle() bool {
	return o != nil && o.Handle != nil
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *IntegrationMonitor) SetHandle(v string) {
	o.Handle = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationMonitor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutoResolveEnabled != nil {
		toSerialize["auto_resolve_enabled"] = o.AutoResolveEnabled
	}
	if o.CaseTypeId != nil {
		toSerialize["case_type_id"] = o.CaseTypeId
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationMonitor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoResolveEnabled *bool   `json:"auto_resolve_enabled,omitempty"`
		CaseTypeId         *string `json:"case_type_id,omitempty"`
		Enabled            *bool   `json:"enabled,omitempty"`
		Handle             *string `json:"handle,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_resolve_enabled", "case_type_id", "enabled", "handle"})
	} else {
		return err
	}
	o.AutoResolveEnabled = all.AutoResolveEnabled
	o.CaseTypeId = all.CaseTypeId
	o.Enabled = all.Enabled
	o.Handle = all.Handle

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
