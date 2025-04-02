// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConvertJobResultsToSignalsDataType Type of payload.
type ConvertJobResultsToSignalsDataType string

// List of ConvertJobResultsToSignalsDataType.
const (
	CONVERTJOBRESULTSTOSIGNALSDATATYPE_HISTORICALDETECTIONSJOBRESULTSIGNALCONVERSION ConvertJobResultsToSignalsDataType = "historicalDetectionsJobResultSignalConversion"
)

var allowedConvertJobResultsToSignalsDataTypeEnumValues = []ConvertJobResultsToSignalsDataType{
	CONVERTJOBRESULTSTOSIGNALSDATATYPE_HISTORICALDETECTIONSJOBRESULTSIGNALCONVERSION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ConvertJobResultsToSignalsDataType) GetAllowedValues() []ConvertJobResultsToSignalsDataType {
	return allowedConvertJobResultsToSignalsDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ConvertJobResultsToSignalsDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ConvertJobResultsToSignalsDataType(value)
	return nil
}

// NewConvertJobResultsToSignalsDataTypeFromValue returns a pointer to a valid ConvertJobResultsToSignalsDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewConvertJobResultsToSignalsDataTypeFromValue(v string) (*ConvertJobResultsToSignalsDataType, error) {
	ev := ConvertJobResultsToSignalsDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ConvertJobResultsToSignalsDataType: valid values are %v", v, allowedConvertJobResultsToSignalsDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ConvertJobResultsToSignalsDataType) IsValid() bool {
	for _, existing := range allowedConvertJobResultsToSignalsDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConvertJobResultsToSignalsDataType value.
func (v ConvertJobResultsToSignalsDataType) Ptr() *ConvertJobResultsToSignalsDataType {
	return &v
}
