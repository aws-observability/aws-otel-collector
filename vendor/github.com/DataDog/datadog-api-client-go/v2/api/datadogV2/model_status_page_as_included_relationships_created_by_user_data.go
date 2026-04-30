// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPageAsIncludedRelationshipsCreatedByUserData The data object identifying the Datadog user who created the status page.
type StatusPageAsIncludedRelationshipsCreatedByUserData struct {
	// The ID of the Datadog user who created the status page.
	Id string `json:"id"`
	// Users resource type.
	Type StatusPagesUserType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatusPageAsIncludedRelationshipsCreatedByUserData instantiates a new StatusPageAsIncludedRelationshipsCreatedByUserData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatusPageAsIncludedRelationshipsCreatedByUserData(id string, typeVar StatusPagesUserType) *StatusPageAsIncludedRelationshipsCreatedByUserData {
	this := StatusPageAsIncludedRelationshipsCreatedByUserData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewStatusPageAsIncludedRelationshipsCreatedByUserDataWithDefaults instantiates a new StatusPageAsIncludedRelationshipsCreatedByUserData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatusPageAsIncludedRelationshipsCreatedByUserDataWithDefaults() *StatusPageAsIncludedRelationshipsCreatedByUserData {
	this := StatusPageAsIncludedRelationshipsCreatedByUserData{}
	var typeVar StatusPagesUserType = STATUSPAGESUSERTYPE_USERS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *StatusPageAsIncludedRelationshipsCreatedByUserData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StatusPageAsIncludedRelationshipsCreatedByUserData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *StatusPageAsIncludedRelationshipsCreatedByUserData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *StatusPageAsIncludedRelationshipsCreatedByUserData) GetType() StatusPagesUserType {
	if o == nil {
		var ret StatusPagesUserType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StatusPageAsIncludedRelationshipsCreatedByUserData) GetTypeOk() (*StatusPagesUserType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StatusPageAsIncludedRelationshipsCreatedByUserData) SetType(v StatusPagesUserType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatusPageAsIncludedRelationshipsCreatedByUserData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StatusPageAsIncludedRelationshipsCreatedByUserData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string              `json:"id"`
		Type *StatusPagesUserType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
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
	o.Id = *all.Id
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
