// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RoutingRuleAttributes Defines the configurable attributes of a routing rule, such as actions, query, time restriction, and urgency.
type RoutingRuleAttributes struct {
	// Specifies the list of actions to perform when the routing rule matches.
	Actions []RoutingRuleAction `json:"actions,omitempty"`
	// Defines the query or condition that triggers this routing rule.
	Query *string `json:"query,omitempty"`
	// Holds time zone information and a list of time restrictions for a routing rule.
	TimeRestriction *TimeRestrictions `json:"time_restriction,omitempty"`
	// Specifies the level of urgency for a routing rule (low, high, or dynamic).
	Urgency *Urgency `json:"urgency,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRoutingRuleAttributes instantiates a new RoutingRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRoutingRuleAttributes() *RoutingRuleAttributes {
	this := RoutingRuleAttributes{}
	return &this
}

// NewRoutingRuleAttributesWithDefaults instantiates a new RoutingRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRoutingRuleAttributesWithDefaults() *RoutingRuleAttributes {
	this := RoutingRuleAttributes{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *RoutingRuleAttributes) GetActions() []RoutingRuleAction {
	if o == nil || o.Actions == nil {
		var ret []RoutingRuleAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleAttributes) GetActionsOk() (*[]RoutingRuleAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return &o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *RoutingRuleAttributes) HasActions() bool {
	return o != nil && o.Actions != nil
}

// SetActions gets a reference to the given []RoutingRuleAction and assigns it to the Actions field.
func (o *RoutingRuleAttributes) SetActions(v []RoutingRuleAction) {
	o.Actions = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *RoutingRuleAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *RoutingRuleAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *RoutingRuleAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetTimeRestriction returns the TimeRestriction field value if set, zero value otherwise.
func (o *RoutingRuleAttributes) GetTimeRestriction() TimeRestrictions {
	if o == nil || o.TimeRestriction == nil {
		var ret TimeRestrictions
		return ret
	}
	return *o.TimeRestriction
}

// GetTimeRestrictionOk returns a tuple with the TimeRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleAttributes) GetTimeRestrictionOk() (*TimeRestrictions, bool) {
	if o == nil || o.TimeRestriction == nil {
		return nil, false
	}
	return o.TimeRestriction, true
}

// HasTimeRestriction returns a boolean if a field has been set.
func (o *RoutingRuleAttributes) HasTimeRestriction() bool {
	return o != nil && o.TimeRestriction != nil
}

// SetTimeRestriction gets a reference to the given TimeRestrictions and assigns it to the TimeRestriction field.
func (o *RoutingRuleAttributes) SetTimeRestriction(v TimeRestrictions) {
	o.TimeRestriction = &v
}

// GetUrgency returns the Urgency field value if set, zero value otherwise.
func (o *RoutingRuleAttributes) GetUrgency() Urgency {
	if o == nil || o.Urgency == nil {
		var ret Urgency
		return ret
	}
	return *o.Urgency
}

// GetUrgencyOk returns a tuple with the Urgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleAttributes) GetUrgencyOk() (*Urgency, bool) {
	if o == nil || o.Urgency == nil {
		return nil, false
	}
	return o.Urgency, true
}

// HasUrgency returns a boolean if a field has been set.
func (o *RoutingRuleAttributes) HasUrgency() bool {
	return o != nil && o.Urgency != nil
}

// SetUrgency gets a reference to the given Urgency and assigns it to the Urgency field.
func (o *RoutingRuleAttributes) SetUrgency(v Urgency) {
	o.Urgency = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RoutingRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.TimeRestriction != nil {
		toSerialize["time_restriction"] = o.TimeRestriction
	}
	if o.Urgency != nil {
		toSerialize["urgency"] = o.Urgency
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RoutingRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Actions         []RoutingRuleAction `json:"actions,omitempty"`
		Query           *string             `json:"query,omitempty"`
		TimeRestriction *TimeRestrictions   `json:"time_restriction,omitempty"`
		Urgency         *Urgency            `json:"urgency,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actions", "query", "time_restriction", "urgency"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Actions = all.Actions
	o.Query = all.Query
	if all.TimeRestriction != nil && all.TimeRestriction.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeRestriction = all.TimeRestriction
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
