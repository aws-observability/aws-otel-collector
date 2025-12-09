// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertEventAttributesLinksItem A link.
type AlertEventAttributesLinksItem struct {
	// The category of the link.
	Category *AlertEventAttributesLinksItemCategory `json:"category,omitempty"`
	// The display text of the link.
	Title *string `json:"title,omitempty"`
	// The URL of the link.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertEventAttributesLinksItem instantiates a new AlertEventAttributesLinksItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertEventAttributesLinksItem() *AlertEventAttributesLinksItem {
	this := AlertEventAttributesLinksItem{}
	return &this
}

// NewAlertEventAttributesLinksItemWithDefaults instantiates a new AlertEventAttributesLinksItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertEventAttributesLinksItemWithDefaults() *AlertEventAttributesLinksItem {
	this := AlertEventAttributesLinksItem{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AlertEventAttributesLinksItem) GetCategory() AlertEventAttributesLinksItemCategory {
	if o == nil || o.Category == nil {
		var ret AlertEventAttributesLinksItemCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventAttributesLinksItem) GetCategoryOk() (*AlertEventAttributesLinksItemCategory, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AlertEventAttributesLinksItem) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given AlertEventAttributesLinksItemCategory and assigns it to the Category field.
func (o *AlertEventAttributesLinksItem) SetCategory(v AlertEventAttributesLinksItemCategory) {
	o.Category = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlertEventAttributesLinksItem) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventAttributesLinksItem) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlertEventAttributesLinksItem) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlertEventAttributesLinksItem) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AlertEventAttributesLinksItem) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventAttributesLinksItem) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AlertEventAttributesLinksItem) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AlertEventAttributesLinksItem) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertEventAttributesLinksItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertEventAttributesLinksItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category *AlertEventAttributesLinksItemCategory `json:"category,omitempty"`
		Title    *string                                `json:"title,omitempty"`
		Url      *string                                `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "title", "url"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Category != nil && !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = all.Category
	}
	o.Title = all.Title
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
