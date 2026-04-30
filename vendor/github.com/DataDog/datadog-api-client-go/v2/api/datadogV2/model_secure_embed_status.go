// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecureEmbedStatus The status of the secure embed share. Active means the shared dashboard is available. Paused means it is not.
type SecureEmbedStatus string

// List of SecureEmbedStatus.
const (
	SECUREEMBEDSTATUS_ACTIVE SecureEmbedStatus = "active"
	SECUREEMBEDSTATUS_PAUSED SecureEmbedStatus = "paused"
)

var allowedSecureEmbedStatusEnumValues = []SecureEmbedStatus{
	SECUREEMBEDSTATUS_ACTIVE,
	SECUREEMBEDSTATUS_PAUSED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecureEmbedStatus) GetAllowedValues() []SecureEmbedStatus {
	return allowedSecureEmbedStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecureEmbedStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecureEmbedStatus(value)
	return nil
}

// NewSecureEmbedStatusFromValue returns a pointer to a valid SecureEmbedStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecureEmbedStatusFromValue(v string) (*SecureEmbedStatus, error) {
	ev := SecureEmbedStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecureEmbedStatus: valid values are %v", v, allowedSecureEmbedStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecureEmbedStatus) IsValid() bool {
	for _, existing := range allowedSecureEmbedStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecureEmbedStatus value.
func (v SecureEmbedStatus) Ptr() *SecureEmbedStatus {
	return &v
}
