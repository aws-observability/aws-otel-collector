// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerImageGroupAttributes Attributes for a Container Image Group.
type ContainerImageGroupAttributes struct {
	// Number of Container Images in the group.
	Count *int64 `json:"count,omitempty"`
	// Name of the Container Image group.
	Name *string `json:"name,omitempty"`
	// Tags from the group name parsed in key/value format.
	Tags interface{} `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewContainerImageGroupAttributes instantiates a new ContainerImageGroupAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerImageGroupAttributes() *ContainerImageGroupAttributes {
	this := ContainerImageGroupAttributes{}
	return &this
}

// NewContainerImageGroupAttributesWithDefaults instantiates a new ContainerImageGroupAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerImageGroupAttributesWithDefaults() *ContainerImageGroupAttributes {
	this := ContainerImageGroupAttributes{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ContainerImageGroupAttributes) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageGroupAttributes) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ContainerImageGroupAttributes) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ContainerImageGroupAttributes) SetCount(v int64) {
	o.Count = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerImageGroupAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageGroupAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerImageGroupAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerImageGroupAttributes) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ContainerImageGroupAttributes) GetTags() interface{} {
	if o == nil || o.Tags == nil {
		var ret interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageGroupAttributes) GetTagsOk() (*interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ContainerImageGroupAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given interface{} and assigns it to the Tags field.
func (o *ContainerImageGroupAttributes) SetTags(v interface{}) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerImageGroupAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
func (o *ContainerImageGroupAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count *int64      `json:"count,omitempty"`
		Name  *string     `json:"name,omitempty"`
		Tags  interface{} `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "name", "tags"})
	} else {
		return err
	}
	o.Count = all.Count
	o.Name = all.Name
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
