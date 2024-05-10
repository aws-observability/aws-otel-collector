// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPSTSDelegateAccountType The type of account.
type GCPSTSDelegateAccountType string

// List of GCPSTSDelegateAccountType.
const (
	GCPSTSDELEGATEACCOUNTTYPE_GCP_STS_DELEGATE GCPSTSDelegateAccountType = "gcp_sts_delegate"
)

var allowedGCPSTSDelegateAccountTypeEnumValues = []GCPSTSDelegateAccountType{
	GCPSTSDELEGATEACCOUNTTYPE_GCP_STS_DELEGATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GCPSTSDelegateAccountType) GetAllowedValues() []GCPSTSDelegateAccountType {
	return allowedGCPSTSDelegateAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GCPSTSDelegateAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GCPSTSDelegateAccountType(value)
	return nil
}

// NewGCPSTSDelegateAccountTypeFromValue returns a pointer to a valid GCPSTSDelegateAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGCPSTSDelegateAccountTypeFromValue(v string) (*GCPSTSDelegateAccountType, error) {
	ev := GCPSTSDelegateAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GCPSTSDelegateAccountType: valid values are %v", v, allowedGCPSTSDelegateAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GCPSTSDelegateAccountType) IsValid() bool {
	for _, existing := range allowedGCPSTSDelegateAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GCPSTSDelegateAccountType value.
func (v GCPSTSDelegateAccountType) Ptr() *GCPSTSDelegateAccountType {
	return &v
}
