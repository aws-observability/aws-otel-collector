// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetVersionEntry A single entry in the version history of a dataset.
type SecurityMonitoringDatasetVersionEntry struct {
	// The list of field changes between this version of the dataset and the previous one.
	Changes []SecurityMonitoringDatasetVersionFieldChange `json:"changes"`
	// The attributes of a Cloud SIEM dataset.
	Dataset SecurityMonitoringDatasetAttributesResponse `json:"dataset"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetVersionEntry instantiates a new SecurityMonitoringDatasetVersionEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetVersionEntry(changes []SecurityMonitoringDatasetVersionFieldChange, dataset SecurityMonitoringDatasetAttributesResponse) *SecurityMonitoringDatasetVersionEntry {
	this := SecurityMonitoringDatasetVersionEntry{}
	this.Changes = changes
	this.Dataset = dataset
	return &this
}

// NewSecurityMonitoringDatasetVersionEntryWithDefaults instantiates a new SecurityMonitoringDatasetVersionEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetVersionEntryWithDefaults() *SecurityMonitoringDatasetVersionEntry {
	this := SecurityMonitoringDatasetVersionEntry{}
	return &this
}

// GetChanges returns the Changes field value.
func (o *SecurityMonitoringDatasetVersionEntry) GetChanges() []SecurityMonitoringDatasetVersionFieldChange {
	if o == nil {
		var ret []SecurityMonitoringDatasetVersionFieldChange
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetVersionEntry) GetChangesOk() (*[]SecurityMonitoringDatasetVersionFieldChange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Changes, true
}

// SetChanges sets field value.
func (o *SecurityMonitoringDatasetVersionEntry) SetChanges(v []SecurityMonitoringDatasetVersionFieldChange) {
	o.Changes = v
}

// GetDataset returns the Dataset field value.
func (o *SecurityMonitoringDatasetVersionEntry) GetDataset() SecurityMonitoringDatasetAttributesResponse {
	if o == nil {
		var ret SecurityMonitoringDatasetAttributesResponse
		return ret
	}
	return o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetVersionEntry) GetDatasetOk() (*SecurityMonitoringDatasetAttributesResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dataset, true
}

// SetDataset sets field value.
func (o *SecurityMonitoringDatasetVersionEntry) SetDataset(v SecurityMonitoringDatasetAttributesResponse) {
	o.Dataset = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetVersionEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["changes"] = o.Changes
	toSerialize["dataset"] = o.Dataset

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringDatasetVersionEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Changes *[]SecurityMonitoringDatasetVersionFieldChange `json:"changes"`
		Dataset *SecurityMonitoringDatasetAttributesResponse   `json:"dataset"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Changes == nil {
		return fmt.Errorf("required field changes missing")
	}
	if all.Dataset == nil {
		return fmt.Errorf("required field dataset missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"changes", "dataset"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Changes = *all.Changes
	if all.Dataset.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Dataset = *all.Dataset

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
