// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSEventBridgeDeleteStatus The event source status "empty".
type AWSEventBridgeDeleteStatus string

// List of AWSEventBridgeDeleteStatus.
const (
	AWSEVENTBRIDGEDELETESTATUS_EMPTY AWSEventBridgeDeleteStatus = "empty"
)

var allowedAWSEventBridgeDeleteStatusEnumValues = []AWSEventBridgeDeleteStatus{
	AWSEVENTBRIDGEDELETESTATUS_EMPTY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSEventBridgeDeleteStatus) GetAllowedValues() []AWSEventBridgeDeleteStatus {
	return allowedAWSEventBridgeDeleteStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSEventBridgeDeleteStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSEventBridgeDeleteStatus(value)
	return nil
}

// NewAWSEventBridgeDeleteStatusFromValue returns a pointer to a valid AWSEventBridgeDeleteStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSEventBridgeDeleteStatusFromValue(v string) (*AWSEventBridgeDeleteStatus, error) {
	ev := AWSEventBridgeDeleteStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSEventBridgeDeleteStatus: valid values are %v", v, allowedAWSEventBridgeDeleteStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSEventBridgeDeleteStatus) IsValid() bool {
	for _, existing := range allowedAWSEventBridgeDeleteStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSEventBridgeDeleteStatus value.
func (v AWSEventBridgeDeleteStatus) Ptr() *AWSEventBridgeDeleteStatus {
	return &v
}
