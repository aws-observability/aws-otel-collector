// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// APIKeyRelationships Resources related to the API key.
type APIKeyRelationships struct {
	// Relationship to user.
	CreatedBy *RelationshipToUser `json:"created_by,omitempty"`
	// Relationship to user.
	ModifiedBy NullableNullableRelationshipToUser `json:"modified_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAPIKeyRelationships instantiates a new APIKeyRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAPIKeyRelationships() *APIKeyRelationships {
	this := APIKeyRelationships{}
	return &this
}

// NewAPIKeyRelationshipsWithDefaults instantiates a new APIKeyRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAPIKeyRelationshipsWithDefaults() *APIKeyRelationships {
	this := APIKeyRelationships{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *APIKeyRelationships) GetCreatedBy() RelationshipToUser {
	if o == nil || o.CreatedBy == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIKeyRelationships) GetCreatedByOk() (*RelationshipToUser, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *APIKeyRelationships) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given RelationshipToUser and assigns it to the CreatedBy field.
func (o *APIKeyRelationships) SetCreatedBy(v RelationshipToUser) {
	o.CreatedBy = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *APIKeyRelationships) GetModifiedBy() NullableRelationshipToUser {
	if o == nil || o.ModifiedBy.Get() == nil {
		var ret NullableRelationshipToUser
		return ret
	}
	return *o.ModifiedBy.Get()
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *APIKeyRelationships) GetModifiedByOk() (*NullableRelationshipToUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedBy.Get(), o.ModifiedBy.IsSet()
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *APIKeyRelationships) HasModifiedBy() bool {
	return o != nil && o.ModifiedBy.IsSet()
}

// SetModifiedBy gets a reference to the given NullableNullableRelationshipToUser and assigns it to the ModifiedBy field.
func (o *APIKeyRelationships) SetModifiedBy(v NullableRelationshipToUser) {
	o.ModifiedBy.Set(&v)
}

// SetModifiedByNil sets the value for ModifiedBy to be an explicit nil.
func (o *APIKeyRelationships) SetModifiedByNil() {
	o.ModifiedBy.Set(nil)
}

// UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil.
func (o *APIKeyRelationships) UnsetModifiedBy() {
	o.ModifiedBy.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o APIKeyRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.ModifiedBy.IsSet() {
		toSerialize["modified_by"] = o.ModifiedBy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *APIKeyRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedBy  *RelationshipToUser                `json:"created_by,omitempty"`
		ModifiedBy NullableNullableRelationshipToUser `json:"modified_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by", "modified_by"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedBy != nil && all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = all.CreatedBy
	o.ModifiedBy = all.ModifiedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
