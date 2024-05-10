// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSEventBridgeCreateStatus The event source status "created".
type AWSEventBridgeCreateStatus string

// List of AWSEventBridgeCreateStatus.
const (
	AWSEVENTBRIDGECREATESTATUS_CREATED AWSEventBridgeCreateStatus = "created"
)

var allowedAWSEventBridgeCreateStatusEnumValues = []AWSEventBridgeCreateStatus{
	AWSEVENTBRIDGECREATESTATUS_CREATED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSEventBridgeCreateStatus) GetAllowedValues() []AWSEventBridgeCreateStatus {
	return allowedAWSEventBridgeCreateStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSEventBridgeCreateStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSEventBridgeCreateStatus(value)
	return nil
}

// NewAWSEventBridgeCreateStatusFromValue returns a pointer to a valid AWSEventBridgeCreateStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSEventBridgeCreateStatusFromValue(v string) (*AWSEventBridgeCreateStatus, error) {
	ev := AWSEventBridgeCreateStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSEventBridgeCreateStatus: valid values are %v", v, allowedAWSEventBridgeCreateStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSEventBridgeCreateStatus) IsValid() bool {
	for _, existing := range allowedAWSEventBridgeCreateStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSEventBridgeCreateStatus value.
func (v AWSEventBridgeCreateStatus) Ptr() *AWSEventBridgeCreateStatus {
	return &v
}
