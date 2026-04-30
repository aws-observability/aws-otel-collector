// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TriggerInvestigationRequestType The resource type for trigger investigation requests.
type TriggerInvestigationRequestType string

// List of TriggerInvestigationRequestType.
const (
	TRIGGERINVESTIGATIONREQUESTTYPE_TRIGGER_INVESTIGATION_REQUEST TriggerInvestigationRequestType = "trigger_investigation_request"
)

var allowedTriggerInvestigationRequestTypeEnumValues = []TriggerInvestigationRequestType{
	TRIGGERINVESTIGATIONREQUESTTYPE_TRIGGER_INVESTIGATION_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TriggerInvestigationRequestType) GetAllowedValues() []TriggerInvestigationRequestType {
	return allowedTriggerInvestigationRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TriggerInvestigationRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TriggerInvestigationRequestType(value)
	return nil
}

// NewTriggerInvestigationRequestTypeFromValue returns a pointer to a valid TriggerInvestigationRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTriggerInvestigationRequestTypeFromValue(v string) (*TriggerInvestigationRequestType, error) {
	ev := TriggerInvestigationRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TriggerInvestigationRequestType: valid values are %v", v, allowedTriggerInvestigationRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TriggerInvestigationRequestType) IsValid() bool {
	for _, existing := range allowedTriggerInvestigationRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TriggerInvestigationRequestType value.
func (v TriggerInvestigationRequestType) Ptr() *TriggerInvestigationRequestType {
	return &v
}
