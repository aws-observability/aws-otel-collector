// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgConnectionTypeEnum Available connection types between organizations.
type OrgConnectionTypeEnum string

// List of OrgConnectionTypeEnum.
const (
	ORGCONNECTIONTYPEENUM_LOGS    OrgConnectionTypeEnum = "logs"
	ORGCONNECTIONTYPEENUM_METRICS OrgConnectionTypeEnum = "metrics"
)

var allowedOrgConnectionTypeEnumEnumValues = []OrgConnectionTypeEnum{
	ORGCONNECTIONTYPEENUM_LOGS,
	ORGCONNECTIONTYPEENUM_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgConnectionTypeEnum) GetAllowedValues() []OrgConnectionTypeEnum {
	return allowedOrgConnectionTypeEnumEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgConnectionTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgConnectionTypeEnum(value)
	return nil
}

// NewOrgConnectionTypeEnumFromValue returns a pointer to a valid OrgConnectionTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgConnectionTypeEnumFromValue(v string) (*OrgConnectionTypeEnum, error) {
	ev := OrgConnectionTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgConnectionTypeEnum: valid values are %v", v, allowedOrgConnectionTypeEnumEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgConnectionTypeEnum) IsValid() bool {
	for _, existing := range allowedOrgConnectionTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgConnectionTypeEnum value.
func (v OrgConnectionTypeEnum) Ptr() *OrgConnectionTypeEnum {
	return &v
}
