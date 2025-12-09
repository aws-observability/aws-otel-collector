// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GreyNoiseIntegrationType The definition of the `GreyNoiseIntegrationType` object.
type GreyNoiseIntegrationType string

// List of GreyNoiseIntegrationType.
const (
	GREYNOISEINTEGRATIONTYPE_GREYNOISE GreyNoiseIntegrationType = "GreyNoise"
)

var allowedGreyNoiseIntegrationTypeEnumValues = []GreyNoiseIntegrationType{
	GREYNOISEINTEGRATIONTYPE_GREYNOISE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GreyNoiseIntegrationType) GetAllowedValues() []GreyNoiseIntegrationType {
	return allowedGreyNoiseIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GreyNoiseIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GreyNoiseIntegrationType(value)
	return nil
}

// NewGreyNoiseIntegrationTypeFromValue returns a pointer to a valid GreyNoiseIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGreyNoiseIntegrationTypeFromValue(v string) (*GreyNoiseIntegrationType, error) {
	ev := GreyNoiseIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GreyNoiseIntegrationType: valid values are %v", v, allowedGreyNoiseIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GreyNoiseIntegrationType) IsValid() bool {
	for _, existing := range allowedGreyNoiseIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GreyNoiseIntegrationType value.
func (v GreyNoiseIntegrationType) Ptr() *GreyNoiseIntegrationType {
	return &v
}
