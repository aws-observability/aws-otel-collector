// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionJSONPathTargetTarget Composed target for `validatesJSONPath` operator.
type SyntheticsAssertionJSONPathTargetTarget struct {
	// The element from the list of results to assert on.  To choose from the first element in the list `firstElementMatches`, every element in the list `everyElementMatches`, at least one element in the list `atLeastOneElementMatches` or the serialized value of the list `serializationMatches`.
	ElementsOperator *string `json:"elementsOperator,omitempty"`
	// The JSON path to assert.
	JsonPath *string `json:"jsonPath,omitempty"`
	// The specific operator to use on the path.
	Operator *string `json:"operator,omitempty"`
	// Value used by the operator in assertions. Can be either a number or string.
	TargetValue *SyntheticsAssertionTargetValue `json:"targetValue,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsAssertionJSONPathTargetTarget instantiates a new SyntheticsAssertionJSONPathTargetTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsAssertionJSONPathTargetTarget() *SyntheticsAssertionJSONPathTargetTarget {
	this := SyntheticsAssertionJSONPathTargetTarget{}
	return &this
}

// NewSyntheticsAssertionJSONPathTargetTargetWithDefaults instantiates a new SyntheticsAssertionJSONPathTargetTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsAssertionJSONPathTargetTargetWithDefaults() *SyntheticsAssertionJSONPathTargetTarget {
	this := SyntheticsAssertionJSONPathTargetTarget{}
	return &this
}

// GetElementsOperator returns the ElementsOperator field value if set, zero value otherwise.
func (o *SyntheticsAssertionJSONPathTargetTarget) GetElementsOperator() string {
	if o == nil || o.ElementsOperator == nil {
		var ret string
		return ret
	}
	return *o.ElementsOperator
}

// GetElementsOperatorOk returns a tuple with the ElementsOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionJSONPathTargetTarget) GetElementsOperatorOk() (*string, bool) {
	if o == nil || o.ElementsOperator == nil {
		return nil, false
	}
	return o.ElementsOperator, true
}

// HasElementsOperator returns a boolean if a field has been set.
func (o *SyntheticsAssertionJSONPathTargetTarget) HasElementsOperator() bool {
	return o != nil && o.ElementsOperator != nil
}

// SetElementsOperator gets a reference to the given string and assigns it to the ElementsOperator field.
func (o *SyntheticsAssertionJSONPathTargetTarget) SetElementsOperator(v string) {
	o.ElementsOperator = &v
}

// GetJsonPath returns the JsonPath field value if set, zero value otherwise.
func (o *SyntheticsAssertionJSONPathTargetTarget) GetJsonPath() string {
	if o == nil || o.JsonPath == nil {
		var ret string
		return ret
	}
	return *o.JsonPath
}

// GetJsonPathOk returns a tuple with the JsonPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionJSONPathTargetTarget) GetJsonPathOk() (*string, bool) {
	if o == nil || o.JsonPath == nil {
		return nil, false
	}
	return o.JsonPath, true
}

// HasJsonPath returns a boolean if a field has been set.
func (o *SyntheticsAssertionJSONPathTargetTarget) HasJsonPath() bool {
	return o != nil && o.JsonPath != nil
}

// SetJsonPath gets a reference to the given string and assigns it to the JsonPath field.
func (o *SyntheticsAssertionJSONPathTargetTarget) SetJsonPath(v string) {
	o.JsonPath = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *SyntheticsAssertionJSONPathTargetTarget) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionJSONPathTargetTarget) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *SyntheticsAssertionJSONPathTargetTarget) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *SyntheticsAssertionJSONPathTargetTarget) SetOperator(v string) {
	o.Operator = &v
}

// GetTargetValue returns the TargetValue field value if set, zero value otherwise.
func (o *SyntheticsAssertionJSONPathTargetTarget) GetTargetValue() SyntheticsAssertionTargetValue {
	if o == nil || o.TargetValue == nil {
		var ret SyntheticsAssertionTargetValue
		return ret
	}
	return *o.TargetValue
}

// GetTargetValueOk returns a tuple with the TargetValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionJSONPathTargetTarget) GetTargetValueOk() (*SyntheticsAssertionTargetValue, bool) {
	if o == nil || o.TargetValue == nil {
		return nil, false
	}
	return o.TargetValue, true
}

// HasTargetValue returns a boolean if a field has been set.
func (o *SyntheticsAssertionJSONPathTargetTarget) HasTargetValue() bool {
	return o != nil && o.TargetValue != nil
}

// SetTargetValue gets a reference to the given SyntheticsAssertionTargetValue and assigns it to the TargetValue field.
func (o *SyntheticsAssertionJSONPathTargetTarget) SetTargetValue(v SyntheticsAssertionTargetValue) {
	o.TargetValue = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsAssertionJSONPathTargetTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ElementsOperator != nil {
		toSerialize["elementsOperator"] = o.ElementsOperator
	}
	if o.JsonPath != nil {
		toSerialize["jsonPath"] = o.JsonPath
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.TargetValue != nil {
		toSerialize["targetValue"] = o.TargetValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsAssertionJSONPathTargetTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ElementsOperator *string                         `json:"elementsOperator,omitempty"`
		JsonPath         *string                         `json:"jsonPath,omitempty"`
		Operator         *string                         `json:"operator,omitempty"`
		TargetValue      *SyntheticsAssertionTargetValue `json:"targetValue,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"elementsOperator", "jsonPath", "operator", "targetValue"})
	} else {
		return err
	}
	o.ElementsOperator = all.ElementsOperator
	o.JsonPath = all.JsonPath
	o.Operator = all.Operator
	o.TargetValue = all.TargetValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
