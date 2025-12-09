// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OktaIntegrationType The definition of the `OktaIntegrationType` object.
type OktaIntegrationType string

// List of OktaIntegrationType.
const (
	OKTAINTEGRATIONTYPE_OKTA OktaIntegrationType = "Okta"
)

var allowedOktaIntegrationTypeEnumValues = []OktaIntegrationType{
	OKTAINTEGRATIONTYPE_OKTA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OktaIntegrationType) GetAllowedValues() []OktaIntegrationType {
	return allowedOktaIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OktaIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OktaIntegrationType(value)
	return nil
}

// NewOktaIntegrationTypeFromValue returns a pointer to a valid OktaIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOktaIntegrationTypeFromValue(v string) (*OktaIntegrationType, error) {
	ev := OktaIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OktaIntegrationType: valid values are %v", v, allowedOktaIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OktaIntegrationType) IsValid() bool {
	for _, existing := range allowedOktaIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OktaIntegrationType value.
func (v OktaIntegrationType) Ptr() *OktaIntegrationType {
	return &v
}
