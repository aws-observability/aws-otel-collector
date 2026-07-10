// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomForecastEntry A monthly entry of a custom budget forecast.
type CustomForecastEntry struct {
	// Forecast amount for the month.
	Amount float64 `json:"amount"`
	// Month the custom forecast entry applies to, in `YYYYMM` format.
	Month int64 `json:"month"`
	// Tag filters that scope this custom forecast entry to specific resources.
	TagFilters []CustomForecastEntryTagFilter `json:"tag_filters"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomForecastEntry instantiates a new CustomForecastEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomForecastEntry(amount float64, month int64, tagFilters []CustomForecastEntryTagFilter) *CustomForecastEntry {
	this := CustomForecastEntry{}
	this.Amount = amount
	this.Month = month
	this.TagFilters = tagFilters
	return &this
}

// NewCustomForecastEntryWithDefaults instantiates a new CustomForecastEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomForecastEntryWithDefaults() *CustomForecastEntry {
	this := CustomForecastEntry{}
	return &this
}

// GetAmount returns the Amount field value.
func (o *CustomForecastEntry) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CustomForecastEntry) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value.
func (o *CustomForecastEntry) SetAmount(v float64) {
	o.Amount = v
}

// GetMonth returns the Month field value.
func (o *CustomForecastEntry) GetMonth() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *CustomForecastEntry) GetMonthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value.
func (o *CustomForecastEntry) SetMonth(v int64) {
	o.Month = v
}

// GetTagFilters returns the TagFilters field value.
func (o *CustomForecastEntry) GetTagFilters() []CustomForecastEntryTagFilter {
	if o == nil {
		var ret []CustomForecastEntryTagFilter
		return ret
	}
	return o.TagFilters
}

// GetTagFiltersOk returns a tuple with the TagFilters field value
// and a boolean to check if the value has been set.
func (o *CustomForecastEntry) GetTagFiltersOk() (*[]CustomForecastEntryTagFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagFilters, true
}

// SetTagFilters sets field value.
func (o *CustomForecastEntry) SetTagFilters(v []CustomForecastEntryTagFilter) {
	o.TagFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomForecastEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["amount"] = o.Amount
	toSerialize["month"] = o.Month
	toSerialize["tag_filters"] = o.TagFilters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomForecastEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Amount     *float64                        `json:"amount"`
		Month      *int64                          `json:"month"`
		TagFilters *[]CustomForecastEntryTagFilter `json:"tag_filters"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Amount == nil {
		return fmt.Errorf("required field amount missing")
	}
	if all.Month == nil {
		return fmt.Errorf("required field month missing")
	}
	if all.TagFilters == nil {
		return fmt.Errorf("required field tag_filters missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"amount", "month", "tag_filters"})
	} else {
		return err
	}
	o.Amount = *all.Amount
	o.Month = *all.Month
	o.TagFilters = *all.TagFilters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
