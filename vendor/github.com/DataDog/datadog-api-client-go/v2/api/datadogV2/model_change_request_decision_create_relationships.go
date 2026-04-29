// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestDecisionCreateRelationships Relationships for creating a change request decision.
type ChangeRequestDecisionCreateRelationships struct {
	// Relationship to a user.
	RequestedUser *ChangeRequestUserRelationship `json:"requested_user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeRequestDecisionCreateRelationships instantiates a new ChangeRequestDecisionCreateRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeRequestDecisionCreateRelationships() *ChangeRequestDecisionCreateRelationships {
	this := ChangeRequestDecisionCreateRelationships{}
	return &this
}

// NewChangeRequestDecisionCreateRelationshipsWithDefaults instantiates a new ChangeRequestDecisionCreateRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeRequestDecisionCreateRelationshipsWithDefaults() *ChangeRequestDecisionCreateRelationships {
	this := ChangeRequestDecisionCreateRelationships{}
	return &this
}

// GetRequestedUser returns the RequestedUser field value if set, zero value otherwise.
func (o *ChangeRequestDecisionCreateRelationships) GetRequestedUser() ChangeRequestUserRelationship {
	if o == nil || o.RequestedUser == nil {
		var ret ChangeRequestUserRelationship
		return ret
	}
	return *o.RequestedUser
}

// GetRequestedUserOk returns a tuple with the RequestedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestDecisionCreateRelationships) GetRequestedUserOk() (*ChangeRequestUserRelationship, bool) {
	if o == nil || o.RequestedUser == nil {
		return nil, false
	}
	return o.RequestedUser, true
}

// HasRequestedUser returns a boolean if a field has been set.
func (o *ChangeRequestDecisionCreateRelationships) HasRequestedUser() bool {
	return o != nil && o.RequestedUser != nil
}

// SetRequestedUser gets a reference to the given ChangeRequestUserRelationship and assigns it to the RequestedUser field.
func (o *ChangeRequestDecisionCreateRelationships) SetRequestedUser(v ChangeRequestUserRelationship) {
	o.RequestedUser = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeRequestDecisionCreateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RequestedUser != nil {
		toSerialize["requested_user"] = o.RequestedUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeRequestDecisionCreateRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RequestedUser *ChangeRequestUserRelationship `json:"requested_user,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"requested_user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.RequestedUser != nil && all.RequestedUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RequestedUser = all.RequestedUser

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
