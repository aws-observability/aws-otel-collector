// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomAttributeConfigCreate Custom attribute config
type CustomAttributeConfigCreate struct {
	// Custom attribute config resource attributes
	Attributes CustomAttributeConfigAttributesCreate `json:"attributes"`
	// Custom attributes config JSON:API resource type
	Type CustomAttributeConfigResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomAttributeConfigCreate instantiates a new CustomAttributeConfigCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomAttributeConfigCreate(attributes CustomAttributeConfigAttributesCreate, typeVar CustomAttributeConfigResourceType) *CustomAttributeConfigCreate {
	this := CustomAttributeConfigCreate{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewCustomAttributeConfigCreateWithDefaults instantiates a new CustomAttributeConfigCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomAttributeConfigCreateWithDefaults() *CustomAttributeConfigCreate {
	this := CustomAttributeConfigCreate{}
	var typeVar CustomAttributeConfigResourceType = CUSTOMATTRIBUTECONFIGRESOURCETYPE_CUSTOM_ATTRIBUTE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *CustomAttributeConfigCreate) GetAttributes() CustomAttributeConfigAttributesCreate {
	if o == nil {
		var ret CustomAttributeConfigAttributesCreate
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfigCreate) GetAttributesOk() (*CustomAttributeConfigAttributesCreate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *CustomAttributeConfigCreate) SetAttributes(v CustomAttributeConfigAttributesCreate) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *CustomAttributeConfigCreate) GetType() CustomAttributeConfigResourceType {
	if o == nil {
		var ret CustomAttributeConfigResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfigCreate) GetTypeOk() (*CustomAttributeConfigResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomAttributeConfigCreate) SetType(v CustomAttributeConfigResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomAttributeConfigCreate) MarshalJSON() ([]byte, error) {
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
func (o *CustomAttributeConfigCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CustomAttributeConfigAttributesCreate `json:"attributes"`
		Type       *CustomAttributeConfigResourceType     `json:"type"`
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
