// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsAwsSPCommitment AWS Savings Plan commitment details.
type CommitmentsAwsSPCommitment struct {
	// The unique identifier of the Savings Plan.
	CommitmentId string `json:"commitment_id"`
	// The hourly committed spend for the Savings Plan.
	CommittedSpendPerHour *float64 `json:"committed_spend_per_hour,omitempty"`
	// The expiration date of the commitment.
	ExpirationDate *string `json:"expiration_date,omitempty"`
	// The payment option for the Savings Plan.
	PurchaseOption string `json:"purchase_option"`
	// The Savings Plan type.
	SavingsPlanType string `json:"savings_plan_type"`
	// The start date of the commitment.
	StartDate *string `json:"start_date,omitempty"`
	// The term length in years.
	TermLength *float64 `json:"term_length,omitempty"`
	// The utilization percentage of the commitment.
	Utilization *float64 `json:"utilization,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsAwsSPCommitment instantiates a new CommitmentsAwsSPCommitment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsAwsSPCommitment(commitmentId string, purchaseOption string, savingsPlanType string) *CommitmentsAwsSPCommitment {
	this := CommitmentsAwsSPCommitment{}
	this.CommitmentId = commitmentId
	this.PurchaseOption = purchaseOption
	this.SavingsPlanType = savingsPlanType
	return &this
}

// NewCommitmentsAwsSPCommitmentWithDefaults instantiates a new CommitmentsAwsSPCommitment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsAwsSPCommitmentWithDefaults() *CommitmentsAwsSPCommitment {
	this := CommitmentsAwsSPCommitment{}
	return &this
}

// GetCommitmentId returns the CommitmentId field value.
func (o *CommitmentsAwsSPCommitment) GetCommitmentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CommitmentId
}

// GetCommitmentIdOk returns a tuple with the CommitmentId field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsSPCommitment) GetCommitmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitmentId, true
}

// SetCommitmentId sets field value.
func (o *CommitmentsAwsSPCommitment) SetCommitmentId(v string) {
	o.CommitmentId = v
}

// GetCommittedSpendPerHour returns the CommittedSpendPerHour field value if set, zero value otherwise.
func (o *CommitmentsAwsSPCommitment) GetCommittedSpendPerHour() float64 {
	if o == nil || o.CommittedSpendPerHour == nil {
		var ret float64
		return ret
	}
	return *o.CommittedSpendPerHour
}

// GetCommittedSpendPerHourOk returns a tuple with the CommittedSpendPerHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsSPCommitment) GetCommittedSpendPerHourOk() (*float64, bool) {
	if o == nil || o.CommittedSpendPerHour == nil {
		return nil, false
	}
	return o.CommittedSpendPerHour, true
}

// HasCommittedSpendPerHour returns a boolean if a field has been set.
func (o *CommitmentsAwsSPCommitment) HasCommittedSpendPerHour() bool {
	return o != nil && o.CommittedSpendPerHour != nil
}

// SetCommittedSpendPerHour gets a reference to the given float64 and assigns it to the CommittedSpendPerHour field.
func (o *CommitmentsAwsSPCommitment) SetCommittedSpendPerHour(v float64) {
	o.CommittedSpendPerHour = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *CommitmentsAwsSPCommitment) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsSPCommitment) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *CommitmentsAwsSPCommitment) HasExpirationDate() bool {
	return o != nil && o.ExpirationDate != nil
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *CommitmentsAwsSPCommitment) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetPurchaseOption returns the PurchaseOption field value.
func (o *CommitmentsAwsSPCommitment) GetPurchaseOption() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PurchaseOption
}

// GetPurchaseOptionOk returns a tuple with the PurchaseOption field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsSPCommitment) GetPurchaseOptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOption, true
}

// SetPurchaseOption sets field value.
func (o *CommitmentsAwsSPCommitment) SetPurchaseOption(v string) {
	o.PurchaseOption = v
}

// GetSavingsPlanType returns the SavingsPlanType field value.
func (o *CommitmentsAwsSPCommitment) GetSavingsPlanType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SavingsPlanType
}

// GetSavingsPlanTypeOk returns a tuple with the SavingsPlanType field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsSPCommitment) GetSavingsPlanTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SavingsPlanType, true
}

// SetSavingsPlanType sets field value.
func (o *CommitmentsAwsSPCommitment) SetSavingsPlanType(v string) {
	o.SavingsPlanType = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CommitmentsAwsSPCommitment) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsSPCommitment) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CommitmentsAwsSPCommitment) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CommitmentsAwsSPCommitment) SetStartDate(v string) {
	o.StartDate = &v
}

// GetTermLength returns the TermLength field value if set, zero value otherwise.
func (o *CommitmentsAwsSPCommitment) GetTermLength() float64 {
	if o == nil || o.TermLength == nil {
		var ret float64
		return ret
	}
	return *o.TermLength
}

// GetTermLengthOk returns a tuple with the TermLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsSPCommitment) GetTermLengthOk() (*float64, bool) {
	if o == nil || o.TermLength == nil {
		return nil, false
	}
	return o.TermLength, true
}

// HasTermLength returns a boolean if a field has been set.
func (o *CommitmentsAwsSPCommitment) HasTermLength() bool {
	return o != nil && o.TermLength != nil
}

// SetTermLength gets a reference to the given float64 and assigns it to the TermLength field.
func (o *CommitmentsAwsSPCommitment) SetTermLength(v float64) {
	o.TermLength = &v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *CommitmentsAwsSPCommitment) GetUtilization() float64 {
	if o == nil || o.Utilization == nil {
		var ret float64
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsSPCommitment) GetUtilizationOk() (*float64, bool) {
	if o == nil || o.Utilization == nil {
		return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *CommitmentsAwsSPCommitment) HasUtilization() bool {
	return o != nil && o.Utilization != nil
}

// SetUtilization gets a reference to the given float64 and assigns it to the Utilization field.
func (o *CommitmentsAwsSPCommitment) SetUtilization(v float64) {
	o.Utilization = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsAwsSPCommitment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["commitment_id"] = o.CommitmentId
	if o.CommittedSpendPerHour != nil {
		toSerialize["committed_spend_per_hour"] = o.CommittedSpendPerHour
	}
	if o.ExpirationDate != nil {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	toSerialize["purchase_option"] = o.PurchaseOption
	toSerialize["savings_plan_type"] = o.SavingsPlanType
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.TermLength != nil {
		toSerialize["term_length"] = o.TermLength
	}
	if o.Utilization != nil {
		toSerialize["utilization"] = o.Utilization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsAwsSPCommitment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommitmentId          *string  `json:"commitment_id"`
		CommittedSpendPerHour *float64 `json:"committed_spend_per_hour,omitempty"`
		ExpirationDate        *string  `json:"expiration_date,omitempty"`
		PurchaseOption        *string  `json:"purchase_option"`
		SavingsPlanType       *string  `json:"savings_plan_type"`
		StartDate             *string  `json:"start_date,omitempty"`
		TermLength            *float64 `json:"term_length,omitempty"`
		Utilization           *float64 `json:"utilization,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CommitmentId == nil {
		return fmt.Errorf("required field commitment_id missing")
	}
	if all.PurchaseOption == nil {
		return fmt.Errorf("required field purchase_option missing")
	}
	if all.SavingsPlanType == nil {
		return fmt.Errorf("required field savings_plan_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commitment_id", "committed_spend_per_hour", "expiration_date", "purchase_option", "savings_plan_type", "start_date", "term_length", "utilization"})
	} else {
		return err
	}
	o.CommitmentId = *all.CommitmentId
	o.CommittedSpendPerHour = all.CommittedSpendPerHour
	o.ExpirationDate = all.ExpirationDate
	o.PurchaseOption = *all.PurchaseOption
	o.SavingsPlanType = *all.SavingsPlanType
	o.StartDate = all.StartDate
	o.TermLength = all.TermLength
	o.Utilization = all.Utilization

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
