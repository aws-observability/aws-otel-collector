// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionQuerySpecConnectionGroup The connection group to use for an action query.
type ActionQuerySpecConnectionGroup struct {
	// The ID of the connection group.
	Id *uuid.UUID `json:"id,omitempty"`
	// The tags of the connection group.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActionQuerySpecConnectionGroup instantiates a new ActionQuerySpecConnectionGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionQuerySpecConnectionGroup() *ActionQuerySpecConnectionGroup {
	this := ActionQuerySpecConnectionGroup{}
	return &this
}

// NewActionQuerySpecConnectionGroupWithDefaults instantiates a new ActionQuerySpecConnectionGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionQuerySpecConnectionGroupWithDefaults() *ActionQuerySpecConnectionGroup {
	this := ActionQuerySpecConnectionGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ActionQuerySpecConnectionGroup) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionQuerySpecConnectionGroup) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ActionQuerySpecConnectionGroup) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *ActionQuerySpecConnectionGroup) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ActionQuerySpecConnectionGroup) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionQuerySpecConnectionGroup) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ActionQuerySpecConnectionGroup) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ActionQuerySpecConnectionGroup) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionQuerySpecConnectionGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ActionQuerySpecConnectionGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *uuid.UUID `json:"id,omitempty"`
		Tags []string   `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "tags"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
