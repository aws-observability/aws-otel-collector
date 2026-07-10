// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSdkConfigAttributes Attributes of the RUM SDK configuration.
type RumSdkConfigAttributes struct {
	// The RUM SDK settings for a configuration.
	Rum RumSdkConfigRumAttributes `json:"rum"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSdkConfigAttributes instantiates a new RumSdkConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSdkConfigAttributes(rum RumSdkConfigRumAttributes) *RumSdkConfigAttributes {
	this := RumSdkConfigAttributes{}
	this.Rum = rum
	return &this
}

// NewRumSdkConfigAttributesWithDefaults instantiates a new RumSdkConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSdkConfigAttributesWithDefaults() *RumSdkConfigAttributes {
	this := RumSdkConfigAttributes{}
	return &this
}

// GetRum returns the Rum field value.
func (o *RumSdkConfigAttributes) GetRum() RumSdkConfigRumAttributes {
	if o == nil {
		var ret RumSdkConfigRumAttributes
		return ret
	}
	return o.Rum
}

// GetRumOk returns a tuple with the Rum field value
// and a boolean to check if the value has been set.
func (o *RumSdkConfigAttributes) GetRumOk() (*RumSdkConfigRumAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rum, true
}

// SetRum sets field value.
func (o *RumSdkConfigAttributes) SetRum(v RumSdkConfigRumAttributes) {
	o.Rum = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSdkConfigAttributes) MarshalJSON() ([]byte, error) {
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
func (o *RumSdkConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rum *RumSdkConfigRumAttributes `json:"rum"`
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
