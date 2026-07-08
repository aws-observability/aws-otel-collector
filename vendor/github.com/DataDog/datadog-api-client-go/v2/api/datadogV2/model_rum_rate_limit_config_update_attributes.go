// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRateLimitConfigUpdateAttributes The RUM rate limit configuration properties to create or update.
type RumRateLimitConfigUpdateAttributes struct {
	// The configuration used when `mode` is `adaptive`.
	Adaptive *RumRateLimitAdaptiveConfig `json:"adaptive,omitempty"`
	// The configuration used when `mode` is `custom`.
	Custom *RumRateLimitCustomConfig `json:"custom,omitempty"`
	// The rate limit mode. `custom` enforces a fixed session limit, while
	// `adaptive` dynamically adjusts retention.
	Mode RumRateLimitMode `json:"mode"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRateLimitConfigUpdateAttributes instantiates a new RumRateLimitConfigUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRateLimitConfigUpdateAttributes(mode RumRateLimitMode) *RumRateLimitConfigUpdateAttributes {
	this := RumRateLimitConfigUpdateAttributes{}
	this.Mode = mode
	return &this
}

// NewRumRateLimitConfigUpdateAttributesWithDefaults instantiates a new RumRateLimitConfigUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRateLimitConfigUpdateAttributesWithDefaults() *RumRateLimitConfigUpdateAttributes {
	this := RumRateLimitConfigUpdateAttributes{}
	return &this
}

// GetAdaptive returns the Adaptive field value if set, zero value otherwise.
func (o *RumRateLimitConfigUpdateAttributes) GetAdaptive() RumRateLimitAdaptiveConfig {
	if o == nil || o.Adaptive == nil {
		var ret RumRateLimitAdaptiveConfig
		return ret
	}
	return *o.Adaptive
}

// GetAdaptiveOk returns a tuple with the Adaptive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRateLimitConfigUpdateAttributes) GetAdaptiveOk() (*RumRateLimitAdaptiveConfig, bool) {
	if o == nil || o.Adaptive == nil {
		return nil, false
	}
	return o.Adaptive, true
}

// HasAdaptive returns a boolean if a field has been set.
func (o *RumRateLimitConfigUpdateAttributes) HasAdaptive() bool {
	return o != nil && o.Adaptive != nil
}

// SetAdaptive gets a reference to the given RumRateLimitAdaptiveConfig and assigns it to the Adaptive field.
func (o *RumRateLimitConfigUpdateAttributes) SetAdaptive(v RumRateLimitAdaptiveConfig) {
	o.Adaptive = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *RumRateLimitConfigUpdateAttributes) GetCustom() RumRateLimitCustomConfig {
	if o == nil || o.Custom == nil {
		var ret RumRateLimitCustomConfig
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRateLimitConfigUpdateAttributes) GetCustomOk() (*RumRateLimitCustomConfig, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *RumRateLimitConfigUpdateAttributes) HasCustom() bool {
	return o != nil && o.Custom != nil
}

// SetCustom gets a reference to the given RumRateLimitCustomConfig and assigns it to the Custom field.
func (o *RumRateLimitConfigUpdateAttributes) SetCustom(v RumRateLimitCustomConfig) {
	o.Custom = &v
}

// GetMode returns the Mode field value.
func (o *RumRateLimitConfigUpdateAttributes) GetMode() RumRateLimitMode {
	if o == nil {
		var ret RumRateLimitMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *RumRateLimitConfigUpdateAttributes) GetModeOk() (*RumRateLimitMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *RumRateLimitConfigUpdateAttributes) SetMode(v RumRateLimitMode) {
	o.Mode = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRateLimitConfigUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Adaptive != nil {
		toSerialize["adaptive"] = o.Adaptive
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	toSerialize["mode"] = o.Mode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumRateLimitConfigUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Adaptive *RumRateLimitAdaptiveConfig `json:"adaptive,omitempty"`
		Custom   *RumRateLimitCustomConfig   `json:"custom,omitempty"`
		Mode     *RumRateLimitMode           `json:"mode"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"adaptive", "custom", "mode"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Adaptive != nil && all.Adaptive.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Adaptive = all.Adaptive
	if all.Custom != nil && all.Custom.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Custom = all.Custom
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
