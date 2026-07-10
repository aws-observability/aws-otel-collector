// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmUnifiedHostType The JSON:API type for unified host resources. The value should always be `unified_host`.
type CsmUnifiedHostType string

// List of CsmUnifiedHostType.
const (
	CSMUNIFIEDHOSTTYPE_UNIFIED_HOST CsmUnifiedHostType = "unified_host"
)

var allowedCsmUnifiedHostTypeEnumValues = []CsmUnifiedHostType{
	CSMUNIFIEDHOSTTYPE_UNIFIED_HOST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CsmUnifiedHostType) GetAllowedValues() []CsmUnifiedHostType {
	return allowedCsmUnifiedHostTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CsmUnifiedHostType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CsmUnifiedHostType(value)
	return nil
}

// NewCsmUnifiedHostTypeFromValue returns a pointer to a valid CsmUnifiedHostType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCsmUnifiedHostTypeFromValue(v string) (*CsmUnifiedHostType, error) {
	ev := CsmUnifiedHostType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CsmUnifiedHostType: valid values are %v", v, allowedCsmUnifiedHostTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CsmUnifiedHostType) IsValid() bool {
	for _, existing := range allowedCsmUnifiedHostTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CsmUnifiedHostType value.
func (v CsmUnifiedHostType) Ptr() *CsmUnifiedHostType {
	return &v
}
