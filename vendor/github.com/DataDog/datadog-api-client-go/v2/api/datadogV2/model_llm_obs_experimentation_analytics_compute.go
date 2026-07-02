// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationAnalyticsCompute A single metric computation definition.
type LLMObsExperimentationAnalyticsCompute struct {
	// Name of the metric to compute.
	Metric string `json:"metric"`
	// Optional alias for this computation in the response.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentationAnalyticsCompute instantiates a new LLMObsExperimentationAnalyticsCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentationAnalyticsCompute(metric string) *LLMObsExperimentationAnalyticsCompute {
	this := LLMObsExperimentationAnalyticsCompute{}
	this.Metric = metric
	return &this
}

// NewLLMObsExperimentationAnalyticsComputeWithDefaults instantiates a new LLMObsExperimentationAnalyticsCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentationAnalyticsComputeWithDefaults() *LLMObsExperimentationAnalyticsCompute {
	this := LLMObsExperimentationAnalyticsCompute{}
	return &this
}

// GetMetric returns the Metric field value.
func (o *LLMObsExperimentationAnalyticsCompute) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationAnalyticsCompute) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *LLMObsExperimentationAnalyticsCompute) SetMetric(v string) {
	o.Metric = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LLMObsExperimentationAnalyticsCompute) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationAnalyticsCompute) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LLMObsExperimentationAnalyticsCompute) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LLMObsExperimentationAnalyticsCompute) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentationAnalyticsCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["metric"] = o.Metric
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentationAnalyticsCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Metric *string `json:"metric"`
		Name   *string `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Metric == nil {
		return fmt.Errorf("required field metric missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"metric", "name"})
	} else {
		return err
	}
	o.Metric = *all.Metric
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
