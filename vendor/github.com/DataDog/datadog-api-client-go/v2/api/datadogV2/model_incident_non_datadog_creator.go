// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNonDatadogCreator Incident's non Datadog creator.
type IncidentNonDatadogCreator struct {
	// Non Datadog creator `48px` image.
	Image48Px *string `json:"image_48_px,omitempty"`
	// Non Datadog creator name.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentNonDatadogCreator instantiates a new IncidentNonDatadogCreator object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentNonDatadogCreator() *IncidentNonDatadogCreator {
	this := IncidentNonDatadogCreator{}
	return &this
}

// NewIncidentNonDatadogCreatorWithDefaults instantiates a new IncidentNonDatadogCreator object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentNonDatadogCreatorWithDefaults() *IncidentNonDatadogCreator {
	this := IncidentNonDatadogCreator{}
	return &this
}

// GetImage48Px returns the Image48Px field value if set, zero value otherwise.
func (o *IncidentNonDatadogCreator) GetImage48Px() string {
	if o == nil || o.Image48Px == nil {
		var ret string
		return ret
	}
	return *o.Image48Px
}

// GetImage48PxOk returns a tuple with the Image48Px field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNonDatadogCreator) GetImage48PxOk() (*string, bool) {
	if o == nil || o.Image48Px == nil {
		return nil, false
	}
	return o.Image48Px, true
}

// HasImage48Px returns a boolean if a field has been set.
func (o *IncidentNonDatadogCreator) HasImage48Px() bool {
	return o != nil && o.Image48Px != nil
}

// SetImage48Px gets a reference to the given string and assigns it to the Image48Px field.
func (o *IncidentNonDatadogCreator) SetImage48Px(v string) {
	o.Image48Px = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IncidentNonDatadogCreator) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNonDatadogCreator) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IncidentNonDatadogCreator) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IncidentNonDatadogCreator) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentNonDatadogCreator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Image48Px != nil {
		toSerialize["image_48_px"] = o.Image48Px
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentNonDatadogCreator) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Image48Px *string `json:"image_48_px,omitempty"`
		Name      *string `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"image_48_px", "name"})
	} else {
		return err
	}
	o.Image48Px = all.Image48Px
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableIncidentNonDatadogCreator handles when a null is used for IncidentNonDatadogCreator.
type NullableIncidentNonDatadogCreator struct {
	value *IncidentNonDatadogCreator
	isSet bool
}

// Get returns the associated value.
func (v NullableIncidentNonDatadogCreator) Get() *IncidentNonDatadogCreator {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableIncidentNonDatadogCreator) Set(val *IncidentNonDatadogCreator) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableIncidentNonDatadogCreator) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableIncidentNonDatadogCreator) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncidentNonDatadogCreator initializes the struct as if Set has been called.
func NewNullableIncidentNonDatadogCreator(val *IncidentNonDatadogCreator) *NullableIncidentNonDatadogCreator {
	return &NullableIncidentNonDatadogCreator{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableIncidentNonDatadogCreator) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableIncidentNonDatadogCreator) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
