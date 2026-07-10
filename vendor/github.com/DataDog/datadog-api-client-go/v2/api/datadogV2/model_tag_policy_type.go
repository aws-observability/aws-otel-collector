// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPolicyType How the policy is enforced. `blocking` rejects telemetry that violates the policy.
// `surfacing` only highlights non-compliant telemetry without blocking it.
type TagPolicyType string

// List of TagPolicyType.
const (
	TAGPOLICYTYPE_BLOCKING  TagPolicyType = "blocking"
	TAGPOLICYTYPE_SURFACING TagPolicyType = "surfacing"
)

var allowedTagPolicyTypeEnumValues = []TagPolicyType{
	TAGPOLICYTYPE_BLOCKING,
	TAGPOLICYTYPE_SURFACING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagPolicyType) GetAllowedValues() []TagPolicyType {
	return allowedTagPolicyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagPolicyType(value)
	return nil
}

// NewTagPolicyTypeFromValue returns a pointer to a valid TagPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagPolicyTypeFromValue(v string) (*TagPolicyType, error) {
	ev := TagPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagPolicyType: valid values are %v", v, allowedTagPolicyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagPolicyType) IsValid() bool {
	for _, existing := range allowedTagPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagPolicyType value.
func (v TagPolicyType) Ptr() *TagPolicyType {
	return &v
}
