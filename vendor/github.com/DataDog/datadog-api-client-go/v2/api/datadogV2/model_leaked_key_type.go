// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LeakedKeyType The definition of LeakedKeyType object.
type LeakedKeyType string

// List of LeakedKeyType.
const (
	LEAKEDKEYTYPE_LEAKED_KEYS LeakedKeyType = "leaked_keys"
)

var allowedLeakedKeyTypeEnumValues = []LeakedKeyType{
	LEAKEDKEYTYPE_LEAKED_KEYS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LeakedKeyType) GetAllowedValues() []LeakedKeyType {
	return allowedLeakedKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LeakedKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LeakedKeyType(value)
	return nil
}

// NewLeakedKeyTypeFromValue returns a pointer to a valid LeakedKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLeakedKeyTypeFromValue(v string) (*LeakedKeyType, error) {
	ev := LeakedKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LeakedKeyType: valid values are %v", v, allowedLeakedKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LeakedKeyType) IsValid() bool {
	for _, existing := range allowedLeakedKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LeakedKeyType value.
func (v LeakedKeyType) Ptr() *LeakedKeyType {
	return &v
}
