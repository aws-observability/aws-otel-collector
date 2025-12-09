// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SignalArchiveReason Reason why a signal has been archived.
type SignalArchiveReason string

// List of SignalArchiveReason.
const (
	SIGNALARCHIVEREASON_NONE                     SignalArchiveReason = "none"
	SIGNALARCHIVEREASON_FALSE_POSITIVE           SignalArchiveReason = "false_positive"
	SIGNALARCHIVEREASON_TESTING_OR_MAINTENANCE   SignalArchiveReason = "testing_or_maintenance"
	SIGNALARCHIVEREASON_INVESTIGATED_CASE_OPENED SignalArchiveReason = "investigated_case_opened"
	SIGNALARCHIVEREASON_TRUE_POSITIVE_BENIGN     SignalArchiveReason = "true_positive_benign"
	SIGNALARCHIVEREASON_TRUE_POSITIVE_MALICIOUS  SignalArchiveReason = "true_positive_malicious"
	SIGNALARCHIVEREASON_OTHER                    SignalArchiveReason = "other"
)

var allowedSignalArchiveReasonEnumValues = []SignalArchiveReason{
	SIGNALARCHIVEREASON_NONE,
	SIGNALARCHIVEREASON_FALSE_POSITIVE,
	SIGNALARCHIVEREASON_TESTING_OR_MAINTENANCE,
	SIGNALARCHIVEREASON_INVESTIGATED_CASE_OPENED,
	SIGNALARCHIVEREASON_TRUE_POSITIVE_BENIGN,
	SIGNALARCHIVEREASON_TRUE_POSITIVE_MALICIOUS,
	SIGNALARCHIVEREASON_OTHER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SignalArchiveReason) GetAllowedValues() []SignalArchiveReason {
	return allowedSignalArchiveReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SignalArchiveReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SignalArchiveReason(value)
	return nil
}

// NewSignalArchiveReasonFromValue returns a pointer to a valid SignalArchiveReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSignalArchiveReasonFromValue(v string) (*SignalArchiveReason, error) {
	ev := SignalArchiveReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SignalArchiveReason: valid values are %v", v, allowedSignalArchiveReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SignalArchiveReason) IsValid() bool {
	for _, existing := range allowedSignalArchiveReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SignalArchiveReason value.
func (v SignalArchiveReason) Ptr() *SignalArchiveReason {
	return &v
}
