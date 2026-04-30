// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RolloutOptions Applied progression options for a progressive rollout.
type RolloutOptions struct {
	// Whether the schedule starts automatically.
	Autostart bool `json:"autostart"`
	// Interval in milliseconds for uniform interval strategies.
	SelectionIntervalMs int64 `json:"selection_interval_ms"`
	// The progression strategy used by a progressive rollout.
	Strategy RolloutStrategy `json:"strategy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRolloutOptions instantiates a new RolloutOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRolloutOptions(autostart bool, selectionIntervalMs int64, strategy RolloutStrategy) *RolloutOptions {
	this := RolloutOptions{}
	this.Autostart = autostart
	this.SelectionIntervalMs = selectionIntervalMs
	this.Strategy = strategy
	return &this
}

// NewRolloutOptionsWithDefaults instantiates a new RolloutOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRolloutOptionsWithDefaults() *RolloutOptions {
	this := RolloutOptions{}
	return &this
}

// GetAutostart returns the Autostart field value.
func (o *RolloutOptions) GetAutostart() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Autostart
}

// GetAutostartOk returns a tuple with the Autostart field value
// and a boolean to check if the value has been set.
func (o *RolloutOptions) GetAutostartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Autostart, true
}

// SetAutostart sets field value.
func (o *RolloutOptions) SetAutostart(v bool) {
	o.Autostart = v
}

// GetSelectionIntervalMs returns the SelectionIntervalMs field value.
func (o *RolloutOptions) GetSelectionIntervalMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SelectionIntervalMs
}

// GetSelectionIntervalMsOk returns a tuple with the SelectionIntervalMs field value
// and a boolean to check if the value has been set.
func (o *RolloutOptions) GetSelectionIntervalMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectionIntervalMs, true
}

// SetSelectionIntervalMs sets field value.
func (o *RolloutOptions) SetSelectionIntervalMs(v int64) {
	o.SelectionIntervalMs = v
}

// GetStrategy returns the Strategy field value.
func (o *RolloutOptions) GetStrategy() RolloutStrategy {
	if o == nil {
		var ret RolloutStrategy
		return ret
	}
	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *RolloutOptions) GetStrategyOk() (*RolloutStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value.
func (o *RolloutOptions) SetStrategy(v RolloutStrategy) {
	o.Strategy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RolloutOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["autostart"] = o.Autostart
	toSerialize["selection_interval_ms"] = o.SelectionIntervalMs
	toSerialize["strategy"] = o.Strategy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RolloutOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Autostart           *bool            `json:"autostart"`
		SelectionIntervalMs *int64           `json:"selection_interval_ms"`
		Strategy            *RolloutStrategy `json:"strategy"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Autostart == nil {
		return fmt.Errorf("required field autostart missing")
	}
	if all.SelectionIntervalMs == nil {
		return fmt.Errorf("required field selection_interval_ms missing")
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
	o.Autostart = *all.Autostart
	o.SelectionIntervalMs = *all.SelectionIntervalMs
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
