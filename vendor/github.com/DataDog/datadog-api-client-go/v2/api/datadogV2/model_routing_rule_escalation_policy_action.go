// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RoutingRuleEscalationPolicyAction Triggers an escalation policy.
type RoutingRuleEscalationPolicyAction struct {
	// The number of minutes before an acknowledged page is re-triggered.
	AckTimeoutMinutes *int64 `json:"ack_timeout_minutes,omitempty"`
	// The ID of the escalation policy to route to.
	PolicyId string `json:"policy_id"`
	// Support hours during which the escalation policy will be executed. Outside of these hours, the escalation policy will be on hold and triggered once the next support hours window starts. This is mutually exclusive with the top-level `time_restriction` field on the routing rule.
	SupportHours *RoutingRuleEscalationPolicyActionSupportHours `json:"support_hours,omitempty"`
	// Indicates that the action pages an escalation policy. This action can be set once per routing rule item, and is mutually exclusive with the top-level `policy_id` field on the routing rule.
	Type RoutingRuleEscalationPolicyActionType `json:"type"`
	// Specifies the level of urgency for a routing rule (low, high, or dynamic).
	Urgency *Urgency `json:"urgency,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRoutingRuleEscalationPolicyAction instantiates a new RoutingRuleEscalationPolicyAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRoutingRuleEscalationPolicyAction(policyId string, typeVar RoutingRuleEscalationPolicyActionType) *RoutingRuleEscalationPolicyAction {
	this := RoutingRuleEscalationPolicyAction{}
	this.PolicyId = policyId
	this.Type = typeVar
	return &this
}

// NewRoutingRuleEscalationPolicyActionWithDefaults instantiates a new RoutingRuleEscalationPolicyAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRoutingRuleEscalationPolicyActionWithDefaults() *RoutingRuleEscalationPolicyAction {
	this := RoutingRuleEscalationPolicyAction{}
	var typeVar RoutingRuleEscalationPolicyActionType = ROUTINGRULEESCALATIONPOLICYACTIONTYPE_ESCALATION_POLICY
	this.Type = typeVar
	return &this
}

// GetAckTimeoutMinutes returns the AckTimeoutMinutes field value if set, zero value otherwise.
func (o *RoutingRuleEscalationPolicyAction) GetAckTimeoutMinutes() int64 {
	if o == nil || o.AckTimeoutMinutes == nil {
		var ret int64
		return ret
	}
	return *o.AckTimeoutMinutes
}

// GetAckTimeoutMinutesOk returns a tuple with the AckTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleEscalationPolicyAction) GetAckTimeoutMinutesOk() (*int64, bool) {
	if o == nil || o.AckTimeoutMinutes == nil {
		return nil, false
	}
	return o.AckTimeoutMinutes, true
}

// HasAckTimeoutMinutes returns a boolean if a field has been set.
func (o *RoutingRuleEscalationPolicyAction) HasAckTimeoutMinutes() bool {
	return o != nil && o.AckTimeoutMinutes != nil
}

// SetAckTimeoutMinutes gets a reference to the given int64 and assigns it to the AckTimeoutMinutes field.
func (o *RoutingRuleEscalationPolicyAction) SetAckTimeoutMinutes(v int64) {
	o.AckTimeoutMinutes = &v
}

// GetPolicyId returns the PolicyId field value.
func (o *RoutingRuleEscalationPolicyAction) GetPolicyId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value
// and a boolean to check if the value has been set.
func (o *RoutingRuleEscalationPolicyAction) GetPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyId, true
}

// SetPolicyId sets field value.
func (o *RoutingRuleEscalationPolicyAction) SetPolicyId(v string) {
	o.PolicyId = v
}

// GetSupportHours returns the SupportHours field value if set, zero value otherwise.
func (o *RoutingRuleEscalationPolicyAction) GetSupportHours() RoutingRuleEscalationPolicyActionSupportHours {
	if o == nil || o.SupportHours == nil {
		var ret RoutingRuleEscalationPolicyActionSupportHours
		return ret
	}
	return *o.SupportHours
}

// GetSupportHoursOk returns a tuple with the SupportHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleEscalationPolicyAction) GetSupportHoursOk() (*RoutingRuleEscalationPolicyActionSupportHours, bool) {
	if o == nil || o.SupportHours == nil {
		return nil, false
	}
	return o.SupportHours, true
}

// HasSupportHours returns a boolean if a field has been set.
func (o *RoutingRuleEscalationPolicyAction) HasSupportHours() bool {
	return o != nil && o.SupportHours != nil
}

// SetSupportHours gets a reference to the given RoutingRuleEscalationPolicyActionSupportHours and assigns it to the SupportHours field.
func (o *RoutingRuleEscalationPolicyAction) SetSupportHours(v RoutingRuleEscalationPolicyActionSupportHours) {
	o.SupportHours = &v
}

// GetType returns the Type field value.
func (o *RoutingRuleEscalationPolicyAction) GetType() RoutingRuleEscalationPolicyActionType {
	if o == nil {
		var ret RoutingRuleEscalationPolicyActionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoutingRuleEscalationPolicyAction) GetTypeOk() (*RoutingRuleEscalationPolicyActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RoutingRuleEscalationPolicyAction) SetType(v RoutingRuleEscalationPolicyActionType) {
	o.Type = v
}

// GetUrgency returns the Urgency field value if set, zero value otherwise.
func (o *RoutingRuleEscalationPolicyAction) GetUrgency() Urgency {
	if o == nil || o.Urgency == nil {
		var ret Urgency
		return ret
	}
	return *o.Urgency
}

// GetUrgencyOk returns a tuple with the Urgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleEscalationPolicyAction) GetUrgencyOk() (*Urgency, bool) {
	if o == nil || o.Urgency == nil {
		return nil, false
	}
	return o.Urgency, true
}

// HasUrgency returns a boolean if a field has been set.
func (o *RoutingRuleEscalationPolicyAction) HasUrgency() bool {
	return o != nil && o.Urgency != nil
}

// SetUrgency gets a reference to the given Urgency and assigns it to the Urgency field.
func (o *RoutingRuleEscalationPolicyAction) SetUrgency(v Urgency) {
	o.Urgency = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RoutingRuleEscalationPolicyAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AckTimeoutMinutes != nil {
		toSerialize["ack_timeout_minutes"] = o.AckTimeoutMinutes
	}
	toSerialize["policy_id"] = o.PolicyId
	if o.SupportHours != nil {
		toSerialize["support_hours"] = o.SupportHours
	}
	toSerialize["type"] = o.Type
	if o.Urgency != nil {
		toSerialize["urgency"] = o.Urgency
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RoutingRuleEscalationPolicyAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AckTimeoutMinutes *int64                                         `json:"ack_timeout_minutes,omitempty"`
		PolicyId          *string                                        `json:"policy_id"`
		SupportHours      *RoutingRuleEscalationPolicyActionSupportHours `json:"support_hours,omitempty"`
		Type              *RoutingRuleEscalationPolicyActionType         `json:"type"`
		Urgency           *Urgency                                       `json:"urgency,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PolicyId == nil {
		return fmt.Errorf("required field policy_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ack_timeout_minutes", "policy_id", "support_hours", "type", "urgency"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AckTimeoutMinutes = all.AckTimeoutMinutes
	o.PolicyId = *all.PolicyId
	if all.SupportHours != nil && all.SupportHours.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SupportHours = all.SupportHours
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if all.Urgency != nil && !all.Urgency.IsValid() {
		hasInvalidField = true
	} else {
		o.Urgency = all.Urgency
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
