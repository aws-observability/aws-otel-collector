// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatUpdateOrganizationHandleRequestAttributes Organization handle attributes for an update request.
type GoogleChatUpdateOrganizationHandleRequestAttributes struct {
	// Organization handle name.
	Name *string `json:"name,omitempty"`
	// Google space resource name.
	SpaceResourceName *string `json:"space_resource_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGoogleChatUpdateOrganizationHandleRequestAttributes instantiates a new GoogleChatUpdateOrganizationHandleRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGoogleChatUpdateOrganizationHandleRequestAttributes() *GoogleChatUpdateOrganizationHandleRequestAttributes {
	this := GoogleChatUpdateOrganizationHandleRequestAttributes{}
	return &this
}

// NewGoogleChatUpdateOrganizationHandleRequestAttributesWithDefaults instantiates a new GoogleChatUpdateOrganizationHandleRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGoogleChatUpdateOrganizationHandleRequestAttributesWithDefaults() *GoogleChatUpdateOrganizationHandleRequestAttributes {
	this := GoogleChatUpdateOrganizationHandleRequestAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GoogleChatUpdateOrganizationHandleRequestAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatUpdateOrganizationHandleRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GoogleChatUpdateOrganizationHandleRequestAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GoogleChatUpdateOrganizationHandleRequestAttributes) SetName(v string) {
	o.Name = &v
}

// GetSpaceResourceName returns the SpaceResourceName field value if set, zero value otherwise.
func (o *GoogleChatUpdateOrganizationHandleRequestAttributes) GetSpaceResourceName() string {
	if o == nil || o.SpaceResourceName == nil {
		var ret string
		return ret
	}
	return *o.SpaceResourceName
}

// GetSpaceResourceNameOk returns a tuple with the SpaceResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatUpdateOrganizationHandleRequestAttributes) GetSpaceResourceNameOk() (*string, bool) {
	if o == nil || o.SpaceResourceName == nil {
		return nil, false
	}
	return o.SpaceResourceName, true
}

// HasSpaceResourceName returns a boolean if a field has been set.
func (o *GoogleChatUpdateOrganizationHandleRequestAttributes) HasSpaceResourceName() bool {
	return o != nil && o.SpaceResourceName != nil
}

// SetSpaceResourceName gets a reference to the given string and assigns it to the SpaceResourceName field.
func (o *GoogleChatUpdateOrganizationHandleRequestAttributes) SetSpaceResourceName(v string) {
	o.SpaceResourceName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GoogleChatUpdateOrganizationHandleRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SpaceResourceName != nil {
		toSerialize["space_resource_name"] = o.SpaceResourceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GoogleChatUpdateOrganizationHandleRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string `json:"name,omitempty"`
		SpaceResourceName *string `json:"space_resource_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "space_resource_name"})
	} else {
		return err
	}
	o.Name = all.Name
	o.SpaceResourceName = all.SpaceResourceName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
