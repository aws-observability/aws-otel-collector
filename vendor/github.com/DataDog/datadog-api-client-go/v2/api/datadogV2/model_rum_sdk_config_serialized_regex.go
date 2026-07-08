// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSdkConfigSerializedRegex A serialized regex used as an extractor in dynamic options.
type RumSdkConfigSerializedRegex struct {
	// The type identifier for a serialized regex. Always `regex`.
	RcSerializedType RumSdkConfigSerializedRegexType `json:"rc_serialized_type"`
	// The regex pattern used for extraction.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSdkConfigSerializedRegex instantiates a new RumSdkConfigSerializedRegex object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSdkConfigSerializedRegex(rcSerializedType RumSdkConfigSerializedRegexType, value string) *RumSdkConfigSerializedRegex {
	this := RumSdkConfigSerializedRegex{}
	this.RcSerializedType = rcSerializedType
	this.Value = value
	return &this
}

// NewRumSdkConfigSerializedRegexWithDefaults instantiates a new RumSdkConfigSerializedRegex object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSdkConfigSerializedRegexWithDefaults() *RumSdkConfigSerializedRegex {
	this := RumSdkConfigSerializedRegex{}
	return &this
}

// GetRcSerializedType returns the RcSerializedType field value.
func (o *RumSdkConfigSerializedRegex) GetRcSerializedType() RumSdkConfigSerializedRegexType {
	if o == nil {
		var ret RumSdkConfigSerializedRegexType
		return ret
	}
	return o.RcSerializedType
}

// GetRcSerializedTypeOk returns a tuple with the RcSerializedType field value
// and a boolean to check if the value has been set.
func (o *RumSdkConfigSerializedRegex) GetRcSerializedTypeOk() (*RumSdkConfigSerializedRegexType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RcSerializedType, true
}

// SetRcSerializedType sets field value.
func (o *RumSdkConfigSerializedRegex) SetRcSerializedType(v RumSdkConfigSerializedRegexType) {
	o.RcSerializedType = v
}

// GetValue returns the Value field value.
func (o *RumSdkConfigSerializedRegex) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RumSdkConfigSerializedRegex) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *RumSdkConfigSerializedRegex) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSdkConfigSerializedRegex) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["rc_serialized_type"] = o.RcSerializedType
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSdkConfigSerializedRegex) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RcSerializedType *RumSdkConfigSerializedRegexType `json:"rc_serialized_type"`
		Value            *string                          `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RcSerializedType == nil {
		return fmt.Errorf("required field rc_serialized_type missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rc_serialized_type", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.RcSerializedType.IsValid() {
		hasInvalidField = true
	} else {
		o.RcSerializedType = *all.RcSerializedType
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
