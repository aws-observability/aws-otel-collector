// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyCreateRequestDataAttributesStepsItems Defines a single escalation step within an escalation policy creation request. Contains assignment strategy, escalation timeout, and a list of targets.
type EscalationPolicyCreateRequestDataAttributesStepsItems struct {
	// Specifies how this escalation step will assign targets (example `default` or `round-robin`).
	Assignment *EscalationPolicyStepAttributesAssignment `json:"assignment,omitempty"`
	// Defines how many seconds to wait before escalating to the next step.
	EscalateAfterSeconds *int64 `json:"escalate_after_seconds,omitempty"`
	// Specifies the collection of escalation targets for this step.
	Targets []EscalationPolicyStepTarget `json:"targets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyCreateRequestDataAttributesStepsItems instantiates a new EscalationPolicyCreateRequestDataAttributesStepsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyCreateRequestDataAttributesStepsItems(targets []EscalationPolicyStepTarget) *EscalationPolicyCreateRequestDataAttributesStepsItems {
	this := EscalationPolicyCreateRequestDataAttributesStepsItems{}
	this.Targets = targets
	return &this
}

// NewEscalationPolicyCreateRequestDataAttributesStepsItemsWithDefaults instantiates a new EscalationPolicyCreateRequestDataAttributesStepsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyCreateRequestDataAttributesStepsItemsWithDefaults() *EscalationPolicyCreateRequestDataAttributesStepsItems {
	this := EscalationPolicyCreateRequestDataAttributesStepsItems{}
	return &this
}

// GetAssignment returns the Assignment field value if set, zero value otherwise.
func (o *EscalationPolicyCreateRequestDataAttributesStepsItems) GetAssignment() EscalationPolicyStepAttributesAssignment {
	if o == nil || o.Assignment == nil {
		var ret EscalationPolicyStepAttributesAssignment
		return ret
	}
	return *o.Assignment
}

// GetAssignmentOk returns a tuple with the Assignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyCreateRequestDataAttributesStepsItems) GetAssignmentOk() (*EscalationPolicyStepAttributesAssignment, bool) {
	if o == nil || o.Assignment == nil {
		return nil, false
	}
	return o.Assignment, true
}

// HasAssignment returns a boolean if a field has been set.
func (o *EscalationPolicyCreateRequestDataAttributesStepsItems) HasAssignment() bool {
	return o != nil && o.Assignment != nil
}

// SetAssignment gets a reference to the given EscalationPolicyStepAttributesAssignment and assigns it to the Assignment field.
func (o *EscalationPolicyCreateRequestDataAttributesStepsItems) SetAssignment(v EscalationPolicyStepAttributesAssignment) {
	o.Assignment = &v
}

// GetEscalateAfterSeconds returns the EscalateAfterSeconds field value if set, zero value otherwise.
func (o *EscalationPolicyCreateRequestDataAttributesStepsItems) GetEscalateAfterSeconds() int64 {
	if o == nil || o.EscalateAfterSeconds == nil {
		var ret int64
		return ret
	}
	return *o.EscalateAfterSeconds
}

// GetEscalateAfterSecondsOk returns a tuple with the EscalateAfterSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyCreateRequestDataAttributesStepsItems) GetEscalateAfterSecondsOk() (*int64, bool) {
	if o == nil || o.EscalateAfterSeconds == nil {
		return nil, false
	}
	return o.EscalateAfterSeconds, true
}

// HasEscalateAfterSeconds returns a boolean if a field has been set.
func (o *EscalationPolicyCreateRequestDataAttributesStepsItems) HasEscalateAfterSeconds() bool {
	return o != nil && o.EscalateAfterSeconds != nil
}

// SetEscalateAfterSeconds gets a reference to the given int64 and assigns it to the EscalateAfterSeconds field.
func (o *EscalationPolicyCreateRequestDataAttributesStepsItems) SetEscalateAfterSeconds(v int64) {
	o.EscalateAfterSeconds = &v
}

// GetTargets returns the Targets field value.
func (o *EscalationPolicyCreateRequestDataAttributesStepsItems) GetTargets() []EscalationPolicyStepTarget {
	if o == nil {
		var ret []EscalationPolicyStepTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyCreateRequestDataAttributesStepsItems) GetTargetsOk() (*[]EscalationPolicyStepTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value.
func (o *EscalationPolicyCreateRequestDataAttributesStepsItems) SetTargets(v []EscalationPolicyStepTarget) {
	o.Targets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyCreateRequestDataAttributesStepsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assignment != nil {
		toSerialize["assignment"] = o.Assignment
	}
	if o.EscalateAfterSeconds != nil {
		toSerialize["escalate_after_seconds"] = o.EscalateAfterSeconds
	}
	toSerialize["targets"] = o.Targets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationPolicyCreateRequestDataAttributesStepsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assignment           *EscalationPolicyStepAttributesAssignment `json:"assignment,omitempty"`
		EscalateAfterSeconds *int64                                    `json:"escalate_after_seconds,omitempty"`
		Targets              *[]EscalationPolicyStepTarget             `json:"targets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Targets == nil {
		return fmt.Errorf("required field targets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignment", "escalate_after_seconds", "targets"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Assignment != nil && !all.Assignment.IsValid() {
		hasInvalidField = true
	} else {
		o.Assignment = all.Assignment
	}
	o.EscalateAfterSeconds = all.EscalateAfterSeconds
	o.Targets = *all.Targets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
