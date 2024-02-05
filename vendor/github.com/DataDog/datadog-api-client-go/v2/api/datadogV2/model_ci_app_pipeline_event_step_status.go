// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppPipelineEventStepStatus The final status of the step.
type CIAppPipelineEventStepStatus string

// List of CIAppPipelineEventStepStatus.
const (
	CIAPPPIPELINEEVENTSTEPSTATUS_SUCCESS CIAppPipelineEventStepStatus = "success"
	CIAPPPIPELINEEVENTSTEPSTATUS_ERROR   CIAppPipelineEventStepStatus = "error"
)

var allowedCIAppPipelineEventStepStatusEnumValues = []CIAppPipelineEventStepStatus{
	CIAPPPIPELINEEVENTSTEPSTATUS_SUCCESS,
	CIAPPPIPELINEEVENTSTEPSTATUS_ERROR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppPipelineEventStepStatus) GetAllowedValues() []CIAppPipelineEventStepStatus {
	return allowedCIAppPipelineEventStepStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppPipelineEventStepStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppPipelineEventStepStatus(value)
	return nil
}

// NewCIAppPipelineEventStepStatusFromValue returns a pointer to a valid CIAppPipelineEventStepStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppPipelineEventStepStatusFromValue(v string) (*CIAppPipelineEventStepStatus, error) {
	ev := CIAppPipelineEventStepStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppPipelineEventStepStatus: valid values are %v", v, allowedCIAppPipelineEventStepStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppPipelineEventStepStatus) IsValid() bool {
	for _, existing := range allowedCIAppPipelineEventStepStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppPipelineEventStepStatus value.
func (v CIAppPipelineEventStepStatus) Ptr() *CIAppPipelineEventStepStatus {
	return &v
}
