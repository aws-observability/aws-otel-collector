// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HistoricalJobDataType Type of payload.
type HistoricalJobDataType string

// List of HistoricalJobDataType.
const (
	HISTORICALJOBDATATYPE_HISTORICALDETECTIONSJOB HistoricalJobDataType = "historicalDetectionsJob"
)

var allowedHistoricalJobDataTypeEnumValues = []HistoricalJobDataType{
	HISTORICALJOBDATATYPE_HISTORICALDETECTIONSJOB,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HistoricalJobDataType) GetAllowedValues() []HistoricalJobDataType {
	return allowedHistoricalJobDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HistoricalJobDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HistoricalJobDataType(value)
	return nil
}

// NewHistoricalJobDataTypeFromValue returns a pointer to a valid HistoricalJobDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHistoricalJobDataTypeFromValue(v string) (*HistoricalJobDataType, error) {
	ev := HistoricalJobDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HistoricalJobDataType: valid values are %v", v, allowedHistoricalJobDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HistoricalJobDataType) IsValid() bool {
	for _, existing := range allowedHistoricalJobDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HistoricalJobDataType value.
func (v HistoricalJobDataType) Ptr() *HistoricalJobDataType {
	return &v
}
