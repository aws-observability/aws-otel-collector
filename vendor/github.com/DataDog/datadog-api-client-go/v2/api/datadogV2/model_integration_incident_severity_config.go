// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationIncidentSeverityConfig Severity configuration for mapping incident priorities to case priorities.
type IntegrationIncidentSeverityConfig struct {
	// Mapping of incident severity values to case priority values.
	PriorityMapping map[string]string `json:"priority_mapping,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationIncidentSeverityConfig instantiates a new IntegrationIncidentSeverityConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationIncidentSeverityConfig() *IntegrationIncidentSeverityConfig {
	this := IntegrationIncidentSeverityConfig{}
	return &this
}

// NewIntegrationIncidentSeverityConfigWithDefaults instantiates a new IntegrationIncidentSeverityConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationIncidentSeverityConfigWithDefaults() *IntegrationIncidentSeverityConfig {
	this := IntegrationIncidentSeverityConfig{}
	return &this
}

// GetPriorityMapping returns the PriorityMapping field value if set, zero value otherwise.
func (o *IntegrationIncidentSeverityConfig) GetPriorityMapping() map[string]string {
	if o == nil || o.PriorityMapping == nil {
		var ret map[string]string
		return ret
	}
	return o.PriorityMapping
}

// GetPriorityMappingOk returns a tuple with the PriorityMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationIncidentSeverityConfig) GetPriorityMappingOk() (*map[string]string, bool) {
	if o == nil || o.PriorityMapping == nil {
		return nil, false
	}
	return &o.PriorityMapping, true
}

// HasPriorityMapping returns a boolean if a field has been set.
func (o *IntegrationIncidentSeverityConfig) HasPriorityMapping() bool {
	return o != nil && o.PriorityMapping != nil
}

// SetPriorityMapping gets a reference to the given map[string]string and assigns it to the PriorityMapping field.
func (o *IntegrationIncidentSeverityConfig) SetPriorityMapping(v map[string]string) {
	o.PriorityMapping = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationIncidentSeverityConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.PriorityMapping != nil {
		toSerialize["priority_mapping"] = o.PriorityMapping
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationIncidentSeverityConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PriorityMapping map[string]string `json:"priority_mapping,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"priority_mapping"})
	} else {
		return err
	}
	o.PriorityMapping = all.PriorityMapping

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
