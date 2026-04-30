// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDeleteDatasetRecordsDataAttributesRequest Attributes for deleting records from an LLM Observability dataset.
type LLMObsDeleteDatasetRecordsDataAttributesRequest struct {
	// List of record IDs to delete.
	RecordIds []string `json:"record_ids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDeleteDatasetRecordsDataAttributesRequest instantiates a new LLMObsDeleteDatasetRecordsDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDeleteDatasetRecordsDataAttributesRequest(recordIds []string) *LLMObsDeleteDatasetRecordsDataAttributesRequest {
	this := LLMObsDeleteDatasetRecordsDataAttributesRequest{}
	this.RecordIds = recordIds
	return &this
}

// NewLLMObsDeleteDatasetRecordsDataAttributesRequestWithDefaults instantiates a new LLMObsDeleteDatasetRecordsDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDeleteDatasetRecordsDataAttributesRequestWithDefaults() *LLMObsDeleteDatasetRecordsDataAttributesRequest {
	this := LLMObsDeleteDatasetRecordsDataAttributesRequest{}
	return &this
}

// GetRecordIds returns the RecordIds field value.
func (o *LLMObsDeleteDatasetRecordsDataAttributesRequest) GetRecordIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RecordIds
}

// GetRecordIdsOk returns a tuple with the RecordIds field value
// and a boolean to check if the value has been set.
func (o *LLMObsDeleteDatasetRecordsDataAttributesRequest) GetRecordIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordIds, true
}

// SetRecordIds sets field value.
func (o *LLMObsDeleteDatasetRecordsDataAttributesRequest) SetRecordIds(v []string) {
	o.RecordIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDeleteDatasetRecordsDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["record_ids"] = o.RecordIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDeleteDatasetRecordsDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RecordIds *[]string `json:"record_ids"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RecordIds == nil {
		return fmt.Errorf("required field record_ids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"record_ids"})
	} else {
		return err
	}
	o.RecordIds = *all.RecordIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
