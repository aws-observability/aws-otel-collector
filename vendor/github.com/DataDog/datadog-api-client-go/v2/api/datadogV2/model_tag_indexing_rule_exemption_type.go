// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleExemptionType The tag indexing rule exemption resource type.
type TagIndexingRuleExemptionType string

// List of TagIndexingRuleExemptionType.
const (
	TAGINDEXINGRULEEXEMPTIONTYPE_TAG_INDEXING_RULE_EXEMPTIONS TagIndexingRuleExemptionType = "tag_indexing_rule_exemptions"
)

var allowedTagIndexingRuleExemptionTypeEnumValues = []TagIndexingRuleExemptionType{
	TAGINDEXINGRULEEXEMPTIONTYPE_TAG_INDEXING_RULE_EXEMPTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagIndexingRuleExemptionType) GetAllowedValues() []TagIndexingRuleExemptionType {
	return allowedTagIndexingRuleExemptionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagIndexingRuleExemptionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagIndexingRuleExemptionType(value)
	return nil
}

// NewTagIndexingRuleExemptionTypeFromValue returns a pointer to a valid TagIndexingRuleExemptionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagIndexingRuleExemptionTypeFromValue(v string) (*TagIndexingRuleExemptionType, error) {
	ev := TagIndexingRuleExemptionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagIndexingRuleExemptionType: valid values are %v", v, allowedTagIndexingRuleExemptionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagIndexingRuleExemptionType) IsValid() bool {
	for _, existing := range allowedTagIndexingRuleExemptionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagIndexingRuleExemptionType value.
func (v TagIndexingRuleExemptionType) Ptr() *TagIndexingRuleExemptionType {
	return &v
}
