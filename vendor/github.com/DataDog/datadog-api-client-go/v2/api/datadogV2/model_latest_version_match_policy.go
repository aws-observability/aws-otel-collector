// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LatestVersionMatchPolicy The policy for matching the latest form version during an upsert operation.
type LatestVersionMatchPolicy string

// List of LatestVersionMatchPolicy.
const (
	LATESTVERSIONMATCHPOLICY_NONE          LatestVersionMatchPolicy = "none"
	LATESTVERSIONMATCHPOLICY_IF_ETAG_MATCH LatestVersionMatchPolicy = "if_etag_match"
)

var allowedLatestVersionMatchPolicyEnumValues = []LatestVersionMatchPolicy{
	LATESTVERSIONMATCHPOLICY_NONE,
	LATESTVERSIONMATCHPOLICY_IF_ETAG_MATCH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LatestVersionMatchPolicy) GetAllowedValues() []LatestVersionMatchPolicy {
	return allowedLatestVersionMatchPolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LatestVersionMatchPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LatestVersionMatchPolicy(value)
	return nil
}

// NewLatestVersionMatchPolicyFromValue returns a pointer to a valid LatestVersionMatchPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLatestVersionMatchPolicyFromValue(v string) (*LatestVersionMatchPolicy, error) {
	ev := LatestVersionMatchPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LatestVersionMatchPolicy: valid values are %v", v, allowedLatestVersionMatchPolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LatestVersionMatchPolicy) IsValid() bool {
	for _, existing := range allowedLatestVersionMatchPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LatestVersionMatchPolicy value.
func (v LatestVersionMatchPolicy) Ptr() *LatestVersionMatchPolicy {
	return &v
}
