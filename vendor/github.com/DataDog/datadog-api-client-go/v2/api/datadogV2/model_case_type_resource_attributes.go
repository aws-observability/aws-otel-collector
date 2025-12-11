// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseTypeResourceAttributes Case Type resource attributes
type CaseTypeResourceAttributes struct {
	// Timestamp of when the case type was deleted
	DeletedAt datadog.NullableTime `json:"deleted_at,omitempty"`
	// Case type description.
	Description *string `json:"description,omitempty"`
	// Case type emoji.
	Emoji *string `json:"emoji,omitempty"`
	// Case type name.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseTypeResourceAttributes instantiates a new CaseTypeResourceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseTypeResourceAttributes(name string) *CaseTypeResourceAttributes {
	this := CaseTypeResourceAttributes{}
	this.Name = name
	return &this
}

// NewCaseTypeResourceAttributesWithDefaults instantiates a new CaseTypeResourceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseTypeResourceAttributesWithDefaults() *CaseTypeResourceAttributes {
	this := CaseTypeResourceAttributes{}
	return &this
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaseTypeResourceAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CaseTypeResourceAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *CaseTypeResourceAttributes) HasDeletedAt() bool {
	return o != nil && o.DeletedAt.IsSet()
}

// SetDeletedAt gets a reference to the given datadog.NullableTime and assigns it to the DeletedAt field.
func (o *CaseTypeResourceAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil.
func (o *CaseTypeResourceAttributes) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil.
func (o *CaseTypeResourceAttributes) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CaseTypeResourceAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTypeResourceAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CaseTypeResourceAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CaseTypeResourceAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEmoji returns the Emoji field value if set, zero value otherwise.
func (o *CaseTypeResourceAttributes) GetEmoji() string {
	if o == nil || o.Emoji == nil {
		var ret string
		return ret
	}
	return *o.Emoji
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTypeResourceAttributes) GetEmojiOk() (*string, bool) {
	if o == nil || o.Emoji == nil {
		return nil, false
	}
	return o.Emoji, true
}

// HasEmoji returns a boolean if a field has been set.
func (o *CaseTypeResourceAttributes) HasEmoji() bool {
	return o != nil && o.Emoji != nil
}

// SetEmoji gets a reference to the given string and assigns it to the Emoji field.
func (o *CaseTypeResourceAttributes) SetEmoji(v string) {
	o.Emoji = &v
}

// GetName returns the Name field value.
func (o *CaseTypeResourceAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CaseTypeResourceAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CaseTypeResourceAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseTypeResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Emoji != nil {
		toSerialize["emoji"] = o.Emoji
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseTypeResourceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeletedAt   datadog.NullableTime `json:"deleted_at,omitempty"`
		Description *string              `json:"description,omitempty"`
		Emoji       *string              `json:"emoji,omitempty"`
		Name        *string              `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deleted_at", "description", "emoji", "name"})
	} else {
		return err
	}
	o.DeletedAt = all.DeletedAt
	o.Description = all.Description
	o.Emoji = all.Emoji
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
