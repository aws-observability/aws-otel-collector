// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkAssertionPacketLossPercentage Packet loss percentage assertion for a Network Path test.
type SyntheticsNetworkAssertionPacketLossPercentage struct {
	// Assertion operator to apply.
	Operator SyntheticsNetworkAssertionOperator `json:"operator"`
	// Target value as a percentage (0 to 1).
	Target float64 `json:"target"`
	// Type of the packet loss percentage assertion.
	Type SyntheticsNetworkAssertionPacketLossPercentageType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsNetworkAssertionPacketLossPercentage instantiates a new SyntheticsNetworkAssertionPacketLossPercentage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsNetworkAssertionPacketLossPercentage(operator SyntheticsNetworkAssertionOperator, target float64, typeVar SyntheticsNetworkAssertionPacketLossPercentageType) *SyntheticsNetworkAssertionPacketLossPercentage {
	this := SyntheticsNetworkAssertionPacketLossPercentage{}
	this.Operator = operator
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewSyntheticsNetworkAssertionPacketLossPercentageWithDefaults instantiates a new SyntheticsNetworkAssertionPacketLossPercentage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsNetworkAssertionPacketLossPercentageWithDefaults() *SyntheticsNetworkAssertionPacketLossPercentage {
	this := SyntheticsNetworkAssertionPacketLossPercentage{}
	var typeVar SyntheticsNetworkAssertionPacketLossPercentageType = SYNTHETICSNETWORKASSERTIONPACKETLOSSPERCENTAGETYPE_PACKET_LOSS_PERCENTAGE
	this.Type = typeVar
	return &this
}

// GetOperator returns the Operator field value.
func (o *SyntheticsNetworkAssertionPacketLossPercentage) GetOperator() SyntheticsNetworkAssertionOperator {
	if o == nil {
		var ret SyntheticsNetworkAssertionOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkAssertionPacketLossPercentage) GetOperatorOk() (*SyntheticsNetworkAssertionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *SyntheticsNetworkAssertionPacketLossPercentage) SetOperator(v SyntheticsNetworkAssertionOperator) {
	o.Operator = v
}

// GetTarget returns the Target field value.
func (o *SyntheticsNetworkAssertionPacketLossPercentage) GetTarget() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkAssertionPacketLossPercentage) GetTargetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *SyntheticsNetworkAssertionPacketLossPercentage) SetTarget(v float64) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *SyntheticsNetworkAssertionPacketLossPercentage) GetType() SyntheticsNetworkAssertionPacketLossPercentageType {
	if o == nil {
		var ret SyntheticsNetworkAssertionPacketLossPercentageType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkAssertionPacketLossPercentage) GetTypeOk() (*SyntheticsNetworkAssertionPacketLossPercentageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsNetworkAssertionPacketLossPercentage) SetType(v SyntheticsNetworkAssertionPacketLossPercentageType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsNetworkAssertionPacketLossPercentage) MarshalJSON() ([]byte, error) {
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
func (o *SyntheticsNetworkAssertionPacketLossPercentage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Operator *SyntheticsNetworkAssertionOperator                 `json:"operator"`
		Target   *float64                                            `json:"target"`
		Type     *SyntheticsNetworkAssertionPacketLossPercentageType `json:"type"`
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
