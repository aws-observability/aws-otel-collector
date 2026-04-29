// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionCohortCriteriaTimeInterval Time interval for cohort criteria.
type RetentionCohortCriteriaTimeInterval struct {
	// Type of time interval for cohort criteria.
	Type RetentionCohortCriteriaTimeIntervalType `json:"type"`
	// Calendar interval definition.
	Value CalendarInterval `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionCohortCriteriaTimeInterval instantiates a new RetentionCohortCriteriaTimeInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionCohortCriteriaTimeInterval(typeVar RetentionCohortCriteriaTimeIntervalType, value CalendarInterval) *RetentionCohortCriteriaTimeInterval {
	this := RetentionCohortCriteriaTimeInterval{}
	this.Type = typeVar
	this.Value = value
	return &this
}

// NewRetentionCohortCriteriaTimeIntervalWithDefaults instantiates a new RetentionCohortCriteriaTimeInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionCohortCriteriaTimeIntervalWithDefaults() *RetentionCohortCriteriaTimeInterval {
	this := RetentionCohortCriteriaTimeInterval{}
	return &this
}

// GetType returns the Type field value.
func (o *RetentionCohortCriteriaTimeInterval) GetType() RetentionCohortCriteriaTimeIntervalType {
	if o == nil {
		var ret RetentionCohortCriteriaTimeIntervalType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RetentionCohortCriteriaTimeInterval) GetTypeOk() (*RetentionCohortCriteriaTimeIntervalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RetentionCohortCriteriaTimeInterval) SetType(v RetentionCohortCriteriaTimeIntervalType) {
	o.Type = v
}

// GetValue returns the Value field value.
func (o *RetentionCohortCriteriaTimeInterval) GetValue() CalendarInterval {
	if o == nil {
		var ret CalendarInterval
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RetentionCohortCriteriaTimeInterval) GetValueOk() (*CalendarInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *RetentionCohortCriteriaTimeInterval) SetValue(v CalendarInterval) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionCohortCriteriaTimeInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionCohortCriteriaTimeInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *RetentionCohortCriteriaTimeIntervalType `json:"type"`
		Value *CalendarInterval                        `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
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
	if all.Value.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Value = *all.Value

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
