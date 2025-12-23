// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConfiguredScheduleTargetRelationships Represents the relationships of a configured schedule target.
type ConfiguredScheduleTargetRelationships struct {
	// Holds the schedule reference for a configured schedule target.
	Schedule ConfiguredScheduleTargetRelationshipsSchedule `json:"schedule"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConfiguredScheduleTargetRelationships instantiates a new ConfiguredScheduleTargetRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConfiguredScheduleTargetRelationships(schedule ConfiguredScheduleTargetRelationshipsSchedule) *ConfiguredScheduleTargetRelationships {
	this := ConfiguredScheduleTargetRelationships{}
	this.Schedule = schedule
	return &this
}

// NewConfiguredScheduleTargetRelationshipsWithDefaults instantiates a new ConfiguredScheduleTargetRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConfiguredScheduleTargetRelationshipsWithDefaults() *ConfiguredScheduleTargetRelationships {
	this := ConfiguredScheduleTargetRelationships{}
	return &this
}

// GetSchedule returns the Schedule field value.
func (o *ConfiguredScheduleTargetRelationships) GetSchedule() ConfiguredScheduleTargetRelationshipsSchedule {
	if o == nil {
		var ret ConfiguredScheduleTargetRelationshipsSchedule
		return ret
	}
	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *ConfiguredScheduleTargetRelationships) GetScheduleOk() (*ConfiguredScheduleTargetRelationshipsSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value.
func (o *ConfiguredScheduleTargetRelationships) SetSchedule(v ConfiguredScheduleTargetRelationshipsSchedule) {
	o.Schedule = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConfiguredScheduleTargetRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["schedule"] = o.Schedule

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConfiguredScheduleTargetRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Schedule *ConfiguredScheduleTargetRelationshipsSchedule `json:"schedule"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Schedule == nil {
		return fmt.Errorf("required field schedule missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"schedule"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Schedule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Schedule = *all.Schedule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
