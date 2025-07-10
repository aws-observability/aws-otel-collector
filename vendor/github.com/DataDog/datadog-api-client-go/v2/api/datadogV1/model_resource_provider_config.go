// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ResourceProviderConfig Configuration settings applied to resources from the specified Azure resource provider.
type ResourceProviderConfig struct {
	// Collect metrics for resources from this provider.
	MetricsEnabled *bool `json:"metrics_enabled,omitempty"`
	// The provider namespace to apply this configuration to.
	Namespace *string `json:"namespace,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewResourceProviderConfig instantiates a new ResourceProviderConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResourceProviderConfig() *ResourceProviderConfig {
	this := ResourceProviderConfig{}
	return &this
}

// NewResourceProviderConfigWithDefaults instantiates a new ResourceProviderConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResourceProviderConfigWithDefaults() *ResourceProviderConfig {
	this := ResourceProviderConfig{}
	return &this
}

// GetMetricsEnabled returns the MetricsEnabled field value if set, zero value otherwise.
func (o *ResourceProviderConfig) GetMetricsEnabled() bool {
	if o == nil || o.MetricsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.MetricsEnabled
}

// GetMetricsEnabledOk returns a tuple with the MetricsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceProviderConfig) GetMetricsEnabledOk() (*bool, bool) {
	if o == nil || o.MetricsEnabled == nil {
		return nil, false
	}
	return o.MetricsEnabled, true
}

// HasMetricsEnabled returns a boolean if a field has been set.
func (o *ResourceProviderConfig) HasMetricsEnabled() bool {
	return o != nil && o.MetricsEnabled != nil
}

// SetMetricsEnabled gets a reference to the given bool and assigns it to the MetricsEnabled field.
func (o *ResourceProviderConfig) SetMetricsEnabled(v bool) {
	o.MetricsEnabled = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ResourceProviderConfig) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceProviderConfig) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ResourceProviderConfig) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ResourceProviderConfig) SetNamespace(v string) {
	o.Namespace = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ResourceProviderConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MetricsEnabled != nil {
		toSerialize["metrics_enabled"] = o.MetricsEnabled
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ResourceProviderConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MetricsEnabled *bool   `json:"metrics_enabled,omitempty"`
		Namespace      *string `json:"namespace,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"metrics_enabled", "namespace"})
	} else {
		return err
	}
	o.MetricsEnabled = all.MetricsEnabled
	o.Namespace = all.Namespace

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
