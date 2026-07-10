// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayMapArithmeticSubProcessor An arithmetic sub-processor for use inside an array-map processor.
// Unlike the top-level arithmetic processor, `is_enabled` is not supported.
type LogsArrayMapArithmeticSubProcessor struct {
	// Arithmetic operation to perform.
	Expression string `json:"expression"`
	// Replace missing attribute values with 0.
	IsReplaceMissing *bool `json:"is_replace_missing,omitempty"`
	// Name of the sub-processor.
	Name *string `json:"name,omitempty"`
	// Target attribute path for the result.
	Target string `json:"target"`
	// Type of logs arithmetic processor.
	Type LogsArithmeticProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArrayMapArithmeticSubProcessor instantiates a new LogsArrayMapArithmeticSubProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArrayMapArithmeticSubProcessor(expression string, target string, typeVar LogsArithmeticProcessorType) *LogsArrayMapArithmeticSubProcessor {
	this := LogsArrayMapArithmeticSubProcessor{}
	this.Expression = expression
	var isReplaceMissing bool = false
	this.IsReplaceMissing = &isReplaceMissing
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewLogsArrayMapArithmeticSubProcessorWithDefaults instantiates a new LogsArrayMapArithmeticSubProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArrayMapArithmeticSubProcessorWithDefaults() *LogsArrayMapArithmeticSubProcessor {
	this := LogsArrayMapArithmeticSubProcessor{}
	var isReplaceMissing bool = false
	this.IsReplaceMissing = &isReplaceMissing
	var typeVar LogsArithmeticProcessorType = LOGSARITHMETICPROCESSORTYPE_ARITHMETIC_PROCESSOR
	this.Type = typeVar
	return &this
}

// GetExpression returns the Expression field value.
func (o *LogsArrayMapArithmeticSubProcessor) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *LogsArrayMapArithmeticSubProcessor) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value.
func (o *LogsArrayMapArithmeticSubProcessor) SetExpression(v string) {
	o.Expression = v
}

// GetIsReplaceMissing returns the IsReplaceMissing field value if set, zero value otherwise.
func (o *LogsArrayMapArithmeticSubProcessor) GetIsReplaceMissing() bool {
	if o == nil || o.IsReplaceMissing == nil {
		var ret bool
		return ret
	}
	return *o.IsReplaceMissing
}

// GetIsReplaceMissingOk returns a tuple with the IsReplaceMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArrayMapArithmeticSubProcessor) GetIsReplaceMissingOk() (*bool, bool) {
	if o == nil || o.IsReplaceMissing == nil {
		return nil, false
	}
	return o.IsReplaceMissing, true
}

// HasIsReplaceMissing returns a boolean if a field has been set.
func (o *LogsArrayMapArithmeticSubProcessor) HasIsReplaceMissing() bool {
	return o != nil && o.IsReplaceMissing != nil
}

// SetIsReplaceMissing gets a reference to the given bool and assigns it to the IsReplaceMissing field.
func (o *LogsArrayMapArithmeticSubProcessor) SetIsReplaceMissing(v bool) {
	o.IsReplaceMissing = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsArrayMapArithmeticSubProcessor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArrayMapArithmeticSubProcessor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsArrayMapArithmeticSubProcessor) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsArrayMapArithmeticSubProcessor) SetName(v string) {
	o.Name = &v
}

// GetTarget returns the Target field value.
func (o *LogsArrayMapArithmeticSubProcessor) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LogsArrayMapArithmeticSubProcessor) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *LogsArrayMapArithmeticSubProcessor) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *LogsArrayMapArithmeticSubProcessor) GetType() LogsArithmeticProcessorType {
	if o == nil {
		var ret LogsArithmeticProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArrayMapArithmeticSubProcessor) GetTypeOk() (*LogsArithmeticProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsArrayMapArithmeticSubProcessor) SetType(v LogsArithmeticProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArrayMapArithmeticSubProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["expression"] = o.Expression
	if o.IsReplaceMissing != nil {
		toSerialize["is_replace_missing"] = o.IsReplaceMissing
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArrayMapArithmeticSubProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Expression       *string                      `json:"expression"`
		IsReplaceMissing *bool                        `json:"is_replace_missing,omitempty"`
		Name             *string                      `json:"name,omitempty"`
		Target           *string                      `json:"target"`
		Type             *LogsArithmeticProcessorType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Expression == nil {
		return fmt.Errorf("required field expression missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"expression", "is_replace_missing", "name", "target", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Expression = *all.Expression
	o.IsReplaceMissing = all.IsReplaceMissing
	o.Name = all.Name
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
