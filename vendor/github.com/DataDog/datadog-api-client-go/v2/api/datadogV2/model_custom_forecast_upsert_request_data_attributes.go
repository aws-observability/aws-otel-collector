// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomForecastUpsertRequestDataAttributes Attributes of a custom forecast upsert request.
type CustomForecastUpsertRequestDataAttributes struct {
	// The UUID of the budget that this custom forecast belongs to.
	BudgetUid string `json:"budget_uid"`
	// Monthly custom forecast entries. An empty list deletes any existing
	// custom forecast for the budget.
	Entries []CustomForecastEntry `json:"entries"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomForecastUpsertRequestDataAttributes instantiates a new CustomForecastUpsertRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomForecastUpsertRequestDataAttributes(budgetUid string, entries []CustomForecastEntry) *CustomForecastUpsertRequestDataAttributes {
	this := CustomForecastUpsertRequestDataAttributes{}
	this.BudgetUid = budgetUid
	this.Entries = entries
	return &this
}

// NewCustomForecastUpsertRequestDataAttributesWithDefaults instantiates a new CustomForecastUpsertRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomForecastUpsertRequestDataAttributesWithDefaults() *CustomForecastUpsertRequestDataAttributes {
	this := CustomForecastUpsertRequestDataAttributes{}
	return &this
}

// GetBudgetUid returns the BudgetUid field value.
func (o *CustomForecastUpsertRequestDataAttributes) GetBudgetUid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BudgetUid
}

// GetBudgetUidOk returns a tuple with the BudgetUid field value
// and a boolean to check if the value has been set.
func (o *CustomForecastUpsertRequestDataAttributes) GetBudgetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetUid, true
}

// SetBudgetUid sets field value.
func (o *CustomForecastUpsertRequestDataAttributes) SetBudgetUid(v string) {
	o.BudgetUid = v
}

// GetEntries returns the Entries field value.
func (o *CustomForecastUpsertRequestDataAttributes) GetEntries() []CustomForecastEntry {
	if o == nil {
		var ret []CustomForecastEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value
// and a boolean to check if the value has been set.
func (o *CustomForecastUpsertRequestDataAttributes) GetEntriesOk() (*[]CustomForecastEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entries, true
}

// SetEntries sets field value.
func (o *CustomForecastUpsertRequestDataAttributes) SetEntries(v []CustomForecastEntry) {
	o.Entries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomForecastUpsertRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["budget_uid"] = o.BudgetUid
	toSerialize["entries"] = o.Entries

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomForecastUpsertRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BudgetUid *string                `json:"budget_uid"`
		Entries   *[]CustomForecastEntry `json:"entries"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BudgetUid == nil {
		return fmt.Errorf("required field budget_uid missing")
	}
	if all.Entries == nil {
		return fmt.Errorf("required field entries missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"budget_uid", "entries"})
	} else {
		return err
	}
	o.BudgetUid = *all.BudgetUid
	o.Entries = *all.Entries

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
