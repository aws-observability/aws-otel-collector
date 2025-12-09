// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetType Resource type, always set to `dataset`.
type DatasetType string

// List of DatasetType.
const (
	DATASETTYPE_DATASET DatasetType = "dataset"
)

var allowedDatasetTypeEnumValues = []DatasetType{
	DATASETTYPE_DATASET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatasetType) GetAllowedValues() []DatasetType {
	return allowedDatasetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatasetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatasetType(value)
	return nil
}

// NewDatasetTypeFromValue returns a pointer to a valid DatasetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatasetTypeFromValue(v string) (*DatasetType, error) {
	ev := DatasetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatasetType: valid values are %v", v, allowedDatasetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatasetType) IsValid() bool {
	for _, existing := range allowedDatasetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatasetType value.
func (v DatasetType) Ptr() *DatasetType {
	return &v
}
