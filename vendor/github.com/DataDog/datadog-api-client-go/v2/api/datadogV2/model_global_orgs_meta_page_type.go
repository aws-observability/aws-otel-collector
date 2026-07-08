// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GlobalOrgsMetaPageType Type of global orgs pagination.
type GlobalOrgsMetaPageType string

// List of GlobalOrgsMetaPageType.
const (
	GLOBALORGSMETAPAGETYPE_CURSOR GlobalOrgsMetaPageType = "cursor"
)

var allowedGlobalOrgsMetaPageTypeEnumValues = []GlobalOrgsMetaPageType{
	GLOBALORGSMETAPAGETYPE_CURSOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GlobalOrgsMetaPageType) GetAllowedValues() []GlobalOrgsMetaPageType {
	return allowedGlobalOrgsMetaPageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GlobalOrgsMetaPageType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GlobalOrgsMetaPageType(value)
	return nil
}

// NewGlobalOrgsMetaPageTypeFromValue returns a pointer to a valid GlobalOrgsMetaPageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGlobalOrgsMetaPageTypeFromValue(v string) (*GlobalOrgsMetaPageType, error) {
	ev := GlobalOrgsMetaPageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GlobalOrgsMetaPageType: valid values are %v", v, allowedGlobalOrgsMetaPageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GlobalOrgsMetaPageType) IsValid() bool {
	for _, existing := range allowedGlobalOrgsMetaPageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GlobalOrgsMetaPageType value.
func (v GlobalOrgsMetaPageType) Ptr() *GlobalOrgsMetaPageType {
	return &v
}
