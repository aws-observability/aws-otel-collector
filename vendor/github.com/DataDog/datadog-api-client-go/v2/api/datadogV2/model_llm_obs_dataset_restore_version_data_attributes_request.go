// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetRestoreVersionDataAttributesRequest Attributes for restoring an LLM Observability dataset to a previous version.
type LLMObsDatasetRestoreVersionDataAttributesRequest struct {
	// Version number of the dataset to restore. Must be between 0 and the current version of the dataset, inclusive.
	DatasetVersion int32 `json:"dataset_version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetRestoreVersionDataAttributesRequest instantiates a new LLMObsDatasetRestoreVersionDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetRestoreVersionDataAttributesRequest(datasetVersion int32) *LLMObsDatasetRestoreVersionDataAttributesRequest {
	this := LLMObsDatasetRestoreVersionDataAttributesRequest{}
	this.DatasetVersion = datasetVersion
	return &this
}

// NewLLMObsDatasetRestoreVersionDataAttributesRequestWithDefaults instantiates a new LLMObsDatasetRestoreVersionDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetRestoreVersionDataAttributesRequestWithDefaults() *LLMObsDatasetRestoreVersionDataAttributesRequest {
	this := LLMObsDatasetRestoreVersionDataAttributesRequest{}
	return &this
}

// GetDatasetVersion returns the DatasetVersion field value.
func (o *LLMObsDatasetRestoreVersionDataAttributesRequest) GetDatasetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.DatasetVersion
}

// GetDatasetVersionOk returns a tuple with the DatasetVersion field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRestoreVersionDataAttributesRequest) GetDatasetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetVersion, true
}

// SetDatasetVersion sets field value.
func (o *LLMObsDatasetRestoreVersionDataAttributesRequest) SetDatasetVersion(v int32) {
	o.DatasetVersion = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetRestoreVersionDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dataset_version"] = o.DatasetVersion

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDatasetRestoreVersionDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasetVersion *int32 `json:"dataset_version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DatasetVersion == nil {
		return fmt.Errorf("required field dataset_version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dataset_version"})
	} else {
		return err
	}
	o.DatasetVersion = *all.DatasetVersion

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
