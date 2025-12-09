// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestOptionsHTTPVersion HTTP version to use for a Synthetic test.
type SyntheticsTestOptionsHTTPVersion string

// List of SyntheticsTestOptionsHTTPVersion.
const (
	SYNTHETICSTESTOPTIONSHTTPVERSION_HTTP1 SyntheticsTestOptionsHTTPVersion = "http1"
	SYNTHETICSTESTOPTIONSHTTPVERSION_HTTP2 SyntheticsTestOptionsHTTPVersion = "http2"
	SYNTHETICSTESTOPTIONSHTTPVERSION_ANY   SyntheticsTestOptionsHTTPVersion = "any"
)

var allowedSyntheticsTestOptionsHTTPVersionEnumValues = []SyntheticsTestOptionsHTTPVersion{
	SYNTHETICSTESTOPTIONSHTTPVERSION_HTTP1,
	SYNTHETICSTESTOPTIONSHTTPVERSION_HTTP2,
	SYNTHETICSTESTOPTIONSHTTPVERSION_ANY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestOptionsHTTPVersion) GetAllowedValues() []SyntheticsTestOptionsHTTPVersion {
	return allowedSyntheticsTestOptionsHTTPVersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestOptionsHTTPVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestOptionsHTTPVersion(value)
	return nil
}

// NewSyntheticsTestOptionsHTTPVersionFromValue returns a pointer to a valid SyntheticsTestOptionsHTTPVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestOptionsHTTPVersionFromValue(v string) (*SyntheticsTestOptionsHTTPVersion, error) {
	ev := SyntheticsTestOptionsHTTPVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestOptionsHTTPVersion: valid values are %v", v, allowedSyntheticsTestOptionsHTTPVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestOptionsHTTPVersion) IsValid() bool {
	for _, existing := range allowedSyntheticsTestOptionsHTTPVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestOptionsHTTPVersion value.
func (v SyntheticsTestOptionsHTTPVersion) Ptr() *SyntheticsTestOptionsHTTPVersion {
	return &v
}
