// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetBatchUpdateUpdateRecord A record update payload as part of a batch update on an LLM Observability dataset.
type LLMObsDatasetBatchUpdateUpdateRecord struct {
	// Represents any valid JSON value.
	ExpectedOutput NullableAnyValue `json:"expected_output,omitempty"`
	// Unique identifier of the record to update.
	Id string `json:"id"`
	// Represents any valid JSON value.
	Input NullableAnyValue `json:"input,omitempty"`
	// Updated metadata associated with the record.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Explicit tag operations for updating records. Operations are applied in order, Remove then Add then Set. `set` is the final override; if specified, the result of `remove` and `add` is discarded.
	TagOperations *LLMObsDatasetRecordTagOperations `json:"tag_operations,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetBatchUpdateUpdateRecord instantiates a new LLMObsDatasetBatchUpdateUpdateRecord object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetBatchUpdateUpdateRecord(id string) *LLMObsDatasetBatchUpdateUpdateRecord {
	this := LLMObsDatasetBatchUpdateUpdateRecord{}
	this.Id = id
	return &this
}

// NewLLMObsDatasetBatchUpdateUpdateRecordWithDefaults instantiates a new LLMObsDatasetBatchUpdateUpdateRecord object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetBatchUpdateUpdateRecordWithDefaults() *LLMObsDatasetBatchUpdateUpdateRecord {
	this := LLMObsDatasetBatchUpdateUpdateRecord{}
	return &this
}

// GetExpectedOutput returns the ExpectedOutput field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsDatasetBatchUpdateUpdateRecord) GetExpectedOutput() AnyValue {
	if o == nil || o.ExpectedOutput.Get() == nil {
		var ret AnyValue
		return ret
	}
	return *o.ExpectedOutput.Get()
}

// GetExpectedOutputOk returns a tuple with the ExpectedOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) GetExpectedOutputOk() (*AnyValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpectedOutput.Get(), o.ExpectedOutput.IsSet()
}

// HasExpectedOutput returns a boolean if a field has been set.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) HasExpectedOutput() bool {
	return o != nil && o.ExpectedOutput.IsSet()
}

// SetExpectedOutput gets a reference to the given NullableAnyValue and assigns it to the ExpectedOutput field.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) SetExpectedOutput(v AnyValue) {
	o.ExpectedOutput.Set(&v)
}

// SetExpectedOutputNil sets the value for ExpectedOutput to be an explicit nil.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) SetExpectedOutputNil() {
	o.ExpectedOutput.Set(nil)
}

// UnsetExpectedOutput ensures that no value is present for ExpectedOutput, not even an explicit nil.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) UnsetExpectedOutput() {
	o.ExpectedOutput.Unset()
}

// GetId returns the Id field value.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) SetId(v string) {
	o.Id = v
}

// GetInput returns the Input field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsDatasetBatchUpdateUpdateRecord) GetInput() AnyValue {
	if o == nil || o.Input.Get() == nil {
		var ret AnyValue
		return ret
	}
	return *o.Input.Get()
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) GetInputOk() (*AnyValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Input.Get(), o.Input.IsSet()
}

// HasInput returns a boolean if a field has been set.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) HasInput() bool {
	return o != nil && o.Input.IsSet()
}

// SetInput gets a reference to the given NullableAnyValue and assigns it to the Input field.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) SetInput(v AnyValue) {
	o.Input.Set(&v)
}

// SetInputNil sets the value for Input to be an explicit nil.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) SetInputNil() {
	o.Input.Set(nil)
}

// UnsetInput ensures that no value is present for Input, not even an explicit nil.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) UnsetInput() {
	o.Input.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetTagOperations returns the TagOperations field value if set, zero value otherwise.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) GetTagOperations() LLMObsDatasetRecordTagOperations {
	if o == nil || o.TagOperations == nil {
		var ret LLMObsDatasetRecordTagOperations
		return ret
	}
	return *o.TagOperations
}

// GetTagOperationsOk returns a tuple with the TagOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) GetTagOperationsOk() (*LLMObsDatasetRecordTagOperations, bool) {
	if o == nil || o.TagOperations == nil {
		return nil, false
	}
	return o.TagOperations, true
}

// HasTagOperations returns a boolean if a field has been set.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) HasTagOperations() bool {
	return o != nil && o.TagOperations != nil
}

// SetTagOperations gets a reference to the given LLMObsDatasetRecordTagOperations and assigns it to the TagOperations field.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) SetTagOperations(v LLMObsDatasetRecordTagOperations) {
	o.TagOperations = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetBatchUpdateUpdateRecord) MarshalJSON() ([]byte, error) {
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
	if o.TagOperations != nil {
		toSerialize["tag_operations"] = o.TagOperations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDatasetBatchUpdateUpdateRecord) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExpectedOutput NullableAnyValue                  `json:"expected_output,omitempty"`
		Id             *string                           `json:"id"`
		Input          NullableAnyValue                  `json:"input,omitempty"`
		Metadata       map[string]interface{}            `json:"metadata,omitempty"`
		TagOperations  *LLMObsDatasetRecordTagOperations `json:"tag_operations,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"expected_output", "id", "input", "metadata", "tag_operations"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ExpectedOutput = all.ExpectedOutput
	o.Id = *all.Id
	o.Input = all.Input
	o.Metadata = all.Metadata
	if all.TagOperations != nil && all.TagOperations.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TagOperations = all.TagOperations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
