// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertEventCustomAttributesLinksItems A link.
type AlertEventCustomAttributesLinksItems struct {
	// The category of the link.
	Category AlertEventCustomAttributesLinksItemsCategory `json:"category"`
	// The display text of the link. Limited to 300 characters.
	Title *string `json:"title,omitempty"`
	// The URL of the link. Limited to 2048 characters.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAlertEventCustomAttributesLinksItems instantiates a new AlertEventCustomAttributesLinksItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertEventCustomAttributesLinksItems(category AlertEventCustomAttributesLinksItemsCategory, url string) *AlertEventCustomAttributesLinksItems {
	this := AlertEventCustomAttributesLinksItems{}
	this.Category = category
	this.Url = url
	return &this
}

// NewAlertEventCustomAttributesLinksItemsWithDefaults instantiates a new AlertEventCustomAttributesLinksItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertEventCustomAttributesLinksItemsWithDefaults() *AlertEventCustomAttributesLinksItems {
	this := AlertEventCustomAttributesLinksItems{}
	return &this
}

// GetCategory returns the Category field value.
func (o *AlertEventCustomAttributesLinksItems) GetCategory() AlertEventCustomAttributesLinksItemsCategory {
	if o == nil {
		var ret AlertEventCustomAttributesLinksItemsCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *AlertEventCustomAttributesLinksItems) GetCategoryOk() (*AlertEventCustomAttributesLinksItemsCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *AlertEventCustomAttributesLinksItems) SetCategory(v AlertEventCustomAttributesLinksItemsCategory) {
	o.Category = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlertEventCustomAttributesLinksItems) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventCustomAttributesLinksItems) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlertEventCustomAttributesLinksItems) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlertEventCustomAttributesLinksItems) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value.
func (o *AlertEventCustomAttributesLinksItems) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AlertEventCustomAttributesLinksItems) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *AlertEventCustomAttributesLinksItems) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertEventCustomAttributesLinksItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	toSerialize["url"] = o.Url
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertEventCustomAttributesLinksItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category *AlertEventCustomAttributesLinksItemsCategory `json:"category"`
		Title    *string                                       `json:"title,omitempty"`
		Url      *string                                       `json:"url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}

	hasInvalidField := false
	if !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = *all.Category
	}
	o.Title = all.Title
	o.Url = *all.Url

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
