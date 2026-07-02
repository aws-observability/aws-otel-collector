// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardShareType Type of dashboard sharing.
type SharedDashboardShareType string

// List of SharedDashboardShareType.
const (
	SHAREDDASHBOARDSHARETYPE_OPEN         SharedDashboardShareType = "open"
	SHAREDDASHBOARDSHARETYPE_INVITE       SharedDashboardShareType = "invite"
	SHAREDDASHBOARDSHARETYPE_EMBED        SharedDashboardShareType = "embed"
	SHAREDDASHBOARDSHARETYPE_SECURE_EMBED SharedDashboardShareType = "secure-embed"
)

var allowedSharedDashboardShareTypeEnumValues = []SharedDashboardShareType{
	SHAREDDASHBOARDSHARETYPE_OPEN,
	SHAREDDASHBOARDSHARETYPE_INVITE,
	SHAREDDASHBOARDSHARETYPE_EMBED,
	SHAREDDASHBOARDSHARETYPE_SECURE_EMBED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SharedDashboardShareType) GetAllowedValues() []SharedDashboardShareType {
	return allowedSharedDashboardShareTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SharedDashboardShareType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SharedDashboardShareType(value)
	return nil
}

// NewSharedDashboardShareTypeFromValue returns a pointer to a valid SharedDashboardShareType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSharedDashboardShareTypeFromValue(v string) (*SharedDashboardShareType, error) {
	ev := SharedDashboardShareType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SharedDashboardShareType: valid values are %v", v, allowedSharedDashboardShareTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SharedDashboardShareType) IsValid() bool {
	for _, existing := range allowedSharedDashboardShareTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SharedDashboardShareType value.
func (v SharedDashboardShareType) Ptr() *SharedDashboardShareType {
	return &v
}
