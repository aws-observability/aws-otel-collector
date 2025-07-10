// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions Configuration for fully redacting sensitive data.
type ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions struct {
	// The `ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions` `replace`.
	Replace string `json:"replace"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions(replace string) *ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions {
	this := ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions{}
	this.Replace = replace
	return &this
}

// NewObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptionsWithDefaults instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptionsWithDefaults() *ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions {
	this := ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions{}
	return &this
}

// GetReplace returns the Replace field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions) GetReplace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions) GetReplaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replace, true
}

// SetReplace sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions) SetReplace(v string) {
	o.Replace = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["replace"] = o.Replace

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Replace *string `json:"replace"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Replace == nil {
		return fmt.Errorf("required field replace missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"replace"})
	} else {
		return err
	}
	o.Replace = *all.Replace

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
