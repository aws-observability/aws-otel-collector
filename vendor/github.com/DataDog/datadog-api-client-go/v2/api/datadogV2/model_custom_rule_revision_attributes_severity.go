// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRuleRevisionAttributesSeverity Rule severity
type CustomRuleRevisionAttributesSeverity string

// List of CustomRuleRevisionAttributesSeverity.
const (
	CUSTOMRULEREVISIONATTRIBUTESSEVERITY_ERROR   CustomRuleRevisionAttributesSeverity = "ERROR"
	CUSTOMRULEREVISIONATTRIBUTESSEVERITY_WARNING CustomRuleRevisionAttributesSeverity = "WARNING"
	CUSTOMRULEREVISIONATTRIBUTESSEVERITY_NOTICE  CustomRuleRevisionAttributesSeverity = "NOTICE"
)

var allowedCustomRuleRevisionAttributesSeverityEnumValues = []CustomRuleRevisionAttributesSeverity{
	CUSTOMRULEREVISIONATTRIBUTESSEVERITY_ERROR,
	CUSTOMRULEREVISIONATTRIBUTESSEVERITY_WARNING,
	CUSTOMRULEREVISIONATTRIBUTESSEVERITY_NOTICE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomRuleRevisionAttributesSeverity) GetAllowedValues() []CustomRuleRevisionAttributesSeverity {
	return allowedCustomRuleRevisionAttributesSeverityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomRuleRevisionAttributesSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomRuleRevisionAttributesSeverity(value)
	return nil
}

// NewCustomRuleRevisionAttributesSeverityFromValue returns a pointer to a valid CustomRuleRevisionAttributesSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomRuleRevisionAttributesSeverityFromValue(v string) (*CustomRuleRevisionAttributesSeverity, error) {
	ev := CustomRuleRevisionAttributesSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomRuleRevisionAttributesSeverity: valid values are %v", v, allowedCustomRuleRevisionAttributesSeverityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomRuleRevisionAttributesSeverity) IsValid() bool {
	for _, existing := range allowedCustomRuleRevisionAttributesSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomRuleRevisionAttributesSeverity value.
func (v CustomRuleRevisionAttributesSeverity) Ptr() *CustomRuleRevisionAttributesSeverity {
	return &v
}
