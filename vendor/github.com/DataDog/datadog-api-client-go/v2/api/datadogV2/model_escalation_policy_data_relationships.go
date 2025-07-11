// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyDataRelationships Represents the relationships for an escalation policy, including references to steps and teams.
type EscalationPolicyDataRelationships struct {
	// Defines the relationship to a collection of steps within an escalation policy. Contains an array of step data references.
	Steps EscalationPolicyDataRelationshipsSteps `json:"steps"`
	// Associates teams with this schedule in a data structure.
	Teams *DataRelationshipsTeams `json:"teams,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyDataRelationships instantiates a new EscalationPolicyDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyDataRelationships(steps EscalationPolicyDataRelationshipsSteps) *EscalationPolicyDataRelationships {
	this := EscalationPolicyDataRelationships{}
	this.Steps = steps
	return &this
}

// NewEscalationPolicyDataRelationshipsWithDefaults instantiates a new EscalationPolicyDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyDataRelationshipsWithDefaults() *EscalationPolicyDataRelationships {
	this := EscalationPolicyDataRelationships{}
	return &this
}

// GetSteps returns the Steps field value.
func (o *EscalationPolicyDataRelationships) GetSteps() EscalationPolicyDataRelationshipsSteps {
	if o == nil {
		var ret EscalationPolicyDataRelationshipsSteps
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyDataRelationships) GetStepsOk() (*EscalationPolicyDataRelationshipsSteps, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Steps, true
}

// SetSteps sets field value.
func (o *EscalationPolicyDataRelationships) SetSteps(v EscalationPolicyDataRelationshipsSteps) {
	o.Steps = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *EscalationPolicyDataRelationships) GetTeams() DataRelationshipsTeams {
	if o == nil || o.Teams == nil {
		var ret DataRelationshipsTeams
		return ret
	}
	return *o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyDataRelationships) GetTeamsOk() (*DataRelationshipsTeams, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *EscalationPolicyDataRelationships) HasTeams() bool {
	return o != nil && o.Teams != nil
}

// SetTeams gets a reference to the given DataRelationshipsTeams and assigns it to the Teams field.
func (o *EscalationPolicyDataRelationships) SetTeams(v DataRelationshipsTeams) {
	o.Teams = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["steps"] = o.Steps
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationPolicyDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Steps *EscalationPolicyDataRelationshipsSteps `json:"steps"`
		Teams *DataRelationshipsTeams                 `json:"teams,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Steps == nil {
		return fmt.Errorf("required field steps missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"steps", "teams"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Steps.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Steps = *all.Steps
	if all.Teams != nil && all.Teams.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Teams = all.Teams

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
