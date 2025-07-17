// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact Configuration for partially redacting matched sensitive data.
type ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact struct {
	// Action type that redacts part of the sensitive data while preserving a configurable number of characters, typically used for masking purposes (e.g., show last 4 digits of a credit card).
	Action ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction `json:"action"`
	// Controls how partial redaction is applied, including character count and direction.
	Options ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions `json:"options"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact(action ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction, options ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions) *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact {
	this := ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact{}
	this.Action = action
	this.Options = options
	return &this
}

// NewObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactWithDefaults instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactWithDefaults() *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact {
	this := ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact{}
	return &this
}

// GetAction returns the Action field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact) GetAction() ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact) GetActionOk() (*ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact) SetAction(v ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction) {
	o.Action = v
}

// GetOptions returns the Options field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact) GetOptions() ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact) GetOptionsOk() (*ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact) SetOptions(v ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions) {
	o.Options = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["options"] = o.Options

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action  *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction  `json:"action"`
		Options *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions `json:"options"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.Options == nil {
		return fmt.Errorf("required field options missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "options"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	if all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = *all.Options

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
