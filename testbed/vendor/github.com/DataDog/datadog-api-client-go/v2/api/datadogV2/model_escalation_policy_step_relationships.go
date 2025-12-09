// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyStepRelationships Represents the relationship of an escalation policy step to its targets.
type EscalationPolicyStepRelationships struct {
	// A list of escalation targets for a step
	Targets *EscalationTargets `json:"targets,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyStepRelationships instantiates a new EscalationPolicyStepRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyStepRelationships() *EscalationPolicyStepRelationships {
	this := EscalationPolicyStepRelationships{}
	return &this
}

// NewEscalationPolicyStepRelationshipsWithDefaults instantiates a new EscalationPolicyStepRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyStepRelationshipsWithDefaults() *EscalationPolicyStepRelationships {
	this := EscalationPolicyStepRelationships{}
	return &this
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *EscalationPolicyStepRelationships) GetTargets() EscalationTargets {
	if o == nil || o.Targets == nil {
		var ret EscalationTargets
		return ret
	}
	return *o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyStepRelationships) GetTargetsOk() (*EscalationTargets, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *EscalationPolicyStepRelationships) HasTargets() bool {
	return o != nil && o.Targets != nil
}

// SetTargets gets a reference to the given EscalationTargets and assigns it to the Targets field.
func (o *EscalationPolicyStepRelationships) SetTargets(v EscalationTargets) {
	o.Targets = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyStepRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Targets != nil {
		toSerialize["targets"] = o.Targets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationPolicyStepRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Targets *EscalationTargets `json:"targets,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"targets"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Targets != nil && all.Targets.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Targets = all.Targets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
