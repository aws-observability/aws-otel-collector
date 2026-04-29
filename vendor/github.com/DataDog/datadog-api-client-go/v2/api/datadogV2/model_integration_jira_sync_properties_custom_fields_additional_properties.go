// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties Synchronization configuration for a Jira custom field.
type IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties struct {
	// The type of synchronization to apply for this custom field.
	SyncType *string `json:"sync_type,omitempty"`
	// Represents any valid JSON value.
	Value NullableAnyValue `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties instantiates a new IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties() *IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties {
	this := IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties{}
	return &this
}

// NewIntegrationJiraSyncPropertiesCustomFieldsAdditionalPropertiesWithDefaults instantiates a new IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationJiraSyncPropertiesCustomFieldsAdditionalPropertiesWithDefaults() *IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties {
	this := IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties{}
	return &this
}

// GetSyncType returns the SyncType field value if set, zero value otherwise.
func (o *IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties) GetSyncType() string {
	if o == nil || o.SyncType == nil {
		var ret string
		return ret
	}
	return *o.SyncType
}

// GetSyncTypeOk returns a tuple with the SyncType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties) GetSyncTypeOk() (*string, bool) {
	if o == nil || o.SyncType == nil {
		return nil, false
	}
	return o.SyncType, true
}

// HasSyncType returns a boolean if a field has been set.
func (o *IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties) HasSyncType() bool {
	return o != nil && o.SyncType != nil
}

// SetSyncType gets a reference to the given string and assigns it to the SyncType field.
func (o *IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties) SetSyncType(v string) {
	o.SyncType = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties) GetValue() AnyValue {
	if o == nil || o.Value.Get() == nil {
		var ret AnyValue
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties) GetValueOk() (*AnyValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties) HasValue() bool {
	return o != nil && o.Value.IsSet()
}

// SetValue gets a reference to the given NullableAnyValue and assigns it to the Value field.
func (o *IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties) SetValue(v AnyValue) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil.
func (o *IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil.
func (o *IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties) UnsetValue() {
	o.Value.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.SyncType != nil {
		toSerialize["sync_type"] = o.SyncType
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SyncType *string          `json:"sync_type,omitempty"`
		Value    NullableAnyValue `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"sync_type", "value"})
	} else {
		return err
	}
	o.SyncType = all.SyncType
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
