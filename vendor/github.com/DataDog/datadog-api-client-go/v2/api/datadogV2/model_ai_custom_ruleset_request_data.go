// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiCustomRulesetRequestData Request data for creating an AI custom ruleset.
type AiCustomRulesetRequestData struct {
	// Attributes for creating an AI custom ruleset.
	Attributes *AiCustomRulesetRequestAttributes `json:"attributes,omitempty"`
	// The ruleset identifier, which must match the name.
	Id *string `json:"id,omitempty"`
	// AI custom ruleset resource type.
	Type *AiCustomRulesetDataType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiCustomRulesetRequestData instantiates a new AiCustomRulesetRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiCustomRulesetRequestData() *AiCustomRulesetRequestData {
	this := AiCustomRulesetRequestData{}
	return &this
}

// NewAiCustomRulesetRequestDataWithDefaults instantiates a new AiCustomRulesetRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiCustomRulesetRequestDataWithDefaults() *AiCustomRulesetRequestData {
	this := AiCustomRulesetRequestData{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AiCustomRulesetRequestData) GetAttributes() AiCustomRulesetRequestAttributes {
	if o == nil || o.Attributes == nil {
		var ret AiCustomRulesetRequestAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiCustomRulesetRequestData) GetAttributesOk() (*AiCustomRulesetRequestAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AiCustomRulesetRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given AiCustomRulesetRequestAttributes and assigns it to the Attributes field.
func (o *AiCustomRulesetRequestData) SetAttributes(v AiCustomRulesetRequestAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AiCustomRulesetRequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiCustomRulesetRequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AiCustomRulesetRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AiCustomRulesetRequestData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AiCustomRulesetRequestData) GetType() AiCustomRulesetDataType {
	if o == nil || o.Type == nil {
		var ret AiCustomRulesetDataType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiCustomRulesetRequestData) GetTypeOk() (*AiCustomRulesetDataType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AiCustomRulesetRequestData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given AiCustomRulesetDataType and assigns it to the Type field.
func (o *AiCustomRulesetRequestData) SetType(v AiCustomRulesetDataType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiCustomRulesetRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiCustomRulesetRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AiCustomRulesetRequestAttributes `json:"attributes,omitempty"`
		Id         *string                           `json:"id,omitempty"`
		Type       *AiCustomRulesetDataType          `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
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
