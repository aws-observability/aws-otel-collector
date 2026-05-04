// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DetachCaseRequestData Data for detaching security findings from their case.
type DetachCaseRequestData struct {
	// Relationships detaching security findings from their case.
	Relationships *DetachCaseRequestDataRelationships `json:"relationships,omitempty"`
	// Cases resource type.
	Type CaseDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDetachCaseRequestData instantiates a new DetachCaseRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDetachCaseRequestData(typeVar CaseDataType) *DetachCaseRequestData {
	this := DetachCaseRequestData{}
	this.Type = typeVar
	return &this
}

// NewDetachCaseRequestDataWithDefaults instantiates a new DetachCaseRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDetachCaseRequestDataWithDefaults() *DetachCaseRequestData {
	this := DetachCaseRequestData{}
	var typeVar CaseDataType = CASEDATATYPE_CASES
	this.Type = typeVar
	return &this
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *DetachCaseRequestData) GetRelationships() DetachCaseRequestDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret DetachCaseRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetachCaseRequestData) GetRelationshipsOk() (*DetachCaseRequestDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *DetachCaseRequestData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given DetachCaseRequestDataRelationships and assigns it to the Relationships field.
func (o *DetachCaseRequestData) SetRelationships(v DetachCaseRequestDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *DetachCaseRequestData) GetType() CaseDataType {
	if o == nil {
		var ret CaseDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DetachCaseRequestData) GetTypeOk() (*CaseDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DetachCaseRequestData) SetType(v CaseDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DetachCaseRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DetachCaseRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Relationships *DetachCaseRequestDataRelationships `json:"relationships,omitempty"`
		Type          *CaseDataType                       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
