// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestUserRelationship Relationship to a user.
type ChangeRequestUserRelationship struct {
	// User relationship data.
	Data NullableChangeRequestUserRelationshipData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeRequestUserRelationship instantiates a new ChangeRequestUserRelationship object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeRequestUserRelationship(data NullableChangeRequestUserRelationshipData) *ChangeRequestUserRelationship {
	this := ChangeRequestUserRelationship{}
	this.Data = data
	return &this
}

// NewChangeRequestUserRelationshipWithDefaults instantiates a new ChangeRequestUserRelationship object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeRequestUserRelationshipWithDefaults() *ChangeRequestUserRelationship {
	this := ChangeRequestUserRelationship{}
	return &this
}

// GetData returns the Data field value.
// If the value is explicit nil, the zero value for ChangeRequestUserRelationshipData will be returned.
func (o *ChangeRequestUserRelationship) GetData() ChangeRequestUserRelationshipData {
	if o == nil || o.Data.Get() == nil {
		var ret ChangeRequestUserRelationshipData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ChangeRequestUserRelationship) GetDataOk() (*ChangeRequestUserRelationshipData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// SetData sets field value.
func (o *ChangeRequestUserRelationship) SetData(v ChangeRequestUserRelationshipData) {
	o.Data.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeRequestUserRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeRequestUserRelationship) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data NullableChangeRequestUserRelationshipData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Data.IsSet() {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
