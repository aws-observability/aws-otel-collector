// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureAccountMetricsConfig Dictionary containing the key `excluded_resource_providers` which has to be a list of Microsoft Azure Resource Provider names.
// This feature is currently being beta tested.
// In order to enable all resource providers for metric collection, pass:
// `metrics_config: {"excluded_resource_providers": []}` (i.e., an empty list for `excluded_resource_providers`).
type AzureAccountMetricsConfig struct {
	// List of Microsoft Azure Resource Providers to exclude from metric collection.
	ExcludedResourceProviders []string `json:"excluded_resource_providers,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAzureAccountMetricsConfig instantiates a new AzureAccountMetricsConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAzureAccountMetricsConfig() *AzureAccountMetricsConfig {
	this := AzureAccountMetricsConfig{}
	return &this
}

// NewAzureAccountMetricsConfigWithDefaults instantiates a new AzureAccountMetricsConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAzureAccountMetricsConfigWithDefaults() *AzureAccountMetricsConfig {
	this := AzureAccountMetricsConfig{}
	return &this
}

// GetExcludedResourceProviders returns the ExcludedResourceProviders field value if set, zero value otherwise.
func (o *AzureAccountMetricsConfig) GetExcludedResourceProviders() []string {
	if o == nil || o.ExcludedResourceProviders == nil {
		var ret []string
		return ret
	}
	return o.ExcludedResourceProviders
}

// GetExcludedResourceProvidersOk returns a tuple with the ExcludedResourceProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccountMetricsConfig) GetExcludedResourceProvidersOk() (*[]string, bool) {
	if o == nil || o.ExcludedResourceProviders == nil {
		return nil, false
	}
	return &o.ExcludedResourceProviders, true
}

// HasExcludedResourceProviders returns a boolean if a field has been set.
func (o *AzureAccountMetricsConfig) HasExcludedResourceProviders() bool {
	return o != nil && o.ExcludedResourceProviders != nil
}

// SetExcludedResourceProviders gets a reference to the given []string and assigns it to the ExcludedResourceProviders field.
func (o *AzureAccountMetricsConfig) SetExcludedResourceProviders(v []string) {
	o.ExcludedResourceProviders = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AzureAccountMetricsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExcludedResourceProviders != nil {
		toSerialize["excluded_resource_providers"] = o.ExcludedResourceProviders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AzureAccountMetricsConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExcludedResourceProviders []string `json:"excluded_resource_providers,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"excluded_resource_providers"})
	} else {
		return err
	}
	o.ExcludedResourceProviders = all.ExcludedResourceProviders

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
