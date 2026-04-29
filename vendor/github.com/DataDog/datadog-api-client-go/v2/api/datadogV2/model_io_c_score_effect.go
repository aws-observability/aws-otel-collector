// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IoCScoreEffect Effect of a scoring factor on the indicator's threat score.
type IoCScoreEffect string

// List of IoCScoreEffect.
const (
	IOCSCOREEFFECT_RAISE_SCORE IoCScoreEffect = "RAISE_SCORE"
	IOCSCOREEFFECT_LOWER_SCORE IoCScoreEffect = "LOWER_SCORE"
	IOCSCOREEFFECT_NO_EFFECT   IoCScoreEffect = "NO_EFFECT"
)

var allowedIoCScoreEffectEnumValues = []IoCScoreEffect{
	IOCSCOREEFFECT_RAISE_SCORE,
	IOCSCOREEFFECT_LOWER_SCORE,
	IOCSCOREEFFECT_NO_EFFECT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IoCScoreEffect) GetAllowedValues() []IoCScoreEffect {
	return allowedIoCScoreEffectEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IoCScoreEffect) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IoCScoreEffect(value)
	return nil
}

// NewIoCScoreEffectFromValue returns a pointer to a valid IoCScoreEffect
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIoCScoreEffectFromValue(v string) (*IoCScoreEffect, error) {
	ev := IoCScoreEffect(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IoCScoreEffect: valid values are %v", v, allowedIoCScoreEffectEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IoCScoreEffect) IsValid() bool {
	for _, existing := range allowedIoCScoreEffectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IoCScoreEffect value.
func (v IoCScoreEffect) Ptr() *IoCScoreEffect {
	return &v
}
