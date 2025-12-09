// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ShiftDataRelationships Relationships for an on-call shift.
type ShiftDataRelationships struct {
	// Defines the relationship between a shift and the user who is working that shift.
	User *ShiftDataRelationshipsUser `json:"user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewShiftDataRelationships instantiates a new ShiftDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewShiftDataRelationships() *ShiftDataRelationships {
	this := ShiftDataRelationships{}
	return &this
}

// NewShiftDataRelationshipsWithDefaults instantiates a new ShiftDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewShiftDataRelationshipsWithDefaults() *ShiftDataRelationships {
	this := ShiftDataRelationships{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ShiftDataRelationships) GetUser() ShiftDataRelationshipsUser {
	if o == nil || o.User == nil {
		var ret ShiftDataRelationshipsUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShiftDataRelationships) GetUserOk() (*ShiftDataRelationshipsUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ShiftDataRelationships) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given ShiftDataRelationshipsUser and assigns it to the User field.
func (o *ShiftDataRelationships) SetUser(v ShiftDataRelationshipsUser) {
	o.User = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ShiftDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ShiftDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		User *ShiftDataRelationshipsUser `json:"user,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.User != nil && all.User.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.User = all.User

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
