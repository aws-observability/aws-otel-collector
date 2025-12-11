// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReorderRuleResourceData The definition of `ReorderRuleResourceData` object.
type ReorderRuleResourceData struct {
	// The `ReorderRuleResourceData` `id`.
	Id *string `json:"id,omitempty"`
	// Arbitrary rule resource type.
	Type ReorderRuleResourceDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReorderRuleResourceData instantiates a new ReorderRuleResourceData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReorderRuleResourceData(typeVar ReorderRuleResourceDataType) *ReorderRuleResourceData {
	this := ReorderRuleResourceData{}
	this.Type = typeVar
	return &this
}

// NewReorderRuleResourceDataWithDefaults instantiates a new ReorderRuleResourceData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReorderRuleResourceDataWithDefaults() *ReorderRuleResourceData {
	this := ReorderRuleResourceData{}
	var typeVar ReorderRuleResourceDataType = REORDERRULERESOURCEDATATYPE_ARBITRARY_RULE
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReorderRuleResourceData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReorderRuleResourceData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReorderRuleResourceData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReorderRuleResourceData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *ReorderRuleResourceData) GetType() ReorderRuleResourceDataType {
	if o == nil {
		var ret ReorderRuleResourceDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReorderRuleResourceData) GetTypeOk() (*ReorderRuleResourceDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ReorderRuleResourceData) SetType(v ReorderRuleResourceDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReorderRuleResourceData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReorderRuleResourceData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                      `json:"id,omitempty"`
		Type *ReorderRuleResourceDataType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
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
