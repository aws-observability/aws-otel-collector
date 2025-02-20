// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3DatadogIntegrationOpsgenie An Opsgenie integration schema.
type EntityV3DatadogIntegrationOpsgenie struct {
	// The region for the Opsgenie integration.
	Region *string `json:"region,omitempty"`
	// The service URL for the Opsgenie integration.
	ServiceUrl string `json:"serviceURL"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3DatadogIntegrationOpsgenie instantiates a new EntityV3DatadogIntegrationOpsgenie object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3DatadogIntegrationOpsgenie(serviceUrl string) *EntityV3DatadogIntegrationOpsgenie {
	this := EntityV3DatadogIntegrationOpsgenie{}
	this.ServiceUrl = serviceUrl
	return &this
}

// NewEntityV3DatadogIntegrationOpsgenieWithDefaults instantiates a new EntityV3DatadogIntegrationOpsgenie object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3DatadogIntegrationOpsgenieWithDefaults() *EntityV3DatadogIntegrationOpsgenie {
	this := EntityV3DatadogIntegrationOpsgenie{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *EntityV3DatadogIntegrationOpsgenie) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3DatadogIntegrationOpsgenie) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *EntityV3DatadogIntegrationOpsgenie) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *EntityV3DatadogIntegrationOpsgenie) SetRegion(v string) {
	o.Region = &v
}

// GetServiceUrl returns the ServiceUrl field value.
func (o *EntityV3DatadogIntegrationOpsgenie) GetServiceUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value
// and a boolean to check if the value has been set.
func (o *EntityV3DatadogIntegrationOpsgenie) GetServiceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceUrl, true
}

// SetServiceUrl sets field value.
func (o *EntityV3DatadogIntegrationOpsgenie) SetServiceUrl(v string) {
	o.ServiceUrl = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3DatadogIntegrationOpsgenie) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	toSerialize["serviceURL"] = o.ServiceUrl
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3DatadogIntegrationOpsgenie) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Region     *string `json:"region,omitempty"`
		ServiceUrl *string `json:"serviceURL"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ServiceUrl == nil {
		return fmt.Errorf("required field serviceURL missing")
	}
	o.Region = all.Region
	o.ServiceUrl = *all.ServiceUrl

	return nil
}
