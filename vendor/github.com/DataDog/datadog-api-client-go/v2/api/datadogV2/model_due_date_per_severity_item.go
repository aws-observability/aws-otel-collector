// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DueDatePerSeverityItem A mapping of a severity level to the number of days until a finding is due.
type DueDatePerSeverityItem struct {
	// The number of days from the reference point until the finding is due.
	DueInDays int64 `json:"due_in_days"`
	// A severity level used to configure due date thresholds.
	Severity DueDateSeverity `json:"severity"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDueDatePerSeverityItem instantiates a new DueDatePerSeverityItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDueDatePerSeverityItem(dueInDays int64, severity DueDateSeverity) *DueDatePerSeverityItem {
	this := DueDatePerSeverityItem{}
	this.DueInDays = dueInDays
	this.Severity = severity
	return &this
}

// NewDueDatePerSeverityItemWithDefaults instantiates a new DueDatePerSeverityItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDueDatePerSeverityItemWithDefaults() *DueDatePerSeverityItem {
	this := DueDatePerSeverityItem{}
	return &this
}

// GetDueInDays returns the DueInDays field value.
func (o *DueDatePerSeverityItem) GetDueInDays() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DueInDays
}

// GetDueInDaysOk returns a tuple with the DueInDays field value
// and a boolean to check if the value has been set.
func (o *DueDatePerSeverityItem) GetDueInDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DueInDays, true
}

// SetDueInDays sets field value.
func (o *DueDatePerSeverityItem) SetDueInDays(v int64) {
	o.DueInDays = v
}

// GetSeverity returns the Severity field value.
func (o *DueDatePerSeverityItem) GetSeverity() DueDateSeverity {
	if o == nil {
		var ret DueDateSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *DueDatePerSeverityItem) GetSeverityOk() (*DueDateSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *DueDatePerSeverityItem) SetSeverity(v DueDateSeverity) {
	o.Severity = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DueDatePerSeverityItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["due_in_days"] = o.DueInDays
	toSerialize["severity"] = o.Severity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DueDatePerSeverityItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DueInDays *int64           `json:"due_in_days"`
		Severity  *DueDateSeverity `json:"severity"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DueInDays == nil {
		return fmt.Errorf("required field due_in_days missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"due_in_days", "severity"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DueInDays = *all.DueInDays
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
