// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPIntegrationType The definition of the `GCPIntegrationType` object.
type GCPIntegrationType string

// List of GCPIntegrationType.
const (
	GCPINTEGRATIONTYPE_GCP GCPIntegrationType = "GCP"
)

var allowedGCPIntegrationTypeEnumValues = []GCPIntegrationType{
	GCPINTEGRATIONTYPE_GCP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GCPIntegrationType) GetAllowedValues() []GCPIntegrationType {
	return allowedGCPIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GCPIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GCPIntegrationType(value)
	return nil
}

// NewGCPIntegrationTypeFromValue returns a pointer to a valid GCPIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGCPIntegrationTypeFromValue(v string) (*GCPIntegrationType, error) {
	ev := GCPIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GCPIntegrationType: valid values are %v", v, allowedGCPIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GCPIntegrationType) IsValid() bool {
	for _, existing := range allowedGCPIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GCPIntegrationType value.
func (v GCPIntegrationType) Ptr() *GCPIntegrationType {
	return &v
}
