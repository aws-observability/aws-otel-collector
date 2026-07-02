// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudInventoryCloudProviderId Cloud provider for this sync configuration (`aws`, `gcp`, or `azure`). For requests, must match the provider block supplied under `attributes`.
type CloudInventoryCloudProviderId string

// List of CloudInventoryCloudProviderId.
const (
	CLOUDINVENTORYCLOUDPROVIDERID_AWS   CloudInventoryCloudProviderId = "aws"
	CLOUDINVENTORYCLOUDPROVIDERID_GCP   CloudInventoryCloudProviderId = "gcp"
	CLOUDINVENTORYCLOUDPROVIDERID_AZURE CloudInventoryCloudProviderId = "azure"
)

var allowedCloudInventoryCloudProviderIdEnumValues = []CloudInventoryCloudProviderId{
	CLOUDINVENTORYCLOUDPROVIDERID_AWS,
	CLOUDINVENTORYCLOUDPROVIDERID_GCP,
	CLOUDINVENTORYCLOUDPROVIDERID_AZURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CloudInventoryCloudProviderId) GetAllowedValues() []CloudInventoryCloudProviderId {
	return allowedCloudInventoryCloudProviderIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudInventoryCloudProviderId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudInventoryCloudProviderId(value)
	return nil
}

// NewCloudInventoryCloudProviderIdFromValue returns a pointer to a valid CloudInventoryCloudProviderId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudInventoryCloudProviderIdFromValue(v string) (*CloudInventoryCloudProviderId, error) {
	ev := CloudInventoryCloudProviderId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudInventoryCloudProviderId: valid values are %v", v, allowedCloudInventoryCloudProviderIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudInventoryCloudProviderId) IsValid() bool {
	for _, existing := range allowedCloudInventoryCloudProviderIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudInventoryCloudProviderId value.
func (v CloudInventoryCloudProviderId) Ptr() *CloudInventoryCloudProviderId {
	return &v
}
