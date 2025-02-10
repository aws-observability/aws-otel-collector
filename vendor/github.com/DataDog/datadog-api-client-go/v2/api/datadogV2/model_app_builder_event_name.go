// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppBuilderEventName The triggering action for the event.
type AppBuilderEventName string

// List of AppBuilderEventName.
const (
	APPBUILDEREVENTNAME_PAGECHANGE          AppBuilderEventName = "pageChange"
	APPBUILDEREVENTNAME_TABLEROWCLICK       AppBuilderEventName = "tableRowClick"
	APPBUILDEREVENTNAME_TABLEROWBUTTONCLICK AppBuilderEventName = "_tableRowButtonClick"
	APPBUILDEREVENTNAME_CHANGE              AppBuilderEventName = "change"
	APPBUILDEREVENTNAME_SUBMIT              AppBuilderEventName = "submit"
	APPBUILDEREVENTNAME_CLICK               AppBuilderEventName = "click"
	APPBUILDEREVENTNAME_TOGGLEOPEN          AppBuilderEventName = "toggleOpen"
	APPBUILDEREVENTNAME_CLOSE               AppBuilderEventName = "close"
	APPBUILDEREVENTNAME_OPEN                AppBuilderEventName = "open"
	APPBUILDEREVENTNAME_EXECUTIONFINISHED   AppBuilderEventName = "executionFinished"
)

var allowedAppBuilderEventNameEnumValues = []AppBuilderEventName{
	APPBUILDEREVENTNAME_PAGECHANGE,
	APPBUILDEREVENTNAME_TABLEROWCLICK,
	APPBUILDEREVENTNAME_TABLEROWBUTTONCLICK,
	APPBUILDEREVENTNAME_CHANGE,
	APPBUILDEREVENTNAME_SUBMIT,
	APPBUILDEREVENTNAME_CLICK,
	APPBUILDEREVENTNAME_TOGGLEOPEN,
	APPBUILDEREVENTNAME_CLOSE,
	APPBUILDEREVENTNAME_OPEN,
	APPBUILDEREVENTNAME_EXECUTIONFINISHED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AppBuilderEventName) GetAllowedValues() []AppBuilderEventName {
	return allowedAppBuilderEventNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AppBuilderEventName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AppBuilderEventName(value)
	return nil
}

// NewAppBuilderEventNameFromValue returns a pointer to a valid AppBuilderEventName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAppBuilderEventNameFromValue(v string) (*AppBuilderEventName, error) {
	ev := AppBuilderEventName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AppBuilderEventName: valid values are %v", v, allowedAppBuilderEventNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AppBuilderEventName) IsValid() bool {
	for _, existing := range allowedAppBuilderEventNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppBuilderEventName value.
func (v AppBuilderEventName) Ptr() *AppBuilderEventName {
	return &v
}
