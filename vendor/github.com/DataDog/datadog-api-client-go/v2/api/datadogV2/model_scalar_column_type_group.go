// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScalarColumnTypeGroup The type of column present for groups.
type ScalarColumnTypeGroup string

// List of ScalarColumnTypeGroup.
const (
	SCALARCOLUMNTYPEGROUP_GROUP ScalarColumnTypeGroup = "group"
)

var allowedScalarColumnTypeGroupEnumValues = []ScalarColumnTypeGroup{
	SCALARCOLUMNTYPEGROUP_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScalarColumnTypeGroup) GetAllowedValues() []ScalarColumnTypeGroup {
	return allowedScalarColumnTypeGroupEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScalarColumnTypeGroup) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScalarColumnTypeGroup(value)
	return nil
}

// NewScalarColumnTypeGroupFromValue returns a pointer to a valid ScalarColumnTypeGroup
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScalarColumnTypeGroupFromValue(v string) (*ScalarColumnTypeGroup, error) {
	ev := ScalarColumnTypeGroup(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScalarColumnTypeGroup: valid values are %v", v, allowedScalarColumnTypeGroupEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScalarColumnTypeGroup) IsValid() bool {
	for _, existing := range allowedScalarColumnTypeGroupEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScalarColumnTypeGroup value.
func (v ScalarColumnTypeGroup) Ptr() *ScalarColumnTypeGroup {
	return &v
}
