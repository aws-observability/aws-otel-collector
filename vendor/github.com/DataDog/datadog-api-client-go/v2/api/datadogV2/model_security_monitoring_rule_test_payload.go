// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleTestPayload - Test a rule.
type SecurityMonitoringRuleTestPayload struct {
	SecurityMonitoringStandardRuleTestPayload *SecurityMonitoringStandardRuleTestPayload

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SecurityMonitoringStandardRuleTestPayloadAsSecurityMonitoringRuleTestPayload is a convenience function that returns SecurityMonitoringStandardRuleTestPayload wrapped in SecurityMonitoringRuleTestPayload.
func SecurityMonitoringStandardRuleTestPayloadAsSecurityMonitoringRuleTestPayload(v *SecurityMonitoringStandardRuleTestPayload) SecurityMonitoringRuleTestPayload {
	return SecurityMonitoringRuleTestPayload{SecurityMonitoringStandardRuleTestPayload: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SecurityMonitoringRuleTestPayload) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SecurityMonitoringStandardRuleTestPayload
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringStandardRuleTestPayload)
	if err == nil {
		if obj.SecurityMonitoringStandardRuleTestPayload != nil && obj.SecurityMonitoringStandardRuleTestPayload.UnparsedObject == nil {
			jsonSecurityMonitoringStandardRuleTestPayload, _ := datadog.Marshal(obj.SecurityMonitoringStandardRuleTestPayload)
			if string(jsonSecurityMonitoringStandardRuleTestPayload) == "{}" { // empty struct
				obj.SecurityMonitoringStandardRuleTestPayload = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringStandardRuleTestPayload = nil
		}
	} else {
		obj.SecurityMonitoringStandardRuleTestPayload = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SecurityMonitoringStandardRuleTestPayload = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SecurityMonitoringRuleTestPayload) MarshalJSON() ([]byte, error) {
	if obj.SecurityMonitoringStandardRuleTestPayload != nil {
		return datadog.Marshal(&obj.SecurityMonitoringStandardRuleTestPayload)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SecurityMonitoringRuleTestPayload) GetActualInstance() interface{} {
	if obj.SecurityMonitoringStandardRuleTestPayload != nil {
		return obj.SecurityMonitoringStandardRuleTestPayload
	}

	// all schemas are nil
	return nil
}
