// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunHistoricalJobRequestDataType Type of data.
type RunHistoricalJobRequestDataType string

// List of RunHistoricalJobRequestDataType.
const (
	RUNHISTORICALJOBREQUESTDATATYPE_HISTORICALDETECTIONSJOBCREATE RunHistoricalJobRequestDataType = "historicalDetectionsJobCreate"
)

var allowedRunHistoricalJobRequestDataTypeEnumValues = []RunHistoricalJobRequestDataType{
	RUNHISTORICALJOBREQUESTDATATYPE_HISTORICALDETECTIONSJOBCREATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RunHistoricalJobRequestDataType) GetAllowedValues() []RunHistoricalJobRequestDataType {
	return allowedRunHistoricalJobRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RunHistoricalJobRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RunHistoricalJobRequestDataType(value)
	return nil
}

// NewRunHistoricalJobRequestDataTypeFromValue returns a pointer to a valid RunHistoricalJobRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRunHistoricalJobRequestDataTypeFromValue(v string) (*RunHistoricalJobRequestDataType, error) {
	ev := RunHistoricalJobRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RunHistoricalJobRequestDataType: valid values are %v", v, allowedRunHistoricalJobRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RunHistoricalJobRequestDataType) IsValid() bool {
	for _, existing := range allowedRunHistoricalJobRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RunHistoricalJobRequestDataType value.
func (v RunHistoricalJobRequestDataType) Ptr() *RunHistoricalJobRequestDataType {
	return &v
}
