// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerGroupAttributes Attributes for a container group.
type ContainerGroupAttributes struct {
	// Number of containers in the group.
	Count *int64 `json:"count,omitempty"`
	// Tags from the group name parsed in key/value format.
	Tags interface{} `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewContainerGroupAttributes instantiates a new ContainerGroupAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerGroupAttributes() *ContainerGroupAttributes {
	this := ContainerGroupAttributes{}
	return &this
}

// NewContainerGroupAttributesWithDefaults instantiates a new ContainerGroupAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerGroupAttributesWithDefaults() *ContainerGroupAttributes {
	this := ContainerGroupAttributes{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ContainerGroupAttributes) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupAttributes) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ContainerGroupAttributes) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ContainerGroupAttributes) SetCount(v int64) {
	o.Count = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ContainerGroupAttributes) GetTags() interface{} {
	if o == nil || o.Tags == nil {
		var ret interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupAttributes) GetTagsOk() (*interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ContainerGroupAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given interface{} and assigns it to the Tags field.
func (o *ContainerGroupAttributes) SetTags(v interface{}) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerGroupAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
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
func (o *ContainerGroupAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count *int64      `json:"count,omitempty"`
		Tags  interface{} `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "tags"})
	} else {
		return err
	}
	o.Count = all.Count
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
