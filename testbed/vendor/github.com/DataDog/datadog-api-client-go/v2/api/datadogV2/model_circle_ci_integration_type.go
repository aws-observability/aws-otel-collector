// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CircleCIIntegrationType The definition of the `CircleCIIntegrationType` object.
type CircleCIIntegrationType string

// List of CircleCIIntegrationType.
const (
	CIRCLECIINTEGRATIONTYPE_CIRCLECI CircleCIIntegrationType = "CircleCI"
)

var allowedCircleCIIntegrationTypeEnumValues = []CircleCIIntegrationType{
	CIRCLECIINTEGRATIONTYPE_CIRCLECI,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CircleCIIntegrationType) GetAllowedValues() []CircleCIIntegrationType {
	return allowedCircleCIIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CircleCIIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CircleCIIntegrationType(value)
	return nil
}

// NewCircleCIIntegrationTypeFromValue returns a pointer to a valid CircleCIIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCircleCIIntegrationTypeFromValue(v string) (*CircleCIIntegrationType, error) {
	ev := CircleCIIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CircleCIIntegrationType: valid values are %v", v, allowedCircleCIIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CircleCIIntegrationType) IsValid() bool {
	for _, existing := range allowedCircleCIIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CircleCIIntegrationType value.
func (v CircleCIIntegrationType) Ptr() *CircleCIIntegrationType {
	return &v
}
