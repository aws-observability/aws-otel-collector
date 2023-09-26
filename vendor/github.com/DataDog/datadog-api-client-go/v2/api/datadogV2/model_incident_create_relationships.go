// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCreateRelationships The relationships the incident will have with other resources once created.
type IncidentCreateRelationships struct {
	// Relationship to user.
	CommanderUser NullableRelationshipToUser `json:"commander_user"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewIncidentCreateRelationships instantiates a new IncidentCreateRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentCreateRelationships(commanderUser NullableRelationshipToUser) *IncidentCreateRelationships {
	this := IncidentCreateRelationships{}
	this.CommanderUser = commanderUser
	return &this
}

// NewIncidentCreateRelationshipsWithDefaults instantiates a new IncidentCreateRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentCreateRelationshipsWithDefaults() *IncidentCreateRelationships {
	this := IncidentCreateRelationships{}
	return &this
}

// GetCommanderUser returns the CommanderUser field value.
func (o *IncidentCreateRelationships) GetCommanderUser() NullableRelationshipToUser {
	if o == nil {
		var ret NullableRelationshipToUser
		return ret
	}
	return o.CommanderUser
}

// GetCommanderUserOk returns a tuple with the CommanderUser field value
// and a boolean to check if the value has been set.
func (o *IncidentCreateRelationships) GetCommanderUserOk() (*NullableRelationshipToUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommanderUser, true
}

// SetCommanderUser sets field value.
func (o *IncidentCreateRelationships) SetCommanderUser(v NullableRelationshipToUser) {
	o.CommanderUser = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentCreateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["commander_user"] = o.CommanderUser

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentCreateRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommanderUser *NullableRelationshipToUser `json:"commander_user"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CommanderUser == nil {
		return fmt.Errorf("required field commander_user missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commander_user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CommanderUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CommanderUser = *all.CommanderUser

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
