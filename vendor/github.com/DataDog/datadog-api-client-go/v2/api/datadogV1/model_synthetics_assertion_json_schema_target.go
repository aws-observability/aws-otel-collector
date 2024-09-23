// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionJSONSchemaTarget An assertion for the `validatesJSONSchema` operator.
type SyntheticsAssertionJSONSchemaTarget struct {
	// Assertion operator to apply.
	Operator SyntheticsAssertionJSONSchemaOperator `json:"operator"`
	// Composed target for `validatesJSONSchema` operator.
	Target *SyntheticsAssertionJSONSchemaTargetTarget `json:"target,omitempty"`
	// Type of the assertion.
	Type SyntheticsAssertionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsAssertionJSONSchemaTarget instantiates a new SyntheticsAssertionJSONSchemaTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsAssertionJSONSchemaTarget(operator SyntheticsAssertionJSONSchemaOperator, typeVar SyntheticsAssertionType) *SyntheticsAssertionJSONSchemaTarget {
	this := SyntheticsAssertionJSONSchemaTarget{}
	this.Operator = operator
	this.Type = typeVar
	return &this
}

// NewSyntheticsAssertionJSONSchemaTargetWithDefaults instantiates a new SyntheticsAssertionJSONSchemaTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsAssertionJSONSchemaTargetWithDefaults() *SyntheticsAssertionJSONSchemaTarget {
	this := SyntheticsAssertionJSONSchemaTarget{}
	return &this
}

// GetOperator returns the Operator field value.
func (o *SyntheticsAssertionJSONSchemaTarget) GetOperator() SyntheticsAssertionJSONSchemaOperator {
	if o == nil {
		var ret SyntheticsAssertionJSONSchemaOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionJSONSchemaTarget) GetOperatorOk() (*SyntheticsAssertionJSONSchemaOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *SyntheticsAssertionJSONSchemaTarget) SetOperator(v SyntheticsAssertionJSONSchemaOperator) {
	o.Operator = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SyntheticsAssertionJSONSchemaTarget) GetTarget() SyntheticsAssertionJSONSchemaTargetTarget {
	if o == nil || o.Target == nil {
		var ret SyntheticsAssertionJSONSchemaTargetTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionJSONSchemaTarget) GetTargetOk() (*SyntheticsAssertionJSONSchemaTargetTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SyntheticsAssertionJSONSchemaTarget) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given SyntheticsAssertionJSONSchemaTargetTarget and assigns it to the Target field.
func (o *SyntheticsAssertionJSONSchemaTarget) SetTarget(v SyntheticsAssertionJSONSchemaTargetTarget) {
	o.Target = &v
}

// GetType returns the Type field value.
func (o *SyntheticsAssertionJSONSchemaTarget) GetType() SyntheticsAssertionType {
	if o == nil {
		var ret SyntheticsAssertionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionJSONSchemaTarget) GetTypeOk() (*SyntheticsAssertionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsAssertionJSONSchemaTarget) SetType(v SyntheticsAssertionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsAssertionJSONSchemaTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["operator"] = o.Operator
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsAssertionJSONSchemaTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Operator *SyntheticsAssertionJSONSchemaOperator     `json:"operator"`
		Target   *SyntheticsAssertionJSONSchemaTargetTarget `json:"target,omitempty"`
		Type     *SyntheticsAssertionType                   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Operator == nil {
		return fmt.Errorf("required field operator missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"operator", "target", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = *all.Operator
	}
	if all.Target != nil && all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = all.Target
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
