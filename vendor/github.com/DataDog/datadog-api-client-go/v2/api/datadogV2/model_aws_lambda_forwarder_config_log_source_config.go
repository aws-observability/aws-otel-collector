// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSLambdaForwarderConfigLogSourceConfig Log source configuration.
type AWSLambdaForwarderConfigLogSourceConfig struct {
	// List of AWS log source tag filters. Defaults to `[]`.
	TagFilters []AWSLogSourceTagFilter `json:"tag_filters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSLambdaForwarderConfigLogSourceConfig instantiates a new AWSLambdaForwarderConfigLogSourceConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSLambdaForwarderConfigLogSourceConfig() *AWSLambdaForwarderConfigLogSourceConfig {
	this := AWSLambdaForwarderConfigLogSourceConfig{}
	return &this
}

// NewAWSLambdaForwarderConfigLogSourceConfigWithDefaults instantiates a new AWSLambdaForwarderConfigLogSourceConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSLambdaForwarderConfigLogSourceConfigWithDefaults() *AWSLambdaForwarderConfigLogSourceConfig {
	this := AWSLambdaForwarderConfigLogSourceConfig{}
	return &this
}

// GetTagFilters returns the TagFilters field value if set, zero value otherwise.
func (o *AWSLambdaForwarderConfigLogSourceConfig) GetTagFilters() []AWSLogSourceTagFilter {
	if o == nil || o.TagFilters == nil {
		var ret []AWSLogSourceTagFilter
		return ret
	}
	return o.TagFilters
}

// GetTagFiltersOk returns a tuple with the TagFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSLambdaForwarderConfigLogSourceConfig) GetTagFiltersOk() (*[]AWSLogSourceTagFilter, bool) {
	if o == nil || o.TagFilters == nil {
		return nil, false
	}
	return &o.TagFilters, true
}

// HasTagFilters returns a boolean if a field has been set.
func (o *AWSLambdaForwarderConfigLogSourceConfig) HasTagFilters() bool {
	return o != nil && o.TagFilters != nil
}

// SetTagFilters gets a reference to the given []AWSLogSourceTagFilter and assigns it to the TagFilters field.
func (o *AWSLambdaForwarderConfigLogSourceConfig) SetTagFilters(v []AWSLogSourceTagFilter) {
	o.TagFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSLambdaForwarderConfigLogSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TagFilters != nil {
		toSerialize["tag_filters"] = o.TagFilters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSLambdaForwarderConfigLogSourceConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TagFilters []AWSLogSourceTagFilter `json:"tag_filters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"tag_filters"})
	} else {
		return err
	}
	o.TagFilters = all.TagFilters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
