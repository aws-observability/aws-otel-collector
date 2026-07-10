// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPolicyResourceType JSON:API resource type for a tag policy.
type TagPolicyResourceType string

// List of TagPolicyResourceType.
const (
	TAGPOLICYRESOURCETYPE_TAG_POLICY TagPolicyResourceType = "tag_policy"
)

var allowedTagPolicyResourceTypeEnumValues = []TagPolicyResourceType{
	TAGPOLICYRESOURCETYPE_TAG_POLICY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagPolicyResourceType) GetAllowedValues() []TagPolicyResourceType {
	return allowedTagPolicyResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagPolicyResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagPolicyResourceType(value)
	return nil
}

// NewTagPolicyResourceTypeFromValue returns a pointer to a valid TagPolicyResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagPolicyResourceTypeFromValue(v string) (*TagPolicyResourceType, error) {
	ev := TagPolicyResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagPolicyResourceType: valid values are %v", v, allowedTagPolicyResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagPolicyResourceType) IsValid() bool {
	for _, existing := range allowedTagPolicyResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagPolicyResourceType value.
func (v TagPolicyResourceType) Ptr() *TagPolicyResourceType {
	return &v
}
