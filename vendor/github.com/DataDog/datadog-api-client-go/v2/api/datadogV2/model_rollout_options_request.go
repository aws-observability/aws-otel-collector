// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RolloutOptionsRequest Rollout options request payload.
type RolloutOptionsRequest struct {
	// Whether the schedule should begin automatically.
	Autostart datadog.NullableBool `json:"autostart,omitempty"`
	// Interval in milliseconds for uniform interval strategies.
	SelectionIntervalMs *int64 `json:"selection_interval_ms,omitempty"`
	// The progression strategy used by a progressive rollout.
	Strategy RolloutStrategy `json:"strategy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRolloutOptionsRequest instantiates a new RolloutOptionsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRolloutOptionsRequest(strategy RolloutStrategy) *RolloutOptionsRequest {
	this := RolloutOptionsRequest{}
	this.Strategy = strategy
	return &this
}

// NewRolloutOptionsRequestWithDefaults instantiates a new RolloutOptionsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRolloutOptionsRequestWithDefaults() *RolloutOptionsRequest {
	this := RolloutOptionsRequest{}
	return &this
}

// GetAutostart returns the Autostart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RolloutOptionsRequest) GetAutostart() bool {
	if o == nil || o.Autostart.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Autostart.Get()
}

// GetAutostartOk returns a tuple with the Autostart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RolloutOptionsRequest) GetAutostartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Autostart.Get(), o.Autostart.IsSet()
}

// HasAutostart returns a boolean if a field has been set.
func (o *RolloutOptionsRequest) HasAutostart() bool {
	return o != nil && o.Autostart.IsSet()
}

// SetAutostart gets a reference to the given datadog.NullableBool and assigns it to the Autostart field.
func (o *RolloutOptionsRequest) SetAutostart(v bool) {
	o.Autostart.Set(&v)
}

// SetAutostartNil sets the value for Autostart to be an explicit nil.
func (o *RolloutOptionsRequest) SetAutostartNil() {
	o.Autostart.Set(nil)
}

// UnsetAutostart ensures that no value is present for Autostart, not even an explicit nil.
func (o *RolloutOptionsRequest) UnsetAutostart() {
	o.Autostart.Unset()
}

// GetSelectionIntervalMs returns the SelectionIntervalMs field value if set, zero value otherwise.
func (o *RolloutOptionsRequest) GetSelectionIntervalMs() int64 {
	if o == nil || o.SelectionIntervalMs == nil {
		var ret int64
		return ret
	}
	return *o.SelectionIntervalMs
}

// GetSelectionIntervalMsOk returns a tuple with the SelectionIntervalMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolloutOptionsRequest) GetSelectionIntervalMsOk() (*int64, bool) {
	if o == nil || o.SelectionIntervalMs == nil {
		return nil, false
	}
	return o.SelectionIntervalMs, true
}

// HasSelectionIntervalMs returns a boolean if a field has been set.
func (o *RolloutOptionsRequest) HasSelectionIntervalMs() bool {
	return o != nil && o.SelectionIntervalMs != nil
}

// SetSelectionIntervalMs gets a reference to the given int64 and assigns it to the SelectionIntervalMs field.
func (o *RolloutOptionsRequest) SetSelectionIntervalMs(v int64) {
	o.SelectionIntervalMs = &v
}

// GetStrategy returns the Strategy field value.
func (o *RolloutOptionsRequest) GetStrategy() RolloutStrategy {
	if o == nil {
		var ret RolloutStrategy
		return ret
	}
	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *RolloutOptionsRequest) GetStrategyOk() (*RolloutStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value.
func (o *RolloutOptionsRequest) SetStrategy(v RolloutStrategy) {
	o.Strategy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RolloutOptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Autostart.IsSet() {
		toSerialize["autostart"] = o.Autostart.Get()
	}
	if o.SelectionIntervalMs != nil {
		toSerialize["selection_interval_ms"] = o.SelectionIntervalMs
	}
	toSerialize["strategy"] = o.Strategy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RolloutOptionsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Autostart           datadog.NullableBool `json:"autostart,omitempty"`
		SelectionIntervalMs *int64               `json:"selection_interval_ms,omitempty"`
		Strategy            *RolloutStrategy     `json:"strategy"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Strategy == nil {
		return fmt.Errorf("required field strategy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"autostart", "selection_interval_ms", "strategy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Autostart = all.Autostart
	o.SelectionIntervalMs = all.SelectionIntervalMs
	if !all.Strategy.IsValid() {
		hasInvalidField = true
	} else {
		o.Strategy = *all.Strategy
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
