// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SBOMComponentType The SBOM component type
type SBOMComponentType string

// List of SBOMComponentType.
const (
	SBOMCOMPONENTTYPE_APPLICATION            SBOMComponentType = "application"
	SBOMCOMPONENTTYPE_CONTAINER              SBOMComponentType = "container"
	SBOMCOMPONENTTYPE_DATA                   SBOMComponentType = "data"
	SBOMCOMPONENTTYPE_DEVICE                 SBOMComponentType = "device"
	SBOMCOMPONENTTYPE_DEVICE_DRIVER          SBOMComponentType = "device-driver"
	SBOMCOMPONENTTYPE_FILE                   SBOMComponentType = "file"
	SBOMCOMPONENTTYPE_FIRMWARE               SBOMComponentType = "firmware"
	SBOMCOMPONENTTYPE_FRAMEWORK              SBOMComponentType = "framework"
	SBOMCOMPONENTTYPE_LIBRARY                SBOMComponentType = "library"
	SBOMCOMPONENTTYPE_MACHINE_LEARNING_MODEL SBOMComponentType = "machine-learning-model"
	SBOMCOMPONENTTYPE_OPERATING_SYSTEM       SBOMComponentType = "operating-system"
	SBOMCOMPONENTTYPE_PLATFORM               SBOMComponentType = "platform"
)

var allowedSBOMComponentTypeEnumValues = []SBOMComponentType{
	SBOMCOMPONENTTYPE_APPLICATION,
	SBOMCOMPONENTTYPE_CONTAINER,
	SBOMCOMPONENTTYPE_DATA,
	SBOMCOMPONENTTYPE_DEVICE,
	SBOMCOMPONENTTYPE_DEVICE_DRIVER,
	SBOMCOMPONENTTYPE_FILE,
	SBOMCOMPONENTTYPE_FIRMWARE,
	SBOMCOMPONENTTYPE_FRAMEWORK,
	SBOMCOMPONENTTYPE_LIBRARY,
	SBOMCOMPONENTTYPE_MACHINE_LEARNING_MODEL,
	SBOMCOMPONENTTYPE_OPERATING_SYSTEM,
	SBOMCOMPONENTTYPE_PLATFORM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SBOMComponentType) GetAllowedValues() []SBOMComponentType {
	return allowedSBOMComponentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SBOMComponentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SBOMComponentType(value)
	return nil
}

// NewSBOMComponentTypeFromValue returns a pointer to a valid SBOMComponentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSBOMComponentTypeFromValue(v string) (*SBOMComponentType, error) {
	ev := SBOMComponentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SBOMComponentType: valid values are %v", v, allowedSBOMComponentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SBOMComponentType) IsValid() bool {
	for _, existing := range allowedSBOMComponentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SBOMComponentType value.
func (v SBOMComponentType) Ptr() *SBOMComponentType {
	return &v
}
