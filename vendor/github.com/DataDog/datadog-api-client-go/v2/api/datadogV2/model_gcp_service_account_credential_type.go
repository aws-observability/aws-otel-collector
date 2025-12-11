// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPServiceAccountCredentialType The definition of the `GCPServiceAccount` object.
type GCPServiceAccountCredentialType string

// List of GCPServiceAccountCredentialType.
const (
	GCPSERVICEACCOUNTCREDENTIALTYPE_GCPSERVICEACCOUNT GCPServiceAccountCredentialType = "GCPServiceAccount"
)

var allowedGCPServiceAccountCredentialTypeEnumValues = []GCPServiceAccountCredentialType{
	GCPSERVICEACCOUNTCREDENTIALTYPE_GCPSERVICEACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GCPServiceAccountCredentialType) GetAllowedValues() []GCPServiceAccountCredentialType {
	return allowedGCPServiceAccountCredentialTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GCPServiceAccountCredentialType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GCPServiceAccountCredentialType(value)
	return nil
}

// NewGCPServiceAccountCredentialTypeFromValue returns a pointer to a valid GCPServiceAccountCredentialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGCPServiceAccountCredentialTypeFromValue(v string) (*GCPServiceAccountCredentialType, error) {
	ev := GCPServiceAccountCredentialType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GCPServiceAccountCredentialType: valid values are %v", v, allowedGCPServiceAccountCredentialTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GCPServiceAccountCredentialType) IsValid() bool {
	for _, existing := range allowedGCPServiceAccountCredentialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GCPServiceAccountCredentialType value.
func (v GCPServiceAccountCredentialType) Ptr() *GCPServiceAccountCredentialType {
	return &v
}
