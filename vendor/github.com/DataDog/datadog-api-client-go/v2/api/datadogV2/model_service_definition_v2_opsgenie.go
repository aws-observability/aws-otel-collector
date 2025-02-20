// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Opsgenie Opsgenie integration for the service.
type ServiceDefinitionV2Opsgenie struct {
	// Opsgenie instance region.
	Region *ServiceDefinitionV2OpsgenieRegion `json:"region,omitempty"`
	// Opsgenie service url.
	ServiceUrl string `json:"service-url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceDefinitionV2Opsgenie instantiates a new ServiceDefinitionV2Opsgenie object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceDefinitionV2Opsgenie(serviceUrl string) *ServiceDefinitionV2Opsgenie {
	this := ServiceDefinitionV2Opsgenie{}
	this.ServiceUrl = serviceUrl
	return &this
}

// NewServiceDefinitionV2OpsgenieWithDefaults instantiates a new ServiceDefinitionV2Opsgenie object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceDefinitionV2OpsgenieWithDefaults() *ServiceDefinitionV2Opsgenie {
	this := ServiceDefinitionV2Opsgenie{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Opsgenie) GetRegion() ServiceDefinitionV2OpsgenieRegion {
	if o == nil || o.Region == nil {
		var ret ServiceDefinitionV2OpsgenieRegion
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Opsgenie) GetRegionOk() (*ServiceDefinitionV2OpsgenieRegion, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Opsgenie) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given ServiceDefinitionV2OpsgenieRegion and assigns it to the Region field.
func (o *ServiceDefinitionV2Opsgenie) SetRegion(v ServiceDefinitionV2OpsgenieRegion) {
	o.Region = &v
}

// GetServiceUrl returns the ServiceUrl field value.
func (o *ServiceDefinitionV2Opsgenie) GetServiceUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Opsgenie) GetServiceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceUrl, true
}

// SetServiceUrl sets field value.
func (o *ServiceDefinitionV2Opsgenie) SetServiceUrl(v string) {
	o.ServiceUrl = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceDefinitionV2Opsgenie) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	toSerialize["service-url"] = o.ServiceUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceDefinitionV2Opsgenie) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Region     *ServiceDefinitionV2OpsgenieRegion `json:"region,omitempty"`
		ServiceUrl *string                            `json:"service-url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ServiceUrl == nil {
		return fmt.Errorf("required field service-url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"region", "service-url"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Region != nil && !all.Region.IsValid() {
		hasInvalidField = true
	} else {
		o.Region = all.Region
	}
	o.ServiceUrl = *all.ServiceUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
