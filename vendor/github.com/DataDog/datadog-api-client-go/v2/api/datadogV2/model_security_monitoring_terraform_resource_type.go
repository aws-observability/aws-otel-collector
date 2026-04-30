// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringTerraformResourceType The type of security monitoring resource to export to Terraform.
type SecurityMonitoringTerraformResourceType string

// List of SecurityMonitoringTerraformResourceType.
const (
	SECURITYMONITORINGTERRAFORMRESOURCETYPE_SUPPRESSIONS    SecurityMonitoringTerraformResourceType = "suppressions"
	SECURITYMONITORINGTERRAFORMRESOURCETYPE_CRITICAL_ASSETS SecurityMonitoringTerraformResourceType = "critical_assets"
)

var allowedSecurityMonitoringTerraformResourceTypeEnumValues = []SecurityMonitoringTerraformResourceType{
	SECURITYMONITORINGTERRAFORMRESOURCETYPE_SUPPRESSIONS,
	SECURITYMONITORINGTERRAFORMRESOURCETYPE_CRITICAL_ASSETS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringTerraformResourceType) GetAllowedValues() []SecurityMonitoringTerraformResourceType {
	return allowedSecurityMonitoringTerraformResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringTerraformResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringTerraformResourceType(value)
	return nil
}

// NewSecurityMonitoringTerraformResourceTypeFromValue returns a pointer to a valid SecurityMonitoringTerraformResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringTerraformResourceTypeFromValue(v string) (*SecurityMonitoringTerraformResourceType, error) {
	ev := SecurityMonitoringTerraformResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringTerraformResourceType: valid values are %v", v, allowedSecurityMonitoringTerraformResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringTerraformResourceType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringTerraformResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringTerraformResourceType value.
func (v SecurityMonitoringTerraformResourceType) Ptr() *SecurityMonitoringTerraformResourceType {
	return &v
}
