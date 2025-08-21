// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetAttributes Dataset metadata and configuration(s).
type DatasetAttributes struct {
	// Timestamp when the dataset was created.
	CreatedAt datadog.NullableTime `json:"created_at,omitempty"`
	// Unique ID of the user who created the dataset.
	CreatedBy *uuid.UUID `json:"created_by,omitempty"`
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

// NewDatasetAttributes instantiates a new DatasetAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetAttributes(name string, principals []string, productFilters []FiltersPerProduct) *DatasetAttributes {
	this := DatasetAttributes{}
	this.Name = name
	this.Principals = principals
	this.ProductFilters = productFilters
	return &this
}

// NewDatasetAttributesWithDefaults instantiates a new DatasetAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetAttributesWithDefaults() *DatasetAttributes {
	this := DatasetAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatasetAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DatasetAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DatasetAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt.IsSet()
}

// SetCreatedAt gets a reference to the given datadog.NullableTime and assigns it to the CreatedAt field.
func (o *DatasetAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil.
func (o *DatasetAttributes) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil.
func (o *DatasetAttributes) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DatasetAttributes) GetCreatedBy() uuid.UUID {
	if o == nil || o.CreatedBy == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetAttributes) GetCreatedByOk() (*uuid.UUID, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DatasetAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given uuid.UUID and assigns it to the CreatedBy field.
func (o *DatasetAttributes) SetCreatedBy(v uuid.UUID) {
	o.CreatedBy = &v
}

// GetName returns the Name field value.
func (o *DatasetAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatasetAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DatasetAttributes) SetName(v string) {
	o.Name = v
}

// GetPrincipals returns the Principals field value.
func (o *DatasetAttributes) GetPrincipals() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value
// and a boolean to check if the value has been set.
func (o *DatasetAttributes) GetPrincipalsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principals, true
}

// SetPrincipals sets field value.
func (o *DatasetAttributes) SetPrincipals(v []string) {
	o.Principals = v
}

// GetProductFilters returns the ProductFilters field value.
func (o *DatasetAttributes) GetProductFilters() []FiltersPerProduct {
	if o == nil {
		var ret []FiltersPerProduct
		return ret
	}
	return o.ProductFilters
}

// GetProductFiltersOk returns a tuple with the ProductFilters field value
// and a boolean to check if the value has been set.
func (o *DatasetAttributes) GetProductFiltersOk() (*[]FiltersPerProduct, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductFilters, true
}

// SetProductFilters sets field value.
func (o *DatasetAttributes) SetProductFilters(v []FiltersPerProduct) {
	o.ProductFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
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
func (o *DatasetAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt      datadog.NullableTime `json:"created_at,omitempty"`
		CreatedBy      *uuid.UUID           `json:"created_by,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "name", "principals", "product_filters"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.Name = *all.Name
	o.Principals = *all.Principals
	o.ProductFilters = *all.ProductFilters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
