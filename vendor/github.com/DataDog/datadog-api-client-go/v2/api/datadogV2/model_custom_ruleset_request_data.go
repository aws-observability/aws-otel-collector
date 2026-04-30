// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRulesetRequestData Data object for a custom ruleset create or update request.
type CustomRulesetRequestData struct {
	// Attributes for creating or updating a custom ruleset.
	Attributes *CustomRulesetRequestDataAttributes `json:"attributes,omitempty"`
	// Ruleset identifier
	Id *string `json:"id,omitempty"`
	// Resource type
	Type *CustomRulesetDataType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomRulesetRequestData instantiates a new CustomRulesetRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomRulesetRequestData() *CustomRulesetRequestData {
	this := CustomRulesetRequestData{}
	return &this
}

// NewCustomRulesetRequestDataWithDefaults instantiates a new CustomRulesetRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomRulesetRequestDataWithDefaults() *CustomRulesetRequestData {
	this := CustomRulesetRequestData{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomRulesetRequestData) GetAttributes() CustomRulesetRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret CustomRulesetRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRulesetRequestData) GetAttributesOk() (*CustomRulesetRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomRulesetRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CustomRulesetRequestDataAttributes and assigns it to the Attributes field.
func (o *CustomRulesetRequestData) SetAttributes(v CustomRulesetRequestDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomRulesetRequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRulesetRequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomRulesetRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomRulesetRequestData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomRulesetRequestData) GetType() CustomRulesetDataType {
	if o == nil || o.Type == nil {
		var ret CustomRulesetDataType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRulesetRequestData) GetTypeOk() (*CustomRulesetDataType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomRulesetRequestData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given CustomRulesetDataType and assigns it to the Type field.
func (o *CustomRulesetRequestData) SetType(v CustomRulesetDataType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomRulesetRequestData) MarshalJSON() ([]byte, error) {
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
func (o *CustomRulesetRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CustomRulesetRequestDataAttributes `json:"attributes,omitempty"`
		Id         *string                             `json:"id,omitempty"`
		Type       *CustomRulesetDataType              `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
