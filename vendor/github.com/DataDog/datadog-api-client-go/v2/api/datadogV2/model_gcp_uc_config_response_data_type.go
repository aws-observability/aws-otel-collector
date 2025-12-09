// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GcpUcConfigResponseDataType Google Cloud Usage Cost config resource type.
type GcpUcConfigResponseDataType string

// List of GcpUcConfigResponseDataType.
const (
	GCPUCCONFIGRESPONSEDATATYPE_GCP_UC_CONFIG GcpUcConfigResponseDataType = "gcp_uc_config"
)

var allowedGcpUcConfigResponseDataTypeEnumValues = []GcpUcConfigResponseDataType{
	GCPUCCONFIGRESPONSEDATATYPE_GCP_UC_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GcpUcConfigResponseDataType) GetAllowedValues() []GcpUcConfigResponseDataType {
	return allowedGcpUcConfigResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GcpUcConfigResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GcpUcConfigResponseDataType(value)
	return nil
}

// NewGcpUcConfigResponseDataTypeFromValue returns a pointer to a valid GcpUcConfigResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGcpUcConfigResponseDataTypeFromValue(v string) (*GcpUcConfigResponseDataType, error) {
	ev := GcpUcConfigResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GcpUcConfigResponseDataType: valid values are %v", v, allowedGcpUcConfigResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GcpUcConfigResponseDataType) IsValid() bool {
	for _, existing := range allowedGcpUcConfigResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GcpUcConfigResponseDataType value.
func (v GcpUcConfigResponseDataType) Ptr() *GcpUcConfigResponseDataType {
	return &v
}
