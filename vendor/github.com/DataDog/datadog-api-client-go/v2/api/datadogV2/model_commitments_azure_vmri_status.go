// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsAzureVMRIStatus Status of an Azure VM Reserved Instance.
type CommitmentsAzureVMRIStatus string

// List of CommitmentsAzureVMRIStatus.
const (
	COMMITMENTSAZUREVMRISTATUS_RUNNING   CommitmentsAzureVMRIStatus = "running"
	COMMITMENTSAZUREVMRISTATUS_EXPIRED   CommitmentsAzureVMRIStatus = "expired"
	COMMITMENTSAZUREVMRISTATUS_CANCELLED CommitmentsAzureVMRIStatus = "cancelled"
)

var allowedCommitmentsAzureVMRIStatusEnumValues = []CommitmentsAzureVMRIStatus{
	COMMITMENTSAZUREVMRISTATUS_RUNNING,
	COMMITMENTSAZUREVMRISTATUS_EXPIRED,
	COMMITMENTSAZUREVMRISTATUS_CANCELLED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CommitmentsAzureVMRIStatus) GetAllowedValues() []CommitmentsAzureVMRIStatus {
	return allowedCommitmentsAzureVMRIStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CommitmentsAzureVMRIStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CommitmentsAzureVMRIStatus(value)
	return nil
}

// NewCommitmentsAzureVMRIStatusFromValue returns a pointer to a valid CommitmentsAzureVMRIStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCommitmentsAzureVMRIStatusFromValue(v string) (*CommitmentsAzureVMRIStatus, error) {
	ev := CommitmentsAzureVMRIStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CommitmentsAzureVMRIStatus: valid values are %v", v, allowedCommitmentsAzureVMRIStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CommitmentsAzureVMRIStatus) IsValid() bool {
	for _, existing := range allowedCommitmentsAzureVMRIStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommitmentsAzureVMRIStatus value.
func (v CommitmentsAzureVMRIStatus) Ptr() *CommitmentsAzureVMRIStatus {
	return &v
}
