// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudInventoryCloudProviderRequestType Always `cloud_provider`.
type CloudInventoryCloudProviderRequestType string

// List of CloudInventoryCloudProviderRequestType.
const (
	CLOUDINVENTORYCLOUDPROVIDERREQUESTTYPE_CLOUD_PROVIDER CloudInventoryCloudProviderRequestType = "cloud_provider"
)

var allowedCloudInventoryCloudProviderRequestTypeEnumValues = []CloudInventoryCloudProviderRequestType{
	CLOUDINVENTORYCLOUDPROVIDERREQUESTTYPE_CLOUD_PROVIDER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CloudInventoryCloudProviderRequestType) GetAllowedValues() []CloudInventoryCloudProviderRequestType {
	return allowedCloudInventoryCloudProviderRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudInventoryCloudProviderRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudInventoryCloudProviderRequestType(value)
	return nil
}

// NewCloudInventoryCloudProviderRequestTypeFromValue returns a pointer to a valid CloudInventoryCloudProviderRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudInventoryCloudProviderRequestTypeFromValue(v string) (*CloudInventoryCloudProviderRequestType, error) {
	ev := CloudInventoryCloudProviderRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudInventoryCloudProviderRequestType: valid values are %v", v, allowedCloudInventoryCloudProviderRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudInventoryCloudProviderRequestType) IsValid() bool {
	for _, existing := range allowedCloudInventoryCloudProviderRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudInventoryCloudProviderRequestType value.
func (v CloudInventoryCloudProviderRequestType) Ptr() *CloudInventoryCloudProviderRequestType {
	return &v
}
