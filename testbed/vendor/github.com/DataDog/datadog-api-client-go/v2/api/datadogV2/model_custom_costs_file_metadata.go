// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomCostsFileMetadata Schema of a Custom Costs metadata.
type CustomCostsFileMetadata struct {
	// Total cost in the cost file.
	BilledCost *float64 `json:"billed_cost,omitempty"`
	// Currency used in the Custom Costs file.
	BillingCurrency *string `json:"billing_currency,omitempty"`
	// Usage charge period of a Custom Costs file.
	ChargePeriod *CustomCostsFileUsageChargePeriod `json:"charge_period,omitempty"`
	// Name of the Custom Costs file.
	Name *string `json:"name,omitempty"`
	// Providers contained in the Custom Costs file.
	ProviderNames []string `json:"provider_names,omitempty"`
	// Status of the Custom Costs file.
	Status *string `json:"status,omitempty"`
	// Timestamp, in millisecond, of the upload time of the Custom Costs file.
	UploadedAt *float64 `json:"uploaded_at,omitempty"`
	// Metadata of the user that has uploaded the Custom Costs file.
	UploadedBy *CustomCostsUser `json:"uploaded_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomCostsFileMetadata instantiates a new CustomCostsFileMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomCostsFileMetadata() *CustomCostsFileMetadata {
	this := CustomCostsFileMetadata{}
	return &this
}

// NewCustomCostsFileMetadataWithDefaults instantiates a new CustomCostsFileMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomCostsFileMetadataWithDefaults() *CustomCostsFileMetadata {
	this := CustomCostsFileMetadata{}
	return &this
}

// GetBilledCost returns the BilledCost field value if set, zero value otherwise.
func (o *CustomCostsFileMetadata) GetBilledCost() float64 {
	if o == nil || o.BilledCost == nil {
		var ret float64
		return ret
	}
	return *o.BilledCost
}

// GetBilledCostOk returns a tuple with the BilledCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileMetadata) GetBilledCostOk() (*float64, bool) {
	if o == nil || o.BilledCost == nil {
		return nil, false
	}
	return o.BilledCost, true
}

// HasBilledCost returns a boolean if a field has been set.
func (o *CustomCostsFileMetadata) HasBilledCost() bool {
	return o != nil && o.BilledCost != nil
}

// SetBilledCost gets a reference to the given float64 and assigns it to the BilledCost field.
func (o *CustomCostsFileMetadata) SetBilledCost(v float64) {
	o.BilledCost = &v
}

// GetBillingCurrency returns the BillingCurrency field value if set, zero value otherwise.
func (o *CustomCostsFileMetadata) GetBillingCurrency() string {
	if o == nil || o.BillingCurrency == nil {
		var ret string
		return ret
	}
	return *o.BillingCurrency
}

// GetBillingCurrencyOk returns a tuple with the BillingCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileMetadata) GetBillingCurrencyOk() (*string, bool) {
	if o == nil || o.BillingCurrency == nil {
		return nil, false
	}
	return o.BillingCurrency, true
}

// HasBillingCurrency returns a boolean if a field has been set.
func (o *CustomCostsFileMetadata) HasBillingCurrency() bool {
	return o != nil && o.BillingCurrency != nil
}

// SetBillingCurrency gets a reference to the given string and assigns it to the BillingCurrency field.
func (o *CustomCostsFileMetadata) SetBillingCurrency(v string) {
	o.BillingCurrency = &v
}

// GetChargePeriod returns the ChargePeriod field value if set, zero value otherwise.
func (o *CustomCostsFileMetadata) GetChargePeriod() CustomCostsFileUsageChargePeriod {
	if o == nil || o.ChargePeriod == nil {
		var ret CustomCostsFileUsageChargePeriod
		return ret
	}
	return *o.ChargePeriod
}

// GetChargePeriodOk returns a tuple with the ChargePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileMetadata) GetChargePeriodOk() (*CustomCostsFileUsageChargePeriod, bool) {
	if o == nil || o.ChargePeriod == nil {
		return nil, false
	}
	return o.ChargePeriod, true
}

// HasChargePeriod returns a boolean if a field has been set.
func (o *CustomCostsFileMetadata) HasChargePeriod() bool {
	return o != nil && o.ChargePeriod != nil
}

// SetChargePeriod gets a reference to the given CustomCostsFileUsageChargePeriod and assigns it to the ChargePeriod field.
func (o *CustomCostsFileMetadata) SetChargePeriod(v CustomCostsFileUsageChargePeriod) {
	o.ChargePeriod = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomCostsFileMetadata) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileMetadata) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomCostsFileMetadata) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomCostsFileMetadata) SetName(v string) {
	o.Name = &v
}

// GetProviderNames returns the ProviderNames field value if set, zero value otherwise.
func (o *CustomCostsFileMetadata) GetProviderNames() []string {
	if o == nil || o.ProviderNames == nil {
		var ret []string
		return ret
	}
	return o.ProviderNames
}

// GetProviderNamesOk returns a tuple with the ProviderNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileMetadata) GetProviderNamesOk() (*[]string, bool) {
	if o == nil || o.ProviderNames == nil {
		return nil, false
	}
	return &o.ProviderNames, true
}

// HasProviderNames returns a boolean if a field has been set.
func (o *CustomCostsFileMetadata) HasProviderNames() bool {
	return o != nil && o.ProviderNames != nil
}

// SetProviderNames gets a reference to the given []string and assigns it to the ProviderNames field.
func (o *CustomCostsFileMetadata) SetProviderNames(v []string) {
	o.ProviderNames = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CustomCostsFileMetadata) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileMetadata) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomCostsFileMetadata) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CustomCostsFileMetadata) SetStatus(v string) {
	o.Status = &v
}

// GetUploadedAt returns the UploadedAt field value if set, zero value otherwise.
func (o *CustomCostsFileMetadata) GetUploadedAt() float64 {
	if o == nil || o.UploadedAt == nil {
		var ret float64
		return ret
	}
	return *o.UploadedAt
}

// GetUploadedAtOk returns a tuple with the UploadedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileMetadata) GetUploadedAtOk() (*float64, bool) {
	if o == nil || o.UploadedAt == nil {
		return nil, false
	}
	return o.UploadedAt, true
}

// HasUploadedAt returns a boolean if a field has been set.
func (o *CustomCostsFileMetadata) HasUploadedAt() bool {
	return o != nil && o.UploadedAt != nil
}

// SetUploadedAt gets a reference to the given float64 and assigns it to the UploadedAt field.
func (o *CustomCostsFileMetadata) SetUploadedAt(v float64) {
	o.UploadedAt = &v
}

// GetUploadedBy returns the UploadedBy field value if set, zero value otherwise.
func (o *CustomCostsFileMetadata) GetUploadedBy() CustomCostsUser {
	if o == nil || o.UploadedBy == nil {
		var ret CustomCostsUser
		return ret
	}
	return *o.UploadedBy
}

// GetUploadedByOk returns a tuple with the UploadedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileMetadata) GetUploadedByOk() (*CustomCostsUser, bool) {
	if o == nil || o.UploadedBy == nil {
		return nil, false
	}
	return o.UploadedBy, true
}

// HasUploadedBy returns a boolean if a field has been set.
func (o *CustomCostsFileMetadata) HasUploadedBy() bool {
	return o != nil && o.UploadedBy != nil
}

// SetUploadedBy gets a reference to the given CustomCostsUser and assigns it to the UploadedBy field.
func (o *CustomCostsFileMetadata) SetUploadedBy(v CustomCostsUser) {
	o.UploadedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomCostsFileMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BilledCost != nil {
		toSerialize["billed_cost"] = o.BilledCost
	}
	if o.BillingCurrency != nil {
		toSerialize["billing_currency"] = o.BillingCurrency
	}
	if o.ChargePeriod != nil {
		toSerialize["charge_period"] = o.ChargePeriod
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProviderNames != nil {
		toSerialize["provider_names"] = o.ProviderNames
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UploadedAt != nil {
		toSerialize["uploaded_at"] = o.UploadedAt
	}
	if o.UploadedBy != nil {
		toSerialize["uploaded_by"] = o.UploadedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomCostsFileMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BilledCost      *float64                          `json:"billed_cost,omitempty"`
		BillingCurrency *string                           `json:"billing_currency,omitempty"`
		ChargePeriod    *CustomCostsFileUsageChargePeriod `json:"charge_period,omitempty"`
		Name            *string                           `json:"name,omitempty"`
		ProviderNames   []string                          `json:"provider_names,omitempty"`
		Status          *string                           `json:"status,omitempty"`
		UploadedAt      *float64                          `json:"uploaded_at,omitempty"`
		UploadedBy      *CustomCostsUser                  `json:"uploaded_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"billed_cost", "billing_currency", "charge_period", "name", "provider_names", "status", "uploaded_at", "uploaded_by"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BilledCost = all.BilledCost
	o.BillingCurrency = all.BillingCurrency
	if all.ChargePeriod != nil && all.ChargePeriod.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ChargePeriod = all.ChargePeriod
	o.Name = all.Name
	o.ProviderNames = all.ProviderNames
	o.Status = all.Status
	o.UploadedAt = all.UploadedAt
	if all.UploadedBy != nil && all.UploadedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UploadedBy = all.UploadedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
