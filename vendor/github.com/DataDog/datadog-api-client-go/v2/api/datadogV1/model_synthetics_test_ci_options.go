// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestCiOptions CI/CD options for a Synthetic test.
type SyntheticsTestCiOptions struct {
	// Execution rule for a Synthetic test.
	ExecutionRule *SyntheticsTestExecutionRule `json:"executionRule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSyntheticsTestCiOptions instantiates a new SyntheticsTestCiOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestCiOptions() *SyntheticsTestCiOptions {
	this := SyntheticsTestCiOptions{}
	return &this
}

// NewSyntheticsTestCiOptionsWithDefaults instantiates a new SyntheticsTestCiOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestCiOptionsWithDefaults() *SyntheticsTestCiOptions {
	this := SyntheticsTestCiOptions{}
	return &this
}

// GetExecutionRule returns the ExecutionRule field value if set, zero value otherwise.
func (o *SyntheticsTestCiOptions) GetExecutionRule() SyntheticsTestExecutionRule {
	if o == nil || o.ExecutionRule == nil {
		var ret SyntheticsTestExecutionRule
		return ret
	}
	return *o.ExecutionRule
}

// GetExecutionRuleOk returns a tuple with the ExecutionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestCiOptions) GetExecutionRuleOk() (*SyntheticsTestExecutionRule, bool) {
	if o == nil || o.ExecutionRule == nil {
		return nil, false
	}
	return o.ExecutionRule, true
}

// HasExecutionRule returns a boolean if a field has been set.
func (o *SyntheticsTestCiOptions) HasExecutionRule() bool {
	return o != nil && o.ExecutionRule != nil
}

// SetExecutionRule gets a reference to the given SyntheticsTestExecutionRule and assigns it to the ExecutionRule field.
func (o *SyntheticsTestCiOptions) SetExecutionRule(v SyntheticsTestExecutionRule) {
	o.ExecutionRule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestCiOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ExecutionRule != nil {
		toSerialize["executionRule"] = o.ExecutionRule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestCiOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExecutionRule *SyntheticsTestExecutionRule `json:"executionRule,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"executionRule"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ExecutionRule != nil && !all.ExecutionRule.IsValid() {
		hasInvalidField = true
	} else {
		o.ExecutionRule = all.ExecutionRule
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
