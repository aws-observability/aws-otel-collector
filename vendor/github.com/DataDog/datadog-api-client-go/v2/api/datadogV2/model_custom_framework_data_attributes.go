// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomFrameworkDataAttributes Framework Data Attributes.
type CustomFrameworkDataAttributes struct {
	// Framework Description
	Description *string `json:"description,omitempty"`
	// Framework Handle
	Handle string `json:"handle"`
	// Framework Icon URL
	IconUrl *string `json:"icon_url,omitempty"`
	// Framework Name
	Name string `json:"name"`
	// Framework Requirements
	Requirements []CustomFrameworkRequirement `json:"requirements"`
	// Framework Version
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomFrameworkDataAttributes instantiates a new CustomFrameworkDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomFrameworkDataAttributes(handle string, name string, requirements []CustomFrameworkRequirement, version string) *CustomFrameworkDataAttributes {
	this := CustomFrameworkDataAttributes{}
	this.Handle = handle
	this.Name = name
	this.Requirements = requirements
	this.Version = version
	return &this
}

// NewCustomFrameworkDataAttributesWithDefaults instantiates a new CustomFrameworkDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomFrameworkDataAttributesWithDefaults() *CustomFrameworkDataAttributes {
	this := CustomFrameworkDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomFrameworkDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFrameworkDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomFrameworkDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomFrameworkDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetHandle returns the Handle field value.
func (o *CustomFrameworkDataAttributes) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *CustomFrameworkDataAttributes) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *CustomFrameworkDataAttributes) SetHandle(v string) {
	o.Handle = v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *CustomFrameworkDataAttributes) GetIconUrl() string {
	if o == nil || o.IconUrl == nil {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFrameworkDataAttributes) GetIconUrlOk() (*string, bool) {
	if o == nil || o.IconUrl == nil {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *CustomFrameworkDataAttributes) HasIconUrl() bool {
	return o != nil && o.IconUrl != nil
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *CustomFrameworkDataAttributes) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetName returns the Name field value.
func (o *CustomFrameworkDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomFrameworkDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CustomFrameworkDataAttributes) SetName(v string) {
	o.Name = v
}

// GetRequirements returns the Requirements field value.
func (o *CustomFrameworkDataAttributes) GetRequirements() []CustomFrameworkRequirement {
	if o == nil {
		var ret []CustomFrameworkRequirement
		return ret
	}
	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value
// and a boolean to check if the value has been set.
func (o *CustomFrameworkDataAttributes) GetRequirementsOk() (*[]CustomFrameworkRequirement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requirements, true
}

// SetRequirements sets field value.
func (o *CustomFrameworkDataAttributes) SetRequirements(v []CustomFrameworkRequirement) {
	o.Requirements = v
}

// GetVersion returns the Version field value.
func (o *CustomFrameworkDataAttributes) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CustomFrameworkDataAttributes) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *CustomFrameworkDataAttributes) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomFrameworkDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["handle"] = o.Handle
	if o.IconUrl != nil {
		toSerialize["icon_url"] = o.IconUrl
	}
	toSerialize["name"] = o.Name
	toSerialize["requirements"] = o.Requirements
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomFrameworkDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description  *string                       `json:"description,omitempty"`
		Handle       *string                       `json:"handle"`
		IconUrl      *string                       `json:"icon_url,omitempty"`
		Name         *string                       `json:"name"`
		Requirements *[]CustomFrameworkRequirement `json:"requirements"`
		Version      *string                       `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Requirements == nil {
		return fmt.Errorf("required field requirements missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "handle", "icon_url", "name", "requirements", "version"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Handle = *all.Handle
	o.IconUrl = all.IconUrl
	o.Name = *all.Name
	o.Requirements = *all.Requirements
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
