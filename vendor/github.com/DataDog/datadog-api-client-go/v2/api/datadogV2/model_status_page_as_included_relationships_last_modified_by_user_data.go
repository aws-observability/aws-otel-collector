// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPageAsIncludedRelationshipsLastModifiedByUserData The data object identifying the Datadog user who last modified the status page.
type StatusPageAsIncludedRelationshipsLastModifiedByUserData struct {
	// The ID of the Datadog user who last modified the status page.
	Id string `json:"id"`
	// Users resource type.
	Type StatusPagesUserType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatusPageAsIncludedRelationshipsLastModifiedByUserData instantiates a new StatusPageAsIncludedRelationshipsLastModifiedByUserData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatusPageAsIncludedRelationshipsLastModifiedByUserData(id string, typeVar StatusPagesUserType) *StatusPageAsIncludedRelationshipsLastModifiedByUserData {
	this := StatusPageAsIncludedRelationshipsLastModifiedByUserData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewStatusPageAsIncludedRelationshipsLastModifiedByUserDataWithDefaults instantiates a new StatusPageAsIncludedRelationshipsLastModifiedByUserData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatusPageAsIncludedRelationshipsLastModifiedByUserDataWithDefaults() *StatusPageAsIncludedRelationshipsLastModifiedByUserData {
	this := StatusPageAsIncludedRelationshipsLastModifiedByUserData{}
	var typeVar StatusPagesUserType = STATUSPAGESUSERTYPE_USERS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *StatusPageAsIncludedRelationshipsLastModifiedByUserData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StatusPageAsIncludedRelationshipsLastModifiedByUserData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *StatusPageAsIncludedRelationshipsLastModifiedByUserData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *StatusPageAsIncludedRelationshipsLastModifiedByUserData) GetType() StatusPagesUserType {
	if o == nil {
		var ret StatusPagesUserType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StatusPageAsIncludedRelationshipsLastModifiedByUserData) GetTypeOk() (*StatusPagesUserType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StatusPageAsIncludedRelationshipsLastModifiedByUserData) SetType(v StatusPagesUserType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatusPageAsIncludedRelationshipsLastModifiedByUserData) MarshalJSON() ([]byte, error) {
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
func (o *StatusPageAsIncludedRelationshipsLastModifiedByUserData) UnmarshalJSON(bytes []byte) (err error) {
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
