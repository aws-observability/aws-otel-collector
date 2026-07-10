// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmAgentlessHostResourceType The type of cloud resource for an agentless host.
type CsmAgentlessHostResourceType string

// List of CsmAgentlessHostResourceType.
const (
	CSMAGENTLESSHOSTRESOURCETYPE_AWS_EC2_INSTANCE               CsmAgentlessHostResourceType = "aws_ec2_instance"
	CSMAGENTLESSHOSTRESOURCETYPE_AZURE_VIRTUAL_MACHINE_INSTANCE CsmAgentlessHostResourceType = "azure_virtual_machine_instance"
	CSMAGENTLESSHOSTRESOURCETYPE_GCP_COMPUTE_INSTANCE           CsmAgentlessHostResourceType = "gcp_compute_instance"
	CSMAGENTLESSHOSTRESOURCETYPE_OCI_INSTANCE                   CsmAgentlessHostResourceType = "oci_instance"
)

var allowedCsmAgentlessHostResourceTypeEnumValues = []CsmAgentlessHostResourceType{
	CSMAGENTLESSHOSTRESOURCETYPE_AWS_EC2_INSTANCE,
	CSMAGENTLESSHOSTRESOURCETYPE_AZURE_VIRTUAL_MACHINE_INSTANCE,
	CSMAGENTLESSHOSTRESOURCETYPE_GCP_COMPUTE_INSTANCE,
	CSMAGENTLESSHOSTRESOURCETYPE_OCI_INSTANCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CsmAgentlessHostResourceType) GetAllowedValues() []CsmAgentlessHostResourceType {
	return allowedCsmAgentlessHostResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CsmAgentlessHostResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CsmAgentlessHostResourceType(value)
	return nil
}

// NewCsmAgentlessHostResourceTypeFromValue returns a pointer to a valid CsmAgentlessHostResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCsmAgentlessHostResourceTypeFromValue(v string) (*CsmAgentlessHostResourceType, error) {
	ev := CsmAgentlessHostResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CsmAgentlessHostResourceType: valid values are %v", v, allowedCsmAgentlessHostResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CsmAgentlessHostResourceType) IsValid() bool {
	for _, existing := range allowedCsmAgentlessHostResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CsmAgentlessHostResourceType value.
func (v CsmAgentlessHostResourceType) Ptr() *CsmAgentlessHostResourceType {
	return &v
}
