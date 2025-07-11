// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyStepAttributes Defines attributes for an escalation policy step, such as assignment strategy and escalation timeout.
type EscalationPolicyStepAttributes struct {
	// Specifies how this escalation step will assign targets (example `default` or `round-robin`).
	Assignment *EscalationPolicyStepAttributesAssignment `json:"assignment,omitempty"`
	// Specifies how many seconds to wait before escalating to the next step.
	EscalateAfterSeconds *int64 `json:"escalate_after_seconds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyStepAttributes instantiates a new EscalationPolicyStepAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyStepAttributes() *EscalationPolicyStepAttributes {
	this := EscalationPolicyStepAttributes{}
	return &this
}

// NewEscalationPolicyStepAttributesWithDefaults instantiates a new EscalationPolicyStepAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyStepAttributesWithDefaults() *EscalationPolicyStepAttributes {
	this := EscalationPolicyStepAttributes{}
	return &this
}

// GetAssignment returns the Assignment field value if set, zero value otherwise.
func (o *EscalationPolicyStepAttributes) GetAssignment() EscalationPolicyStepAttributesAssignment {
	if o == nil || o.Assignment == nil {
		var ret EscalationPolicyStepAttributesAssignment
		return ret
	}
	return *o.Assignment
}

// GetAssignmentOk returns a tuple with the Assignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyStepAttributes) GetAssignmentOk() (*EscalationPolicyStepAttributesAssignment, bool) {
	if o == nil || o.Assignment == nil {
		return nil, false
	}
	return o.Assignment, true
}

// HasAssignment returns a boolean if a field has been set.
func (o *EscalationPolicyStepAttributes) HasAssignment() bool {
	return o != nil && o.Assignment != nil
}

// SetAssignment gets a reference to the given EscalationPolicyStepAttributesAssignment and assigns it to the Assignment field.
func (o *EscalationPolicyStepAttributes) SetAssignment(v EscalationPolicyStepAttributesAssignment) {
	o.Assignment = &v
}

// GetEscalateAfterSeconds returns the EscalateAfterSeconds field value if set, zero value otherwise.
func (o *EscalationPolicyStepAttributes) GetEscalateAfterSeconds() int64 {
	if o == nil || o.EscalateAfterSeconds == nil {
		var ret int64
		return ret
	}
	return *o.EscalateAfterSeconds
}

// GetEscalateAfterSecondsOk returns a tuple with the EscalateAfterSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyStepAttributes) GetEscalateAfterSecondsOk() (*int64, bool) {
	if o == nil || o.EscalateAfterSeconds == nil {
		return nil, false
	}
	return o.EscalateAfterSeconds, true
}

// HasEscalateAfterSeconds returns a boolean if a field has been set.
func (o *EscalationPolicyStepAttributes) HasEscalateAfterSeconds() bool {
	return o != nil && o.EscalateAfterSeconds != nil
}

// SetEscalateAfterSeconds gets a reference to the given int64 and assigns it to the EscalateAfterSeconds field.
func (o *EscalationPolicyStepAttributes) SetEscalateAfterSeconds(v int64) {
	o.EscalateAfterSeconds = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyStepAttributes) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationPolicyStepAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assignment           *EscalationPolicyStepAttributesAssignment `json:"assignment,omitempty"`
		EscalateAfterSeconds *int64                                    `json:"escalate_after_seconds,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignment", "escalate_after_seconds"})
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
