// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateAppTagsRequestDataAttributes Attributes for replacing an app's tags.
type UpdateAppTagsRequestDataAttributes struct {
	// The full list of tags that should be set on the app. Existing tags not present in this list are removed.
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateAppTagsRequestDataAttributes instantiates a new UpdateAppTagsRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateAppTagsRequestDataAttributes(tags []string) *UpdateAppTagsRequestDataAttributes {
	this := UpdateAppTagsRequestDataAttributes{}
	this.Tags = tags
	return &this
}

// NewUpdateAppTagsRequestDataAttributesWithDefaults instantiates a new UpdateAppTagsRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateAppTagsRequestDataAttributesWithDefaults() *UpdateAppTagsRequestDataAttributes {
	this := UpdateAppTagsRequestDataAttributes{}
	return &this
}

// GetTags returns the Tags field value.
func (o *UpdateAppTagsRequestDataAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *UpdateAppTagsRequestDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *UpdateAppTagsRequestDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateAppTagsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateAppTagsRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Tags *[]string `json:"tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"tags"})
	} else {
		return err
	}
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
