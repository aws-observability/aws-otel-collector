// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetAttributesResponse Dataset metadata and configuration(s).
type DatasetAttributesResponse struct {
	// Timestamp when the dataset was created.
	CreatedAt datadog.NullableTime `json:"created_at,omitempty"`
	// Unique ID of the user who created the dataset.
	CreatedBy *uuid.UUID `json:"created_by,omitempty"`
	// Name of the dataset.
	Name *string `json:"name,omitempty"`
	// List of access principals, formatted as `principal_type:id`. Principal can be 'team' or 'role'.
	Principals []string `json:"principals,omitempty"`
	// List of product-specific filters.
	ProductFilters []FiltersPerProduct `json:"product_filters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatasetAttributesResponse instantiates a new DatasetAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetAttributesResponse() *DatasetAttributesResponse {
	this := DatasetAttributesResponse{}
	return &this
}

// NewDatasetAttributesResponseWithDefaults instantiates a new DatasetAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetAttributesResponseWithDefaults() *DatasetAttributesResponse {
	this := DatasetAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatasetAttributesResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DatasetAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DatasetAttributesResponse) HasCreatedAt() bool {
	return o != nil && o.CreatedAt.IsSet()
}

// SetCreatedAt gets a reference to the given datadog.NullableTime and assigns it to the CreatedAt field.
func (o *DatasetAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil.
func (o *DatasetAttributesResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil.
func (o *DatasetAttributesResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DatasetAttributesResponse) GetCreatedBy() uuid.UUID {
	if o == nil || o.CreatedBy == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetAttributesResponse) GetCreatedByOk() (*uuid.UUID, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DatasetAttributesResponse) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given uuid.UUID and assigns it to the CreatedBy field.
func (o *DatasetAttributesResponse) SetCreatedBy(v uuid.UUID) {
	o.CreatedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DatasetAttributesResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DatasetAttributesResponse) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DatasetAttributesResponse) SetName(v string) {
	o.Name = &v
}

// GetPrincipals returns the Principals field value if set, zero value otherwise.
func (o *DatasetAttributesResponse) GetPrincipals() []string {
	if o == nil || o.Principals == nil {
		var ret []string
		return ret
	}
	return o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetAttributesResponse) GetPrincipalsOk() (*[]string, bool) {
	if o == nil || o.Principals == nil {
		return nil, false
	}
	return &o.Principals, true
}

// HasPrincipals returns a boolean if a field has been set.
func (o *DatasetAttributesResponse) HasPrincipals() bool {
	return o != nil && o.Principals != nil
}

// SetPrincipals gets a reference to the given []string and assigns it to the Principals field.
func (o *DatasetAttributesResponse) SetPrincipals(v []string) {
	o.Principals = v
}

// GetProductFilters returns the ProductFilters field value if set, zero value otherwise.
func (o *DatasetAttributesResponse) GetProductFilters() []FiltersPerProduct {
	if o == nil || o.ProductFilters == nil {
		var ret []FiltersPerProduct
		return ret
	}
	return o.ProductFilters
}

// GetProductFiltersOk returns a tuple with the ProductFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetAttributesResponse) GetProductFiltersOk() (*[]FiltersPerProduct, bool) {
	if o == nil || o.ProductFilters == nil {
		return nil, false
	}
	return &o.ProductFilters, true
}

// HasProductFilters returns a boolean if a field has been set.
func (o *DatasetAttributesResponse) HasProductFilters() bool {
	return o != nil && o.ProductFilters != nil
}

// SetProductFilters gets a reference to the given []FiltersPerProduct and assigns it to the ProductFilters field.
func (o *DatasetAttributesResponse) SetProductFilters(v []FiltersPerProduct) {
	o.ProductFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetAttributesResponse) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Principals != nil {
		toSerialize["principals"] = o.Principals
	}
	if o.ProductFilters != nil {
		toSerialize["product_filters"] = o.ProductFilters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt      datadog.NullableTime `json:"created_at,omitempty"`
		CreatedBy      *uuid.UUID           `json:"created_by,omitempty"`
		Name           *string              `json:"name,omitempty"`
		Principals     []string             `json:"principals,omitempty"`
		ProductFilters []FiltersPerProduct  `json:"product_filters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "name", "principals", "product_filters"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.Name = all.Name
	o.Principals = all.Principals
	o.ProductFilters = all.ProductFilters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
