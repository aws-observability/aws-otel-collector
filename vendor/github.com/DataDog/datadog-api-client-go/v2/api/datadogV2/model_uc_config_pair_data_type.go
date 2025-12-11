// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UCConfigPairDataType Azure UC configs resource type.
type UCConfigPairDataType string

// List of UCConfigPairDataType.
const (
	UCCONFIGPAIRDATATYPE_AZURE_UC_CONFIGS UCConfigPairDataType = "azure_uc_configs"
)

var allowedUCConfigPairDataTypeEnumValues = []UCConfigPairDataType{
	UCCONFIGPAIRDATATYPE_AZURE_UC_CONFIGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UCConfigPairDataType) GetAllowedValues() []UCConfigPairDataType {
	return allowedUCConfigPairDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UCConfigPairDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UCConfigPairDataType(value)
	return nil
}

// NewUCConfigPairDataTypeFromValue returns a pointer to a valid UCConfigPairDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUCConfigPairDataTypeFromValue(v string) (*UCConfigPairDataType, error) {
	ev := UCConfigPairDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UCConfigPairDataType: valid values are %v", v, allowedUCConfigPairDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UCConfigPairDataType) IsValid() bool {
	for _, existing := range allowedUCConfigPairDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UCConfigPairDataType value.
func (v UCConfigPairDataType) Ptr() *UCConfigPairDataType {
	return &v
}
