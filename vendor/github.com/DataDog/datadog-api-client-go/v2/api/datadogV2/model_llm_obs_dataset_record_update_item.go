// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetRecordUpdateItem A record update payload for an LLM Observability dataset.
type LLMObsDatasetRecordUpdateItem struct {
	// Represents any valid JSON value.
	ExpectedOutput NullableAnyValue `json:"expected_output,omitempty"`
	// Unique identifier of the record to update.
	Id string `json:"id"`
	// Represents any valid JSON value.
	Input NullableAnyValue `json:"input,omitempty"`
	// Updated metadata associated with the record.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetRecordUpdateItem instantiates a new LLMObsDatasetRecordUpdateItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetRecordUpdateItem(id string) *LLMObsDatasetRecordUpdateItem {
	this := LLMObsDatasetRecordUpdateItem{}
	this.Id = id
	return &this
}

// NewLLMObsDatasetRecordUpdateItemWithDefaults instantiates a new LLMObsDatasetRecordUpdateItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetRecordUpdateItemWithDefaults() *LLMObsDatasetRecordUpdateItem {
	this := LLMObsDatasetRecordUpdateItem{}
	return &this
}

// GetExpectedOutput returns the ExpectedOutput field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsDatasetRecordUpdateItem) GetExpectedOutput() AnyValue {
	if o == nil || o.ExpectedOutput.Get() == nil {
		var ret AnyValue
		return ret
	}
	return *o.ExpectedOutput.Get()
}

// GetExpectedOutputOk returns a tuple with the ExpectedOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsDatasetRecordUpdateItem) GetExpectedOutputOk() (*AnyValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpectedOutput.Get(), o.ExpectedOutput.IsSet()
}

// HasExpectedOutput returns a boolean if a field has been set.
func (o *LLMObsDatasetRecordUpdateItem) HasExpectedOutput() bool {
	return o != nil && o.ExpectedOutput.IsSet()
}

// SetExpectedOutput gets a reference to the given NullableAnyValue and assigns it to the ExpectedOutput field.
func (o *LLMObsDatasetRecordUpdateItem) SetExpectedOutput(v AnyValue) {
	o.ExpectedOutput.Set(&v)
}

// SetExpectedOutputNil sets the value for ExpectedOutput to be an explicit nil.
func (o *LLMObsDatasetRecordUpdateItem) SetExpectedOutputNil() {
	o.ExpectedOutput.Set(nil)
}

// UnsetExpectedOutput ensures that no value is present for ExpectedOutput, not even an explicit nil.
func (o *LLMObsDatasetRecordUpdateItem) UnsetExpectedOutput() {
	o.ExpectedOutput.Unset()
}

// GetId returns the Id field value.
func (o *LLMObsDatasetRecordUpdateItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordUpdateItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsDatasetRecordUpdateItem) SetId(v string) {
	o.Id = v
}

// GetInput returns the Input field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsDatasetRecordUpdateItem) GetInput() AnyValue {
	if o == nil || o.Input.Get() == nil {
		var ret AnyValue
		return ret
	}
	return *o.Input.Get()
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsDatasetRecordUpdateItem) GetInputOk() (*AnyValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Input.Get(), o.Input.IsSet()
}

// HasInput returns a boolean if a field has been set.
func (o *LLMObsDatasetRecordUpdateItem) HasInput() bool {
	return o != nil && o.Input.IsSet()
}

// SetInput gets a reference to the given NullableAnyValue and assigns it to the Input field.
func (o *LLMObsDatasetRecordUpdateItem) SetInput(v AnyValue) {
	o.Input.Set(&v)
}

// SetInputNil sets the value for Input to be an explicit nil.
func (o *LLMObsDatasetRecordUpdateItem) SetInputNil() {
	o.Input.Set(nil)
}

// UnsetInput ensures that no value is present for Input, not even an explicit nil.
func (o *LLMObsDatasetRecordUpdateItem) UnsetInput() {
	o.Input.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LLMObsDatasetRecordUpdateItem) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordUpdateItem) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LLMObsDatasetRecordUpdateItem) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LLMObsDatasetRecordUpdateItem) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetRecordUpdateItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExpectedOutput.IsSet() {
		toSerialize["expected_output"] = o.ExpectedOutput.Get()
	}
	toSerialize["id"] = o.Id
	if o.Input.IsSet() {
		toSerialize["input"] = o.Input.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDatasetRecordUpdateItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExpectedOutput NullableAnyValue       `json:"expected_output,omitempty"`
		Id             *string                `json:"id"`
		Input          NullableAnyValue       `json:"input,omitempty"`
		Metadata       map[string]interface{} `json:"metadata,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"expected_output", "id", "input", "metadata"})
	} else {
		return err
	}
	o.ExpectedOutput = all.ExpectedOutput
	o.Id = *all.Id
	o.Input = all.Input
	o.Metadata = all.Metadata

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
