// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkAssertionJitter Jitter assertion for a Network Path test.
type SyntheticsNetworkAssertionJitter struct {
	// Assertion operator to apply.
	Operator SyntheticsNetworkAssertionOperator `json:"operator"`
	// Target value in milliseconds.
	Target float64 `json:"target"`
	// Type of the jitter assertion.
	Type SyntheticsNetworkAssertionJitterType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsNetworkAssertionJitter instantiates a new SyntheticsNetworkAssertionJitter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsNetworkAssertionJitter(operator SyntheticsNetworkAssertionOperator, target float64, typeVar SyntheticsNetworkAssertionJitterType) *SyntheticsNetworkAssertionJitter {
	this := SyntheticsNetworkAssertionJitter{}
	this.Operator = operator
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewSyntheticsNetworkAssertionJitterWithDefaults instantiates a new SyntheticsNetworkAssertionJitter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsNetworkAssertionJitterWithDefaults() *SyntheticsNetworkAssertionJitter {
	this := SyntheticsNetworkAssertionJitter{}
	var typeVar SyntheticsNetworkAssertionJitterType = SYNTHETICSNETWORKASSERTIONJITTERTYPE_JITTER
	this.Type = typeVar
	return &this
}

// GetOperator returns the Operator field value.
func (o *SyntheticsNetworkAssertionJitter) GetOperator() SyntheticsNetworkAssertionOperator {
	if o == nil {
		var ret SyntheticsNetworkAssertionOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkAssertionJitter) GetOperatorOk() (*SyntheticsNetworkAssertionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *SyntheticsNetworkAssertionJitter) SetOperator(v SyntheticsNetworkAssertionOperator) {
	o.Operator = v
}

// GetTarget returns the Target field value.
func (o *SyntheticsNetworkAssertionJitter) GetTarget() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkAssertionJitter) GetTargetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *SyntheticsNetworkAssertionJitter) SetTarget(v float64) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *SyntheticsNetworkAssertionJitter) GetType() SyntheticsNetworkAssertionJitterType {
	if o == nil {
		var ret SyntheticsNetworkAssertionJitterType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkAssertionJitter) GetTypeOk() (*SyntheticsNetworkAssertionJitterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsNetworkAssertionJitter) SetType(v SyntheticsNetworkAssertionJitterType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsNetworkAssertionJitter) MarshalJSON() ([]byte, error) {
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
func (o *SyntheticsNetworkAssertionJitter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Operator *SyntheticsNetworkAssertionOperator   `json:"operator"`
		Target   *float64                              `json:"target"`
		Type     *SyntheticsNetworkAssertionJitterType `json:"type"`
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
