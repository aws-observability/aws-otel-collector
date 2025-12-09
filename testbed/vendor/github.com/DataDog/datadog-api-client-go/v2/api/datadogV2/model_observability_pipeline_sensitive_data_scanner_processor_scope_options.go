// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions Fields to which the scope rule applies.
type ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions struct {
	// The `ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions` `fields`.
	Fields []string `json:"fields"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSensitiveDataScannerProcessorScopeOptions instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSensitiveDataScannerProcessorScopeOptions(fields []string) *ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions {
	this := ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions{}
	this.Fields = fields
	return &this
}

// NewObservabilityPipelineSensitiveDataScannerProcessorScopeOptionsWithDefaults instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSensitiveDataScannerProcessorScopeOptionsWithDefaults() *ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions {
	this := ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions{}
	return &this
}

// GetFields returns the Fields field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions) GetFieldsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions) SetFields(v []string) {
	o.Fields = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["fields"] = o.Fields

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields *[]string `json:"fields"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Fields == nil {
		return fmt.Errorf("required field fields missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields"})
	} else {
		return err
	}
	o.Fields = *all.Fields

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
