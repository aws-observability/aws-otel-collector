// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleMeetConfigurationReference A reference to a Google Meet Configuration resource.
type GoogleMeetConfigurationReference struct {
	// The Google Meet configuration relationship data object.
	Data NullableGoogleMeetConfigurationReferenceData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGoogleMeetConfigurationReference instantiates a new GoogleMeetConfigurationReference object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGoogleMeetConfigurationReference(data NullableGoogleMeetConfigurationReferenceData) *GoogleMeetConfigurationReference {
	this := GoogleMeetConfigurationReference{}
	this.Data = data
	return &this
}

// NewGoogleMeetConfigurationReferenceWithDefaults instantiates a new GoogleMeetConfigurationReference object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGoogleMeetConfigurationReferenceWithDefaults() *GoogleMeetConfigurationReference {
	this := GoogleMeetConfigurationReference{}
	return &this
}

// GetData returns the Data field value.
// If the value is explicit nil, the zero value for GoogleMeetConfigurationReferenceData will be returned.
func (o *GoogleMeetConfigurationReference) GetData() GoogleMeetConfigurationReferenceData {
	if o == nil || o.Data.Get() == nil {
		var ret GoogleMeetConfigurationReferenceData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GoogleMeetConfigurationReference) GetDataOk() (*GoogleMeetConfigurationReferenceData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// SetData sets field value.
func (o *GoogleMeetConfigurationReference) SetData(v GoogleMeetConfigurationReferenceData) {
	o.Data.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o GoogleMeetConfigurationReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GoogleMeetConfigurationReference) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data NullableGoogleMeetConfigurationReferenceData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Data.IsSet() {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableGoogleMeetConfigurationReference handles when a null is used for GoogleMeetConfigurationReference.
type NullableGoogleMeetConfigurationReference struct {
	value *GoogleMeetConfigurationReference
	isSet bool
}

// Get returns the associated value.
func (v NullableGoogleMeetConfigurationReference) Get() *GoogleMeetConfigurationReference {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableGoogleMeetConfigurationReference) Set(val *GoogleMeetConfigurationReference) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableGoogleMeetConfigurationReference) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableGoogleMeetConfigurationReference) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableGoogleMeetConfigurationReference initializes the struct as if Set has been called.
func NewNullableGoogleMeetConfigurationReference(val *GoogleMeetConfigurationReference) *NullableGoogleMeetConfigurationReference {
	return &NullableGoogleMeetConfigurationReference{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableGoogleMeetConfigurationReference) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableGoogleMeetConfigurationReference) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
