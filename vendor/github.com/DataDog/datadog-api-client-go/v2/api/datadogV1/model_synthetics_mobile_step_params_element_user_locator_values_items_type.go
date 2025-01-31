// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileStepParamsElementUserLocatorValuesItemsType Type of a user locator.
type SyntheticsMobileStepParamsElementUserLocatorValuesItemsType string

// List of SyntheticsMobileStepParamsElementUserLocatorValuesItemsType.
const (
	SYNTHETICSMOBILESTEPPARAMSELEMENTUSERLOCATORVALUESITEMSTYPE_ACCESSIBILITY_ID     SyntheticsMobileStepParamsElementUserLocatorValuesItemsType = "accessibility-id"
	SYNTHETICSMOBILESTEPPARAMSELEMENTUSERLOCATORVALUESITEMSTYPE_ID                   SyntheticsMobileStepParamsElementUserLocatorValuesItemsType = "id"
	SYNTHETICSMOBILESTEPPARAMSELEMENTUSERLOCATORVALUESITEMSTYPE_IOS_PREDICATE_STRING SyntheticsMobileStepParamsElementUserLocatorValuesItemsType = "ios-predicate-string"
	SYNTHETICSMOBILESTEPPARAMSELEMENTUSERLOCATORVALUESITEMSTYPE_IOS_CLASS_CHAIN      SyntheticsMobileStepParamsElementUserLocatorValuesItemsType = "ios-class-chain"
	SYNTHETICSMOBILESTEPPARAMSELEMENTUSERLOCATORVALUESITEMSTYPE_XPATH                SyntheticsMobileStepParamsElementUserLocatorValuesItemsType = "xpath"
)

var allowedSyntheticsMobileStepParamsElementUserLocatorValuesItemsTypeEnumValues = []SyntheticsMobileStepParamsElementUserLocatorValuesItemsType{
	SYNTHETICSMOBILESTEPPARAMSELEMENTUSERLOCATORVALUESITEMSTYPE_ACCESSIBILITY_ID,
	SYNTHETICSMOBILESTEPPARAMSELEMENTUSERLOCATORVALUESITEMSTYPE_ID,
	SYNTHETICSMOBILESTEPPARAMSELEMENTUSERLOCATORVALUESITEMSTYPE_IOS_PREDICATE_STRING,
	SYNTHETICSMOBILESTEPPARAMSELEMENTUSERLOCATORVALUESITEMSTYPE_IOS_CLASS_CHAIN,
	SYNTHETICSMOBILESTEPPARAMSELEMENTUSERLOCATORVALUESITEMSTYPE_XPATH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsMobileStepParamsElementUserLocatorValuesItemsType) GetAllowedValues() []SyntheticsMobileStepParamsElementUserLocatorValuesItemsType {
	return allowedSyntheticsMobileStepParamsElementUserLocatorValuesItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsMobileStepParamsElementUserLocatorValuesItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsMobileStepParamsElementUserLocatorValuesItemsType(value)
	return nil
}

// NewSyntheticsMobileStepParamsElementUserLocatorValuesItemsTypeFromValue returns a pointer to a valid SyntheticsMobileStepParamsElementUserLocatorValuesItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsMobileStepParamsElementUserLocatorValuesItemsTypeFromValue(v string) (*SyntheticsMobileStepParamsElementUserLocatorValuesItemsType, error) {
	ev := SyntheticsMobileStepParamsElementUserLocatorValuesItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsMobileStepParamsElementUserLocatorValuesItemsType: valid values are %v", v, allowedSyntheticsMobileStepParamsElementUserLocatorValuesItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsMobileStepParamsElementUserLocatorValuesItemsType) IsValid() bool {
	for _, existing := range allowedSyntheticsMobileStepParamsElementUserLocatorValuesItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsMobileStepParamsElementUserLocatorValuesItemsType value.
func (v SyntheticsMobileStepParamsElementUserLocatorValuesItemsType) Ptr() *SyntheticsMobileStepParamsElementUserLocatorValuesItemsType {
	return &v
}
