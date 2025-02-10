// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSNamespaceTagFilter AWS Metrics Collection tag filters list. Defaults to `[]`.
// The array of custom AWS resource tags (in the form `key:value`) defines a filter that Datadog uses when collecting metrics from a specified service.
// Wildcards, such as `?` (match a single character) and `*` (match multiple characters), and exclusion using `!` before the tag are supported.
// For EC2, only hosts that match one of the defined tags will be imported into Datadog. The rest will be ignored.
// For example, `env:production,instance-type:c?.*,!region:us-east-1`.
type AWSNamespaceTagFilter struct {
	// The AWS service for which the tag filters defined in `tags` will be applied.
	Namespace *string `json:"namespace,omitempty"`
	// The AWS resource tags to filter on for the service specified by `namespace`.
	Tags datadog.NullableList[string] `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSNamespaceTagFilter instantiates a new AWSNamespaceTagFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSNamespaceTagFilter() *AWSNamespaceTagFilter {
	this := AWSNamespaceTagFilter{}
	return &this
}

// NewAWSNamespaceTagFilterWithDefaults instantiates a new AWSNamespaceTagFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSNamespaceTagFilterWithDefaults() *AWSNamespaceTagFilter {
	this := AWSNamespaceTagFilter{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *AWSNamespaceTagFilter) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSNamespaceTagFilter) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *AWSNamespaceTagFilter) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *AWSNamespaceTagFilter) SetNamespace(v string) {
	o.Namespace = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSNamespaceTagFilter) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AWSNamespaceTagFilter) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *AWSNamespaceTagFilter) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.NullableList[string] and assigns it to the Tags field.
func (o *AWSNamespaceTagFilter) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *AWSNamespaceTagFilter) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *AWSNamespaceTagFilter) UnsetTags() {
	o.Tags.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSNamespaceTagFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSNamespaceTagFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Namespace *string                      `json:"namespace,omitempty"`
		Tags      datadog.NullableList[string] `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"namespace", "tags"})
	} else {
		return err
	}
	o.Namespace = all.Namespace
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
