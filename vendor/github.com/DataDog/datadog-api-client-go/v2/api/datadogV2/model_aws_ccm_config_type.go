// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCcmConfigType AWS CCM Config resource type.
type AWSCcmConfigType string

// List of AWSCcmConfigType.
const (
	AWSCCMCONFIGTYPE_CCM_CONFIG AWSCcmConfigType = "ccm_config"
)

var allowedAWSCcmConfigTypeEnumValues = []AWSCcmConfigType{
	AWSCCMCONFIGTYPE_CCM_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSCcmConfigType) GetAllowedValues() []AWSCcmConfigType {
	return allowedAWSCcmConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSCcmConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSCcmConfigType(value)
	return nil
}

// NewAWSCcmConfigTypeFromValue returns a pointer to a valid AWSCcmConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSCcmConfigTypeFromValue(v string) (*AWSCcmConfigType, error) {
	ev := AWSCcmConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSCcmConfigType: valid values are %v", v, allowedAWSCcmConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSCcmConfigType) IsValid() bool {
	for _, existing := range allowedAWSCcmConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSCcmConfigType value.
func (v AWSCcmConfigType) Ptr() *AWSCcmConfigType {
	return &v
}
