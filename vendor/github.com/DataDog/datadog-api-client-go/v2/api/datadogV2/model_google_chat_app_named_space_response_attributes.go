// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatAppNamedSpaceResponseAttributes Google Chat space attributes.
type GoogleChatAppNamedSpaceResponseAttributes struct {
	// Google space display name.
	DisplayName *string `json:"display_name,omitempty"`
	// Organization binding ID.
	OrganizationBindingId *string `json:"organization_binding_id,omitempty"`
	// Google space resource name.
	ResourceName *string `json:"resource_name,omitempty"`
	// Google space URI.
	SpaceUri *string `json:"space_uri,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGoogleChatAppNamedSpaceResponseAttributes instantiates a new GoogleChatAppNamedSpaceResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGoogleChatAppNamedSpaceResponseAttributes() *GoogleChatAppNamedSpaceResponseAttributes {
	this := GoogleChatAppNamedSpaceResponseAttributes{}
	return &this
}

// NewGoogleChatAppNamedSpaceResponseAttributesWithDefaults instantiates a new GoogleChatAppNamedSpaceResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGoogleChatAppNamedSpaceResponseAttributesWithDefaults() *GoogleChatAppNamedSpaceResponseAttributes {
	this := GoogleChatAppNamedSpaceResponseAttributes{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GoogleChatAppNamedSpaceResponseAttributes) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatAppNamedSpaceResponseAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GoogleChatAppNamedSpaceResponseAttributes) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GoogleChatAppNamedSpaceResponseAttributes) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetOrganizationBindingId returns the OrganizationBindingId field value if set, zero value otherwise.
func (o *GoogleChatAppNamedSpaceResponseAttributes) GetOrganizationBindingId() string {
	if o == nil || o.OrganizationBindingId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationBindingId
}

// GetOrganizationBindingIdOk returns a tuple with the OrganizationBindingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatAppNamedSpaceResponseAttributes) GetOrganizationBindingIdOk() (*string, bool) {
	if o == nil || o.OrganizationBindingId == nil {
		return nil, false
	}
	return o.OrganizationBindingId, true
}

// HasOrganizationBindingId returns a boolean if a field has been set.
func (o *GoogleChatAppNamedSpaceResponseAttributes) HasOrganizationBindingId() bool {
	return o != nil && o.OrganizationBindingId != nil
}

// SetOrganizationBindingId gets a reference to the given string and assigns it to the OrganizationBindingId field.
func (o *GoogleChatAppNamedSpaceResponseAttributes) SetOrganizationBindingId(v string) {
	o.OrganizationBindingId = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *GoogleChatAppNamedSpaceResponseAttributes) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatAppNamedSpaceResponseAttributes) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *GoogleChatAppNamedSpaceResponseAttributes) HasResourceName() bool {
	return o != nil && o.ResourceName != nil
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *GoogleChatAppNamedSpaceResponseAttributes) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetSpaceUri returns the SpaceUri field value if set, zero value otherwise.
func (o *GoogleChatAppNamedSpaceResponseAttributes) GetSpaceUri() string {
	if o == nil || o.SpaceUri == nil {
		var ret string
		return ret
	}
	return *o.SpaceUri
}

// GetSpaceUriOk returns a tuple with the SpaceUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatAppNamedSpaceResponseAttributes) GetSpaceUriOk() (*string, bool) {
	if o == nil || o.SpaceUri == nil {
		return nil, false
	}
	return o.SpaceUri, true
}

// HasSpaceUri returns a boolean if a field has been set.
func (o *GoogleChatAppNamedSpaceResponseAttributes) HasSpaceUri() bool {
	return o != nil && o.SpaceUri != nil
}

// SetSpaceUri gets a reference to the given string and assigns it to the SpaceUri field.
func (o *GoogleChatAppNamedSpaceResponseAttributes) SetSpaceUri(v string) {
	o.SpaceUri = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GoogleChatAppNamedSpaceResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.OrganizationBindingId != nil {
		toSerialize["organization_binding_id"] = o.OrganizationBindingId
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	if o.SpaceUri != nil {
		toSerialize["space_uri"] = o.SpaceUri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GoogleChatAppNamedSpaceResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName           *string `json:"display_name,omitempty"`
		OrganizationBindingId *string `json:"organization_binding_id,omitempty"`
		ResourceName          *string `json:"resource_name,omitempty"`
		SpaceUri              *string `json:"space_uri,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "organization_binding_id", "resource_name", "space_uri"})
	} else {
		return err
	}
	o.DisplayName = all.DisplayName
	o.OrganizationBindingId = all.OrganizationBindingId
	o.ResourceName = all.ResourceName
	o.SpaceUri = all.SpaceUri

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
