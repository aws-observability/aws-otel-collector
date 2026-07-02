// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetUpdateType The type of resource for a dataset update request.
type SecurityMonitoringDatasetUpdateType string

// List of SecurityMonitoringDatasetUpdateType.
const (
	SECURITYMONITORINGDATASETUPDATETYPE_DATASET_UPDATE SecurityMonitoringDatasetUpdateType = "datasetUpdate"
)

var allowedSecurityMonitoringDatasetUpdateTypeEnumValues = []SecurityMonitoringDatasetUpdateType{
	SECURITYMONITORINGDATASETUPDATETYPE_DATASET_UPDATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringDatasetUpdateType) GetAllowedValues() []SecurityMonitoringDatasetUpdateType {
	return allowedSecurityMonitoringDatasetUpdateTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringDatasetUpdateType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringDatasetUpdateType(value)
	return nil
}

// NewSecurityMonitoringDatasetUpdateTypeFromValue returns a pointer to a valid SecurityMonitoringDatasetUpdateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringDatasetUpdateTypeFromValue(v string) (*SecurityMonitoringDatasetUpdateType, error) {
	ev := SecurityMonitoringDatasetUpdateType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringDatasetUpdateType: valid values are %v", v, allowedSecurityMonitoringDatasetUpdateTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringDatasetUpdateType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringDatasetUpdateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringDatasetUpdateType value.
func (v SecurityMonitoringDatasetUpdateType) Ptr() *SecurityMonitoringDatasetUpdateType {
	return &v
}
