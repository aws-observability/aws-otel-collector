// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPMonitoredResourceConfig Configuration for a GCP monitored resource.
type GCPMonitoredResourceConfig struct {
	// List of filters to limit the monitored resources that are pulled into Datadog by using tags.
	// Only monitored resources that apply to specified filters are imported into Datadog.
	Filters []string `json:"filters,omitempty"`
	// The GCP monitored resource type. Only a subset of resource types are supported.
	Type *GCPMonitoredResourceConfigType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGCPMonitoredResourceConfig instantiates a new GCPMonitoredResourceConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPMonitoredResourceConfig() *GCPMonitoredResourceConfig {
	this := GCPMonitoredResourceConfig{}
	return &this
}

// NewGCPMonitoredResourceConfigWithDefaults instantiates a new GCPMonitoredResourceConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPMonitoredResourceConfigWithDefaults() *GCPMonitoredResourceConfig {
	this := GCPMonitoredResourceConfig{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *GCPMonitoredResourceConfig) GetFilters() []string {
	if o == nil || o.Filters == nil {
		var ret []string
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPMonitoredResourceConfig) GetFiltersOk() (*[]string, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *GCPMonitoredResourceConfig) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given []string and assigns it to the Filters field.
func (o *GCPMonitoredResourceConfig) SetFilters(v []string) {
	o.Filters = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GCPMonitoredResourceConfig) GetType() GCPMonitoredResourceConfigType {
	if o == nil || o.Type == nil {
		var ret GCPMonitoredResourceConfigType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPMonitoredResourceConfig) GetTypeOk() (*GCPMonitoredResourceConfigType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GCPMonitoredResourceConfig) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given GCPMonitoredResourceConfigType and assigns it to the Type field.
func (o *GCPMonitoredResourceConfig) SetType(v GCPMonitoredResourceConfigType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPMonitoredResourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GCPMonitoredResourceConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filters []string                        `json:"filters,omitempty"`
		Type    *GCPMonitoredResourceConfigType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filters", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Filters = all.Filters
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
