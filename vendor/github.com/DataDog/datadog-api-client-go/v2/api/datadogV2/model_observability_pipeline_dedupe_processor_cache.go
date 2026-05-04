// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDedupeProcessorCache Configuration for the cache used to detect duplicates.
type ObservabilityPipelineDedupeProcessorCache struct {
	// The number of events to cache for duplicate detection.
	NumEvents int64 `json:"num_events"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineDedupeProcessorCache instantiates a new ObservabilityPipelineDedupeProcessorCache object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineDedupeProcessorCache(numEvents int64) *ObservabilityPipelineDedupeProcessorCache {
	this := ObservabilityPipelineDedupeProcessorCache{}
	this.NumEvents = numEvents
	return &this
}

// NewObservabilityPipelineDedupeProcessorCacheWithDefaults instantiates a new ObservabilityPipelineDedupeProcessorCache object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineDedupeProcessorCacheWithDefaults() *ObservabilityPipelineDedupeProcessorCache {
	this := ObservabilityPipelineDedupeProcessorCache{}
	return &this
}

// GetNumEvents returns the NumEvents field value.
func (o *ObservabilityPipelineDedupeProcessorCache) GetNumEvents() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.NumEvents
}

// GetNumEventsOk returns a tuple with the NumEvents field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDedupeProcessorCache) GetNumEventsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumEvents, true
}

// SetNumEvents sets field value.
func (o *ObservabilityPipelineDedupeProcessorCache) SetNumEvents(v int64) {
	o.NumEvents = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineDedupeProcessorCache) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["num_events"] = o.NumEvents

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineDedupeProcessorCache) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NumEvents *int64 `json:"num_events"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.NumEvents == nil {
		return fmt.Errorf("required field num_events missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"num_events"})
	} else {
		return err
	}
	o.NumEvents = *all.NumEvents

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
