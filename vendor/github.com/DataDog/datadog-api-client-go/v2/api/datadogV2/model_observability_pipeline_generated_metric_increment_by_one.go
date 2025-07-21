// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGeneratedMetricIncrementByOne Strategy that increments a generated metric by one for each matching event.
type ObservabilityPipelineGeneratedMetricIncrementByOne struct {
	// Increments the metric by 1 for each matching event.
	Strategy ObservabilityPipelineGeneratedMetricIncrementByOneStrategy `json:"strategy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineGeneratedMetricIncrementByOne instantiates a new ObservabilityPipelineGeneratedMetricIncrementByOne object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineGeneratedMetricIncrementByOne(strategy ObservabilityPipelineGeneratedMetricIncrementByOneStrategy) *ObservabilityPipelineGeneratedMetricIncrementByOne {
	this := ObservabilityPipelineGeneratedMetricIncrementByOne{}
	this.Strategy = strategy
	return &this
}

// NewObservabilityPipelineGeneratedMetricIncrementByOneWithDefaults instantiates a new ObservabilityPipelineGeneratedMetricIncrementByOne object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineGeneratedMetricIncrementByOneWithDefaults() *ObservabilityPipelineGeneratedMetricIncrementByOne {
	this := ObservabilityPipelineGeneratedMetricIncrementByOne{}
	return &this
}

// GetStrategy returns the Strategy field value.
func (o *ObservabilityPipelineGeneratedMetricIncrementByOne) GetStrategy() ObservabilityPipelineGeneratedMetricIncrementByOneStrategy {
	if o == nil {
		var ret ObservabilityPipelineGeneratedMetricIncrementByOneStrategy
		return ret
	}
	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGeneratedMetricIncrementByOne) GetStrategyOk() (*ObservabilityPipelineGeneratedMetricIncrementByOneStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value.
func (o *ObservabilityPipelineGeneratedMetricIncrementByOne) SetStrategy(v ObservabilityPipelineGeneratedMetricIncrementByOneStrategy) {
	o.Strategy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineGeneratedMetricIncrementByOne) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["strategy"] = o.Strategy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineGeneratedMetricIncrementByOne) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Strategy *ObservabilityPipelineGeneratedMetricIncrementByOneStrategy `json:"strategy"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Strategy == nil {
		return fmt.Errorf("required field strategy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"strategy"})
	} else {
		return err
	}

	hasInvalidField := false
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
