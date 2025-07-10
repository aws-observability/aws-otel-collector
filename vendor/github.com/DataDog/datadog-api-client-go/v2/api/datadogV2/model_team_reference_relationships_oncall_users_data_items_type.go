// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamReferenceRelationshipsOncallUsersDataItemsType Users resource type.
type TeamReferenceRelationshipsOncallUsersDataItemsType string

// List of TeamReferenceRelationshipsOncallUsersDataItemsType.
const (
	TEAMREFERENCERELATIONSHIPSONCALLUSERSDATAITEMSTYPE_USERS TeamReferenceRelationshipsOncallUsersDataItemsType = "users"
)

var allowedTeamReferenceRelationshipsOncallUsersDataItemsTypeEnumValues = []TeamReferenceRelationshipsOncallUsersDataItemsType{
	TEAMREFERENCERELATIONSHIPSONCALLUSERSDATAITEMSTYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamReferenceRelationshipsOncallUsersDataItemsType) GetAllowedValues() []TeamReferenceRelationshipsOncallUsersDataItemsType {
	return allowedTeamReferenceRelationshipsOncallUsersDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamReferenceRelationshipsOncallUsersDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamReferenceRelationshipsOncallUsersDataItemsType(value)
	return nil
}

// NewTeamReferenceRelationshipsOncallUsersDataItemsTypeFromValue returns a pointer to a valid TeamReferenceRelationshipsOncallUsersDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamReferenceRelationshipsOncallUsersDataItemsTypeFromValue(v string) (*TeamReferenceRelationshipsOncallUsersDataItemsType, error) {
	ev := TeamReferenceRelationshipsOncallUsersDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamReferenceRelationshipsOncallUsersDataItemsType: valid values are %v", v, allowedTeamReferenceRelationshipsOncallUsersDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamReferenceRelationshipsOncallUsersDataItemsType) IsValid() bool {
	for _, existing := range allowedTeamReferenceRelationshipsOncallUsersDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamReferenceRelationshipsOncallUsersDataItemsType value.
func (v TeamReferenceRelationshipsOncallUsersDataItemsType) Ptr() *TeamReferenceRelationshipsOncallUsersDataItemsType {
	return &v
}
