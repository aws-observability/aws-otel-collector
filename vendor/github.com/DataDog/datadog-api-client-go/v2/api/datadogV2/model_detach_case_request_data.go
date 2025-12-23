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
	// The unique identifier of the detachment request.
	Id *string `json:"id,omitempty"`
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

// GetId returns the Id field value if set, zero value otherwise.
func (o *DetachCaseRequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetachCaseRequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DetachCaseRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DetachCaseRequestData) SetId(v string) {
	o.Id = &v
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
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
		Id            *string                             `json:"id,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
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
