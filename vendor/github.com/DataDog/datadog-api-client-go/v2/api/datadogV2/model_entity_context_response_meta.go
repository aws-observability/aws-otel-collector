// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityContextResponseMeta Metadata returned alongside the entity context response.
type EntityContextResponseMeta struct {
	// Pagination metadata for the entity context response.
	Page EntityContextPage `json:"page"`
	// The total number of entities matching the query, irrespective of pagination.
	TotalCount int32 `json:"total_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityContextResponseMeta instantiates a new EntityContextResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityContextResponseMeta(page EntityContextPage, totalCount int32) *EntityContextResponseMeta {
	this := EntityContextResponseMeta{}
	this.Page = page
	this.TotalCount = totalCount
	return &this
}

// NewEntityContextResponseMetaWithDefaults instantiates a new EntityContextResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityContextResponseMetaWithDefaults() *EntityContextResponseMeta {
	this := EntityContextResponseMeta{}
	return &this
}

// GetPage returns the Page field value.
func (o *EntityContextResponseMeta) GetPage() EntityContextPage {
	if o == nil {
		var ret EntityContextPage
		return ret
	}
	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *EntityContextResponseMeta) GetPageOk() (*EntityContextPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value.
func (o *EntityContextResponseMeta) SetPage(v EntityContextPage) {
	o.Page = v
}

// GetTotalCount returns the TotalCount field value.
func (o *EntityContextResponseMeta) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *EntityContextResponseMeta) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value.
func (o *EntityContextResponseMeta) SetTotalCount(v int32) {
	o.TotalCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityContextResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["page"] = o.Page
	toSerialize["total_count"] = o.TotalCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityContextResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Page       *EntityContextPage `json:"page"`
		TotalCount *int32             `json:"total_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Page == nil {
		return fmt.Errorf("required field page missing")
	}
	if all.TotalCount == nil {
		return fmt.Errorf("required field total_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"page", "total_count"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = *all.Page
	o.TotalCount = *all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
