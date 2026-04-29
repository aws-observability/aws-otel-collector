// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDatadogLogsDestinationRoute Defines how the `datadog_logs` destination routes matching logs to a Datadog site using a specific API key.
type ObservabilityPipelineDatadogLogsDestinationRoute struct {
	// Name of the environment variable or secret that stores the Datadog API key used by this route.
	ApiKeyKey *string `json:"api_key_key,omitempty"`
	// A Datadog search query that determines which logs are forwarded using this route.
	Include *string `json:"include,omitempty"`
	// Unique identifier for this route within the destination.
	RouteId *string `json:"route_id,omitempty"`
	// Datadog site where matching logs are sent (for example, `us1`).
	Site *string `json:"site,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineDatadogLogsDestinationRoute instantiates a new ObservabilityPipelineDatadogLogsDestinationRoute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineDatadogLogsDestinationRoute() *ObservabilityPipelineDatadogLogsDestinationRoute {
	this := ObservabilityPipelineDatadogLogsDestinationRoute{}
	return &this
}

// NewObservabilityPipelineDatadogLogsDestinationRouteWithDefaults instantiates a new ObservabilityPipelineDatadogLogsDestinationRoute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineDatadogLogsDestinationRouteWithDefaults() *ObservabilityPipelineDatadogLogsDestinationRoute {
	this := ObservabilityPipelineDatadogLogsDestinationRoute{}
	return &this
}

// GetApiKeyKey returns the ApiKeyKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) GetApiKeyKey() string {
	if o == nil || o.ApiKeyKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKeyKey
}

// GetApiKeyKeyOk returns a tuple with the ApiKeyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) GetApiKeyKeyOk() (*string, bool) {
	if o == nil || o.ApiKeyKey == nil {
		return nil, false
	}
	return o.ApiKeyKey, true
}

// HasApiKeyKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) HasApiKeyKey() bool {
	return o != nil && o.ApiKeyKey != nil
}

// SetApiKeyKey gets a reference to the given string and assigns it to the ApiKeyKey field.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) SetApiKeyKey(v string) {
	o.ApiKeyKey = &v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) GetInclude() string {
	if o == nil || o.Include == nil {
		var ret string
		return ret
	}
	return *o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) GetIncludeOk() (*string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) HasInclude() bool {
	return o != nil && o.Include != nil
}

// SetInclude gets a reference to the given string and assigns it to the Include field.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) SetInclude(v string) {
	o.Include = &v
}

// GetRouteId returns the RouteId field value if set, zero value otherwise.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) GetRouteId() string {
	if o == nil || o.RouteId == nil {
		var ret string
		return ret
	}
	return *o.RouteId
}

// GetRouteIdOk returns a tuple with the RouteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) GetRouteIdOk() (*string, bool) {
	if o == nil || o.RouteId == nil {
		return nil, false
	}
	return o.RouteId, true
}

// HasRouteId returns a boolean if a field has been set.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) HasRouteId() bool {
	return o != nil && o.RouteId != nil
}

// SetRouteId gets a reference to the given string and assigns it to the RouteId field.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) SetRouteId(v string) {
	o.RouteId = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) GetSite() string {
	if o == nil || o.Site == nil {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) GetSiteOk() (*string, bool) {
	if o == nil || o.Site == nil {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) HasSite() bool {
	return o != nil && o.Site != nil
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) SetSite(v string) {
	o.Site = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineDatadogLogsDestinationRoute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApiKeyKey != nil {
		toSerialize["api_key_key"] = o.ApiKeyKey
	}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}
	if o.RouteId != nil {
		toSerialize["route_id"] = o.RouteId
	}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineDatadogLogsDestinationRoute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKeyKey *string `json:"api_key_key,omitempty"`
		Include   *string `json:"include,omitempty"`
		RouteId   *string `json:"route_id,omitempty"`
		Site      *string `json:"site,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key_key", "include", "route_id", "site"})
	} else {
		return err
	}
	o.ApiKeyKey = all.ApiKeyKey
	o.Include = all.Include
	o.RouteId = all.RouteId
	o.Site = all.Site

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
