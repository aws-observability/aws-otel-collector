// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayProcessorOperationAppend Operation that appends a value to a target array attribute.
type LogsArrayProcessorOperationAppend struct {
	// Remove or preserve the remapped source element.
	PreserveSource *bool `json:"preserve_source,omitempty"`
	// Attribute path containing the value to append.
	Source string `json:"source"`
	// Attribute path of the array to append to.
	Target string `json:"target"`
	// Operation type.
	Type LogsArrayProcessorOperationAppendType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArrayProcessorOperationAppend instantiates a new LogsArrayProcessorOperationAppend object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArrayProcessorOperationAppend(source string, target string, typeVar LogsArrayProcessorOperationAppendType) *LogsArrayProcessorOperationAppend {
	this := LogsArrayProcessorOperationAppend{}
	var preserveSource bool = true
	this.PreserveSource = &preserveSource
	this.Source = source
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewLogsArrayProcessorOperationAppendWithDefaults instantiates a new LogsArrayProcessorOperationAppend object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArrayProcessorOperationAppendWithDefaults() *LogsArrayProcessorOperationAppend {
	this := LogsArrayProcessorOperationAppend{}
	var preserveSource bool = true
	this.PreserveSource = &preserveSource
	return &this
}

// GetPreserveSource returns the PreserveSource field value if set, zero value otherwise.
func (o *LogsArrayProcessorOperationAppend) GetPreserveSource() bool {
	if o == nil || o.PreserveSource == nil {
		var ret bool
		return ret
	}
	return *o.PreserveSource
}

// GetPreserveSourceOk returns a tuple with the PreserveSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationAppend) GetPreserveSourceOk() (*bool, bool) {
	if o == nil || o.PreserveSource == nil {
		return nil, false
	}
	return o.PreserveSource, true
}

// HasPreserveSource returns a boolean if a field has been set.
func (o *LogsArrayProcessorOperationAppend) HasPreserveSource() bool {
	return o != nil && o.PreserveSource != nil
}

// SetPreserveSource gets a reference to the given bool and assigns it to the PreserveSource field.
func (o *LogsArrayProcessorOperationAppend) SetPreserveSource(v bool) {
	o.PreserveSource = &v
}

// GetSource returns the Source field value.
func (o *LogsArrayProcessorOperationAppend) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationAppend) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *LogsArrayProcessorOperationAppend) SetSource(v string) {
	o.Source = v
}

// GetTarget returns the Target field value.
func (o *LogsArrayProcessorOperationAppend) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationAppend) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *LogsArrayProcessorOperationAppend) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *LogsArrayProcessorOperationAppend) GetType() LogsArrayProcessorOperationAppendType {
	if o == nil {
		var ret LogsArrayProcessorOperationAppendType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationAppend) GetTypeOk() (*LogsArrayProcessorOperationAppendType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsArrayProcessorOperationAppend) SetType(v LogsArrayProcessorOperationAppendType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArrayProcessorOperationAppend) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.PreserveSource != nil {
		toSerialize["preserve_source"] = o.PreserveSource
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
func (o *LogsArrayProcessorOperationAppend) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PreserveSource *bool                                  `json:"preserve_source,omitempty"`
		Source         *string                                `json:"source"`
		Target         *string                                `json:"target"`
		Type           *LogsArrayProcessorOperationAppendType `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"preserve_source", "source", "target", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.PreserveSource = all.PreserveSource
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
