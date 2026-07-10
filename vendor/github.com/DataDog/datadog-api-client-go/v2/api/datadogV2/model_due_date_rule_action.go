// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DueDateRuleAction The action to take when the due date rule matches a finding.
type DueDateRuleAction struct {
	// A list of severity-to-due-date mappings. Each severity may appear at most once.
	DueDaysPerSeverity []DueDatePerSeverityItem `json:"due_days_per_severity"`
	// The reference point from which the due date is calculated. When `fix_available` is selected but not applicable to the finding type, `first_seen` is used instead.
	DueFrom DueDateFrom `json:"due_from"`
	// An optional description providing more context for the due date assignment.
	ReasonDescription *string `json:"reason_description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDueDateRuleAction instantiates a new DueDateRuleAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDueDateRuleAction(dueDaysPerSeverity []DueDatePerSeverityItem, dueFrom DueDateFrom) *DueDateRuleAction {
	this := DueDateRuleAction{}
	this.DueDaysPerSeverity = dueDaysPerSeverity
	this.DueFrom = dueFrom
	return &this
}

// NewDueDateRuleActionWithDefaults instantiates a new DueDateRuleAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDueDateRuleActionWithDefaults() *DueDateRuleAction {
	this := DueDateRuleAction{}
	return &this
}

// GetDueDaysPerSeverity returns the DueDaysPerSeverity field value.
func (o *DueDateRuleAction) GetDueDaysPerSeverity() []DueDatePerSeverityItem {
	if o == nil {
		var ret []DueDatePerSeverityItem
		return ret
	}
	return o.DueDaysPerSeverity
}

// GetDueDaysPerSeverityOk returns a tuple with the DueDaysPerSeverity field value
// and a boolean to check if the value has been set.
func (o *DueDateRuleAction) GetDueDaysPerSeverityOk() (*[]DueDatePerSeverityItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DueDaysPerSeverity, true
}

// SetDueDaysPerSeverity sets field value.
func (o *DueDateRuleAction) SetDueDaysPerSeverity(v []DueDatePerSeverityItem) {
	o.DueDaysPerSeverity = v
}

// GetDueFrom returns the DueFrom field value.
func (o *DueDateRuleAction) GetDueFrom() DueDateFrom {
	if o == nil {
		var ret DueDateFrom
		return ret
	}
	return o.DueFrom
}

// GetDueFromOk returns a tuple with the DueFrom field value
// and a boolean to check if the value has been set.
func (o *DueDateRuleAction) GetDueFromOk() (*DueDateFrom, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DueFrom, true
}

// SetDueFrom sets field value.
func (o *DueDateRuleAction) SetDueFrom(v DueDateFrom) {
	o.DueFrom = v
}

// GetReasonDescription returns the ReasonDescription field value if set, zero value otherwise.
func (o *DueDateRuleAction) GetReasonDescription() string {
	if o == nil || o.ReasonDescription == nil {
		var ret string
		return ret
	}
	return *o.ReasonDescription
}

// GetReasonDescriptionOk returns a tuple with the ReasonDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DueDateRuleAction) GetReasonDescriptionOk() (*string, bool) {
	if o == nil || o.ReasonDescription == nil {
		return nil, false
	}
	return o.ReasonDescription, true
}

// HasReasonDescription returns a boolean if a field has been set.
func (o *DueDateRuleAction) HasReasonDescription() bool {
	return o != nil && o.ReasonDescription != nil
}

// SetReasonDescription gets a reference to the given string and assigns it to the ReasonDescription field.
func (o *DueDateRuleAction) SetReasonDescription(v string) {
	o.ReasonDescription = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DueDateRuleAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["due_days_per_severity"] = o.DueDaysPerSeverity
	toSerialize["due_from"] = o.DueFrom
	if o.ReasonDescription != nil {
		toSerialize["reason_description"] = o.ReasonDescription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DueDateRuleAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DueDaysPerSeverity *[]DueDatePerSeverityItem `json:"due_days_per_severity"`
		DueFrom            *DueDateFrom              `json:"due_from"`
		ReasonDescription  *string                   `json:"reason_description,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DueDaysPerSeverity == nil {
		return fmt.Errorf("required field due_days_per_severity missing")
	}
	if all.DueFrom == nil {
		return fmt.Errorf("required field due_from missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"due_days_per_severity", "due_from", "reason_description"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DueDaysPerSeverity = *all.DueDaysPerSeverity
	if !all.DueFrom.IsValid() {
		hasInvalidField = true
	} else {
		o.DueFrom = *all.DueFrom
	}
	o.ReasonDescription = all.ReasonDescription

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
