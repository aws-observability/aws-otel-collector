// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketDestinationFramingBytes Event data is not delimited at all.
type ObservabilityPipelineSocketDestinationFramingBytes struct {
	// The definition of `ObservabilityPipelineSocketDestinationFramingBytesMethod` object.
	Method ObservabilityPipelineSocketDestinationFramingBytesMethod `json:"method"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSocketDestinationFramingBytes instantiates a new ObservabilityPipelineSocketDestinationFramingBytes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSocketDestinationFramingBytes(method ObservabilityPipelineSocketDestinationFramingBytesMethod) *ObservabilityPipelineSocketDestinationFramingBytes {
	this := ObservabilityPipelineSocketDestinationFramingBytes{}
	this.Method = method
	return &this
}

// NewObservabilityPipelineSocketDestinationFramingBytesWithDefaults instantiates a new ObservabilityPipelineSocketDestinationFramingBytes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSocketDestinationFramingBytesWithDefaults() *ObservabilityPipelineSocketDestinationFramingBytes {
	this := ObservabilityPipelineSocketDestinationFramingBytes{}
	return &this
}

// GetMethod returns the Method field value.
func (o *ObservabilityPipelineSocketDestinationFramingBytes) GetMethod() ObservabilityPipelineSocketDestinationFramingBytesMethod {
	if o == nil {
		var ret ObservabilityPipelineSocketDestinationFramingBytesMethod
		return ret
	}
	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketDestinationFramingBytes) GetMethodOk() (*ObservabilityPipelineSocketDestinationFramingBytesMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value.
func (o *ObservabilityPipelineSocketDestinationFramingBytes) SetMethod(v ObservabilityPipelineSocketDestinationFramingBytesMethod) {
	o.Method = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSocketDestinationFramingBytes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["method"] = o.Method

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSocketDestinationFramingBytes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Method *ObservabilityPipelineSocketDestinationFramingBytesMethod `json:"method"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Method == nil {
		return fmt.Errorf("required field method missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"method"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Method.IsValid() {
		hasInvalidField = true
	} else {
		o.Method = *all.Method
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
