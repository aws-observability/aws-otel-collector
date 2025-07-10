// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleUpdateRequestDataAttributesLayersItemsMembersItems Defines a single member within a layer during an update request, referring
// to a specific user.
type ScheduleUpdateRequestDataAttributesLayersItemsMembersItems struct {
	// Identifies the user who is assigned to this member object. Only `id` is required.
	User *ScheduleUpdateRequestDataAttributesLayersItemsMembersItemsUser `json:"user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleUpdateRequestDataAttributesLayersItemsMembersItems instantiates a new ScheduleUpdateRequestDataAttributesLayersItemsMembersItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleUpdateRequestDataAttributesLayersItemsMembersItems() *ScheduleUpdateRequestDataAttributesLayersItemsMembersItems {
	this := ScheduleUpdateRequestDataAttributesLayersItemsMembersItems{}
	return &this
}

// NewScheduleUpdateRequestDataAttributesLayersItemsMembersItemsWithDefaults instantiates a new ScheduleUpdateRequestDataAttributesLayersItemsMembersItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleUpdateRequestDataAttributesLayersItemsMembersItemsWithDefaults() *ScheduleUpdateRequestDataAttributesLayersItemsMembersItems {
	this := ScheduleUpdateRequestDataAttributesLayersItemsMembersItems{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsMembersItems) GetUser() ScheduleUpdateRequestDataAttributesLayersItemsMembersItemsUser {
	if o == nil || o.User == nil {
		var ret ScheduleUpdateRequestDataAttributesLayersItemsMembersItemsUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsMembersItems) GetUserOk() (*ScheduleUpdateRequestDataAttributesLayersItemsMembersItemsUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsMembersItems) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given ScheduleUpdateRequestDataAttributesLayersItemsMembersItemsUser and assigns it to the User field.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsMembersItems) SetUser(v ScheduleUpdateRequestDataAttributesLayersItemsMembersItemsUser) {
	o.User = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleUpdateRequestDataAttributesLayersItemsMembersItems) MarshalJSON() ([]byte, error) {
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
func (o *ScheduleUpdateRequestDataAttributesLayersItemsMembersItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		User *ScheduleUpdateRequestDataAttributesLayersItemsMembersItemsUser `json:"user,omitempty"`
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
