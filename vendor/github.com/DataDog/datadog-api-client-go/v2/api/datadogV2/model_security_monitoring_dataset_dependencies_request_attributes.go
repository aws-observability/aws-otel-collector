// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetDependenciesRequestAttributes The attributes of a dataset dependencies request.
type SecurityMonitoringDatasetDependenciesRequestAttributes struct {
	// The list of dataset UUIDs to query dependencies for. Must contain between 1 and 100 items.
	DatasetIds []string `json:"datasetIds"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetDependenciesRequestAttributes instantiates a new SecurityMonitoringDatasetDependenciesRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetDependenciesRequestAttributes(datasetIds []string) *SecurityMonitoringDatasetDependenciesRequestAttributes {
	this := SecurityMonitoringDatasetDependenciesRequestAttributes{}
	this.DatasetIds = datasetIds
	return &this
}

// NewSecurityMonitoringDatasetDependenciesRequestAttributesWithDefaults instantiates a new SecurityMonitoringDatasetDependenciesRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetDependenciesRequestAttributesWithDefaults() *SecurityMonitoringDatasetDependenciesRequestAttributes {
	this := SecurityMonitoringDatasetDependenciesRequestAttributes{}
	return &this
}

// GetDatasetIds returns the DatasetIds field value.
func (o *SecurityMonitoringDatasetDependenciesRequestAttributes) GetDatasetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DatasetIds
}

// GetDatasetIdsOk returns a tuple with the DatasetIds field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDependenciesRequestAttributes) GetDatasetIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetIds, true
}

// SetDatasetIds sets field value.
func (o *SecurityMonitoringDatasetDependenciesRequestAttributes) SetDatasetIds(v []string) {
	o.DatasetIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetDependenciesRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["datasetIds"] = o.DatasetIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringDatasetDependenciesRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasetIds *[]string `json:"datasetIds"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DatasetIds == nil {
		return fmt.Errorf("required field datasetIds missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"datasetIds"})
	} else {
		return err
	}
	o.DatasetIds = *all.DatasetIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
