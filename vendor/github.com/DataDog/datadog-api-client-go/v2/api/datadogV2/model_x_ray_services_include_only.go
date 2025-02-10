// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// XRayServicesIncludeOnly Include only these services. Defaults to `[]`.
type XRayServicesIncludeOnly struct {
	// Include only these services.
	IncludeOnly []string `json:"include_only"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewXRayServicesIncludeOnly instantiates a new XRayServicesIncludeOnly object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewXRayServicesIncludeOnly(includeOnly []string) *XRayServicesIncludeOnly {
	this := XRayServicesIncludeOnly{}
	this.IncludeOnly = includeOnly
	return &this
}

// NewXRayServicesIncludeOnlyWithDefaults instantiates a new XRayServicesIncludeOnly object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewXRayServicesIncludeOnlyWithDefaults() *XRayServicesIncludeOnly {
	this := XRayServicesIncludeOnly{}
	return &this
}

// GetIncludeOnly returns the IncludeOnly field value.
func (o *XRayServicesIncludeOnly) GetIncludeOnly() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IncludeOnly
}

// GetIncludeOnlyOk returns a tuple with the IncludeOnly field value
// and a boolean to check if the value has been set.
func (o *XRayServicesIncludeOnly) GetIncludeOnlyOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeOnly, true
}

// SetIncludeOnly sets field value.
func (o *XRayServicesIncludeOnly) SetIncludeOnly(v []string) {
	o.IncludeOnly = v
}

// MarshalJSON serializes the struct using spec logic.
func (o XRayServicesIncludeOnly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["include_only"] = o.IncludeOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *XRayServicesIncludeOnly) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeOnly *[]string `json:"include_only"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IncludeOnly == nil {
		return fmt.Errorf("required field include_only missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include_only"})
	} else {
		return err
	}
	o.IncludeOnly = *all.IncludeOnly

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableXRayServicesIncludeOnly handles when a null is used for XRayServicesIncludeOnly.
type NullableXRayServicesIncludeOnly struct {
	value *XRayServicesIncludeOnly
	isSet bool
}

// Get returns the associated value.
func (v NullableXRayServicesIncludeOnly) Get() *XRayServicesIncludeOnly {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableXRayServicesIncludeOnly) Set(val *XRayServicesIncludeOnly) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableXRayServicesIncludeOnly) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableXRayServicesIncludeOnly) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableXRayServicesIncludeOnly initializes the struct as if Set has been called.
func NewNullableXRayServicesIncludeOnly(val *XRayServicesIncludeOnly) *NullableXRayServicesIncludeOnly {
	return &NullableXRayServicesIncludeOnly{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableXRayServicesIncludeOnly) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableXRayServicesIncludeOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
