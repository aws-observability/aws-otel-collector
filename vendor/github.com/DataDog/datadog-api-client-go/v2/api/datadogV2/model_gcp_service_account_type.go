// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPServiceAccountType The type of account.
type GCPServiceAccountType string

// List of GCPServiceAccountType.
const (
	GCPSERVICEACCOUNTTYPE_GCP_SERVICE_ACCOUNT GCPServiceAccountType = "gcp_service_account"
)

var allowedGCPServiceAccountTypeEnumValues = []GCPServiceAccountType{
	GCPSERVICEACCOUNTTYPE_GCP_SERVICE_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GCPServiceAccountType) GetAllowedValues() []GCPServiceAccountType {
	return allowedGCPServiceAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GCPServiceAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GCPServiceAccountType(value)
	return nil
}

// NewGCPServiceAccountTypeFromValue returns a pointer to a valid GCPServiceAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGCPServiceAccountTypeFromValue(v string) (*GCPServiceAccountType, error) {
	ev := GCPServiceAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GCPServiceAccountType: valid values are %v", v, allowedGCPServiceAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GCPServiceAccountType) IsValid() bool {
	for _, existing := range allowedGCPServiceAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GCPServiceAccountType value.
func (v GCPServiceAccountType) Ptr() *GCPServiceAccountType {
	return &v
}
