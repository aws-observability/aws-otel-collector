// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GcpScanOptionsInputUpdateDataType GCP scan options resource type.
type GcpScanOptionsInputUpdateDataType string

// List of GcpScanOptionsInputUpdateDataType.
const (
	GCPSCANOPTIONSINPUTUPDATEDATATYPE_GCP_SCAN_OPTIONS GcpScanOptionsInputUpdateDataType = "gcp_scan_options"
)

var allowedGcpScanOptionsInputUpdateDataTypeEnumValues = []GcpScanOptionsInputUpdateDataType{
	GCPSCANOPTIONSINPUTUPDATEDATATYPE_GCP_SCAN_OPTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GcpScanOptionsInputUpdateDataType) GetAllowedValues() []GcpScanOptionsInputUpdateDataType {
	return allowedGcpScanOptionsInputUpdateDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GcpScanOptionsInputUpdateDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GcpScanOptionsInputUpdateDataType(value)
	return nil
}

// NewGcpScanOptionsInputUpdateDataTypeFromValue returns a pointer to a valid GcpScanOptionsInputUpdateDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGcpScanOptionsInputUpdateDataTypeFromValue(v string) (*GcpScanOptionsInputUpdateDataType, error) {
	ev := GcpScanOptionsInputUpdateDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GcpScanOptionsInputUpdateDataType: valid values are %v", v, allowedGcpScanOptionsInputUpdateDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GcpScanOptionsInputUpdateDataType) IsValid() bool {
	for _, existing := range allowedGcpScanOptionsInputUpdateDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GcpScanOptionsInputUpdateDataType value.
func (v GcpScanOptionsInputUpdateDataType) Ptr() *GcpScanOptionsInputUpdateDataType {
	return &v
}
