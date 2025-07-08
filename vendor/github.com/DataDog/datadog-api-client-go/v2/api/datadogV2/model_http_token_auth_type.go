// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPTokenAuthType The definition of `HTTPTokenAuthType` object.
type HTTPTokenAuthType string

// List of HTTPTokenAuthType.
const (
	HTTPTOKENAUTHTYPE_HTTPTOKENAUTH HTTPTokenAuthType = "HTTPTokenAuth"
)

var allowedHTTPTokenAuthTypeEnumValues = []HTTPTokenAuthType{
	HTTPTOKENAUTHTYPE_HTTPTOKENAUTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HTTPTokenAuthType) GetAllowedValues() []HTTPTokenAuthType {
	return allowedHTTPTokenAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HTTPTokenAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HTTPTokenAuthType(value)
	return nil
}

// NewHTTPTokenAuthTypeFromValue returns a pointer to a valid HTTPTokenAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHTTPTokenAuthTypeFromValue(v string) (*HTTPTokenAuthType, error) {
	ev := HTTPTokenAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HTTPTokenAuthType: valid values are %v", v, allowedHTTPTokenAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HTTPTokenAuthType) IsValid() bool {
	for _, existing := range allowedHTTPTokenAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HTTPTokenAuthType value.
func (v HTTPTokenAuthType) Ptr() *HTTPTokenAuthType {
	return &v
}
