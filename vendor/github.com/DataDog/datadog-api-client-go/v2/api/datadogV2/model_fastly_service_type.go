// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FastlyServiceType The JSON:API type for this API. Should always be `fastly-services`.
type FastlyServiceType string

// List of FastlyServiceType.
const (
	FASTLYSERVICETYPE_FASTLY_SERVICES FastlyServiceType = "fastly-services"
)

var allowedFastlyServiceTypeEnumValues = []FastlyServiceType{
	FASTLYSERVICETYPE_FASTLY_SERVICES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FastlyServiceType) GetAllowedValues() []FastlyServiceType {
	return allowedFastlyServiceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FastlyServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FastlyServiceType(value)
	return nil
}

// NewFastlyServiceTypeFromValue returns a pointer to a valid FastlyServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFastlyServiceTypeFromValue(v string) (*FastlyServiceType, error) {
	ev := FastlyServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FastlyServiceType: valid values are %v", v, allowedFastlyServiceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FastlyServiceType) IsValid() bool {
	for _, existing := range allowedFastlyServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FastlyServiceType value.
func (v FastlyServiceType) Ptr() *FastlyServiceType {
	return &v
}
