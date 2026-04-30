// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BudgetWithEntriesDataAttributesEntriesItems A single monthly budget entry defining the allocated amount and optional tag filters for a specific month.
type BudgetWithEntriesDataAttributesEntriesItems struct {
	// The budgeted amount for this entry.
	Amount *float64 `json:"amount,omitempty"`
	// The month this budget entry applies to, in YYYYMM format.
	Month *int64 `json:"month,omitempty"`
	// The list of tag filters that scope this budget entry to specific resources.
	TagFilters []BudgetWithEntriesDataAttributesEntriesItemsTagFiltersItems `json:"tag_filters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBudgetWithEntriesDataAttributesEntriesItems instantiates a new BudgetWithEntriesDataAttributesEntriesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBudgetWithEntriesDataAttributesEntriesItems() *BudgetWithEntriesDataAttributesEntriesItems {
	this := BudgetWithEntriesDataAttributesEntriesItems{}
	return &this
}

// NewBudgetWithEntriesDataAttributesEntriesItemsWithDefaults instantiates a new BudgetWithEntriesDataAttributesEntriesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBudgetWithEntriesDataAttributesEntriesItemsWithDefaults() *BudgetWithEntriesDataAttributesEntriesItems {
	this := BudgetWithEntriesDataAttributesEntriesItems{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BudgetWithEntriesDataAttributesEntriesItems) GetAmount() float64 {
	if o == nil || o.Amount == nil {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesDataAttributesEntriesItems) GetAmountOk() (*float64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BudgetWithEntriesDataAttributesEntriesItems) HasAmount() bool {
	return o != nil && o.Amount != nil
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *BudgetWithEntriesDataAttributesEntriesItems) SetAmount(v float64) {
	o.Amount = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *BudgetWithEntriesDataAttributesEntriesItems) GetMonth() int64 {
	if o == nil || o.Month == nil {
		var ret int64
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesDataAttributesEntriesItems) GetMonthOk() (*int64, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *BudgetWithEntriesDataAttributesEntriesItems) HasMonth() bool {
	return o != nil && o.Month != nil
}

// SetMonth gets a reference to the given int64 and assigns it to the Month field.
func (o *BudgetWithEntriesDataAttributesEntriesItems) SetMonth(v int64) {
	o.Month = &v
}

// GetTagFilters returns the TagFilters field value if set, zero value otherwise.
func (o *BudgetWithEntriesDataAttributesEntriesItems) GetTagFilters() []BudgetWithEntriesDataAttributesEntriesItemsTagFiltersItems {
	if o == nil || o.TagFilters == nil {
		var ret []BudgetWithEntriesDataAttributesEntriesItemsTagFiltersItems
		return ret
	}
	return o.TagFilters
}

// GetTagFiltersOk returns a tuple with the TagFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesDataAttributesEntriesItems) GetTagFiltersOk() (*[]BudgetWithEntriesDataAttributesEntriesItemsTagFiltersItems, bool) {
	if o == nil || o.TagFilters == nil {
		return nil, false
	}
	return &o.TagFilters, true
}

// HasTagFilters returns a boolean if a field has been set.
func (o *BudgetWithEntriesDataAttributesEntriesItems) HasTagFilters() bool {
	return o != nil && o.TagFilters != nil
}

// SetTagFilters gets a reference to the given []BudgetWithEntriesDataAttributesEntriesItemsTagFiltersItems and assigns it to the TagFilters field.
func (o *BudgetWithEntriesDataAttributesEntriesItems) SetTagFilters(v []BudgetWithEntriesDataAttributesEntriesItemsTagFiltersItems) {
	o.TagFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BudgetWithEntriesDataAttributesEntriesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Month != nil {
		toSerialize["month"] = o.Month
	}
	if o.TagFilters != nil {
		toSerialize["tag_filters"] = o.TagFilters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BudgetWithEntriesDataAttributesEntriesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Amount     *float64                                                     `json:"amount,omitempty"`
		Month      *int64                                                       `json:"month,omitempty"`
		TagFilters []BudgetWithEntriesDataAttributesEntriesItemsTagFiltersItems `json:"tag_filters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"amount", "month", "tag_filters"})
	} else {
		return err
	}
	o.Amount = all.Amount
	o.Month = all.Month
	o.TagFilters = all.TagFilters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
