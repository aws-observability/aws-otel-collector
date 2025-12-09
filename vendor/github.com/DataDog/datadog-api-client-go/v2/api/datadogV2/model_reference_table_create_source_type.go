// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReferenceTableCreateSourceType The source type for creating reference table data. Only these source types can be created through this API.
type ReferenceTableCreateSourceType string

// List of ReferenceTableCreateSourceType.
const (
	REFERENCETABLECREATESOURCETYPE_LOCAL_FILE ReferenceTableCreateSourceType = "LOCAL_FILE"
	REFERENCETABLECREATESOURCETYPE_S3         ReferenceTableCreateSourceType = "S3"
	REFERENCETABLECREATESOURCETYPE_GCS        ReferenceTableCreateSourceType = "GCS"
	REFERENCETABLECREATESOURCETYPE_AZURE      ReferenceTableCreateSourceType = "AZURE"
)

var allowedReferenceTableCreateSourceTypeEnumValues = []ReferenceTableCreateSourceType{
	REFERENCETABLECREATESOURCETYPE_LOCAL_FILE,
	REFERENCETABLECREATESOURCETYPE_S3,
	REFERENCETABLECREATESOURCETYPE_GCS,
	REFERENCETABLECREATESOURCETYPE_AZURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReferenceTableCreateSourceType) GetAllowedValues() []ReferenceTableCreateSourceType {
	return allowedReferenceTableCreateSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReferenceTableCreateSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReferenceTableCreateSourceType(value)
	return nil
}

// NewReferenceTableCreateSourceTypeFromValue returns a pointer to a valid ReferenceTableCreateSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReferenceTableCreateSourceTypeFromValue(v string) (*ReferenceTableCreateSourceType, error) {
	ev := ReferenceTableCreateSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReferenceTableCreateSourceType: valid values are %v", v, allowedReferenceTableCreateSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReferenceTableCreateSourceType) IsValid() bool {
	for _, existing := range allowedReferenceTableCreateSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReferenceTableCreateSourceType value.
func (v ReferenceTableCreateSourceType) Ptr() *ReferenceTableCreateSourceType {
	return &v
}
