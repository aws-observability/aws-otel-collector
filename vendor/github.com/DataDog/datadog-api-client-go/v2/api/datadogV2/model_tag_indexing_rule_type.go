// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleType The tag indexing rule resource type.
type TagIndexingRuleType string

// List of TagIndexingRuleType.
const (
	TAGINDEXINGRULETYPE_TAG_INDEXING_RULES TagIndexingRuleType = "tag_indexing_rules"
)

var allowedTagIndexingRuleTypeEnumValues = []TagIndexingRuleType{
	TAGINDEXINGRULETYPE_TAG_INDEXING_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagIndexingRuleType) GetAllowedValues() []TagIndexingRuleType {
	return allowedTagIndexingRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagIndexingRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagIndexingRuleType(value)
	return nil
}

// NewTagIndexingRuleTypeFromValue returns a pointer to a valid TagIndexingRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagIndexingRuleTypeFromValue(v string) (*TagIndexingRuleType, error) {
	ev := TagIndexingRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagIndexingRuleType: valid values are %v", v, allowedTagIndexingRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagIndexingRuleType) IsValid() bool {
	for _, existing := range allowedTagIndexingRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagIndexingRuleType value.
func (v TagIndexingRuleType) Ptr() *TagIndexingRuleType {
	return &v
}
