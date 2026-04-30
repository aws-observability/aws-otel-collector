// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BatchRowsQueryDataType Resource type identifier for batch queries of reference table rows.
type BatchRowsQueryDataType string

// List of BatchRowsQueryDataType.
const (
	BATCHROWSQUERYDATATYPE_REFERENCE_TABLES_BATCH_ROWS_QUERY BatchRowsQueryDataType = "reference-tables-batch-rows-query"
)

var allowedBatchRowsQueryDataTypeEnumValues = []BatchRowsQueryDataType{
	BATCHROWSQUERYDATATYPE_REFERENCE_TABLES_BATCH_ROWS_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BatchRowsQueryDataType) GetAllowedValues() []BatchRowsQueryDataType {
	return allowedBatchRowsQueryDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BatchRowsQueryDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BatchRowsQueryDataType(value)
	return nil
}

// NewBatchRowsQueryDataTypeFromValue returns a pointer to a valid BatchRowsQueryDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBatchRowsQueryDataTypeFromValue(v string) (*BatchRowsQueryDataType, error) {
	ev := BatchRowsQueryDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BatchRowsQueryDataType: valid values are %v", v, allowedBatchRowsQueryDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BatchRowsQueryDataType) IsValid() bool {
	for _, existing := range allowedBatchRowsQueryDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BatchRowsQueryDataType value.
func (v BatchRowsQueryDataType) Ptr() *BatchRowsQueryDataType {
	return &v
}
