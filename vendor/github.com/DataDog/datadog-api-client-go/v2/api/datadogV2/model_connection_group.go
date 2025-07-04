// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConnectionGroup The definition of `ConnectionGroup` object.
type ConnectionGroup struct {
	// The `ConnectionGroup` `connectionGroupId`.
	ConnectionGroupId string `json:"connectionGroupId"`
	// The `ConnectionGroup` `label`.
	Label string `json:"label"`
	// The `ConnectionGroup` `tags`.
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConnectionGroup instantiates a new ConnectionGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConnectionGroup(connectionGroupId string, label string, tags []string) *ConnectionGroup {
	this := ConnectionGroup{}
	this.ConnectionGroupId = connectionGroupId
	this.Label = label
	this.Tags = tags
	return &this
}

// NewConnectionGroupWithDefaults instantiates a new ConnectionGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConnectionGroupWithDefaults() *ConnectionGroup {
	this := ConnectionGroup{}
	return &this
}

// GetConnectionGroupId returns the ConnectionGroupId field value.
func (o *ConnectionGroup) GetConnectionGroupId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConnectionGroupId
}

// GetConnectionGroupIdOk returns a tuple with the ConnectionGroupId field value
// and a boolean to check if the value has been set.
func (o *ConnectionGroup) GetConnectionGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionGroupId, true
}

// SetConnectionGroupId sets field value.
func (o *ConnectionGroup) SetConnectionGroupId(v string) {
	o.ConnectionGroupId = v
}

// GetLabel returns the Label field value.
func (o *ConnectionGroup) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ConnectionGroup) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *ConnectionGroup) SetLabel(v string) {
	o.Label = v
}

// GetTags returns the Tags field value.
func (o *ConnectionGroup) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ConnectionGroup) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *ConnectionGroup) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConnectionGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["connectionGroupId"] = o.ConnectionGroupId
	toSerialize["label"] = o.Label
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConnectionGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConnectionGroupId *string   `json:"connectionGroupId"`
		Label             *string   `json:"label"`
		Tags              *[]string `json:"tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConnectionGroupId == nil {
		return fmt.Errorf("required field connectionGroupId missing")
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connectionGroupId", "label", "tags"})
	} else {
		return err
	}
	o.ConnectionGroupId = *all.ConnectionGroupId
	o.Label = *all.Label
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
