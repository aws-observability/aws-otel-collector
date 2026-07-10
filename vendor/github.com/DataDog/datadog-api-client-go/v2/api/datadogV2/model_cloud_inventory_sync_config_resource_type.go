// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudInventorySyncConfigResourceType Always `sync_configs`.
type CloudInventorySyncConfigResourceType string

// List of CloudInventorySyncConfigResourceType.
const (
	CLOUDINVENTORYSYNCCONFIGRESOURCETYPE_SYNC_CONFIGS CloudInventorySyncConfigResourceType = "sync_configs"
)

var allowedCloudInventorySyncConfigResourceTypeEnumValues = []CloudInventorySyncConfigResourceType{
	CLOUDINVENTORYSYNCCONFIGRESOURCETYPE_SYNC_CONFIGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CloudInventorySyncConfigResourceType) GetAllowedValues() []CloudInventorySyncConfigResourceType {
	return allowedCloudInventorySyncConfigResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudInventorySyncConfigResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudInventorySyncConfigResourceType(value)
	return nil
}

// NewCloudInventorySyncConfigResourceTypeFromValue returns a pointer to a valid CloudInventorySyncConfigResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudInventorySyncConfigResourceTypeFromValue(v string) (*CloudInventorySyncConfigResourceType, error) {
	ev := CloudInventorySyncConfigResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudInventorySyncConfigResourceType: valid values are %v", v, allowedCloudInventorySyncConfigResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudInventorySyncConfigResourceType) IsValid() bool {
	for _, existing := range allowedCloudInventorySyncConfigResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudInventorySyncConfigResourceType value.
func (v CloudInventorySyncConfigResourceType) Ptr() *CloudInventorySyncConfigResourceType {
	return &v
}
