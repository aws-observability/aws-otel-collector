// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VirusTotalIntegrationType The definition of the `VirusTotalIntegrationType` object.
type VirusTotalIntegrationType string

// List of VirusTotalIntegrationType.
const (
	VIRUSTOTALINTEGRATIONTYPE_VIRUSTOTAL VirusTotalIntegrationType = "VirusTotal"
)

var allowedVirusTotalIntegrationTypeEnumValues = []VirusTotalIntegrationType{
	VIRUSTOTALINTEGRATIONTYPE_VIRUSTOTAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *VirusTotalIntegrationType) GetAllowedValues() []VirusTotalIntegrationType {
	return allowedVirusTotalIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *VirusTotalIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = VirusTotalIntegrationType(value)
	return nil
}

// NewVirusTotalIntegrationTypeFromValue returns a pointer to a valid VirusTotalIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewVirusTotalIntegrationTypeFromValue(v string) (*VirusTotalIntegrationType, error) {
	ev := VirusTotalIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for VirusTotalIntegrationType: valid values are %v", v, allowedVirusTotalIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v VirusTotalIntegrationType) IsValid() bool {
	for _, existing := range allowedVirusTotalIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VirusTotalIntegrationType value.
func (v VirusTotalIntegrationType) Ptr() *VirusTotalIntegrationType {
	return &v
}
