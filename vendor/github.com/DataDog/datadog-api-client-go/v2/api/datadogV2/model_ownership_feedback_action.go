// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipFeedbackAction The feedback action to apply to an inference.
type OwnershipFeedbackAction string

// List of OwnershipFeedbackAction.
const (
	OWNERSHIPFEEDBACKACTION_CONFIRM OwnershipFeedbackAction = "confirm"
	OWNERSHIPFEEDBACKACTION_REJECT  OwnershipFeedbackAction = "reject"
	OWNERSHIPFEEDBACKACTION_CORRECT OwnershipFeedbackAction = "correct"
	OWNERSHIPFEEDBACKACTION_PERSIST OwnershipFeedbackAction = "persist"
)

var allowedOwnershipFeedbackActionEnumValues = []OwnershipFeedbackAction{
	OWNERSHIPFEEDBACKACTION_CONFIRM,
	OWNERSHIPFEEDBACKACTION_REJECT,
	OWNERSHIPFEEDBACKACTION_CORRECT,
	OWNERSHIPFEEDBACKACTION_PERSIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OwnershipFeedbackAction) GetAllowedValues() []OwnershipFeedbackAction {
	return allowedOwnershipFeedbackActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OwnershipFeedbackAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OwnershipFeedbackAction(value)
	return nil
}

// NewOwnershipFeedbackActionFromValue returns a pointer to a valid OwnershipFeedbackAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOwnershipFeedbackActionFromValue(v string) (*OwnershipFeedbackAction, error) {
	ev := OwnershipFeedbackAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OwnershipFeedbackAction: valid values are %v", v, allowedOwnershipFeedbackActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OwnershipFeedbackAction) IsValid() bool {
	for _, existing := range allowedOwnershipFeedbackActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OwnershipFeedbackAction value.
func (v OwnershipFeedbackAction) Ptr() *OwnershipFeedbackAction {
	return &v
}
