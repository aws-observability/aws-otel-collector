// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorActionHash Configuration for hashing matched sensitive values.
type ObservabilityPipelineSensitiveDataScannerProcessorActionHash struct {
	// Action type that replaces the matched sensitive data with a hashed representation, preserving structure while securing content.
	Action ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction `json:"action"`
	// The `ObservabilityPipelineSensitiveDataScannerProcessorActionHash` `options`.
	Options interface{} `json:"options,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSensitiveDataScannerProcessorActionHash instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorActionHash object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSensitiveDataScannerProcessorActionHash(action ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction) *ObservabilityPipelineSensitiveDataScannerProcessorActionHash {
	this := ObservabilityPipelineSensitiveDataScannerProcessorActionHash{}
	this.Action = action
	return &this
}

// NewObservabilityPipelineSensitiveDataScannerProcessorActionHashWithDefaults instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorActionHash object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSensitiveDataScannerProcessorActionHashWithDefaults() *ObservabilityPipelineSensitiveDataScannerProcessorActionHash {
	this := ObservabilityPipelineSensitiveDataScannerProcessorActionHash{}
	return &this
}

// GetAction returns the Action field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionHash) GetAction() ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionHash) GetActionOk() (*ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionHash) SetAction(v ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction) {
	o.Action = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionHash) GetOptions() interface{} {
	if o == nil || o.Options == nil {
		var ret interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionHash) GetOptionsOk() (*interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionHash) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given interface{} and assigns it to the Options field.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionHash) SetOptions(v interface{}) {
	o.Options = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSensitiveDataScannerProcessorActionHash) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionHash) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action  *ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction `json:"action"`
		Options interface{}                                                         `json:"options,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
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
	o.Options = all.Options

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
