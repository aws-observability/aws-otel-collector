// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipEvidenceType The type of the ownership evidence resource. The value should always be `ownership_evidence`.
type OwnershipEvidenceType string

// List of OwnershipEvidenceType.
const (
	OWNERSHIPEVIDENCETYPE_OWNERSHIP_EVIDENCE OwnershipEvidenceType = "ownership_evidence"
)

var allowedOwnershipEvidenceTypeEnumValues = []OwnershipEvidenceType{
	OWNERSHIPEVIDENCETYPE_OWNERSHIP_EVIDENCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OwnershipEvidenceType) GetAllowedValues() []OwnershipEvidenceType {
	return allowedOwnershipEvidenceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OwnershipEvidenceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OwnershipEvidenceType(value)
	return nil
}

// NewOwnershipEvidenceTypeFromValue returns a pointer to a valid OwnershipEvidenceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOwnershipEvidenceTypeFromValue(v string) (*OwnershipEvidenceType, error) {
	ev := OwnershipEvidenceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OwnershipEvidenceType: valid values are %v", v, allowedOwnershipEvidenceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OwnershipEvidenceType) IsValid() bool {
	for _, existing := range allowedOwnershipEvidenceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OwnershipEvidenceType value.
func (v OwnershipEvidenceType) Ptr() *OwnershipEvidenceType {
	return &v
}
