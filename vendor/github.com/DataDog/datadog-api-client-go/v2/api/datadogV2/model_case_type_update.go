// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseTypeUpdate Data object for updating a case type.
type CaseTypeUpdate struct {
	// Attributes of a case type, which define a classification category for cases. Organizations use case types to model different workflows (for example, Security Incident, Bug Report, Change Request).
	Attributes *CaseTypeResourceAttributes `json:"attributes,omitempty"`
	// JSON:API resource type for case types.
	Type CaseTypeResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseTypeUpdate instantiates a new CaseTypeUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseTypeUpdate(typeVar CaseTypeResourceType) *CaseTypeUpdate {
	this := CaseTypeUpdate{}
	this.Type = typeVar
	return &this
}

// NewCaseTypeUpdateWithDefaults instantiates a new CaseTypeUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseTypeUpdateWithDefaults() *CaseTypeUpdate {
	this := CaseTypeUpdate{}
	var typeVar CaseTypeResourceType = CASETYPERESOURCETYPE_CASE_TYPE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CaseTypeUpdate) GetAttributes() CaseTypeResourceAttributes {
	if o == nil || o.Attributes == nil {
		var ret CaseTypeResourceAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTypeUpdate) GetAttributesOk() (*CaseTypeResourceAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CaseTypeUpdate) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CaseTypeResourceAttributes and assigns it to the Attributes field.
func (o *CaseTypeUpdate) SetAttributes(v CaseTypeResourceAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value.
func (o *CaseTypeUpdate) GetType() CaseTypeResourceType {
	if o == nil {
		var ret CaseTypeResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CaseTypeUpdate) GetTypeOk() (*CaseTypeResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CaseTypeUpdate) SetType(v CaseTypeResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseTypeUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseTypeUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CaseTypeResourceAttributes `json:"attributes,omitempty"`
		Type       *CaseTypeResourceType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
