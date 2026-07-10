// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RoutingRuleEscalationPolicyActionSupportHours Support hours during which the escalation policy will be executed. Outside of these hours, the escalation policy will be on hold and triggered once the next support hours window starts. This is mutually exclusive with the top-level `time_restriction` field on the routing rule.
type RoutingRuleEscalationPolicyActionSupportHours struct {
	// The list of support hours time windows.
	Restrictions []TimeRestriction `json:"restrictions,omitempty"`
	// The time zone in which the support hours are expressed.
	TimeZone string `json:"time_zone"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRoutingRuleEscalationPolicyActionSupportHours instantiates a new RoutingRuleEscalationPolicyActionSupportHours object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRoutingRuleEscalationPolicyActionSupportHours(timeZone string) *RoutingRuleEscalationPolicyActionSupportHours {
	this := RoutingRuleEscalationPolicyActionSupportHours{}
	this.TimeZone = timeZone
	return &this
}

// NewRoutingRuleEscalationPolicyActionSupportHoursWithDefaults instantiates a new RoutingRuleEscalationPolicyActionSupportHours object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRoutingRuleEscalationPolicyActionSupportHoursWithDefaults() *RoutingRuleEscalationPolicyActionSupportHours {
	this := RoutingRuleEscalationPolicyActionSupportHours{}
	return &this
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *RoutingRuleEscalationPolicyActionSupportHours) GetRestrictions() []TimeRestriction {
	if o == nil || o.Restrictions == nil {
		var ret []TimeRestriction
		return ret
	}
	return o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleEscalationPolicyActionSupportHours) GetRestrictionsOk() (*[]TimeRestriction, bool) {
	if o == nil || o.Restrictions == nil {
		return nil, false
	}
	return &o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *RoutingRuleEscalationPolicyActionSupportHours) HasRestrictions() bool {
	return o != nil && o.Restrictions != nil
}

// SetRestrictions gets a reference to the given []TimeRestriction and assigns it to the Restrictions field.
func (o *RoutingRuleEscalationPolicyActionSupportHours) SetRestrictions(v []TimeRestriction) {
	o.Restrictions = v
}

// GetTimeZone returns the TimeZone field value.
func (o *RoutingRuleEscalationPolicyActionSupportHours) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *RoutingRuleEscalationPolicyActionSupportHours) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value.
func (o *RoutingRuleEscalationPolicyActionSupportHours) SetTimeZone(v string) {
	o.TimeZone = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RoutingRuleEscalationPolicyActionSupportHours) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Restrictions != nil {
		toSerialize["restrictions"] = o.Restrictions
	}
	toSerialize["time_zone"] = o.TimeZone

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RoutingRuleEscalationPolicyActionSupportHours) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Restrictions []TimeRestriction `json:"restrictions,omitempty"`
		TimeZone     *string           `json:"time_zone"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TimeZone == nil {
		return fmt.Errorf("required field time_zone missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"restrictions", "time_zone"})
	} else {
		return err
	}
	o.Restrictions = all.Restrictions
	o.TimeZone = *all.TimeZone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
