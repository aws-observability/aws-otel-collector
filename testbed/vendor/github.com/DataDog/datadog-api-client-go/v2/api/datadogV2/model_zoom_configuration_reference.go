// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ZoomConfigurationReference A reference to a Zoom configuration resource.
type ZoomConfigurationReference struct {
	// The Zoom configuration relationship data object.
	Data NullableZoomConfigurationReferenceData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewZoomConfigurationReference instantiates a new ZoomConfigurationReference object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewZoomConfigurationReference(data NullableZoomConfigurationReferenceData) *ZoomConfigurationReference {
	this := ZoomConfigurationReference{}
	this.Data = data
	return &this
}

// NewZoomConfigurationReferenceWithDefaults instantiates a new ZoomConfigurationReference object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewZoomConfigurationReferenceWithDefaults() *ZoomConfigurationReference {
	this := ZoomConfigurationReference{}
	return &this
}

// GetData returns the Data field value.
// If the value is explicit nil, the zero value for ZoomConfigurationReferenceData will be returned.
func (o *ZoomConfigurationReference) GetData() ZoomConfigurationReferenceData {
	if o == nil || o.Data.Get() == nil {
		var ret ZoomConfigurationReferenceData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ZoomConfigurationReference) GetDataOk() (*ZoomConfigurationReferenceData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// SetData sets field value.
func (o *ZoomConfigurationReference) SetData(v ZoomConfigurationReferenceData) {
	o.Data.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o ZoomConfigurationReference) MarshalJSON() ([]byte, error) {
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
func (o *ZoomConfigurationReference) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data NullableZoomConfigurationReferenceData `json:"data"`
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

// NullableZoomConfigurationReference handles when a null is used for ZoomConfigurationReference.
type NullableZoomConfigurationReference struct {
	value *ZoomConfigurationReference
	isSet bool
}

// Get returns the associated value.
func (v NullableZoomConfigurationReference) Get() *ZoomConfigurationReference {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableZoomConfigurationReference) Set(val *ZoomConfigurationReference) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableZoomConfigurationReference) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableZoomConfigurationReference) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableZoomConfigurationReference initializes the struct as if Set has been called.
func NewNullableZoomConfigurationReference(val *ZoomConfigurationReference) *NullableZoomConfigurationReference {
	return &NullableZoomConfigurationReference{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableZoomConfigurationReference) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableZoomConfigurationReference) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
