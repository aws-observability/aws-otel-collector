// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabFacetKeysAttributes Available facet key names for filtering resources.
type ModelLabFacetKeysAttributes struct {
	// The list of available metric facet keys.
	Metrics datadog.NullableList[string] `json:"metrics"`
	// The list of available parameter facet keys.
	Parameters []string `json:"parameters"`
	// The list of available tag facet keys.
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabFacetKeysAttributes instantiates a new ModelLabFacetKeysAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabFacetKeysAttributes(metrics datadog.NullableList[string], parameters []string, tags []string) *ModelLabFacetKeysAttributes {
	this := ModelLabFacetKeysAttributes{}
	this.Metrics = metrics
	this.Parameters = parameters
	this.Tags = tags
	return &this
}

// NewModelLabFacetKeysAttributesWithDefaults instantiates a new ModelLabFacetKeysAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabFacetKeysAttributesWithDefaults() *ModelLabFacetKeysAttributes {
	this := ModelLabFacetKeysAttributes{}
	return &this
}

// GetMetrics returns the Metrics field value.
// If the value is explicit nil, the zero value for []string will be returned.
func (o *ModelLabFacetKeysAttributes) GetMetrics() []string {
	if o == nil || o.Metrics.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Metrics.Get()
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabFacetKeysAttributes) GetMetricsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metrics.Get(), o.Metrics.IsSet()
}

// SetMetrics sets field value.
func (o *ModelLabFacetKeysAttributes) SetMetrics(v []string) {
	o.Metrics.Set(&v)
}

// GetParameters returns the Parameters field value.
func (o *ModelLabFacetKeysAttributes) GetParameters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *ModelLabFacetKeysAttributes) GetParametersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value.
func (o *ModelLabFacetKeysAttributes) SetParameters(v []string) {
	o.Parameters = v
}

// GetTags returns the Tags field value.
func (o *ModelLabFacetKeysAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ModelLabFacetKeysAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *ModelLabFacetKeysAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabFacetKeysAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["metrics"] = o.Metrics.Get()
	toSerialize["parameters"] = o.Parameters
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModelLabFacetKeysAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Metrics    datadog.NullableList[string] `json:"metrics"`
		Parameters *[]string                    `json:"parameters"`
		Tags       *[]string                    `json:"tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Metrics.IsSet() {
		return fmt.Errorf("required field metrics missing")
	}
	if all.Parameters == nil {
		return fmt.Errorf("required field parameters missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"metrics", "parameters", "tags"})
	} else {
		return err
	}
	o.Metrics = all.Metrics
	o.Parameters = *all.Parameters
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
