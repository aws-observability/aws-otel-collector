// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateOutcomesAsyncRequestItem Scorecard outcome for a single entity and rule.
type UpdateOutcomesAsyncRequestItem struct {
	// The unique reference for an IDP entity.
	EntityReference string `json:"entity_reference"`
	// Any remarks regarding the scorecard rule's evaluation. Supports HTML hyperlinks.
	Remarks *string `json:"remarks,omitempty"`
	// The unique ID for a scorecard rule.
	RuleId string `json:"rule_id"`
	// The state of the rule evaluation.
	State State `json:"state"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateOutcomesAsyncRequestItem instantiates a new UpdateOutcomesAsyncRequestItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateOutcomesAsyncRequestItem(entityReference string, ruleId string, state State) *UpdateOutcomesAsyncRequestItem {
	this := UpdateOutcomesAsyncRequestItem{}
	this.EntityReference = entityReference
	this.RuleId = ruleId
	this.State = state
	return &this
}

// NewUpdateOutcomesAsyncRequestItemWithDefaults instantiates a new UpdateOutcomesAsyncRequestItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateOutcomesAsyncRequestItemWithDefaults() *UpdateOutcomesAsyncRequestItem {
	this := UpdateOutcomesAsyncRequestItem{}
	return &this
}

// GetEntityReference returns the EntityReference field value.
func (o *UpdateOutcomesAsyncRequestItem) GetEntityReference() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EntityReference
}

// GetEntityReferenceOk returns a tuple with the EntityReference field value
// and a boolean to check if the value has been set.
func (o *UpdateOutcomesAsyncRequestItem) GetEntityReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityReference, true
}

// SetEntityReference sets field value.
func (o *UpdateOutcomesAsyncRequestItem) SetEntityReference(v string) {
	o.EntityReference = v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *UpdateOutcomesAsyncRequestItem) GetRemarks() string {
	if o == nil || o.Remarks == nil {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOutcomesAsyncRequestItem) GetRemarksOk() (*string, bool) {
	if o == nil || o.Remarks == nil {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *UpdateOutcomesAsyncRequestItem) HasRemarks() bool {
	return o != nil && o.Remarks != nil
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *UpdateOutcomesAsyncRequestItem) SetRemarks(v string) {
	o.Remarks = &v
}

// GetRuleId returns the RuleId field value.
func (o *UpdateOutcomesAsyncRequestItem) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *UpdateOutcomesAsyncRequestItem) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value.
func (o *UpdateOutcomesAsyncRequestItem) SetRuleId(v string) {
	o.RuleId = v
}

// GetState returns the State field value.
func (o *UpdateOutcomesAsyncRequestItem) GetState() State {
	if o == nil {
		var ret State
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *UpdateOutcomesAsyncRequestItem) GetStateOk() (*State, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *UpdateOutcomesAsyncRequestItem) SetState(v State) {
	o.State = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateOutcomesAsyncRequestItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["entity_reference"] = o.EntityReference
	if o.Remarks != nil {
		toSerialize["remarks"] = o.Remarks
	}
	toSerialize["rule_id"] = o.RuleId
	toSerialize["state"] = o.State

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateOutcomesAsyncRequestItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EntityReference *string `json:"entity_reference"`
		Remarks         *string `json:"remarks,omitempty"`
		RuleId          *string `json:"rule_id"`
		State           *State  `json:"state"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EntityReference == nil {
		return fmt.Errorf("required field entity_reference missing")
	}
	if all.RuleId == nil {
		return fmt.Errorf("required field rule_id missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"entity_reference", "remarks", "rule_id", "state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EntityReference = *all.EntityReference
	o.Remarks = all.Remarks
	o.RuleId = *all.RuleId
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
