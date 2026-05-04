// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineElasticsearchDestinationCompression Compression configuration for the Elasticsearch destination.
type ObservabilityPipelineElasticsearchDestinationCompression struct {
	// The compression algorithm applied when sending data to Elasticsearch.
	Algorithm ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm `json:"algorithm"`
	// The compression level. Only applicable for `gzip`, `zlib`, and `zstd` algorithms.
	Level *int64 `json:"level,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineElasticsearchDestinationCompression instantiates a new ObservabilityPipelineElasticsearchDestinationCompression object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineElasticsearchDestinationCompression(algorithm ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm) *ObservabilityPipelineElasticsearchDestinationCompression {
	this := ObservabilityPipelineElasticsearchDestinationCompression{}
	this.Algorithm = algorithm
	return &this
}

// NewObservabilityPipelineElasticsearchDestinationCompressionWithDefaults instantiates a new ObservabilityPipelineElasticsearchDestinationCompression object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineElasticsearchDestinationCompressionWithDefaults() *ObservabilityPipelineElasticsearchDestinationCompression {
	this := ObservabilityPipelineElasticsearchDestinationCompression{}
	return &this
}

// GetAlgorithm returns the Algorithm field value.
func (o *ObservabilityPipelineElasticsearchDestinationCompression) GetAlgorithm() ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm {
	if o == nil {
		var ret ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm
		return ret
	}
	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestinationCompression) GetAlgorithmOk() (*ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value.
func (o *ObservabilityPipelineElasticsearchDestinationCompression) SetAlgorithm(v ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm) {
	o.Algorithm = v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *ObservabilityPipelineElasticsearchDestinationCompression) GetLevel() int64 {
	if o == nil || o.Level == nil {
		var ret int64
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestinationCompression) GetLevelOk() (*int64, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *ObservabilityPipelineElasticsearchDestinationCompression) HasLevel() bool {
	return o != nil && o.Level != nil
}

// SetLevel gets a reference to the given int64 and assigns it to the Level field.
func (o *ObservabilityPipelineElasticsearchDestinationCompression) SetLevel(v int64) {
	o.Level = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineElasticsearchDestinationCompression) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["algorithm"] = o.Algorithm
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineElasticsearchDestinationCompression) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Algorithm *ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm `json:"algorithm"`
		Level     *int64                                                             `json:"level,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Algorithm == nil {
		return fmt.Errorf("required field algorithm missing")
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
	o.Level = all.Level

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
