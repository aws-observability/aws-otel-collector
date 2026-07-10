// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleImpossibleTravelOptions Options on impossible travel detection method.
type SecurityMonitoringRuleImpossibleTravelOptions struct {
	// If true, signals are suppressed for the first 24 hours. In that time, Datadog learns the user's regular
	// access locations. This can be helpful to reduce noise and infer VPN usage or credentialed API access.
	BaselineUserLocations *bool `json:"baselineUserLocations,omitempty"`
	// The duration in days during which Datadog learns the user's regular access locations. After this period, signals are generated for accesses from unknown locations.
	BaselineUserLocationsDuration datadog.NullableInt32 `json:"baselineUserLocationsDuration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleImpossibleTravelOptions instantiates a new SecurityMonitoringRuleImpossibleTravelOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleImpossibleTravelOptions() *SecurityMonitoringRuleImpossibleTravelOptions {
	this := SecurityMonitoringRuleImpossibleTravelOptions{}
	return &this
}

// NewSecurityMonitoringRuleImpossibleTravelOptionsWithDefaults instantiates a new SecurityMonitoringRuleImpossibleTravelOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleImpossibleTravelOptionsWithDefaults() *SecurityMonitoringRuleImpossibleTravelOptions {
	this := SecurityMonitoringRuleImpossibleTravelOptions{}
	return &this
}

// GetBaselineUserLocations returns the BaselineUserLocations field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleImpossibleTravelOptions) GetBaselineUserLocations() bool {
	if o == nil || o.BaselineUserLocations == nil {
		var ret bool
		return ret
	}
	return *o.BaselineUserLocations
}

// GetBaselineUserLocationsOk returns a tuple with the BaselineUserLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleImpossibleTravelOptions) GetBaselineUserLocationsOk() (*bool, bool) {
	if o == nil || o.BaselineUserLocations == nil {
		return nil, false
	}
	return o.BaselineUserLocations, true
}

// HasBaselineUserLocations returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleImpossibleTravelOptions) HasBaselineUserLocations() bool {
	return o != nil && o.BaselineUserLocations != nil
}

// SetBaselineUserLocations gets a reference to the given bool and assigns it to the BaselineUserLocations field.
func (o *SecurityMonitoringRuleImpossibleTravelOptions) SetBaselineUserLocations(v bool) {
	o.BaselineUserLocations = &v
}

// GetBaselineUserLocationsDuration returns the BaselineUserLocationsDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityMonitoringRuleImpossibleTravelOptions) GetBaselineUserLocationsDuration() int32 {
	if o == nil || o.BaselineUserLocationsDuration.Get() == nil {
		var ret int32
		return ret
	}
	return *o.BaselineUserLocationsDuration.Get()
}

// GetBaselineUserLocationsDurationOk returns a tuple with the BaselineUserLocationsDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityMonitoringRuleImpossibleTravelOptions) GetBaselineUserLocationsDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaselineUserLocationsDuration.Get(), o.BaselineUserLocationsDuration.IsSet()
}

// HasBaselineUserLocationsDuration returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleImpossibleTravelOptions) HasBaselineUserLocationsDuration() bool {
	return o != nil && o.BaselineUserLocationsDuration.IsSet()
}

// SetBaselineUserLocationsDuration gets a reference to the given datadog.NullableInt32 and assigns it to the BaselineUserLocationsDuration field.
func (o *SecurityMonitoringRuleImpossibleTravelOptions) SetBaselineUserLocationsDuration(v int32) {
	o.BaselineUserLocationsDuration.Set(&v)
}

// SetBaselineUserLocationsDurationNil sets the value for BaselineUserLocationsDuration to be an explicit nil.
func (o *SecurityMonitoringRuleImpossibleTravelOptions) SetBaselineUserLocationsDurationNil() {
	o.BaselineUserLocationsDuration.Set(nil)
}

// UnsetBaselineUserLocationsDuration ensures that no value is present for BaselineUserLocationsDuration, not even an explicit nil.
func (o *SecurityMonitoringRuleImpossibleTravelOptions) UnsetBaselineUserLocationsDuration() {
	o.BaselineUserLocationsDuration.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleImpossibleTravelOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BaselineUserLocations != nil {
		toSerialize["baselineUserLocations"] = o.BaselineUserLocations
	}
	if o.BaselineUserLocationsDuration.IsSet() {
		toSerialize["baselineUserLocationsDuration"] = o.BaselineUserLocationsDuration.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleImpossibleTravelOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaselineUserLocations         *bool                 `json:"baselineUserLocations,omitempty"`
		BaselineUserLocationsDuration datadog.NullableInt32 `json:"baselineUserLocationsDuration,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"baselineUserLocations", "baselineUserLocationsDuration"})
	} else {
		return err
	}
	o.BaselineUserLocations = all.BaselineUserLocations
	o.BaselineUserLocationsDuration = all.BaselineUserLocationsDuration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
