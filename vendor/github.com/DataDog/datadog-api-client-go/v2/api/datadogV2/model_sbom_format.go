// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SBOMFormat The SBOM standard
type SBOMFormat string

// List of SBOMFormat.
const (
	SBOMFORMAT_CYCLONEDX SBOMFormat = "CycloneDX"
	SBOMFORMAT_SPDX      SBOMFormat = "SPDX"
)

var allowedSBOMFormatEnumValues = []SBOMFormat{
	SBOMFORMAT_CYCLONEDX,
	SBOMFORMAT_SPDX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SBOMFormat) GetAllowedValues() []SBOMFormat {
	return allowedSBOMFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SBOMFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SBOMFormat(value)
	return nil
}

// NewSBOMFormatFromValue returns a pointer to a valid SBOMFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSBOMFormatFromValue(v string) (*SBOMFormat, error) {
	ev := SBOMFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SBOMFormat: valid values are %v", v, allowedSBOMFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SBOMFormat) IsValid() bool {
	for _, existing := range allowedSBOMFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SBOMFormat value.
func (v SBOMFormat) Ptr() *SBOMFormat {
	return &v
}
