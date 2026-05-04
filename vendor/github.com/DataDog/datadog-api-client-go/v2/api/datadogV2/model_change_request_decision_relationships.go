// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestDecisionRelationships Relationships of a change request decision.
type ChangeRequestDecisionRelationships struct {
	// Relationship to a user.
	ModifiedBy ChangeRequestUserRelationship `json:"modified_by"`
	// Relationship to a user.
	RequestedByUser ChangeRequestUserRelationship `json:"requested_by_user"`
	// Relationship to a user.
	RequestedUser ChangeRequestUserRelationship `json:"requested_user"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeRequestDecisionRelationships instantiates a new ChangeRequestDecisionRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeRequestDecisionRelationships(modifiedBy ChangeRequestUserRelationship, requestedByUser ChangeRequestUserRelationship, requestedUser ChangeRequestUserRelationship) *ChangeRequestDecisionRelationships {
	this := ChangeRequestDecisionRelationships{}
	this.ModifiedBy = modifiedBy
	this.RequestedByUser = requestedByUser
	this.RequestedUser = requestedUser
	return &this
}

// NewChangeRequestDecisionRelationshipsWithDefaults instantiates a new ChangeRequestDecisionRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeRequestDecisionRelationshipsWithDefaults() *ChangeRequestDecisionRelationships {
	this := ChangeRequestDecisionRelationships{}
	return &this
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *ChangeRequestDecisionRelationships) GetModifiedBy() ChangeRequestUserRelationship {
	if o == nil {
		var ret ChangeRequestUserRelationship
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestDecisionRelationships) GetModifiedByOk() (*ChangeRequestUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *ChangeRequestDecisionRelationships) SetModifiedBy(v ChangeRequestUserRelationship) {
	o.ModifiedBy = v
}

// GetRequestedByUser returns the RequestedByUser field value.
func (o *ChangeRequestDecisionRelationships) GetRequestedByUser() ChangeRequestUserRelationship {
	if o == nil {
		var ret ChangeRequestUserRelationship
		return ret
	}
	return o.RequestedByUser
}

// GetRequestedByUserOk returns a tuple with the RequestedByUser field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestDecisionRelationships) GetRequestedByUserOk() (*ChangeRequestUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedByUser, true
}

// SetRequestedByUser sets field value.
func (o *ChangeRequestDecisionRelationships) SetRequestedByUser(v ChangeRequestUserRelationship) {
	o.RequestedByUser = v
}

// GetRequestedUser returns the RequestedUser field value.
func (o *ChangeRequestDecisionRelationships) GetRequestedUser() ChangeRequestUserRelationship {
	if o == nil {
		var ret ChangeRequestUserRelationship
		return ret
	}
	return o.RequestedUser
}

// GetRequestedUserOk returns a tuple with the RequestedUser field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestDecisionRelationships) GetRequestedUserOk() (*ChangeRequestUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedUser, true
}

// SetRequestedUser sets field value.
func (o *ChangeRequestDecisionRelationships) SetRequestedUser(v ChangeRequestUserRelationship) {
	o.RequestedUser = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeRequestDecisionRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["modified_by"] = o.ModifiedBy
	toSerialize["requested_by_user"] = o.RequestedByUser
	toSerialize["requested_user"] = o.RequestedUser

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeRequestDecisionRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ModifiedBy      *ChangeRequestUserRelationship `json:"modified_by"`
		RequestedByUser *ChangeRequestUserRelationship `json:"requested_by_user"`
		RequestedUser   *ChangeRequestUserRelationship `json:"requested_user"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ModifiedBy == nil {
		return fmt.Errorf("required field modified_by missing")
	}
	if all.RequestedByUser == nil {
		return fmt.Errorf("required field requested_by_user missing")
	}
	if all.RequestedUser == nil {
		return fmt.Errorf("required field requested_user missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"modified_by", "requested_by_user", "requested_user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ModifiedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ModifiedBy = *all.ModifiedBy
	if all.RequestedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RequestedByUser = *all.RequestedByUser
	if all.RequestedUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RequestedUser = *all.RequestedUser

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
