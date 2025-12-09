// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPMonitoredResourceConfigType The GCP monitored resource type. Only a subset of resource types are supported.
type GCPMonitoredResourceConfigType string

// List of GCPMonitoredResourceConfigType.
const (
	GCPMONITOREDRESOURCECONFIGTYPE_CLOUD_FUNCTION     GCPMonitoredResourceConfigType = "cloud_function"
	GCPMONITOREDRESOURCECONFIGTYPE_CLOUD_RUN_REVISION GCPMonitoredResourceConfigType = "cloud_run_revision"
	GCPMONITOREDRESOURCECONFIGTYPE_GCE_INSTANCE       GCPMonitoredResourceConfigType = "gce_instance"
)

var allowedGCPMonitoredResourceConfigTypeEnumValues = []GCPMonitoredResourceConfigType{
	GCPMONITOREDRESOURCECONFIGTYPE_CLOUD_FUNCTION,
	GCPMONITOREDRESOURCECONFIGTYPE_CLOUD_RUN_REVISION,
	GCPMONITOREDRESOURCECONFIGTYPE_GCE_INSTANCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GCPMonitoredResourceConfigType) GetAllowedValues() []GCPMonitoredResourceConfigType {
	return allowedGCPMonitoredResourceConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GCPMonitoredResourceConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GCPMonitoredResourceConfigType(value)
	return nil
}

// NewGCPMonitoredResourceConfigTypeFromValue returns a pointer to a valid GCPMonitoredResourceConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGCPMonitoredResourceConfigTypeFromValue(v string) (*GCPMonitoredResourceConfigType, error) {
	ev := GCPMonitoredResourceConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GCPMonitoredResourceConfigType: valid values are %v", v, allowedGCPMonitoredResourceConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GCPMonitoredResourceConfigType) IsValid() bool {
	for _, existing := range allowedGCPMonitoredResourceConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GCPMonitoredResourceConfigType value.
func (v GCPMonitoredResourceConfigType) Ptr() *GCPMonitoredResourceConfigType {
	return &v
}
