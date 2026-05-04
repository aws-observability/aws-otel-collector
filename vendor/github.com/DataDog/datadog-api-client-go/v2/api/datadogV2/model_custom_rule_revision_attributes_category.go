// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRuleRevisionAttributesCategory Rule category
type CustomRuleRevisionAttributesCategory string

// List of CustomRuleRevisionAttributesCategory.
const (
	CUSTOMRULEREVISIONATTRIBUTESCATEGORY_SECURITY       CustomRuleRevisionAttributesCategory = "SECURITY"
	CUSTOMRULEREVISIONATTRIBUTESCATEGORY_BEST_PRACTICES CustomRuleRevisionAttributesCategory = "BEST_PRACTICES"
	CUSTOMRULEREVISIONATTRIBUTESCATEGORY_CODE_STYLE     CustomRuleRevisionAttributesCategory = "CODE_STYLE"
	CUSTOMRULEREVISIONATTRIBUTESCATEGORY_ERROR_PRONE    CustomRuleRevisionAttributesCategory = "ERROR_PRONE"
	CUSTOMRULEREVISIONATTRIBUTESCATEGORY_PERFORMANCE    CustomRuleRevisionAttributesCategory = "PERFORMANCE"
)

var allowedCustomRuleRevisionAttributesCategoryEnumValues = []CustomRuleRevisionAttributesCategory{
	CUSTOMRULEREVISIONATTRIBUTESCATEGORY_SECURITY,
	CUSTOMRULEREVISIONATTRIBUTESCATEGORY_BEST_PRACTICES,
	CUSTOMRULEREVISIONATTRIBUTESCATEGORY_CODE_STYLE,
	CUSTOMRULEREVISIONATTRIBUTESCATEGORY_ERROR_PRONE,
	CUSTOMRULEREVISIONATTRIBUTESCATEGORY_PERFORMANCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomRuleRevisionAttributesCategory) GetAllowedValues() []CustomRuleRevisionAttributesCategory {
	return allowedCustomRuleRevisionAttributesCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomRuleRevisionAttributesCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomRuleRevisionAttributesCategory(value)
	return nil
}

// NewCustomRuleRevisionAttributesCategoryFromValue returns a pointer to a valid CustomRuleRevisionAttributesCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomRuleRevisionAttributesCategoryFromValue(v string) (*CustomRuleRevisionAttributesCategory, error) {
	ev := CustomRuleRevisionAttributesCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomRuleRevisionAttributesCategory: valid values are %v", v, allowedCustomRuleRevisionAttributesCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomRuleRevisionAttributesCategory) IsValid() bool {
	for _, existing := range allowedCustomRuleRevisionAttributesCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomRuleRevisionAttributesCategory value.
func (v CustomRuleRevisionAttributesCategory) Ptr() *CustomRuleRevisionAttributesCategory {
	return &v
}
