// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionTarget An assertion which uses a simple target.
type SyntheticsAssertionTarget struct {
	// Assertion operator to apply.
	Operator SyntheticsAssertionOperator `json:"operator"`
	// The associated assertion property.
	Property *string `json:"property,omitempty"`
	// Value used by the operator in assertions. Can be either a number or string.
	Target SyntheticsAssertionTargetValue `json:"target"`
	// Timings scope for response time assertions.
	TimingsScope *SyntheticsAssertionTimingsScope `json:"timingsScope,omitempty"`
	// Type of the assertion.
	Type SyntheticsAssertionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsAssertionTarget instantiates a new SyntheticsAssertionTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsAssertionTarget(operator SyntheticsAssertionOperator, target SyntheticsAssertionTargetValue, typeVar SyntheticsAssertionType) *SyntheticsAssertionTarget {
	this := SyntheticsAssertionTarget{}
	this.Operator = operator
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewSyntheticsAssertionTargetWithDefaults instantiates a new SyntheticsAssertionTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsAssertionTargetWithDefaults() *SyntheticsAssertionTarget {
	this := SyntheticsAssertionTarget{}
	return &this
}

// GetOperator returns the Operator field value.
func (o *SyntheticsAssertionTarget) GetOperator() SyntheticsAssertionOperator {
	if o == nil {
		var ret SyntheticsAssertionOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionTarget) GetOperatorOk() (*SyntheticsAssertionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *SyntheticsAssertionTarget) SetOperator(v SyntheticsAssertionOperator) {
	o.Operator = v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *SyntheticsAssertionTarget) GetProperty() string {
	if o == nil || o.Property == nil {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionTarget) GetPropertyOk() (*string, bool) {
	if o == nil || o.Property == nil {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *SyntheticsAssertionTarget) HasProperty() bool {
	return o != nil && o.Property != nil
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *SyntheticsAssertionTarget) SetProperty(v string) {
	o.Property = &v
}

// GetTarget returns the Target field value.
func (o *SyntheticsAssertionTarget) GetTarget() SyntheticsAssertionTargetValue {
	if o == nil {
		var ret SyntheticsAssertionTargetValue
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionTarget) GetTargetOk() (*SyntheticsAssertionTargetValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *SyntheticsAssertionTarget) SetTarget(v SyntheticsAssertionTargetValue) {
	o.Target = v
}

// GetTimingsScope returns the TimingsScope field value if set, zero value otherwise.
func (o *SyntheticsAssertionTarget) GetTimingsScope() SyntheticsAssertionTimingsScope {
	if o == nil || o.TimingsScope == nil {
		var ret SyntheticsAssertionTimingsScope
		return ret
	}
	return *o.TimingsScope
}

// GetTimingsScopeOk returns a tuple with the TimingsScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionTarget) GetTimingsScopeOk() (*SyntheticsAssertionTimingsScope, bool) {
	if o == nil || o.TimingsScope == nil {
		return nil, false
	}
	return o.TimingsScope, true
}

// HasTimingsScope returns a boolean if a field has been set.
func (o *SyntheticsAssertionTarget) HasTimingsScope() bool {
	return o != nil && o.TimingsScope != nil
}

// SetTimingsScope gets a reference to the given SyntheticsAssertionTimingsScope and assigns it to the TimingsScope field.
func (o *SyntheticsAssertionTarget) SetTimingsScope(v SyntheticsAssertionTimingsScope) {
	o.TimingsScope = &v
}

// GetType returns the Type field value.
func (o *SyntheticsAssertionTarget) GetType() SyntheticsAssertionType {
	if o == nil {
		var ret SyntheticsAssertionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionTarget) GetTypeOk() (*SyntheticsAssertionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsAssertionTarget) SetType(v SyntheticsAssertionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsAssertionTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["operator"] = o.Operator
	if o.Property != nil {
		toSerialize["property"] = o.Property
	}
	toSerialize["target"] = o.Target
	if o.TimingsScope != nil {
		toSerialize["timingsScope"] = o.TimingsScope
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsAssertionTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Operator     *SyntheticsAssertionOperator     `json:"operator"`
		Property     *string                          `json:"property,omitempty"`
		Target       *SyntheticsAssertionTargetValue  `json:"target"`
		TimingsScope *SyntheticsAssertionTimingsScope `json:"timingsScope,omitempty"`
		Type         *SyntheticsAssertionType         `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Operator == nil {
		return fmt.Errorf("required field operator missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"operator", "property", "target", "timingsScope", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = *all.Operator
	}
	o.Property = all.Property
	o.Target = *all.Target
	if all.TimingsScope != nil && !all.TimingsScope.IsValid() {
		hasInvalidField = true
	} else {
		o.TimingsScope = all.TimingsScope
	}
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
