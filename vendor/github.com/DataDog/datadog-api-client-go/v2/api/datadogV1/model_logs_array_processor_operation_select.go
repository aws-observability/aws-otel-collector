// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayProcessorOperationSelect Operation that finds an object in a `source` array using a `filter`, and then extracts a specific value into the `target` attribute.
type LogsArrayProcessorOperationSelect struct {
	// Filter condition expressed as `key:value` used to find the matching element.
	Filter string `json:"filter"`
	// Attribute path of the array to search into.
	Source string `json:"source"`
	// Attribute that receives the extracted value.
	Target string `json:"target"`
	// Operation type.
	Type LogsArrayProcessorOperationSelectType `json:"type"`
	// Key of the value to extract from the matching element.
	ValueToExtract string `json:"value_to_extract"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArrayProcessorOperationSelect instantiates a new LogsArrayProcessorOperationSelect object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArrayProcessorOperationSelect(filter string, source string, target string, typeVar LogsArrayProcessorOperationSelectType, valueToExtract string) *LogsArrayProcessorOperationSelect {
	this := LogsArrayProcessorOperationSelect{}
	this.Filter = filter
	this.Source = source
	this.Target = target
	this.Type = typeVar
	this.ValueToExtract = valueToExtract
	return &this
}

// NewLogsArrayProcessorOperationSelectWithDefaults instantiates a new LogsArrayProcessorOperationSelect object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArrayProcessorOperationSelectWithDefaults() *LogsArrayProcessorOperationSelect {
	this := LogsArrayProcessorOperationSelect{}
	return &this
}

// GetFilter returns the Filter field value.
func (o *LogsArrayProcessorOperationSelect) GetFilter() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationSelect) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value.
func (o *LogsArrayProcessorOperationSelect) SetFilter(v string) {
	o.Filter = v
}

// GetSource returns the Source field value.
func (o *LogsArrayProcessorOperationSelect) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationSelect) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *LogsArrayProcessorOperationSelect) SetSource(v string) {
	o.Source = v
}

// GetTarget returns the Target field value.
func (o *LogsArrayProcessorOperationSelect) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationSelect) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *LogsArrayProcessorOperationSelect) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *LogsArrayProcessorOperationSelect) GetType() LogsArrayProcessorOperationSelectType {
	if o == nil {
		var ret LogsArrayProcessorOperationSelectType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationSelect) GetTypeOk() (*LogsArrayProcessorOperationSelectType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsArrayProcessorOperationSelect) SetType(v LogsArrayProcessorOperationSelectType) {
	o.Type = v
}

// GetValueToExtract returns the ValueToExtract field value.
func (o *LogsArrayProcessorOperationSelect) GetValueToExtract() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ValueToExtract
}

// GetValueToExtractOk returns a tuple with the ValueToExtract field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessorOperationSelect) GetValueToExtractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueToExtract, true
}

// SetValueToExtract sets field value.
func (o *LogsArrayProcessorOperationSelect) SetValueToExtract(v string) {
	o.ValueToExtract = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArrayProcessorOperationSelect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["filter"] = o.Filter
	toSerialize["source"] = o.Source
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type
	toSerialize["value_to_extract"] = o.ValueToExtract

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArrayProcessorOperationSelect) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter         *string                                `json:"filter"`
		Source         *string                                `json:"source"`
		Target         *string                                `json:"target"`
		Type           *LogsArrayProcessorOperationSelectType `json:"type"`
		ValueToExtract *string                                `json:"value_to_extract"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Filter == nil {
		return fmt.Errorf("required field filter missing")
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
	if all.ValueToExtract == nil {
		return fmt.Errorf("required field value_to_extract missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter", "source", "target", "type", "value_to_extract"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Filter = *all.Filter
	o.Source = *all.Source
	o.Target = *all.Target
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.ValueToExtract = *all.ValueToExtract

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
