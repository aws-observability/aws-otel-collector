// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsFastTestType Type of the Synthetic fast test that produced this result.
type SyntheticsFastTestType string

// List of SyntheticsFastTestType.
const (
	SYNTHETICSFASTTESTTYPE_FAST_API     SyntheticsFastTestType = "fast-api"
	SYNTHETICSFASTTESTTYPE_FAST_BROWSER SyntheticsFastTestType = "fast-browser"
)

var allowedSyntheticsFastTestTypeEnumValues = []SyntheticsFastTestType{
	SYNTHETICSFASTTESTTYPE_FAST_API,
	SYNTHETICSFASTTESTTYPE_FAST_BROWSER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsFastTestType) GetAllowedValues() []SyntheticsFastTestType {
	return allowedSyntheticsFastTestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsFastTestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsFastTestType(value)
	return nil
}

// NewSyntheticsFastTestTypeFromValue returns a pointer to a valid SyntheticsFastTestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsFastTestTypeFromValue(v string) (*SyntheticsFastTestType, error) {
	ev := SyntheticsFastTestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsFastTestType: valid values are %v", v, allowedSyntheticsFastTestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsFastTestType) IsValid() bool {
	for _, existing := range allowedSyntheticsFastTestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsFastTestType value.
func (v SyntheticsFastTestType) Ptr() *SyntheticsFastTestType {
	return &v
}
