// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestUpdateData Data object to update a change request.
type ChangeRequestUpdateData struct {
	// Attributes for updating a change request.
	Attributes *ChangeRequestUpdateAttributes `json:"attributes,omitempty"`
	// Relationships for updating a change request.
	Relationships *ChangeRequestUpdateRelationships `json:"relationships,omitempty"`
	// Change request resource type.
	Type ChangeRequestResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeRequestUpdateData instantiates a new ChangeRequestUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeRequestUpdateData(typeVar ChangeRequestResourceType) *ChangeRequestUpdateData {
	this := ChangeRequestUpdateData{}
	this.Type = typeVar
	return &this
}

// NewChangeRequestUpdateDataWithDefaults instantiates a new ChangeRequestUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeRequestUpdateDataWithDefaults() *ChangeRequestUpdateData {
	this := ChangeRequestUpdateData{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ChangeRequestUpdateData) GetAttributes() ChangeRequestUpdateAttributes {
	if o == nil || o.Attributes == nil {
		var ret ChangeRequestUpdateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestUpdateData) GetAttributesOk() (*ChangeRequestUpdateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ChangeRequestUpdateData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ChangeRequestUpdateAttributes and assigns it to the Attributes field.
func (o *ChangeRequestUpdateData) SetAttributes(v ChangeRequestUpdateAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ChangeRequestUpdateData) GetRelationships() ChangeRequestUpdateRelationships {
	if o == nil || o.Relationships == nil {
		var ret ChangeRequestUpdateRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestUpdateData) GetRelationshipsOk() (*ChangeRequestUpdateRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ChangeRequestUpdateData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given ChangeRequestUpdateRelationships and assigns it to the Relationships field.
func (o *ChangeRequestUpdateData) SetRelationships(v ChangeRequestUpdateRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *ChangeRequestUpdateData) GetType() ChangeRequestResourceType {
	if o == nil {
		var ret ChangeRequestResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestUpdateData) GetTypeOk() (*ChangeRequestResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ChangeRequestUpdateData) SetType(v ChangeRequestResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeRequestUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
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
func (o *ChangeRequestUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *ChangeRequestUpdateAttributes    `json:"attributes,omitempty"`
		Relationships *ChangeRequestUpdateRelationships `json:"relationships,omitempty"`
		Type          *ChangeRequestResourceType        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
