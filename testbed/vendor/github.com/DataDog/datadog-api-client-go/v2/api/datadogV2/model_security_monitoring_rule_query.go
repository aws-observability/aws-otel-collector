// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleQuery - Query for matching rule.
type SecurityMonitoringRuleQuery struct {
	SecurityMonitoringStandardRuleQuery *SecurityMonitoringStandardRuleQuery
	SecurityMonitoringSignalRuleQuery   *SecurityMonitoringSignalRuleQuery

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SecurityMonitoringStandardRuleQueryAsSecurityMonitoringRuleQuery is a convenience function that returns SecurityMonitoringStandardRuleQuery wrapped in SecurityMonitoringRuleQuery.
func SecurityMonitoringStandardRuleQueryAsSecurityMonitoringRuleQuery(v *SecurityMonitoringStandardRuleQuery) SecurityMonitoringRuleQuery {
	return SecurityMonitoringRuleQuery{SecurityMonitoringStandardRuleQuery: v}
}

// SecurityMonitoringSignalRuleQueryAsSecurityMonitoringRuleQuery is a convenience function that returns SecurityMonitoringSignalRuleQuery wrapped in SecurityMonitoringRuleQuery.
func SecurityMonitoringSignalRuleQueryAsSecurityMonitoringRuleQuery(v *SecurityMonitoringSignalRuleQuery) SecurityMonitoringRuleQuery {
	return SecurityMonitoringRuleQuery{SecurityMonitoringSignalRuleQuery: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SecurityMonitoringRuleQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SecurityMonitoringStandardRuleQuery
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringStandardRuleQuery)
	if err == nil {
		if obj.SecurityMonitoringStandardRuleQuery != nil && obj.SecurityMonitoringStandardRuleQuery.UnparsedObject == nil {
			jsonSecurityMonitoringStandardRuleQuery, _ := datadog.Marshal(obj.SecurityMonitoringStandardRuleQuery)
			if string(jsonSecurityMonitoringStandardRuleQuery) == "{}" && string(data) != "{}" { // empty struct
				obj.SecurityMonitoringStandardRuleQuery = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringStandardRuleQuery = nil
		}
	} else {
		obj.SecurityMonitoringStandardRuleQuery = nil
	}

	// try to unmarshal data into SecurityMonitoringSignalRuleQuery
	err = datadog.Unmarshal(data, &obj.SecurityMonitoringSignalRuleQuery)
	if err == nil {
		if obj.SecurityMonitoringSignalRuleQuery != nil && obj.SecurityMonitoringSignalRuleQuery.UnparsedObject == nil {
			jsonSecurityMonitoringSignalRuleQuery, _ := datadog.Marshal(obj.SecurityMonitoringSignalRuleQuery)
			if string(jsonSecurityMonitoringSignalRuleQuery) == "{}" { // empty struct
				obj.SecurityMonitoringSignalRuleQuery = nil
			} else {
				match++
			}
		} else {
			obj.SecurityMonitoringSignalRuleQuery = nil
		}
	} else {
		obj.SecurityMonitoringSignalRuleQuery = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SecurityMonitoringStandardRuleQuery = nil
		obj.SecurityMonitoringSignalRuleQuery = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SecurityMonitoringRuleQuery) MarshalJSON() ([]byte, error) {
	if obj.SecurityMonitoringStandardRuleQuery != nil {
		return datadog.Marshal(&obj.SecurityMonitoringStandardRuleQuery)
	}

	if obj.SecurityMonitoringSignalRuleQuery != nil {
		return datadog.Marshal(&obj.SecurityMonitoringSignalRuleQuery)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SecurityMonitoringRuleQuery) GetActualInstance() interface{} {
	if obj.SecurityMonitoringStandardRuleQuery != nil {
		return obj.SecurityMonitoringStandardRuleQuery
	}

	if obj.SecurityMonitoringSignalRuleQuery != nil {
		return obj.SecurityMonitoringSignalRuleQuery
	}

	// all schemas are nil
	return nil
}
