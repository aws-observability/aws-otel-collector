// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3DatadogIntegrationPagerduty A PagerDuty integration schema.
type EntityV3DatadogIntegrationPagerduty struct {
	// The service URL for the PagerDuty integration.
	ServiceUrl string `json:"serviceURL"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3DatadogIntegrationPagerduty instantiates a new EntityV3DatadogIntegrationPagerduty object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3DatadogIntegrationPagerduty(serviceUrl string) *EntityV3DatadogIntegrationPagerduty {
	this := EntityV3DatadogIntegrationPagerduty{}
	this.ServiceUrl = serviceUrl
	return &this
}

// NewEntityV3DatadogIntegrationPagerdutyWithDefaults instantiates a new EntityV3DatadogIntegrationPagerduty object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3DatadogIntegrationPagerdutyWithDefaults() *EntityV3DatadogIntegrationPagerduty {
	this := EntityV3DatadogIntegrationPagerduty{}
	return &this
}

// GetServiceUrl returns the ServiceUrl field value.
func (o *EntityV3DatadogIntegrationPagerduty) GetServiceUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value
// and a boolean to check if the value has been set.
func (o *EntityV3DatadogIntegrationPagerduty) GetServiceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceUrl, true
}

// SetServiceUrl sets field value.
func (o *EntityV3DatadogIntegrationPagerduty) SetServiceUrl(v string) {
	o.ServiceUrl = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3DatadogIntegrationPagerduty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["serviceURL"] = o.ServiceUrl
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3DatadogIntegrationPagerduty) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ServiceUrl *string `json:"serviceURL"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ServiceUrl == nil {
		return fmt.Errorf("required field serviceURL missing")
	}
	o.ServiceUrl = *all.ServiceUrl

	return nil
}
