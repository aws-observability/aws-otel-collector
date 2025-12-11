// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FastlyIntegrationType The definition of the `FastlyIntegrationType` object.
type FastlyIntegrationType string

// List of FastlyIntegrationType.
const (
	FASTLYINTEGRATIONTYPE_FASTLY FastlyIntegrationType = "Fastly"
)

var allowedFastlyIntegrationTypeEnumValues = []FastlyIntegrationType{
	FASTLYINTEGRATIONTYPE_FASTLY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FastlyIntegrationType) GetAllowedValues() []FastlyIntegrationType {
	return allowedFastlyIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FastlyIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FastlyIntegrationType(value)
	return nil
}

// NewFastlyIntegrationTypeFromValue returns a pointer to a valid FastlyIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFastlyIntegrationTypeFromValue(v string) (*FastlyIntegrationType, error) {
	ev := FastlyIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FastlyIntegrationType: valid values are %v", v, allowedFastlyIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FastlyIntegrationType) IsValid() bool {
	for _, existing := range allowedFastlyIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FastlyIntegrationType value.
func (v FastlyIntegrationType) Ptr() *FastlyIntegrationType {
	return &v
}
