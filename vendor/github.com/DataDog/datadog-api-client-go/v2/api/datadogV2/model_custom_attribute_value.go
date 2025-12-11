// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomAttributeValue Custom attribute values
type CustomAttributeValue struct {
	// If true, value must be an array
	IsMulti bool `json:"is_multi"`
	// Custom attributes type
	Type CustomAttributeType `json:"type"`
	// Union of supported value for a custom attribute
	Value CustomAttributeValuesUnion `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomAttributeValue instantiates a new CustomAttributeValue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomAttributeValue(isMulti bool, typeVar CustomAttributeType, value CustomAttributeValuesUnion) *CustomAttributeValue {
	this := CustomAttributeValue{}
	this.IsMulti = isMulti
	this.Type = typeVar
	this.Value = value
	return &this
}

// NewCustomAttributeValueWithDefaults instantiates a new CustomAttributeValue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomAttributeValueWithDefaults() *CustomAttributeValue {
	this := CustomAttributeValue{}
	return &this
}

// GetIsMulti returns the IsMulti field value.
func (o *CustomAttributeValue) GetIsMulti() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsMulti
}

// GetIsMultiOk returns a tuple with the IsMulti field value
// and a boolean to check if the value has been set.
func (o *CustomAttributeValue) GetIsMultiOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMulti, true
}

// SetIsMulti sets field value.
func (o *CustomAttributeValue) SetIsMulti(v bool) {
	o.IsMulti = v
}

// GetType returns the Type field value.
func (o *CustomAttributeValue) GetType() CustomAttributeType {
	if o == nil {
		var ret CustomAttributeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomAttributeValue) GetTypeOk() (*CustomAttributeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomAttributeValue) SetType(v CustomAttributeType) {
	o.Type = v
}

// GetValue returns the Value field value.
func (o *CustomAttributeValue) GetValue() CustomAttributeValuesUnion {
	if o == nil {
		var ret CustomAttributeValuesUnion
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CustomAttributeValue) GetValueOk() (*CustomAttributeValuesUnion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *CustomAttributeValue) SetValue(v CustomAttributeValuesUnion) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomAttributeValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["is_multi"] = o.IsMulti
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomAttributeValue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsMulti *bool                       `json:"is_multi"`
		Type    *CustomAttributeType        `json:"type"`
		Value   *CustomAttributeValuesUnion `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsMulti == nil {
		return fmt.Errorf("required field is_multi missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_multi", "type", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.IsMulti = *all.IsMulti
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
