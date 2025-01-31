// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsBatchStatus Determines whether the batch has passed, failed, or is in progress.
type SyntheticsBatchStatus string

// List of SyntheticsBatchStatus.
const (
	SYNTHETICSBATCHSTATUS_PASSED  SyntheticsBatchStatus = "passed"
	SYNTHETICSBATCHSTATUS_SKIPPED SyntheticsBatchStatus = "skipped"
	SYNTHETICSBATCHSTATUS_FAILED  SyntheticsBatchStatus = "failed"
)

var allowedSyntheticsBatchStatusEnumValues = []SyntheticsBatchStatus{
	SYNTHETICSBATCHSTATUS_PASSED,
	SYNTHETICSBATCHSTATUS_SKIPPED,
	SYNTHETICSBATCHSTATUS_FAILED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsBatchStatus) GetAllowedValues() []SyntheticsBatchStatus {
	return allowedSyntheticsBatchStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsBatchStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsBatchStatus(value)
	return nil
}

// NewSyntheticsBatchStatusFromValue returns a pointer to a valid SyntheticsBatchStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsBatchStatusFromValue(v string) (*SyntheticsBatchStatus, error) {
	ev := SyntheticsBatchStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsBatchStatus: valid values are %v", v, allowedSyntheticsBatchStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsBatchStatus) IsValid() bool {
	for _, existing := range allowedSyntheticsBatchStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsBatchStatus value.
func (v SyntheticsBatchStatus) Ptr() *SyntheticsBatchStatus {
	return &v
}
