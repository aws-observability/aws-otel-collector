// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude Excludes specific fields from sensitive data scanning.
type ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude struct {
	// Fields to which the scope rule applies.
	Options ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions `json:"options"`
	// Excludes specific fields from processing.
	Target ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget `json:"target"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSensitiveDataScannerProcessorScopeExclude instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSensitiveDataScannerProcessorScopeExclude(options ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions, target ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget) *ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude {
	this := ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude{}
	this.Options = options
	this.Target = target
	return &this
}

// NewObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeWithDefaults instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeWithDefaults() *ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude {
	this := ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude{}
	return &this
}

// GetOptions returns the Options field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude) GetOptions() ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude) GetOptionsOk() (*ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude) SetOptions(v ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions) {
	o.Options = v
}

// GetTarget returns the Target field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude) GetTarget() ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude) GetTargetOk() (*ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude) SetTarget(v ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget) {
	o.Target = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["options"] = o.Options
	toSerialize["target"] = o.Target

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Options *ObservabilityPipelineSensitiveDataScannerProcessorScopeOptions       `json:"options"`
		Target  *ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget `json:"target"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Options == nil {
		return fmt.Errorf("required field options missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"options", "target"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = *all.Options
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
