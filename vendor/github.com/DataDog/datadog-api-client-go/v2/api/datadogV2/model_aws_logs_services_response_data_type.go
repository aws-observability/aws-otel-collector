// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSLogsServicesResponseDataType The `AWSLogsServicesResponseData` `type`.
type AWSLogsServicesResponseDataType string

// List of AWSLogsServicesResponseDataType.
const (
	AWSLOGSSERVICESRESPONSEDATATYPE_LOGS_SERVICES AWSLogsServicesResponseDataType = "logs_services"
)

var allowedAWSLogsServicesResponseDataTypeEnumValues = []AWSLogsServicesResponseDataType{
	AWSLOGSSERVICESRESPONSEDATATYPE_LOGS_SERVICES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSLogsServicesResponseDataType) GetAllowedValues() []AWSLogsServicesResponseDataType {
	return allowedAWSLogsServicesResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSLogsServicesResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSLogsServicesResponseDataType(value)
	return nil
}

// NewAWSLogsServicesResponseDataTypeFromValue returns a pointer to a valid AWSLogsServicesResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSLogsServicesResponseDataTypeFromValue(v string) (*AWSLogsServicesResponseDataType, error) {
	ev := AWSLogsServicesResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSLogsServicesResponseDataType: valid values are %v", v, allowedAWSLogsServicesResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSLogsServicesResponseDataType) IsValid() bool {
	for _, existing := range allowedAWSLogsServicesResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSLogsServicesResponseDataType value.
func (v AWSLogsServicesResponseDataType) Ptr() *AWSLogsServicesResponseDataType {
	return &v
}
