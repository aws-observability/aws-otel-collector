// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileStepParamsDirection The direction of the scroll for a `scrollToElement` step type.
type SyntheticsMobileStepParamsDirection string

// List of SyntheticsMobileStepParamsDirection.
const (
	SYNTHETICSMOBILESTEPPARAMSDIRECTION_UP    SyntheticsMobileStepParamsDirection = "up"
	SYNTHETICSMOBILESTEPPARAMSDIRECTION_DOWN  SyntheticsMobileStepParamsDirection = "down"
	SYNTHETICSMOBILESTEPPARAMSDIRECTION_LEFT  SyntheticsMobileStepParamsDirection = "left"
	SYNTHETICSMOBILESTEPPARAMSDIRECTION_RIGHT SyntheticsMobileStepParamsDirection = "right"
)

var allowedSyntheticsMobileStepParamsDirectionEnumValues = []SyntheticsMobileStepParamsDirection{
	SYNTHETICSMOBILESTEPPARAMSDIRECTION_UP,
	SYNTHETICSMOBILESTEPPARAMSDIRECTION_DOWN,
	SYNTHETICSMOBILESTEPPARAMSDIRECTION_LEFT,
	SYNTHETICSMOBILESTEPPARAMSDIRECTION_RIGHT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsMobileStepParamsDirection) GetAllowedValues() []SyntheticsMobileStepParamsDirection {
	return allowedSyntheticsMobileStepParamsDirectionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsMobileStepParamsDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsMobileStepParamsDirection(value)
	return nil
}

// NewSyntheticsMobileStepParamsDirectionFromValue returns a pointer to a valid SyntheticsMobileStepParamsDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsMobileStepParamsDirectionFromValue(v string) (*SyntheticsMobileStepParamsDirection, error) {
	ev := SyntheticsMobileStepParamsDirection(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsMobileStepParamsDirection: valid values are %v", v, allowedSyntheticsMobileStepParamsDirectionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsMobileStepParamsDirection) IsValid() bool {
	for _, existing := range allowedSyntheticsMobileStepParamsDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsMobileStepParamsDirection value.
func (v SyntheticsMobileStepParamsDirection) Ptr() *SyntheticsMobileStepParamsDirection {
	return &v
}
