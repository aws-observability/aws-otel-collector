// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsScanOptionsAttributes Attributes for the AWS scan options.
type AwsScanOptionsAttributes struct {
	// Indicates if scanning of Lambda functions is enabled.
	Lambda *bool `json:"lambda,omitempty"`
	// Indicates if scanning for sensitive data is enabled.
	SensitiveData *bool `json:"sensitive_data,omitempty"`
	// Indicates if scanning for vulnerabilities in containers is enabled.
	VulnContainersOs *bool `json:"vuln_containers_os,omitempty"`
	// Indicates if scanning for vulnerabilities in hosts is enabled.
	VulnHostOs *bool `json:"vuln_host_os,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAwsScanOptionsAttributes instantiates a new AwsScanOptionsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAwsScanOptionsAttributes() *AwsScanOptionsAttributes {
	this := AwsScanOptionsAttributes{}
	return &this
}

// NewAwsScanOptionsAttributesWithDefaults instantiates a new AwsScanOptionsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAwsScanOptionsAttributesWithDefaults() *AwsScanOptionsAttributes {
	this := AwsScanOptionsAttributes{}
	return &this
}

// GetLambda returns the Lambda field value if set, zero value otherwise.
func (o *AwsScanOptionsAttributes) GetLambda() bool {
	if o == nil || o.Lambda == nil {
		var ret bool
		return ret
	}
	return *o.Lambda
}

// GetLambdaOk returns a tuple with the Lambda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsScanOptionsAttributes) GetLambdaOk() (*bool, bool) {
	if o == nil || o.Lambda == nil {
		return nil, false
	}
	return o.Lambda, true
}

// HasLambda returns a boolean if a field has been set.
func (o *AwsScanOptionsAttributes) HasLambda() bool {
	return o != nil && o.Lambda != nil
}

// SetLambda gets a reference to the given bool and assigns it to the Lambda field.
func (o *AwsScanOptionsAttributes) SetLambda(v bool) {
	o.Lambda = &v
}

// GetSensitiveData returns the SensitiveData field value if set, zero value otherwise.
func (o *AwsScanOptionsAttributes) GetSensitiveData() bool {
	if o == nil || o.SensitiveData == nil {
		var ret bool
		return ret
	}
	return *o.SensitiveData
}

// GetSensitiveDataOk returns a tuple with the SensitiveData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsScanOptionsAttributes) GetSensitiveDataOk() (*bool, bool) {
	if o == nil || o.SensitiveData == nil {
		return nil, false
	}
	return o.SensitiveData, true
}

// HasSensitiveData returns a boolean if a field has been set.
func (o *AwsScanOptionsAttributes) HasSensitiveData() bool {
	return o != nil && o.SensitiveData != nil
}

// SetSensitiveData gets a reference to the given bool and assigns it to the SensitiveData field.
func (o *AwsScanOptionsAttributes) SetSensitiveData(v bool) {
	o.SensitiveData = &v
}

// GetVulnContainersOs returns the VulnContainersOs field value if set, zero value otherwise.
func (o *AwsScanOptionsAttributes) GetVulnContainersOs() bool {
	if o == nil || o.VulnContainersOs == nil {
		var ret bool
		return ret
	}
	return *o.VulnContainersOs
}

// GetVulnContainersOsOk returns a tuple with the VulnContainersOs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsScanOptionsAttributes) GetVulnContainersOsOk() (*bool, bool) {
	if o == nil || o.VulnContainersOs == nil {
		return nil, false
	}
	return o.VulnContainersOs, true
}

// HasVulnContainersOs returns a boolean if a field has been set.
func (o *AwsScanOptionsAttributes) HasVulnContainersOs() bool {
	return o != nil && o.VulnContainersOs != nil
}

// SetVulnContainersOs gets a reference to the given bool and assigns it to the VulnContainersOs field.
func (o *AwsScanOptionsAttributes) SetVulnContainersOs(v bool) {
	o.VulnContainersOs = &v
}

// GetVulnHostOs returns the VulnHostOs field value if set, zero value otherwise.
func (o *AwsScanOptionsAttributes) GetVulnHostOs() bool {
	if o == nil || o.VulnHostOs == nil {
		var ret bool
		return ret
	}
	return *o.VulnHostOs
}

// GetVulnHostOsOk returns a tuple with the VulnHostOs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsScanOptionsAttributes) GetVulnHostOsOk() (*bool, bool) {
	if o == nil || o.VulnHostOs == nil {
		return nil, false
	}
	return o.VulnHostOs, true
}

// HasVulnHostOs returns a boolean if a field has been set.
func (o *AwsScanOptionsAttributes) HasVulnHostOs() bool {
	return o != nil && o.VulnHostOs != nil
}

// SetVulnHostOs gets a reference to the given bool and assigns it to the VulnHostOs field.
func (o *AwsScanOptionsAttributes) SetVulnHostOs(v bool) {
	o.VulnHostOs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AwsScanOptionsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Lambda != nil {
		toSerialize["lambda"] = o.Lambda
	}
	if o.SensitiveData != nil {
		toSerialize["sensitive_data"] = o.SensitiveData
	}
	if o.VulnContainersOs != nil {
		toSerialize["vuln_containers_os"] = o.VulnContainersOs
	}
	if o.VulnHostOs != nil {
		toSerialize["vuln_host_os"] = o.VulnHostOs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AwsScanOptionsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Lambda           *bool `json:"lambda,omitempty"`
		SensitiveData    *bool `json:"sensitive_data,omitempty"`
		VulnContainersOs *bool `json:"vuln_containers_os,omitempty"`
		VulnHostOs       *bool `json:"vuln_host_os,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"lambda", "sensitive_data", "vuln_containers_os", "vuln_host_os"})
	} else {
		return err
	}
	o.Lambda = all.Lambda
	o.SensitiveData = all.SensitiveData
	o.VulnContainersOs = all.VulnContainersOs
	o.VulnHostOs = all.VulnHostOs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
