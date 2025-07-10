// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LayerAttributesRestrictionsItemsStartDay Defines the start day of the restriction within a Layer.
type LayerAttributesRestrictionsItemsStartDay string

// List of LayerAttributesRestrictionsItemsStartDay.
const (
	LAYERATTRIBUTESRESTRICTIONSITEMSSTARTDAY_MONDAY    LayerAttributesRestrictionsItemsStartDay = "monday"
	LAYERATTRIBUTESRESTRICTIONSITEMSSTARTDAY_TUESDAY   LayerAttributesRestrictionsItemsStartDay = "tuesday"
	LAYERATTRIBUTESRESTRICTIONSITEMSSTARTDAY_WEDNESDAY LayerAttributesRestrictionsItemsStartDay = "wednesday"
	LAYERATTRIBUTESRESTRICTIONSITEMSSTARTDAY_THURSDAY  LayerAttributesRestrictionsItemsStartDay = "thursday"
	LAYERATTRIBUTESRESTRICTIONSITEMSSTARTDAY_FRIDAY    LayerAttributesRestrictionsItemsStartDay = "friday"
	LAYERATTRIBUTESRESTRICTIONSITEMSSTARTDAY_SATURDAY  LayerAttributesRestrictionsItemsStartDay = "saturday"
	LAYERATTRIBUTESRESTRICTIONSITEMSSTARTDAY_SUNDAY    LayerAttributesRestrictionsItemsStartDay = "sunday"
)

var allowedLayerAttributesRestrictionsItemsStartDayEnumValues = []LayerAttributesRestrictionsItemsStartDay{
	LAYERATTRIBUTESRESTRICTIONSITEMSSTARTDAY_MONDAY,
	LAYERATTRIBUTESRESTRICTIONSITEMSSTARTDAY_TUESDAY,
	LAYERATTRIBUTESRESTRICTIONSITEMSSTARTDAY_WEDNESDAY,
	LAYERATTRIBUTESRESTRICTIONSITEMSSTARTDAY_THURSDAY,
	LAYERATTRIBUTESRESTRICTIONSITEMSSTARTDAY_FRIDAY,
	LAYERATTRIBUTESRESTRICTIONSITEMSSTARTDAY_SATURDAY,
	LAYERATTRIBUTESRESTRICTIONSITEMSSTARTDAY_SUNDAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LayerAttributesRestrictionsItemsStartDay) GetAllowedValues() []LayerAttributesRestrictionsItemsStartDay {
	return allowedLayerAttributesRestrictionsItemsStartDayEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LayerAttributesRestrictionsItemsStartDay) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LayerAttributesRestrictionsItemsStartDay(value)
	return nil
}

// NewLayerAttributesRestrictionsItemsStartDayFromValue returns a pointer to a valid LayerAttributesRestrictionsItemsStartDay
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLayerAttributesRestrictionsItemsStartDayFromValue(v string) (*LayerAttributesRestrictionsItemsStartDay, error) {
	ev := LayerAttributesRestrictionsItemsStartDay(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LayerAttributesRestrictionsItemsStartDay: valid values are %v", v, allowedLayerAttributesRestrictionsItemsStartDayEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LayerAttributesRestrictionsItemsStartDay) IsValid() bool {
	for _, existing := range allowedLayerAttributesRestrictionsItemsStartDayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LayerAttributesRestrictionsItemsStartDay value.
func (v LayerAttributesRestrictionsItemsStartDay) Ptr() *LayerAttributesRestrictionsItemsStartDay {
	return &v
}
