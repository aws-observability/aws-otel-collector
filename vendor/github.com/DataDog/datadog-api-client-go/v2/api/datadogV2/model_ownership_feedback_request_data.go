// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipFeedbackRequestData The data wrapper for an ownership feedback request.
type OwnershipFeedbackRequestData struct {
	// The attributes of an ownership feedback request.
	Attributes OwnershipFeedbackRequestAttributes `json:"attributes"`
	// The type of the ownership feedback request resource. The value should always be `ownership_feedback`.
	Type OwnershipFeedbackType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOwnershipFeedbackRequestData instantiates a new OwnershipFeedbackRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOwnershipFeedbackRequestData(attributes OwnershipFeedbackRequestAttributes, typeVar OwnershipFeedbackType) *OwnershipFeedbackRequestData {
	this := OwnershipFeedbackRequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewOwnershipFeedbackRequestDataWithDefaults instantiates a new OwnershipFeedbackRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOwnershipFeedbackRequestDataWithDefaults() *OwnershipFeedbackRequestData {
	this := OwnershipFeedbackRequestData{}
	var typeVar OwnershipFeedbackType = OWNERSHIPFEEDBACKTYPE_OWNERSHIP_FEEDBACK
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *OwnershipFeedbackRequestData) GetAttributes() OwnershipFeedbackRequestAttributes {
	if o == nil {
		var ret OwnershipFeedbackRequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OwnershipFeedbackRequestData) GetAttributesOk() (*OwnershipFeedbackRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *OwnershipFeedbackRequestData) SetAttributes(v OwnershipFeedbackRequestAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *OwnershipFeedbackRequestData) GetType() OwnershipFeedbackType {
	if o == nil {
		var ret OwnershipFeedbackType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OwnershipFeedbackRequestData) GetTypeOk() (*OwnershipFeedbackType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OwnershipFeedbackRequestData) SetType(v OwnershipFeedbackType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OwnershipFeedbackRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OwnershipFeedbackRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *OwnershipFeedbackRequestAttributes `json:"attributes"`
		Type       *OwnershipFeedbackType              `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
