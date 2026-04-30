// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BudgetWithEntriesDataAttributes The attributes of a budget including all its monthly entries.
type BudgetWithEntriesDataAttributes struct {
	// The timestamp when the budget was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The ID of the user that created the budget.
	CreatedBy *string `json:"created_by,omitempty"`
	// The month when the budget ends, in YYYYMM format.
	EndMonth *int64 `json:"end_month,omitempty"`
	// The list of monthly budget entries.
	Entries []BudgetWithEntriesDataAttributesEntriesItems `json:"entries,omitempty"`
	// The cost query used to track spending against the budget.
	MetricsQuery *string `json:"metrics_query,omitempty"`
	// The name of the budget.
	Name *string `json:"name,omitempty"`
	// The ID of the organization the budget belongs to.
	OrgId *int64 `json:"org_id,omitempty"`
	// The month when the budget starts, in YYYYMM format.
	StartMonth *int64 `json:"start_month,omitempty"`
	// The total budget amount across all entries.
	TotalAmount *float64 `json:"total_amount,omitempty"`
	// The timestamp when the budget was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// The ID of the user that last updated the budget.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBudgetWithEntriesDataAttributes instantiates a new BudgetWithEntriesDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBudgetWithEntriesDataAttributes() *BudgetWithEntriesDataAttributes {
	this := BudgetWithEntriesDataAttributes{}
	return &this
}

// NewBudgetWithEntriesDataAttributesWithDefaults instantiates a new BudgetWithEntriesDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBudgetWithEntriesDataAttributesWithDefaults() *BudgetWithEntriesDataAttributes {
	this := BudgetWithEntriesDataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BudgetWithEntriesDataAttributes) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesDataAttributes) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BudgetWithEntriesDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *BudgetWithEntriesDataAttributes) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BudgetWithEntriesDataAttributes) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesDataAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BudgetWithEntriesDataAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *BudgetWithEntriesDataAttributes) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetEndMonth returns the EndMonth field value if set, zero value otherwise.
func (o *BudgetWithEntriesDataAttributes) GetEndMonth() int64 {
	if o == nil || o.EndMonth == nil {
		var ret int64
		return ret
	}
	return *o.EndMonth
}

// GetEndMonthOk returns a tuple with the EndMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesDataAttributes) GetEndMonthOk() (*int64, bool) {
	if o == nil || o.EndMonth == nil {
		return nil, false
	}
	return o.EndMonth, true
}

// HasEndMonth returns a boolean if a field has been set.
func (o *BudgetWithEntriesDataAttributes) HasEndMonth() bool {
	return o != nil && o.EndMonth != nil
}

// SetEndMonth gets a reference to the given int64 and assigns it to the EndMonth field.
func (o *BudgetWithEntriesDataAttributes) SetEndMonth(v int64) {
	o.EndMonth = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *BudgetWithEntriesDataAttributes) GetEntries() []BudgetWithEntriesDataAttributesEntriesItems {
	if o == nil || o.Entries == nil {
		var ret []BudgetWithEntriesDataAttributesEntriesItems
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesDataAttributes) GetEntriesOk() (*[]BudgetWithEntriesDataAttributesEntriesItems, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return &o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *BudgetWithEntriesDataAttributes) HasEntries() bool {
	return o != nil && o.Entries != nil
}

// SetEntries gets a reference to the given []BudgetWithEntriesDataAttributesEntriesItems and assigns it to the Entries field.
func (o *BudgetWithEntriesDataAttributes) SetEntries(v []BudgetWithEntriesDataAttributesEntriesItems) {
	o.Entries = v
}

// GetMetricsQuery returns the MetricsQuery field value if set, zero value otherwise.
func (o *BudgetWithEntriesDataAttributes) GetMetricsQuery() string {
	if o == nil || o.MetricsQuery == nil {
		var ret string
		return ret
	}
	return *o.MetricsQuery
}

// GetMetricsQueryOk returns a tuple with the MetricsQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesDataAttributes) GetMetricsQueryOk() (*string, bool) {
	if o == nil || o.MetricsQuery == nil {
		return nil, false
	}
	return o.MetricsQuery, true
}

// HasMetricsQuery returns a boolean if a field has been set.
func (o *BudgetWithEntriesDataAttributes) HasMetricsQuery() bool {
	return o != nil && o.MetricsQuery != nil
}

// SetMetricsQuery gets a reference to the given string and assigns it to the MetricsQuery field.
func (o *BudgetWithEntriesDataAttributes) SetMetricsQuery(v string) {
	o.MetricsQuery = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BudgetWithEntriesDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BudgetWithEntriesDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BudgetWithEntriesDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *BudgetWithEntriesDataAttributes) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesDataAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *BudgetWithEntriesDataAttributes) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *BudgetWithEntriesDataAttributes) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetStartMonth returns the StartMonth field value if set, zero value otherwise.
func (o *BudgetWithEntriesDataAttributes) GetStartMonth() int64 {
	if o == nil || o.StartMonth == nil {
		var ret int64
		return ret
	}
	return *o.StartMonth
}

// GetStartMonthOk returns a tuple with the StartMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesDataAttributes) GetStartMonthOk() (*int64, bool) {
	if o == nil || o.StartMonth == nil {
		return nil, false
	}
	return o.StartMonth, true
}

// HasStartMonth returns a boolean if a field has been set.
func (o *BudgetWithEntriesDataAttributes) HasStartMonth() bool {
	return o != nil && o.StartMonth != nil
}

// SetStartMonth gets a reference to the given int64 and assigns it to the StartMonth field.
func (o *BudgetWithEntriesDataAttributes) SetStartMonth(v int64) {
	o.StartMonth = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *BudgetWithEntriesDataAttributes) GetTotalAmount() float64 {
	if o == nil || o.TotalAmount == nil {
		var ret float64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesDataAttributes) GetTotalAmountOk() (*float64, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *BudgetWithEntriesDataAttributes) HasTotalAmount() bool {
	return o != nil && o.TotalAmount != nil
}

// SetTotalAmount gets a reference to the given float64 and assigns it to the TotalAmount field.
func (o *BudgetWithEntriesDataAttributes) SetTotalAmount(v float64) {
	o.TotalAmount = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *BudgetWithEntriesDataAttributes) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesDataAttributes) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *BudgetWithEntriesDataAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *BudgetWithEntriesDataAttributes) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *BudgetWithEntriesDataAttributes) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetWithEntriesDataAttributes) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *BudgetWithEntriesDataAttributes) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *BudgetWithEntriesDataAttributes) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BudgetWithEntriesDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.EndMonth != nil {
		toSerialize["end_month"] = o.EndMonth
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	if o.MetricsQuery != nil {
		toSerialize["metrics_query"] = o.MetricsQuery
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.StartMonth != nil {
		toSerialize["start_month"] = o.StartMonth
	}
	if o.TotalAmount != nil {
		toSerialize["total_amount"] = o.TotalAmount
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BudgetWithEntriesDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt    *int64                                        `json:"created_at,omitempty"`
		CreatedBy    *string                                       `json:"created_by,omitempty"`
		EndMonth     *int64                                        `json:"end_month,omitempty"`
		Entries      []BudgetWithEntriesDataAttributesEntriesItems `json:"entries,omitempty"`
		MetricsQuery *string                                       `json:"metrics_query,omitempty"`
		Name         *string                                       `json:"name,omitempty"`
		OrgId        *int64                                        `json:"org_id,omitempty"`
		StartMonth   *int64                                        `json:"start_month,omitempty"`
		TotalAmount  *float64                                      `json:"total_amount,omitempty"`
		UpdatedAt    *int64                                        `json:"updated_at,omitempty"`
		UpdatedBy    *string                                       `json:"updated_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "end_month", "entries", "metrics_query", "name", "org_id", "start_month", "total_amount", "updated_at", "updated_by"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.EndMonth = all.EndMonth
	o.Entries = all.Entries
	o.MetricsQuery = all.MetricsQuery
	o.Name = all.Name
	o.OrgId = all.OrgId
	o.StartMonth = all.StartMonth
	o.TotalAmount = all.TotalAmount
	o.UpdatedAt = all.UpdatedAt
	o.UpdatedBy = all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
