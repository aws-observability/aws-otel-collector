// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumPermanentRetentionFilterID The identifier of a permanent RUM retention filter.
type RumPermanentRetentionFilterID string

// List of RumPermanentRetentionFilterID.
const (
	RUMPERMANENTRETENTIONFILTERID_RUM_APM_FLAT_SAMPLING  RumPermanentRetentionFilterID = "rum_apm_flat_sampling"
	RUMPERMANENTRETENTIONFILTERID_SYNTHETICS_SESSIONS    RumPermanentRetentionFilterID = "synthetics_sessions"
	RUMPERMANENTRETENTIONFILTERID_FORCED_REPLAY_SESSIONS RumPermanentRetentionFilterID = "forced_replay_sessions"
)

var allowedRumPermanentRetentionFilterIDEnumValues = []RumPermanentRetentionFilterID{
	RUMPERMANENTRETENTIONFILTERID_RUM_APM_FLAT_SAMPLING,
	RUMPERMANENTRETENTIONFILTERID_SYNTHETICS_SESSIONS,
	RUMPERMANENTRETENTIONFILTERID_FORCED_REPLAY_SESSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumPermanentRetentionFilterID) GetAllowedValues() []RumPermanentRetentionFilterID {
	return allowedRumPermanentRetentionFilterIDEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumPermanentRetentionFilterID) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumPermanentRetentionFilterID(value)
	return nil
}

// NewRumPermanentRetentionFilterIDFromValue returns a pointer to a valid RumPermanentRetentionFilterID
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumPermanentRetentionFilterIDFromValue(v string) (*RumPermanentRetentionFilterID, error) {
	ev := RumPermanentRetentionFilterID(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumPermanentRetentionFilterID: valid values are %v", v, allowedRumPermanentRetentionFilterIDEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumPermanentRetentionFilterID) IsValid() bool {
	for _, existing := range allowedRumPermanentRetentionFilterIDEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumPermanentRetentionFilterID value.
func (v RumPermanentRetentionFilterID) Ptr() *RumPermanentRetentionFilterID {
	return &v
}
