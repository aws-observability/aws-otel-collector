// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListAppsResponseDataItemsAttributes Basic information about the app such as name, description, and tags.
type ListAppsResponseDataItemsAttributes struct {
	// A human-readable description for the app.
	Description *string `json:"description,omitempty"`
	// Whether the app is marked as a favorite by the current user.
	Favorite *bool `json:"favorite,omitempty"`
	// The name of the app.
	Name *string `json:"name,omitempty"`
	// Whether the app is enabled for use in the Datadog self-service hub.
	SelfService *bool `json:"selfService,omitempty"`
	// A list of tags for the app, which can be used to filter apps.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListAppsResponseDataItemsAttributes instantiates a new ListAppsResponseDataItemsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListAppsResponseDataItemsAttributes() *ListAppsResponseDataItemsAttributes {
	this := ListAppsResponseDataItemsAttributes{}
	return &this
}

// NewListAppsResponseDataItemsAttributesWithDefaults instantiates a new ListAppsResponseDataItemsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListAppsResponseDataItemsAttributesWithDefaults() *ListAppsResponseDataItemsAttributes {
	this := ListAppsResponseDataItemsAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListAppsResponseDataItemsAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppsResponseDataItemsAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ListAppsResponseDataItemsAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListAppsResponseDataItemsAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *ListAppsResponseDataItemsAttributes) GetFavorite() bool {
	if o == nil || o.Favorite == nil {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppsResponseDataItemsAttributes) GetFavoriteOk() (*bool, bool) {
	if o == nil || o.Favorite == nil {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *ListAppsResponseDataItemsAttributes) HasFavorite() bool {
	return o != nil && o.Favorite != nil
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *ListAppsResponseDataItemsAttributes) SetFavorite(v bool) {
	o.Favorite = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListAppsResponseDataItemsAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppsResponseDataItemsAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListAppsResponseDataItemsAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListAppsResponseDataItemsAttributes) SetName(v string) {
	o.Name = &v
}

// GetSelfService returns the SelfService field value if set, zero value otherwise.
func (o *ListAppsResponseDataItemsAttributes) GetSelfService() bool {
	if o == nil || o.SelfService == nil {
		var ret bool
		return ret
	}
	return *o.SelfService
}

// GetSelfServiceOk returns a tuple with the SelfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppsResponseDataItemsAttributes) GetSelfServiceOk() (*bool, bool) {
	if o == nil || o.SelfService == nil {
		return nil, false
	}
	return o.SelfService, true
}

// HasSelfService returns a boolean if a field has been set.
func (o *ListAppsResponseDataItemsAttributes) HasSelfService() bool {
	return o != nil && o.SelfService != nil
}

// SetSelfService gets a reference to the given bool and assigns it to the SelfService field.
func (o *ListAppsResponseDataItemsAttributes) SetSelfService(v bool) {
	o.SelfService = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ListAppsResponseDataItemsAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppsResponseDataItemsAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ListAppsResponseDataItemsAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ListAppsResponseDataItemsAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListAppsResponseDataItemsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Favorite != nil {
		toSerialize["favorite"] = o.Favorite
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SelfService != nil {
		toSerialize["selfService"] = o.SelfService
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListAppsResponseDataItemsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string  `json:"description,omitempty"`
		Favorite    *bool    `json:"favorite,omitempty"`
		Name        *string  `json:"name,omitempty"`
		SelfService *bool    `json:"selfService,omitempty"`
		Tags        []string `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "favorite", "name", "selfService", "tags"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Favorite = all.Favorite
	o.Name = all.Name
	o.SelfService = all.SelfService
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
