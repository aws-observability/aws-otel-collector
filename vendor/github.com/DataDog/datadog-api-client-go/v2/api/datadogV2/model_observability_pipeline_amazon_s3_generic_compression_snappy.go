// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3GenericCompressionSnappy Snappy compression.
type ObservabilityPipelineAmazonS3GenericCompressionSnappy struct {
	// The compression type. Always `snappy`.
	Algorithm ObservabilityPipelineAmazonS3GenericCompressionSnappyType `json:"algorithm"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAmazonS3GenericCompressionSnappy instantiates a new ObservabilityPipelineAmazonS3GenericCompressionSnappy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAmazonS3GenericCompressionSnappy(algorithm ObservabilityPipelineAmazonS3GenericCompressionSnappyType) *ObservabilityPipelineAmazonS3GenericCompressionSnappy {
	this := ObservabilityPipelineAmazonS3GenericCompressionSnappy{}
	this.Algorithm = algorithm
	return &this
}

// NewObservabilityPipelineAmazonS3GenericCompressionSnappyWithDefaults instantiates a new ObservabilityPipelineAmazonS3GenericCompressionSnappy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAmazonS3GenericCompressionSnappyWithDefaults() *ObservabilityPipelineAmazonS3GenericCompressionSnappy {
	this := ObservabilityPipelineAmazonS3GenericCompressionSnappy{}
	var algorithm ObservabilityPipelineAmazonS3GenericCompressionSnappyType = OBSERVABILITYPIPELINEAMAZONS3GENERICCOMPRESSIONSNAPPYTYPE_SNAPPY
	this.Algorithm = algorithm
	return &this
}

// GetAlgorithm returns the Algorithm field value.
func (o *ObservabilityPipelineAmazonS3GenericCompressionSnappy) GetAlgorithm() ObservabilityPipelineAmazonS3GenericCompressionSnappyType {
	if o == nil {
		var ret ObservabilityPipelineAmazonS3GenericCompressionSnappyType
		return ret
	}
	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericCompressionSnappy) GetAlgorithmOk() (*ObservabilityPipelineAmazonS3GenericCompressionSnappyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value.
func (o *ObservabilityPipelineAmazonS3GenericCompressionSnappy) SetAlgorithm(v ObservabilityPipelineAmazonS3GenericCompressionSnappyType) {
	o.Algorithm = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAmazonS3GenericCompressionSnappy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["algorithm"] = o.Algorithm

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAmazonS3GenericCompressionSnappy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Algorithm *ObservabilityPipelineAmazonS3GenericCompressionSnappyType `json:"algorithm"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Algorithm == nil {
		return fmt.Errorf("required field algorithm missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"algorithm"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Algorithm.IsValid() {
		hasInvalidField = true
	} else {
		o.Algorithm = *all.Algorithm
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
