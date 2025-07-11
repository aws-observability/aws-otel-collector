// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGeneratedMetricIncrementByField Strategy that increments a generated metric based on the value of a log field.
type ObservabilityPipelineGeneratedMetricIncrementByField struct {
	// Name of the log field containing the numeric value to increment the metric by.
	Field string `json:"field"`
	// Uses a numeric field in the log event as the metric increment.
	Strategy ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy `json:"strategy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineGeneratedMetricIncrementByField instantiates a new ObservabilityPipelineGeneratedMetricIncrementByField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineGeneratedMetricIncrementByField(field string, strategy ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy) *ObservabilityPipelineGeneratedMetricIncrementByField {
	this := ObservabilityPipelineGeneratedMetricIncrementByField{}
	this.Field = field
	this.Strategy = strategy
	return &this
}

// NewObservabilityPipelineGeneratedMetricIncrementByFieldWithDefaults instantiates a new ObservabilityPipelineGeneratedMetricIncrementByField object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineGeneratedMetricIncrementByFieldWithDefaults() *ObservabilityPipelineGeneratedMetricIncrementByField {
	this := ObservabilityPipelineGeneratedMetricIncrementByField{}
	return &this
}

// GetField returns the Field field value.
func (o *ObservabilityPipelineGeneratedMetricIncrementByField) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGeneratedMetricIncrementByField) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *ObservabilityPipelineGeneratedMetricIncrementByField) SetField(v string) {
	o.Field = v
}

// GetStrategy returns the Strategy field value.
func (o *ObservabilityPipelineGeneratedMetricIncrementByField) GetStrategy() ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy {
	if o == nil {
		var ret ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy
		return ret
	}
	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGeneratedMetricIncrementByField) GetStrategyOk() (*ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value.
func (o *ObservabilityPipelineGeneratedMetricIncrementByField) SetStrategy(v ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy) {
	o.Strategy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineGeneratedMetricIncrementByField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["field"] = o.Field
	toSerialize["strategy"] = o.Strategy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineGeneratedMetricIncrementByField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Field    *string                                                       `json:"field"`
		Strategy *ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy `json:"strategy"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}
	if all.Strategy == nil {
		return fmt.Errorf("required field strategy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"field", "strategy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Field = *all.Field
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
