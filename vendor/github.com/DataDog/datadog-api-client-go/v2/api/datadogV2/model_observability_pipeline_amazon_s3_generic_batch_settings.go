// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3GenericBatchSettings Event batching settings
type ObservabilityPipelineAmazonS3GenericBatchSettings struct {
	// Maximum batch size in bytes.
	BatchSize *int64 `json:"batch_size,omitempty"`
	// Maximum number of seconds to wait before flushing the batch.
	TimeoutSecs *int64 `json:"timeout_secs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAmazonS3GenericBatchSettings instantiates a new ObservabilityPipelineAmazonS3GenericBatchSettings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAmazonS3GenericBatchSettings() *ObservabilityPipelineAmazonS3GenericBatchSettings {
	this := ObservabilityPipelineAmazonS3GenericBatchSettings{}
	return &this
}

// NewObservabilityPipelineAmazonS3GenericBatchSettingsWithDefaults instantiates a new ObservabilityPipelineAmazonS3GenericBatchSettings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAmazonS3GenericBatchSettingsWithDefaults() *ObservabilityPipelineAmazonS3GenericBatchSettings {
	this := ObservabilityPipelineAmazonS3GenericBatchSettings{}
	return &this
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonS3GenericBatchSettings) GetBatchSize() int64 {
	if o == nil || o.BatchSize == nil {
		var ret int64
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericBatchSettings) GetBatchSizeOk() (*int64, bool) {
	if o == nil || o.BatchSize == nil {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonS3GenericBatchSettings) HasBatchSize() bool {
	return o != nil && o.BatchSize != nil
}

// SetBatchSize gets a reference to the given int64 and assigns it to the BatchSize field.
func (o *ObservabilityPipelineAmazonS3GenericBatchSettings) SetBatchSize(v int64) {
	o.BatchSize = &v
}

// GetTimeoutSecs returns the TimeoutSecs field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonS3GenericBatchSettings) GetTimeoutSecs() int64 {
	if o == nil || o.TimeoutSecs == nil {
		var ret int64
		return ret
	}
	return *o.TimeoutSecs
}

// GetTimeoutSecsOk returns a tuple with the TimeoutSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericBatchSettings) GetTimeoutSecsOk() (*int64, bool) {
	if o == nil || o.TimeoutSecs == nil {
		return nil, false
	}
	return o.TimeoutSecs, true
}

// HasTimeoutSecs returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonS3GenericBatchSettings) HasTimeoutSecs() bool {
	return o != nil && o.TimeoutSecs != nil
}

// SetTimeoutSecs gets a reference to the given int64 and assigns it to the TimeoutSecs field.
func (o *ObservabilityPipelineAmazonS3GenericBatchSettings) SetTimeoutSecs(v int64) {
	o.TimeoutSecs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAmazonS3GenericBatchSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BatchSize != nil {
		toSerialize["batch_size"] = o.BatchSize
	}
	if o.TimeoutSecs != nil {
		toSerialize["timeout_secs"] = o.TimeoutSecs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAmazonS3GenericBatchSettings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BatchSize   *int64 `json:"batch_size,omitempty"`
		TimeoutSecs *int64 `json:"timeout_secs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"batch_size", "timeout_secs"})
	} else {
		return err
	}
	o.BatchSize = all.BatchSize
	o.TimeoutSecs = all.TimeoutSecs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
