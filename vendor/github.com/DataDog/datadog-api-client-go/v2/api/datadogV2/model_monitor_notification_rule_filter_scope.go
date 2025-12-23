// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorNotificationRuleFilterScope Filter monitor notifications. A monitor notification must match the scope.
type MonitorNotificationRuleFilterScope struct {
	// A scope composed of one or several key:value pairs, which can be used to filter monitor notifications on monitor and group tags.
	Scope string `json:"scope"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMonitorNotificationRuleFilterScope instantiates a new MonitorNotificationRuleFilterScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorNotificationRuleFilterScope(scope string) *MonitorNotificationRuleFilterScope {
	this := MonitorNotificationRuleFilterScope{}
	this.Scope = scope
	return &this
}

// NewMonitorNotificationRuleFilterScopeWithDefaults instantiates a new MonitorNotificationRuleFilterScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorNotificationRuleFilterScopeWithDefaults() *MonitorNotificationRuleFilterScope {
	this := MonitorNotificationRuleFilterScope{}
	return &this
}

// GetScope returns the Scope field value.
func (o *MonitorNotificationRuleFilterScope) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleFilterScope) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *MonitorNotificationRuleFilterScope) SetScope(v string) {
	o.Scope = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorNotificationRuleFilterScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["scope"] = o.Scope
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorNotificationRuleFilterScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Scope *string `json:"scope"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	o.Scope = *all.Scope

	return nil
}
