// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDiskBufferOptions Options for configuring a disk buffer.
type ObservabilityPipelineDiskBufferOptions struct {
	// Maximum size of the disk buffer.
	MaxSize int64 `json:"max_size"`
	// The type of the buffer that will be configured, a disk buffer.
	Type *ObservabilityPipelineBufferOptionsDiskType `json:"type,omitempty"`
	// Behavior when the buffer is full (block and stop accepting new events, or drop new events)
	WhenFull *ObservabilityPipelineBufferOptionsWhenFull `json:"when_full,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineDiskBufferOptions instantiates a new ObservabilityPipelineDiskBufferOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineDiskBufferOptions(maxSize int64) *ObservabilityPipelineDiskBufferOptions {
	this := ObservabilityPipelineDiskBufferOptions{}
	this.MaxSize = maxSize
	var typeVar ObservabilityPipelineBufferOptionsDiskType = OBSERVABILITYPIPELINEBUFFEROPTIONSDISKTYPE_DISK
	this.Type = &typeVar
	var whenFull ObservabilityPipelineBufferOptionsWhenFull = OBSERVABILITYPIPELINEBUFFEROPTIONSWHENFULL_BLOCK
	this.WhenFull = &whenFull
	return &this
}

// NewObservabilityPipelineDiskBufferOptionsWithDefaults instantiates a new ObservabilityPipelineDiskBufferOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineDiskBufferOptionsWithDefaults() *ObservabilityPipelineDiskBufferOptions {
	this := ObservabilityPipelineDiskBufferOptions{}
	var typeVar ObservabilityPipelineBufferOptionsDiskType = OBSERVABILITYPIPELINEBUFFEROPTIONSDISKTYPE_DISK
	this.Type = &typeVar
	var whenFull ObservabilityPipelineBufferOptionsWhenFull = OBSERVABILITYPIPELINEBUFFEROPTIONSWHENFULL_BLOCK
	this.WhenFull = &whenFull
	return &this
}

// GetMaxSize returns the MaxSize field value.
func (o *ObservabilityPipelineDiskBufferOptions) GetMaxSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDiskBufferOptions) GetMaxSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxSize, true
}

// SetMaxSize sets field value.
func (o *ObservabilityPipelineDiskBufferOptions) SetMaxSize(v int64) {
	o.MaxSize = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ObservabilityPipelineDiskBufferOptions) GetType() ObservabilityPipelineBufferOptionsDiskType {
	if o == nil || o.Type == nil {
		var ret ObservabilityPipelineBufferOptionsDiskType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDiskBufferOptions) GetTypeOk() (*ObservabilityPipelineBufferOptionsDiskType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ObservabilityPipelineDiskBufferOptions) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ObservabilityPipelineBufferOptionsDiskType and assigns it to the Type field.
func (o *ObservabilityPipelineDiskBufferOptions) SetType(v ObservabilityPipelineBufferOptionsDiskType) {
	o.Type = &v
}

// GetWhenFull returns the WhenFull field value if set, zero value otherwise.
func (o *ObservabilityPipelineDiskBufferOptions) GetWhenFull() ObservabilityPipelineBufferOptionsWhenFull {
	if o == nil || o.WhenFull == nil {
		var ret ObservabilityPipelineBufferOptionsWhenFull
		return ret
	}
	return *o.WhenFull
}

// GetWhenFullOk returns a tuple with the WhenFull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDiskBufferOptions) GetWhenFullOk() (*ObservabilityPipelineBufferOptionsWhenFull, bool) {
	if o == nil || o.WhenFull == nil {
		return nil, false
	}
	return o.WhenFull, true
}

// HasWhenFull returns a boolean if a field has been set.
func (o *ObservabilityPipelineDiskBufferOptions) HasWhenFull() bool {
	return o != nil && o.WhenFull != nil
}

// SetWhenFull gets a reference to the given ObservabilityPipelineBufferOptionsWhenFull and assigns it to the WhenFull field.
func (o *ObservabilityPipelineDiskBufferOptions) SetWhenFull(v ObservabilityPipelineBufferOptionsWhenFull) {
	o.WhenFull = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineDiskBufferOptions) MarshalJSON() ([]byte, error) {
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
func (o *ObservabilityPipelineDiskBufferOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MaxSize  *int64                                      `json:"max_size"`
		Type     *ObservabilityPipelineBufferOptionsDiskType `json:"type,omitempty"`
		WhenFull *ObservabilityPipelineBufferOptionsWhenFull `json:"when_full,omitempty"`
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
