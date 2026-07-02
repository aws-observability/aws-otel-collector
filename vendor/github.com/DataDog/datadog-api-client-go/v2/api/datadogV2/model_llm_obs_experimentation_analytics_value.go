// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationAnalyticsValue A single analytics result bucket.
type LLMObsExperimentationAnalyticsValue struct {
	// The group-by field values for this bucket.
	By map[string]interface{} `json:"by,omitempty"`
	// Computed metric values for this bucket.
	Metrics map[string]interface{} `json:"metrics"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentationAnalyticsValue instantiates a new LLMObsExperimentationAnalyticsValue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentationAnalyticsValue(metrics map[string]interface{}) *LLMObsExperimentationAnalyticsValue {
	this := LLMObsExperimentationAnalyticsValue{}
	this.Metrics = metrics
	return &this
}

// NewLLMObsExperimentationAnalyticsValueWithDefaults instantiates a new LLMObsExperimentationAnalyticsValue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentationAnalyticsValueWithDefaults() *LLMObsExperimentationAnalyticsValue {
	this := LLMObsExperimentationAnalyticsValue{}
	return &this
}

// GetBy returns the By field value if set, zero value otherwise.
func (o *LLMObsExperimentationAnalyticsValue) GetBy() map[string]interface{} {
	if o == nil || o.By == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.By
}

// GetByOk returns a tuple with the By field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationAnalyticsValue) GetByOk() (*map[string]interface{}, bool) {
	if o == nil || o.By == nil {
		return nil, false
	}
	return &o.By, true
}

// HasBy returns a boolean if a field has been set.
func (o *LLMObsExperimentationAnalyticsValue) HasBy() bool {
	return o != nil && o.By != nil
}

// SetBy gets a reference to the given map[string]interface{} and assigns it to the By field.
func (o *LLMObsExperimentationAnalyticsValue) SetBy(v map[string]interface{}) {
	o.By = v
}

// GetMetrics returns the Metrics field value.
func (o *LLMObsExperimentationAnalyticsValue) GetMetrics() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationAnalyticsValue) GetMetricsOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value.
func (o *LLMObsExperimentationAnalyticsValue) SetMetrics(v map[string]interface{}) {
	o.Metrics = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentationAnalyticsValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.By != nil {
		toSerialize["by"] = o.By
	}
	toSerialize["metrics"] = o.Metrics

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentationAnalyticsValue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		By      map[string]interface{}  `json:"by,omitempty"`
		Metrics *map[string]interface{} `json:"metrics"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Metrics == nil {
		return fmt.Errorf("required field metrics missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"by", "metrics"})
	} else {
		return err
	}
	o.By = all.By
	o.Metrics = *all.Metrics

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
