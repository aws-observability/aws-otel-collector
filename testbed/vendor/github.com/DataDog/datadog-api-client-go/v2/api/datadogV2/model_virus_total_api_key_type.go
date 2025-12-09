// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VirusTotalAPIKeyType The definition of the `VirusTotalAPIKey` object.
type VirusTotalAPIKeyType string

// List of VirusTotalAPIKeyType.
const (
	VIRUSTOTALAPIKEYTYPE_VIRUSTOTALAPIKEY VirusTotalAPIKeyType = "VirusTotalAPIKey"
)

var allowedVirusTotalAPIKeyTypeEnumValues = []VirusTotalAPIKeyType{
	VIRUSTOTALAPIKEYTYPE_VIRUSTOTALAPIKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *VirusTotalAPIKeyType) GetAllowedValues() []VirusTotalAPIKeyType {
	return allowedVirusTotalAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *VirusTotalAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = VirusTotalAPIKeyType(value)
	return nil
}

// NewVirusTotalAPIKeyTypeFromValue returns a pointer to a valid VirusTotalAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewVirusTotalAPIKeyTypeFromValue(v string) (*VirusTotalAPIKeyType, error) {
	ev := VirusTotalAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for VirusTotalAPIKeyType: valid values are %v", v, allowedVirusTotalAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v VirusTotalAPIKeyType) IsValid() bool {
	for _, existing := range allowedVirusTotalAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VirusTotalAPIKeyType value.
func (v VirusTotalAPIKeyType) Ptr() *VirusTotalAPIKeyType {
	return &v
}
