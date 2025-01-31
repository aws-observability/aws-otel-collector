// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppPipelineEventPipelineInProgressStatus The in progress status of the pipeline.
type CIAppPipelineEventPipelineInProgressStatus string

// List of CIAppPipelineEventPipelineInProgressStatus.
const (
	CIAPPPIPELINEEVENTPIPELINEINPROGRESSSTATUS_RUNNING CIAppPipelineEventPipelineInProgressStatus = "running"
)

var allowedCIAppPipelineEventPipelineInProgressStatusEnumValues = []CIAppPipelineEventPipelineInProgressStatus{
	CIAPPPIPELINEEVENTPIPELINEINPROGRESSSTATUS_RUNNING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppPipelineEventPipelineInProgressStatus) GetAllowedValues() []CIAppPipelineEventPipelineInProgressStatus {
	return allowedCIAppPipelineEventPipelineInProgressStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppPipelineEventPipelineInProgressStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppPipelineEventPipelineInProgressStatus(value)
	return nil
}

// NewCIAppPipelineEventPipelineInProgressStatusFromValue returns a pointer to a valid CIAppPipelineEventPipelineInProgressStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppPipelineEventPipelineInProgressStatusFromValue(v string) (*CIAppPipelineEventPipelineInProgressStatus, error) {
	ev := CIAppPipelineEventPipelineInProgressStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppPipelineEventPipelineInProgressStatus: valid values are %v", v, allowedCIAppPipelineEventPipelineInProgressStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppPipelineEventPipelineInProgressStatus) IsValid() bool {
	for _, existing := range allowedCIAppPipelineEventPipelineInProgressStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppPipelineEventPipelineInProgressStatus value.
func (v CIAppPipelineEventPipelineInProgressStatus) Ptr() *CIAppPipelineEventPipelineInProgressStatus {
	return &v
}
