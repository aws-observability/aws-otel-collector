// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppBuilderEventType The response to the event.
type AppBuilderEventType string

// List of AppBuilderEventType.
const (
	APPBUILDEREVENTTYPE_CUSTOM                AppBuilderEventType = "custom"
	APPBUILDEREVENTTYPE_SETCOMPONENTSTATE     AppBuilderEventType = "setComponentState"
	APPBUILDEREVENTTYPE_TRIGGERQUERY          AppBuilderEventType = "triggerQuery"
	APPBUILDEREVENTTYPE_OPENMODAL             AppBuilderEventType = "openModal"
	APPBUILDEREVENTTYPE_CLOSEMODAL            AppBuilderEventType = "closeModal"
	APPBUILDEREVENTTYPE_OPENURL               AppBuilderEventType = "openUrl"
	APPBUILDEREVENTTYPE_DOWNLOADFILE          AppBuilderEventType = "downloadFile"
	APPBUILDEREVENTTYPE_SETSTATEVARIABLEVALUE AppBuilderEventType = "setStateVariableValue"
)

var allowedAppBuilderEventTypeEnumValues = []AppBuilderEventType{
	APPBUILDEREVENTTYPE_CUSTOM,
	APPBUILDEREVENTTYPE_SETCOMPONENTSTATE,
	APPBUILDEREVENTTYPE_TRIGGERQUERY,
	APPBUILDEREVENTTYPE_OPENMODAL,
	APPBUILDEREVENTTYPE_CLOSEMODAL,
	APPBUILDEREVENTTYPE_OPENURL,
	APPBUILDEREVENTTYPE_DOWNLOADFILE,
	APPBUILDEREVENTTYPE_SETSTATEVARIABLEVALUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AppBuilderEventType) GetAllowedValues() []AppBuilderEventType {
	return allowedAppBuilderEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AppBuilderEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AppBuilderEventType(value)
	return nil
}

// NewAppBuilderEventTypeFromValue returns a pointer to a valid AppBuilderEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAppBuilderEventTypeFromValue(v string) (*AppBuilderEventType, error) {
	ev := AppBuilderEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AppBuilderEventType: valid values are %v", v, allowedAppBuilderEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AppBuilderEventType) IsValid() bool {
	for _, existing := range allowedAppBuilderEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppBuilderEventType value.
func (v AppBuilderEventType) Ptr() *AppBuilderEventType {
	return &v
}
