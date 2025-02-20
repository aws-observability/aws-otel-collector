// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricNotebookType Notebook resource type.
type MetricNotebookType string

// List of MetricNotebookType.
const (
	METRICNOTEBOOKTYPE_NOTEBOOKS MetricNotebookType = "notebooks"
)

var allowedMetricNotebookTypeEnumValues = []MetricNotebookType{
	METRICNOTEBOOKTYPE_NOTEBOOKS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricNotebookType) GetAllowedValues() []MetricNotebookType {
	return allowedMetricNotebookTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricNotebookType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricNotebookType(value)
	return nil
}

// NewMetricNotebookTypeFromValue returns a pointer to a valid MetricNotebookType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricNotebookTypeFromValue(v string) (*MetricNotebookType, error) {
	ev := MetricNotebookType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricNotebookType: valid values are %v", v, allowedMetricNotebookTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricNotebookType) IsValid() bool {
	for _, existing := range allowedMetricNotebookTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricNotebookType value.
func (v MetricNotebookType) Ptr() *MetricNotebookType {
	return &v
}
