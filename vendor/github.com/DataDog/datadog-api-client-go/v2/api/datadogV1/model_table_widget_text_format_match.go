// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableWidgetTextFormatMatch Match rule for the table widget text format.
type TableWidgetTextFormatMatch struct {
	// Match or compare option.
	Type TableWidgetTextFormatMatchType `json:"type"`
	// Table Widget Match String.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTableWidgetTextFormatMatch instantiates a new TableWidgetTextFormatMatch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableWidgetTextFormatMatch(typeVar TableWidgetTextFormatMatchType, value string) *TableWidgetTextFormatMatch {
	this := TableWidgetTextFormatMatch{}
	this.Type = typeVar
	this.Value = value
	return &this
}

// NewTableWidgetTextFormatMatchWithDefaults instantiates a new TableWidgetTextFormatMatch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableWidgetTextFormatMatchWithDefaults() *TableWidgetTextFormatMatch {
	this := TableWidgetTextFormatMatch{}
	return &this
}

// GetType returns the Type field value.
func (o *TableWidgetTextFormatMatch) GetType() TableWidgetTextFormatMatchType {
	if o == nil {
		var ret TableWidgetTextFormatMatchType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TableWidgetTextFormatMatch) GetTypeOk() (*TableWidgetTextFormatMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TableWidgetTextFormatMatch) SetType(v TableWidgetTextFormatMatchType) {
	o.Type = v
}

// GetValue returns the Value field value.
func (o *TableWidgetTextFormatMatch) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TableWidgetTextFormatMatch) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *TableWidgetTextFormatMatch) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableWidgetTextFormatMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableWidgetTextFormatMatch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *TableWidgetTextFormatMatchType `json:"type"`
		Value *string                         `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type", "value"})
	} else {
		return err
	}

	hasInvalidField := false
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
