// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketSourceFramingCharacterDelimited Byte frames which are delimited by a chosen character.
type ObservabilityPipelineSocketSourceFramingCharacterDelimited struct {
	// A single ASCII character used to delimit events.
	Delimiter string `json:"delimiter"`
	// Byte frames which are delimited by a chosen character.
	Method ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod `json:"method"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSocketSourceFramingCharacterDelimited instantiates a new ObservabilityPipelineSocketSourceFramingCharacterDelimited object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSocketSourceFramingCharacterDelimited(delimiter string, method ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod) *ObservabilityPipelineSocketSourceFramingCharacterDelimited {
	this := ObservabilityPipelineSocketSourceFramingCharacterDelimited{}
	this.Delimiter = delimiter
	this.Method = method
	return &this
}

// NewObservabilityPipelineSocketSourceFramingCharacterDelimitedWithDefaults instantiates a new ObservabilityPipelineSocketSourceFramingCharacterDelimited object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSocketSourceFramingCharacterDelimitedWithDefaults() *ObservabilityPipelineSocketSourceFramingCharacterDelimited {
	this := ObservabilityPipelineSocketSourceFramingCharacterDelimited{}
	return &this
}

// GetDelimiter returns the Delimiter field value.
func (o *ObservabilityPipelineSocketSourceFramingCharacterDelimited) GetDelimiter() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Delimiter
}

// GetDelimiterOk returns a tuple with the Delimiter field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketSourceFramingCharacterDelimited) GetDelimiterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delimiter, true
}

// SetDelimiter sets field value.
func (o *ObservabilityPipelineSocketSourceFramingCharacterDelimited) SetDelimiter(v string) {
	o.Delimiter = v
}

// GetMethod returns the Method field value.
func (o *ObservabilityPipelineSocketSourceFramingCharacterDelimited) GetMethod() ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod {
	if o == nil {
		var ret ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod
		return ret
	}
	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketSourceFramingCharacterDelimited) GetMethodOk() (*ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value.
func (o *ObservabilityPipelineSocketSourceFramingCharacterDelimited) SetMethod(v ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod) {
	o.Method = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSocketSourceFramingCharacterDelimited) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["delimiter"] = o.Delimiter
	toSerialize["method"] = o.Method

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSocketSourceFramingCharacterDelimited) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Delimiter *string                                                           `json:"delimiter"`
		Method    *ObservabilityPipelineSocketSourceFramingCharacterDelimitedMethod `json:"method"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Delimiter == nil {
		return fmt.Errorf("required field delimiter missing")
	}
	if all.Method == nil {
		return fmt.Errorf("required field method missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"delimiter", "method"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Delimiter = *all.Delimiter
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
