// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SuiteJsonPatchType Type for a JSON Patch request on a Synthetic test suite, `suites_json_patch`.
type SuiteJsonPatchType string

// List of SuiteJsonPatchType.
const (
	SUITEJSONPATCHTYPE_SUITES_JSON_PATCH SuiteJsonPatchType = "suites_json_patch"
)

var allowedSuiteJsonPatchTypeEnumValues = []SuiteJsonPatchType{
	SUITEJSONPATCHTYPE_SUITES_JSON_PATCH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SuiteJsonPatchType) GetAllowedValues() []SuiteJsonPatchType {
	return allowedSuiteJsonPatchTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SuiteJsonPatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SuiteJsonPatchType(value)
	return nil
}

// NewSuiteJsonPatchTypeFromValue returns a pointer to a valid SuiteJsonPatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSuiteJsonPatchTypeFromValue(v string) (*SuiteJsonPatchType, error) {
	ev := SuiteJsonPatchType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SuiteJsonPatchType: valid values are %v", v, allowedSuiteJsonPatchTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SuiteJsonPatchType) IsValid() bool {
	for _, existing := range allowedSuiteJsonPatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SuiteJsonPatchType value.
func (v SuiteJsonPatchType) Ptr() *SuiteJsonPatchType {
	return &v
}
