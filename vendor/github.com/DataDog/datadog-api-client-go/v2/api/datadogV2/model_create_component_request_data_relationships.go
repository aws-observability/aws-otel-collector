// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateComponentRequestDataRelationships The supported relationships for creating a component.
type CreateComponentRequestDataRelationships struct {
	// The group to create the component within.
	Group *CreateComponentRequestDataRelationshipsGroup `json:"group,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateComponentRequestDataRelationships instantiates a new CreateComponentRequestDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateComponentRequestDataRelationships() *CreateComponentRequestDataRelationships {
	this := CreateComponentRequestDataRelationships{}
	return &this
}

// NewCreateComponentRequestDataRelationshipsWithDefaults instantiates a new CreateComponentRequestDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateComponentRequestDataRelationshipsWithDefaults() *CreateComponentRequestDataRelationships {
	this := CreateComponentRequestDataRelationships{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *CreateComponentRequestDataRelationships) GetGroup() CreateComponentRequestDataRelationshipsGroup {
	if o == nil || o.Group == nil {
		var ret CreateComponentRequestDataRelationshipsGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateComponentRequestDataRelationships) GetGroupOk() (*CreateComponentRequestDataRelationshipsGroup, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *CreateComponentRequestDataRelationships) HasGroup() bool {
	return o != nil && o.Group != nil
}

// SetGroup gets a reference to the given CreateComponentRequestDataRelationshipsGroup and assigns it to the Group field.
func (o *CreateComponentRequestDataRelationships) SetGroup(v CreateComponentRequestDataRelationshipsGroup) {
	o.Group = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateComponentRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateComponentRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Group *CreateComponentRequestDataRelationshipsGroup `json:"group,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"group"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Group != nil && all.Group.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Group = all.Group

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
