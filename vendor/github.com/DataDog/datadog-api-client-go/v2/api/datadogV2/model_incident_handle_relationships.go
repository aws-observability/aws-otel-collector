// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentHandleRelationships Relationships associated with an incident handle response, including linked users and incident type.
type IncidentHandleRelationships struct {
	// A single relationship object for an incident handle, wrapping the related resource data.
	CommanderUser *IncidentHandleRelationship `json:"commander_user,omitempty"`
	// A single relationship object for an incident handle, wrapping the related resource data.
	CreatedByUser IncidentHandleRelationship `json:"created_by_user"`
	// A single relationship object for an incident handle, wrapping the related resource data.
	IncidentType IncidentHandleRelationship `json:"incident_type"`
	// A single relationship object for an incident handle, wrapping the related resource data.
	LastModifiedByUser IncidentHandleRelationship `json:"last_modified_by_user"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentHandleRelationships instantiates a new IncidentHandleRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentHandleRelationships(createdByUser IncidentHandleRelationship, incidentType IncidentHandleRelationship, lastModifiedByUser IncidentHandleRelationship) *IncidentHandleRelationships {
	this := IncidentHandleRelationships{}
	this.CreatedByUser = createdByUser
	this.IncidentType = incidentType
	this.LastModifiedByUser = lastModifiedByUser
	return &this
}

// NewIncidentHandleRelationshipsWithDefaults instantiates a new IncidentHandleRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentHandleRelationshipsWithDefaults() *IncidentHandleRelationships {
	this := IncidentHandleRelationships{}
	return &this
}

// GetCommanderUser returns the CommanderUser field value if set, zero value otherwise.
func (o *IncidentHandleRelationships) GetCommanderUser() IncidentHandleRelationship {
	if o == nil || o.CommanderUser == nil {
		var ret IncidentHandleRelationship
		return ret
	}
	return *o.CommanderUser
}

// GetCommanderUserOk returns a tuple with the CommanderUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentHandleRelationships) GetCommanderUserOk() (*IncidentHandleRelationship, bool) {
	if o == nil || o.CommanderUser == nil {
		return nil, false
	}
	return o.CommanderUser, true
}

// HasCommanderUser returns a boolean if a field has been set.
func (o *IncidentHandleRelationships) HasCommanderUser() bool {
	return o != nil && o.CommanderUser != nil
}

// SetCommanderUser gets a reference to the given IncidentHandleRelationship and assigns it to the CommanderUser field.
func (o *IncidentHandleRelationships) SetCommanderUser(v IncidentHandleRelationship) {
	o.CommanderUser = &v
}

// GetCreatedByUser returns the CreatedByUser field value.
func (o *IncidentHandleRelationships) GetCreatedByUser() IncidentHandleRelationship {
	if o == nil {
		var ret IncidentHandleRelationship
		return ret
	}
	return o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleRelationships) GetCreatedByUserOk() (*IncidentHandleRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByUser, true
}

// SetCreatedByUser sets field value.
func (o *IncidentHandleRelationships) SetCreatedByUser(v IncidentHandleRelationship) {
	o.CreatedByUser = v
}

// GetIncidentType returns the IncidentType field value.
func (o *IncidentHandleRelationships) GetIncidentType() IncidentHandleRelationship {
	if o == nil {
		var ret IncidentHandleRelationship
		return ret
	}
	return o.IncidentType
}

// GetIncidentTypeOk returns a tuple with the IncidentType field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleRelationships) GetIncidentTypeOk() (*IncidentHandleRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentType, true
}

// SetIncidentType sets field value.
func (o *IncidentHandleRelationships) SetIncidentType(v IncidentHandleRelationship) {
	o.IncidentType = v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value.
func (o *IncidentHandleRelationships) GetLastModifiedByUser() IncidentHandleRelationship {
	if o == nil {
		var ret IncidentHandleRelationship
		return ret
	}
	return o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleRelationships) GetLastModifiedByUserOk() (*IncidentHandleRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedByUser, true
}

// SetLastModifiedByUser sets field value.
func (o *IncidentHandleRelationships) SetLastModifiedByUser(v IncidentHandleRelationship) {
	o.LastModifiedByUser = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentHandleRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CommanderUser != nil {
		toSerialize["commander_user"] = o.CommanderUser
	}
	toSerialize["created_by_user"] = o.CreatedByUser
	toSerialize["incident_type"] = o.IncidentType
	toSerialize["last_modified_by_user"] = o.LastModifiedByUser

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentHandleRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommanderUser      *IncidentHandleRelationship `json:"commander_user,omitempty"`
		CreatedByUser      *IncidentHandleRelationship `json:"created_by_user"`
		IncidentType       *IncidentHandleRelationship `json:"incident_type"`
		LastModifiedByUser *IncidentHandleRelationship `json:"last_modified_by_user"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedByUser == nil {
		return fmt.Errorf("required field created_by_user missing")
	}
	if all.IncidentType == nil {
		return fmt.Errorf("required field incident_type missing")
	}
	if all.LastModifiedByUser == nil {
		return fmt.Errorf("required field last_modified_by_user missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commander_user", "created_by_user", "incident_type", "last_modified_by_user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CommanderUser != nil && all.CommanderUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CommanderUser = all.CommanderUser
	if all.CreatedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedByUser = *all.CreatedByUser
	if all.IncidentType.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncidentType = *all.IncidentType
	if all.LastModifiedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastModifiedByUser = *all.LastModifiedByUser

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableIncidentHandleRelationships handles when a null is used for IncidentHandleRelationships.
type NullableIncidentHandleRelationships struct {
	value *IncidentHandleRelationships
	isSet bool
}

// Get returns the associated value.
func (v NullableIncidentHandleRelationships) Get() *IncidentHandleRelationships {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableIncidentHandleRelationships) Set(val *IncidentHandleRelationships) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableIncidentHandleRelationships) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableIncidentHandleRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncidentHandleRelationships initializes the struct as if Set has been called.
func NewNullableIncidentHandleRelationships(val *IncidentHandleRelationships) *NullableIncidentHandleRelationships {
	return &NullableIncidentHandleRelationships{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableIncidentHandleRelationships) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableIncidentHandleRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
