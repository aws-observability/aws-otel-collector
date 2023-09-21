// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
=======
	"github.com/goccy/go-json"
>>>>>>> main

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationKeyRelationships Resources related to the application key.
type ApplicationKeyRelationships struct {
	// Relationship to user.
	OwnedBy *RelationshipToUser `json:"owned_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewApplicationKeyRelationships instantiates a new ApplicationKeyRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationKeyRelationships() *ApplicationKeyRelationships {
	this := ApplicationKeyRelationships{}
	return &this
}

// NewApplicationKeyRelationshipsWithDefaults instantiates a new ApplicationKeyRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationKeyRelationshipsWithDefaults() *ApplicationKeyRelationships {
	this := ApplicationKeyRelationships{}
	return &this
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *ApplicationKeyRelationships) GetOwnedBy() RelationshipToUser {
	if o == nil || o.OwnedBy == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationKeyRelationships) GetOwnedByOk() (*RelationshipToUser, bool) {
	if o == nil || o.OwnedBy == nil {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *ApplicationKeyRelationships) HasOwnedBy() bool {
	return o != nil && o.OwnedBy != nil
}

// SetOwnedBy gets a reference to the given RelationshipToUser and assigns it to the OwnedBy field.
func (o *ApplicationKeyRelationships) SetOwnedBy(v RelationshipToUser) {
	o.OwnedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationKeyRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.OwnedBy != nil {
		toSerialize["owned_by"] = o.OwnedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationKeyRelationships) UnmarshalJSON(bytes []byte) (err error) {
<<<<<<< HEAD
	raw := map[string]interface{}{}
=======
>>>>>>> main
	all := struct {
		OwnedBy *RelationshipToUser `json:"owned_by,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
<<<<<<< HEAD
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
=======
		return json.Unmarshal(bytes, &o.UnparsedObject)
>>>>>>> main
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"owned_by"})
	} else {
		return err
	}
<<<<<<< HEAD
	if all.OwnedBy != nil && all.OwnedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.OwnedBy = all.OwnedBy
=======

	hasInvalidField := false
	if all.OwnedBy != nil && all.OwnedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OwnedBy = all.OwnedBy

>>>>>>> main
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

<<<<<<< HEAD
=======
	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

>>>>>>> main
	return nil
}
