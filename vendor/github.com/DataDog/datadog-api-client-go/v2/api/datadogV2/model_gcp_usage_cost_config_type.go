// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPUsageCostConfigType Type of Google Cloud Usage Cost config.
type GCPUsageCostConfigType string

// List of GCPUsageCostConfigType.
const (
	GCPUSAGECOSTCONFIGTYPE_GCP_UC_CONFIG GCPUsageCostConfigType = "gcp_uc_config"
)

var allowedGCPUsageCostConfigTypeEnumValues = []GCPUsageCostConfigType{
	GCPUSAGECOSTCONFIGTYPE_GCP_UC_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GCPUsageCostConfigType) GetAllowedValues() []GCPUsageCostConfigType {
	return allowedGCPUsageCostConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GCPUsageCostConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GCPUsageCostConfigType(value)
	return nil
}

// NewGCPUsageCostConfigTypeFromValue returns a pointer to a valid GCPUsageCostConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGCPUsageCostConfigTypeFromValue(v string) (*GCPUsageCostConfigType, error) {
	ev := GCPUsageCostConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GCPUsageCostConfigType: valid values are %v", v, allowedGCPUsageCostConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GCPUsageCostConfigType) IsValid() bool {
	for _, existing := range allowedGCPUsageCostConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GCPUsageCostConfigType value.
func (v GCPUsageCostConfigType) Ptr() *GCPUsageCostConfigType {
	return &v
}
