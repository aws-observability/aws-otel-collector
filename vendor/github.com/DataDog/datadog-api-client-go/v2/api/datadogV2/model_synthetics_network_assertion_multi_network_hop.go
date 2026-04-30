// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkAssertionMultiNetworkHop Multi-network hop assertion for a Network Path test.
type SyntheticsNetworkAssertionMultiNetworkHop struct {
	// Assertion operator to apply.
	Operator SyntheticsNetworkAssertionOperator `json:"operator"`
	// The associated assertion property.
	Property SyntheticsNetworkAssertionProperty `json:"property"`
	// Target value in number of hops.
	Target float64 `json:"target"`
	// Type of the multi-network hop assertion.
	Type SyntheticsNetworkAssertionMultiNetworkHopType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsNetworkAssertionMultiNetworkHop instantiates a new SyntheticsNetworkAssertionMultiNetworkHop object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsNetworkAssertionMultiNetworkHop(operator SyntheticsNetworkAssertionOperator, property SyntheticsNetworkAssertionProperty, target float64, typeVar SyntheticsNetworkAssertionMultiNetworkHopType) *SyntheticsNetworkAssertionMultiNetworkHop {
	this := SyntheticsNetworkAssertionMultiNetworkHop{}
	this.Operator = operator
	this.Property = property
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewSyntheticsNetworkAssertionMultiNetworkHopWithDefaults instantiates a new SyntheticsNetworkAssertionMultiNetworkHop object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsNetworkAssertionMultiNetworkHopWithDefaults() *SyntheticsNetworkAssertionMultiNetworkHop {
	this := SyntheticsNetworkAssertionMultiNetworkHop{}
	var typeVar SyntheticsNetworkAssertionMultiNetworkHopType = SYNTHETICSNETWORKASSERTIONMULTINETWORKHOPTYPE_MULTI_NETWORK_HOP
	this.Type = typeVar
	return &this
}

// GetOperator returns the Operator field value.
func (o *SyntheticsNetworkAssertionMultiNetworkHop) GetOperator() SyntheticsNetworkAssertionOperator {
	if o == nil {
		var ret SyntheticsNetworkAssertionOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkAssertionMultiNetworkHop) GetOperatorOk() (*SyntheticsNetworkAssertionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *SyntheticsNetworkAssertionMultiNetworkHop) SetOperator(v SyntheticsNetworkAssertionOperator) {
	o.Operator = v
}

// GetProperty returns the Property field value.
func (o *SyntheticsNetworkAssertionMultiNetworkHop) GetProperty() SyntheticsNetworkAssertionProperty {
	if o == nil {
		var ret SyntheticsNetworkAssertionProperty
		return ret
	}
	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkAssertionMultiNetworkHop) GetPropertyOk() (*SyntheticsNetworkAssertionProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value.
func (o *SyntheticsNetworkAssertionMultiNetworkHop) SetProperty(v SyntheticsNetworkAssertionProperty) {
	o.Property = v
}

// GetTarget returns the Target field value.
func (o *SyntheticsNetworkAssertionMultiNetworkHop) GetTarget() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkAssertionMultiNetworkHop) GetTargetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *SyntheticsNetworkAssertionMultiNetworkHop) SetTarget(v float64) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *SyntheticsNetworkAssertionMultiNetworkHop) GetType() SyntheticsNetworkAssertionMultiNetworkHopType {
	if o == nil {
		var ret SyntheticsNetworkAssertionMultiNetworkHopType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkAssertionMultiNetworkHop) GetTypeOk() (*SyntheticsNetworkAssertionMultiNetworkHopType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsNetworkAssertionMultiNetworkHop) SetType(v SyntheticsNetworkAssertionMultiNetworkHopType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsNetworkAssertionMultiNetworkHop) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["operator"] = o.Operator
	toSerialize["property"] = o.Property
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsNetworkAssertionMultiNetworkHop) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Operator *SyntheticsNetworkAssertionOperator            `json:"operator"`
		Property *SyntheticsNetworkAssertionProperty            `json:"property"`
		Target   *float64                                       `json:"target"`
		Type     *SyntheticsNetworkAssertionMultiNetworkHopType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Operator == nil {
		return fmt.Errorf("required field operator missing")
	}
	if all.Property == nil {
		return fmt.Errorf("required field property missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"operator", "property", "target", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = *all.Operator
	}
	if !all.Property.IsValid() {
		hasInvalidField = true
	} else {
		o.Property = *all.Property
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
