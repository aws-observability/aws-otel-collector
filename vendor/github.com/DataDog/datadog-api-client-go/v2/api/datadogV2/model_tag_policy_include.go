// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPolicyInclude A related resource to include alongside a tag policy in the response. Currently the only supported value is `score`.
type TagPolicyInclude string

// List of TagPolicyInclude.
const (
	TAGPOLICYINCLUDE_SCORE TagPolicyInclude = "score"
)

var allowedTagPolicyIncludeEnumValues = []TagPolicyInclude{
	TAGPOLICYINCLUDE_SCORE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagPolicyInclude) GetAllowedValues() []TagPolicyInclude {
	return allowedTagPolicyIncludeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagPolicyInclude) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagPolicyInclude(value)
	return nil
}

// NewTagPolicyIncludeFromValue returns a pointer to a valid TagPolicyInclude
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagPolicyIncludeFromValue(v string) (*TagPolicyInclude, error) {
	ev := TagPolicyInclude(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagPolicyInclude: valid values are %v", v, allowedTagPolicyIncludeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagPolicyInclude) IsValid() bool {
	for _, existing := range allowedTagPolicyIncludeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagPolicyInclude value.
func (v TagPolicyInclude) Ptr() *TagPolicyInclude {
	return &v
}
