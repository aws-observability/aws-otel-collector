// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSNamespaceFiltersExcludeOnly Exclude only these namespaces from metrics collection.
// Defaults to `["AWS/SQS", "AWS/ElasticMapReduce", "AWS/Usage"]`.
// `AWS/SQS`, `AWS/ElasticMapReduce`, and `AWS/Usage` are excluded by default
// to reduce your AWS CloudWatch costs from `GetMetricData` API calls.
type AWSNamespaceFiltersExcludeOnly struct {
	// Exclude only these namespaces from metrics collection.
	// Defaults to `["AWS/SQS", "AWS/ElasticMapReduce", "AWS/Usage"]`.
	// `AWS/SQS`, `AWS/ElasticMapReduce`, and `AWS/Usage` are excluded by default
	// to reduce your AWS CloudWatch costs from `GetMetricData` API calls.
	ExcludeOnly []string `json:"exclude_only"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSNamespaceFiltersExcludeOnly instantiates a new AWSNamespaceFiltersExcludeOnly object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSNamespaceFiltersExcludeOnly(excludeOnly []string) *AWSNamespaceFiltersExcludeOnly {
	this := AWSNamespaceFiltersExcludeOnly{}
	this.ExcludeOnly = excludeOnly
	return &this
}

// NewAWSNamespaceFiltersExcludeOnlyWithDefaults instantiates a new AWSNamespaceFiltersExcludeOnly object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSNamespaceFiltersExcludeOnlyWithDefaults() *AWSNamespaceFiltersExcludeOnly {
	this := AWSNamespaceFiltersExcludeOnly{}
	return &this
}

// GetExcludeOnly returns the ExcludeOnly field value.
func (o *AWSNamespaceFiltersExcludeOnly) GetExcludeOnly() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeOnly
}

// GetExcludeOnlyOk returns a tuple with the ExcludeOnly field value
// and a boolean to check if the value has been set.
func (o *AWSNamespaceFiltersExcludeOnly) GetExcludeOnlyOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExcludeOnly, true
}

// SetExcludeOnly sets field value.
func (o *AWSNamespaceFiltersExcludeOnly) SetExcludeOnly(v []string) {
	o.ExcludeOnly = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSNamespaceFiltersExcludeOnly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["exclude_only"] = o.ExcludeOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSNamespaceFiltersExcludeOnly) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExcludeOnly *[]string `json:"exclude_only"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ExcludeOnly == nil {
		return fmt.Errorf("required field exclude_only missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"exclude_only"})
	} else {
		return err
	}
	o.ExcludeOnly = *all.ExcludeOnly

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
