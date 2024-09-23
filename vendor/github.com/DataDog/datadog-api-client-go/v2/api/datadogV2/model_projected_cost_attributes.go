// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProjectedCostAttributes Projected Cost attributes data.
type ProjectedCostAttributes struct {
	// The account name.
	AccountName *string `json:"account_name,omitempty"`
	// The account public ID.
	AccountPublicId *string `json:"account_public_id,omitempty"`
	// List of charges data reported for the requested month.
	Charges []ChargebackBreakdown `json:"charges,omitempty"`
	// The month requested.
	Date *time.Time `json:"date,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The total projected cost of products for the month.
	ProjectedTotalCost *float64 `json:"projected_total_cost,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// The region of the Datadog instance that the organization belongs to.
	Region *string `json:"region,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProjectedCostAttributes instantiates a new ProjectedCostAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProjectedCostAttributes() *ProjectedCostAttributes {
	this := ProjectedCostAttributes{}
	return &this
}

// NewProjectedCostAttributesWithDefaults instantiates a new ProjectedCostAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProjectedCostAttributesWithDefaults() *ProjectedCostAttributes {
	this := ProjectedCostAttributes{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *ProjectedCostAttributes) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectedCostAttributes) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *ProjectedCostAttributes) HasAccountName() bool {
	return o != nil && o.AccountName != nil
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *ProjectedCostAttributes) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAccountPublicId returns the AccountPublicId field value if set, zero value otherwise.
func (o *ProjectedCostAttributes) GetAccountPublicId() string {
	if o == nil || o.AccountPublicId == nil {
		var ret string
		return ret
	}
	return *o.AccountPublicId
}

// GetAccountPublicIdOk returns a tuple with the AccountPublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectedCostAttributes) GetAccountPublicIdOk() (*string, bool) {
	if o == nil || o.AccountPublicId == nil {
		return nil, false
	}
	return o.AccountPublicId, true
}

// HasAccountPublicId returns a boolean if a field has been set.
func (o *ProjectedCostAttributes) HasAccountPublicId() bool {
	return o != nil && o.AccountPublicId != nil
}

// SetAccountPublicId gets a reference to the given string and assigns it to the AccountPublicId field.
func (o *ProjectedCostAttributes) SetAccountPublicId(v string) {
	o.AccountPublicId = &v
}

// GetCharges returns the Charges field value if set, zero value otherwise.
func (o *ProjectedCostAttributes) GetCharges() []ChargebackBreakdown {
	if o == nil || o.Charges == nil {
		var ret []ChargebackBreakdown
		return ret
	}
	return o.Charges
}

// GetChargesOk returns a tuple with the Charges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectedCostAttributes) GetChargesOk() (*[]ChargebackBreakdown, bool) {
	if o == nil || o.Charges == nil {
		return nil, false
	}
	return &o.Charges, true
}

// HasCharges returns a boolean if a field has been set.
func (o *ProjectedCostAttributes) HasCharges() bool {
	return o != nil && o.Charges != nil
}

// SetCharges gets a reference to the given []ChargebackBreakdown and assigns it to the Charges field.
func (o *ProjectedCostAttributes) SetCharges(v []ChargebackBreakdown) {
	o.Charges = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *ProjectedCostAttributes) GetDate() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectedCostAttributes) GetDateOk() (*time.Time, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *ProjectedCostAttributes) HasDate() bool {
	return o != nil && o.Date != nil
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *ProjectedCostAttributes) SetDate(v time.Time) {
	o.Date = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *ProjectedCostAttributes) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectedCostAttributes) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *ProjectedCostAttributes) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *ProjectedCostAttributes) SetOrgName(v string) {
	o.OrgName = &v
}

// GetProjectedTotalCost returns the ProjectedTotalCost field value if set, zero value otherwise.
func (o *ProjectedCostAttributes) GetProjectedTotalCost() float64 {
	if o == nil || o.ProjectedTotalCost == nil {
		var ret float64
		return ret
	}
	return *o.ProjectedTotalCost
}

// GetProjectedTotalCostOk returns a tuple with the ProjectedTotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectedCostAttributes) GetProjectedTotalCostOk() (*float64, bool) {
	if o == nil || o.ProjectedTotalCost == nil {
		return nil, false
	}
	return o.ProjectedTotalCost, true
}

// HasProjectedTotalCost returns a boolean if a field has been set.
func (o *ProjectedCostAttributes) HasProjectedTotalCost() bool {
	return o != nil && o.ProjectedTotalCost != nil
}

// SetProjectedTotalCost gets a reference to the given float64 and assigns it to the ProjectedTotalCost field.
func (o *ProjectedCostAttributes) SetProjectedTotalCost(v float64) {
	o.ProjectedTotalCost = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *ProjectedCostAttributes) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectedCostAttributes) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *ProjectedCostAttributes) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *ProjectedCostAttributes) SetPublicId(v string) {
	o.PublicId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ProjectedCostAttributes) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectedCostAttributes) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ProjectedCostAttributes) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ProjectedCostAttributes) SetRegion(v string) {
	o.Region = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProjectedCostAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	if o.AccountPublicId != nil {
		toSerialize["account_public_id"] = o.AccountPublicId
	}
	if o.Charges != nil {
		toSerialize["charges"] = o.Charges
	}
	if o.Date != nil {
		if o.Date.Nanosecond() == 0 {
			toSerialize["date"] = o.Date.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["date"] = o.Date.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.ProjectedTotalCost != nil {
		toSerialize["projected_total_cost"] = o.ProjectedTotalCost
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProjectedCostAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountName        *string               `json:"account_name,omitempty"`
		AccountPublicId    *string               `json:"account_public_id,omitempty"`
		Charges            []ChargebackBreakdown `json:"charges,omitempty"`
		Date               *time.Time            `json:"date,omitempty"`
		OrgName            *string               `json:"org_name,omitempty"`
		ProjectedTotalCost *float64              `json:"projected_total_cost,omitempty"`
		PublicId           *string               `json:"public_id,omitempty"`
		Region             *string               `json:"region,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_name", "account_public_id", "charges", "date", "org_name", "projected_total_cost", "public_id", "region"})
	} else {
		return err
	}
	o.AccountName = all.AccountName
	o.AccountPublicId = all.AccountPublicId
	o.Charges = all.Charges
	o.Date = all.Date
	o.OrgName = all.OrgName
	o.ProjectedTotalCost = all.ProjectedTotalCost
	o.PublicId = all.PublicId
	o.Region = all.Region

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
