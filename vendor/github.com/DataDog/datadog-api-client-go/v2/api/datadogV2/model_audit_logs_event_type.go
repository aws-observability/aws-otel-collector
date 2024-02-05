// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuditLogsEventType Type of the event.
type AuditLogsEventType string

// List of AuditLogsEventType.
const (
	AUDITLOGSEVENTTYPE_Audit AuditLogsEventType = "audit"
)

var allowedAuditLogsEventTypeEnumValues = []AuditLogsEventType{
	AUDITLOGSEVENTTYPE_Audit,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AuditLogsEventType) GetAllowedValues() []AuditLogsEventType {
	return allowedAuditLogsEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AuditLogsEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AuditLogsEventType(value)
	return nil
}

// NewAuditLogsEventTypeFromValue returns a pointer to a valid AuditLogsEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAuditLogsEventTypeFromValue(v string) (*AuditLogsEventType, error) {
	ev := AuditLogsEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AuditLogsEventType: valid values are %v", v, allowedAuditLogsEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AuditLogsEventType) IsValid() bool {
	for _, existing := range allowedAuditLogsEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuditLogsEventType value.
func (v AuditLogsEventType) Ptr() *AuditLogsEventType {
	return &v
}
