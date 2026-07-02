// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetType The type of resource for a dataset response.
type SecurityMonitoringDatasetType string

// List of SecurityMonitoringDatasetType.
const (
	SECURITYMONITORINGDATASETTYPE_DATASET SecurityMonitoringDatasetType = "dataset"
)

var allowedSecurityMonitoringDatasetTypeEnumValues = []SecurityMonitoringDatasetType{
	SECURITYMONITORINGDATASETTYPE_DATASET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringDatasetType) GetAllowedValues() []SecurityMonitoringDatasetType {
	return allowedSecurityMonitoringDatasetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringDatasetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringDatasetType(value)
	return nil
}

// NewSecurityMonitoringDatasetTypeFromValue returns a pointer to a valid SecurityMonitoringDatasetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringDatasetTypeFromValue(v string) (*SecurityMonitoringDatasetType, error) {
	ev := SecurityMonitoringDatasetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringDatasetType: valid values are %v", v, allowedSecurityMonitoringDatasetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringDatasetType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringDatasetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringDatasetType value.
func (v SecurityMonitoringDatasetType) Ptr() *SecurityMonitoringDatasetType {
	return &v
}
