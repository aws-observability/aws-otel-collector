// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetBatchUpdateInsertRecord A record to insert as part of a batch update on an LLM Observability dataset.
type LLMObsDatasetBatchUpdateInsertRecord struct {
	// Represents any valid JSON value.
	ExpectedOutput NullableAnyValue `json:"expected_output,omitempty"`
	// Optional user-provided identifier for the record. If omitted, the server generates an identifier.
	Id *string `json:"id,omitempty"`
	// Represents any valid JSON value.
	Input NullableAnyValue `json:"input"`
	// Arbitrary metadata associated with the record.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Explicit tag operations for updating records. Operations are applied in order, Remove then Add then Set. `set` is the final override; if specified, the result of `remove` and `add` is discarded.
	TagOperations *LLMObsDatasetRecordTagOperations `json:"tag_operations,omitempty"`
	// List of tag strings.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetBatchUpdateInsertRecord instantiates a new LLMObsDatasetBatchUpdateInsertRecord object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetBatchUpdateInsertRecord(input NullableAnyValue) *LLMObsDatasetBatchUpdateInsertRecord {
	this := LLMObsDatasetBatchUpdateInsertRecord{}
	this.Input = input
	return &this
}

// NewLLMObsDatasetBatchUpdateInsertRecordWithDefaults instantiates a new LLMObsDatasetBatchUpdateInsertRecord object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetBatchUpdateInsertRecordWithDefaults() *LLMObsDatasetBatchUpdateInsertRecord {
	this := LLMObsDatasetBatchUpdateInsertRecord{}
	return &this
}

// GetExpectedOutput returns the ExpectedOutput field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsDatasetBatchUpdateInsertRecord) GetExpectedOutput() AnyValue {
	if o == nil || o.ExpectedOutput.Get() == nil {
		var ret AnyValue
		return ret
	}
	return *o.ExpectedOutput.Get()
}

// GetExpectedOutputOk returns a tuple with the ExpectedOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsDatasetBatchUpdateInsertRecord) GetExpectedOutputOk() (*AnyValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpectedOutput.Get(), o.ExpectedOutput.IsSet()
}

// HasExpectedOutput returns a boolean if a field has been set.
func (o *LLMObsDatasetBatchUpdateInsertRecord) HasExpectedOutput() bool {
	return o != nil && o.ExpectedOutput.IsSet()
}

// SetExpectedOutput gets a reference to the given NullableAnyValue and assigns it to the ExpectedOutput field.
func (o *LLMObsDatasetBatchUpdateInsertRecord) SetExpectedOutput(v AnyValue) {
	o.ExpectedOutput.Set(&v)
}

// SetExpectedOutputNil sets the value for ExpectedOutput to be an explicit nil.
func (o *LLMObsDatasetBatchUpdateInsertRecord) SetExpectedOutputNil() {
	o.ExpectedOutput.Set(nil)
}

// UnsetExpectedOutput ensures that no value is present for ExpectedOutput, not even an explicit nil.
func (o *LLMObsDatasetBatchUpdateInsertRecord) UnsetExpectedOutput() {
	o.ExpectedOutput.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LLMObsDatasetBatchUpdateInsertRecord) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetBatchUpdateInsertRecord) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LLMObsDatasetBatchUpdateInsertRecord) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LLMObsDatasetBatchUpdateInsertRecord) SetId(v string) {
	o.Id = &v
}

// GetInput returns the Input field value.
// If the value is explicit nil, the zero value for AnyValue will be returned.
func (o *LLMObsDatasetBatchUpdateInsertRecord) GetInput() AnyValue {
	if o == nil || o.Input.Get() == nil {
		var ret AnyValue
		return ret
	}
	return *o.Input.Get()
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsDatasetBatchUpdateInsertRecord) GetInputOk() (*AnyValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Input.Get(), o.Input.IsSet()
}

// SetInput sets field value.
func (o *LLMObsDatasetBatchUpdateInsertRecord) SetInput(v AnyValue) {
	o.Input.Set(&v)
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LLMObsDatasetBatchUpdateInsertRecord) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetBatchUpdateInsertRecord) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LLMObsDatasetBatchUpdateInsertRecord) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LLMObsDatasetBatchUpdateInsertRecord) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetTagOperations returns the TagOperations field value if set, zero value otherwise.
func (o *LLMObsDatasetBatchUpdateInsertRecord) GetTagOperations() LLMObsDatasetRecordTagOperations {
	if o == nil || o.TagOperations == nil {
		var ret LLMObsDatasetRecordTagOperations
		return ret
	}
	return *o.TagOperations
}

// GetTagOperationsOk returns a tuple with the TagOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetBatchUpdateInsertRecord) GetTagOperationsOk() (*LLMObsDatasetRecordTagOperations, bool) {
	if o == nil || o.TagOperations == nil {
		return nil, false
	}
	return o.TagOperations, true
}

// HasTagOperations returns a boolean if a field has been set.
func (o *LLMObsDatasetBatchUpdateInsertRecord) HasTagOperations() bool {
	return o != nil && o.TagOperations != nil
}

// SetTagOperations gets a reference to the given LLMObsDatasetRecordTagOperations and assigns it to the TagOperations field.
func (o *LLMObsDatasetBatchUpdateInsertRecord) SetTagOperations(v LLMObsDatasetRecordTagOperations) {
	o.TagOperations = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LLMObsDatasetBatchUpdateInsertRecord) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetBatchUpdateInsertRecord) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LLMObsDatasetBatchUpdateInsertRecord) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LLMObsDatasetBatchUpdateInsertRecord) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetBatchUpdateInsertRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExpectedOutput.IsSet() {
		toSerialize["expected_output"] = o.ExpectedOutput.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["input"] = o.Input.Get()
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.TagOperations != nil {
		toSerialize["tag_operations"] = o.TagOperations
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDatasetBatchUpdateInsertRecord) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExpectedOutput NullableAnyValue                  `json:"expected_output,omitempty"`
		Id             *string                           `json:"id,omitempty"`
		Input          NullableAnyValue                  `json:"input"`
		Metadata       map[string]interface{}            `json:"metadata,omitempty"`
		TagOperations  *LLMObsDatasetRecordTagOperations `json:"tag_operations,omitempty"`
		Tags           []string                          `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Input.IsSet() {
		return fmt.Errorf("required field input missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"expected_output", "id", "input", "metadata", "tag_operations", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ExpectedOutput = all.ExpectedOutput
	o.Id = all.Id
	o.Input = all.Input
	o.Metadata = all.Metadata
	if all.TagOperations != nil && all.TagOperations.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TagOperations = all.TagOperations
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
