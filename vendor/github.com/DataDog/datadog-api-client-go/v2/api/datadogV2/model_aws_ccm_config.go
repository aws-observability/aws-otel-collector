// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCcmConfig AWS Cloud Cost Management config.
type AWSCcmConfig struct {
	// List of data export configurations for Cost and Usage Reports.
	DataExportConfigs []DataExportConfig `json:"data_export_configs"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSCcmConfig instantiates a new AWSCcmConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSCcmConfig(dataExportConfigs []DataExportConfig) *AWSCcmConfig {
	this := AWSCcmConfig{}
	this.DataExportConfigs = dataExportConfigs
	return &this
}

// NewAWSCcmConfigWithDefaults instantiates a new AWSCcmConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSCcmConfigWithDefaults() *AWSCcmConfig {
	this := AWSCcmConfig{}
	return &this
}

// GetDataExportConfigs returns the DataExportConfigs field value.
func (o *AWSCcmConfig) GetDataExportConfigs() []DataExportConfig {
	if o == nil {
		var ret []DataExportConfig
		return ret
	}
	return o.DataExportConfigs
}

// GetDataExportConfigsOk returns a tuple with the DataExportConfigs field value
// and a boolean to check if the value has been set.
func (o *AWSCcmConfig) GetDataExportConfigsOk() (*[]DataExportConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataExportConfigs, true
}

// SetDataExportConfigs sets field value.
func (o *AWSCcmConfig) SetDataExportConfigs(v []DataExportConfig) {
	o.DataExportConfigs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSCcmConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_export_configs"] = o.DataExportConfigs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSCcmConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataExportConfigs *[]DataExportConfig `json:"data_export_configs"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataExportConfigs == nil {
		return fmt.Errorf("required field data_export_configs missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_export_configs"})
	} else {
		return err
	}
	o.DataExportConfigs = *all.DataExportConfigs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
