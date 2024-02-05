// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudConfigurationRuleType The rule type.
type CloudConfigurationRuleType string

// List of CloudConfigurationRuleType.
const (
	CLOUDCONFIGURATIONRULETYPE_CLOUD_CONFIGURATION CloudConfigurationRuleType = "cloud_configuration"
)

var allowedCloudConfigurationRuleTypeEnumValues = []CloudConfigurationRuleType{
	CLOUDCONFIGURATIONRULETYPE_CLOUD_CONFIGURATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CloudConfigurationRuleType) GetAllowedValues() []CloudConfigurationRuleType {
	return allowedCloudConfigurationRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudConfigurationRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudConfigurationRuleType(value)
	return nil
}

// NewCloudConfigurationRuleTypeFromValue returns a pointer to a valid CloudConfigurationRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudConfigurationRuleTypeFromValue(v string) (*CloudConfigurationRuleType, error) {
	ev := CloudConfigurationRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudConfigurationRuleType: valid values are %v", v, allowedCloudConfigurationRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudConfigurationRuleType) IsValid() bool {
	for _, existing := range allowedCloudConfigurationRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudConfigurationRuleType value.
func (v CloudConfigurationRuleType) Ptr() *CloudConfigurationRuleType {
	return &v
}
