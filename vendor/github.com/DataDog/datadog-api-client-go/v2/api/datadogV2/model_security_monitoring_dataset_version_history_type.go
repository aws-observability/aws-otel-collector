// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetVersionHistoryType The type of resource for a dataset version history response.
type SecurityMonitoringDatasetVersionHistoryType string

// List of SecurityMonitoringDatasetVersionHistoryType.
const (
	SECURITYMONITORINGDATASETVERSIONHISTORYTYPE_DATASET_VERSION_HISTORY SecurityMonitoringDatasetVersionHistoryType = "dataset_version_history"
)

var allowedSecurityMonitoringDatasetVersionHistoryTypeEnumValues = []SecurityMonitoringDatasetVersionHistoryType{
	SECURITYMONITORINGDATASETVERSIONHISTORYTYPE_DATASET_VERSION_HISTORY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringDatasetVersionHistoryType) GetAllowedValues() []SecurityMonitoringDatasetVersionHistoryType {
	return allowedSecurityMonitoringDatasetVersionHistoryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringDatasetVersionHistoryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringDatasetVersionHistoryType(value)
	return nil
}

// NewSecurityMonitoringDatasetVersionHistoryTypeFromValue returns a pointer to a valid SecurityMonitoringDatasetVersionHistoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringDatasetVersionHistoryTypeFromValue(v string) (*SecurityMonitoringDatasetVersionHistoryType, error) {
	ev := SecurityMonitoringDatasetVersionHistoryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringDatasetVersionHistoryType: valid values are %v", v, allowedSecurityMonitoringDatasetVersionHistoryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringDatasetVersionHistoryType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringDatasetVersionHistoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringDatasetVersionHistoryType value.
func (v SecurityMonitoringDatasetVersionHistoryType) Ptr() *SecurityMonitoringDatasetVersionHistoryType {
	return &v
}
