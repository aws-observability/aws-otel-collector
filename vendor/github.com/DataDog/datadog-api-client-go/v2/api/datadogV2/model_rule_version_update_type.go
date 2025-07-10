// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RuleVersionUpdateType The type of change.
type RuleVersionUpdateType string

// List of RuleVersionUpdateType.
const (
	RULEVERSIONUPDATETYPE_CREATE RuleVersionUpdateType = "create"
	RULEVERSIONUPDATETYPE_UPDATE RuleVersionUpdateType = "update"
	RULEVERSIONUPDATETYPE_DELETE RuleVersionUpdateType = "delete"
)

var allowedRuleVersionUpdateTypeEnumValues = []RuleVersionUpdateType{
	RULEVERSIONUPDATETYPE_CREATE,
	RULEVERSIONUPDATETYPE_UPDATE,
	RULEVERSIONUPDATETYPE_DELETE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RuleVersionUpdateType) GetAllowedValues() []RuleVersionUpdateType {
	return allowedRuleVersionUpdateTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RuleVersionUpdateType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RuleVersionUpdateType(value)
	return nil
}

// NewRuleVersionUpdateTypeFromValue returns a pointer to a valid RuleVersionUpdateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRuleVersionUpdateTypeFromValue(v string) (*RuleVersionUpdateType, error) {
	ev := RuleVersionUpdateType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RuleVersionUpdateType: valid values are %v", v, allowedRuleVersionUpdateTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RuleVersionUpdateType) IsValid() bool {
	for _, existing := range allowedRuleVersionUpdateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RuleVersionUpdateType value.
func (v RuleVersionUpdateType) Ptr() *RuleVersionUpdateType {
	return &v
}
