// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamOnCallRespondersDataRelationshipsRespondersDataItemsType Identifies the resource type for individual user entities associated with on-call response.
type TeamOnCallRespondersDataRelationshipsRespondersDataItemsType string

// List of TeamOnCallRespondersDataRelationshipsRespondersDataItemsType.
const (
	TEAMONCALLRESPONDERSDATARELATIONSHIPSRESPONDERSDATAITEMSTYPE_USERS TeamOnCallRespondersDataRelationshipsRespondersDataItemsType = "users"
)

var allowedTeamOnCallRespondersDataRelationshipsRespondersDataItemsTypeEnumValues = []TeamOnCallRespondersDataRelationshipsRespondersDataItemsType{
	TEAMONCALLRESPONDERSDATARELATIONSHIPSRESPONDERSDATAITEMSTYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamOnCallRespondersDataRelationshipsRespondersDataItemsType) GetAllowedValues() []TeamOnCallRespondersDataRelationshipsRespondersDataItemsType {
	return allowedTeamOnCallRespondersDataRelationshipsRespondersDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamOnCallRespondersDataRelationshipsRespondersDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamOnCallRespondersDataRelationshipsRespondersDataItemsType(value)
	return nil
}

// NewTeamOnCallRespondersDataRelationshipsRespondersDataItemsTypeFromValue returns a pointer to a valid TeamOnCallRespondersDataRelationshipsRespondersDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamOnCallRespondersDataRelationshipsRespondersDataItemsTypeFromValue(v string) (*TeamOnCallRespondersDataRelationshipsRespondersDataItemsType, error) {
	ev := TeamOnCallRespondersDataRelationshipsRespondersDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamOnCallRespondersDataRelationshipsRespondersDataItemsType: valid values are %v", v, allowedTeamOnCallRespondersDataRelationshipsRespondersDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamOnCallRespondersDataRelationshipsRespondersDataItemsType) IsValid() bool {
	for _, existing := range allowedTeamOnCallRespondersDataRelationshipsRespondersDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamOnCallRespondersDataRelationshipsRespondersDataItemsType value.
func (v TeamOnCallRespondersDataRelationshipsRespondersDataItemsType) Ptr() *TeamOnCallRespondersDataRelationshipsRespondersDataItemsType {
	return &v
}
