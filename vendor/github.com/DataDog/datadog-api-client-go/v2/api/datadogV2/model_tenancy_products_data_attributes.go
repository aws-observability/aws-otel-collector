// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TenancyProductsDataAttributes Attributes of an OCI tenancy product resource, containing the list of available products and their enablement status.
type TenancyProductsDataAttributes struct {
	// List of Datadog products and their enablement status for the tenancy.
	Products []TenancyProductsDataAttributesProductsItems `json:"products,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTenancyProductsDataAttributes instantiates a new TenancyProductsDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTenancyProductsDataAttributes() *TenancyProductsDataAttributes {
	this := TenancyProductsDataAttributes{}
	return &this
}

// NewTenancyProductsDataAttributesWithDefaults instantiates a new TenancyProductsDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTenancyProductsDataAttributesWithDefaults() *TenancyProductsDataAttributes {
	this := TenancyProductsDataAttributes{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *TenancyProductsDataAttributes) GetProducts() []TenancyProductsDataAttributesProductsItems {
	if o == nil || o.Products == nil {
		var ret []TenancyProductsDataAttributesProductsItems
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyProductsDataAttributes) GetProductsOk() (*[]TenancyProductsDataAttributesProductsItems, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return &o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *TenancyProductsDataAttributes) HasProducts() bool {
	return o != nil && o.Products != nil
}

// SetProducts gets a reference to the given []TenancyProductsDataAttributesProductsItems and assigns it to the Products field.
func (o *TenancyProductsDataAttributes) SetProducts(v []TenancyProductsDataAttributesProductsItems) {
	o.Products = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TenancyProductsDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Products != nil {
		toSerialize["products"] = o.Products
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TenancyProductsDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Products []TenancyProductsDataAttributesProductsItems `json:"products,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"products"})
	} else {
		return err
	}
	o.Products = all.Products

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
