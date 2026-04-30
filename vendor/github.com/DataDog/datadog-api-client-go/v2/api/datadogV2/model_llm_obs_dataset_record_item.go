// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetRecordItem A single record to append to an LLM Observability dataset.
type LLMObsDatasetRecordItem struct {
	// Represents any valid JSON value.
	ExpectedOutput NullableAnyValue `json:"expected_output,omitempty"`
	// Represents any valid JSON value.
	Input NullableAnyValue `json:"input"`
	// Arbitrary metadata associated with the record.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetRecordItem instantiates a new LLMObsDatasetRecordItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetRecordItem(input NullableAnyValue) *LLMObsDatasetRecordItem {
	this := LLMObsDatasetRecordItem{}
	this.Input = input
	return &this
}

// NewLLMObsDatasetRecordItemWithDefaults instantiates a new LLMObsDatasetRecordItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetRecordItemWithDefaults() *LLMObsDatasetRecordItem {
	this := LLMObsDatasetRecordItem{}
	return &this
}

// GetExpectedOutput returns the ExpectedOutput field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsDatasetRecordItem) GetExpectedOutput() AnyValue {
	if o == nil || o.ExpectedOutput.Get() == nil {
		var ret AnyValue
		return ret
	}
	return *o.ExpectedOutput.Get()
}

// GetExpectedOutputOk returns a tuple with the ExpectedOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsDatasetRecordItem) GetExpectedOutputOk() (*AnyValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpectedOutput.Get(), o.ExpectedOutput.IsSet()
}

// HasExpectedOutput returns a boolean if a field has been set.
func (o *LLMObsDatasetRecordItem) HasExpectedOutput() bool {
	return o != nil && o.ExpectedOutput.IsSet()
}

// SetExpectedOutput gets a reference to the given NullableAnyValue and assigns it to the ExpectedOutput field.
func (o *LLMObsDatasetRecordItem) SetExpectedOutput(v AnyValue) {
	o.ExpectedOutput.Set(&v)
}

// SetExpectedOutputNil sets the value for ExpectedOutput to be an explicit nil.
func (o *LLMObsDatasetRecordItem) SetExpectedOutputNil() {
	o.ExpectedOutput.Set(nil)
}

// UnsetExpectedOutput ensures that no value is present for ExpectedOutput, not even an explicit nil.
func (o *LLMObsDatasetRecordItem) UnsetExpectedOutput() {
	o.ExpectedOutput.Unset()
}

// GetInput returns the Input field value.
// If the value is explicit nil, the zero value for AnyValue will be returned.
func (o *LLMObsDatasetRecordItem) GetInput() AnyValue {
	if o == nil || o.Input.Get() == nil {
		var ret AnyValue
		return ret
	}
	return *o.Input.Get()
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsDatasetRecordItem) GetInputOk() (*AnyValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Input.Get(), o.Input.IsSet()
}

// SetInput sets field value.
func (o *LLMObsDatasetRecordItem) SetInput(v AnyValue) {
	o.Input.Set(&v)
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LLMObsDatasetRecordItem) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordItem) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LLMObsDatasetRecordItem) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LLMObsDatasetRecordItem) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetRecordItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExpectedOutput.IsSet() {
		toSerialize["expected_output"] = o.ExpectedOutput.Get()
	}
	toSerialize["input"] = o.Input.Get()
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDatasetRecordItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExpectedOutput NullableAnyValue       `json:"expected_output,omitempty"`
		Input          NullableAnyValue       `json:"input"`
		Metadata       map[string]interface{} `json:"metadata,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Input.IsSet() {
		return fmt.Errorf("required field input missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"expected_output", "input", "metadata"})
	} else {
		return err
	}
	o.ExpectedOutput = all.ExpectedOutput
	o.Input = all.Input
	o.Metadata = all.Metadata

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
