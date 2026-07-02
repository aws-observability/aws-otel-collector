// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseBulkResourceType JSON:API resource type for bulk case operations.
type CaseBulkResourceType string

// List of CaseBulkResourceType.
const (
	CASEBULKRESOURCETYPE_BULK CaseBulkResourceType = "bulk"
)

var allowedCaseBulkResourceTypeEnumValues = []CaseBulkResourceType{
	CASEBULKRESOURCETYPE_BULK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseBulkResourceType) GetAllowedValues() []CaseBulkResourceType {
	return allowedCaseBulkResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseBulkResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseBulkResourceType(value)
	return nil
}

// NewCaseBulkResourceTypeFromValue returns a pointer to a valid CaseBulkResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseBulkResourceTypeFromValue(v string) (*CaseBulkResourceType, error) {
	ev := CaseBulkResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseBulkResourceType: valid values are %v", v, allowedCaseBulkResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseBulkResourceType) IsValid() bool {
	for _, existing := range allowedCaseBulkResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseBulkResourceType value.
func (v CaseBulkResourceType) Ptr() *CaseBulkResourceType {
	return &v
}
