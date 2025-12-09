// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FiltersPerProduct Product-specific filters for the dataset.
type FiltersPerProduct struct {
	// Defines the list of tag-based filters used to restrict access to telemetry data for a specific product.
	// These filters act as access control rules. Each filter must follow the tag query syntax used by
	// Datadog (such as `@tag.key:value`), and only one tag or attribute may be used to define the access strategy
	// per telemetry type.
	Filters []string `json:"filters"`
	// Name of the product the dataset is for. Possible values are 'apm', 'rum',
	// 'metrics', 'logs', 'error_tracking', and 'cloud_cost'.
	Product string `json:"product"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFiltersPerProduct instantiates a new FiltersPerProduct object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFiltersPerProduct(filters []string, product string) *FiltersPerProduct {
	this := FiltersPerProduct{}
	this.Filters = filters
	this.Product = product
	return &this
}

// NewFiltersPerProductWithDefaults instantiates a new FiltersPerProduct object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFiltersPerProductWithDefaults() *FiltersPerProduct {
	this := FiltersPerProduct{}
	return &this
}

// GetFilters returns the Filters field value.
func (o *FiltersPerProduct) GetFilters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *FiltersPerProduct) GetFiltersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value.
func (o *FiltersPerProduct) SetFilters(v []string) {
	o.Filters = v
}

// GetProduct returns the Product field value.
func (o *FiltersPerProduct) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *FiltersPerProduct) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value.
func (o *FiltersPerProduct) SetProduct(v string) {
	o.Product = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FiltersPerProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["filters"] = o.Filters
	toSerialize["product"] = o.Product

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FiltersPerProduct) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filters *[]string `json:"filters"`
		Product *string   `json:"product"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Filters == nil {
		return fmt.Errorf("required field filters missing")
	}
	if all.Product == nil {
		return fmt.Errorf("required field product missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filters", "product"})
	} else {
		return err
	}
	o.Filters = *all.Filters
	o.Product = *all.Product

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
