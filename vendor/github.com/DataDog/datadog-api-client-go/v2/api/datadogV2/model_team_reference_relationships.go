// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamReferenceRelationships Collects the key relationship fields for a team reference, specifically on-call users.
type TeamReferenceRelationships struct {
	// Defines which users are on-call within a team, stored as an array of references.
	OncallUsers *TeamReferenceRelationshipsOncallUsers `json:"oncall_users,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamReferenceRelationships instantiates a new TeamReferenceRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamReferenceRelationships() *TeamReferenceRelationships {
	this := TeamReferenceRelationships{}
	return &this
}

// NewTeamReferenceRelationshipsWithDefaults instantiates a new TeamReferenceRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamReferenceRelationshipsWithDefaults() *TeamReferenceRelationships {
	this := TeamReferenceRelationships{}
	return &this
}

// GetOncallUsers returns the OncallUsers field value if set, zero value otherwise.
func (o *TeamReferenceRelationships) GetOncallUsers() TeamReferenceRelationshipsOncallUsers {
	if o == nil || o.OncallUsers == nil {
		var ret TeamReferenceRelationshipsOncallUsers
		return ret
	}
	return *o.OncallUsers
}

// GetOncallUsersOk returns a tuple with the OncallUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamReferenceRelationships) GetOncallUsersOk() (*TeamReferenceRelationshipsOncallUsers, bool) {
	if o == nil || o.OncallUsers == nil {
		return nil, false
	}
	return o.OncallUsers, true
}

// HasOncallUsers returns a boolean if a field has been set.
func (o *TeamReferenceRelationships) HasOncallUsers() bool {
	return o != nil && o.OncallUsers != nil
}

// SetOncallUsers gets a reference to the given TeamReferenceRelationshipsOncallUsers and assigns it to the OncallUsers field.
func (o *TeamReferenceRelationships) SetOncallUsers(v TeamReferenceRelationshipsOncallUsers) {
	o.OncallUsers = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamReferenceRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OncallUsers != nil {
		toSerialize["oncall_users"] = o.OncallUsers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamReferenceRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OncallUsers *TeamReferenceRelationshipsOncallUsers `json:"oncall_users,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"oncall_users"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OncallUsers != nil && all.OncallUsers.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OncallUsers = all.OncallUsers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
