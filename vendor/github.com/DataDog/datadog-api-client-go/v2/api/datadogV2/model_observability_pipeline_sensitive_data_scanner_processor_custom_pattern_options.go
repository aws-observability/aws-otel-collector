// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions Options for defining a custom regex pattern.
type ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions struct {
	// A regular expression used to detect sensitive values. Must be a valid regex.
	Rule string `json:"rule"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions(rule string) *ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions {
	this := ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions{}
	this.Rule = rule
	return &this
}

// NewObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptionsWithDefaults instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptionsWithDefaults() *ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions {
	this := ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions{}
	return &this
}

// GetRule returns the Rule field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions) GetRule() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions) GetRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions) SetRule(v string) {
	o.Rule = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["rule"] = o.Rule

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rule *string `json:"rule"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Rule == nil {
		return fmt.Errorf("required field rule missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rule"})
	} else {
		return err
	}
	o.Rule = *all.Rule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
