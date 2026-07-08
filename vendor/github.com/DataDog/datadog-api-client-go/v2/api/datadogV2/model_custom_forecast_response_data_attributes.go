// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomForecastResponseDataAttributes Attributes of a custom forecast.
type CustomForecastResponseDataAttributes struct {
	// The UUID of the budget that this custom forecast belongs to.
	BudgetUid string `json:"budget_uid"`
	// Timestamp the custom forecast was created, in Unix milliseconds.
	CreatedAt int64 `json:"created_at"`
	// The id of the user that created the custom forecast.
	CreatedBy string `json:"created_by"`
	// Monthly custom forecast entries.
	Entries []CustomForecastEntry `json:"entries"`
	// Timestamp the custom forecast was last updated, in Unix milliseconds.
	UpdatedAt int64 `json:"updated_at"`
	// The id of the user that last updated the custom forecast.
	UpdatedBy string `json:"updated_by"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomForecastResponseDataAttributes instantiates a new CustomForecastResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomForecastResponseDataAttributes(budgetUid string, createdAt int64, createdBy string, entries []CustomForecastEntry, updatedAt int64, updatedBy string) *CustomForecastResponseDataAttributes {
	this := CustomForecastResponseDataAttributes{}
	this.BudgetUid = budgetUid
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Entries = entries
	this.UpdatedAt = updatedAt
	this.UpdatedBy = updatedBy
	return &this
}

// NewCustomForecastResponseDataAttributesWithDefaults instantiates a new CustomForecastResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomForecastResponseDataAttributesWithDefaults() *CustomForecastResponseDataAttributes {
	this := CustomForecastResponseDataAttributes{}
	return &this
}

// GetBudgetUid returns the BudgetUid field value.
func (o *CustomForecastResponseDataAttributes) GetBudgetUid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BudgetUid
}

// GetBudgetUidOk returns a tuple with the BudgetUid field value
// and a boolean to check if the value has been set.
func (o *CustomForecastResponseDataAttributes) GetBudgetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetUid, true
}

// SetBudgetUid sets field value.
func (o *CustomForecastResponseDataAttributes) SetBudgetUid(v string) {
	o.BudgetUid = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *CustomForecastResponseDataAttributes) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomForecastResponseDataAttributes) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *CustomForecastResponseDataAttributes) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *CustomForecastResponseDataAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *CustomForecastResponseDataAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *CustomForecastResponseDataAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetEntries returns the Entries field value.
func (o *CustomForecastResponseDataAttributes) GetEntries() []CustomForecastEntry {
	if o == nil {
		var ret []CustomForecastEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value
// and a boolean to check if the value has been set.
func (o *CustomForecastResponseDataAttributes) GetEntriesOk() (*[]CustomForecastEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entries, true
}

// SetEntries sets field value.
func (o *CustomForecastResponseDataAttributes) SetEntries(v []CustomForecastEntry) {
	o.Entries = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *CustomForecastResponseDataAttributes) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomForecastResponseDataAttributes) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *CustomForecastResponseDataAttributes) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value.
func (o *CustomForecastResponseDataAttributes) GetUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *CustomForecastResponseDataAttributes) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value.
func (o *CustomForecastResponseDataAttributes) SetUpdatedBy(v string) {
	o.UpdatedBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomForecastResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["budget_uid"] = o.BudgetUid
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["entries"] = o.Entries
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["updated_by"] = o.UpdatedBy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomForecastResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BudgetUid *string                `json:"budget_uid"`
		CreatedAt *int64                 `json:"created_at"`
		CreatedBy *string                `json:"created_by"`
		Entries   *[]CustomForecastEntry `json:"entries"`
		UpdatedAt *int64                 `json:"updated_at"`
		UpdatedBy *string                `json:"updated_by"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BudgetUid == nil {
		return fmt.Errorf("required field budget_uid missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Entries == nil {
		return fmt.Errorf("required field entries missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.UpdatedBy == nil {
		return fmt.Errorf("required field updated_by missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"budget_uid", "created_at", "created_by", "entries", "updated_at", "updated_by"})
	} else {
		return err
	}
	o.BudgetUid = *all.BudgetUid
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Entries = *all.Entries
	o.UpdatedAt = *all.UpdatedAt
	o.UpdatedBy = *all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
