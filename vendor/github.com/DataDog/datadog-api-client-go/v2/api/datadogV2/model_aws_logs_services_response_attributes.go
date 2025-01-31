// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSLogsServicesResponseAttributes AWS Logs Services response body
type AWSLogsServicesResponseAttributes struct {
	// List of AWS services that can send logs to Datadog
	LogsServices []string `json:"logs_services"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSLogsServicesResponseAttributes instantiates a new AWSLogsServicesResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSLogsServicesResponseAttributes(logsServices []string) *AWSLogsServicesResponseAttributes {
	this := AWSLogsServicesResponseAttributes{}
	this.LogsServices = logsServices
	return &this
}

// NewAWSLogsServicesResponseAttributesWithDefaults instantiates a new AWSLogsServicesResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSLogsServicesResponseAttributesWithDefaults() *AWSLogsServicesResponseAttributes {
	this := AWSLogsServicesResponseAttributes{}
	return &this
}

// GetLogsServices returns the LogsServices field value.
func (o *AWSLogsServicesResponseAttributes) GetLogsServices() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.LogsServices
}

// GetLogsServicesOk returns a tuple with the LogsServices field value
// and a boolean to check if the value has been set.
func (o *AWSLogsServicesResponseAttributes) GetLogsServicesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogsServices, true
}

// SetLogsServices sets field value.
func (o *AWSLogsServicesResponseAttributes) SetLogsServices(v []string) {
	o.LogsServices = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSLogsServicesResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["logs_services"] = o.LogsServices

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSLogsServicesResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LogsServices *[]string `json:"logs_services"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LogsServices == nil {
		return fmt.Errorf("required field logs_services missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"logs_services"})
	} else {
		return err
	}
	o.LogsServices = *all.LogsServices

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
