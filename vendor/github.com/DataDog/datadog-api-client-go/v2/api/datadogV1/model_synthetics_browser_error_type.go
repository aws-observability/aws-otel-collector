// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsBrowserErrorType Error type returned by a browser test.
type SyntheticsBrowserErrorType string

// List of SyntheticsBrowserErrorType.
const (
	SYNTHETICSBROWSERERRORTYPE_NETWORK SyntheticsBrowserErrorType = "network"
	SYNTHETICSBROWSERERRORTYPE_JS      SyntheticsBrowserErrorType = "js"
)

var allowedSyntheticsBrowserErrorTypeEnumValues = []SyntheticsBrowserErrorType{
	SYNTHETICSBROWSERERRORTYPE_NETWORK,
	SYNTHETICSBROWSERERRORTYPE_JS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsBrowserErrorType) GetAllowedValues() []SyntheticsBrowserErrorType {
	return allowedSyntheticsBrowserErrorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsBrowserErrorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsBrowserErrorType(value)
	return nil
}

// NewSyntheticsBrowserErrorTypeFromValue returns a pointer to a valid SyntheticsBrowserErrorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsBrowserErrorTypeFromValue(v string) (*SyntheticsBrowserErrorType, error) {
	ev := SyntheticsBrowserErrorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsBrowserErrorType: valid values are %v", v, allowedSyntheticsBrowserErrorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsBrowserErrorType) IsValid() bool {
	for _, existing := range allowedSyntheticsBrowserErrorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsBrowserErrorType value.
func (v SyntheticsBrowserErrorType) Ptr() *SyntheticsBrowserErrorType {
	return &v
}
