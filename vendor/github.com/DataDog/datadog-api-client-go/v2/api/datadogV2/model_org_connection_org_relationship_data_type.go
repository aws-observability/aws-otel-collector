// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgConnectionOrgRelationshipDataType The type of the organization relationship.
type OrgConnectionOrgRelationshipDataType string

// List of OrgConnectionOrgRelationshipDataType.
const (
	ORGCONNECTIONORGRELATIONSHIPDATATYPE_ORGS OrgConnectionOrgRelationshipDataType = "orgs"
)

var allowedOrgConnectionOrgRelationshipDataTypeEnumValues = []OrgConnectionOrgRelationshipDataType{
	ORGCONNECTIONORGRELATIONSHIPDATATYPE_ORGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgConnectionOrgRelationshipDataType) GetAllowedValues() []OrgConnectionOrgRelationshipDataType {
	return allowedOrgConnectionOrgRelationshipDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgConnectionOrgRelationshipDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgConnectionOrgRelationshipDataType(value)
	return nil
}

// NewOrgConnectionOrgRelationshipDataTypeFromValue returns a pointer to a valid OrgConnectionOrgRelationshipDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgConnectionOrgRelationshipDataTypeFromValue(v string) (*OrgConnectionOrgRelationshipDataType, error) {
	ev := OrgConnectionOrgRelationshipDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgConnectionOrgRelationshipDataType: valid values are %v", v, allowedOrgConnectionOrgRelationshipDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgConnectionOrgRelationshipDataType) IsValid() bool {
	for _, existing := range allowedOrgConnectionOrgRelationshipDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgConnectionOrgRelationshipDataType value.
func (v OrgConnectionOrgRelationshipDataType) Ptr() *OrgConnectionOrgRelationshipDataType {
	return &v
}
