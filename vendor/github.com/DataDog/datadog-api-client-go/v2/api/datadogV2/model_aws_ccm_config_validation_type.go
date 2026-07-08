// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCcmConfigValidationType AWS CCM config validation resource type.
type AWSCcmConfigValidationType string

// List of AWSCcmConfigValidationType.
const (
	AWSCCMCONFIGVALIDATIONTYPE_CCM_CONFIG_VALIDATION AWSCcmConfigValidationType = "ccm_config_validation"
)

var allowedAWSCcmConfigValidationTypeEnumValues = []AWSCcmConfigValidationType{
	AWSCCMCONFIGVALIDATIONTYPE_CCM_CONFIG_VALIDATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSCcmConfigValidationType) GetAllowedValues() []AWSCcmConfigValidationType {
	return allowedAWSCcmConfigValidationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSCcmConfigValidationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSCcmConfigValidationType(value)
	return nil
}

// NewAWSCcmConfigValidationTypeFromValue returns a pointer to a valid AWSCcmConfigValidationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSCcmConfigValidationTypeFromValue(v string) (*AWSCcmConfigValidationType, error) {
	ev := AWSCcmConfigValidationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSCcmConfigValidationType: valid values are %v", v, allowedAWSCcmConfigValidationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSCcmConfigValidationType) IsValid() bool {
	for _, existing := range allowedAWSCcmConfigValidationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSCcmConfigValidationType value.
func (v AWSCcmConfigValidationType) Ptr() *AWSCcmConfigValidationType {
	return &v
}
