// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsSuiteOptions Object describing the extra options for a Synthetic suite.
type SyntheticsSuiteOptions struct {
	// Percentage of critical tests failure needed for a suite to fail.
	AlertingThreshold *float64 `json:"alerting_threshold,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsSuiteOptions instantiates a new SyntheticsSuiteOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsSuiteOptions() *SyntheticsSuiteOptions {
	this := SyntheticsSuiteOptions{}
	return &this
}

// NewSyntheticsSuiteOptionsWithDefaults instantiates a new SyntheticsSuiteOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsSuiteOptionsWithDefaults() *SyntheticsSuiteOptions {
	this := SyntheticsSuiteOptions{}
	return &this
}

// GetAlertingThreshold returns the AlertingThreshold field value if set, zero value otherwise.
func (o *SyntheticsSuiteOptions) GetAlertingThreshold() float64 {
	if o == nil || o.AlertingThreshold == nil {
		var ret float64
		return ret
	}
	return *o.AlertingThreshold
}

// GetAlertingThresholdOk returns a tuple with the AlertingThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSuiteOptions) GetAlertingThresholdOk() (*float64, bool) {
	if o == nil || o.AlertingThreshold == nil {
		return nil, false
	}
	return o.AlertingThreshold, true
}

// HasAlertingThreshold returns a boolean if a field has been set.
func (o *SyntheticsSuiteOptions) HasAlertingThreshold() bool {
	return o != nil && o.AlertingThreshold != nil
}

// SetAlertingThreshold gets a reference to the given float64 and assigns it to the AlertingThreshold field.
func (o *SyntheticsSuiteOptions) SetAlertingThreshold(v float64) {
	o.AlertingThreshold = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsSuiteOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AlertingThreshold != nil {
		toSerialize["alerting_threshold"] = o.AlertingThreshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsSuiteOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlertingThreshold *float64 `json:"alerting_threshold,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alerting_threshold"})
	} else {
		return err
	}
	o.AlertingThreshold = all.AlertingThreshold

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
