// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Dot1Pagerduty PagerDuty integration for the service.
type ServiceDefinitionV2Dot1Pagerduty struct {
	// PagerDuty service url.
	ServiceUrl *string `json:"service-url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceDefinitionV2Dot1Pagerduty instantiates a new ServiceDefinitionV2Dot1Pagerduty object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceDefinitionV2Dot1Pagerduty() *ServiceDefinitionV2Dot1Pagerduty {
	this := ServiceDefinitionV2Dot1Pagerduty{}
	return &this
}

// NewServiceDefinitionV2Dot1PagerdutyWithDefaults instantiates a new ServiceDefinitionV2Dot1Pagerduty object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceDefinitionV2Dot1PagerdutyWithDefaults() *ServiceDefinitionV2Dot1Pagerduty {
	this := ServiceDefinitionV2Dot1Pagerduty{}
	return &this
}

// GetServiceUrl returns the ServiceUrl field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot1Pagerduty) GetServiceUrl() string {
	if o == nil || o.ServiceUrl == nil {
		var ret string
		return ret
	}
	return *o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1Pagerduty) GetServiceUrlOk() (*string, bool) {
	if o == nil || o.ServiceUrl == nil {
		return nil, false
	}
	return o.ServiceUrl, true
}

// HasServiceUrl returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot1Pagerduty) HasServiceUrl() bool {
	return o != nil && o.ServiceUrl != nil
}

// SetServiceUrl gets a reference to the given string and assigns it to the ServiceUrl field.
func (o *ServiceDefinitionV2Dot1Pagerduty) SetServiceUrl(v string) {
	o.ServiceUrl = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceDefinitionV2Dot1Pagerduty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ServiceUrl != nil {
		toSerialize["service-url"] = o.ServiceUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceDefinitionV2Dot1Pagerduty) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ServiceUrl *string `json:"service-url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"service-url"})
	} else {
		return err
	}
	o.ServiceUrl = all.ServiceUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
