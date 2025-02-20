// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleConvertPayload - Convert a rule from JSON to Terraform.
type SecurityMonitoringRuleConvertPayload struct {
	SecurityMonitoringStandardRulePayload *SecurityMonitoringStandardRulePayload
	SecurityMonitoringSignalRulePayload   *SecurityMonitoringSignalRulePayload

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SecurityMonitoringStandardRulePayloadAsSecurityMonitoringRuleConvertPayload is a convenience function that returns SecurityMonitoringStandardRulePayload wrapped in SecurityMonitoringRuleConvertPayload.
func SecurityMonitoringStandardRulePayloadAsSecurityMonitoringRuleConvertPayload(v *SecurityMonitoringStandardRulePayload) SecurityMonitoringRuleConvertPayload {
	return SecurityMonitoringRuleConvertPayload{SecurityMonitoringStandardRulePayload: v}
}

// SecurityMonitoringSignalRulePayloadAsSecurityMonitoringRuleConvertPayload is a convenience function that returns SecurityMonitoringSignalRulePayload wrapped in SecurityMonitoringRuleConvertPayload.
func SecurityMonitoringSignalRulePayloadAsSecurityMonitoringRuleConvertPayload(v *SecurityMonitoringSignalRulePayload) SecurityMonitoringRuleConvertPayload {
	return SecurityMonitoringRuleConvertPayload{SecurityMonitoringSignalRulePayload: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SecurityMonitoringRuleConvertPayload) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SecurityMonitoringStandardRulePayload
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringStandardRulePayload)
	if err == nil {
		if obj.SecurityMonitoringStandardRulePayload != nil && obj.SecurityMonitoringStandardRulePayload.UnparsedObject == nil {
			jsonSecurityMonitoringStandardRulePayload, _ := datadog.Marshal(obj.SecurityMonitoringStandardRulePayload)
			if string(jsonSecurityMonitoringStandardRulePayload) == "{}" { // empty struct
				obj.SecurityMonitoringStandardRulePayload = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringStandardRulePayload = nil
		}
	} else {
		obj.SecurityMonitoringStandardRulePayload = nil
	}

	// try to unmarshal data into SecurityMonitoringSignalRulePayload
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringSignalRulePayload)
	if err == nil {
		if obj.SecurityMonitoringSignalRulePayload != nil && obj.SecurityMonitoringSignalRulePayload.UnparsedObject == nil {
			jsonSecurityMonitoringSignalRulePayload, _ := datadog.Marshal(obj.SecurityMonitoringSignalRulePayload)
			if string(jsonSecurityMonitoringSignalRulePayload) == "{}" { // empty struct
				obj.SecurityMonitoringSignalRulePayload = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringSignalRulePayload = nil
		}
	} else {
		obj.SecurityMonitoringSignalRulePayload = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SecurityMonitoringStandardRulePayload = nil
		obj.SecurityMonitoringSignalRulePayload = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SecurityMonitoringRuleConvertPayload) MarshalJSON() ([]byte, error) {
	if obj.SecurityMonitoringStandardRulePayload != nil {
		return datadog.Marshal(&obj.SecurityMonitoringStandardRulePayload)
	}

	if obj.SecurityMonitoringSignalRulePayload != nil {
		return datadog.Marshal(&obj.SecurityMonitoringSignalRulePayload)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SecurityMonitoringRuleConvertPayload) GetActualInstance() interface{} {
	if obj.SecurityMonitoringStandardRulePayload != nil {
		return obj.SecurityMonitoringStandardRulePayload
	}

	if obj.SecurityMonitoringSignalRulePayload != nil {
		return obj.SecurityMonitoringSignalRulePayload
	}

	// all schemas are nil
	return nil
}
