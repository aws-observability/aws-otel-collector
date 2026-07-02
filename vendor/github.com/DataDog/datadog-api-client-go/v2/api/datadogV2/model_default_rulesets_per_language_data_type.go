// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DefaultRulesetsPerLanguageDataType Default rulesets per language resource type.
type DefaultRulesetsPerLanguageDataType string

// List of DefaultRulesetsPerLanguageDataType.
const (
	DEFAULTRULESETSPERLANGUAGEDATATYPE_DEFAULT_RULESETS_PER_LANGUAGE DefaultRulesetsPerLanguageDataType = "defaultRulesetsPerLanguage"
)

var allowedDefaultRulesetsPerLanguageDataTypeEnumValues = []DefaultRulesetsPerLanguageDataType{
	DEFAULTRULESETSPERLANGUAGEDATATYPE_DEFAULT_RULESETS_PER_LANGUAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DefaultRulesetsPerLanguageDataType) GetAllowedValues() []DefaultRulesetsPerLanguageDataType {
	return allowedDefaultRulesetsPerLanguageDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DefaultRulesetsPerLanguageDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DefaultRulesetsPerLanguageDataType(value)
	return nil
}

// NewDefaultRulesetsPerLanguageDataTypeFromValue returns a pointer to a valid DefaultRulesetsPerLanguageDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDefaultRulesetsPerLanguageDataTypeFromValue(v string) (*DefaultRulesetsPerLanguageDataType, error) {
	ev := DefaultRulesetsPerLanguageDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DefaultRulesetsPerLanguageDataType: valid values are %v", v, allowedDefaultRulesetsPerLanguageDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DefaultRulesetsPerLanguageDataType) IsValid() bool {
	for _, existing := range allowedDefaultRulesetsPerLanguageDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DefaultRulesetsPerLanguageDataType value.
func (v DefaultRulesetsPerLanguageDataType) Ptr() *DefaultRulesetsPerLanguageDataType {
	return &v
}
