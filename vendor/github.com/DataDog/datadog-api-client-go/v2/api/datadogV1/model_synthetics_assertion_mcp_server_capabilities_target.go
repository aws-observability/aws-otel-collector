// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionMCPServerCapabilitiesTarget An assertion that checks that an MCP server advertises the expected capabilities.
type SyntheticsAssertionMCPServerCapabilitiesTarget struct {
	// Assertion operator to apply.
	Operator SyntheticsAssertionOperator `json:"operator"`
	// List of MCP server capabilities to assert against.
	Target []SyntheticsMCPServerCapability `json:"target"`
	// Type of the assertion.
	Type SyntheticsAssertionMCPServerCapabilitiesType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsAssertionMCPServerCapabilitiesTarget instantiates a new SyntheticsAssertionMCPServerCapabilitiesTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsAssertionMCPServerCapabilitiesTarget(operator SyntheticsAssertionOperator, target []SyntheticsMCPServerCapability, typeVar SyntheticsAssertionMCPServerCapabilitiesType) *SyntheticsAssertionMCPServerCapabilitiesTarget {
	this := SyntheticsAssertionMCPServerCapabilitiesTarget{}
	this.Operator = operator
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewSyntheticsAssertionMCPServerCapabilitiesTargetWithDefaults instantiates a new SyntheticsAssertionMCPServerCapabilitiesTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsAssertionMCPServerCapabilitiesTargetWithDefaults() *SyntheticsAssertionMCPServerCapabilitiesTarget {
	this := SyntheticsAssertionMCPServerCapabilitiesTarget{}
	return &this
}

// GetOperator returns the Operator field value.
func (o *SyntheticsAssertionMCPServerCapabilitiesTarget) GetOperator() SyntheticsAssertionOperator {
	if o == nil {
		var ret SyntheticsAssertionOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionMCPServerCapabilitiesTarget) GetOperatorOk() (*SyntheticsAssertionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *SyntheticsAssertionMCPServerCapabilitiesTarget) SetOperator(v SyntheticsAssertionOperator) {
	o.Operator = v
}

// GetTarget returns the Target field value.
func (o *SyntheticsAssertionMCPServerCapabilitiesTarget) GetTarget() []SyntheticsMCPServerCapability {
	if o == nil {
		var ret []SyntheticsMCPServerCapability
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionMCPServerCapabilitiesTarget) GetTargetOk() (*[]SyntheticsMCPServerCapability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *SyntheticsAssertionMCPServerCapabilitiesTarget) SetTarget(v []SyntheticsMCPServerCapability) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *SyntheticsAssertionMCPServerCapabilitiesTarget) GetType() SyntheticsAssertionMCPServerCapabilitiesType {
	if o == nil {
		var ret SyntheticsAssertionMCPServerCapabilitiesType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionMCPServerCapabilitiesTarget) GetTypeOk() (*SyntheticsAssertionMCPServerCapabilitiesType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsAssertionMCPServerCapabilitiesTarget) SetType(v SyntheticsAssertionMCPServerCapabilitiesType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsAssertionMCPServerCapabilitiesTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["operator"] = o.Operator
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsAssertionMCPServerCapabilitiesTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Operator *SyntheticsAssertionOperator                  `json:"operator"`
		Target   *[]SyntheticsMCPServerCapability              `json:"target"`
		Type     *SyntheticsAssertionMCPServerCapabilitiesType `json:"type"`
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
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
