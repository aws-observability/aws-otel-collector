// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorAsset Represents key links tied to a monitor to help users take action on alerts.
// This feature is in Preview and only available to users with the feature enabled.
type MonitorAsset struct {
	// Indicates the type of asset this entity represents on a monitor.
	Category MonitorAssetCategory `json:"category"`
	// Name for the monitor asset
	Name string `json:"name"`
	// Represents the identifier of the internal Datadog resource that this asset represents. IDs in this field should be passed in as strings.
	ResourceKey *string `json:"resource_key,omitempty"`
	// Type of internal Datadog resource associated with a monitor asset.
	ResourceType *MonitorAssetResourceType `json:"resource_type,omitempty"`
	// URL link for the asset. For links with an internal resource type set, this should be the relative path to where the Datadog domain is appended internally. For external links, this should be the full URL path.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorAsset instantiates a new MonitorAsset object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorAsset(category MonitorAssetCategory, name string, url string) *MonitorAsset {
	this := MonitorAsset{}
	this.Category = category
	this.Name = name
	this.Url = url
	return &this
}

// NewMonitorAssetWithDefaults instantiates a new MonitorAsset object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorAssetWithDefaults() *MonitorAsset {
	this := MonitorAsset{}
	return &this
}

// GetCategory returns the Category field value.
func (o *MonitorAsset) GetCategory() MonitorAssetCategory {
	if o == nil {
		var ret MonitorAssetCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *MonitorAsset) GetCategoryOk() (*MonitorAssetCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *MonitorAsset) SetCategory(v MonitorAssetCategory) {
	o.Category = v
}

// GetName returns the Name field value.
func (o *MonitorAsset) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MonitorAsset) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MonitorAsset) SetName(v string) {
	o.Name = v
}

// GetResourceKey returns the ResourceKey field value if set, zero value otherwise.
func (o *MonitorAsset) GetResourceKey() string {
	if o == nil || o.ResourceKey == nil {
		var ret string
		return ret
	}
	return *o.ResourceKey
}

// GetResourceKeyOk returns a tuple with the ResourceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorAsset) GetResourceKeyOk() (*string, bool) {
	if o == nil || o.ResourceKey == nil {
		return nil, false
	}
	return o.ResourceKey, true
}

// HasResourceKey returns a boolean if a field has been set.
func (o *MonitorAsset) HasResourceKey() bool {
	return o != nil && o.ResourceKey != nil
}

// SetResourceKey gets a reference to the given string and assigns it to the ResourceKey field.
func (o *MonitorAsset) SetResourceKey(v string) {
	o.ResourceKey = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *MonitorAsset) GetResourceType() MonitorAssetResourceType {
	if o == nil || o.ResourceType == nil {
		var ret MonitorAssetResourceType
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorAsset) GetResourceTypeOk() (*MonitorAssetResourceType, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *MonitorAsset) HasResourceType() bool {
	return o != nil && o.ResourceType != nil
}

// SetResourceType gets a reference to the given MonitorAssetResourceType and assigns it to the ResourceType field.
func (o *MonitorAsset) SetResourceType(v MonitorAssetResourceType) {
	o.ResourceType = &v
}

// GetUrl returns the Url field value.
func (o *MonitorAsset) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *MonitorAsset) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *MonitorAsset) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorAsset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["name"] = o.Name
	if o.ResourceKey != nil {
		toSerialize["resource_key"] = o.ResourceKey
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorAsset) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category     *MonitorAssetCategory     `json:"category"`
		Name         *string                   `json:"name"`
		ResourceKey  *string                   `json:"resource_key,omitempty"`
		ResourceType *MonitorAssetResourceType `json:"resource_type,omitempty"`
		Url          *string                   `json:"url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "name", "resource_key", "resource_type", "url"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = *all.Category
	}
	o.Name = *all.Name
	o.ResourceKey = all.ResourceKey
	if all.ResourceType != nil && !all.ResourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ResourceType = all.ResourceType
	}
	o.Url = *all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
