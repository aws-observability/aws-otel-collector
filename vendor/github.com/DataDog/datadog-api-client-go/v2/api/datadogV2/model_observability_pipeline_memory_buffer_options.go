// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMemoryBufferOptions Options for configuring a memory buffer by byte size.
type ObservabilityPipelineMemoryBufferOptions struct {
	// Maximum size of the memory buffer.
	MaxSize int64 `json:"max_size"`
	// The type of the buffer that will be configured, a memory buffer.
	Type *ObservabilityPipelineBufferOptionsMemoryType `json:"type,omitempty"`
	// Behavior when the buffer is full (block and stop accepting new events, or drop new events)
	WhenFull *ObservabilityPipelineBufferOptionsWhenFull `json:"when_full,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineMemoryBufferOptions instantiates a new ObservabilityPipelineMemoryBufferOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineMemoryBufferOptions(maxSize int64) *ObservabilityPipelineMemoryBufferOptions {
	this := ObservabilityPipelineMemoryBufferOptions{}
	this.MaxSize = maxSize
	var typeVar ObservabilityPipelineBufferOptionsMemoryType = OBSERVABILITYPIPELINEBUFFEROPTIONSMEMORYTYPE_MEMORY
	this.Type = &typeVar
	var whenFull ObservabilityPipelineBufferOptionsWhenFull = OBSERVABILITYPIPELINEBUFFEROPTIONSWHENFULL_BLOCK
	this.WhenFull = &whenFull
	return &this
}

// NewObservabilityPipelineMemoryBufferOptionsWithDefaults instantiates a new ObservabilityPipelineMemoryBufferOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineMemoryBufferOptionsWithDefaults() *ObservabilityPipelineMemoryBufferOptions {
	this := ObservabilityPipelineMemoryBufferOptions{}
	var typeVar ObservabilityPipelineBufferOptionsMemoryType = OBSERVABILITYPIPELINEBUFFEROPTIONSMEMORYTYPE_MEMORY
	this.Type = &typeVar
	var whenFull ObservabilityPipelineBufferOptionsWhenFull = OBSERVABILITYPIPELINEBUFFEROPTIONSWHENFULL_BLOCK
	this.WhenFull = &whenFull
	return &this
}

// GetMaxSize returns the MaxSize field value.
func (o *ObservabilityPipelineMemoryBufferOptions) GetMaxSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMemoryBufferOptions) GetMaxSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxSize, true
}

// SetMaxSize sets field value.
func (o *ObservabilityPipelineMemoryBufferOptions) SetMaxSize(v int64) {
	o.MaxSize = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ObservabilityPipelineMemoryBufferOptions) GetType() ObservabilityPipelineBufferOptionsMemoryType {
	if o == nil || o.Type == nil {
		var ret ObservabilityPipelineBufferOptionsMemoryType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMemoryBufferOptions) GetTypeOk() (*ObservabilityPipelineBufferOptionsMemoryType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ObservabilityPipelineMemoryBufferOptions) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ObservabilityPipelineBufferOptionsMemoryType and assigns it to the Type field.
func (o *ObservabilityPipelineMemoryBufferOptions) SetType(v ObservabilityPipelineBufferOptionsMemoryType) {
	o.Type = &v
}

// GetWhenFull returns the WhenFull field value if set, zero value otherwise.
func (o *ObservabilityPipelineMemoryBufferOptions) GetWhenFull() ObservabilityPipelineBufferOptionsWhenFull {
	if o == nil || o.WhenFull == nil {
		var ret ObservabilityPipelineBufferOptionsWhenFull
		return ret
	}
	return *o.WhenFull
}

// GetWhenFullOk returns a tuple with the WhenFull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMemoryBufferOptions) GetWhenFullOk() (*ObservabilityPipelineBufferOptionsWhenFull, bool) {
	if o == nil || o.WhenFull == nil {
		return nil, false
	}
	return o.WhenFull, true
}

// HasWhenFull returns a boolean if a field has been set.
func (o *ObservabilityPipelineMemoryBufferOptions) HasWhenFull() bool {
	return o != nil && o.WhenFull != nil
}

// SetWhenFull gets a reference to the given ObservabilityPipelineBufferOptionsWhenFull and assigns it to the WhenFull field.
func (o *ObservabilityPipelineMemoryBufferOptions) SetWhenFull(v ObservabilityPipelineBufferOptionsWhenFull) {
	o.WhenFull = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineMemoryBufferOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["max_size"] = o.MaxSize
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.WhenFull != nil {
		toSerialize["when_full"] = o.WhenFull
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineMemoryBufferOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MaxSize  *int64                                        `json:"max_size"`
		Type     *ObservabilityPipelineBufferOptionsMemoryType `json:"type,omitempty"`
		WhenFull *ObservabilityPipelineBufferOptionsWhenFull   `json:"when_full,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MaxSize == nil {
		return fmt.Errorf("required field max_size missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max_size", "type", "when_full"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MaxSize = *all.MaxSize
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	if all.WhenFull != nil && !all.WhenFull.IsValid() {
		hasInvalidField = true
	} else {
		o.WhenFull = all.WhenFull
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
