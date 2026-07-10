// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipFeedbackType The type of the ownership feedback request resource. The value should always be `ownership_feedback`.
type OwnershipFeedbackType string

// List of OwnershipFeedbackType.
const (
	OWNERSHIPFEEDBACKTYPE_OWNERSHIP_FEEDBACK OwnershipFeedbackType = "ownership_feedback"
)

var allowedOwnershipFeedbackTypeEnumValues = []OwnershipFeedbackType{
	OWNERSHIPFEEDBACKTYPE_OWNERSHIP_FEEDBACK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OwnershipFeedbackType) GetAllowedValues() []OwnershipFeedbackType {
	return allowedOwnershipFeedbackTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OwnershipFeedbackType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OwnershipFeedbackType(value)
	return nil
}

// NewOwnershipFeedbackTypeFromValue returns a pointer to a valid OwnershipFeedbackType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOwnershipFeedbackTypeFromValue(v string) (*OwnershipFeedbackType, error) {
	ev := OwnershipFeedbackType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OwnershipFeedbackType: valid values are %v", v, allowedOwnershipFeedbackTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OwnershipFeedbackType) IsValid() bool {
	for _, existing := range allowedOwnershipFeedbackTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OwnershipFeedbackType value.
func (v OwnershipFeedbackType) Ptr() *OwnershipFeedbackType {
	return &v
}
