// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3GenericCompressionGzip Gzip compression.
type ObservabilityPipelineAmazonS3GenericCompressionGzip struct {
	// The compression type. Always `gzip`.
	Algorithm ObservabilityPipelineAmazonS3GenericCompressionGzipType `json:"algorithm"`
	// Gzip compression level.
	Level int64 `json:"level"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAmazonS3GenericCompressionGzip instantiates a new ObservabilityPipelineAmazonS3GenericCompressionGzip object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAmazonS3GenericCompressionGzip(algorithm ObservabilityPipelineAmazonS3GenericCompressionGzipType, level int64) *ObservabilityPipelineAmazonS3GenericCompressionGzip {
	this := ObservabilityPipelineAmazonS3GenericCompressionGzip{}
	this.Algorithm = algorithm
	this.Level = level
	return &this
}

// NewObservabilityPipelineAmazonS3GenericCompressionGzipWithDefaults instantiates a new ObservabilityPipelineAmazonS3GenericCompressionGzip object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAmazonS3GenericCompressionGzipWithDefaults() *ObservabilityPipelineAmazonS3GenericCompressionGzip {
	this := ObservabilityPipelineAmazonS3GenericCompressionGzip{}
	var algorithm ObservabilityPipelineAmazonS3GenericCompressionGzipType = OBSERVABILITYPIPELINEAMAZONS3GENERICCOMPRESSIONGZIPTYPE_GZIP
	this.Algorithm = algorithm
	return &this
}

// GetAlgorithm returns the Algorithm field value.
func (o *ObservabilityPipelineAmazonS3GenericCompressionGzip) GetAlgorithm() ObservabilityPipelineAmazonS3GenericCompressionGzipType {
	if o == nil {
		var ret ObservabilityPipelineAmazonS3GenericCompressionGzipType
		return ret
	}
	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericCompressionGzip) GetAlgorithmOk() (*ObservabilityPipelineAmazonS3GenericCompressionGzipType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value.
func (o *ObservabilityPipelineAmazonS3GenericCompressionGzip) SetAlgorithm(v ObservabilityPipelineAmazonS3GenericCompressionGzipType) {
	o.Algorithm = v
}

// GetLevel returns the Level field value.
func (o *ObservabilityPipelineAmazonS3GenericCompressionGzip) GetLevel() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericCompressionGzip) GetLevelOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value.
func (o *ObservabilityPipelineAmazonS3GenericCompressionGzip) SetLevel(v int64) {
	o.Level = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAmazonS3GenericCompressionGzip) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["algorithm"] = o.Algorithm
	toSerialize["level"] = o.Level

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAmazonS3GenericCompressionGzip) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Algorithm *ObservabilityPipelineAmazonS3GenericCompressionGzipType `json:"algorithm"`
		Level     *int64                                                   `json:"level"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Algorithm == nil {
		return fmt.Errorf("required field algorithm missing")
	}
	if all.Level == nil {
		return fmt.Errorf("required field level missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"algorithm", "level"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Algorithm.IsValid() {
		hasInvalidField = true
	} else {
		o.Algorithm = *all.Algorithm
	}
	o.Level = *all.Level

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
