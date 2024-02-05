// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppPipelineEventJobStatus The final status of the job.
type CIAppPipelineEventJobStatus string

// List of CIAppPipelineEventJobStatus.
const (
	CIAPPPIPELINEEVENTJOBSTATUS_SUCCESS  CIAppPipelineEventJobStatus = "success"
	CIAPPPIPELINEEVENTJOBSTATUS_ERROR    CIAppPipelineEventJobStatus = "error"
	CIAPPPIPELINEEVENTJOBSTATUS_CANCELED CIAppPipelineEventJobStatus = "canceled"
	CIAPPPIPELINEEVENTJOBSTATUS_SKIPPED  CIAppPipelineEventJobStatus = "skipped"
)

var allowedCIAppPipelineEventJobStatusEnumValues = []CIAppPipelineEventJobStatus{
	CIAPPPIPELINEEVENTJOBSTATUS_SUCCESS,
	CIAPPPIPELINEEVENTJOBSTATUS_ERROR,
	CIAPPPIPELINEEVENTJOBSTATUS_CANCELED,
	CIAPPPIPELINEEVENTJOBSTATUS_SKIPPED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppPipelineEventJobStatus) GetAllowedValues() []CIAppPipelineEventJobStatus {
	return allowedCIAppPipelineEventJobStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppPipelineEventJobStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppPipelineEventJobStatus(value)
	return nil
}

// NewCIAppPipelineEventJobStatusFromValue returns a pointer to a valid CIAppPipelineEventJobStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppPipelineEventJobStatusFromValue(v string) (*CIAppPipelineEventJobStatus, error) {
	ev := CIAppPipelineEventJobStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppPipelineEventJobStatus: valid values are %v", v, allowedCIAppPipelineEventJobStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppPipelineEventJobStatus) IsValid() bool {
	for _, existing := range allowedCIAppPipelineEventJobStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppPipelineEventJobStatus value.
func (v CIAppPipelineEventJobStatus) Ptr() *CIAppPipelineEventJobStatus {
	return &v
}
