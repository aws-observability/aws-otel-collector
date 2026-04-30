// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WildcardWidgetSpecificationType Type of specification used by the wildcard widget.
type WildcardWidgetSpecificationType string

// List of WildcardWidgetSpecificationType.
const (
	WILDCARDWIDGETSPECIFICATIONTYPE_VEGA      WildcardWidgetSpecificationType = "vega"
	WILDCARDWIDGETSPECIFICATIONTYPE_VEGA_LITE WildcardWidgetSpecificationType = "vega-lite"
)

var allowedWildcardWidgetSpecificationTypeEnumValues = []WildcardWidgetSpecificationType{
	WILDCARDWIDGETSPECIFICATIONTYPE_VEGA,
	WILDCARDWIDGETSPECIFICATIONTYPE_VEGA_LITE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WildcardWidgetSpecificationType) GetAllowedValues() []WildcardWidgetSpecificationType {
	return allowedWildcardWidgetSpecificationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WildcardWidgetSpecificationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WildcardWidgetSpecificationType(value)
	return nil
}

// NewWildcardWidgetSpecificationTypeFromValue returns a pointer to a valid WildcardWidgetSpecificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWildcardWidgetSpecificationTypeFromValue(v string) (*WildcardWidgetSpecificationType, error) {
	ev := WildcardWidgetSpecificationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WildcardWidgetSpecificationType: valid values are %v", v, allowedWildcardWidgetSpecificationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WildcardWidgetSpecificationType) IsValid() bool {
	for _, existing := range allowedWildcardWidgetSpecificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WildcardWidgetSpecificationType value.
func (v WildcardWidgetSpecificationType) Ptr() *WildcardWidgetSpecificationType {
	return &v
}
