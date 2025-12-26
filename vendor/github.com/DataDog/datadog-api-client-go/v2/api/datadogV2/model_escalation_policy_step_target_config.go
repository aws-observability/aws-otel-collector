// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyStepTargetConfig Configuration for an escalation target, such as schedule position.
type EscalationPolicyStepTargetConfig struct {
	// Schedule-specific configuration for an escalation target.
	Schedule *EscalationPolicyStepTargetConfigSchedule `json:"schedule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyStepTargetConfig instantiates a new EscalationPolicyStepTargetConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyStepTargetConfig() *EscalationPolicyStepTargetConfig {
	this := EscalationPolicyStepTargetConfig{}
	return &this
}

// NewEscalationPolicyStepTargetConfigWithDefaults instantiates a new EscalationPolicyStepTargetConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyStepTargetConfigWithDefaults() *EscalationPolicyStepTargetConfig {
	this := EscalationPolicyStepTargetConfig{}
	return &this
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *EscalationPolicyStepTargetConfig) GetSchedule() EscalationPolicyStepTargetConfigSchedule {
	if o == nil || o.Schedule == nil {
		var ret EscalationPolicyStepTargetConfigSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyStepTargetConfig) GetScheduleOk() (*EscalationPolicyStepTargetConfigSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *EscalationPolicyStepTargetConfig) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given EscalationPolicyStepTargetConfigSchedule and assigns it to the Schedule field.
func (o *EscalationPolicyStepTargetConfig) SetSchedule(v EscalationPolicyStepTargetConfigSchedule) {
	o.Schedule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyStepTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationPolicyStepTargetConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Schedule *EscalationPolicyStepTargetConfigSchedule `json:"schedule,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"schedule"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Schedule != nil && all.Schedule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Schedule = all.Schedule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
