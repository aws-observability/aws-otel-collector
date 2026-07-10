// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPolicySource The telemetry source that a tag policy applies to.
type TagPolicySource string

// List of TagPolicySource.
const (
	TAGPOLICYSOURCE_LOGS    TagPolicySource = "logs"
	TAGPOLICYSOURCE_SPANS   TagPolicySource = "spans"
	TAGPOLICYSOURCE_METRICS TagPolicySource = "metrics"
	TAGPOLICYSOURCE_RUM     TagPolicySource = "rum"
	TAGPOLICYSOURCE_FEED    TagPolicySource = "feed"
)

var allowedTagPolicySourceEnumValues = []TagPolicySource{
	TAGPOLICYSOURCE_LOGS,
	TAGPOLICYSOURCE_SPANS,
	TAGPOLICYSOURCE_METRICS,
	TAGPOLICYSOURCE_RUM,
	TAGPOLICYSOURCE_FEED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagPolicySource) GetAllowedValues() []TagPolicySource {
	return allowedTagPolicySourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagPolicySource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagPolicySource(value)
	return nil
}

// NewTagPolicySourceFromValue returns a pointer to a valid TagPolicySource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagPolicySourceFromValue(v string) (*TagPolicySource, error) {
	ev := TagPolicySource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagPolicySource: valid values are %v", v, allowedTagPolicySourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagPolicySource) IsValid() bool {
	for _, existing := range allowedTagPolicySourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagPolicySource value.
func (v TagPolicySource) Ptr() *TagPolicySource {
	return &v
}
