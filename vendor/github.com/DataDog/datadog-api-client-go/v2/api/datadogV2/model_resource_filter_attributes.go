// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ResourceFilterAttributes Attributes of a resource filter.
type ResourceFilterAttributes struct {
	// A map of cloud provider names (e.g., "aws", "gcp", "azure") to a map of account/resource IDs and their associated tag filters.
	CloudProvider map[string]map[string][]string `json:"cloud_provider"`
	// The UUID of the resource filter.
	Uuid *string `json:"uuid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewResourceFilterAttributes instantiates a new ResourceFilterAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResourceFilterAttributes(cloudProvider map[string]map[string][]string) *ResourceFilterAttributes {
	this := ResourceFilterAttributes{}
	this.CloudProvider = cloudProvider
	return &this
}

// NewResourceFilterAttributesWithDefaults instantiates a new ResourceFilterAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResourceFilterAttributesWithDefaults() *ResourceFilterAttributes {
	this := ResourceFilterAttributes{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value.
func (o *ResourceFilterAttributes) GetCloudProvider() map[string]map[string][]string {
	if o == nil {
		var ret map[string]map[string][]string
		return ret
	}
	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ResourceFilterAttributes) GetCloudProviderOk() (*map[string]map[string][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value.
func (o *ResourceFilterAttributes) SetCloudProvider(v map[string]map[string][]string) {
	o.CloudProvider = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ResourceFilterAttributes) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceFilterAttributes) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ResourceFilterAttributes) HasUuid() bool {
	return o != nil && o.Uuid != nil
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ResourceFilterAttributes) SetUuid(v string) {
	o.Uuid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ResourceFilterAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cloud_provider"] = o.CloudProvider
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ResourceFilterAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CloudProvider *map[string]map[string][]string `json:"cloud_provider"`
		Uuid          *string                         `json:"uuid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CloudProvider == nil {
		return fmt.Errorf("required field cloud_provider missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cloud_provider", "uuid"})
	} else {
		return err
	}
	o.CloudProvider = *all.CloudProvider
	o.Uuid = all.Uuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
