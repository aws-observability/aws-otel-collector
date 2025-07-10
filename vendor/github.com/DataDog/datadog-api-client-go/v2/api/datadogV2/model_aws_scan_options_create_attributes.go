// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsScanOptionsCreateAttributes Attributes for the AWS scan options to create.
type AwsScanOptionsCreateAttributes struct {
	// Indicates if scanning of Lambda functions is enabled.
	Lambda bool `json:"lambda"`
	// Indicates if scanning for sensitive data is enabled.
	SensitiveData bool `json:"sensitive_data"`
	// Indicates if scanning for vulnerabilities in containers is enabled.
	VulnContainersOs bool `json:"vuln_containers_os"`
	// Indicates if scanning for vulnerabilities in hosts is enabled.
	VulnHostOs bool `json:"vuln_host_os"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAwsScanOptionsCreateAttributes instantiates a new AwsScanOptionsCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAwsScanOptionsCreateAttributes(lambda bool, sensitiveData bool, vulnContainersOs bool, vulnHostOs bool) *AwsScanOptionsCreateAttributes {
	this := AwsScanOptionsCreateAttributes{}
	this.Lambda = lambda
	this.SensitiveData = sensitiveData
	this.VulnContainersOs = vulnContainersOs
	this.VulnHostOs = vulnHostOs
	return &this
}

// NewAwsScanOptionsCreateAttributesWithDefaults instantiates a new AwsScanOptionsCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAwsScanOptionsCreateAttributesWithDefaults() *AwsScanOptionsCreateAttributes {
	this := AwsScanOptionsCreateAttributes{}
	return &this
}

// GetLambda returns the Lambda field value.
func (o *AwsScanOptionsCreateAttributes) GetLambda() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Lambda
}

// GetLambdaOk returns a tuple with the Lambda field value
// and a boolean to check if the value has been set.
func (o *AwsScanOptionsCreateAttributes) GetLambdaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lambda, true
}

// SetLambda sets field value.
func (o *AwsScanOptionsCreateAttributes) SetLambda(v bool) {
	o.Lambda = v
}

// GetSensitiveData returns the SensitiveData field value.
func (o *AwsScanOptionsCreateAttributes) GetSensitiveData() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SensitiveData
}

// GetSensitiveDataOk returns a tuple with the SensitiveData field value
// and a boolean to check if the value has been set.
func (o *AwsScanOptionsCreateAttributes) GetSensitiveDataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SensitiveData, true
}

// SetSensitiveData sets field value.
func (o *AwsScanOptionsCreateAttributes) SetSensitiveData(v bool) {
	o.SensitiveData = v
}

// GetVulnContainersOs returns the VulnContainersOs field value.
func (o *AwsScanOptionsCreateAttributes) GetVulnContainersOs() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.VulnContainersOs
}

// GetVulnContainersOsOk returns a tuple with the VulnContainersOs field value
// and a boolean to check if the value has been set.
func (o *AwsScanOptionsCreateAttributes) GetVulnContainersOsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnContainersOs, true
}

// SetVulnContainersOs sets field value.
func (o *AwsScanOptionsCreateAttributes) SetVulnContainersOs(v bool) {
	o.VulnContainersOs = v
}

// GetVulnHostOs returns the VulnHostOs field value.
func (o *AwsScanOptionsCreateAttributes) GetVulnHostOs() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.VulnHostOs
}

// GetVulnHostOsOk returns a tuple with the VulnHostOs field value
// and a boolean to check if the value has been set.
func (o *AwsScanOptionsCreateAttributes) GetVulnHostOsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnHostOs, true
}

// SetVulnHostOs sets field value.
func (o *AwsScanOptionsCreateAttributes) SetVulnHostOs(v bool) {
	o.VulnHostOs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AwsScanOptionsCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["lambda"] = o.Lambda
	toSerialize["sensitive_data"] = o.SensitiveData
	toSerialize["vuln_containers_os"] = o.VulnContainersOs
	toSerialize["vuln_host_os"] = o.VulnHostOs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AwsScanOptionsCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Lambda           *bool `json:"lambda"`
		SensitiveData    *bool `json:"sensitive_data"`
		VulnContainersOs *bool `json:"vuln_containers_os"`
		VulnHostOs       *bool `json:"vuln_host_os"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Lambda == nil {
		return fmt.Errorf("required field lambda missing")
	}
	if all.SensitiveData == nil {
		return fmt.Errorf("required field sensitive_data missing")
	}
	if all.VulnContainersOs == nil {
		return fmt.Errorf("required field vuln_containers_os missing")
	}
	if all.VulnHostOs == nil {
		return fmt.Errorf("required field vuln_host_os missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"lambda", "sensitive_data", "vuln_containers_os", "vuln_host_os"})
	} else {
		return err
	}
	o.Lambda = *all.Lambda
	o.SensitiveData = *all.SensitiveData
	o.VulnContainersOs = *all.VulnContainersOs
	o.VulnHostOs = *all.VulnHostOs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
