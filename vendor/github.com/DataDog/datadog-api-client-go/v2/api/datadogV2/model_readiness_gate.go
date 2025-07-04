// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReadinessGate Used to merge multiple branches into a single branch.
type ReadinessGate struct {
	// The definition of `ReadinessGateThresholdType` object.
	ThresholdType ReadinessGateThresholdType `json:"thresholdType"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReadinessGate instantiates a new ReadinessGate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReadinessGate(thresholdType ReadinessGateThresholdType) *ReadinessGate {
	this := ReadinessGate{}
	this.ThresholdType = thresholdType
	return &this
}

// NewReadinessGateWithDefaults instantiates a new ReadinessGate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReadinessGateWithDefaults() *ReadinessGate {
	this := ReadinessGate{}
	return &this
}

// GetThresholdType returns the ThresholdType field value.
func (o *ReadinessGate) GetThresholdType() ReadinessGateThresholdType {
	if o == nil {
		var ret ReadinessGateThresholdType
		return ret
	}
	return o.ThresholdType
}

// GetThresholdTypeOk returns a tuple with the ThresholdType field value
// and a boolean to check if the value has been set.
func (o *ReadinessGate) GetThresholdTypeOk() (*ReadinessGateThresholdType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThresholdType, true
}

// SetThresholdType sets field value.
func (o *ReadinessGate) SetThresholdType(v ReadinessGateThresholdType) {
	o.ThresholdType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReadinessGate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["thresholdType"] = o.ThresholdType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReadinessGate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ThresholdType *ReadinessGateThresholdType `json:"thresholdType"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ThresholdType == nil {
		return fmt.Errorf("required field thresholdType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"thresholdType"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.ThresholdType.IsValid() {
		hasInvalidField = true
	} else {
		o.ThresholdType = *all.ThresholdType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
