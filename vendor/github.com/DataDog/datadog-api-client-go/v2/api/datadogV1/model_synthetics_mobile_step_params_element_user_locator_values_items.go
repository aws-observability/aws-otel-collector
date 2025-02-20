// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileStepParamsElementUserLocatorValuesItems A single user locator object.
type SyntheticsMobileStepParamsElementUserLocatorValuesItems struct {
	// Type of a user locator.
	Type *SyntheticsMobileStepParamsElementUserLocatorValuesItemsType `json:"type,omitempty"`
	// Value of a user locator.
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileStepParamsElementUserLocatorValuesItems instantiates a new SyntheticsMobileStepParamsElementUserLocatorValuesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileStepParamsElementUserLocatorValuesItems() *SyntheticsMobileStepParamsElementUserLocatorValuesItems {
	this := SyntheticsMobileStepParamsElementUserLocatorValuesItems{}
	return &this
}

// NewSyntheticsMobileStepParamsElementUserLocatorValuesItemsWithDefaults instantiates a new SyntheticsMobileStepParamsElementUserLocatorValuesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileStepParamsElementUserLocatorValuesItemsWithDefaults() *SyntheticsMobileStepParamsElementUserLocatorValuesItems {
	this := SyntheticsMobileStepParamsElementUserLocatorValuesItems{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsElementUserLocatorValuesItems) GetType() SyntheticsMobileStepParamsElementUserLocatorValuesItemsType {
	if o == nil || o.Type == nil {
		var ret SyntheticsMobileStepParamsElementUserLocatorValuesItemsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsElementUserLocatorValuesItems) GetTypeOk() (*SyntheticsMobileStepParamsElementUserLocatorValuesItemsType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsElementUserLocatorValuesItems) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SyntheticsMobileStepParamsElementUserLocatorValuesItemsType and assigns it to the Type field.
func (o *SyntheticsMobileStepParamsElementUserLocatorValuesItems) SetType(v SyntheticsMobileStepParamsElementUserLocatorValuesItemsType) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsElementUserLocatorValuesItems) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsElementUserLocatorValuesItems) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsElementUserLocatorValuesItems) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SyntheticsMobileStepParamsElementUserLocatorValuesItems) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileStepParamsElementUserLocatorValuesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileStepParamsElementUserLocatorValuesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *SyntheticsMobileStepParamsElementUserLocatorValuesItemsType `json:"type,omitempty"`
		Value *string                                                      `json:"value,omitempty"`
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
