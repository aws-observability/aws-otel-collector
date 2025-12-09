// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketDestinationFramingCharacterDelimited Each log event is separated using the specified delimiter character.
type ObservabilityPipelineSocketDestinationFramingCharacterDelimited struct {
	// A single ASCII character used as a delimiter.
	Delimiter string `json:"delimiter"`
	// The definition of `ObservabilityPipelineSocketDestinationFramingCharacterDelimitedMethod` object.
	Method ObservabilityPipelineSocketDestinationFramingCharacterDelimitedMethod `json:"method"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSocketDestinationFramingCharacterDelimited instantiates a new ObservabilityPipelineSocketDestinationFramingCharacterDelimited object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSocketDestinationFramingCharacterDelimited(delimiter string, method ObservabilityPipelineSocketDestinationFramingCharacterDelimitedMethod) *ObservabilityPipelineSocketDestinationFramingCharacterDelimited {
	this := ObservabilityPipelineSocketDestinationFramingCharacterDelimited{}
	this.Delimiter = delimiter
	this.Method = method
	return &this
}

// NewObservabilityPipelineSocketDestinationFramingCharacterDelimitedWithDefaults instantiates a new ObservabilityPipelineSocketDestinationFramingCharacterDelimited object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSocketDestinationFramingCharacterDelimitedWithDefaults() *ObservabilityPipelineSocketDestinationFramingCharacterDelimited {
	this := ObservabilityPipelineSocketDestinationFramingCharacterDelimited{}
	return &this
}

// GetDelimiter returns the Delimiter field value.
func (o *ObservabilityPipelineSocketDestinationFramingCharacterDelimited) GetDelimiter() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Delimiter
}

// GetDelimiterOk returns a tuple with the Delimiter field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketDestinationFramingCharacterDelimited) GetDelimiterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delimiter, true
}

// SetDelimiter sets field value.
func (o *ObservabilityPipelineSocketDestinationFramingCharacterDelimited) SetDelimiter(v string) {
	o.Delimiter = v
}

// GetMethod returns the Method field value.
func (o *ObservabilityPipelineSocketDestinationFramingCharacterDelimited) GetMethod() ObservabilityPipelineSocketDestinationFramingCharacterDelimitedMethod {
	if o == nil {
		var ret ObservabilityPipelineSocketDestinationFramingCharacterDelimitedMethod
		return ret
	}
	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketDestinationFramingCharacterDelimited) GetMethodOk() (*ObservabilityPipelineSocketDestinationFramingCharacterDelimitedMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value.
func (o *ObservabilityPipelineSocketDestinationFramingCharacterDelimited) SetMethod(v ObservabilityPipelineSocketDestinationFramingCharacterDelimitedMethod) {
	o.Method = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSocketDestinationFramingCharacterDelimited) MarshalJSON() ([]byte, error) {
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
func (o *ObservabilityPipelineSocketDestinationFramingCharacterDelimited) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Delimiter *string                                                                `json:"delimiter"`
		Method    *ObservabilityPipelineSocketDestinationFramingCharacterDelimitedMethod `json:"method"`
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
