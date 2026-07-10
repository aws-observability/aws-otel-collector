// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisRuleResponse The result of applying a single static analysis rule to the analyzed source code.
type AnalysisRuleResponse struct {
	// A list of error messages encountered while executing the rule.
	Errors []string `json:"errors"`
	// An error message if the rule execution failed, or null if execution succeeded.
	ExecutionError datadog.NullableString `json:"execution_error"`
	// The time taken to execute the rule, in milliseconds.
	ExecutionTimeMs int64 `json:"execution_time_ms"`
	// The identifier of the rule that produced this response.
	Identifier string `json:"identifier"`
	// The raw output produced by the rule engine during execution.
	Output string `json:"output"`
	// The list of violations found by this rule.
	Violations []AnalysisViolation `json:"violations"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnalysisRuleResponse instantiates a new AnalysisRuleResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnalysisRuleResponse(errors []string, executionError datadog.NullableString, executionTimeMs int64, identifier string, output string, violations []AnalysisViolation) *AnalysisRuleResponse {
	this := AnalysisRuleResponse{}
	this.Errors = errors
	this.ExecutionError = executionError
	this.ExecutionTimeMs = executionTimeMs
	this.Identifier = identifier
	this.Output = output
	this.Violations = violations
	return &this
}

// NewAnalysisRuleResponseWithDefaults instantiates a new AnalysisRuleResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnalysisRuleResponseWithDefaults() *AnalysisRuleResponse {
	this := AnalysisRuleResponse{}
	return &this
}

// GetErrors returns the Errors field value.
func (o *AnalysisRuleResponse) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *AnalysisRuleResponse) GetErrorsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value.
func (o *AnalysisRuleResponse) SetErrors(v []string) {
	o.Errors = v
}

// GetExecutionError returns the ExecutionError field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *AnalysisRuleResponse) GetExecutionError() string {
	if o == nil || o.ExecutionError.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExecutionError.Get()
}

// GetExecutionErrorOk returns a tuple with the ExecutionError field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AnalysisRuleResponse) GetExecutionErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionError.Get(), o.ExecutionError.IsSet()
}

// SetExecutionError sets field value.
func (o *AnalysisRuleResponse) SetExecutionError(v string) {
	o.ExecutionError.Set(&v)
}

// GetExecutionTimeMs returns the ExecutionTimeMs field value.
func (o *AnalysisRuleResponse) GetExecutionTimeMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ExecutionTimeMs
}

// GetExecutionTimeMsOk returns a tuple with the ExecutionTimeMs field value
// and a boolean to check if the value has been set.
func (o *AnalysisRuleResponse) GetExecutionTimeMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionTimeMs, true
}

// SetExecutionTimeMs sets field value.
func (o *AnalysisRuleResponse) SetExecutionTimeMs(v int64) {
	o.ExecutionTimeMs = v
}

// GetIdentifier returns the Identifier field value.
func (o *AnalysisRuleResponse) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *AnalysisRuleResponse) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value.
func (o *AnalysisRuleResponse) SetIdentifier(v string) {
	o.Identifier = v
}

// GetOutput returns the Output field value.
func (o *AnalysisRuleResponse) GetOutput() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value
// and a boolean to check if the value has been set.
func (o *AnalysisRuleResponse) GetOutputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Output, true
}

// SetOutput sets field value.
func (o *AnalysisRuleResponse) SetOutput(v string) {
	o.Output = v
}

// GetViolations returns the Violations field value.
func (o *AnalysisRuleResponse) GetViolations() []AnalysisViolation {
	if o == nil {
		var ret []AnalysisViolation
		return ret
	}
	return o.Violations
}

// GetViolationsOk returns a tuple with the Violations field value
// and a boolean to check if the value has been set.
func (o *AnalysisRuleResponse) GetViolationsOk() (*[]AnalysisViolation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Violations, true
}

// SetViolations sets field value.
func (o *AnalysisRuleResponse) SetViolations(v []AnalysisViolation) {
	o.Violations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnalysisRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["errors"] = o.Errors
	toSerialize["execution_error"] = o.ExecutionError.Get()
	toSerialize["execution_time_ms"] = o.ExecutionTimeMs
	toSerialize["identifier"] = o.Identifier
	toSerialize["output"] = o.Output
	toSerialize["violations"] = o.Violations

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnalysisRuleResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Errors          *[]string              `json:"errors"`
		ExecutionError  datadog.NullableString `json:"execution_error"`
		ExecutionTimeMs *int64                 `json:"execution_time_ms"`
		Identifier      *string                `json:"identifier"`
		Output          *string                `json:"output"`
		Violations      *[]AnalysisViolation   `json:"violations"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Errors == nil {
		return fmt.Errorf("required field errors missing")
	}
	if !all.ExecutionError.IsSet() {
		return fmt.Errorf("required field execution_error missing")
	}
	if all.ExecutionTimeMs == nil {
		return fmt.Errorf("required field execution_time_ms missing")
	}
	if all.Identifier == nil {
		return fmt.Errorf("required field identifier missing")
	}
	if all.Output == nil {
		return fmt.Errorf("required field output missing")
	}
	if all.Violations == nil {
		return fmt.Errorf("required field violations missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"errors", "execution_error", "execution_time_ms", "identifier", "output", "violations"})
	} else {
		return err
	}
	o.Errors = *all.Errors
	o.ExecutionError = all.ExecutionError
	o.ExecutionTimeMs = *all.ExecutionTimeMs
	o.Identifier = *all.Identifier
	o.Output = *all.Output
	o.Violations = *all.Violations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
