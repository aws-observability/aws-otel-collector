// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionFilterEventType The type of RUM events to filter on.
type RumRetentionFilterEventType string

// List of RumRetentionFilterEventType.
const (
	RUMRETENTIONFILTEREVENTTYPE_SESSION   RumRetentionFilterEventType = "session"
	RUMRETENTIONFILTEREVENTTYPE_VIEW      RumRetentionFilterEventType = "view"
	RUMRETENTIONFILTEREVENTTYPE_ACTION    RumRetentionFilterEventType = "action"
	RUMRETENTIONFILTEREVENTTYPE_ERROR     RumRetentionFilterEventType = "error"
	RUMRETENTIONFILTEREVENTTYPE_RESOURCE  RumRetentionFilterEventType = "resource"
	RUMRETENTIONFILTEREVENTTYPE_LONG_TASK RumRetentionFilterEventType = "long_task"
	RUMRETENTIONFILTEREVENTTYPE_VITAL     RumRetentionFilterEventType = "vital"
)

var allowedRumRetentionFilterEventTypeEnumValues = []RumRetentionFilterEventType{
	RUMRETENTIONFILTEREVENTTYPE_SESSION,
	RUMRETENTIONFILTEREVENTTYPE_VIEW,
	RUMRETENTIONFILTEREVENTTYPE_ACTION,
	RUMRETENTIONFILTEREVENTTYPE_ERROR,
	RUMRETENTIONFILTEREVENTTYPE_RESOURCE,
	RUMRETENTIONFILTEREVENTTYPE_LONG_TASK,
	RUMRETENTIONFILTEREVENTTYPE_VITAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRetentionFilterEventType) GetAllowedValues() []RumRetentionFilterEventType {
	return allowedRumRetentionFilterEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRetentionFilterEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRetentionFilterEventType(value)
	return nil
}

// NewRumRetentionFilterEventTypeFromValue returns a pointer to a valid RumRetentionFilterEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRetentionFilterEventTypeFromValue(v string) (*RumRetentionFilterEventType, error) {
	ev := RumRetentionFilterEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRetentionFilterEventType: valid values are %v", v, allowedRumRetentionFilterEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRetentionFilterEventType) IsValid() bool {
	for _, existing := range allowedRumRetentionFilterEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRetentionFilterEventType value.
func (v RumRetentionFilterEventType) Ptr() *RumRetentionFilterEventType {
	return &v
}
