// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppKeyRegistrationData Data related to the app key registration.
type AppKeyRegistrationData struct {
	// The app key registration identifier
	Id *uuid.UUID `json:"id,omitempty"`
	// The definition of `AppKeyRegistrationDataType` object.
	Type AppKeyRegistrationDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAppKeyRegistrationData instantiates a new AppKeyRegistrationData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAppKeyRegistrationData(typeVar AppKeyRegistrationDataType) *AppKeyRegistrationData {
	this := AppKeyRegistrationData{}
	this.Type = typeVar
	return &this
}

// NewAppKeyRegistrationDataWithDefaults instantiates a new AppKeyRegistrationData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAppKeyRegistrationDataWithDefaults() *AppKeyRegistrationData {
	this := AppKeyRegistrationData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppKeyRegistrationData) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppKeyRegistrationData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppKeyRegistrationData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *AppKeyRegistrationData) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *AppKeyRegistrationData) GetType() AppKeyRegistrationDataType {
	if o == nil {
		var ret AppKeyRegistrationDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppKeyRegistrationData) GetTypeOk() (*AppKeyRegistrationDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AppKeyRegistrationData) SetType(v AppKeyRegistrationDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AppKeyRegistrationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AppKeyRegistrationData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *uuid.UUID                  `json:"id,omitempty"`
		Type *AppKeyRegistrationDataType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
