// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunThreatHuntingJobRequestDataType Type of data.
type RunThreatHuntingJobRequestDataType string

// List of RunThreatHuntingJobRequestDataType.
const (
	RUNTHREATHUNTINGJOBREQUESTDATATYPE_HISTORICALDETECTIONSJOBCREATE RunThreatHuntingJobRequestDataType = "historicalDetectionsJobCreate"
)

var allowedRunThreatHuntingJobRequestDataTypeEnumValues = []RunThreatHuntingJobRequestDataType{
	RUNTHREATHUNTINGJOBREQUESTDATATYPE_HISTORICALDETECTIONSJOBCREATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RunThreatHuntingJobRequestDataType) GetAllowedValues() []RunThreatHuntingJobRequestDataType {
	return allowedRunThreatHuntingJobRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RunThreatHuntingJobRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RunThreatHuntingJobRequestDataType(value)
	return nil
}

// NewRunThreatHuntingJobRequestDataTypeFromValue returns a pointer to a valid RunThreatHuntingJobRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRunThreatHuntingJobRequestDataTypeFromValue(v string) (*RunThreatHuntingJobRequestDataType, error) {
	ev := RunThreatHuntingJobRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RunThreatHuntingJobRequestDataType: valid values are %v", v, allowedRunThreatHuntingJobRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RunThreatHuntingJobRequestDataType) IsValid() bool {
	for _, existing := range allowedRunThreatHuntingJobRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RunThreatHuntingJobRequestDataType value.
func (v RunThreatHuntingJobRequestDataType) Ptr() *RunThreatHuntingJobRequestDataType {
	return &v
}
