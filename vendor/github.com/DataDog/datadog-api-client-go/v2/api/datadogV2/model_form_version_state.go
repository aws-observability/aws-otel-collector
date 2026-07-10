// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormVersionState The state of a form version.
type FormVersionState string

// List of FormVersionState.
const (
	FORMVERSIONSTATE_DRAFT  FormVersionState = "draft"
	FORMVERSIONSTATE_FROZEN FormVersionState = "frozen"
)

var allowedFormVersionStateEnumValues = []FormVersionState{
	FORMVERSIONSTATE_DRAFT,
	FORMVERSIONSTATE_FROZEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormVersionState) GetAllowedValues() []FormVersionState {
	return allowedFormVersionStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormVersionState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormVersionState(value)
	return nil
}

// NewFormVersionStateFromValue returns a pointer to a valid FormVersionState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormVersionStateFromValue(v string) (*FormVersionState, error) {
	ev := FormVersionState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormVersionState: valid values are %v", v, allowedFormVersionStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormVersionState) IsValid() bool {
	for _, existing := range allowedFormVersionStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormVersionState value.
func (v FormVersionState) Ptr() *FormVersionState {
	return &v
}
