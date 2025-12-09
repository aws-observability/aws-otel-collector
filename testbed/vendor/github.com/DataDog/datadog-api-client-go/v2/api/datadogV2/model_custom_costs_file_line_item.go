// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomCostsFileLineItem Line item details from a Custom Costs file.
type CustomCostsFileLineItem struct {
	// Total cost in the cost file.
	BilledCost *float64 `json:"BilledCost,omitempty"`
	// Currency used in the Custom Costs file.
	BillingCurrency *string `json:"BillingCurrency,omitempty"`
	// Description for the line item cost.
	ChargeDescription *string `json:"ChargeDescription,omitempty"`
	// End date of the usage charge.
	ChargePeriodEnd *string `json:"ChargePeriodEnd,omitempty"`
	// Start date of the usage charge.
	ChargePeriodStart *string `json:"ChargePeriodStart,omitempty"`
	// Name of the provider for the line item.
	ProviderName *string `json:"ProviderName,omitempty"`
	// Additional tags for the line item.
	Tags map[string]string `json:"Tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomCostsFileLineItem instantiates a new CustomCostsFileLineItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomCostsFileLineItem() *CustomCostsFileLineItem {
	this := CustomCostsFileLineItem{}
	return &this
}

// NewCustomCostsFileLineItemWithDefaults instantiates a new CustomCostsFileLineItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomCostsFileLineItemWithDefaults() *CustomCostsFileLineItem {
	this := CustomCostsFileLineItem{}
	return &this
}

// GetBilledCost returns the BilledCost field value if set, zero value otherwise.
func (o *CustomCostsFileLineItem) GetBilledCost() float64 {
	if o == nil || o.BilledCost == nil {
		var ret float64
		return ret
	}
	return *o.BilledCost
}

// GetBilledCostOk returns a tuple with the BilledCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileLineItem) GetBilledCostOk() (*float64, bool) {
	if o == nil || o.BilledCost == nil {
		return nil, false
	}
	return o.BilledCost, true
}

// HasBilledCost returns a boolean if a field has been set.
func (o *CustomCostsFileLineItem) HasBilledCost() bool {
	return o != nil && o.BilledCost != nil
}

// SetBilledCost gets a reference to the given float64 and assigns it to the BilledCost field.
func (o *CustomCostsFileLineItem) SetBilledCost(v float64) {
	o.BilledCost = &v
}

// GetBillingCurrency returns the BillingCurrency field value if set, zero value otherwise.
func (o *CustomCostsFileLineItem) GetBillingCurrency() string {
	if o == nil || o.BillingCurrency == nil {
		var ret string
		return ret
	}
	return *o.BillingCurrency
}

// GetBillingCurrencyOk returns a tuple with the BillingCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileLineItem) GetBillingCurrencyOk() (*string, bool) {
	if o == nil || o.BillingCurrency == nil {
		return nil, false
	}
	return o.BillingCurrency, true
}

// HasBillingCurrency returns a boolean if a field has been set.
func (o *CustomCostsFileLineItem) HasBillingCurrency() bool {
	return o != nil && o.BillingCurrency != nil
}

// SetBillingCurrency gets a reference to the given string and assigns it to the BillingCurrency field.
func (o *CustomCostsFileLineItem) SetBillingCurrency(v string) {
	o.BillingCurrency = &v
}

// GetChargeDescription returns the ChargeDescription field value if set, zero value otherwise.
func (o *CustomCostsFileLineItem) GetChargeDescription() string {
	if o == nil || o.ChargeDescription == nil {
		var ret string
		return ret
	}
	return *o.ChargeDescription
}

// GetChargeDescriptionOk returns a tuple with the ChargeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileLineItem) GetChargeDescriptionOk() (*string, bool) {
	if o == nil || o.ChargeDescription == nil {
		return nil, false
	}
	return o.ChargeDescription, true
}

// HasChargeDescription returns a boolean if a field has been set.
func (o *CustomCostsFileLineItem) HasChargeDescription() bool {
	return o != nil && o.ChargeDescription != nil
}

// SetChargeDescription gets a reference to the given string and assigns it to the ChargeDescription field.
func (o *CustomCostsFileLineItem) SetChargeDescription(v string) {
	o.ChargeDescription = &v
}

// GetChargePeriodEnd returns the ChargePeriodEnd field value if set, zero value otherwise.
func (o *CustomCostsFileLineItem) GetChargePeriodEnd() string {
	if o == nil || o.ChargePeriodEnd == nil {
		var ret string
		return ret
	}
	return *o.ChargePeriodEnd
}

// GetChargePeriodEndOk returns a tuple with the ChargePeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileLineItem) GetChargePeriodEndOk() (*string, bool) {
	if o == nil || o.ChargePeriodEnd == nil {
		return nil, false
	}
	return o.ChargePeriodEnd, true
}

// HasChargePeriodEnd returns a boolean if a field has been set.
func (o *CustomCostsFileLineItem) HasChargePeriodEnd() bool {
	return o != nil && o.ChargePeriodEnd != nil
}

// SetChargePeriodEnd gets a reference to the given string and assigns it to the ChargePeriodEnd field.
func (o *CustomCostsFileLineItem) SetChargePeriodEnd(v string) {
	o.ChargePeriodEnd = &v
}

// GetChargePeriodStart returns the ChargePeriodStart field value if set, zero value otherwise.
func (o *CustomCostsFileLineItem) GetChargePeriodStart() string {
	if o == nil || o.ChargePeriodStart == nil {
		var ret string
		return ret
	}
	return *o.ChargePeriodStart
}

// GetChargePeriodStartOk returns a tuple with the ChargePeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileLineItem) GetChargePeriodStartOk() (*string, bool) {
	if o == nil || o.ChargePeriodStart == nil {
		return nil, false
	}
	return o.ChargePeriodStart, true
}

// HasChargePeriodStart returns a boolean if a field has been set.
func (o *CustomCostsFileLineItem) HasChargePeriodStart() bool {
	return o != nil && o.ChargePeriodStart != nil
}

// SetChargePeriodStart gets a reference to the given string and assigns it to the ChargePeriodStart field.
func (o *CustomCostsFileLineItem) SetChargePeriodStart(v string) {
	o.ChargePeriodStart = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *CustomCostsFileLineItem) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileLineItem) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *CustomCostsFileLineItem) HasProviderName() bool {
	return o != nil && o.ProviderName != nil
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *CustomCostsFileLineItem) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CustomCostsFileLineItem) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileLineItem) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CustomCostsFileLineItem) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *CustomCostsFileLineItem) SetTags(v map[string]string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomCostsFileLineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BilledCost != nil {
		toSerialize["BilledCost"] = o.BilledCost
	}
	if o.BillingCurrency != nil {
		toSerialize["BillingCurrency"] = o.BillingCurrency
	}
	if o.ChargeDescription != nil {
		toSerialize["ChargeDescription"] = o.ChargeDescription
	}
	if o.ChargePeriodEnd != nil {
		toSerialize["ChargePeriodEnd"] = o.ChargePeriodEnd
	}
	if o.ChargePeriodStart != nil {
		toSerialize["ChargePeriodStart"] = o.ChargePeriodStart
	}
	if o.ProviderName != nil {
		toSerialize["ProviderName"] = o.ProviderName
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomCostsFileLineItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BilledCost        *float64          `json:"BilledCost,omitempty"`
		BillingCurrency   *string           `json:"BillingCurrency,omitempty"`
		ChargeDescription *string           `json:"ChargeDescription,omitempty"`
		ChargePeriodEnd   *string           `json:"ChargePeriodEnd,omitempty"`
		ChargePeriodStart *string           `json:"ChargePeriodStart,omitempty"`
		ProviderName      *string           `json:"ProviderName,omitempty"`
		Tags              map[string]string `json:"Tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"BilledCost", "BillingCurrency", "ChargeDescription", "ChargePeriodEnd", "ChargePeriodStart", "ProviderName", "Tags"})
	} else {
		return err
	}
	o.BilledCost = all.BilledCost
	o.BillingCurrency = all.BillingCurrency
	o.ChargeDescription = all.ChargeDescription
	o.ChargePeriodEnd = all.ChargePeriodEnd
	o.ChargePeriodStart = all.ChargePeriodStart
	o.ProviderName = all.ProviderName
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
