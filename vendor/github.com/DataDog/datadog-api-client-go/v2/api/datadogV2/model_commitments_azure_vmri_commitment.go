// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsAzureVMRICommitment Azure Virtual Machine Reserved Instance commitment details.
type CommitmentsAzureVMRICommitment struct {
	// The display name of the Azure reservation.
	BenefitName string `json:"benefit_name"`
	// The unique identifier of the Reserved Instance.
	CommitmentId string `json:"commitment_id"`
	// The expiration date of the commitment.
	ExpirationDate *string `json:"expiration_date,omitempty"`
	// The Azure VM instance type.
	InstanceType string `json:"instance_type"`
	// The Azure meter sub-category for the reservation.
	MeterSubCategory string `json:"meter_sub_category"`
	// The Azure region of the Reserved Instance.
	Region string `json:"region"`
	// The start date of the commitment.
	StartDate *string `json:"start_date,omitempty"`
	// Status of an Azure VM Reserved Instance.
	Status CommitmentsAzureVMRIStatus `json:"status"`
	// The term length in years.
	TermLength *float64 `json:"term_length,omitempty"`
	// The utilization percentage of the commitment.
	Utilization *float64 `json:"utilization,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsAzureVMRICommitment instantiates a new CommitmentsAzureVMRICommitment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsAzureVMRICommitment(benefitName string, commitmentId string, instanceType string, meterSubCategory string, region string, status CommitmentsAzureVMRIStatus) *CommitmentsAzureVMRICommitment {
	this := CommitmentsAzureVMRICommitment{}
	this.BenefitName = benefitName
	this.CommitmentId = commitmentId
	this.InstanceType = instanceType
	this.MeterSubCategory = meterSubCategory
	this.Region = region
	this.Status = status
	return &this
}

// NewCommitmentsAzureVMRICommitmentWithDefaults instantiates a new CommitmentsAzureVMRICommitment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsAzureVMRICommitmentWithDefaults() *CommitmentsAzureVMRICommitment {
	this := CommitmentsAzureVMRICommitment{}
	return &this
}

// GetBenefitName returns the BenefitName field value.
func (o *CommitmentsAzureVMRICommitment) GetBenefitName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BenefitName
}

// GetBenefitNameOk returns a tuple with the BenefitName field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureVMRICommitment) GetBenefitNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BenefitName, true
}

// SetBenefitName sets field value.
func (o *CommitmentsAzureVMRICommitment) SetBenefitName(v string) {
	o.BenefitName = v
}

// GetCommitmentId returns the CommitmentId field value.
func (o *CommitmentsAzureVMRICommitment) GetCommitmentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CommitmentId
}

// GetCommitmentIdOk returns a tuple with the CommitmentId field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureVMRICommitment) GetCommitmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitmentId, true
}

// SetCommitmentId sets field value.
func (o *CommitmentsAzureVMRICommitment) SetCommitmentId(v string) {
	o.CommitmentId = v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *CommitmentsAzureVMRICommitment) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureVMRICommitment) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *CommitmentsAzureVMRICommitment) HasExpirationDate() bool {
	return o != nil && o.ExpirationDate != nil
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *CommitmentsAzureVMRICommitment) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetInstanceType returns the InstanceType field value.
func (o *CommitmentsAzureVMRICommitment) GetInstanceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureVMRICommitment) GetInstanceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value.
func (o *CommitmentsAzureVMRICommitment) SetInstanceType(v string) {
	o.InstanceType = v
}

// GetMeterSubCategory returns the MeterSubCategory field value.
func (o *CommitmentsAzureVMRICommitment) GetMeterSubCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MeterSubCategory
}

// GetMeterSubCategoryOk returns a tuple with the MeterSubCategory field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureVMRICommitment) GetMeterSubCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeterSubCategory, true
}

// SetMeterSubCategory sets field value.
func (o *CommitmentsAzureVMRICommitment) SetMeterSubCategory(v string) {
	o.MeterSubCategory = v
}

// GetRegion returns the Region field value.
func (o *CommitmentsAzureVMRICommitment) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureVMRICommitment) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *CommitmentsAzureVMRICommitment) SetRegion(v string) {
	o.Region = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CommitmentsAzureVMRICommitment) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureVMRICommitment) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CommitmentsAzureVMRICommitment) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CommitmentsAzureVMRICommitment) SetStartDate(v string) {
	o.StartDate = &v
}

// GetStatus returns the Status field value.
func (o *CommitmentsAzureVMRICommitment) GetStatus() CommitmentsAzureVMRIStatus {
	if o == nil {
		var ret CommitmentsAzureVMRIStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureVMRICommitment) GetStatusOk() (*CommitmentsAzureVMRIStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *CommitmentsAzureVMRICommitment) SetStatus(v CommitmentsAzureVMRIStatus) {
	o.Status = v
}

// GetTermLength returns the TermLength field value if set, zero value otherwise.
func (o *CommitmentsAzureVMRICommitment) GetTermLength() float64 {
	if o == nil || o.TermLength == nil {
		var ret float64
		return ret
	}
	return *o.TermLength
}

// GetTermLengthOk returns a tuple with the TermLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureVMRICommitment) GetTermLengthOk() (*float64, bool) {
	if o == nil || o.TermLength == nil {
		return nil, false
	}
	return o.TermLength, true
}

// HasTermLength returns a boolean if a field has been set.
func (o *CommitmentsAzureVMRICommitment) HasTermLength() bool {
	return o != nil && o.TermLength != nil
}

// SetTermLength gets a reference to the given float64 and assigns it to the TermLength field.
func (o *CommitmentsAzureVMRICommitment) SetTermLength(v float64) {
	o.TermLength = &v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *CommitmentsAzureVMRICommitment) GetUtilization() float64 {
	if o == nil || o.Utilization == nil {
		var ret float64
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureVMRICommitment) GetUtilizationOk() (*float64, bool) {
	if o == nil || o.Utilization == nil {
		return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *CommitmentsAzureVMRICommitment) HasUtilization() bool {
	return o != nil && o.Utilization != nil
}

// SetUtilization gets a reference to the given float64 and assigns it to the Utilization field.
func (o *CommitmentsAzureVMRICommitment) SetUtilization(v float64) {
	o.Utilization = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsAzureVMRICommitment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["benefit_name"] = o.BenefitName
	toSerialize["commitment_id"] = o.CommitmentId
	if o.ExpirationDate != nil {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	toSerialize["instance_type"] = o.InstanceType
	toSerialize["meter_sub_category"] = o.MeterSubCategory
	toSerialize["region"] = o.Region
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	toSerialize["status"] = o.Status
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
func (o *CommitmentsAzureVMRICommitment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BenefitName      *string                     `json:"benefit_name"`
		CommitmentId     *string                     `json:"commitment_id"`
		ExpirationDate   *string                     `json:"expiration_date,omitempty"`
		InstanceType     *string                     `json:"instance_type"`
		MeterSubCategory *string                     `json:"meter_sub_category"`
		Region           *string                     `json:"region"`
		StartDate        *string                     `json:"start_date,omitempty"`
		Status           *CommitmentsAzureVMRIStatus `json:"status"`
		TermLength       *float64                    `json:"term_length,omitempty"`
		Utilization      *float64                    `json:"utilization,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BenefitName == nil {
		return fmt.Errorf("required field benefit_name missing")
	}
	if all.CommitmentId == nil {
		return fmt.Errorf("required field commitment_id missing")
	}
	if all.InstanceType == nil {
		return fmt.Errorf("required field instance_type missing")
	}
	if all.MeterSubCategory == nil {
		return fmt.Errorf("required field meter_sub_category missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"benefit_name", "commitment_id", "expiration_date", "instance_type", "meter_sub_category", "region", "start_date", "status", "term_length", "utilization"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BenefitName = *all.BenefitName
	o.CommitmentId = *all.CommitmentId
	o.ExpirationDate = all.ExpirationDate
	o.InstanceType = *all.InstanceType
	o.MeterSubCategory = *all.MeterSubCategory
	o.Region = *all.Region
	o.StartDate = all.StartDate
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.TermLength = all.TermLength
	o.Utilization = all.Utilization

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
