// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipInferenceStatus The lifecycle status of an ownership inference.
type OwnershipInferenceStatus string

// List of OwnershipInferenceStatus.
const (
	OWNERSHIPINFERENCESTATUS_SUGGESTED  OwnershipInferenceStatus = "suggested"
	OWNERSHIPINFERENCESTATUS_PERSISTED  OwnershipInferenceStatus = "persisted"
	OWNERSHIPINFERENCESTATUS_OVERRIDDEN OwnershipInferenceStatus = "overridden"
	OWNERSHIPINFERENCESTATUS_FAILED     OwnershipInferenceStatus = "failed"
	OWNERSHIPINFERENCESTATUS_UNKNOWN    OwnershipInferenceStatus = "unknown"
)

var allowedOwnershipInferenceStatusEnumValues = []OwnershipInferenceStatus{
	OWNERSHIPINFERENCESTATUS_SUGGESTED,
	OWNERSHIPINFERENCESTATUS_PERSISTED,
	OWNERSHIPINFERENCESTATUS_OVERRIDDEN,
	OWNERSHIPINFERENCESTATUS_FAILED,
	OWNERSHIPINFERENCESTATUS_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OwnershipInferenceStatus) GetAllowedValues() []OwnershipInferenceStatus {
	return allowedOwnershipInferenceStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OwnershipInferenceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OwnershipInferenceStatus(value)
	return nil
}

// NewOwnershipInferenceStatusFromValue returns a pointer to a valid OwnershipInferenceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOwnershipInferenceStatusFromValue(v string) (*OwnershipInferenceStatus, error) {
	ev := OwnershipInferenceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OwnershipInferenceStatus: valid values are %v", v, allowedOwnershipInferenceStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OwnershipInferenceStatus) IsValid() bool {
	for _, existing := range allowedOwnershipInferenceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OwnershipInferenceStatus value.
func (v OwnershipInferenceStatus) Ptr() *OwnershipInferenceStatus {
	return &v
}
