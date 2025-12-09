// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSTracesConfig AWS Traces Collection config.
type AWSTracesConfig struct {
	// AWS X-Ray services to collect traces from. Defaults to `include_only`.
	XrayServices *XRayServicesList `json:"xray_services,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSTracesConfig instantiates a new AWSTracesConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSTracesConfig() *AWSTracesConfig {
	this := AWSTracesConfig{}
	return &this
}

// NewAWSTracesConfigWithDefaults instantiates a new AWSTracesConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSTracesConfigWithDefaults() *AWSTracesConfig {
	this := AWSTracesConfig{}
	return &this
}

// GetXrayServices returns the XrayServices field value if set, zero value otherwise.
func (o *AWSTracesConfig) GetXrayServices() XRayServicesList {
	if o == nil || o.XrayServices == nil {
		var ret XRayServicesList
		return ret
	}
	return *o.XrayServices
}

// GetXrayServicesOk returns a tuple with the XrayServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSTracesConfig) GetXrayServicesOk() (*XRayServicesList, bool) {
	if o == nil || o.XrayServices == nil {
		return nil, false
	}
	return o.XrayServices, true
}

// HasXrayServices returns a boolean if a field has been set.
func (o *AWSTracesConfig) HasXrayServices() bool {
	return o != nil && o.XrayServices != nil
}

// SetXrayServices gets a reference to the given XRayServicesList and assigns it to the XrayServices field.
func (o *AWSTracesConfig) SetXrayServices(v XRayServicesList) {
	o.XrayServices = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSTracesConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.XrayServices != nil {
		toSerialize["xray_services"] = o.XrayServices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSTracesConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		XrayServices *XRayServicesList `json:"xray_services,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"xray_services"})
	} else {
		return err
	}
	o.XrayServices = all.XrayServices

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
