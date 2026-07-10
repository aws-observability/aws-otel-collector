// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRateLimitAdaptiveConfig The configuration used when `mode` is `adaptive`.
type RumRateLimitAdaptiveConfig struct {
	// The maximum fraction of sessions to retain, in the range `(0, 1]`.
	MaxRetentionRate float64 `json:"max_retention_rate"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRateLimitAdaptiveConfig instantiates a new RumRateLimitAdaptiveConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRateLimitAdaptiveConfig(maxRetentionRate float64) *RumRateLimitAdaptiveConfig {
	this := RumRateLimitAdaptiveConfig{}
	this.MaxRetentionRate = maxRetentionRate
	return &this
}

// NewRumRateLimitAdaptiveConfigWithDefaults instantiates a new RumRateLimitAdaptiveConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRateLimitAdaptiveConfigWithDefaults() *RumRateLimitAdaptiveConfig {
	this := RumRateLimitAdaptiveConfig{}
	return &this
}

// GetMaxRetentionRate returns the MaxRetentionRate field value.
func (o *RumRateLimitAdaptiveConfig) GetMaxRetentionRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MaxRetentionRate
}

// GetMaxRetentionRateOk returns a tuple with the MaxRetentionRate field value
// and a boolean to check if the value has been set.
func (o *RumRateLimitAdaptiveConfig) GetMaxRetentionRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxRetentionRate, true
}

// SetMaxRetentionRate sets field value.
func (o *RumRateLimitAdaptiveConfig) SetMaxRetentionRate(v float64) {
	o.MaxRetentionRate = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRateLimitAdaptiveConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["max_retention_rate"] = o.MaxRetentionRate

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumRateLimitAdaptiveConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MaxRetentionRate *float64 `json:"max_retention_rate"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MaxRetentionRate == nil {
		return fmt.Errorf("required field max_retention_rate missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max_retention_rate"})
	} else {
		return err
	}
	o.MaxRetentionRate = *all.MaxRetentionRate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
