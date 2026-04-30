// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatCreateOrganizationHandleRequestAttributes Organization handle attributes for a create request.
type GoogleChatCreateOrganizationHandleRequestAttributes struct {
	// Organization handle name.
	Name string `json:"name"`
	// Google space resource name.
	SpaceResourceName string `json:"space_resource_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGoogleChatCreateOrganizationHandleRequestAttributes instantiates a new GoogleChatCreateOrganizationHandleRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGoogleChatCreateOrganizationHandleRequestAttributes(name string, spaceResourceName string) *GoogleChatCreateOrganizationHandleRequestAttributes {
	this := GoogleChatCreateOrganizationHandleRequestAttributes{}
	this.Name = name
	this.SpaceResourceName = spaceResourceName
	return &this
}

// NewGoogleChatCreateOrganizationHandleRequestAttributesWithDefaults instantiates a new GoogleChatCreateOrganizationHandleRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGoogleChatCreateOrganizationHandleRequestAttributesWithDefaults() *GoogleChatCreateOrganizationHandleRequestAttributes {
	this := GoogleChatCreateOrganizationHandleRequestAttributes{}
	return &this
}

// GetName returns the Name field value.
func (o *GoogleChatCreateOrganizationHandleRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GoogleChatCreateOrganizationHandleRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GoogleChatCreateOrganizationHandleRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetSpaceResourceName returns the SpaceResourceName field value.
func (o *GoogleChatCreateOrganizationHandleRequestAttributes) GetSpaceResourceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpaceResourceName
}

// GetSpaceResourceNameOk returns a tuple with the SpaceResourceName field value
// and a boolean to check if the value has been set.
func (o *GoogleChatCreateOrganizationHandleRequestAttributes) GetSpaceResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceResourceName, true
}

// SetSpaceResourceName sets field value.
func (o *GoogleChatCreateOrganizationHandleRequestAttributes) SetSpaceResourceName(v string) {
	o.SpaceResourceName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GoogleChatCreateOrganizationHandleRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["space_resource_name"] = o.SpaceResourceName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GoogleChatCreateOrganizationHandleRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string `json:"name"`
		SpaceResourceName *string `json:"space_resource_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.SpaceResourceName == nil {
		return fmt.Errorf("required field space_resource_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "space_resource_name"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.SpaceResourceName = *all.SpaceResourceName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
