// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VersionHistoryUpdateType The type of change.
type VersionHistoryUpdateType string

// List of VersionHistoryUpdateType.
const (
	VERSIONHISTORYUPDATETYPE_CREATE VersionHistoryUpdateType = "create"
	VERSIONHISTORYUPDATETYPE_UPDATE VersionHistoryUpdateType = "update"
	VERSIONHISTORYUPDATETYPE_DELETE VersionHistoryUpdateType = "delete"
)

var allowedVersionHistoryUpdateTypeEnumValues = []VersionHistoryUpdateType{
	VERSIONHISTORYUPDATETYPE_CREATE,
	VERSIONHISTORYUPDATETYPE_UPDATE,
	VERSIONHISTORYUPDATETYPE_DELETE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *VersionHistoryUpdateType) GetAllowedValues() []VersionHistoryUpdateType {
	return allowedVersionHistoryUpdateTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *VersionHistoryUpdateType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = VersionHistoryUpdateType(value)
	return nil
}

// NewVersionHistoryUpdateTypeFromValue returns a pointer to a valid VersionHistoryUpdateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewVersionHistoryUpdateTypeFromValue(v string) (*VersionHistoryUpdateType, error) {
	ev := VersionHistoryUpdateType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for VersionHistoryUpdateType: valid values are %v", v, allowedVersionHistoryUpdateTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v VersionHistoryUpdateType) IsValid() bool {
	for _, existing := range allowedVersionHistoryUpdateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VersionHistoryUpdateType value.
func (v VersionHistoryUpdateType) Ptr() *VersionHistoryUpdateType {
	return &v
}
