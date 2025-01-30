// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSNewExternalIDResponseDataType The `AWSNewExternalIDResponseData` `type`.
type AWSNewExternalIDResponseDataType string

// List of AWSNewExternalIDResponseDataType.
const (
	AWSNEWEXTERNALIDRESPONSEDATATYPE_EXTERNAL_ID AWSNewExternalIDResponseDataType = "external_id"
)

var allowedAWSNewExternalIDResponseDataTypeEnumValues = []AWSNewExternalIDResponseDataType{
	AWSNEWEXTERNALIDRESPONSEDATATYPE_EXTERNAL_ID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSNewExternalIDResponseDataType) GetAllowedValues() []AWSNewExternalIDResponseDataType {
	return allowedAWSNewExternalIDResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSNewExternalIDResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSNewExternalIDResponseDataType(value)
	return nil
}

// NewAWSNewExternalIDResponseDataTypeFromValue returns a pointer to a valid AWSNewExternalIDResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSNewExternalIDResponseDataTypeFromValue(v string) (*AWSNewExternalIDResponseDataType, error) {
	ev := AWSNewExternalIDResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSNewExternalIDResponseDataType: valid values are %v", v, allowedAWSNewExternalIDResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSNewExternalIDResponseDataType) IsValid() bool {
	for _, existing := range allowedAWSNewExternalIDResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSNewExternalIDResponseDataType value.
func (v AWSNewExternalIDResponseDataType) Ptr() *AWSNewExternalIDResponseDataType {
	return &v
}
