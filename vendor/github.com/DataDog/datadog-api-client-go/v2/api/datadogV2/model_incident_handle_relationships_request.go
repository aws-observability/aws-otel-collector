// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentHandleRelationshipsRequest Relationships to associate with an incident handle in a create or update request.
type IncidentHandleRelationshipsRequest struct {
	// A single relationship object for an incident handle, wrapping the related resource data.
	CommanderUser *IncidentHandleRelationship `json:"commander_user,omitempty"`
	// A single relationship object for an incident handle, wrapping the related resource data.
	IncidentType IncidentHandleRelationship `json:"incident_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentHandleRelationshipsRequest instantiates a new IncidentHandleRelationshipsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentHandleRelationshipsRequest(incidentType IncidentHandleRelationship) *IncidentHandleRelationshipsRequest {
	this := IncidentHandleRelationshipsRequest{}
	this.IncidentType = incidentType
	return &this
}

// NewIncidentHandleRelationshipsRequestWithDefaults instantiates a new IncidentHandleRelationshipsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentHandleRelationshipsRequestWithDefaults() *IncidentHandleRelationshipsRequest {
	this := IncidentHandleRelationshipsRequest{}
	return &this
}

// GetCommanderUser returns the CommanderUser field value if set, zero value otherwise.
func (o *IncidentHandleRelationshipsRequest) GetCommanderUser() IncidentHandleRelationship {
	if o == nil || o.CommanderUser == nil {
		var ret IncidentHandleRelationship
		return ret
	}
	return *o.CommanderUser
}

// GetCommanderUserOk returns a tuple with the CommanderUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentHandleRelationshipsRequest) GetCommanderUserOk() (*IncidentHandleRelationship, bool) {
	if o == nil || o.CommanderUser == nil {
		return nil, false
	}
	return o.CommanderUser, true
}

// HasCommanderUser returns a boolean if a field has been set.
func (o *IncidentHandleRelationshipsRequest) HasCommanderUser() bool {
	return o != nil && o.CommanderUser != nil
}

// SetCommanderUser gets a reference to the given IncidentHandleRelationship and assigns it to the CommanderUser field.
func (o *IncidentHandleRelationshipsRequest) SetCommanderUser(v IncidentHandleRelationship) {
	o.CommanderUser = &v
}

// GetIncidentType returns the IncidentType field value.
func (o *IncidentHandleRelationshipsRequest) GetIncidentType() IncidentHandleRelationship {
	if o == nil {
		var ret IncidentHandleRelationship
		return ret
	}
	return o.IncidentType
}

// GetIncidentTypeOk returns a tuple with the IncidentType field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleRelationshipsRequest) GetIncidentTypeOk() (*IncidentHandleRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentType, true
}

// SetIncidentType sets field value.
func (o *IncidentHandleRelationshipsRequest) SetIncidentType(v IncidentHandleRelationship) {
	o.IncidentType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentHandleRelationshipsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CommanderUser != nil {
		toSerialize["commander_user"] = o.CommanderUser
	}
	toSerialize["incident_type"] = o.IncidentType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentHandleRelationshipsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommanderUser *IncidentHandleRelationship `json:"commander_user,omitempty"`
		IncidentType  *IncidentHandleRelationship `json:"incident_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IncidentType == nil {
		return fmt.Errorf("required field incident_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commander_user", "incident_type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CommanderUser != nil && all.CommanderUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CommanderUser = all.CommanderUser
	if all.IncidentType.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncidentType = *all.IncidentType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableIncidentHandleRelationshipsRequest handles when a null is used for IncidentHandleRelationshipsRequest.
type NullableIncidentHandleRelationshipsRequest struct {
	value *IncidentHandleRelationshipsRequest
	isSet bool
}

// Get returns the associated value.
func (v NullableIncidentHandleRelationshipsRequest) Get() *IncidentHandleRelationshipsRequest {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableIncidentHandleRelationshipsRequest) Set(val *IncidentHandleRelationshipsRequest) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableIncidentHandleRelationshipsRequest) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableIncidentHandleRelationshipsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncidentHandleRelationshipsRequest initializes the struct as if Set has been called.
func NewNullableIncidentHandleRelationshipsRequest(val *IncidentHandleRelationshipsRequest) *NullableIncidentHandleRelationshipsRequest {
	return &NullableIncidentHandleRelationshipsRequest{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableIncidentHandleRelationshipsRequest) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableIncidentHandleRelationshipsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
