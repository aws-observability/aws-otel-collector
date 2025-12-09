// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudflareGlobalAPITokenType The definition of the `CloudflareGlobalAPIToken` object.
type CloudflareGlobalAPITokenType string

// List of CloudflareGlobalAPITokenType.
const (
	CLOUDFLAREGLOBALAPITOKENTYPE_CLOUDFLAREGLOBALAPITOKEN CloudflareGlobalAPITokenType = "CloudflareGlobalAPIToken"
)

var allowedCloudflareGlobalAPITokenTypeEnumValues = []CloudflareGlobalAPITokenType{
	CLOUDFLAREGLOBALAPITOKENTYPE_CLOUDFLAREGLOBALAPITOKEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CloudflareGlobalAPITokenType) GetAllowedValues() []CloudflareGlobalAPITokenType {
	return allowedCloudflareGlobalAPITokenTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudflareGlobalAPITokenType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudflareGlobalAPITokenType(value)
	return nil
}

// NewCloudflareGlobalAPITokenTypeFromValue returns a pointer to a valid CloudflareGlobalAPITokenType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudflareGlobalAPITokenTypeFromValue(v string) (*CloudflareGlobalAPITokenType, error) {
	ev := CloudflareGlobalAPITokenType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudflareGlobalAPITokenType: valid values are %v", v, allowedCloudflareGlobalAPITokenTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudflareGlobalAPITokenType) IsValid() bool {
	for _, existing := range allowedCloudflareGlobalAPITokenTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudflareGlobalAPITokenType value.
func (v CloudflareGlobalAPITokenType) Ptr() *CloudflareGlobalAPITokenType {
	return &v
}
