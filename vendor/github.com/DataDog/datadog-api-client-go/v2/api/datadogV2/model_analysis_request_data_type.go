// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisRequestDataType Analysis request resource type.
type AnalysisRequestDataType string

// List of AnalysisRequestDataType.
const (
	ANALYSISREQUESTDATATYPE_ANALYSIS_REQUEST AnalysisRequestDataType = "analysis_request"
)

var allowedAnalysisRequestDataTypeEnumValues = []AnalysisRequestDataType{
	ANALYSISREQUESTDATATYPE_ANALYSIS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AnalysisRequestDataType) GetAllowedValues() []AnalysisRequestDataType {
	return allowedAnalysisRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AnalysisRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AnalysisRequestDataType(value)
	return nil
}

// NewAnalysisRequestDataTypeFromValue returns a pointer to a valid AnalysisRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAnalysisRequestDataTypeFromValue(v string) (*AnalysisRequestDataType, error) {
	ev := AnalysisRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AnalysisRequestDataType: valid values are %v", v, allowedAnalysisRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AnalysisRequestDataType) IsValid() bool {
	for _, existing := range allowedAnalysisRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnalysisRequestDataType value.
func (v AnalysisRequestDataType) Ptr() *AnalysisRequestDataType {
	return &v
}
