// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmCloudProvider The cloud provider of a host resource.
type CsmCloudProvider string

// List of CsmCloudProvider.
const (
	CSMCLOUDPROVIDER_AWS   CsmCloudProvider = "aws"
	CSMCLOUDPROVIDER_GCP   CsmCloudProvider = "gcp"
	CSMCLOUDPROVIDER_AZURE CsmCloudProvider = "azure"
	CSMCLOUDPROVIDER_OCI   CsmCloudProvider = "oci"
)

var allowedCsmCloudProviderEnumValues = []CsmCloudProvider{
	CSMCLOUDPROVIDER_AWS,
	CSMCLOUDPROVIDER_GCP,
	CSMCLOUDPROVIDER_AZURE,
	CSMCLOUDPROVIDER_OCI,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CsmCloudProvider) GetAllowedValues() []CsmCloudProvider {
	return allowedCsmCloudProviderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CsmCloudProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CsmCloudProvider(value)
	return nil
}

// NewCsmCloudProviderFromValue returns a pointer to a valid CsmCloudProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCsmCloudProviderFromValue(v string) (*CsmCloudProvider, error) {
	ev := CsmCloudProvider(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CsmCloudProvider: valid values are %v", v, allowedCsmCloudProviderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CsmCloudProvider) IsValid() bool {
	for _, existing := range allowedCsmCloudProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CsmCloudProvider value.
func (v CsmCloudProvider) Ptr() *CsmCloudProvider {
	return &v
}
