// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HourlyUsageType Usage type that is being measured.
type HourlyUsageType string

// List of HourlyUsageType.
const (
	HOURLYUSAGETYPE_APP_SEC_HOST_COUNT                       HourlyUsageType = "app_sec_host_count"
	HOURLYUSAGETYPE_OBSERVABILITY_PIPELINES_BYTES_PROCESSSED HourlyUsageType = "observability_pipelines_bytes_processed"
	HOURLYUSAGETYPE_LAMBDA_TRACED_INVOCATIONS_COUNT          HourlyUsageType = "lambda_traced_invocations_count"
)

var allowedHourlyUsageTypeEnumValues = []HourlyUsageType{
	HOURLYUSAGETYPE_APP_SEC_HOST_COUNT,
	HOURLYUSAGETYPE_OBSERVABILITY_PIPELINES_BYTES_PROCESSSED,
	HOURLYUSAGETYPE_LAMBDA_TRACED_INVOCATIONS_COUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HourlyUsageType) GetAllowedValues() []HourlyUsageType {
	return allowedHourlyUsageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HourlyUsageType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HourlyUsageType(value)
	return nil
}

// NewHourlyUsageTypeFromValue returns a pointer to a valid HourlyUsageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHourlyUsageTypeFromValue(v string) (*HourlyUsageType, error) {
	ev := HourlyUsageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HourlyUsageType: valid values are %v", v, allowedHourlyUsageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HourlyUsageType) IsValid() bool {
	for _, existing := range allowedHourlyUsageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HourlyUsageType value.
func (v HourlyUsageType) Ptr() *HourlyUsageType {
	return &v
}
