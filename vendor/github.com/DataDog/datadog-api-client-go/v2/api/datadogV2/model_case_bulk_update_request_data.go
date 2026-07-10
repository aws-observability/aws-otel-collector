// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseBulkUpdateRequestData Data object wrapping the bulk update type and attributes.
type CaseBulkUpdateRequestData struct {
	// Attributes for the bulk update, specifying which cases to update and the action to apply.
	Attributes CaseBulkUpdateRequestAttributes `json:"attributes"`
	// JSON:API resource type for bulk case operations.
	Type CaseBulkResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseBulkUpdateRequestData instantiates a new CaseBulkUpdateRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseBulkUpdateRequestData(attributes CaseBulkUpdateRequestAttributes, typeVar CaseBulkResourceType) *CaseBulkUpdateRequestData {
	this := CaseBulkUpdateRequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewCaseBulkUpdateRequestDataWithDefaults instantiates a new CaseBulkUpdateRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseBulkUpdateRequestDataWithDefaults() *CaseBulkUpdateRequestData {
	this := CaseBulkUpdateRequestData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *CaseBulkUpdateRequestData) GetAttributes() CaseBulkUpdateRequestAttributes {
	if o == nil {
		var ret CaseBulkUpdateRequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CaseBulkUpdateRequestData) GetAttributesOk() (*CaseBulkUpdateRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *CaseBulkUpdateRequestData) SetAttributes(v CaseBulkUpdateRequestAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *CaseBulkUpdateRequestData) GetType() CaseBulkResourceType {
	if o == nil {
		var ret CaseBulkResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CaseBulkUpdateRequestData) GetTypeOk() (*CaseBulkResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CaseBulkUpdateRequestData) SetType(v CaseBulkResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseBulkUpdateRequestData) MarshalJSON() ([]byte, error) {
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
func (o *CaseBulkUpdateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CaseBulkUpdateRequestAttributes `json:"attributes"`
		Type       *CaseBulkResourceType            `json:"type"`
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
