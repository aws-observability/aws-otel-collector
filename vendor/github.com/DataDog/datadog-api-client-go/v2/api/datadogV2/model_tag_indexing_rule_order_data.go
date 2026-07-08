// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleOrderData Data object for the reorder operation.
type TagIndexingRuleOrderData struct {
	// Attributes for the reorder operation.
	Attributes TagIndexingRuleOrderAttributes `json:"attributes"`
	// The tag indexing rule resource type.
	Type TagIndexingRuleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagIndexingRuleOrderData instantiates a new TagIndexingRuleOrderData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagIndexingRuleOrderData(attributes TagIndexingRuleOrderAttributes, typeVar TagIndexingRuleType) *TagIndexingRuleOrderData {
	this := TagIndexingRuleOrderData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewTagIndexingRuleOrderDataWithDefaults instantiates a new TagIndexingRuleOrderData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagIndexingRuleOrderDataWithDefaults() *TagIndexingRuleOrderData {
	this := TagIndexingRuleOrderData{}
	var typeVar TagIndexingRuleType = TAGINDEXINGRULETYPE_TAG_INDEXING_RULES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *TagIndexingRuleOrderData) GetAttributes() TagIndexingRuleOrderAttributes {
	if o == nil {
		var ret TagIndexingRuleOrderAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleOrderData) GetAttributesOk() (*TagIndexingRuleOrderAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *TagIndexingRuleOrderData) SetAttributes(v TagIndexingRuleOrderAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *TagIndexingRuleOrderData) GetType() TagIndexingRuleType {
	if o == nil {
		var ret TagIndexingRuleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleOrderData) GetTypeOk() (*TagIndexingRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TagIndexingRuleOrderData) SetType(v TagIndexingRuleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagIndexingRuleOrderData) MarshalJSON() ([]byte, error) {
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
func (o *TagIndexingRuleOrderData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *TagIndexingRuleOrderAttributes `json:"attributes"`
		Type       *TagIndexingRuleType            `json:"type"`
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
