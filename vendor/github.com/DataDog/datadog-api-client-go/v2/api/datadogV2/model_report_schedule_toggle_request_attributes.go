// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleToggleRequestAttributes The status to set on the report schedule.
type ReportScheduleToggleRequestAttributes struct {
	// Whether the schedule is currently delivering reports (`active`) or paused (`inactive`).
	Status ReportScheduleStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportScheduleToggleRequestAttributes instantiates a new ReportScheduleToggleRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportScheduleToggleRequestAttributes(status ReportScheduleStatus) *ReportScheduleToggleRequestAttributes {
	this := ReportScheduleToggleRequestAttributes{}
	this.Status = status
	return &this
}

// NewReportScheduleToggleRequestAttributesWithDefaults instantiates a new ReportScheduleToggleRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportScheduleToggleRequestAttributesWithDefaults() *ReportScheduleToggleRequestAttributes {
	this := ReportScheduleToggleRequestAttributes{}
	return &this
}

// GetStatus returns the Status field value.
func (o *ReportScheduleToggleRequestAttributes) GetStatus() ReportScheduleStatus {
	if o == nil {
		var ret ReportScheduleStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleToggleRequestAttributes) GetStatusOk() (*ReportScheduleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ReportScheduleToggleRequestAttributes) SetStatus(v ReportScheduleStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportScheduleToggleRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReportScheduleToggleRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status *ReportScheduleStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"status"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
