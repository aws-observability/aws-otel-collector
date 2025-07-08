// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleUpdateRequestData Contains all data needed to update an existing schedule, including its attributes (such as name and time zone) and any relationships to teams.
type ScheduleUpdateRequestData struct {
	// Defines the updatable attributes for a schedule, such as name, time zone, and layers.
	Attributes ScheduleUpdateRequestDataAttributes `json:"attributes"`
	// The ID of the schedule to be updated.
	Id string `json:"id"`
	// Houses relationships for the schedule update, typically referencing teams.
	Relationships *ScheduleUpdateRequestDataRelationships `json:"relationships,omitempty"`
	// Schedules resource type.
	Type ScheduleUpdateRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleUpdateRequestData instantiates a new ScheduleUpdateRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleUpdateRequestData(attributes ScheduleUpdateRequestDataAttributes, id string, typeVar ScheduleUpdateRequestDataType) *ScheduleUpdateRequestData {
	this := ScheduleUpdateRequestData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewScheduleUpdateRequestDataWithDefaults instantiates a new ScheduleUpdateRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleUpdateRequestDataWithDefaults() *ScheduleUpdateRequestData {
	this := ScheduleUpdateRequestData{}
	var typeVar ScheduleUpdateRequestDataType = SCHEDULEUPDATEREQUESTDATATYPE_SCHEDULES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ScheduleUpdateRequestData) GetAttributes() ScheduleUpdateRequestDataAttributes {
	if o == nil {
		var ret ScheduleUpdateRequestDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestData) GetAttributesOk() (*ScheduleUpdateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ScheduleUpdateRequestData) SetAttributes(v ScheduleUpdateRequestDataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *ScheduleUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ScheduleUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestData) GetRelationships() ScheduleUpdateRequestDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ScheduleUpdateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestData) GetRelationshipsOk() (*ScheduleUpdateRequestDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given ScheduleUpdateRequestDataRelationships and assigns it to the Relationships field.
func (o *ScheduleUpdateRequestData) SetRelationships(v ScheduleUpdateRequestDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *ScheduleUpdateRequestData) GetType() ScheduleUpdateRequestDataType {
	if o == nil {
		var ret ScheduleUpdateRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestData) GetTypeOk() (*ScheduleUpdateRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ScheduleUpdateRequestData) SetType(v ScheduleUpdateRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleUpdateRequestData) MarshalJSON() ([]byte, error) {
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
func (o *ScheduleUpdateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *ScheduleUpdateRequestDataAttributes    `json:"attributes"`
		Id            *string                                 `json:"id"`
		Relationships *ScheduleUpdateRequestDataRelationships `json:"relationships,omitempty"`
		Type          *ScheduleUpdateRequestDataType          `json:"type"`
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
