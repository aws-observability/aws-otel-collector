// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ThreatHuntingJobDataType Type of payload.
type ThreatHuntingJobDataType string

// List of ThreatHuntingJobDataType.
const (
	THREATHUNTINGJOBDATATYPE_HISTORICALDETECTIONSJOB ThreatHuntingJobDataType = "historicalDetectionsJob"
)

var allowedThreatHuntingJobDataTypeEnumValues = []ThreatHuntingJobDataType{
	THREATHUNTINGJOBDATATYPE_HISTORICALDETECTIONSJOB,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ThreatHuntingJobDataType) GetAllowedValues() []ThreatHuntingJobDataType {
	return allowedThreatHuntingJobDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ThreatHuntingJobDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ThreatHuntingJobDataType(value)
	return nil
}

// NewThreatHuntingJobDataTypeFromValue returns a pointer to a valid ThreatHuntingJobDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewThreatHuntingJobDataTypeFromValue(v string) (*ThreatHuntingJobDataType, error) {
	ev := ThreatHuntingJobDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ThreatHuntingJobDataType: valid values are %v", v, allowedThreatHuntingJobDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ThreatHuntingJobDataType) IsValid() bool {
	for _, existing := range allowedThreatHuntingJobDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ThreatHuntingJobDataType value.
func (v ThreatHuntingJobDataType) Ptr() *ThreatHuntingJobDataType {
	return &v
}
