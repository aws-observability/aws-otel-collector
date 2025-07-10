// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompletionCondition The definition of `CompletionCondition` object.
type CompletionCondition struct {
	// The `CompletionCondition` `operand1`.
	Operand1 interface{} `json:"operand1"`
	// The `CompletionCondition` `operand2`.
	Operand2 interface{} `json:"operand2,omitempty"`
	// The definition of `CompletionConditionOperator` object.
	Operator CompletionConditionOperator `json:"operator"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCompletionCondition instantiates a new CompletionCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCompletionCondition(operand1 interface{}, operator CompletionConditionOperator) *CompletionCondition {
	this := CompletionCondition{}
	this.Operand1 = operand1
	this.Operator = operator
	return &this
}

// NewCompletionConditionWithDefaults instantiates a new CompletionCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCompletionConditionWithDefaults() *CompletionCondition {
	this := CompletionCondition{}
	return &this
}

// GetOperand1 returns the Operand1 field value.
func (o *CompletionCondition) GetOperand1() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Operand1
}

// GetOperand1Ok returns a tuple with the Operand1 field value
// and a boolean to check if the value has been set.
func (o *CompletionCondition) GetOperand1Ok() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operand1, true
}

// SetOperand1 sets field value.
func (o *CompletionCondition) SetOperand1(v interface{}) {
	o.Operand1 = v
}

// GetOperand2 returns the Operand2 field value if set, zero value otherwise.
func (o *CompletionCondition) GetOperand2() interface{} {
	if o == nil || o.Operand2 == nil {
		var ret interface{}
		return ret
	}
	return o.Operand2
}

// GetOperand2Ok returns a tuple with the Operand2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletionCondition) GetOperand2Ok() (*interface{}, bool) {
	if o == nil || o.Operand2 == nil {
		return nil, false
	}
	return &o.Operand2, true
}

// HasOperand2 returns a boolean if a field has been set.
func (o *CompletionCondition) HasOperand2() bool {
	return o != nil && o.Operand2 != nil
}

// SetOperand2 gets a reference to the given interface{} and assigns it to the Operand2 field.
func (o *CompletionCondition) SetOperand2(v interface{}) {
	o.Operand2 = v
}

// GetOperator returns the Operator field value.
func (o *CompletionCondition) GetOperator() CompletionConditionOperator {
	if o == nil {
		var ret CompletionConditionOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *CompletionCondition) GetOperatorOk() (*CompletionConditionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *CompletionCondition) SetOperator(v CompletionConditionOperator) {
	o.Operator = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CompletionCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["operand1"] = o.Operand1
	if o.Operand2 != nil {
		toSerialize["operand2"] = o.Operand2
	}
	toSerialize["operator"] = o.Operator

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CompletionCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Operand1 *interface{}                 `json:"operand1"`
		Operand2 interface{}                  `json:"operand2,omitempty"`
		Operator *CompletionConditionOperator `json:"operator"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Operand1 == nil {
		return fmt.Errorf("required field operand1 missing")
	}
	if all.Operator == nil {
		return fmt.Errorf("required field operator missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"operand1", "operand2", "operator"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Operand1 = *all.Operand1
	o.Operand2 = all.Operand2
	if !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = *all.Operator
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
