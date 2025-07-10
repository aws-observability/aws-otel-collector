// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleMember Represents a single member entry in a schedule, referencing a specific user.
type ScheduleMember struct {
	// The unique identifier for this schedule member.
	Id *string `json:"id,omitempty"`
	// Defines relationships for a schedule member, primarily referencing a single user.
	Relationships *ScheduleMemberRelationships `json:"relationships,omitempty"`
	// Schedule Members resource type.
	Type ScheduleMemberType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleMember instantiates a new ScheduleMember object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleMember(typeVar ScheduleMemberType) *ScheduleMember {
	this := ScheduleMember{}
	this.Type = typeVar
	return &this
}

// NewScheduleMemberWithDefaults instantiates a new ScheduleMember object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleMemberWithDefaults() *ScheduleMember {
	this := ScheduleMember{}
	var typeVar ScheduleMemberType = SCHEDULEMEMBERTYPE_MEMBERS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScheduleMember) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleMember) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScheduleMember) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScheduleMember) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ScheduleMember) GetRelationships() ScheduleMemberRelationships {
	if o == nil || o.Relationships == nil {
		var ret ScheduleMemberRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleMember) GetRelationshipsOk() (*ScheduleMemberRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ScheduleMember) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given ScheduleMemberRelationships and assigns it to the Relationships field.
func (o *ScheduleMember) SetRelationships(v ScheduleMemberRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *ScheduleMember) GetType() ScheduleMemberType {
	if o == nil {
		var ret ScheduleMemberType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScheduleMember) GetTypeOk() (*ScheduleMemberType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ScheduleMember) SetType(v ScheduleMemberType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleMember) MarshalJSON() ([]byte, error) {
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
func (o *ScheduleMember) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id            *string                      `json:"id,omitempty"`
		Relationships *ScheduleMemberRelationships `json:"relationships,omitempty"`
		Type          *ScheduleMemberType          `json:"type"`
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
