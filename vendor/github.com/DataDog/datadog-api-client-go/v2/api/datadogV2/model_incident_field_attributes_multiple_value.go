// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentFieldAttributesMultipleValue A field with potentially multiple values selected.
type IncidentFieldAttributesMultipleValue struct {
	// Type of the multiple value field definitions.
	Type *IncidentFieldAttributesValueType `json:"type,omitempty"`
	// The multiple values selected for this field.
	Value datadog.NullableList[string] `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentFieldAttributesMultipleValue instantiates a new IncidentFieldAttributesMultipleValue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentFieldAttributesMultipleValue() *IncidentFieldAttributesMultipleValue {
	this := IncidentFieldAttributesMultipleValue{}
	var typeVar IncidentFieldAttributesValueType = INCIDENTFIELDATTRIBUTESVALUETYPE_MULTISELECT
	this.Type = &typeVar
	return &this
}

// NewIncidentFieldAttributesMultipleValueWithDefaults instantiates a new IncidentFieldAttributesMultipleValue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentFieldAttributesMultipleValueWithDefaults() *IncidentFieldAttributesMultipleValue {
	this := IncidentFieldAttributesMultipleValue{}
	var typeVar IncidentFieldAttributesValueType = INCIDENTFIELDATTRIBUTESVALUETYPE_MULTISELECT
	this.Type = &typeVar
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IncidentFieldAttributesMultipleValue) GetType() IncidentFieldAttributesValueType {
	if o == nil || o.Type == nil {
		var ret IncidentFieldAttributesValueType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentFieldAttributesMultipleValue) GetTypeOk() (*IncidentFieldAttributesValueType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IncidentFieldAttributesMultipleValue) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given IncidentFieldAttributesValueType and assigns it to the Type field.
func (o *IncidentFieldAttributesMultipleValue) SetType(v IncidentFieldAttributesValueType) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentFieldAttributesMultipleValue) GetValue() []string {
	if o == nil || o.Value.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentFieldAttributesMultipleValue) GetValueOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *IncidentFieldAttributesMultipleValue) HasValue() bool {
	return o != nil && o.Value.IsSet()
}

// SetValue gets a reference to the given datadog.NullableList[string] and assigns it to the Value field.
func (o *IncidentFieldAttributesMultipleValue) SetValue(v []string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil.
func (o *IncidentFieldAttributesMultipleValue) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil.
func (o *IncidentFieldAttributesMultipleValue) UnsetValue() {
	o.Value.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentFieldAttributesMultipleValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
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
func (o *IncidentFieldAttributesMultipleValue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *IncidentFieldAttributesValueType `json:"type,omitempty"`
		Value datadog.NullableList[string]      `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
