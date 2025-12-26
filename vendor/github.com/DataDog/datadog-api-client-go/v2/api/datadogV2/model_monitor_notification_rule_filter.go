// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorNotificationRuleFilter - Filter used to associate the notification rule with monitors.
type MonitorNotificationRuleFilter struct {
	MonitorNotificationRuleFilterTags  *MonitorNotificationRuleFilterTags
	MonitorNotificationRuleFilterScope *MonitorNotificationRuleFilterScope

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MonitorNotificationRuleFilterTagsAsMonitorNotificationRuleFilter is a convenience function that returns MonitorNotificationRuleFilterTags wrapped in MonitorNotificationRuleFilter.
func MonitorNotificationRuleFilterTagsAsMonitorNotificationRuleFilter(v *MonitorNotificationRuleFilterTags) MonitorNotificationRuleFilter {
	return MonitorNotificationRuleFilter{MonitorNotificationRuleFilterTags: v}
}

// MonitorNotificationRuleFilterScopeAsMonitorNotificationRuleFilter is a convenience function that returns MonitorNotificationRuleFilterScope wrapped in MonitorNotificationRuleFilter.
func MonitorNotificationRuleFilterScopeAsMonitorNotificationRuleFilter(v *MonitorNotificationRuleFilterScope) MonitorNotificationRuleFilter {
	return MonitorNotificationRuleFilter{MonitorNotificationRuleFilterScope: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *MonitorNotificationRuleFilter) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MonitorNotificationRuleFilterTags
	err = datadog.Unmarshal(data, &obj.MonitorNotificationRuleFilterTags)
	if err == nil {
		if obj.MonitorNotificationRuleFilterTags != nil && obj.MonitorNotificationRuleFilterTags.UnparsedObject == nil {
			jsonMonitorNotificationRuleFilterTags, _ := datadog.Marshal(obj.MonitorNotificationRuleFilterTags)
			if string(jsonMonitorNotificationRuleFilterTags) == "{}" { // empty struct
				obj.MonitorNotificationRuleFilterTags = nil
			} else {
				match++
			}
		} else {
			obj.MonitorNotificationRuleFilterTags = nil
		}
	} else {
		obj.MonitorNotificationRuleFilterTags = nil
	}

	// try to unmarshal data into MonitorNotificationRuleFilterScope
	err = datadog.Unmarshal(data, &obj.MonitorNotificationRuleFilterScope)
	if err == nil {
		if obj.MonitorNotificationRuleFilterScope != nil && obj.MonitorNotificationRuleFilterScope.UnparsedObject == nil {
			jsonMonitorNotificationRuleFilterScope, _ := datadog.Marshal(obj.MonitorNotificationRuleFilterScope)
			if string(jsonMonitorNotificationRuleFilterScope) == "{}" { // empty struct
				obj.MonitorNotificationRuleFilterScope = nil
			} else {
				match++
			}
		} else {
			obj.MonitorNotificationRuleFilterScope = nil
		}
	} else {
		obj.MonitorNotificationRuleFilterScope = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MonitorNotificationRuleFilterTags = nil
		obj.MonitorNotificationRuleFilterScope = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj MonitorNotificationRuleFilter) MarshalJSON() ([]byte, error) {
	if obj.MonitorNotificationRuleFilterTags != nil {
		return datadog.Marshal(&obj.MonitorNotificationRuleFilterTags)
	}

	if obj.MonitorNotificationRuleFilterScope != nil {
		return datadog.Marshal(&obj.MonitorNotificationRuleFilterScope)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *MonitorNotificationRuleFilter) GetActualInstance() interface{} {
	if obj.MonitorNotificationRuleFilterTags != nil {
		return obj.MonitorNotificationRuleFilterTags
	}

	if obj.MonitorNotificationRuleFilterScope != nil {
		return obj.MonitorNotificationRuleFilterScope
	}

	// all schemas are nil
	return nil
}
