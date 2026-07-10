// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabRunStatus The status of a Model Lab run.
type ModelLabRunStatus string

// List of ModelLabRunStatus.
const (
	MODELLABRUNSTATUS_PENDING      ModelLabRunStatus = "pending"
	MODELLABRUNSTATUS_RUNNING      ModelLabRunStatus = "running"
	MODELLABRUNSTATUS_COMPLETED    ModelLabRunStatus = "completed"
	MODELLABRUNSTATUS_FAILED       ModelLabRunStatus = "failed"
	MODELLABRUNSTATUS_KILLED       ModelLabRunStatus = "killed"
	MODELLABRUNSTATUS_UNRESPONSIVE ModelLabRunStatus = "unresponsive"
	MODELLABRUNSTATUS_PAUSED       ModelLabRunStatus = "paused"
)

var allowedModelLabRunStatusEnumValues = []ModelLabRunStatus{
	MODELLABRUNSTATUS_PENDING,
	MODELLABRUNSTATUS_RUNNING,
	MODELLABRUNSTATUS_COMPLETED,
	MODELLABRUNSTATUS_FAILED,
	MODELLABRUNSTATUS_KILLED,
	MODELLABRUNSTATUS_UNRESPONSIVE,
	MODELLABRUNSTATUS_PAUSED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ModelLabRunStatus) GetAllowedValues() []ModelLabRunStatus {
	return allowedModelLabRunStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ModelLabRunStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ModelLabRunStatus(value)
	return nil
}

// NewModelLabRunStatusFromValue returns a pointer to a valid ModelLabRunStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewModelLabRunStatusFromValue(v string) (*ModelLabRunStatus, error) {
	ev := ModelLabRunStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ModelLabRunStatus: valid values are %v", v, allowedModelLabRunStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ModelLabRunStatus) IsValid() bool {
	for _, existing := range allowedModelLabRunStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModelLabRunStatus value.
func (v ModelLabRunStatus) Ptr() *ModelLabRunStatus {
	return &v
}
