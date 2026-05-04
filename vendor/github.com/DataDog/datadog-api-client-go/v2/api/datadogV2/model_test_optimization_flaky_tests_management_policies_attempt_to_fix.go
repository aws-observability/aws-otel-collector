// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationFlakyTestsManagementPoliciesAttemptToFix Configuration for the attempt-to-fix Flaky Tests Management policy.
type TestOptimizationFlakyTestsManagementPoliciesAttemptToFix struct {
	// Number of retries when attempting to fix a flaky test. Must be greater than 0.
	Retries *int64 `json:"retries,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTestOptimizationFlakyTestsManagementPoliciesAttemptToFix instantiates a new TestOptimizationFlakyTestsManagementPoliciesAttemptToFix object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTestOptimizationFlakyTestsManagementPoliciesAttemptToFix() *TestOptimizationFlakyTestsManagementPoliciesAttemptToFix {
	this := TestOptimizationFlakyTestsManagementPoliciesAttemptToFix{}
	return &this
}

// NewTestOptimizationFlakyTestsManagementPoliciesAttemptToFixWithDefaults instantiates a new TestOptimizationFlakyTestsManagementPoliciesAttemptToFix object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTestOptimizationFlakyTestsManagementPoliciesAttemptToFixWithDefaults() *TestOptimizationFlakyTestsManagementPoliciesAttemptToFix {
	this := TestOptimizationFlakyTestsManagementPoliciesAttemptToFix{}
	return &this
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttemptToFix) GetRetries() int64 {
	if o == nil || o.Retries == nil {
		var ret int64
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttemptToFix) GetRetriesOk() (*int64, bool) {
	if o == nil || o.Retries == nil {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttemptToFix) HasRetries() bool {
	return o != nil && o.Retries != nil
}

// SetRetries gets a reference to the given int64 and assigns it to the Retries field.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttemptToFix) SetRetries(v int64) {
	o.Retries = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TestOptimizationFlakyTestsManagementPoliciesAttemptToFix) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Retries != nil {
		toSerialize["retries"] = o.Retries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttemptToFix) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Retries *int64 `json:"retries,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"retries"})
	} else {
		return err
	}
	o.Retries = all.Retries

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
