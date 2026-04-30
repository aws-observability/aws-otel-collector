// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AutoCloseInactiveCases Auto-close inactive cases settings.
type AutoCloseInactiveCases struct {
	// Whether auto-close is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Maximum inactive time in seconds before auto-closing.
	MaxInactiveTimeInSecs *int64 `json:"max_inactive_time_in_secs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutoCloseInactiveCases instantiates a new AutoCloseInactiveCases object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutoCloseInactiveCases() *AutoCloseInactiveCases {
	this := AutoCloseInactiveCases{}
	return &this
}

// NewAutoCloseInactiveCasesWithDefaults instantiates a new AutoCloseInactiveCases object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutoCloseInactiveCasesWithDefaults() *AutoCloseInactiveCases {
	this := AutoCloseInactiveCases{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AutoCloseInactiveCases) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoCloseInactiveCases) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AutoCloseInactiveCases) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AutoCloseInactiveCases) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaxInactiveTimeInSecs returns the MaxInactiveTimeInSecs field value if set, zero value otherwise.
func (o *AutoCloseInactiveCases) GetMaxInactiveTimeInSecs() int64 {
	if o == nil || o.MaxInactiveTimeInSecs == nil {
		var ret int64
		return ret
	}
	return *o.MaxInactiveTimeInSecs
}

// GetMaxInactiveTimeInSecsOk returns a tuple with the MaxInactiveTimeInSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoCloseInactiveCases) GetMaxInactiveTimeInSecsOk() (*int64, bool) {
	if o == nil || o.MaxInactiveTimeInSecs == nil {
		return nil, false
	}
	return o.MaxInactiveTimeInSecs, true
}

// HasMaxInactiveTimeInSecs returns a boolean if a field has been set.
func (o *AutoCloseInactiveCases) HasMaxInactiveTimeInSecs() bool {
	return o != nil && o.MaxInactiveTimeInSecs != nil
}

// SetMaxInactiveTimeInSecs gets a reference to the given int64 and assigns it to the MaxInactiveTimeInSecs field.
func (o *AutoCloseInactiveCases) SetMaxInactiveTimeInSecs(v int64) {
	o.MaxInactiveTimeInSecs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AutoCloseInactiveCases) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.MaxInactiveTimeInSecs != nil {
		toSerialize["max_inactive_time_in_secs"] = o.MaxInactiveTimeInSecs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutoCloseInactiveCases) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled               *bool  `json:"enabled,omitempty"`
		MaxInactiveTimeInSecs *int64 `json:"max_inactive_time_in_secs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "max_inactive_time_in_secs"})
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.MaxInactiveTimeInSecs = all.MaxInactiveTimeInSecs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
