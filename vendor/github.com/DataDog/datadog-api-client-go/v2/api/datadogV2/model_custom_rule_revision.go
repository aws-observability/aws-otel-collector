// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRuleRevision A specific revision of a custom static analysis rule.
type CustomRuleRevision struct {
	// Attributes of a custom rule revision, including code, metadata, and test cases.
	Attributes CustomRuleRevisionAttributes `json:"attributes"`
	// Revision identifier
	Id string `json:"id"`
	// Resource type
	Type CustomRuleRevisionDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomRuleRevision instantiates a new CustomRuleRevision object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomRuleRevision(attributes CustomRuleRevisionAttributes, id string, typeVar CustomRuleRevisionDataType) *CustomRuleRevision {
	this := CustomRuleRevision{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewCustomRuleRevisionWithDefaults instantiates a new CustomRuleRevision object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomRuleRevisionWithDefaults() *CustomRuleRevision {
	this := CustomRuleRevision{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *CustomRuleRevision) GetAttributes() CustomRuleRevisionAttributes {
	if o == nil {
		var ret CustomRuleRevisionAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevision) GetAttributesOk() (*CustomRuleRevisionAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *CustomRuleRevision) SetAttributes(v CustomRuleRevisionAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *CustomRuleRevision) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevision) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CustomRuleRevision) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *CustomRuleRevision) GetType() CustomRuleRevisionDataType {
	if o == nil {
		var ret CustomRuleRevisionDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevision) GetTypeOk() (*CustomRuleRevisionDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomRuleRevision) SetType(v CustomRuleRevisionDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomRuleRevision) MarshalJSON() ([]byte, error) {
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
func (o *CustomRuleRevision) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CustomRuleRevisionAttributes `json:"attributes"`
		Id         *string                       `json:"id"`
		Type       *CustomRuleRevisionDataType   `json:"type"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
