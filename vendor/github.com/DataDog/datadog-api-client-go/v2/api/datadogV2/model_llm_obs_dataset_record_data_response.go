// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetRecordDataResponse A single LLM Observability dataset record.
type LLMObsDatasetRecordDataResponse struct {
	// Timestamp when the record was created.
	CreatedAt time.Time `json:"created_at"`
	// Identifier of the dataset this record belongs to.
	DatasetId string `json:"dataset_id"`
	// Represents any valid JSON value.
	ExpectedOutput NullableAnyValue `json:"expected_output"`
	// Unique identifier of the record.
	Id string `json:"id"`
	// Represents any valid JSON value.
	Input NullableAnyValue `json:"input"`
	// Arbitrary metadata associated with the record.
	Metadata map[string]interface{} `json:"metadata"`
	// Timestamp when the record was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetRecordDataResponse instantiates a new LLMObsDatasetRecordDataResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetRecordDataResponse(createdAt time.Time, datasetId string, expectedOutput NullableAnyValue, id string, input NullableAnyValue, metadata map[string]interface{}, updatedAt time.Time) *LLMObsDatasetRecordDataResponse {
	this := LLMObsDatasetRecordDataResponse{}
	this.CreatedAt = createdAt
	this.DatasetId = datasetId
	this.ExpectedOutput = expectedOutput
	this.Id = id
	this.Input = input
	this.Metadata = metadata
	this.UpdatedAt = updatedAt
	return &this
}

// NewLLMObsDatasetRecordDataResponseWithDefaults instantiates a new LLMObsDatasetRecordDataResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetRecordDataResponseWithDefaults() *LLMObsDatasetRecordDataResponse {
	this := LLMObsDatasetRecordDataResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsDatasetRecordDataResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordDataResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsDatasetRecordDataResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDatasetId returns the DatasetId field value.
func (o *LLMObsDatasetRecordDataResponse) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordDataResponse) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value.
func (o *LLMObsDatasetRecordDataResponse) SetDatasetId(v string) {
	o.DatasetId = v
}

// GetExpectedOutput returns the ExpectedOutput field value.
// If the value is explicit nil, the zero value for AnyValue will be returned.
func (o *LLMObsDatasetRecordDataResponse) GetExpectedOutput() AnyValue {
	if o == nil || o.ExpectedOutput.Get() == nil {
		var ret AnyValue
		return ret
	}
	return *o.ExpectedOutput.Get()
}

// GetExpectedOutputOk returns a tuple with the ExpectedOutput field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsDatasetRecordDataResponse) GetExpectedOutputOk() (*AnyValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpectedOutput.Get(), o.ExpectedOutput.IsSet()
}

// SetExpectedOutput sets field value.
func (o *LLMObsDatasetRecordDataResponse) SetExpectedOutput(v AnyValue) {
	o.ExpectedOutput.Set(&v)
}

// GetId returns the Id field value.
func (o *LLMObsDatasetRecordDataResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordDataResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsDatasetRecordDataResponse) SetId(v string) {
	o.Id = v
}

// GetInput returns the Input field value.
// If the value is explicit nil, the zero value for AnyValue will be returned.
func (o *LLMObsDatasetRecordDataResponse) GetInput() AnyValue {
	if o == nil || o.Input.Get() == nil {
		var ret AnyValue
		return ret
	}
	return *o.Input.Get()
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsDatasetRecordDataResponse) GetInputOk() (*AnyValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Input.Get(), o.Input.IsSet()
}

// SetInput sets field value.
func (o *LLMObsDatasetRecordDataResponse) SetInput(v AnyValue) {
	o.Input.Set(&v)
}

// GetMetadata returns the Metadata field value.
// If the value is explicit nil, the zero value for map[string]interface{} will be returned.
func (o *LLMObsDatasetRecordDataResponse) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsDatasetRecordDataResponse) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value.
func (o *LLMObsDatasetRecordDataResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *LLMObsDatasetRecordDataResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordDataResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *LLMObsDatasetRecordDataResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetRecordDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["dataset_id"] = o.DatasetId
	toSerialize["expected_output"] = o.ExpectedOutput.Get()
	toSerialize["id"] = o.Id
	toSerialize["input"] = o.Input.Get()
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDatasetRecordDataResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt      *time.Time             `json:"created_at"`
		DatasetId      *string                `json:"dataset_id"`
		ExpectedOutput NullableAnyValue       `json:"expected_output"`
		Id             *string                `json:"id"`
		Input          NullableAnyValue       `json:"input"`
		Metadata       map[string]interface{} `json:"metadata"`
		UpdatedAt      *time.Time             `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.DatasetId == nil {
		return fmt.Errorf("required field dataset_id missing")
	}
	if !all.ExpectedOutput.IsSet() {
		return fmt.Errorf("required field expected_output missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if !all.Input.IsSet() {
		return fmt.Errorf("required field input missing")
	}
	if all.Metadata == nil {
		return fmt.Errorf("required field metadata missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "dataset_id", "expected_output", "id", "input", "metadata", "updated_at"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.DatasetId = *all.DatasetId
	o.ExpectedOutput = all.ExpectedOutput
	o.Id = *all.Id
	o.Input = all.Input
	o.Metadata = all.Metadata
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
