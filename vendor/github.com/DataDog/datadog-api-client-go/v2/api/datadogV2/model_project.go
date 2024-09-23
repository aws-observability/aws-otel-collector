// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Project A Project
type Project struct {
	// Project attributes
	Attributes ProjectAttributes `json:"attributes"`
	// The Project's identifier
	Id string `json:"id"`
	// Project relationships
	Relationships *ProjectRelationships `json:"relationships,omitempty"`
	// Project resource type
	Type ProjectResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProject instantiates a new Project object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProject(attributes ProjectAttributes, id string, typeVar ProjectResourceType) *Project {
	this := Project{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewProjectWithDefaults instantiates a new Project object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProjectWithDefaults() *Project {
	this := Project{}
	var typeVar ProjectResourceType = PROJECTRESOURCETYPE_PROJECT
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *Project) GetAttributes() ProjectAttributes {
	if o == nil {
		var ret ProjectAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *Project) GetAttributesOk() (*ProjectAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *Project) SetAttributes(v ProjectAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *Project) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Project) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *Project) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *Project) GetRelationships() ProjectRelationships {
	if o == nil || o.Relationships == nil {
		var ret ProjectRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetRelationshipsOk() (*ProjectRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *Project) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given ProjectRelationships and assigns it to the Relationships field.
func (o *Project) SetRelationships(v ProjectRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *Project) GetType() ProjectResourceType {
	if o == nil {
		var ret ProjectResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Project) GetTypeOk() (*ProjectResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Project) SetType(v ProjectResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
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
func (o *Project) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *ProjectAttributes    `json:"attributes"`
		Id            *string               `json:"id"`
		Relationships *ProjectRelationships `json:"relationships,omitempty"`
		Type          *ProjectResourceType  `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
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
