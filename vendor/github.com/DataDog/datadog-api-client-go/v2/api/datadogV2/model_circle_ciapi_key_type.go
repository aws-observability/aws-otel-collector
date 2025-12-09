// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CircleCIAPIKeyType The definition of the `CircleCIAPIKey` object.
type CircleCIAPIKeyType string

// List of CircleCIAPIKeyType.
const (
	CIRCLECIAPIKEYTYPE_CIRCLECIAPIKEY CircleCIAPIKeyType = "CircleCIAPIKey"
)

var allowedCircleCIAPIKeyTypeEnumValues = []CircleCIAPIKeyType{
	CIRCLECIAPIKEYTYPE_CIRCLECIAPIKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CircleCIAPIKeyType) GetAllowedValues() []CircleCIAPIKeyType {
	return allowedCircleCIAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CircleCIAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CircleCIAPIKeyType(value)
	return nil
}

// NewCircleCIAPIKeyTypeFromValue returns a pointer to a valid CircleCIAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCircleCIAPIKeyTypeFromValue(v string) (*CircleCIAPIKeyType, error) {
	ev := CircleCIAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CircleCIAPIKeyType: valid values are %v", v, allowedCircleCIAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CircleCIAPIKeyType) IsValid() bool {
	for _, existing := range allowedCircleCIAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CircleCIAPIKeyType value.
func (v CircleCIAPIKeyType) Ptr() *CircleCIAPIKeyType {
	return &v
}
