// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerGroupAttributes Attributes of the Sensitive Data Scanner group.
type SensitiveDataScannerGroupAttributes struct {
	// Description of the group.
	Description *string `json:"description,omitempty"`
	// Filter for the Scanning Group.
	Filter *SensitiveDataScannerFilter `json:"filter,omitempty"`
	// Whether or not the group is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the group.
	Name *string `json:"name,omitempty"`
	// List of products the scanning group applies.
	ProductList []SensitiveDataScannerProduct `json:"product_list,omitempty"`
	// List of sampling rates per product type.
	Samplings []SensitiveDataScannerSamplings `json:"samplings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSensitiveDataScannerGroupAttributes instantiates a new SensitiveDataScannerGroupAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerGroupAttributes() *SensitiveDataScannerGroupAttributes {
	this := SensitiveDataScannerGroupAttributes{}
	return &this
}

// NewSensitiveDataScannerGroupAttributesWithDefaults instantiates a new SensitiveDataScannerGroupAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerGroupAttributesWithDefaults() *SensitiveDataScannerGroupAttributes {
	this := SensitiveDataScannerGroupAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SensitiveDataScannerGroupAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerGroupAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SensitiveDataScannerGroupAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SensitiveDataScannerGroupAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *SensitiveDataScannerGroupAttributes) GetFilter() SensitiveDataScannerFilter {
	if o == nil || o.Filter == nil {
		var ret SensitiveDataScannerFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerGroupAttributes) GetFilterOk() (*SensitiveDataScannerFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *SensitiveDataScannerGroupAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given SensitiveDataScannerFilter and assigns it to the Filter field.
func (o *SensitiveDataScannerGroupAttributes) SetFilter(v SensitiveDataScannerFilter) {
	o.Filter = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *SensitiveDataScannerGroupAttributes) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerGroupAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *SensitiveDataScannerGroupAttributes) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *SensitiveDataScannerGroupAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SensitiveDataScannerGroupAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerGroupAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SensitiveDataScannerGroupAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SensitiveDataScannerGroupAttributes) SetName(v string) {
	o.Name = &v
}

// GetProductList returns the ProductList field value if set, zero value otherwise.
func (o *SensitiveDataScannerGroupAttributes) GetProductList() []SensitiveDataScannerProduct {
	if o == nil || o.ProductList == nil {
		var ret []SensitiveDataScannerProduct
		return ret
	}
	return o.ProductList
}

// GetProductListOk returns a tuple with the ProductList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerGroupAttributes) GetProductListOk() (*[]SensitiveDataScannerProduct, bool) {
	if o == nil || o.ProductList == nil {
		return nil, false
	}
	return &o.ProductList, true
}

// HasProductList returns a boolean if a field has been set.
func (o *SensitiveDataScannerGroupAttributes) HasProductList() bool {
	return o != nil && o.ProductList != nil
}

// SetProductList gets a reference to the given []SensitiveDataScannerProduct and assigns it to the ProductList field.
func (o *SensitiveDataScannerGroupAttributes) SetProductList(v []SensitiveDataScannerProduct) {
	o.ProductList = v
}

// GetSamplings returns the Samplings field value if set, zero value otherwise.
func (o *SensitiveDataScannerGroupAttributes) GetSamplings() []SensitiveDataScannerSamplings {
	if o == nil || o.Samplings == nil {
		var ret []SensitiveDataScannerSamplings
		return ret
	}
	return o.Samplings
}

// GetSamplingsOk returns a tuple with the Samplings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerGroupAttributes) GetSamplingsOk() (*[]SensitiveDataScannerSamplings, bool) {
	if o == nil || o.Samplings == nil {
		return nil, false
	}
	return &o.Samplings, true
}

// HasSamplings returns a boolean if a field has been set.
func (o *SensitiveDataScannerGroupAttributes) HasSamplings() bool {
	return o != nil && o.Samplings != nil
}

// SetSamplings gets a reference to the given []SensitiveDataScannerSamplings and assigns it to the Samplings field.
func (o *SensitiveDataScannerGroupAttributes) SetSamplings(v []SensitiveDataScannerSamplings) {
	o.Samplings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerGroupAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProductList != nil {
		toSerialize["product_list"] = o.ProductList
	}
	if o.Samplings != nil {
		toSerialize["samplings"] = o.Samplings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerGroupAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                         `json:"description,omitempty"`
		Filter      *SensitiveDataScannerFilter     `json:"filter,omitempty"`
		IsEnabled   *bool                           `json:"is_enabled,omitempty"`
		Name        *string                         `json:"name,omitempty"`
		ProductList []SensitiveDataScannerProduct   `json:"product_list,omitempty"`
		Samplings   []SensitiveDataScannerSamplings `json:"samplings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "filter", "is_enabled", "name", "product_list", "samplings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	if all.Filter != nil && all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = all.Filter
	o.IsEnabled = all.IsEnabled
	o.Name = all.Name
	o.ProductList = all.ProductList
	o.Samplings = all.Samplings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
