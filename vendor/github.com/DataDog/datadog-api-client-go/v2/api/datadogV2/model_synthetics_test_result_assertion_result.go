// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultAssertionResult An individual assertion result from a Synthetic test.
type SyntheticsTestResultAssertionResult struct {
	// Actual value observed during the test. Its type depends on the assertion type.
	Actual interface{} `json:"actual,omitempty"`
	// Error message if the assertion failed.
	ErrorMessage *string `json:"error_message,omitempty"`
	// Expected value for the assertion. Its type depends on the assertion type.
	Expected interface{} `json:"expected,omitempty"`
	// Operator used for the assertion (for example, `is`, `contains`).
	Operator *string `json:"operator,omitempty"`
	// Property targeted by the assertion, when applicable.
	Property *string `json:"property,omitempty"`
	// Target value for the assertion. Its type depends on the assertion type.
	Target interface{} `json:"target,omitempty"`
	// JSON path or XPath evaluated for the assertion.
	TargetPath *string `json:"target_path,omitempty"`
	// Operator used for the target path assertion.
	TargetPathOperator *string `json:"target_path_operator,omitempty"`
	// Type of the assertion (for example, `responseTime`, `statusCode`, `body`).
	Type *string `json:"type,omitempty"`
	// Whether the assertion passed.
	Valid *bool `json:"valid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultAssertionResult instantiates a new SyntheticsTestResultAssertionResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultAssertionResult() *SyntheticsTestResultAssertionResult {
	this := SyntheticsTestResultAssertionResult{}
	return &this
}

// NewSyntheticsTestResultAssertionResultWithDefaults instantiates a new SyntheticsTestResultAssertionResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultAssertionResultWithDefaults() *SyntheticsTestResultAssertionResult {
	this := SyntheticsTestResultAssertionResult{}
	return &this
}

// GetActual returns the Actual field value if set, zero value otherwise.
func (o *SyntheticsTestResultAssertionResult) GetActual() interface{} {
	if o == nil || o.Actual == nil {
		var ret interface{}
		return ret
	}
	return o.Actual
}

// GetActualOk returns a tuple with the Actual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAssertionResult) GetActualOk() (*interface{}, bool) {
	if o == nil || o.Actual == nil {
		return nil, false
	}
	return &o.Actual, true
}

// HasActual returns a boolean if a field has been set.
func (o *SyntheticsTestResultAssertionResult) HasActual() bool {
	return o != nil && o.Actual != nil
}

// SetActual gets a reference to the given interface{} and assigns it to the Actual field.
func (o *SyntheticsTestResultAssertionResult) SetActual(v interface{}) {
	o.Actual = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *SyntheticsTestResultAssertionResult) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAssertionResult) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *SyntheticsTestResultAssertionResult) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *SyntheticsTestResultAssertionResult) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetExpected returns the Expected field value if set, zero value otherwise.
func (o *SyntheticsTestResultAssertionResult) GetExpected() interface{} {
	if o == nil || o.Expected == nil {
		var ret interface{}
		return ret
	}
	return o.Expected
}

// GetExpectedOk returns a tuple with the Expected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAssertionResult) GetExpectedOk() (*interface{}, bool) {
	if o == nil || o.Expected == nil {
		return nil, false
	}
	return &o.Expected, true
}

// HasExpected returns a boolean if a field has been set.
func (o *SyntheticsTestResultAssertionResult) HasExpected() bool {
	return o != nil && o.Expected != nil
}

// SetExpected gets a reference to the given interface{} and assigns it to the Expected field.
func (o *SyntheticsTestResultAssertionResult) SetExpected(v interface{}) {
	o.Expected = v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *SyntheticsTestResultAssertionResult) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAssertionResult) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *SyntheticsTestResultAssertionResult) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *SyntheticsTestResultAssertionResult) SetOperator(v string) {
	o.Operator = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *SyntheticsTestResultAssertionResult) GetProperty() string {
	if o == nil || o.Property == nil {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAssertionResult) GetPropertyOk() (*string, bool) {
	if o == nil || o.Property == nil {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *SyntheticsTestResultAssertionResult) HasProperty() bool {
	return o != nil && o.Property != nil
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *SyntheticsTestResultAssertionResult) SetProperty(v string) {
	o.Property = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SyntheticsTestResultAssertionResult) GetTarget() interface{} {
	if o == nil || o.Target == nil {
		var ret interface{}
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAssertionResult) GetTargetOk() (*interface{}, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return &o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SyntheticsTestResultAssertionResult) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given interface{} and assigns it to the Target field.
func (o *SyntheticsTestResultAssertionResult) SetTarget(v interface{}) {
	o.Target = v
}

// GetTargetPath returns the TargetPath field value if set, zero value otherwise.
func (o *SyntheticsTestResultAssertionResult) GetTargetPath() string {
	if o == nil || o.TargetPath == nil {
		var ret string
		return ret
	}
	return *o.TargetPath
}

// GetTargetPathOk returns a tuple with the TargetPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAssertionResult) GetTargetPathOk() (*string, bool) {
	if o == nil || o.TargetPath == nil {
		return nil, false
	}
	return o.TargetPath, true
}

// HasTargetPath returns a boolean if a field has been set.
func (o *SyntheticsTestResultAssertionResult) HasTargetPath() bool {
	return o != nil && o.TargetPath != nil
}

// SetTargetPath gets a reference to the given string and assigns it to the TargetPath field.
func (o *SyntheticsTestResultAssertionResult) SetTargetPath(v string) {
	o.TargetPath = &v
}

// GetTargetPathOperator returns the TargetPathOperator field value if set, zero value otherwise.
func (o *SyntheticsTestResultAssertionResult) GetTargetPathOperator() string {
	if o == nil || o.TargetPathOperator == nil {
		var ret string
		return ret
	}
	return *o.TargetPathOperator
}

// GetTargetPathOperatorOk returns a tuple with the TargetPathOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAssertionResult) GetTargetPathOperatorOk() (*string, bool) {
	if o == nil || o.TargetPathOperator == nil {
		return nil, false
	}
	return o.TargetPathOperator, true
}

// HasTargetPathOperator returns a boolean if a field has been set.
func (o *SyntheticsTestResultAssertionResult) HasTargetPathOperator() bool {
	return o != nil && o.TargetPathOperator != nil
}

// SetTargetPathOperator gets a reference to the given string and assigns it to the TargetPathOperator field.
func (o *SyntheticsTestResultAssertionResult) SetTargetPathOperator(v string) {
	o.TargetPathOperator = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsTestResultAssertionResult) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAssertionResult) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsTestResultAssertionResult) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SyntheticsTestResultAssertionResult) SetType(v string) {
	o.Type = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *SyntheticsTestResultAssertionResult) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAssertionResult) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *SyntheticsTestResultAssertionResult) HasValid() bool {
	return o != nil && o.Valid != nil
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *SyntheticsTestResultAssertionResult) SetValid(v bool) {
	o.Valid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultAssertionResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Actual != nil {
		toSerialize["actual"] = o.Actual
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.Expected != nil {
		toSerialize["expected"] = o.Expected
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.Property != nil {
		toSerialize["property"] = o.Property
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.TargetPath != nil {
		toSerialize["target_path"] = o.TargetPath
	}
	if o.TargetPathOperator != nil {
		toSerialize["target_path_operator"] = o.TargetPathOperator
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultAssertionResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Actual             interface{} `json:"actual,omitempty"`
		ErrorMessage       *string     `json:"error_message,omitempty"`
		Expected           interface{} `json:"expected,omitempty"`
		Operator           *string     `json:"operator,omitempty"`
		Property           *string     `json:"property,omitempty"`
		Target             interface{} `json:"target,omitempty"`
		TargetPath         *string     `json:"target_path,omitempty"`
		TargetPathOperator *string     `json:"target_path_operator,omitempty"`
		Type               *string     `json:"type,omitempty"`
		Valid              *bool       `json:"valid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actual", "error_message", "expected", "operator", "property", "target", "target_path", "target_path_operator", "type", "valid"})
	} else {
		return err
	}
	o.Actual = all.Actual
	o.ErrorMessage = all.ErrorMessage
	o.Expected = all.Expected
	o.Operator = all.Operator
	o.Property = all.Property
	o.Target = all.Target
	o.TargetPath = all.TargetPath
	o.TargetPathOperator = all.TargetPathOperator
	o.Type = all.Type
	o.Valid = all.Valid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
