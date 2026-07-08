// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SourcemapsListMeta Pagination metadata for the source maps list response.
type SourcemapsListMeta struct {
	// Page information for the source maps list response.
	Page SourcemapsListMetaPage `json:"page"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSourcemapsListMeta instantiates a new SourcemapsListMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourcemapsListMeta(page SourcemapsListMetaPage) *SourcemapsListMeta {
	this := SourcemapsListMeta{}
	this.Page = page
	return &this
}

// NewSourcemapsListMetaWithDefaults instantiates a new SourcemapsListMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourcemapsListMetaWithDefaults() *SourcemapsListMeta {
	this := SourcemapsListMeta{}
	return &this
}

// GetPage returns the Page field value.
func (o *SourcemapsListMeta) GetPage() SourcemapsListMetaPage {
	if o == nil {
		var ret SourcemapsListMetaPage
		return ret
	}
	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *SourcemapsListMeta) GetPageOk() (*SourcemapsListMetaPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value.
func (o *SourcemapsListMeta) SetPage(v SourcemapsListMetaPage) {
	o.Page = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourcemapsListMeta) MarshalJSON() ([]byte, error) {
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
func (o *SourcemapsListMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Page *SourcemapsListMetaPage `json:"page"`
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
