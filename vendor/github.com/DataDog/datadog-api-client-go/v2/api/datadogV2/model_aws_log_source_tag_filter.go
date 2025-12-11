// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSLogSourceTagFilter AWS log source tag filter list. Defaults to `[]`.
// Array of log source to AWS resource tag mappings. Each mapping contains a log source and its
// associated AWS resource tags (in `key:value` format) used to filter logs submitted to Datadog.
// Tag filters are applied for tags on the AWS resource emitting logs; tags associated with the
// log storage entity (such as a CloudWatch Log Group or S3 Bucket) are not considered.
// For more information on resource tag filter syntax,
// [see AWS resource exclusion](https://docs.datadoghq.com/account_management/billing/aws/#aws-resource-exclusion)
// in the AWS integration billing page.
type AWSLogSourceTagFilter struct {
	// The AWS log source to which the tag filters defined in `tags` are applied.
	Source *string `json:"source,omitempty"`
	// The AWS resource tags to filter on for the log source specified by `source`.
	Tags datadog.NullableList[string] `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSLogSourceTagFilter instantiates a new AWSLogSourceTagFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSLogSourceTagFilter() *AWSLogSourceTagFilter {
	this := AWSLogSourceTagFilter{}
	return &this
}

// NewAWSLogSourceTagFilterWithDefaults instantiates a new AWSLogSourceTagFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSLogSourceTagFilterWithDefaults() *AWSLogSourceTagFilter {
	this := AWSLogSourceTagFilter{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AWSLogSourceTagFilter) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSLogSourceTagFilter) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AWSLogSourceTagFilter) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *AWSLogSourceTagFilter) SetSource(v string) {
	o.Source = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSLogSourceTagFilter) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AWSLogSourceTagFilter) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *AWSLogSourceTagFilter) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.NullableList[string] and assigns it to the Tags field.
func (o *AWSLogSourceTagFilter) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *AWSLogSourceTagFilter) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *AWSLogSourceTagFilter) UnsetTags() {
	o.Tags.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSLogSourceTagFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
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
func (o *AWSLogSourceTagFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Source *string                      `json:"source,omitempty"`
		Tags   datadog.NullableList[string] `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"source", "tags"})
	} else {
		return err
	}
	o.Source = all.Source
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
