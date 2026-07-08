// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsAwsEC2RICommitment AWS EC2 Reserved Instance commitment details.
type CommitmentsAwsEC2RICommitment struct {
	// The availability zone of the reservation.
	AvailabilityZone *string `json:"availability_zone,omitempty"`
	// The unique identifier of the Reserved Instance.
	CommitmentId string `json:"commitment_id"`
	// The expiration date of the commitment.
	ExpirationDate *string `json:"expiration_date,omitempty"`
	// The EC2 instance type.
	InstanceType string `json:"instance_type"`
	// The number of Normalized Capacity Units.
	NumberOfNfus *float64 `json:"number_of_nfus,omitempty"`
	// The number of reserved instances.
	NumberOfReservations *float64 `json:"number_of_reservations,omitempty"`
	// The offering class of the Reserved Instance.
	OfferingClass string `json:"offering_class"`
	// The operating system of the Reserved Instance.
	OperatingSystem string `json:"operating_system"`
	// The payment option for the Reserved Instance.
	PurchaseOption string `json:"purchase_option"`
	// The AWS region of the Reserved Instance.
	Region string `json:"region"`
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

// NewCommitmentsAwsEC2RICommitment instantiates a new CommitmentsAwsEC2RICommitment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsAwsEC2RICommitment(commitmentId string, instanceType string, offeringClass string, operatingSystem string, purchaseOption string, region string) *CommitmentsAwsEC2RICommitment {
	this := CommitmentsAwsEC2RICommitment{}
	this.CommitmentId = commitmentId
	this.InstanceType = instanceType
	this.OfferingClass = offeringClass
	this.OperatingSystem = operatingSystem
	this.PurchaseOption = purchaseOption
	this.Region = region
	return &this
}

// NewCommitmentsAwsEC2RICommitmentWithDefaults instantiates a new CommitmentsAwsEC2RICommitment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsAwsEC2RICommitmentWithDefaults() *CommitmentsAwsEC2RICommitment {
	this := CommitmentsAwsEC2RICommitment{}
	return &this
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *CommitmentsAwsEC2RICommitment) GetAvailabilityZone() string {
	if o == nil || o.AvailabilityZone == nil {
		var ret string
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsEC2RICommitment) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil || o.AvailabilityZone == nil {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *CommitmentsAwsEC2RICommitment) HasAvailabilityZone() bool {
	return o != nil && o.AvailabilityZone != nil
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *CommitmentsAwsEC2RICommitment) SetAvailabilityZone(v string) {
	o.AvailabilityZone = &v
}

// GetCommitmentId returns the CommitmentId field value.
func (o *CommitmentsAwsEC2RICommitment) GetCommitmentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CommitmentId
}

// GetCommitmentIdOk returns a tuple with the CommitmentId field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsEC2RICommitment) GetCommitmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitmentId, true
}

// SetCommitmentId sets field value.
func (o *CommitmentsAwsEC2RICommitment) SetCommitmentId(v string) {
	o.CommitmentId = v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *CommitmentsAwsEC2RICommitment) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsEC2RICommitment) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *CommitmentsAwsEC2RICommitment) HasExpirationDate() bool {
	return o != nil && o.ExpirationDate != nil
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *CommitmentsAwsEC2RICommitment) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetInstanceType returns the InstanceType field value.
func (o *CommitmentsAwsEC2RICommitment) GetInstanceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsEC2RICommitment) GetInstanceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value.
func (o *CommitmentsAwsEC2RICommitment) SetInstanceType(v string) {
	o.InstanceType = v
}

// GetNumberOfNfus returns the NumberOfNfus field value if set, zero value otherwise.
func (o *CommitmentsAwsEC2RICommitment) GetNumberOfNfus() float64 {
	if o == nil || o.NumberOfNfus == nil {
		var ret float64
		return ret
	}
	return *o.NumberOfNfus
}

// GetNumberOfNfusOk returns a tuple with the NumberOfNfus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsEC2RICommitment) GetNumberOfNfusOk() (*float64, bool) {
	if o == nil || o.NumberOfNfus == nil {
		return nil, false
	}
	return o.NumberOfNfus, true
}

// HasNumberOfNfus returns a boolean if a field has been set.
func (o *CommitmentsAwsEC2RICommitment) HasNumberOfNfus() bool {
	return o != nil && o.NumberOfNfus != nil
}

// SetNumberOfNfus gets a reference to the given float64 and assigns it to the NumberOfNfus field.
func (o *CommitmentsAwsEC2RICommitment) SetNumberOfNfus(v float64) {
	o.NumberOfNfus = &v
}

// GetNumberOfReservations returns the NumberOfReservations field value if set, zero value otherwise.
func (o *CommitmentsAwsEC2RICommitment) GetNumberOfReservations() float64 {
	if o == nil || o.NumberOfReservations == nil {
		var ret float64
		return ret
	}
	return *o.NumberOfReservations
}

// GetNumberOfReservationsOk returns a tuple with the NumberOfReservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsEC2RICommitment) GetNumberOfReservationsOk() (*float64, bool) {
	if o == nil || o.NumberOfReservations == nil {
		return nil, false
	}
	return o.NumberOfReservations, true
}

// HasNumberOfReservations returns a boolean if a field has been set.
func (o *CommitmentsAwsEC2RICommitment) HasNumberOfReservations() bool {
	return o != nil && o.NumberOfReservations != nil
}

// SetNumberOfReservations gets a reference to the given float64 and assigns it to the NumberOfReservations field.
func (o *CommitmentsAwsEC2RICommitment) SetNumberOfReservations(v float64) {
	o.NumberOfReservations = &v
}

// GetOfferingClass returns the OfferingClass field value.
func (o *CommitmentsAwsEC2RICommitment) GetOfferingClass() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OfferingClass
}

// GetOfferingClassOk returns a tuple with the OfferingClass field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsEC2RICommitment) GetOfferingClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferingClass, true
}

// SetOfferingClass sets field value.
func (o *CommitmentsAwsEC2RICommitment) SetOfferingClass(v string) {
	o.OfferingClass = v
}

// GetOperatingSystem returns the OperatingSystem field value.
func (o *CommitmentsAwsEC2RICommitment) GetOperatingSystem() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsEC2RICommitment) GetOperatingSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperatingSystem, true
}

// SetOperatingSystem sets field value.
func (o *CommitmentsAwsEC2RICommitment) SetOperatingSystem(v string) {
	o.OperatingSystem = v
}

// GetPurchaseOption returns the PurchaseOption field value.
func (o *CommitmentsAwsEC2RICommitment) GetPurchaseOption() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PurchaseOption
}

// GetPurchaseOptionOk returns a tuple with the PurchaseOption field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsEC2RICommitment) GetPurchaseOptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOption, true
}

// SetPurchaseOption sets field value.
func (o *CommitmentsAwsEC2RICommitment) SetPurchaseOption(v string) {
	o.PurchaseOption = v
}

// GetRegion returns the Region field value.
func (o *CommitmentsAwsEC2RICommitment) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsEC2RICommitment) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *CommitmentsAwsEC2RICommitment) SetRegion(v string) {
	o.Region = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CommitmentsAwsEC2RICommitment) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsEC2RICommitment) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CommitmentsAwsEC2RICommitment) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CommitmentsAwsEC2RICommitment) SetStartDate(v string) {
	o.StartDate = &v
}

// GetTermLength returns the TermLength field value if set, zero value otherwise.
func (o *CommitmentsAwsEC2RICommitment) GetTermLength() float64 {
	if o == nil || o.TermLength == nil {
		var ret float64
		return ret
	}
	return *o.TermLength
}

// GetTermLengthOk returns a tuple with the TermLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsEC2RICommitment) GetTermLengthOk() (*float64, bool) {
	if o == nil || o.TermLength == nil {
		return nil, false
	}
	return o.TermLength, true
}

// HasTermLength returns a boolean if a field has been set.
func (o *CommitmentsAwsEC2RICommitment) HasTermLength() bool {
	return o != nil && o.TermLength != nil
}

// SetTermLength gets a reference to the given float64 and assigns it to the TermLength field.
func (o *CommitmentsAwsEC2RICommitment) SetTermLength(v float64) {
	o.TermLength = &v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *CommitmentsAwsEC2RICommitment) GetUtilization() float64 {
	if o == nil || o.Utilization == nil {
		var ret float64
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsEC2RICommitment) GetUtilizationOk() (*float64, bool) {
	if o == nil || o.Utilization == nil {
		return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *CommitmentsAwsEC2RICommitment) HasUtilization() bool {
	return o != nil && o.Utilization != nil
}

// SetUtilization gets a reference to the given float64 and assigns it to the Utilization field.
func (o *CommitmentsAwsEC2RICommitment) SetUtilization(v float64) {
	o.Utilization = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsAwsEC2RICommitment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AvailabilityZone != nil {
		toSerialize["availability_zone"] = o.AvailabilityZone
	}
	toSerialize["commitment_id"] = o.CommitmentId
	if o.ExpirationDate != nil {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	toSerialize["instance_type"] = o.InstanceType
	if o.NumberOfNfus != nil {
		toSerialize["number_of_nfus"] = o.NumberOfNfus
	}
	if o.NumberOfReservations != nil {
		toSerialize["number_of_reservations"] = o.NumberOfReservations
	}
	toSerialize["offering_class"] = o.OfferingClass
	toSerialize["operating_system"] = o.OperatingSystem
	toSerialize["purchase_option"] = o.PurchaseOption
	toSerialize["region"] = o.Region
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
func (o *CommitmentsAwsEC2RICommitment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvailabilityZone     *string  `json:"availability_zone,omitempty"`
		CommitmentId         *string  `json:"commitment_id"`
		ExpirationDate       *string  `json:"expiration_date,omitempty"`
		InstanceType         *string  `json:"instance_type"`
		NumberOfNfus         *float64 `json:"number_of_nfus,omitempty"`
		NumberOfReservations *float64 `json:"number_of_reservations,omitempty"`
		OfferingClass        *string  `json:"offering_class"`
		OperatingSystem      *string  `json:"operating_system"`
		PurchaseOption       *string  `json:"purchase_option"`
		Region               *string  `json:"region"`
		StartDate            *string  `json:"start_date,omitempty"`
		TermLength           *float64 `json:"term_length,omitempty"`
		Utilization          *float64 `json:"utilization,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CommitmentId == nil {
		return fmt.Errorf("required field commitment_id missing")
	}
	if all.InstanceType == nil {
		return fmt.Errorf("required field instance_type missing")
	}
	if all.OfferingClass == nil {
		return fmt.Errorf("required field offering_class missing")
	}
	if all.OperatingSystem == nil {
		return fmt.Errorf("required field operating_system missing")
	}
	if all.PurchaseOption == nil {
		return fmt.Errorf("required field purchase_option missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"availability_zone", "commitment_id", "expiration_date", "instance_type", "number_of_nfus", "number_of_reservations", "offering_class", "operating_system", "purchase_option", "region", "start_date", "term_length", "utilization"})
	} else {
		return err
	}
	o.AvailabilityZone = all.AvailabilityZone
	o.CommitmentId = *all.CommitmentId
	o.ExpirationDate = all.ExpirationDate
	o.InstanceType = *all.InstanceType
	o.NumberOfNfus = all.NumberOfNfus
	o.NumberOfReservations = all.NumberOfReservations
	o.OfferingClass = *all.OfferingClass
	o.OperatingSystem = *all.OperatingSystem
	o.PurchaseOption = *all.PurchaseOption
	o.Region = *all.Region
	o.StartDate = all.StartDate
	o.TermLength = all.TermLength
	o.Utilization = all.Utilization

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
