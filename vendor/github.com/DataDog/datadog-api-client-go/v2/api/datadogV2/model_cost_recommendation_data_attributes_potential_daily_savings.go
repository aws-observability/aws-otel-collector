// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationDataAttributesPotentialDailySavings Estimated daily savings if the recommendation is applied.
type CostRecommendationDataAttributesPotentialDailySavings struct {
	// Numeric amount of the potential daily savings.
	Amount *float64 `json:"amount,omitempty"`
	// ISO 4217 currency code for the savings amount.
	Currency *string `json:"currency,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationDataAttributesPotentialDailySavings instantiates a new CostRecommendationDataAttributesPotentialDailySavings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationDataAttributesPotentialDailySavings() *CostRecommendationDataAttributesPotentialDailySavings {
	this := CostRecommendationDataAttributesPotentialDailySavings{}
	return &this
}

// NewCostRecommendationDataAttributesPotentialDailySavingsWithDefaults instantiates a new CostRecommendationDataAttributesPotentialDailySavings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationDataAttributesPotentialDailySavingsWithDefaults() *CostRecommendationDataAttributesPotentialDailySavings {
	this := CostRecommendationDataAttributesPotentialDailySavings{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CostRecommendationDataAttributesPotentialDailySavings) GetAmount() float64 {
	if o == nil || o.Amount == nil {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationDataAttributesPotentialDailySavings) GetAmountOk() (*float64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CostRecommendationDataAttributesPotentialDailySavings) HasAmount() bool {
	return o != nil && o.Amount != nil
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *CostRecommendationDataAttributesPotentialDailySavings) SetAmount(v float64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CostRecommendationDataAttributesPotentialDailySavings) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationDataAttributesPotentialDailySavings) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CostRecommendationDataAttributesPotentialDailySavings) HasCurrency() bool {
	return o != nil && o.Currency != nil
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CostRecommendationDataAttributesPotentialDailySavings) SetCurrency(v string) {
	o.Currency = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationDataAttributesPotentialDailySavings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationDataAttributesPotentialDailySavings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Amount   *float64 `json:"amount,omitempty"`
		Currency *string  `json:"currency,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"amount", "currency"})
	} else {
		return err
	}
	o.Amount = all.Amount
	o.Currency = all.Currency

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
