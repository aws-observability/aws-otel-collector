// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetRecordsMutationData Response containing records after a create or update operation.
type LLMObsDatasetRecordsMutationData struct {
	// List of affected dataset records.
	Records []LLMObsDatasetRecordDataResponse `json:"records"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetRecordsMutationData instantiates a new LLMObsDatasetRecordsMutationData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetRecordsMutationData(records []LLMObsDatasetRecordDataResponse) *LLMObsDatasetRecordsMutationData {
	this := LLMObsDatasetRecordsMutationData{}
	this.Records = records
	return &this
}

// NewLLMObsDatasetRecordsMutationDataWithDefaults instantiates a new LLMObsDatasetRecordsMutationData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetRecordsMutationDataWithDefaults() *LLMObsDatasetRecordsMutationData {
	this := LLMObsDatasetRecordsMutationData{}
	return &this
}

// GetRecords returns the Records field value.
func (o *LLMObsDatasetRecordsMutationData) GetRecords() []LLMObsDatasetRecordDataResponse {
	if o == nil {
		var ret []LLMObsDatasetRecordDataResponse
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordsMutationData) GetRecordsOk() (*[]LLMObsDatasetRecordDataResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Records, true
}

// SetRecords sets field value.
func (o *LLMObsDatasetRecordsMutationData) SetRecords(v []LLMObsDatasetRecordDataResponse) {
	o.Records = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetRecordsMutationData) MarshalJSON() ([]byte, error) {
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
func (o *LLMObsDatasetRecordsMutationData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Records *[]LLMObsDatasetRecordDataResponse `json:"records"`
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
