// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsApiMultistepSubtestAttributes Attributes of a Synthetic API multistep subtest.
type SyntheticsApiMultistepSubtestAttributes struct {
	// Name of the subtest.
	Name *string `json:"name,omitempty"`
	// The public ID of the subtest.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsApiMultistepSubtestAttributes instantiates a new SyntheticsApiMultistepSubtestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsApiMultistepSubtestAttributes() *SyntheticsApiMultistepSubtestAttributes {
	this := SyntheticsApiMultistepSubtestAttributes{}
	return &this
}

// NewSyntheticsApiMultistepSubtestAttributesWithDefaults instantiates a new SyntheticsApiMultistepSubtestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsApiMultistepSubtestAttributesWithDefaults() *SyntheticsApiMultistepSubtestAttributes {
	this := SyntheticsApiMultistepSubtestAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsApiMultistepSubtestAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsApiMultistepSubtestAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsApiMultistepSubtestAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsApiMultistepSubtestAttributes) SetName(v string) {
	o.Name = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *SyntheticsApiMultistepSubtestAttributes) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsApiMultistepSubtestAttributes) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *SyntheticsApiMultistepSubtestAttributes) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *SyntheticsApiMultistepSubtestAttributes) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsApiMultistepSubtestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsApiMultistepSubtestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string `json:"name,omitempty"`
		PublicId *string `json:"public_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "public_id"})
	} else {
		return err
	}
	o.Name = all.Name
	o.PublicId = all.PublicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
