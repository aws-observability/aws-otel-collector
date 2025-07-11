// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorActionRedact Configuration for completely redacting matched sensitive data.
type ObservabilityPipelineSensitiveDataScannerProcessorActionRedact struct {
	// Action type that completely replaces the matched sensitive data with a fixed replacement string to remove all visibility.
	Action ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction `json:"action"`
	// Configuration for fully redacting sensitive data.
	Options ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions `json:"options"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSensitiveDataScannerProcessorActionRedact instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorActionRedact object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSensitiveDataScannerProcessorActionRedact(action ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction, options ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions) *ObservabilityPipelineSensitiveDataScannerProcessorActionRedact {
	this := ObservabilityPipelineSensitiveDataScannerProcessorActionRedact{}
	this.Action = action
	this.Options = options
	return &this
}

// NewObservabilityPipelineSensitiveDataScannerProcessorActionRedactWithDefaults instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorActionRedact object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSensitiveDataScannerProcessorActionRedactWithDefaults() *ObservabilityPipelineSensitiveDataScannerProcessorActionRedact {
	this := ObservabilityPipelineSensitiveDataScannerProcessorActionRedact{}
	return &this
}

// GetAction returns the Action field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionRedact) GetAction() ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionRedact) GetActionOk() (*ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionRedact) SetAction(v ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction) {
	o.Action = v
}

// GetOptions returns the Options field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionRedact) GetOptions() ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionRedact) GetOptionsOk() (*ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionRedact) SetOptions(v ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions) {
	o.Options = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSensitiveDataScannerProcessorActionRedact) MarshalJSON() ([]byte, error) {
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
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionRedact) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action  *ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction  `json:"action"`
		Options *ObservabilityPipelineSensitiveDataScannerProcessorActionRedactOptions `json:"options"`
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
