// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImportFieldAttributesSingleValue A field with a single value selected.
type IncidentImportFieldAttributesSingleValue struct {
	// The single value selected for this field.
	Value datadog.NullableString `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewIncidentImportFieldAttributesSingleValue instantiates a new IncidentImportFieldAttributesSingleValue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentImportFieldAttributesSingleValue() *IncidentImportFieldAttributesSingleValue {
	this := IncidentImportFieldAttributesSingleValue{}
	return &this
}

// NewIncidentImportFieldAttributesSingleValueWithDefaults instantiates a new IncidentImportFieldAttributesSingleValue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentImportFieldAttributesSingleValueWithDefaults() *IncidentImportFieldAttributesSingleValue {
	this := IncidentImportFieldAttributesSingleValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportFieldAttributesSingleValue) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportFieldAttributesSingleValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *IncidentImportFieldAttributesSingleValue) HasValue() bool {
	return o != nil && o.Value.IsSet()
}

// SetValue gets a reference to the given datadog.NullableString and assigns it to the Value field.
func (o *IncidentImportFieldAttributesSingleValue) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil.
func (o *IncidentImportFieldAttributesSingleValue) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil.
func (o *IncidentImportFieldAttributesSingleValue) UnsetValue() {
	o.Value.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentImportFieldAttributesSingleValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentImportFieldAttributesSingleValue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Value datadog.NullableString `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Value = all.Value

	return nil
}
