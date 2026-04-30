// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsSuiteTestAlertingCriticality Alerting criticality for each the test.
type SyntheticsSuiteTestAlertingCriticality string

// List of SyntheticsSuiteTestAlertingCriticality.
const (
	SYNTHETICSSUITETESTALERTINGCRITICALITY_IGNORE   SyntheticsSuiteTestAlertingCriticality = "ignore"
	SYNTHETICSSUITETESTALERTINGCRITICALITY_CRITICAL SyntheticsSuiteTestAlertingCriticality = "critical"
)

var allowedSyntheticsSuiteTestAlertingCriticalityEnumValues = []SyntheticsSuiteTestAlertingCriticality{
	SYNTHETICSSUITETESTALERTINGCRITICALITY_IGNORE,
	SYNTHETICSSUITETESTALERTINGCRITICALITY_CRITICAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsSuiteTestAlertingCriticality) GetAllowedValues() []SyntheticsSuiteTestAlertingCriticality {
	return allowedSyntheticsSuiteTestAlertingCriticalityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsSuiteTestAlertingCriticality) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsSuiteTestAlertingCriticality(value)
	return nil
}

// NewSyntheticsSuiteTestAlertingCriticalityFromValue returns a pointer to a valid SyntheticsSuiteTestAlertingCriticality
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsSuiteTestAlertingCriticalityFromValue(v string) (*SyntheticsSuiteTestAlertingCriticality, error) {
	ev := SyntheticsSuiteTestAlertingCriticality(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsSuiteTestAlertingCriticality: valid values are %v", v, allowedSyntheticsSuiteTestAlertingCriticalityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsSuiteTestAlertingCriticality) IsValid() bool {
	for _, existing := range allowedSyntheticsSuiteTestAlertingCriticalityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsSuiteTestAlertingCriticality value.
func (v SyntheticsSuiteTestAlertingCriticality) Ptr() *SyntheticsSuiteTestAlertingCriticality {
	return &v
}
