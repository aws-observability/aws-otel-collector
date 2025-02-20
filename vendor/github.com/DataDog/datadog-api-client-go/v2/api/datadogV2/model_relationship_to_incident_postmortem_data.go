// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RelationshipToIncidentPostmortemData The postmortem relationship data.
type RelationshipToIncidentPostmortemData struct {
	// A unique identifier that represents the postmortem.
	Id string `json:"id"`
	// Incident postmortem resource type.
	Type IncidentPostmortemType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRelationshipToIncidentPostmortemData instantiates a new RelationshipToIncidentPostmortemData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRelationshipToIncidentPostmortemData(id string, typeVar IncidentPostmortemType) *RelationshipToIncidentPostmortemData {
	this := RelationshipToIncidentPostmortemData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewRelationshipToIncidentPostmortemDataWithDefaults instantiates a new RelationshipToIncidentPostmortemData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRelationshipToIncidentPostmortemDataWithDefaults() *RelationshipToIncidentPostmortemData {
	this := RelationshipToIncidentPostmortemData{}
	var typeVar IncidentPostmortemType = INCIDENTPOSTMORTEMTYPE_INCIDENT_POSTMORTEMS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *RelationshipToIncidentPostmortemData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RelationshipToIncidentPostmortemData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *RelationshipToIncidentPostmortemData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *RelationshipToIncidentPostmortemData) GetType() IncidentPostmortemType {
	if o == nil {
		var ret IncidentPostmortemType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RelationshipToIncidentPostmortemData) GetTypeOk() (*IncidentPostmortemType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RelationshipToIncidentPostmortemData) SetType(v IncidentPostmortemType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RelationshipToIncidentPostmortemData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RelationshipToIncidentPostmortemData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                 `json:"id"`
		Type *IncidentPostmortemType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
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
