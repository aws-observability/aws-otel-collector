// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit A cardinality override for a specific tag key within a per-metric limit.
type ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit struct {
	// How the per-tag override is applied. `limit_override` enforces a custom limit on the tag; `excluded` skips the tag from cardinality tracking.
	Mode ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode `json:"mode"`
	// The tag key this override applies to.
	TagKey string `json:"tag_key"`
	// The maximum number of distinct values allowed for this tag. Required when `mode` is `limit_override`. Must be omitted when `mode` is `excluded`.
	ValueLimit *int64 `json:"value_limit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit instantiates a new ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit(mode ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode, tagKey string) *ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit {
	this := ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit{}
	this.Mode = mode
	this.TagKey = tagKey
	return &this
}

// NewObservabilityPipelineTagCardinalityLimitProcessorPerTagLimitWithDefaults instantiates a new ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineTagCardinalityLimitProcessorPerTagLimitWithDefaults() *ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit {
	this := ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit{}
	return &this
}

// GetMode returns the Mode field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit) GetMode() ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode {
	if o == nil {
		var ret ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit) GetModeOk() (*ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit) SetMode(v ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode) {
	o.Mode = v
}

// GetTagKey returns the TagKey field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit) GetTagKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagKey, true
}

// SetTagKey sets field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit) SetTagKey(v string) {
	o.TagKey = v
}

// GetValueLimit returns the ValueLimit field value if set, zero value otherwise.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit) GetValueLimit() int64 {
	if o == nil || o.ValueLimit == nil {
		var ret int64
		return ret
	}
	return *o.ValueLimit
}

// GetValueLimitOk returns a tuple with the ValueLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit) GetValueLimitOk() (*int64, bool) {
	if o == nil || o.ValueLimit == nil {
		return nil, false
	}
	return o.ValueLimit, true
}

// HasValueLimit returns a boolean if a field has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit) HasValueLimit() bool {
	return o != nil && o.ValueLimit != nil
}

// SetValueLimit gets a reference to the given int64 and assigns it to the ValueLimit field.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit) SetValueLimit(v int64) {
	o.ValueLimit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["mode"] = o.Mode
	toSerialize["tag_key"] = o.TagKey
	if o.ValueLimit != nil {
		toSerialize["value_limit"] = o.ValueLimit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineTagCardinalityLimitProcessorPerTagLimit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode       *ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode `json:"mode"`
		TagKey     *string                                                      `json:"tag_key"`
		ValueLimit *int64                                                       `json:"value_limit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.TagKey == nil {
		return fmt.Errorf("required field tag_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"mode", "tag_key", "value_limit"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	o.TagKey = *all.TagKey
	o.ValueLimit = all.ValueLimit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
