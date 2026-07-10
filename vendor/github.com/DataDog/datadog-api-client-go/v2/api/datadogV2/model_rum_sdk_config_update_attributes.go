// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSdkConfigUpdateAttributes Attributes of the RUM SDK configuration to update.
type RumSdkConfigUpdateAttributes struct {
	// The RUM SDK settings to apply when updating a configuration.
	Rum RumSdkConfigRumUpdateAttributes `json:"rum"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSdkConfigUpdateAttributes instantiates a new RumSdkConfigUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSdkConfigUpdateAttributes(rum RumSdkConfigRumUpdateAttributes) *RumSdkConfigUpdateAttributes {
	this := RumSdkConfigUpdateAttributes{}
	this.Rum = rum
	return &this
}

// NewRumSdkConfigUpdateAttributesWithDefaults instantiates a new RumSdkConfigUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSdkConfigUpdateAttributesWithDefaults() *RumSdkConfigUpdateAttributes {
	this := RumSdkConfigUpdateAttributes{}
	return &this
}

// GetRum returns the Rum field value.
func (o *RumSdkConfigUpdateAttributes) GetRum() RumSdkConfigRumUpdateAttributes {
	if o == nil {
		var ret RumSdkConfigRumUpdateAttributes
		return ret
	}
	return o.Rum
}

// GetRumOk returns a tuple with the Rum field value
// and a boolean to check if the value has been set.
func (o *RumSdkConfigUpdateAttributes) GetRumOk() (*RumSdkConfigRumUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rum, true
}

// SetRum sets field value.
func (o *RumSdkConfigUpdateAttributes) SetRum(v RumSdkConfigRumUpdateAttributes) {
	o.Rum = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSdkConfigUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["rum"] = o.Rum

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSdkConfigUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rum *RumSdkConfigRumUpdateAttributes `json:"rum"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Rum == nil {
		return fmt.Errorf("required field rum missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rum"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Rum.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rum = *all.Rum

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
