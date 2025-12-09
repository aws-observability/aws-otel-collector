// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpsgenieServiceType Opsgenie service resource type.
type OpsgenieServiceType string

// List of OpsgenieServiceType.
const (
	OPSGENIESERVICETYPE_OPSGENIE_SERVICE OpsgenieServiceType = "opsgenie-service"
)

var allowedOpsgenieServiceTypeEnumValues = []OpsgenieServiceType{
	OPSGENIESERVICETYPE_OPSGENIE_SERVICE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OpsgenieServiceType) GetAllowedValues() []OpsgenieServiceType {
	return allowedOpsgenieServiceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OpsgenieServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OpsgenieServiceType(value)
	return nil
}

// NewOpsgenieServiceTypeFromValue returns a pointer to a valid OpsgenieServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOpsgenieServiceTypeFromValue(v string) (*OpsgenieServiceType, error) {
	ev := OpsgenieServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OpsgenieServiceType: valid values are %v", v, allowedOpsgenieServiceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OpsgenieServiceType) IsValid() bool {
	for _, existing := range allowedOpsgenieServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpsgenieServiceType value.
func (v OpsgenieServiceType) Ptr() *OpsgenieServiceType {
	return &v
}
