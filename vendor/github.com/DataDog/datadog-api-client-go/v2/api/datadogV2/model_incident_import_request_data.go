// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImportRequestData Incident data for an import request.
type IncidentImportRequestData struct {
	// The incident's attributes for an import request.
	Attributes IncidentImportRequestAttributes `json:"attributes"`
	// The relationships for an incident import request.
	Relationships *IncidentImportRelationships `json:"relationships,omitempty"`
	// Incident resource type.
	Type IncidentType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentImportRequestData instantiates a new IncidentImportRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentImportRequestData(attributes IncidentImportRequestAttributes, typeVar IncidentType) *IncidentImportRequestData {
	this := IncidentImportRequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewIncidentImportRequestDataWithDefaults instantiates a new IncidentImportRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentImportRequestDataWithDefaults() *IncidentImportRequestData {
	this := IncidentImportRequestData{}
	var typeVar IncidentType = INCIDENTTYPE_INCIDENTS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *IncidentImportRequestData) GetAttributes() IncidentImportRequestAttributes {
	if o == nil {
		var ret IncidentImportRequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentImportRequestData) GetAttributesOk() (*IncidentImportRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *IncidentImportRequestData) SetAttributes(v IncidentImportRequestAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *IncidentImportRequestData) GetRelationships() IncidentImportRelationships {
	if o == nil || o.Relationships == nil {
		var ret IncidentImportRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImportRequestData) GetRelationshipsOk() (*IncidentImportRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentImportRequestData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given IncidentImportRelationships and assigns it to the Relationships field.
func (o *IncidentImportRequestData) SetRelationships(v IncidentImportRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *IncidentImportRequestData) GetType() IncidentType {
	if o == nil {
		var ret IncidentType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentImportRequestData) GetTypeOk() (*IncidentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentImportRequestData) SetType(v IncidentType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentImportRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
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
func (o *IncidentImportRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentImportRequestAttributes `json:"attributes"`
		Relationships *IncidentImportRelationships     `json:"relationships,omitempty"`
		Type          *IncidentType                    `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
