// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableWidgetTextFormatReplaceSubstring Match Sub-string definition.
type TableWidgetTextFormatReplaceSubstring struct {
	// Text that will be replaced.
	Substring string `json:"substring"`
	// Table widget text format replace sub-string type.
	Type TableWidgetTextFormatReplaceSubstringType `json:"type"`
	// Text that will replace original sub-string.
	With string `json:"with"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTableWidgetTextFormatReplaceSubstring instantiates a new TableWidgetTextFormatReplaceSubstring object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableWidgetTextFormatReplaceSubstring(substring string, typeVar TableWidgetTextFormatReplaceSubstringType, with string) *TableWidgetTextFormatReplaceSubstring {
	this := TableWidgetTextFormatReplaceSubstring{}
	this.Substring = substring
	this.Type = typeVar
	this.With = with
	return &this
}

// NewTableWidgetTextFormatReplaceSubstringWithDefaults instantiates a new TableWidgetTextFormatReplaceSubstring object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableWidgetTextFormatReplaceSubstringWithDefaults() *TableWidgetTextFormatReplaceSubstring {
	this := TableWidgetTextFormatReplaceSubstring{}
	return &this
}

// GetSubstring returns the Substring field value.
func (o *TableWidgetTextFormatReplaceSubstring) GetSubstring() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Substring
}

// GetSubstringOk returns a tuple with the Substring field value
// and a boolean to check if the value has been set.
func (o *TableWidgetTextFormatReplaceSubstring) GetSubstringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Substring, true
}

// SetSubstring sets field value.
func (o *TableWidgetTextFormatReplaceSubstring) SetSubstring(v string) {
	o.Substring = v
}

// GetType returns the Type field value.
func (o *TableWidgetTextFormatReplaceSubstring) GetType() TableWidgetTextFormatReplaceSubstringType {
	if o == nil {
		var ret TableWidgetTextFormatReplaceSubstringType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TableWidgetTextFormatReplaceSubstring) GetTypeOk() (*TableWidgetTextFormatReplaceSubstringType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TableWidgetTextFormatReplaceSubstring) SetType(v TableWidgetTextFormatReplaceSubstringType) {
	o.Type = v
}

// GetWith returns the With field value.
func (o *TableWidgetTextFormatReplaceSubstring) GetWith() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.With
}

// GetWithOk returns a tuple with the With field value
// and a boolean to check if the value has been set.
func (o *TableWidgetTextFormatReplaceSubstring) GetWithOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.With, true
}

// SetWith sets field value.
func (o *TableWidgetTextFormatReplaceSubstring) SetWith(v string) {
	o.With = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableWidgetTextFormatReplaceSubstring) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["substring"] = o.Substring
	toSerialize["type"] = o.Type
	toSerialize["with"] = o.With

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableWidgetTextFormatReplaceSubstring) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Substring *string                                    `json:"substring"`
		Type      *TableWidgetTextFormatReplaceSubstringType `json:"type"`
		With      *string                                    `json:"with"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Substring == nil {
		return fmt.Errorf("required field substring missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.With == nil {
		return fmt.Errorf("required field with missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"substring", "type", "with"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Substring = *all.Substring
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
