// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AssigneeRequestData Data of the assignee request.
type AssigneeRequestData struct {
	// Attributes of the assignee request.
	Attributes *AssigneeRequestDataAttributes `json:"attributes,omitempty"`
	// Unique identifier of the assignee request.
	Id *string `json:"id,omitempty"`
	// Relationships of the assignee request.
	Relationships AssigneeRequestDataRelationships `json:"relationships"`
	// Assignee resource type.
	Type AssigneeDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAssigneeRequestData instantiates a new AssigneeRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssigneeRequestData(relationships AssigneeRequestDataRelationships, typeVar AssigneeDataType) *AssigneeRequestData {
	this := AssigneeRequestData{}
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewAssigneeRequestDataWithDefaults instantiates a new AssigneeRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAssigneeRequestDataWithDefaults() *AssigneeRequestData {
	this := AssigneeRequestData{}
	var typeVar AssigneeDataType = ASSIGNEEDATATYPE_ASSIGNEE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AssigneeRequestData) GetAttributes() AssigneeRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret AssigneeRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssigneeRequestData) GetAttributesOk() (*AssigneeRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AssigneeRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given AssigneeRequestDataAttributes and assigns it to the Attributes field.
func (o *AssigneeRequestData) SetAttributes(v AssigneeRequestDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssigneeRequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssigneeRequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssigneeRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssigneeRequestData) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value.
func (o *AssigneeRequestData) GetRelationships() AssigneeRequestDataRelationships {
	if o == nil {
		var ret AssigneeRequestDataRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AssigneeRequestData) GetRelationshipsOk() (*AssigneeRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *AssigneeRequestData) SetRelationships(v AssigneeRequestDataRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *AssigneeRequestData) GetType() AssigneeDataType {
	if o == nil {
		var ret AssigneeDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AssigneeRequestData) GetTypeOk() (*AssigneeDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AssigneeRequestData) SetType(v AssigneeDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AssigneeRequestData) MarshalJSON() ([]byte, error) {
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
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AssigneeRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *AssigneeRequestDataAttributes    `json:"attributes,omitempty"`
		Id            *string                           `json:"id,omitempty"`
		Relationships *AssigneeRequestDataRelationships `json:"relationships"`
		Type          *AssigneeDataType                 `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = *all.Relationships
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
