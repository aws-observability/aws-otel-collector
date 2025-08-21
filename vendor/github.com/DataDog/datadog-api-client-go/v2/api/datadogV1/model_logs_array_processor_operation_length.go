// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayProcessorOperationLength Operation that computes the length of a `source` array and stores the result in the `target` attribute.
type LogsArrayProcessorOperationLength struct {
	// Attribute path of the array to measure.
	Source string `json:"source"`
	// Attribute that receives the computed length.
	Target string `json:"target"`
	// Operation type.
	Type LogsArrayProcessorOperationLengthType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArrayProcessorOperationLength instantiates a new LogsArrayProcessorOperationLength object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArrayProcessorOperationLength(source string, target string, typeVar LogsArrayProcessorOperationLengthType) *LogsArrayProcessorOperationLength {
	this := LogsArrayProcessorOperationLength{}
	this.Source = source
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewLogsArrayProcessorOperationLengthWithDefaults instantiates a new LogsArrayProcessorOperationLength object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArrayProcessorOperationLengthWithDefaults() *LogsArrayProcessorOperationLength {
	this := LogsArrayProcessorOperationLength{}
	return &this
}

// GetSource returns the Source field value.
func (o *LogsArrayProcessorOperationLength) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationLength) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *LogsArrayProcessorOperationLength) SetSource(v string) {
	o.Source = v
}

// GetTarget returns the Target field value.
func (o *LogsArrayProcessorOperationLength) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationLength) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *LogsArrayProcessorOperationLength) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *LogsArrayProcessorOperationLength) GetType() LogsArrayProcessorOperationLengthType {
	if o == nil {
		var ret LogsArrayProcessorOperationLengthType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationLength) GetTypeOk() (*LogsArrayProcessorOperationLengthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsArrayProcessorOperationLength) SetType(v LogsArrayProcessorOperationLengthType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArrayProcessorOperationLength) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["source"] = o.Source
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArrayProcessorOperationLength) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Source *string                                `json:"source"`
		Target *string                                `json:"target"`
		Type   *LogsArrayProcessorOperationLengthType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"source", "target", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Source = *all.Source
	o.Target = *all.Target
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
