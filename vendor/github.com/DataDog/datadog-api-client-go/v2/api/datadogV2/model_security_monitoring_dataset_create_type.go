// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetCreateType The type of resource for a dataset create request.
type SecurityMonitoringDatasetCreateType string

// List of SecurityMonitoringDatasetCreateType.
const (
	SECURITYMONITORINGDATASETCREATETYPE_DATASET_CREATE SecurityMonitoringDatasetCreateType = "datasetCreate"
)

var allowedSecurityMonitoringDatasetCreateTypeEnumValues = []SecurityMonitoringDatasetCreateType{
	SECURITYMONITORINGDATASETCREATETYPE_DATASET_CREATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringDatasetCreateType) GetAllowedValues() []SecurityMonitoringDatasetCreateType {
	return allowedSecurityMonitoringDatasetCreateTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringDatasetCreateType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringDatasetCreateType(value)
	return nil
}

// NewSecurityMonitoringDatasetCreateTypeFromValue returns a pointer to a valid SecurityMonitoringDatasetCreateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringDatasetCreateTypeFromValue(v string) (*SecurityMonitoringDatasetCreateType, error) {
	ev := SecurityMonitoringDatasetCreateType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringDatasetCreateType: valid values are %v", v, allowedSecurityMonitoringDatasetCreateTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringDatasetCreateType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringDatasetCreateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringDatasetCreateType value.
func (v SecurityMonitoringDatasetCreateType) Ptr() *SecurityMonitoringDatasetCreateType {
	return &v
}
