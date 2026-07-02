// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetBatchUpdateDataAttributesRequest Attributes for batch-updating records in an LLM Observability dataset.
type LLMObsDatasetBatchUpdateDataAttributesRequest struct {
	// Whether to create a new dataset version when applying the batch update. Defaults to `true`.
	CreateNewVersion *bool `json:"create_new_version,omitempty"`
	// Record IDs to delete.
	DeleteRecords []string `json:"delete_records,omitempty"`
	// Records to insert.
	InsertRecords []LLMObsDatasetBatchUpdateInsertRecord `json:"insert_records,omitempty"`
	// List of tag strings.
	Tags []string `json:"tags,omitempty"`
	// Records to update by ID.
	UpdateRecords []LLMObsDatasetBatchUpdateUpdateRecord `json:"update_records,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetBatchUpdateDataAttributesRequest instantiates a new LLMObsDatasetBatchUpdateDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetBatchUpdateDataAttributesRequest() *LLMObsDatasetBatchUpdateDataAttributesRequest {
	this := LLMObsDatasetBatchUpdateDataAttributesRequest{}
	return &this
}

// NewLLMObsDatasetBatchUpdateDataAttributesRequestWithDefaults instantiates a new LLMObsDatasetBatchUpdateDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetBatchUpdateDataAttributesRequestWithDefaults() *LLMObsDatasetBatchUpdateDataAttributesRequest {
	this := LLMObsDatasetBatchUpdateDataAttributesRequest{}
	return &this
}

// GetCreateNewVersion returns the CreateNewVersion field value if set, zero value otherwise.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) GetCreateNewVersion() bool {
	if o == nil || o.CreateNewVersion == nil {
		var ret bool
		return ret
	}
	return *o.CreateNewVersion
}

// GetCreateNewVersionOk returns a tuple with the CreateNewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) GetCreateNewVersionOk() (*bool, bool) {
	if o == nil || o.CreateNewVersion == nil {
		return nil, false
	}
	return o.CreateNewVersion, true
}

// HasCreateNewVersion returns a boolean if a field has been set.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) HasCreateNewVersion() bool {
	return o != nil && o.CreateNewVersion != nil
}

// SetCreateNewVersion gets a reference to the given bool and assigns it to the CreateNewVersion field.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) SetCreateNewVersion(v bool) {
	o.CreateNewVersion = &v
}

// GetDeleteRecords returns the DeleteRecords field value if set, zero value otherwise.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) GetDeleteRecords() []string {
	if o == nil || o.DeleteRecords == nil {
		var ret []string
		return ret
	}
	return o.DeleteRecords
}

// GetDeleteRecordsOk returns a tuple with the DeleteRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) GetDeleteRecordsOk() (*[]string, bool) {
	if o == nil || o.DeleteRecords == nil {
		return nil, false
	}
	return &o.DeleteRecords, true
}

// HasDeleteRecords returns a boolean if a field has been set.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) HasDeleteRecords() bool {
	return o != nil && o.DeleteRecords != nil
}

// SetDeleteRecords gets a reference to the given []string and assigns it to the DeleteRecords field.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) SetDeleteRecords(v []string) {
	o.DeleteRecords = v
}

// GetInsertRecords returns the InsertRecords field value if set, zero value otherwise.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) GetInsertRecords() []LLMObsDatasetBatchUpdateInsertRecord {
	if o == nil || o.InsertRecords == nil {
		var ret []LLMObsDatasetBatchUpdateInsertRecord
		return ret
	}
	return o.InsertRecords
}

// GetInsertRecordsOk returns a tuple with the InsertRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) GetInsertRecordsOk() (*[]LLMObsDatasetBatchUpdateInsertRecord, bool) {
	if o == nil || o.InsertRecords == nil {
		return nil, false
	}
	return &o.InsertRecords, true
}

// HasInsertRecords returns a boolean if a field has been set.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) HasInsertRecords() bool {
	return o != nil && o.InsertRecords != nil
}

// SetInsertRecords gets a reference to the given []LLMObsDatasetBatchUpdateInsertRecord and assigns it to the InsertRecords field.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) SetInsertRecords(v []LLMObsDatasetBatchUpdateInsertRecord) {
	o.InsertRecords = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) SetTags(v []string) {
	o.Tags = v
}

// GetUpdateRecords returns the UpdateRecords field value if set, zero value otherwise.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) GetUpdateRecords() []LLMObsDatasetBatchUpdateUpdateRecord {
	if o == nil || o.UpdateRecords == nil {
		var ret []LLMObsDatasetBatchUpdateUpdateRecord
		return ret
	}
	return o.UpdateRecords
}

// GetUpdateRecordsOk returns a tuple with the UpdateRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) GetUpdateRecordsOk() (*[]LLMObsDatasetBatchUpdateUpdateRecord, bool) {
	if o == nil || o.UpdateRecords == nil {
		return nil, false
	}
	return &o.UpdateRecords, true
}

// HasUpdateRecords returns a boolean if a field has been set.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) HasUpdateRecords() bool {
	return o != nil && o.UpdateRecords != nil
}

// SetUpdateRecords gets a reference to the given []LLMObsDatasetBatchUpdateUpdateRecord and assigns it to the UpdateRecords field.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) SetUpdateRecords(v []LLMObsDatasetBatchUpdateUpdateRecord) {
	o.UpdateRecords = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetBatchUpdateDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreateNewVersion != nil {
		toSerialize["create_new_version"] = o.CreateNewVersion
	}
	if o.DeleteRecords != nil {
		toSerialize["delete_records"] = o.DeleteRecords
	}
	if o.InsertRecords != nil {
		toSerialize["insert_records"] = o.InsertRecords
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdateRecords != nil {
		toSerialize["update_records"] = o.UpdateRecords
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDatasetBatchUpdateDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreateNewVersion *bool                                  `json:"create_new_version,omitempty"`
		DeleteRecords    []string                               `json:"delete_records,omitempty"`
		InsertRecords    []LLMObsDatasetBatchUpdateInsertRecord `json:"insert_records,omitempty"`
		Tags             []string                               `json:"tags,omitempty"`
		UpdateRecords    []LLMObsDatasetBatchUpdateUpdateRecord `json:"update_records,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"create_new_version", "delete_records", "insert_records", "tags", "update_records"})
	} else {
		return err
	}
	o.CreateNewVersion = all.CreateNewVersion
	o.DeleteRecords = all.DeleteRecords
	o.InsertRecords = all.InsertRecords
	o.Tags = all.Tags
	o.UpdateRecords = all.UpdateRecords

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
