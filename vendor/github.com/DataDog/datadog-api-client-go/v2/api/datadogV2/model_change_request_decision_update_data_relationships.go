// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestDecisionUpdateDataRelationships Relationships for updating a change request decision.
type ChangeRequestDecisionUpdateDataRelationships struct {
	// Relationship to change request decisions.
	ChangeRequestDecisions ChangeRequestDecisionsRelationship `json:"change_request_decisions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeRequestDecisionUpdateDataRelationships instantiates a new ChangeRequestDecisionUpdateDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeRequestDecisionUpdateDataRelationships(changeRequestDecisions ChangeRequestDecisionsRelationship) *ChangeRequestDecisionUpdateDataRelationships {
	this := ChangeRequestDecisionUpdateDataRelationships{}
	this.ChangeRequestDecisions = changeRequestDecisions
	return &this
}

// NewChangeRequestDecisionUpdateDataRelationshipsWithDefaults instantiates a new ChangeRequestDecisionUpdateDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeRequestDecisionUpdateDataRelationshipsWithDefaults() *ChangeRequestDecisionUpdateDataRelationships {
	this := ChangeRequestDecisionUpdateDataRelationships{}
	return &this
}

// GetChangeRequestDecisions returns the ChangeRequestDecisions field value.
func (o *ChangeRequestDecisionUpdateDataRelationships) GetChangeRequestDecisions() ChangeRequestDecisionsRelationship {
	if o == nil {
		var ret ChangeRequestDecisionsRelationship
		return ret
	}
	return o.ChangeRequestDecisions
}

// GetChangeRequestDecisionsOk returns a tuple with the ChangeRequestDecisions field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestDecisionUpdateDataRelationships) GetChangeRequestDecisionsOk() (*ChangeRequestDecisionsRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeRequestDecisions, true
}

// SetChangeRequestDecisions sets field value.
func (o *ChangeRequestDecisionUpdateDataRelationships) SetChangeRequestDecisions(v ChangeRequestDecisionsRelationship) {
	o.ChangeRequestDecisions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeRequestDecisionUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["change_request_decisions"] = o.ChangeRequestDecisions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeRequestDecisionUpdateDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChangeRequestDecisions *ChangeRequestDecisionsRelationship `json:"change_request_decisions"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ChangeRequestDecisions == nil {
		return fmt.Errorf("required field change_request_decisions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"change_request_decisions"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ChangeRequestDecisions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ChangeRequestDecisions = *all.ChangeRequestDecisions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
