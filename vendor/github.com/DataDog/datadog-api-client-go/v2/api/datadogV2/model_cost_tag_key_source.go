// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagKeySource A Cloud Cost Management tag key paired with the sources that produced it.
type CostTagKeySource struct {
	// Attributes of a Cloud Cost Management tag source.
	Attributes CostTagKeySourceAttributes `json:"attributes"`
	// The tag key identifier. Equal to the empty-tag sentinel `__empty_tag_key__` when the tag key is empty.
	Id string `json:"id"`
	// Type of the Cloud Cost Management tag source resource.
	Type CostTagKeySourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostTagKeySource instantiates a new CostTagKeySource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostTagKeySource(attributes CostTagKeySourceAttributes, id string, typeVar CostTagKeySourceType) *CostTagKeySource {
	this := CostTagKeySource{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewCostTagKeySourceWithDefaults instantiates a new CostTagKeySource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostTagKeySourceWithDefaults() *CostTagKeySource {
	this := CostTagKeySource{}
	var typeVar CostTagKeySourceType = COSTTAGKEYSOURCETYPE_COST_TAG_KEY_SOURCE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *CostTagKeySource) GetAttributes() CostTagKeySourceAttributes {
	if o == nil {
		var ret CostTagKeySourceAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CostTagKeySource) GetAttributesOk() (*CostTagKeySourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *CostTagKeySource) SetAttributes(v CostTagKeySourceAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *CostTagKeySource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CostTagKeySource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CostTagKeySource) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *CostTagKeySource) GetType() CostTagKeySourceType {
	if o == nil {
		var ret CostTagKeySourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CostTagKeySource) GetTypeOk() (*CostTagKeySourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CostTagKeySource) SetType(v CostTagKeySourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostTagKeySource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostTagKeySource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CostTagKeySourceAttributes `json:"attributes"`
		Id         *string                     `json:"id"`
		Type       *CostTagKeySourceType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
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
