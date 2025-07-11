// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NumberFormatUnitCustom Custom unit.
type NumberFormatUnitCustom struct {
	// The label for the custom unit.
	Label *string `json:"label,omitempty"`
	// The type of custom unit.
	Type *NumberFormatUnitCustomType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNumberFormatUnitCustom instantiates a new NumberFormatUnitCustom object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNumberFormatUnitCustom() *NumberFormatUnitCustom {
	this := NumberFormatUnitCustom{}
	return &this
}

// NewNumberFormatUnitCustomWithDefaults instantiates a new NumberFormatUnitCustom object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNumberFormatUnitCustomWithDefaults() *NumberFormatUnitCustom {
	this := NumberFormatUnitCustom{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *NumberFormatUnitCustom) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormatUnitCustom) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *NumberFormatUnitCustom) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *NumberFormatUnitCustom) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NumberFormatUnitCustom) GetType() NumberFormatUnitCustomType {
	if o == nil || o.Type == nil {
		var ret NumberFormatUnitCustomType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormatUnitCustom) GetTypeOk() (*NumberFormatUnitCustomType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NumberFormatUnitCustom) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given NumberFormatUnitCustomType and assigns it to the Type field.
func (o *NumberFormatUnitCustom) SetType(v NumberFormatUnitCustomType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NumberFormatUnitCustom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NumberFormatUnitCustom) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Label *string                     `json:"label,omitempty"`
		Type  *NumberFormatUnitCustomType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"label", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Label = all.Label
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
