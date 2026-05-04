// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeletedSuiteResponseDataAttributes Attributes of a deleted Synthetic test suite, including deletion timestamp and public ID.
type DeletedSuiteResponseDataAttributes struct {
	// Deletion timestamp of the Synthetic suite ID.
	DeletedAt *string `json:"deleted_at,omitempty"`
	// The Synthetic suite ID deleted.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeletedSuiteResponseDataAttributes instantiates a new DeletedSuiteResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeletedSuiteResponseDataAttributes() *DeletedSuiteResponseDataAttributes {
	this := DeletedSuiteResponseDataAttributes{}
	return &this
}

// NewDeletedSuiteResponseDataAttributesWithDefaults instantiates a new DeletedSuiteResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeletedSuiteResponseDataAttributesWithDefaults() *DeletedSuiteResponseDataAttributes {
	this := DeletedSuiteResponseDataAttributes{}
	return &this
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *DeletedSuiteResponseDataAttributes) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedSuiteResponseDataAttributes) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *DeletedSuiteResponseDataAttributes) HasDeletedAt() bool {
	return o != nil && o.DeletedAt != nil
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *DeletedSuiteResponseDataAttributes) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *DeletedSuiteResponseDataAttributes) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedSuiteResponseDataAttributes) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *DeletedSuiteResponseDataAttributes) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *DeletedSuiteResponseDataAttributes) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeletedSuiteResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeletedSuiteResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeletedAt *string `json:"deleted_at,omitempty"`
		PublicId  *string `json:"public_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deleted_at", "public_id"})
	} else {
		return err
	}
	o.DeletedAt = all.DeletedAt
	o.PublicId = all.PublicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
