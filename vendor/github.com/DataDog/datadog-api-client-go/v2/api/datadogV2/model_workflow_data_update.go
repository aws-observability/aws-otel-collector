// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowDataUpdate Data related to the workflow being updated.
type WorkflowDataUpdate struct {
	// The definition of `WorkflowDataUpdateAttributes` object.
	Attributes WorkflowDataUpdateAttributes `json:"attributes"`
	// The workflow identifier
	Id *string `json:"id,omitempty"`
	// The definition of `WorkflowDataRelationships` object.
	Relationships *WorkflowDataRelationships `json:"relationships,omitempty"`
	// The definition of `WorkflowDataType` object.
	Type WorkflowDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowDataUpdate instantiates a new WorkflowDataUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowDataUpdate(attributes WorkflowDataUpdateAttributes, typeVar WorkflowDataType) *WorkflowDataUpdate {
	this := WorkflowDataUpdate{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewWorkflowDataUpdateWithDefaults instantiates a new WorkflowDataUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowDataUpdateWithDefaults() *WorkflowDataUpdate {
	this := WorkflowDataUpdate{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *WorkflowDataUpdate) GetAttributes() WorkflowDataUpdateAttributes {
	if o == nil {
		var ret WorkflowDataUpdateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WorkflowDataUpdate) GetAttributesOk() (*WorkflowDataUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *WorkflowDataUpdate) SetAttributes(v WorkflowDataUpdateAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowDataUpdate) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDataUpdate) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowDataUpdate) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowDataUpdate) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WorkflowDataUpdate) GetRelationships() WorkflowDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret WorkflowDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDataUpdate) GetRelationshipsOk() (*WorkflowDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WorkflowDataUpdate) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given WorkflowDataRelationships and assigns it to the Relationships field.
func (o *WorkflowDataUpdate) SetRelationships(v WorkflowDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *WorkflowDataUpdate) GetType() WorkflowDataType {
	if o == nil {
		var ret WorkflowDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WorkflowDataUpdate) GetTypeOk() (*WorkflowDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WorkflowDataUpdate) SetType(v WorkflowDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowDataUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
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
func (o *WorkflowDataUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *WorkflowDataUpdateAttributes `json:"attributes"`
		Id            *string                       `json:"id,omitempty"`
		Relationships *WorkflowDataRelationships    `json:"relationships,omitempty"`
		Type          *WorkflowDataType             `json:"type"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
