// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ValidateAPIKeyStatus Status of the validation. Always `ok` when both the API key and the application key are valid.
type ValidateAPIKeyStatus string

// List of ValidateAPIKeyStatus.
const (
	VALIDATEAPIKEYSTATUS_OK ValidateAPIKeyStatus = "ok"
)

var allowedValidateAPIKeyStatusEnumValues = []ValidateAPIKeyStatus{
	VALIDATEAPIKEYSTATUS_OK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ValidateAPIKeyStatus) GetAllowedValues() []ValidateAPIKeyStatus {
	return allowedValidateAPIKeyStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ValidateAPIKeyStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ValidateAPIKeyStatus(value)
	return nil
}

// NewValidateAPIKeyStatusFromValue returns a pointer to a valid ValidateAPIKeyStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewValidateAPIKeyStatusFromValue(v string) (*ValidateAPIKeyStatus, error) {
	ev := ValidateAPIKeyStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ValidateAPIKeyStatus: valid values are %v", v, allowedValidateAPIKeyStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ValidateAPIKeyStatus) IsValid() bool {
	for _, existing := range allowedValidateAPIKeyStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ValidateAPIKeyStatus value.
func (v ValidateAPIKeyStatus) Ptr() *ValidateAPIKeyStatus {
	return &v
}
