// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetRecordsDataAttributesRequest Attributes for appending records to an LLM Observability dataset.
type LLMObsDatasetRecordsDataAttributesRequest struct {
	// Whether to deduplicate records before appending. Defaults to `true`.
	Deduplicate *bool `json:"deduplicate,omitempty"`
	// List of records to append to the dataset.
	Records []LLMObsDatasetRecordItem `json:"records"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetRecordsDataAttributesRequest instantiates a new LLMObsDatasetRecordsDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetRecordsDataAttributesRequest(records []LLMObsDatasetRecordItem) *LLMObsDatasetRecordsDataAttributesRequest {
	this := LLMObsDatasetRecordsDataAttributesRequest{}
	this.Records = records
	return &this
}

// NewLLMObsDatasetRecordsDataAttributesRequestWithDefaults instantiates a new LLMObsDatasetRecordsDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetRecordsDataAttributesRequestWithDefaults() *LLMObsDatasetRecordsDataAttributesRequest {
	this := LLMObsDatasetRecordsDataAttributesRequest{}
	return &this
}

// GetDeduplicate returns the Deduplicate field value if set, zero value otherwise.
func (o *LLMObsDatasetRecordsDataAttributesRequest) GetDeduplicate() bool {
	if o == nil || o.Deduplicate == nil {
		var ret bool
		return ret
	}
	return *o.Deduplicate
}

// GetDeduplicateOk returns a tuple with the Deduplicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordsDataAttributesRequest) GetDeduplicateOk() (*bool, bool) {
	if o == nil || o.Deduplicate == nil {
		return nil, false
	}
	return o.Deduplicate, true
}

// HasDeduplicate returns a boolean if a field has been set.
func (o *LLMObsDatasetRecordsDataAttributesRequest) HasDeduplicate() bool {
	return o != nil && o.Deduplicate != nil
}

// SetDeduplicate gets a reference to the given bool and assigns it to the Deduplicate field.
func (o *LLMObsDatasetRecordsDataAttributesRequest) SetDeduplicate(v bool) {
	o.Deduplicate = &v
}

// GetRecords returns the Records field value.
func (o *LLMObsDatasetRecordsDataAttributesRequest) GetRecords() []LLMObsDatasetRecordItem {
	if o == nil {
		var ret []LLMObsDatasetRecordItem
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordsDataAttributesRequest) GetRecordsOk() (*[]LLMObsDatasetRecordItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Records, true
}

// SetRecords sets field value.
func (o *LLMObsDatasetRecordsDataAttributesRequest) SetRecords(v []LLMObsDatasetRecordItem) {
	o.Records = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetRecordsDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Deduplicate != nil {
		toSerialize["deduplicate"] = o.Deduplicate
	}
	toSerialize["records"] = o.Records

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDatasetRecordsDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Deduplicate *bool                      `json:"deduplicate,omitempty"`
		Records     *[]LLMObsDatasetRecordItem `json:"records"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Records == nil {
		return fmt.Errorf("required field records missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deduplicate", "records"})
	} else {
		return err
	}
	o.Deduplicate = all.Deduplicate
	o.Records = *all.Records

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
