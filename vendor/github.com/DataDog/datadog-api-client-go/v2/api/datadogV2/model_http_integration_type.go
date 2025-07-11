// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPIntegrationType The definition of `HTTPIntegrationType` object.
type HTTPIntegrationType string

// List of HTTPIntegrationType.
const (
	HTTPINTEGRATIONTYPE_HTTP HTTPIntegrationType = "HTTP"
)

var allowedHTTPIntegrationTypeEnumValues = []HTTPIntegrationType{
	HTTPINTEGRATIONTYPE_HTTP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HTTPIntegrationType) GetAllowedValues() []HTTPIntegrationType {
	return allowedHTTPIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HTTPIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HTTPIntegrationType(value)
	return nil
}

// NewHTTPIntegrationTypeFromValue returns a pointer to a valid HTTPIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHTTPIntegrationTypeFromValue(v string) (*HTTPIntegrationType, error) {
	ev := HTTPIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HTTPIntegrationType: valid values are %v", v, allowedHTTPIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HTTPIntegrationType) IsValid() bool {
	for _, existing := range allowedHTTPIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HTTPIntegrationType value.
func (v HTTPIntegrationType) Ptr() *HTTPIntegrationType {
	return &v
}
