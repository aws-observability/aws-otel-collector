// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetAttributesRequest Dataset metadata and configurations.
type DatasetAttributesRequest struct {
	// Name of the dataset.
	Name string `json:"name"`
	// List of access principals, formatted as `principal_type:id`. Principal can be 'team' or 'role'.
	Principals []string `json:"principals"`
	// List of product-specific filters.
	ProductFilters []FiltersPerProduct `json:"product_filters"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatasetAttributesRequest instantiates a new DatasetAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetAttributesRequest(name string, principals []string, productFilters []FiltersPerProduct) *DatasetAttributesRequest {
	this := DatasetAttributesRequest{}
	this.Name = name
	this.Principals = principals
	this.ProductFilters = productFilters
	return &this
}

// NewDatasetAttributesRequestWithDefaults instantiates a new DatasetAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetAttributesRequestWithDefaults() *DatasetAttributesRequest {
	this := DatasetAttributesRequest{}
	return &this
}

// GetName returns the Name field value.
func (o *DatasetAttributesRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatasetAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DatasetAttributesRequest) SetName(v string) {
	o.Name = v
}

// GetPrincipals returns the Principals field value.
func (o *DatasetAttributesRequest) GetPrincipals() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value
// and a boolean to check if the value has been set.
func (o *DatasetAttributesRequest) GetPrincipalsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principals, true
}

// SetPrincipals sets field value.
func (o *DatasetAttributesRequest) SetPrincipals(v []string) {
	o.Principals = v
}

// GetProductFilters returns the ProductFilters field value.
func (o *DatasetAttributesRequest) GetProductFilters() []FiltersPerProduct {
	if o == nil {
		var ret []FiltersPerProduct
		return ret
	}
	return o.ProductFilters
}

// GetProductFiltersOk returns a tuple with the ProductFilters field value
// and a boolean to check if the value has been set.
func (o *DatasetAttributesRequest) GetProductFiltersOk() (*[]FiltersPerProduct, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductFilters, true
}

// SetProductFilters sets field value.
func (o *DatasetAttributesRequest) SetProductFilters(v []FiltersPerProduct) {
	o.ProductFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["principals"] = o.Principals
	toSerialize["product_filters"] = o.ProductFilters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string              `json:"name"`
		Principals     *[]string            `json:"principals"`
		ProductFilters *[]FiltersPerProduct `json:"product_filters"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Principals == nil {
		return fmt.Errorf("required field principals missing")
	}
	if all.ProductFilters == nil {
		return fmt.Errorf("required field product_filters missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "principals", "product_filters"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Principals = *all.Principals
	o.ProductFilters = *all.ProductFilters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
