// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisResponseDataType Analysis response resource type.
type AnalysisResponseDataType string

// List of AnalysisResponseDataType.
const (
	ANALYSISRESPONSEDATATYPE_SERVER_REQUEST AnalysisResponseDataType = "server_request"
)

var allowedAnalysisResponseDataTypeEnumValues = []AnalysisResponseDataType{
	ANALYSISRESPONSEDATATYPE_SERVER_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AnalysisResponseDataType) GetAllowedValues() []AnalysisResponseDataType {
	return allowedAnalysisResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AnalysisResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AnalysisResponseDataType(value)
	return nil
}

// NewAnalysisResponseDataTypeFromValue returns a pointer to a valid AnalysisResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAnalysisResponseDataTypeFromValue(v string) (*AnalysisResponseDataType, error) {
	ev := AnalysisResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AnalysisResponseDataType: valid values are %v", v, allowedAnalysisResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AnalysisResponseDataType) IsValid() bool {
	for _, existing := range allowedAnalysisResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnalysisResponseDataType value.
func (v AnalysisResponseDataType) Ptr() *AnalysisResponseDataType {
	return &v
}
