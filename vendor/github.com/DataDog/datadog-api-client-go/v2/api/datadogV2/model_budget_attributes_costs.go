// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BudgetAttributesCosts Aggregated cost data for the budget over the requested period.
type BudgetAttributesCosts struct {
	// The total actual cost. Present only when `actual=true` is requested.
	Actual datadog.NullableFloat64 `json:"actual,omitempty"`
	// The total budgeted amount over the requested period.
	Amount datadog.NullableFloat64 `json:"amount,omitempty"`
	// The total forecast cost, with any custom forecast overrides applied. Present only when `forecast=true` is requested.
	Forecast datadog.NullableFloat64 `json:"forecast,omitempty"`
	// The out-of-the-box ML forecast before custom overrides. Present only when `forecast=true` is requested.
	OotbForecast datadog.NullableFloat64 `json:"ootb_forecast,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBudgetAttributesCosts instantiates a new BudgetAttributesCosts object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBudgetAttributesCosts() *BudgetAttributesCosts {
	this := BudgetAttributesCosts{}
	return &this
}

// NewBudgetAttributesCostsWithDefaults instantiates a new BudgetAttributesCosts object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBudgetAttributesCostsWithDefaults() *BudgetAttributesCosts {
	this := BudgetAttributesCosts{}
	return &this
}

// GetActual returns the Actual field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BudgetAttributesCosts) GetActual() float64 {
	if o == nil || o.Actual.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Actual.Get()
}

// GetActualOk returns a tuple with the Actual field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BudgetAttributesCosts) GetActualOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actual.Get(), o.Actual.IsSet()
}

// HasActual returns a boolean if a field has been set.
func (o *BudgetAttributesCosts) HasActual() bool {
	return o != nil && o.Actual.IsSet()
}

// SetActual gets a reference to the given datadog.NullableFloat64 and assigns it to the Actual field.
func (o *BudgetAttributesCosts) SetActual(v float64) {
	o.Actual.Set(&v)
}

// SetActualNil sets the value for Actual to be an explicit nil.
func (o *BudgetAttributesCosts) SetActualNil() {
	o.Actual.Set(nil)
}

// UnsetActual ensures that no value is present for Actual, not even an explicit nil.
func (o *BudgetAttributesCosts) UnsetActual() {
	o.Actual.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BudgetAttributesCosts) GetAmount() float64 {
	if o == nil || o.Amount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BudgetAttributesCosts) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *BudgetAttributesCosts) HasAmount() bool {
	return o != nil && o.Amount.IsSet()
}

// SetAmount gets a reference to the given datadog.NullableFloat64 and assigns it to the Amount field.
func (o *BudgetAttributesCosts) SetAmount(v float64) {
	o.Amount.Set(&v)
}

// SetAmountNil sets the value for Amount to be an explicit nil.
func (o *BudgetAttributesCosts) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil.
func (o *BudgetAttributesCosts) UnsetAmount() {
	o.Amount.Unset()
}

// GetForecast returns the Forecast field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BudgetAttributesCosts) GetForecast() float64 {
	if o == nil || o.Forecast.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Forecast.Get()
}

// GetForecastOk returns a tuple with the Forecast field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BudgetAttributesCosts) GetForecastOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Forecast.Get(), o.Forecast.IsSet()
}

// HasForecast returns a boolean if a field has been set.
func (o *BudgetAttributesCosts) HasForecast() bool {
	return o != nil && o.Forecast.IsSet()
}

// SetForecast gets a reference to the given datadog.NullableFloat64 and assigns it to the Forecast field.
func (o *BudgetAttributesCosts) SetForecast(v float64) {
	o.Forecast.Set(&v)
}

// SetForecastNil sets the value for Forecast to be an explicit nil.
func (o *BudgetAttributesCosts) SetForecastNil() {
	o.Forecast.Set(nil)
}

// UnsetForecast ensures that no value is present for Forecast, not even an explicit nil.
func (o *BudgetAttributesCosts) UnsetForecast() {
	o.Forecast.Unset()
}

// GetOotbForecast returns the OotbForecast field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BudgetAttributesCosts) GetOotbForecast() float64 {
	if o == nil || o.OotbForecast.Get() == nil {
		var ret float64
		return ret
	}
	return *o.OotbForecast.Get()
}

// GetOotbForecastOk returns a tuple with the OotbForecast field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BudgetAttributesCosts) GetOotbForecastOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OotbForecast.Get(), o.OotbForecast.IsSet()
}

// HasOotbForecast returns a boolean if a field has been set.
func (o *BudgetAttributesCosts) HasOotbForecast() bool {
	return o != nil && o.OotbForecast.IsSet()
}

// SetOotbForecast gets a reference to the given datadog.NullableFloat64 and assigns it to the OotbForecast field.
func (o *BudgetAttributesCosts) SetOotbForecast(v float64) {
	o.OotbForecast.Set(&v)
}

// SetOotbForecastNil sets the value for OotbForecast to be an explicit nil.
func (o *BudgetAttributesCosts) SetOotbForecastNil() {
	o.OotbForecast.Set(nil)
}

// UnsetOotbForecast ensures that no value is present for OotbForecast, not even an explicit nil.
func (o *BudgetAttributesCosts) UnsetOotbForecast() {
	o.OotbForecast.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o BudgetAttributesCosts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Actual.IsSet() {
		toSerialize["actual"] = o.Actual.Get()
	}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.Forecast.IsSet() {
		toSerialize["forecast"] = o.Forecast.Get()
	}
	if o.OotbForecast.IsSet() {
		toSerialize["ootb_forecast"] = o.OotbForecast.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BudgetAttributesCosts) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Actual       datadog.NullableFloat64 `json:"actual,omitempty"`
		Amount       datadog.NullableFloat64 `json:"amount,omitempty"`
		Forecast     datadog.NullableFloat64 `json:"forecast,omitempty"`
		OotbForecast datadog.NullableFloat64 `json:"ootb_forecast,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actual", "amount", "forecast", "ootb_forecast"})
	} else {
		return err
	}
	o.Actual = all.Actual
	o.Amount = all.Amount
	o.Forecast = all.Forecast
	o.OotbForecast = all.OotbForecast

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
