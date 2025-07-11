// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleCreateRequestData The core data wrapper for creating a schedule, encompassing attributes, relationships, and the resource type.
type ScheduleCreateRequestData struct {
	// Describes the main attributes for creating a new schedule, including name, layers, and time zone.
	Attributes ScheduleCreateRequestDataAttributes `json:"attributes"`
	// Gathers relationship objects for the schedule creation request, including the teams to associate.
	Relationships *ScheduleCreateRequestDataRelationships `json:"relationships,omitempty"`
	// Schedules resource type.
	Type ScheduleCreateRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleCreateRequestData instantiates a new ScheduleCreateRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleCreateRequestData(attributes ScheduleCreateRequestDataAttributes, typeVar ScheduleCreateRequestDataType) *ScheduleCreateRequestData {
	this := ScheduleCreateRequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewScheduleCreateRequestDataWithDefaults instantiates a new ScheduleCreateRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleCreateRequestDataWithDefaults() *ScheduleCreateRequestData {
	this := ScheduleCreateRequestData{}
	var typeVar ScheduleCreateRequestDataType = SCHEDULECREATEREQUESTDATATYPE_SCHEDULES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ScheduleCreateRequestData) GetAttributes() ScheduleCreateRequestDataAttributes {
	if o == nil {
		var ret ScheduleCreateRequestDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestData) GetAttributesOk() (*ScheduleCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ScheduleCreateRequestData) SetAttributes(v ScheduleCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ScheduleCreateRequestData) GetRelationships() ScheduleCreateRequestDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ScheduleCreateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestData) GetRelationshipsOk() (*ScheduleCreateRequestDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ScheduleCreateRequestData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given ScheduleCreateRequestDataRelationships and assigns it to the Relationships field.
func (o *ScheduleCreateRequestData) SetRelationships(v ScheduleCreateRequestDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *ScheduleCreateRequestData) GetType() ScheduleCreateRequestDataType {
	if o == nil {
		var ret ScheduleCreateRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestData) GetTypeOk() (*ScheduleCreateRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ScheduleCreateRequestData) SetType(v ScheduleCreateRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleCreateRequestData) MarshalJSON() ([]byte, error) {
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
func (o *ScheduleCreateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *ScheduleCreateRequestDataAttributes    `json:"attributes"`
		Relationships *ScheduleCreateRequestDataRelationships `json:"relationships,omitempty"`
		Type          *ScheduleCreateRequestDataType          `json:"type"`
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
