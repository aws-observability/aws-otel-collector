// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImportRelationships The relationships for an incident import request.
type IncidentImportRelationships struct {
	// Relationship to user.
	CommanderUser NullableNullableRelationshipToUser `json:"commander_user,omitempty"`
	// Relationship to user.
	DeclaredByUser NullableNullableRelationshipToUser `json:"declared_by_user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentImportRelationships instantiates a new IncidentImportRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentImportRelationships() *IncidentImportRelationships {
	this := IncidentImportRelationships{}
	return &this
}

// NewIncidentImportRelationshipsWithDefaults instantiates a new IncidentImportRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentImportRelationshipsWithDefaults() *IncidentImportRelationships {
	this := IncidentImportRelationships{}
	return &this
}

// GetCommanderUser returns the CommanderUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportRelationships) GetCommanderUser() NullableRelationshipToUser {
	if o == nil || o.CommanderUser.Get() == nil {
		var ret NullableRelationshipToUser
		return ret
	}
	return *o.CommanderUser.Get()
}

// GetCommanderUserOk returns a tuple with the CommanderUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportRelationships) GetCommanderUserOk() (*NullableRelationshipToUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommanderUser.Get(), o.CommanderUser.IsSet()
}

// HasCommanderUser returns a boolean if a field has been set.
func (o *IncidentImportRelationships) HasCommanderUser() bool {
	return o != nil && o.CommanderUser.IsSet()
}

// SetCommanderUser gets a reference to the given NullableNullableRelationshipToUser and assigns it to the CommanderUser field.
func (o *IncidentImportRelationships) SetCommanderUser(v NullableRelationshipToUser) {
	o.CommanderUser.Set(&v)
}

// SetCommanderUserNil sets the value for CommanderUser to be an explicit nil.
func (o *IncidentImportRelationships) SetCommanderUserNil() {
	o.CommanderUser.Set(nil)
}

// UnsetCommanderUser ensures that no value is present for CommanderUser, not even an explicit nil.
func (o *IncidentImportRelationships) UnsetCommanderUser() {
	o.CommanderUser.Unset()
}

// GetDeclaredByUser returns the DeclaredByUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportRelationships) GetDeclaredByUser() NullableRelationshipToUser {
	if o == nil || o.DeclaredByUser.Get() == nil {
		var ret NullableRelationshipToUser
		return ret
	}
	return *o.DeclaredByUser.Get()
}

// GetDeclaredByUserOk returns a tuple with the DeclaredByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportRelationships) GetDeclaredByUserOk() (*NullableRelationshipToUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeclaredByUser.Get(), o.DeclaredByUser.IsSet()
}

// HasDeclaredByUser returns a boolean if a field has been set.
func (o *IncidentImportRelationships) HasDeclaredByUser() bool {
	return o != nil && o.DeclaredByUser.IsSet()
}

// SetDeclaredByUser gets a reference to the given NullableNullableRelationshipToUser and assigns it to the DeclaredByUser field.
func (o *IncidentImportRelationships) SetDeclaredByUser(v NullableRelationshipToUser) {
	o.DeclaredByUser.Set(&v)
}

// SetDeclaredByUserNil sets the value for DeclaredByUser to be an explicit nil.
func (o *IncidentImportRelationships) SetDeclaredByUserNil() {
	o.DeclaredByUser.Set(nil)
}

// UnsetDeclaredByUser ensures that no value is present for DeclaredByUser, not even an explicit nil.
func (o *IncidentImportRelationships) UnsetDeclaredByUser() {
	o.DeclaredByUser.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentImportRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CommanderUser.IsSet() {
		toSerialize["commander_user"] = o.CommanderUser.Get()
	}
	if o.DeclaredByUser.IsSet() {
		toSerialize["declared_by_user"] = o.DeclaredByUser.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentImportRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommanderUser  NullableNullableRelationshipToUser `json:"commander_user,omitempty"`
		DeclaredByUser NullableNullableRelationshipToUser `json:"declared_by_user,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commander_user", "declared_by_user"})
	} else {
		return err
	}
	o.CommanderUser = all.CommanderUser
	o.DeclaredByUser = all.DeclaredByUser

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
