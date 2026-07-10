// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampleLogGenerationDuration How long the subscription should remain active before expiring.
type SampleLogGenerationDuration string

// List of SampleLogGenerationDuration.
const (
	SAMPLELOGGENERATIONDURATION_ONE_HOUR   SampleLogGenerationDuration = "1h"
	SAMPLELOGGENERATIONDURATION_ONE_DAY    SampleLogGenerationDuration = "1d"
	SAMPLELOGGENERATIONDURATION_THREE_DAYS SampleLogGenerationDuration = "3d"
	SAMPLELOGGENERATIONDURATION_SEVEN_DAYS SampleLogGenerationDuration = "7d"
)

var allowedSampleLogGenerationDurationEnumValues = []SampleLogGenerationDuration{
	SAMPLELOGGENERATIONDURATION_ONE_HOUR,
	SAMPLELOGGENERATIONDURATION_ONE_DAY,
	SAMPLELOGGENERATIONDURATION_THREE_DAYS,
	SAMPLELOGGENERATIONDURATION_SEVEN_DAYS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SampleLogGenerationDuration) GetAllowedValues() []SampleLogGenerationDuration {
	return allowedSampleLogGenerationDurationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SampleLogGenerationDuration) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SampleLogGenerationDuration(value)
	return nil
}

// NewSampleLogGenerationDurationFromValue returns a pointer to a valid SampleLogGenerationDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSampleLogGenerationDurationFromValue(v string) (*SampleLogGenerationDuration, error) {
	ev := SampleLogGenerationDuration(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SampleLogGenerationDuration: valid values are %v", v, allowedSampleLogGenerationDurationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SampleLogGenerationDuration) IsValid() bool {
	for _, existing := range allowedSampleLogGenerationDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SampleLogGenerationDuration value.
func (v SampleLogGenerationDuration) Ptr() *SampleLogGenerationDuration {
	return &v
}
