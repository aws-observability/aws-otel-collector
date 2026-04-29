// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplitArrayProcessorArrayConfig Configuration for a single array split operation.
type ObservabilityPipelineSplitArrayProcessorArrayConfig struct {
	// The path to the array field to split.
	Field string `json:"field"`
	// A Datadog search query used to determine which logs this array split operation targets.
	Include string `json:"include"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSplitArrayProcessorArrayConfig instantiates a new ObservabilityPipelineSplitArrayProcessorArrayConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSplitArrayProcessorArrayConfig(field string, include string) *ObservabilityPipelineSplitArrayProcessorArrayConfig {
	this := ObservabilityPipelineSplitArrayProcessorArrayConfig{}
	this.Field = field
	this.Include = include
	return &this
}

// NewObservabilityPipelineSplitArrayProcessorArrayConfigWithDefaults instantiates a new ObservabilityPipelineSplitArrayProcessorArrayConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSplitArrayProcessorArrayConfigWithDefaults() *ObservabilityPipelineSplitArrayProcessorArrayConfig {
	this := ObservabilityPipelineSplitArrayProcessorArrayConfig{}
	return &this
}

// GetField returns the Field field value.
func (o *ObservabilityPipelineSplitArrayProcessorArrayConfig) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplitArrayProcessorArrayConfig) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *ObservabilityPipelineSplitArrayProcessorArrayConfig) SetField(v string) {
	o.Field = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineSplitArrayProcessorArrayConfig) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplitArrayProcessorArrayConfig) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineSplitArrayProcessorArrayConfig) SetInclude(v string) {
	o.Include = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSplitArrayProcessorArrayConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["field"] = o.Field
	toSerialize["include"] = o.Include

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSplitArrayProcessorArrayConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Field   *string `json:"field"`
		Include *string `json:"include"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"field", "include"})
	} else {
		return err
	}
	o.Field = *all.Field
	o.Include = *all.Include

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
