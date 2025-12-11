// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ZoomConfigurationReferenceData The Zoom configuration relationship data object.
type ZoomConfigurationReferenceData struct {
	// The unique identifier of the Zoom configuration.
	Id string `json:"id"`
	// The type of the Zoom configuration.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewZoomConfigurationReferenceData instantiates a new ZoomConfigurationReferenceData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewZoomConfigurationReferenceData(id string, typeVar string) *ZoomConfigurationReferenceData {
	this := ZoomConfigurationReferenceData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewZoomConfigurationReferenceDataWithDefaults instantiates a new ZoomConfigurationReferenceData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewZoomConfigurationReferenceDataWithDefaults() *ZoomConfigurationReferenceData {
	this := ZoomConfigurationReferenceData{}
	return &this
}

// GetId returns the Id field value.
func (o *ZoomConfigurationReferenceData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ZoomConfigurationReferenceData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ZoomConfigurationReferenceData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ZoomConfigurationReferenceData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ZoomConfigurationReferenceData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ZoomConfigurationReferenceData) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ZoomConfigurationReferenceData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ZoomConfigurationReferenceData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string `json:"id"`
		Type *string `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableZoomConfigurationReferenceData handles when a null is used for ZoomConfigurationReferenceData.
type NullableZoomConfigurationReferenceData struct {
	value *ZoomConfigurationReferenceData
	isSet bool
}

// Get returns the associated value.
func (v NullableZoomConfigurationReferenceData) Get() *ZoomConfigurationReferenceData {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableZoomConfigurationReferenceData) Set(val *ZoomConfigurationReferenceData) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableZoomConfigurationReferenceData) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableZoomConfigurationReferenceData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableZoomConfigurationReferenceData initializes the struct as if Set has been called.
func NewNullableZoomConfigurationReferenceData(val *ZoomConfigurationReferenceData) *NullableZoomConfigurationReferenceData {
	return &NullableZoomConfigurationReferenceData{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableZoomConfigurationReferenceData) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableZoomConfigurationReferenceData) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
