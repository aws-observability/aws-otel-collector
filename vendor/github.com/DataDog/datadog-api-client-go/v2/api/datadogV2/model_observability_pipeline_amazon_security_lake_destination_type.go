// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonSecurityLakeDestinationType The destination type. Always `amazon_security_lake`.
type ObservabilityPipelineAmazonSecurityLakeDestinationType string

// List of ObservabilityPipelineAmazonSecurityLakeDestinationType.
const (
	OBSERVABILITYPIPELINEAMAZONSECURITYLAKEDESTINATIONTYPE_AMAZON_SECURITY_LAKE ObservabilityPipelineAmazonSecurityLakeDestinationType = "amazon_security_lake"
)

var allowedObservabilityPipelineAmazonSecurityLakeDestinationTypeEnumValues = []ObservabilityPipelineAmazonSecurityLakeDestinationType{
	OBSERVABILITYPIPELINEAMAZONSECURITYLAKEDESTINATIONTYPE_AMAZON_SECURITY_LAKE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonSecurityLakeDestinationType) GetAllowedValues() []ObservabilityPipelineAmazonSecurityLakeDestinationType {
	return allowedObservabilityPipelineAmazonSecurityLakeDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonSecurityLakeDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonSecurityLakeDestinationType(value)
	return nil
}

// NewObservabilityPipelineAmazonSecurityLakeDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineAmazonSecurityLakeDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonSecurityLakeDestinationTypeFromValue(v string) (*ObservabilityPipelineAmazonSecurityLakeDestinationType, error) {
	ev := ObservabilityPipelineAmazonSecurityLakeDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonSecurityLakeDestinationType: valid values are %v", v, allowedObservabilityPipelineAmazonSecurityLakeDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonSecurityLakeDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonSecurityLakeDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonSecurityLakeDestinationType value.
func (v ObservabilityPipelineAmazonSecurityLakeDestinationType) Ptr() *ObservabilityPipelineAmazonSecurityLakeDestinationType {
	return &v
}
