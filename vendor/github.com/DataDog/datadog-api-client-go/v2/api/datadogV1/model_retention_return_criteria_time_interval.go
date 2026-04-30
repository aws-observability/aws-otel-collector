// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionReturnCriteriaTimeInterval Time interval for return criteria.
type RetentionReturnCriteriaTimeInterval struct {
	// Type of time interval for return criteria.
	Type RetentionReturnCriteriaTimeIntervalType `json:"type"`
	// Unit of time for retention return criteria interval.
	Unit RetentionReturnCriteriaTimeIntervalUnit `json:"unit"`
	// Value of the time interval.
	Value float64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionReturnCriteriaTimeInterval instantiates a new RetentionReturnCriteriaTimeInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionReturnCriteriaTimeInterval(typeVar RetentionReturnCriteriaTimeIntervalType, unit RetentionReturnCriteriaTimeIntervalUnit, value float64) *RetentionReturnCriteriaTimeInterval {
	this := RetentionReturnCriteriaTimeInterval{}
	this.Type = typeVar
	this.Unit = unit
	this.Value = value
	return &this
}

// NewRetentionReturnCriteriaTimeIntervalWithDefaults instantiates a new RetentionReturnCriteriaTimeInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionReturnCriteriaTimeIntervalWithDefaults() *RetentionReturnCriteriaTimeInterval {
	this := RetentionReturnCriteriaTimeInterval{}
	return &this
}

// GetType returns the Type field value.
func (o *RetentionReturnCriteriaTimeInterval) GetType() RetentionReturnCriteriaTimeIntervalType {
	if o == nil {
		var ret RetentionReturnCriteriaTimeIntervalType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RetentionReturnCriteriaTimeInterval) GetTypeOk() (*RetentionReturnCriteriaTimeIntervalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RetentionReturnCriteriaTimeInterval) SetType(v RetentionReturnCriteriaTimeIntervalType) {
	o.Type = v
}

// GetUnit returns the Unit field value.
func (o *RetentionReturnCriteriaTimeInterval) GetUnit() RetentionReturnCriteriaTimeIntervalUnit {
	if o == nil {
		var ret RetentionReturnCriteriaTimeIntervalUnit
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *RetentionReturnCriteriaTimeInterval) GetUnitOk() (*RetentionReturnCriteriaTimeIntervalUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *RetentionReturnCriteriaTimeInterval) SetUnit(v RetentionReturnCriteriaTimeIntervalUnit) {
	o.Unit = v
}

// GetValue returns the Value field value.
func (o *RetentionReturnCriteriaTimeInterval) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RetentionReturnCriteriaTimeInterval) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *RetentionReturnCriteriaTimeInterval) SetValue(v float64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionReturnCriteriaTimeInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["unit"] = o.Unit
	toSerialize["value"] = o.Value
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionReturnCriteriaTimeInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *RetentionReturnCriteriaTimeIntervalType `json:"type"`
		Unit  *RetentionReturnCriteriaTimeIntervalUnit `json:"unit"`
		Value *float64                                 `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Unit == nil {
		return fmt.Errorf("required field unit missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if !all.Unit.IsValid() {
		hasInvalidField = true
	} else {
		o.Unit = *all.Unit
	}
	o.Value = *all.Value

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
