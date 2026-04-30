// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule Automatic quarantine triggering rule based on a time window.
type TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule struct {
	// Whether this auto-quarantine rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Time window in seconds over which flakiness is evaluated. Must be greater than 0.
	WindowSeconds *int64 `json:"window_seconds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule instantiates a new TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule() *TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule {
	this := TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule{}
	return &this
}

// NewTestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRuleWithDefaults instantiates a new TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRuleWithDefaults() *TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule {
	this := TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetWindowSeconds returns the WindowSeconds field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule) GetWindowSeconds() int64 {
	if o == nil || o.WindowSeconds == nil {
		var ret int64
		return ret
	}
	return *o.WindowSeconds
}

// GetWindowSecondsOk returns a tuple with the WindowSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule) GetWindowSecondsOk() (*int64, bool) {
	if o == nil || o.WindowSeconds == nil {
		return nil, false
	}
	return o.WindowSeconds, true
}

// HasWindowSeconds returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule) HasWindowSeconds() bool {
	return o != nil && o.WindowSeconds != nil
}

// SetWindowSeconds gets a reference to the given int64 and assigns it to the WindowSeconds field.
func (o *TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule) SetWindowSeconds(v int64) {
	o.WindowSeconds = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.WindowSeconds != nil {
		toSerialize["window_seconds"] = o.WindowSeconds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled       *bool  `json:"enabled,omitempty"`
		WindowSeconds *int64 `json:"window_seconds,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "window_seconds"})
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.WindowSeconds = all.WindowSeconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
