// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSEventBridgeType Amazon EventBridge resource type.
type AWSEventBridgeType string

// List of AWSEventBridgeType.
const (
	AWSEVENTBRIDGETYPE_EVENT_BRIDGE AWSEventBridgeType = "event_bridge"
)

var allowedAWSEventBridgeTypeEnumValues = []AWSEventBridgeType{
	AWSEVENTBRIDGETYPE_EVENT_BRIDGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSEventBridgeType) GetAllowedValues() []AWSEventBridgeType {
	return allowedAWSEventBridgeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSEventBridgeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSEventBridgeType(value)
	return nil
}

// NewAWSEventBridgeTypeFromValue returns a pointer to a valid AWSEventBridgeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSEventBridgeTypeFromValue(v string) (*AWSEventBridgeType, error) {
	ev := AWSEventBridgeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSEventBridgeType: valid values are %v", v, allowedAWSEventBridgeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSEventBridgeType) IsValid() bool {
	for _, existing := range allowedAWSEventBridgeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSEventBridgeType value.
func (v AWSEventBridgeType) Ptr() *AWSEventBridgeType {
	return &v
}
