// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AsanaAccessTokenType The definition of the `AsanaAccessToken` object.
type AsanaAccessTokenType string

// List of AsanaAccessTokenType.
const (
	ASANAACCESSTOKENTYPE_ASANAACCESSTOKEN AsanaAccessTokenType = "AsanaAccessToken"
)

var allowedAsanaAccessTokenTypeEnumValues = []AsanaAccessTokenType{
	ASANAACCESSTOKENTYPE_ASANAACCESSTOKEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AsanaAccessTokenType) GetAllowedValues() []AsanaAccessTokenType {
	return allowedAsanaAccessTokenTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AsanaAccessTokenType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AsanaAccessTokenType(value)
	return nil
}

// NewAsanaAccessTokenTypeFromValue returns a pointer to a valid AsanaAccessTokenType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAsanaAccessTokenTypeFromValue(v string) (*AsanaAccessTokenType, error) {
	ev := AsanaAccessTokenType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AsanaAccessTokenType: valid values are %v", v, allowedAsanaAccessTokenTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AsanaAccessTokenType) IsValid() bool {
	for _, existing := range allowedAsanaAccessTokenTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AsanaAccessTokenType value.
func (v AsanaAccessTokenType) Ptr() *AsanaAccessTokenType {
	return &v
}
