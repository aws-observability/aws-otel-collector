// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAccountPartition AWS partition your AWS account is scoped to. Defaults to `aws`.
// See [Partitions](https://docs.aws.amazon.com/whitepapers/latest/aws-fault-isolation-boundaries/partitions.html) in the AWS documentation for more information.
type AWSAccountPartition string

// List of AWSAccountPartition.
const (
	AWSACCOUNTPARTITION_AWS        AWSAccountPartition = "aws"
	AWSACCOUNTPARTITION_AWS_CN     AWSAccountPartition = "aws-cn"
	AWSACCOUNTPARTITION_AWS_US_GOV AWSAccountPartition = "aws-us-gov"
)

var allowedAWSAccountPartitionEnumValues = []AWSAccountPartition{
	AWSACCOUNTPARTITION_AWS,
	AWSACCOUNTPARTITION_AWS_CN,
	AWSACCOUNTPARTITION_AWS_US_GOV,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSAccountPartition) GetAllowedValues() []AWSAccountPartition {
	return allowedAWSAccountPartitionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSAccountPartition) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSAccountPartition(value)
	return nil
}

// NewAWSAccountPartitionFromValue returns a pointer to a valid AWSAccountPartition
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSAccountPartitionFromValue(v string) (*AWSAccountPartition, error) {
	ev := AWSAccountPartition(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSAccountPartition: valid values are %v", v, allowedAWSAccountPartitionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSAccountPartition) IsValid() bool {
	for _, existing := range allowedAWSAccountPartitionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSAccountPartition value.
func (v AWSAccountPartition) Ptr() *AWSAccountPartition {
	return &v
}
