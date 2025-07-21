// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorScopeAll Applies scanning across all available fields.
type ObservabilityPipelineSensitiveDataScannerProcessorScopeAll struct {
	// Applies the rule to all fields.
	Target ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget `json:"target"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSensitiveDataScannerProcessorScopeAll instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorScopeAll object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSensitiveDataScannerProcessorScopeAll(target ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget) *ObservabilityPipelineSensitiveDataScannerProcessorScopeAll {
	this := ObservabilityPipelineSensitiveDataScannerProcessorScopeAll{}
	this.Target = target
	return &this
}

// NewObservabilityPipelineSensitiveDataScannerProcessorScopeAllWithDefaults instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorScopeAll object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSensitiveDataScannerProcessorScopeAllWithDefaults() *ObservabilityPipelineSensitiveDataScannerProcessorScopeAll {
	this := ObservabilityPipelineSensitiveDataScannerProcessorScopeAll{}
	return &this
}

// GetTarget returns the Target field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeAll) GetTarget() ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeAll) GetTargetOk() (*ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeAll) SetTarget(v ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget) {
	o.Target = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSensitiveDataScannerProcessorScopeAll) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["target"] = o.Target

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeAll) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Target *ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget `json:"target"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"target"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Target.IsValid() {
		hasInvalidField = true
	} else {
		o.Target = *all.Target
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
