// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetVersionHistoryAttributes The attributes of a dataset version history response.
type SecurityMonitoringDatasetVersionHistoryAttributes struct {
	// The total number of versions available for this dataset.
	Count int64 `json:"count"`
	// A map from version number (as a string) to the dataset state at that version.
	Data map[string]SecurityMonitoringDatasetVersionEntry `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetVersionHistoryAttributes instantiates a new SecurityMonitoringDatasetVersionHistoryAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetVersionHistoryAttributes(count int64, data map[string]SecurityMonitoringDatasetVersionEntry) *SecurityMonitoringDatasetVersionHistoryAttributes {
	this := SecurityMonitoringDatasetVersionHistoryAttributes{}
	this.Count = count
	this.Data = data
	return &this
}

// NewSecurityMonitoringDatasetVersionHistoryAttributesWithDefaults instantiates a new SecurityMonitoringDatasetVersionHistoryAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetVersionHistoryAttributesWithDefaults() *SecurityMonitoringDatasetVersionHistoryAttributes {
	this := SecurityMonitoringDatasetVersionHistoryAttributes{}
	return &this
}

// GetCount returns the Count field value.
func (o *SecurityMonitoringDatasetVersionHistoryAttributes) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetVersionHistoryAttributes) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *SecurityMonitoringDatasetVersionHistoryAttributes) SetCount(v int64) {
	o.Count = v
}

// GetData returns the Data field value.
func (o *SecurityMonitoringDatasetVersionHistoryAttributes) GetData() map[string]SecurityMonitoringDatasetVersionEntry {
	if o == nil {
		var ret map[string]SecurityMonitoringDatasetVersionEntry
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetVersionHistoryAttributes) GetDataOk() (*map[string]SecurityMonitoringDatasetVersionEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *SecurityMonitoringDatasetVersionHistoryAttributes) SetData(v map[string]SecurityMonitoringDatasetVersionEntry) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetVersionHistoryAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["count"] = o.Count
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringDatasetVersionHistoryAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count *int64                                            `json:"count"`
		Data  *map[string]SecurityMonitoringDatasetVersionEntry `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "data"})
	} else {
		return err
	}
	o.Count = *all.Count
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
