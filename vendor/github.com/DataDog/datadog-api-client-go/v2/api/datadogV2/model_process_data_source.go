// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProcessDataSource A data source for process-level infrastructure metrics.
type ProcessDataSource string

// List of ProcessDataSource.
const (
	PROCESSDATASOURCE_PROCESS ProcessDataSource = "process"
)

var allowedProcessDataSourceEnumValues = []ProcessDataSource{
	PROCESSDATASOURCE_PROCESS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProcessDataSource) GetAllowedValues() []ProcessDataSource {
	return allowedProcessDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProcessDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProcessDataSource(value)
	return nil
}

// NewProcessDataSourceFromValue returns a pointer to a valid ProcessDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProcessDataSourceFromValue(v string) (*ProcessDataSource, error) {
	ev := ProcessDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProcessDataSource: valid values are %v", v, allowedProcessDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProcessDataSource) IsValid() bool {
	for _, existing := range allowedProcessDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProcessDataSource value.
func (v ProcessDataSource) Ptr() *ProcessDataSource {
	return &v
}
