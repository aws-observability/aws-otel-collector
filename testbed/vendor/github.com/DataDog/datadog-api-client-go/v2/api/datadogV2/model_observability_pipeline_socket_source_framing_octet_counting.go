// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketSourceFramingOctetCounting Byte frames according to the octet counting format as per RFC6587.
type ObservabilityPipelineSocketSourceFramingOctetCounting struct {
	// Byte frames according to the octet counting format as per RFC6587.
	Method ObservabilityPipelineSocketSourceFramingOctetCountingMethod `json:"method"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSocketSourceFramingOctetCounting instantiates a new ObservabilityPipelineSocketSourceFramingOctetCounting object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSocketSourceFramingOctetCounting(method ObservabilityPipelineSocketSourceFramingOctetCountingMethod) *ObservabilityPipelineSocketSourceFramingOctetCounting {
	this := ObservabilityPipelineSocketSourceFramingOctetCounting{}
	this.Method = method
	return &this
}

// NewObservabilityPipelineSocketSourceFramingOctetCountingWithDefaults instantiates a new ObservabilityPipelineSocketSourceFramingOctetCounting object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSocketSourceFramingOctetCountingWithDefaults() *ObservabilityPipelineSocketSourceFramingOctetCounting {
	this := ObservabilityPipelineSocketSourceFramingOctetCounting{}
	return &this
}

// GetMethod returns the Method field value.
func (o *ObservabilityPipelineSocketSourceFramingOctetCounting) GetMethod() ObservabilityPipelineSocketSourceFramingOctetCountingMethod {
	if o == nil {
		var ret ObservabilityPipelineSocketSourceFramingOctetCountingMethod
		return ret
	}
	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketSourceFramingOctetCounting) GetMethodOk() (*ObservabilityPipelineSocketSourceFramingOctetCountingMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value.
func (o *ObservabilityPipelineSocketSourceFramingOctetCounting) SetMethod(v ObservabilityPipelineSocketSourceFramingOctetCountingMethod) {
	o.Method = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSocketSourceFramingOctetCounting) MarshalJSON() ([]byte, error) {
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
func (o *ObservabilityPipelineSocketSourceFramingOctetCounting) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Method *ObservabilityPipelineSocketSourceFramingOctetCountingMethod `json:"method"`
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
