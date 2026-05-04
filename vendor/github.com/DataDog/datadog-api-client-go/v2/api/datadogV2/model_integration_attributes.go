// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAttributes Attributes for an integration.
type IntegrationAttributes struct {
	// List of categories associated with the integration.
	Categories []string `json:"categories"`
	// A description of the integration.
	Description string `json:"description"`
	// Whether the integration is installed.
	Installed bool `json:"installed"`
	// The name of the integration.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationAttributes instantiates a new IntegrationAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationAttributes(categories []string, description string, installed bool, title string) *IntegrationAttributes {
	this := IntegrationAttributes{}
	this.Categories = categories
	this.Description = description
	this.Installed = installed
	this.Title = title
	return &this
}

// NewIntegrationAttributesWithDefaults instantiates a new IntegrationAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationAttributesWithDefaults() *IntegrationAttributes {
	this := IntegrationAttributes{}
	return &this
}

// GetCategories returns the Categories field value.
func (o *IntegrationAttributes) GetCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *IntegrationAttributes) GetCategoriesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Categories, true
}

// SetCategories sets field value.
func (o *IntegrationAttributes) SetCategories(v []string) {
	o.Categories = v
}

// GetDescription returns the Description field value.
func (o *IntegrationAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *IntegrationAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *IntegrationAttributes) SetDescription(v string) {
	o.Description = v
}

// GetInstalled returns the Installed field value.
func (o *IntegrationAttributes) GetInstalled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Installed
}

// GetInstalledOk returns a tuple with the Installed field value
// and a boolean to check if the value has been set.
func (o *IntegrationAttributes) GetInstalledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Installed, true
}

// SetInstalled sets field value.
func (o *IntegrationAttributes) SetInstalled(v bool) {
	o.Installed = v
}

// GetTitle returns the Title field value.
func (o *IntegrationAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *IntegrationAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *IntegrationAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["categories"] = o.Categories
	toSerialize["description"] = o.Description
	toSerialize["installed"] = o.Installed
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Categories  *[]string `json:"categories"`
		Description *string   `json:"description"`
		Installed   *bool     `json:"installed"`
		Title       *string   `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Categories == nil {
		return fmt.Errorf("required field categories missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Installed == nil {
		return fmt.Errorf("required field installed missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"categories", "description", "installed", "title"})
	} else {
		return err
	}
	o.Categories = *all.Categories
	o.Description = *all.Description
	o.Installed = *all.Installed
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
