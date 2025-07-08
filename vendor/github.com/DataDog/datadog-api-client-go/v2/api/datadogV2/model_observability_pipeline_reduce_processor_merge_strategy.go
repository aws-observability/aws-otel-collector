// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineReduceProcessorMergeStrategy Defines how a specific field should be merged across grouped events.
type ObservabilityPipelineReduceProcessorMergeStrategy struct {
	// The field path in the log event.
	Path string `json:"path"`
	// The merge strategy to apply.
	Strategy ObservabilityPipelineReduceProcessorMergeStrategyStrategy `json:"strategy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineReduceProcessorMergeStrategy instantiates a new ObservabilityPipelineReduceProcessorMergeStrategy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineReduceProcessorMergeStrategy(path string, strategy ObservabilityPipelineReduceProcessorMergeStrategyStrategy) *ObservabilityPipelineReduceProcessorMergeStrategy {
	this := ObservabilityPipelineReduceProcessorMergeStrategy{}
	this.Path = path
	this.Strategy = strategy
	return &this
}

// NewObservabilityPipelineReduceProcessorMergeStrategyWithDefaults instantiates a new ObservabilityPipelineReduceProcessorMergeStrategy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineReduceProcessorMergeStrategyWithDefaults() *ObservabilityPipelineReduceProcessorMergeStrategy {
	this := ObservabilityPipelineReduceProcessorMergeStrategy{}
	return &this
}

// GetPath returns the Path field value.
func (o *ObservabilityPipelineReduceProcessorMergeStrategy) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineReduceProcessorMergeStrategy) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value.
func (o *ObservabilityPipelineReduceProcessorMergeStrategy) SetPath(v string) {
	o.Path = v
}

// GetStrategy returns the Strategy field value.
func (o *ObservabilityPipelineReduceProcessorMergeStrategy) GetStrategy() ObservabilityPipelineReduceProcessorMergeStrategyStrategy {
	if o == nil {
		var ret ObservabilityPipelineReduceProcessorMergeStrategyStrategy
		return ret
	}
	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineReduceProcessorMergeStrategy) GetStrategyOk() (*ObservabilityPipelineReduceProcessorMergeStrategyStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value.
func (o *ObservabilityPipelineReduceProcessorMergeStrategy) SetStrategy(v ObservabilityPipelineReduceProcessorMergeStrategyStrategy) {
	o.Strategy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineReduceProcessorMergeStrategy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["path"] = o.Path
	toSerialize["strategy"] = o.Strategy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineReduceProcessorMergeStrategy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Path     *string                                                    `json:"path"`
		Strategy *ObservabilityPipelineReduceProcessorMergeStrategyStrategy `json:"strategy"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Path == nil {
		return fmt.Errorf("required field path missing")
	}
	if all.Strategy == nil {
		return fmt.Errorf("required field strategy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"path", "strategy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Path = *all.Path
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
