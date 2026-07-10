// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetDependentsType The type of resource for a dataset dependents entry.
type SecurityMonitoringDatasetDependentsType string

// List of SecurityMonitoringDatasetDependentsType.
const (
	SECURITYMONITORINGDATASETDEPENDENTSTYPE_DATASET_DEPENDENTS SecurityMonitoringDatasetDependentsType = "datasetDependents"
)

var allowedSecurityMonitoringDatasetDependentsTypeEnumValues = []SecurityMonitoringDatasetDependentsType{
	SECURITYMONITORINGDATASETDEPENDENTSTYPE_DATASET_DEPENDENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringDatasetDependentsType) GetAllowedValues() []SecurityMonitoringDatasetDependentsType {
	return allowedSecurityMonitoringDatasetDependentsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringDatasetDependentsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringDatasetDependentsType(value)
	return nil
}

// NewSecurityMonitoringDatasetDependentsTypeFromValue returns a pointer to a valid SecurityMonitoringDatasetDependentsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringDatasetDependentsTypeFromValue(v string) (*SecurityMonitoringDatasetDependentsType, error) {
	ev := SecurityMonitoringDatasetDependentsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringDatasetDependentsType: valid values are %v", v, allowedSecurityMonitoringDatasetDependentsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringDatasetDependentsType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringDatasetDependentsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringDatasetDependentsType value.
func (v SecurityMonitoringDatasetDependentsType) Ptr() *SecurityMonitoringDatasetDependentsType {
	return &v
}
