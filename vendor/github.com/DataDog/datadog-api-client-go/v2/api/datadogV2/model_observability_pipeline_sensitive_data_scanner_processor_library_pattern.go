// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern Specifies a pattern from Datadog’s sensitive data detection library to match known sensitive data types.
type ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern struct {
	// Options for selecting a predefined library pattern and enabling keyword support.
	Options ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternOptions `json:"options"`
	// Indicates that a predefined library pattern is used.
	Type ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern(options ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternOptions, typeVar ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType) *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern {
	this := ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern{}
	this.Options = options
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternWithDefaults instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternWithDefaults() *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern {
	this := ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern{}
	return &this
}

// GetOptions returns the Options field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern) GetOptions() ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternOptions {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern) GetOptionsOk() (*ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern) SetOptions(v ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternOptions) {
	o.Options = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern) GetType() ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern) GetTypeOk() (*ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern) SetType(v ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern) MarshalJSON() ([]byte, error) {
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
func (o *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPattern) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Options *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternOptions `json:"options"`
		Type    *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType    `json:"type"`
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
