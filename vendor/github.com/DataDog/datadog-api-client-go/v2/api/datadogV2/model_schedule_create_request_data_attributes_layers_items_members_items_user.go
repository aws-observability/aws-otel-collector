// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser Identifies the user participating in this layer as a single object with an `id`.
type ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser struct {
	// The user's ID.
	Id *string `json:"id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser instantiates a new ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser() *ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser {
	this := ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser{}
	return &this
}

// NewScheduleCreateRequestDataAttributesLayersItemsMembersItemsUserWithDefaults instantiates a new ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleCreateRequestDataAttributesLayersItemsMembersItemsUserWithDefaults() *ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser {
	this := ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser) SetId(v string) {
	o.Id = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleCreateRequestDataAttributesLayersItemsMembersItemsUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id *string `json:"id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id"})
	} else {
		return err
	}
	o.Id = all.Id

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
