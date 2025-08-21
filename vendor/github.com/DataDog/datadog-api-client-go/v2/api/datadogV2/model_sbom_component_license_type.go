// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SBOMComponentLicenseType The SBOM component license type.
type SBOMComponentLicenseType string

// List of SBOMComponentLicenseType.
const (
	SBOMCOMPONENTLICENSETYPE_NETWORK_STRONG_COPYLEFT SBOMComponentLicenseType = "network_strong_copyleft"
	SBOMCOMPONENTLICENSETYPE_NON_STANDARD_COPYLEFT   SBOMComponentLicenseType = "non_standard_copyleft"
	SBOMCOMPONENTLICENSETYPE_OTHER_NON_FREE          SBOMComponentLicenseType = "other_non_free"
	SBOMCOMPONENTLICENSETYPE_OTHER_NON_STANDARD      SBOMComponentLicenseType = "other_non_standard"
	SBOMCOMPONENTLICENSETYPE_PERMISSIVE              SBOMComponentLicenseType = "permissive"
	SBOMCOMPONENTLICENSETYPE_PUBLIC_DOMAIN           SBOMComponentLicenseType = "public_domain"
	SBOMCOMPONENTLICENSETYPE_STRONG_COPYLEFT         SBOMComponentLicenseType = "strong_copyleft"
	SBOMCOMPONENTLICENSETYPE_WEAK_COPYLEFT           SBOMComponentLicenseType = "weak_copyleft"
)

var allowedSBOMComponentLicenseTypeEnumValues = []SBOMComponentLicenseType{
	SBOMCOMPONENTLICENSETYPE_NETWORK_STRONG_COPYLEFT,
	SBOMCOMPONENTLICENSETYPE_NON_STANDARD_COPYLEFT,
	SBOMCOMPONENTLICENSETYPE_OTHER_NON_FREE,
	SBOMCOMPONENTLICENSETYPE_OTHER_NON_STANDARD,
	SBOMCOMPONENTLICENSETYPE_PERMISSIVE,
	SBOMCOMPONENTLICENSETYPE_PUBLIC_DOMAIN,
	SBOMCOMPONENTLICENSETYPE_STRONG_COPYLEFT,
	SBOMCOMPONENTLICENSETYPE_WEAK_COPYLEFT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SBOMComponentLicenseType) GetAllowedValues() []SBOMComponentLicenseType {
	return allowedSBOMComponentLicenseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SBOMComponentLicenseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SBOMComponentLicenseType(value)
	return nil
}

// NewSBOMComponentLicenseTypeFromValue returns a pointer to a valid SBOMComponentLicenseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSBOMComponentLicenseTypeFromValue(v string) (*SBOMComponentLicenseType, error) {
	ev := SBOMComponentLicenseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SBOMComponentLicenseType: valid values are %v", v, allowedSBOMComponentLicenseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SBOMComponentLicenseType) IsValid() bool {
	for _, existing := range allowedSBOMComponentLicenseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SBOMComponentLicenseType value.
func (v SBOMComponentLicenseType) Ptr() *SBOMComponentLicenseType {
	return &v
}
