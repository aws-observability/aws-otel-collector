// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConfluentResourceResponseAttributes Model representation of a Confluent Cloud resource.
type ConfluentResourceResponseAttributes struct {
	// Enable the `custom.consumer_lag_offset` metric, which contains extra metric tags.
	EnableCustomMetrics *bool `json:"enable_custom_metrics,omitempty"`
	// The ID associated with the Confluent resource.
	Id *string `json:"id,omitempty"`
	// The resource type of the Resource. Can be `kafka`, `connector`, `ksql`, or `schema_registry`.
	ResourceType string `json:"resource_type"`
	// A list of strings representing tags. Can be a single key, or key-value pairs separated by a colon.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConfluentResourceResponseAttributes instantiates a new ConfluentResourceResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConfluentResourceResponseAttributes(resourceType string) *ConfluentResourceResponseAttributes {
	this := ConfluentResourceResponseAttributes{}
	var enableCustomMetrics bool = false
	this.EnableCustomMetrics = &enableCustomMetrics
	this.ResourceType = resourceType
	return &this
}

// NewConfluentResourceResponseAttributesWithDefaults instantiates a new ConfluentResourceResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConfluentResourceResponseAttributesWithDefaults() *ConfluentResourceResponseAttributes {
	this := ConfluentResourceResponseAttributes{}
	var enableCustomMetrics bool = false
	this.EnableCustomMetrics = &enableCustomMetrics
	return &this
}

// GetEnableCustomMetrics returns the EnableCustomMetrics field value if set, zero value otherwise.
func (o *ConfluentResourceResponseAttributes) GetEnableCustomMetrics() bool {
	if o == nil || o.EnableCustomMetrics == nil {
		var ret bool
		return ret
	}
	return *o.EnableCustomMetrics
}

// GetEnableCustomMetricsOk returns a tuple with the EnableCustomMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfluentResourceResponseAttributes) GetEnableCustomMetricsOk() (*bool, bool) {
	if o == nil || o.EnableCustomMetrics == nil {
		return nil, false
	}
	return o.EnableCustomMetrics, true
}

// HasEnableCustomMetrics returns a boolean if a field has been set.
func (o *ConfluentResourceResponseAttributes) HasEnableCustomMetrics() bool {
	return o != nil && o.EnableCustomMetrics != nil
}

// SetEnableCustomMetrics gets a reference to the given bool and assigns it to the EnableCustomMetrics field.
func (o *ConfluentResourceResponseAttributes) SetEnableCustomMetrics(v bool) {
	o.EnableCustomMetrics = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfluentResourceResponseAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfluentResourceResponseAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfluentResourceResponseAttributes) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConfluentResourceResponseAttributes) SetId(v string) {
	o.Id = &v
}

// GetResourceType returns the ResourceType field value.
func (o *ConfluentResourceResponseAttributes) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ConfluentResourceResponseAttributes) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *ConfluentResourceResponseAttributes) SetResourceType(v string) {
	o.ResourceType = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConfluentResourceResponseAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfluentResourceResponseAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConfluentResourceResponseAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ConfluentResourceResponseAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConfluentResourceResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EnableCustomMetrics != nil {
		toSerialize["enable_custom_metrics"] = o.EnableCustomMetrics
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["resource_type"] = o.ResourceType
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConfluentResourceResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnableCustomMetrics *bool    `json:"enable_custom_metrics,omitempty"`
		Id                  *string  `json:"id,omitempty"`
		ResourceType        *string  `json:"resource_type"`
		Tags                []string `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ResourceType == nil {
		return fmt.Errorf("required field resource_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enable_custom_metrics", "id", "resource_type", "tags"})
	} else {
		return err
	}
	o.EnableCustomMetrics = all.EnableCustomMetrics
	o.Id = all.Id
	o.ResourceType = *all.ResourceType
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
