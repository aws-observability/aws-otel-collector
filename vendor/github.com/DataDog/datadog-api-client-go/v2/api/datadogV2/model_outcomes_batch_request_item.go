// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OutcomesBatchRequestItem Scorecard outcome for a specific rule, for a given service within a batched update.
type OutcomesBatchRequestItem struct {
	// Any remarks regarding the scorecard rule's evaluation, and supports HTML hyperlinks.
	Remarks *string `json:"remarks,omitempty"`
	// The unique ID for a scorecard rule.
	RuleId string `json:"rule_id"`
	// The unique name for a service in the catalog.
	ServiceName string `json:"service_name"`
	// The state of the rule evaluation.
	State State `json:"state"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOutcomesBatchRequestItem instantiates a new OutcomesBatchRequestItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOutcomesBatchRequestItem(ruleId string, serviceName string, state State) *OutcomesBatchRequestItem {
	this := OutcomesBatchRequestItem{}
	this.RuleId = ruleId
	this.ServiceName = serviceName
	this.State = state
	return &this
}

// NewOutcomesBatchRequestItemWithDefaults instantiates a new OutcomesBatchRequestItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOutcomesBatchRequestItemWithDefaults() *OutcomesBatchRequestItem {
	this := OutcomesBatchRequestItem{}
	return &this
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *OutcomesBatchRequestItem) GetRemarks() string {
	if o == nil || o.Remarks == nil {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesBatchRequestItem) GetRemarksOk() (*string, bool) {
	if o == nil || o.Remarks == nil {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *OutcomesBatchRequestItem) HasRemarks() bool {
	return o != nil && o.Remarks != nil
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *OutcomesBatchRequestItem) SetRemarks(v string) {
	o.Remarks = &v
}

// GetRuleId returns the RuleId field value.
func (o *OutcomesBatchRequestItem) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *OutcomesBatchRequestItem) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value.
func (o *OutcomesBatchRequestItem) SetRuleId(v string) {
	o.RuleId = v
}

// GetServiceName returns the ServiceName field value.
func (o *OutcomesBatchRequestItem) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *OutcomesBatchRequestItem) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value.
func (o *OutcomesBatchRequestItem) SetServiceName(v string) {
	o.ServiceName = v
}

// GetState returns the State field value.
func (o *OutcomesBatchRequestItem) GetState() State {
	if o == nil {
		var ret State
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *OutcomesBatchRequestItem) GetStateOk() (*State, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *OutcomesBatchRequestItem) SetState(v State) {
	o.State = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OutcomesBatchRequestItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Remarks != nil {
		toSerialize["remarks"] = o.Remarks
	}
	toSerialize["rule_id"] = o.RuleId
	toSerialize["service_name"] = o.ServiceName
	toSerialize["state"] = o.State

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OutcomesBatchRequestItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Remarks     *string `json:"remarks,omitempty"`
		RuleId      *string `json:"rule_id"`
		ServiceName *string `json:"service_name"`
		State       *State  `json:"state"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RuleId == nil {
		return fmt.Errorf("required field rule_id missing")
	}
	if all.ServiceName == nil {
		return fmt.Errorf("required field service_name missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"remarks", "rule_id", "service_name", "state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Remarks = all.Remarks
	o.RuleId = *all.RuleId
	o.ServiceName = *all.ServiceName
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
