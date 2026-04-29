// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestHistoryPolicyId The policy that triggered this status change.
type FlakyTestHistoryPolicyId string

// List of FlakyTestHistoryPolicyId.
const (
	FLAKYTESTHISTORYPOLICYID_MANUAL                  FlakyTestHistoryPolicyId = "ftm_policy.manual"
	FLAKYTESTHISTORYPOLICYID_FIXED                   FlakyTestHistoryPolicyId = "ftm_policy.fixed"
	FLAKYTESTHISTORYPOLICYID_DISABLE_FAILURE_RATE    FlakyTestHistoryPolicyId = "ftm_policy.disable.failure_rate"
	FLAKYTESTHISTORYPOLICYID_DISABLE_BRANCH_FLAKE    FlakyTestHistoryPolicyId = "ftm_policy.disable.branch_flake"
	FLAKYTESTHISTORYPOLICYID_DISABLE_DAYS_ACTIVE     FlakyTestHistoryPolicyId = "ftm_policy.disable.days_active"
	FLAKYTESTHISTORYPOLICYID_QUARANTINE_FAILURE_RATE FlakyTestHistoryPolicyId = "ftm_policy.quarantine.failure_rate"
	FLAKYTESTHISTORYPOLICYID_QUARANTINE_BRANCH_FLAKE FlakyTestHistoryPolicyId = "ftm_policy.quarantine.branch_flake"
	FLAKYTESTHISTORYPOLICYID_QUARANTINE_DAYS_ACTIVE  FlakyTestHistoryPolicyId = "ftm_policy.quarantine.days_active"
	FLAKYTESTHISTORYPOLICYID_UNKNOWN                 FlakyTestHistoryPolicyId = "unknown"
)

var allowedFlakyTestHistoryPolicyIdEnumValues = []FlakyTestHistoryPolicyId{
	FLAKYTESTHISTORYPOLICYID_MANUAL,
	FLAKYTESTHISTORYPOLICYID_FIXED,
	FLAKYTESTHISTORYPOLICYID_DISABLE_FAILURE_RATE,
	FLAKYTESTHISTORYPOLICYID_DISABLE_BRANCH_FLAKE,
	FLAKYTESTHISTORYPOLICYID_DISABLE_DAYS_ACTIVE,
	FLAKYTESTHISTORYPOLICYID_QUARANTINE_FAILURE_RATE,
	FLAKYTESTHISTORYPOLICYID_QUARANTINE_BRANCH_FLAKE,
	FLAKYTESTHISTORYPOLICYID_QUARANTINE_DAYS_ACTIVE,
	FLAKYTESTHISTORYPOLICYID_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FlakyTestHistoryPolicyId) GetAllowedValues() []FlakyTestHistoryPolicyId {
	return allowedFlakyTestHistoryPolicyIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FlakyTestHistoryPolicyId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FlakyTestHistoryPolicyId(value)
	return nil
}

// NewFlakyTestHistoryPolicyIdFromValue returns a pointer to a valid FlakyTestHistoryPolicyId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFlakyTestHistoryPolicyIdFromValue(v string) (*FlakyTestHistoryPolicyId, error) {
	ev := FlakyTestHistoryPolicyId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FlakyTestHistoryPolicyId: valid values are %v", v, allowedFlakyTestHistoryPolicyIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FlakyTestHistoryPolicyId) IsValid() bool {
	for _, existing := range allowedFlakyTestHistoryPolicyIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlakyTestHistoryPolicyId value.
func (v FlakyTestHistoryPolicyId) Ptr() *FlakyTestHistoryPolicyId {
	return &v
}
