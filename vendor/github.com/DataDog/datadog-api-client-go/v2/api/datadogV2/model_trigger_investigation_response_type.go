// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TriggerInvestigationResponseType The resource type for trigger investigation responses.
type TriggerInvestigationResponseType string

// List of TriggerInvestigationResponseType.
const (
	TRIGGERINVESTIGATIONRESPONSETYPE_TRIGGER_INVESTIGATION_RESPONSE TriggerInvestigationResponseType = "trigger_investigation_response"
)

var allowedTriggerInvestigationResponseTypeEnumValues = []TriggerInvestigationResponseType{
	TRIGGERINVESTIGATIONRESPONSETYPE_TRIGGER_INVESTIGATION_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TriggerInvestigationResponseType) GetAllowedValues() []TriggerInvestigationResponseType {
	return allowedTriggerInvestigationResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TriggerInvestigationResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TriggerInvestigationResponseType(value)
	return nil
}

// NewTriggerInvestigationResponseTypeFromValue returns a pointer to a valid TriggerInvestigationResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTriggerInvestigationResponseTypeFromValue(v string) (*TriggerInvestigationResponseType, error) {
	ev := TriggerInvestigationResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TriggerInvestigationResponseType: valid values are %v", v, allowedTriggerInvestigationResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TriggerInvestigationResponseType) IsValid() bool {
	for _, existing := range allowedTriggerInvestigationResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TriggerInvestigationResponseType value.
func (v TriggerInvestigationResponseType) Ptr() *TriggerInvestigationResponseType {
	return &v
}
