// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORAListFailuresRequestData The JSON:API data.
type DORAListFailuresRequestData struct {
	// Attributes to get a list of failures.
	Attributes DORAListFailuresRequestAttributes `json:"attributes"`
	// The definition of `DORAListFailuresRequestDataType` object.
	Type *DORAListFailuresRequestDataType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDORAListFailuresRequestData instantiates a new DORAListFailuresRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORAListFailuresRequestData(attributes DORAListFailuresRequestAttributes) *DORAListFailuresRequestData {
	this := DORAListFailuresRequestData{}
	this.Attributes = attributes
	return &this
}

// NewDORAListFailuresRequestDataWithDefaults instantiates a new DORAListFailuresRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORAListFailuresRequestDataWithDefaults() *DORAListFailuresRequestData {
	this := DORAListFailuresRequestData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *DORAListFailuresRequestData) GetAttributes() DORAListFailuresRequestAttributes {
	if o == nil {
		var ret DORAListFailuresRequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DORAListFailuresRequestData) GetAttributesOk() (*DORAListFailuresRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *DORAListFailuresRequestData) SetAttributes(v DORAListFailuresRequestAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DORAListFailuresRequestData) GetType() DORAListFailuresRequestDataType {
	if o == nil || o.Type == nil {
		var ret DORAListFailuresRequestDataType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAListFailuresRequestData) GetTypeOk() (*DORAListFailuresRequestDataType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DORAListFailuresRequestData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given DORAListFailuresRequestDataType and assigns it to the Type field.
func (o *DORAListFailuresRequestData) SetType(v DORAListFailuresRequestDataType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORAListFailuresRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DORAListFailuresRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *DORAListFailuresRequestAttributes `json:"attributes"`
		Type       *DORAListFailuresRequestDataType   `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
