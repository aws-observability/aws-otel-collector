// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BudgetAttributes The attributes of a budget.
type BudgetAttributes struct {
	// The timestamp when the budget was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The id of the user that created the budget.
	CreatedBy *string `json:"created_by,omitempty"`
	// The month when the budget ends.
	EndMonth *int64 `json:"end_month,omitempty"`
	// The entries of the budget.
	Entries []BudgetEntry `json:"entries,omitempty"`
	// The cost query used to track against the budget.
	MetricsQuery *string `json:"metrics_query,omitempty"`
	// The name of the budget.
	Name *string `json:"name,omitempty"`
	// The id of the org the budget belongs to.
	OrgId *int64 `json:"org_id,omitempty"`
	// The month when the budget starts.
	StartMonth *int64 `json:"start_month,omitempty"`
	// The sum of all budget entries' amounts.
	TotalAmount *float64 `json:"total_amount,omitempty"`
	// The timestamp when the budget was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// The id of the user that created the budget.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBudgetAttributes instantiates a new BudgetAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBudgetAttributes() *BudgetAttributes {
	this := BudgetAttributes{}
	return &this
}

// NewBudgetAttributesWithDefaults instantiates a new BudgetAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBudgetAttributesWithDefaults() *BudgetAttributes {
	this := BudgetAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BudgetAttributes) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributes) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BudgetAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *BudgetAttributes) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BudgetAttributes) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BudgetAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *BudgetAttributes) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetEndMonth returns the EndMonth field value if set, zero value otherwise.
func (o *BudgetAttributes) GetEndMonth() int64 {
	if o == nil || o.EndMonth == nil {
		var ret int64
		return ret
	}
	return *o.EndMonth
}

// GetEndMonthOk returns a tuple with the EndMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributes) GetEndMonthOk() (*int64, bool) {
	if o == nil || o.EndMonth == nil {
		return nil, false
	}
	return o.EndMonth, true
}

// HasEndMonth returns a boolean if a field has been set.
func (o *BudgetAttributes) HasEndMonth() bool {
	return o != nil && o.EndMonth != nil
}

// SetEndMonth gets a reference to the given int64 and assigns it to the EndMonth field.
func (o *BudgetAttributes) SetEndMonth(v int64) {
	o.EndMonth = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *BudgetAttributes) GetEntries() []BudgetEntry {
	if o == nil || o.Entries == nil {
		var ret []BudgetEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributes) GetEntriesOk() (*[]BudgetEntry, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return &o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *BudgetAttributes) HasEntries() bool {
	return o != nil && o.Entries != nil
}

// SetEntries gets a reference to the given []BudgetEntry and assigns it to the Entries field.
func (o *BudgetAttributes) SetEntries(v []BudgetEntry) {
	o.Entries = v
}

// GetMetricsQuery returns the MetricsQuery field value if set, zero value otherwise.
func (o *BudgetAttributes) GetMetricsQuery() string {
	if o == nil || o.MetricsQuery == nil {
		var ret string
		return ret
	}
	return *o.MetricsQuery
}

// GetMetricsQueryOk returns a tuple with the MetricsQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributes) GetMetricsQueryOk() (*string, bool) {
	if o == nil || o.MetricsQuery == nil {
		return nil, false
	}
	return o.MetricsQuery, true
}

// HasMetricsQuery returns a boolean if a field has been set.
func (o *BudgetAttributes) HasMetricsQuery() bool {
	return o != nil && o.MetricsQuery != nil
}

// SetMetricsQuery gets a reference to the given string and assigns it to the MetricsQuery field.
func (o *BudgetAttributes) SetMetricsQuery(v string) {
	o.MetricsQuery = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BudgetAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BudgetAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BudgetAttributes) SetName(v string) {
	o.Name = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *BudgetAttributes) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *BudgetAttributes) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *BudgetAttributes) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetStartMonth returns the StartMonth field value if set, zero value otherwise.
func (o *BudgetAttributes) GetStartMonth() int64 {
	if o == nil || o.StartMonth == nil {
		var ret int64
		return ret
	}
	return *o.StartMonth
}

// GetStartMonthOk returns a tuple with the StartMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributes) GetStartMonthOk() (*int64, bool) {
	if o == nil || o.StartMonth == nil {
		return nil, false
	}
	return o.StartMonth, true
}

// HasStartMonth returns a boolean if a field has been set.
func (o *BudgetAttributes) HasStartMonth() bool {
	return o != nil && o.StartMonth != nil
}

// SetStartMonth gets a reference to the given int64 and assigns it to the StartMonth field.
func (o *BudgetAttributes) SetStartMonth(v int64) {
	o.StartMonth = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *BudgetAttributes) GetTotalAmount() float64 {
	if o == nil || o.TotalAmount == nil {
		var ret float64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributes) GetTotalAmountOk() (*float64, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *BudgetAttributes) HasTotalAmount() bool {
	return o != nil && o.TotalAmount != nil
}

// SetTotalAmount gets a reference to the given float64 and assigns it to the TotalAmount field.
func (o *BudgetAttributes) SetTotalAmount(v float64) {
	o.TotalAmount = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *BudgetAttributes) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributes) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *BudgetAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *BudgetAttributes) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *BudgetAttributes) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributes) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *BudgetAttributes) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *BudgetAttributes) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BudgetAttributes) MarshalJSON() ([]byte, error) {
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
func (o *BudgetAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt    *int64        `json:"created_at,omitempty"`
		CreatedBy    *string       `json:"created_by,omitempty"`
		EndMonth     *int64        `json:"end_month,omitempty"`
		Entries      []BudgetEntry `json:"entries,omitempty"`
		MetricsQuery *string       `json:"metrics_query,omitempty"`
		Name         *string       `json:"name,omitempty"`
		OrgId        *int64        `json:"org_id,omitempty"`
		StartMonth   *int64        `json:"start_month,omitempty"`
		TotalAmount  *float64      `json:"total_amount,omitempty"`
		UpdatedAt    *int64        `json:"updated_at,omitempty"`
		UpdatedBy    *string       `json:"updated_by,omitempty"`
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
