// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetRecordsUpdateDataAttributesRequest Attributes for updating records in an LLM Observability dataset.
type LLMObsDatasetRecordsUpdateDataAttributesRequest struct {
	// List of records to update.
	Records []LLMObsDatasetRecordUpdateItem `json:"records"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetRecordsUpdateDataAttributesRequest instantiates a new LLMObsDatasetRecordsUpdateDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetRecordsUpdateDataAttributesRequest(records []LLMObsDatasetRecordUpdateItem) *LLMObsDatasetRecordsUpdateDataAttributesRequest {
	this := LLMObsDatasetRecordsUpdateDataAttributesRequest{}
	this.Records = records
	return &this
}

// NewLLMObsDatasetRecordsUpdateDataAttributesRequestWithDefaults instantiates a new LLMObsDatasetRecordsUpdateDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetRecordsUpdateDataAttributesRequestWithDefaults() *LLMObsDatasetRecordsUpdateDataAttributesRequest {
	this := LLMObsDatasetRecordsUpdateDataAttributesRequest{}
	return &this
}

// GetRecords returns the Records field value.
func (o *LLMObsDatasetRecordsUpdateDataAttributesRequest) GetRecords() []LLMObsDatasetRecordUpdateItem {
	if o == nil {
		var ret []LLMObsDatasetRecordUpdateItem
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordsUpdateDataAttributesRequest) GetRecordsOk() (*[]LLMObsDatasetRecordUpdateItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Records, true
}

// SetRecords sets field value.
func (o *LLMObsDatasetRecordsUpdateDataAttributesRequest) SetRecords(v []LLMObsDatasetRecordUpdateItem) {
	o.Records = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetRecordsUpdateDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["records"] = o.Records

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDatasetRecordsUpdateDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Records *[]LLMObsDatasetRecordUpdateItem `json:"records"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Records == nil {
		return fmt.Errorf("required field records missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"records"})
	} else {
		return err
	}
	o.Records = *all.Records

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
