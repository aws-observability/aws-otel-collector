// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnapshotUpdateRequestDataType Snapshots resource type.
type SnapshotUpdateRequestDataType string

// List of SnapshotUpdateRequestDataType.
const (
	SNAPSHOTUPDATEREQUESTDATATYPE_SNAPSHOTS SnapshotUpdateRequestDataType = "snapshots"
)

var allowedSnapshotUpdateRequestDataTypeEnumValues = []SnapshotUpdateRequestDataType{
	SNAPSHOTUPDATEREQUESTDATATYPE_SNAPSHOTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SnapshotUpdateRequestDataType) GetAllowedValues() []SnapshotUpdateRequestDataType {
	return allowedSnapshotUpdateRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SnapshotUpdateRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SnapshotUpdateRequestDataType(value)
	return nil
}

// NewSnapshotUpdateRequestDataTypeFromValue returns a pointer to a valid SnapshotUpdateRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSnapshotUpdateRequestDataTypeFromValue(v string) (*SnapshotUpdateRequestDataType, error) {
	ev := SnapshotUpdateRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SnapshotUpdateRequestDataType: valid values are %v", v, allowedSnapshotUpdateRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SnapshotUpdateRequestDataType) IsValid() bool {
	for _, existing := range allowedSnapshotUpdateRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SnapshotUpdateRequestDataType value.
func (v SnapshotUpdateRequestDataType) Ptr() *SnapshotUpdateRequestDataType {
	return &v
}
