// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatastoreItemConflictMode How to handle conflicts when inserting items that already exist in the datastore.
type DatastoreItemConflictMode string

// List of DatastoreItemConflictMode.
const (
	DATASTOREITEMCONFLICTMODE_FAIL_ON_CONFLICT      DatastoreItemConflictMode = "fail_on_conflict"
	DATASTOREITEMCONFLICTMODE_OVERWRITE_ON_CONFLICT DatastoreItemConflictMode = "overwrite_on_conflict"
)

var allowedDatastoreItemConflictModeEnumValues = []DatastoreItemConflictMode{
	DATASTOREITEMCONFLICTMODE_FAIL_ON_CONFLICT,
	DATASTOREITEMCONFLICTMODE_OVERWRITE_ON_CONFLICT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatastoreItemConflictMode) GetAllowedValues() []DatastoreItemConflictMode {
	return allowedDatastoreItemConflictModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatastoreItemConflictMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatastoreItemConflictMode(value)
	return nil
}

// NewDatastoreItemConflictModeFromValue returns a pointer to a valid DatastoreItemConflictMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatastoreItemConflictModeFromValue(v string) (*DatastoreItemConflictMode, error) {
	ev := DatastoreItemConflictMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatastoreItemConflictMode: valid values are %v", v, allowedDatastoreItemConflictModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatastoreItemConflictMode) IsValid() bool {
	for _, existing := range allowedDatastoreItemConflictModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatastoreItemConflictMode value.
func (v DatastoreItemConflictMode) Ptr() *DatastoreItemConflictMode {
	return &v
}
