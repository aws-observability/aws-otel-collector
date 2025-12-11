// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleMeetConfigurationReferenceData The Google Meet configuration relationship data object.
type GoogleMeetConfigurationReferenceData struct {
	// The unique identifier of the Google Meet configuration.
	Id string `json:"id"`
	// The type of the Google Meet configuration.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGoogleMeetConfigurationReferenceData instantiates a new GoogleMeetConfigurationReferenceData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGoogleMeetConfigurationReferenceData(id string, typeVar string) *GoogleMeetConfigurationReferenceData {
	this := GoogleMeetConfigurationReferenceData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewGoogleMeetConfigurationReferenceDataWithDefaults instantiates a new GoogleMeetConfigurationReferenceData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGoogleMeetConfigurationReferenceDataWithDefaults() *GoogleMeetConfigurationReferenceData {
	this := GoogleMeetConfigurationReferenceData{}
	return &this
}

// GetId returns the Id field value.
func (o *GoogleMeetConfigurationReferenceData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GoogleMeetConfigurationReferenceData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *GoogleMeetConfigurationReferenceData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *GoogleMeetConfigurationReferenceData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GoogleMeetConfigurationReferenceData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GoogleMeetConfigurationReferenceData) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GoogleMeetConfigurationReferenceData) MarshalJSON() ([]byte, error) {
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
func (o *GoogleMeetConfigurationReferenceData) UnmarshalJSON(bytes []byte) (err error) {
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

// NullableGoogleMeetConfigurationReferenceData handles when a null is used for GoogleMeetConfigurationReferenceData.
type NullableGoogleMeetConfigurationReferenceData struct {
	value *GoogleMeetConfigurationReferenceData
	isSet bool
}

// Get returns the associated value.
func (v NullableGoogleMeetConfigurationReferenceData) Get() *GoogleMeetConfigurationReferenceData {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableGoogleMeetConfigurationReferenceData) Set(val *GoogleMeetConfigurationReferenceData) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableGoogleMeetConfigurationReferenceData) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableGoogleMeetConfigurationReferenceData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableGoogleMeetConfigurationReferenceData initializes the struct as if Set has been called.
func NewNullableGoogleMeetConfigurationReferenceData(val *GoogleMeetConfigurationReferenceData) *NullableGoogleMeetConfigurationReferenceData {
	return &NullableGoogleMeetConfigurationReferenceData{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableGoogleMeetConfigurationReferenceData) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableGoogleMeetConfigurationReferenceData) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
