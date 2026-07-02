// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmHostFacetInfoAttributes Attributes of a facet info response, containing the value distribution for the requested facet.
type CsmHostFacetInfoAttributes struct {
	// The list of facet value entries for the current page.
	Items []CsmHostFacetInfoItem `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCsmHostFacetInfoAttributes instantiates a new CsmHostFacetInfoAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCsmHostFacetInfoAttributes(items []CsmHostFacetInfoItem) *CsmHostFacetInfoAttributes {
	this := CsmHostFacetInfoAttributes{}
	this.Items = items
	return &this
}

// NewCsmHostFacetInfoAttributesWithDefaults instantiates a new CsmHostFacetInfoAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCsmHostFacetInfoAttributesWithDefaults() *CsmHostFacetInfoAttributes {
	this := CsmHostFacetInfoAttributes{}
	return &this
}

// GetItems returns the Items field value.
func (o *CsmHostFacetInfoAttributes) GetItems() []CsmHostFacetInfoItem {
	if o == nil {
		var ret []CsmHostFacetInfoItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CsmHostFacetInfoAttributes) GetItemsOk() (*[]CsmHostFacetInfoItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *CsmHostFacetInfoAttributes) SetItems(v []CsmHostFacetInfoItem) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CsmHostFacetInfoAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CsmHostFacetInfoAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items *[]CsmHostFacetInfoItem `json:"items"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"items"})
	} else {
		return err
	}
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
