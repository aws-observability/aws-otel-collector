// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableWidgetTextFormatReplaceAll Match All definition.
type TableWidgetTextFormatReplaceAll struct {
	// Table widget text format replace all type.
	Type TableWidgetTextFormatReplaceAllType `json:"type"`
	// Replace All type.
	With string `json:"with"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTableWidgetTextFormatReplaceAll instantiates a new TableWidgetTextFormatReplaceAll object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableWidgetTextFormatReplaceAll(typeVar TableWidgetTextFormatReplaceAllType, with string) *TableWidgetTextFormatReplaceAll {
	this := TableWidgetTextFormatReplaceAll{}
	this.Type = typeVar
	this.With = with
	return &this
}

// NewTableWidgetTextFormatReplaceAllWithDefaults instantiates a new TableWidgetTextFormatReplaceAll object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableWidgetTextFormatReplaceAllWithDefaults() *TableWidgetTextFormatReplaceAll {
	this := TableWidgetTextFormatReplaceAll{}
	return &this
}

// GetType returns the Type field value.
func (o *TableWidgetTextFormatReplaceAll) GetType() TableWidgetTextFormatReplaceAllType {
	if o == nil {
		var ret TableWidgetTextFormatReplaceAllType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TableWidgetTextFormatReplaceAll) GetTypeOk() (*TableWidgetTextFormatReplaceAllType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TableWidgetTextFormatReplaceAll) SetType(v TableWidgetTextFormatReplaceAllType) {
	o.Type = v
}

// GetWith returns the With field value.
func (o *TableWidgetTextFormatReplaceAll) GetWith() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.With
}

// GetWithOk returns a tuple with the With field value
// and a boolean to check if the value has been set.
func (o *TableWidgetTextFormatReplaceAll) GetWithOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.With, true
}

// SetWith sets field value.
func (o *TableWidgetTextFormatReplaceAll) SetWith(v string) {
	o.With = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableWidgetTextFormatReplaceAll) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["with"] = o.With

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableWidgetTextFormatReplaceAll) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *TableWidgetTextFormatReplaceAllType `json:"type"`
		With *string                              `json:"with"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.With == nil {
		return fmt.Errorf("required field with missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type", "with"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.With = *all.With

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
