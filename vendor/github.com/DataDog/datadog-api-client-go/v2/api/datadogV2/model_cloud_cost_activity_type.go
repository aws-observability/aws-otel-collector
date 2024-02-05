// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudCostActivityType Type of Cloud Cost Activity.
type CloudCostActivityType string

// List of CloudCostActivityType.
const (
	CLOUDCOSTACTIVITYTYPE_CLOUD_COST_ACTIVITY CloudCostActivityType = "cloud_cost_activity"
)

var allowedCloudCostActivityTypeEnumValues = []CloudCostActivityType{
	CLOUDCOSTACTIVITYTYPE_CLOUD_COST_ACTIVITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CloudCostActivityType) GetAllowedValues() []CloudCostActivityType {
	return allowedCloudCostActivityTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudCostActivityType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudCostActivityType(value)
	return nil
}

// NewCloudCostActivityTypeFromValue returns a pointer to a valid CloudCostActivityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudCostActivityTypeFromValue(v string) (*CloudCostActivityType, error) {
	ev := CloudCostActivityType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudCostActivityType: valid values are %v", v, allowedCloudCostActivityTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudCostActivityType) IsValid() bool {
	for _, existing := range allowedCloudCostActivityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudCostActivityType value.
func (v CloudCostActivityType) Ptr() *CloudCostActivityType {
	return &v
}
