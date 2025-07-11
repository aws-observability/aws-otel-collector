// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern Defines a custom regex-based pattern for identifying sensitive data in logs.
type ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern struct {
	// Options for defining a custom regex pattern.
	Options ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions `json:"options"`
	// Indicates a custom regular expression is used for matching.
	Type ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSensitiveDataScannerProcessorCustomPattern instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSensitiveDataScannerProcessorCustomPattern(options ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions, typeVar ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType) *ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern {
	this := ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern{}
	this.Options = options
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineSensitiveDataScannerProcessorCustomPatternWithDefaults instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSensitiveDataScannerProcessorCustomPatternWithDefaults() *ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern {
	this := ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern{}
	return &this
}

// GetOptions returns the Options field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern) GetOptions() ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern) GetOptionsOk() (*ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern) SetOptions(v ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions) {
	o.Options = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern) GetType() ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern) GetTypeOk() (*ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern) SetType(v ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["options"] = o.Options
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorCustomPattern) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Options *ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternOptions `json:"options"`
		Type    *ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Options == nil {
		return fmt.Errorf("required field options missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"options", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = *all.Options
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
