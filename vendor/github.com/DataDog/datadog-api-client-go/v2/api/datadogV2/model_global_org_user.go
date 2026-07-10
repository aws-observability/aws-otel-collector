// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GlobalOrgUser User information for a global organization association.
type GlobalOrgUser struct {
	// The handle of the user.
	Handle string `json:"handle"`
	// The UUID of the user.
	Uuid uuid.UUID `json:"uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGlobalOrgUser instantiates a new GlobalOrgUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGlobalOrgUser(handle string, uuid uuid.UUID) *GlobalOrgUser {
	this := GlobalOrgUser{}
	this.Handle = handle
	this.Uuid = uuid
	return &this
}

// NewGlobalOrgUserWithDefaults instantiates a new GlobalOrgUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGlobalOrgUserWithDefaults() *GlobalOrgUser {
	this := GlobalOrgUser{}
	return &this
}

// GetHandle returns the Handle field value.
func (o *GlobalOrgUser) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *GlobalOrgUser) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *GlobalOrgUser) SetHandle(v string) {
	o.Handle = v
}

// GetUuid returns the Uuid field value.
func (o *GlobalOrgUser) GetUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *GlobalOrgUser) GetUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value.
func (o *GlobalOrgUser) SetUuid(v uuid.UUID) {
	o.Uuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GlobalOrgUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["handle"] = o.Handle
	toSerialize["uuid"] = o.Uuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GlobalOrgUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Handle *string    `json:"handle"`
		Uuid   *uuid.UUID `json:"uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.Uuid == nil {
		return fmt.Errorf("required field uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"handle", "uuid"})
	} else {
		return err
	}
	o.Handle = *all.Handle
	o.Uuid = *all.Uuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
