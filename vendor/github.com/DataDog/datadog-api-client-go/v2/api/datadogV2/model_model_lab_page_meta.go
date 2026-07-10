// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabPageMeta Pagination metadata for a list response.
type ModelLabPageMeta struct {
	// Pagination details for a list response.
	Page ModelLabPageMetaPage `json:"page"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabPageMeta instantiates a new ModelLabPageMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabPageMeta(page ModelLabPageMetaPage) *ModelLabPageMeta {
	this := ModelLabPageMeta{}
	this.Page = page
	return &this
}

// NewModelLabPageMetaWithDefaults instantiates a new ModelLabPageMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabPageMetaWithDefaults() *ModelLabPageMeta {
	this := ModelLabPageMeta{}
	return &this
}

// GetPage returns the Page field value.
func (o *ModelLabPageMeta) GetPage() ModelLabPageMetaPage {
	if o == nil {
		var ret ModelLabPageMetaPage
		return ret
	}
	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *ModelLabPageMeta) GetPageOk() (*ModelLabPageMetaPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value.
func (o *ModelLabPageMeta) SetPage(v ModelLabPageMetaPage) {
	o.Page = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabPageMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["page"] = o.Page

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModelLabPageMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Page *ModelLabPageMetaPage `json:"page"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Page == nil {
		return fmt.Errorf("required field page missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"page"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = *all.Page

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
