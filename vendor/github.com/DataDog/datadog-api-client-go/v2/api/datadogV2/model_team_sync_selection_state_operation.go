// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamSyncSelectionStateOperation The operation to perform on the selected hierarchy.
// When set to `include`, synchronization covers the
// referenced teams or organizations.
type TeamSyncSelectionStateOperation string

// List of TeamSyncSelectionStateOperation.
const (
	TEAMSYNCSELECTIONSTATEOPERATION_INCLUDE TeamSyncSelectionStateOperation = "include"
)

var allowedTeamSyncSelectionStateOperationEnumValues = []TeamSyncSelectionStateOperation{
	TEAMSYNCSELECTIONSTATEOPERATION_INCLUDE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamSyncSelectionStateOperation) GetAllowedValues() []TeamSyncSelectionStateOperation {
	return allowedTeamSyncSelectionStateOperationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamSyncSelectionStateOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamSyncSelectionStateOperation(value)
	return nil
}

// NewTeamSyncSelectionStateOperationFromValue returns a pointer to a valid TeamSyncSelectionStateOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamSyncSelectionStateOperationFromValue(v string) (*TeamSyncSelectionStateOperation, error) {
	ev := TeamSyncSelectionStateOperation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamSyncSelectionStateOperation: valid values are %v", v, allowedTeamSyncSelectionStateOperationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamSyncSelectionStateOperation) IsValid() bool {
	for _, existing := range allowedTeamSyncSelectionStateOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamSyncSelectionStateOperation value.
func (v TeamSyncSelectionStateOperation) Ptr() *TeamSyncSelectionStateOperation {
	return &v
}
