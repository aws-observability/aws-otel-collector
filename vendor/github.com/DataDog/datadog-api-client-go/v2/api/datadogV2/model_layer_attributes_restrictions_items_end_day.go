// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LayerAttributesRestrictionsItemsEndDay Defines the end day of the restriction within a Layer.
type LayerAttributesRestrictionsItemsEndDay string

// List of LayerAttributesRestrictionsItemsEndDay.
const (
	LAYERATTRIBUTESRESTRICTIONSITEMSENDDAY_MONDAY    LayerAttributesRestrictionsItemsEndDay = "monday"
	LAYERATTRIBUTESRESTRICTIONSITEMSENDDAY_TUESDAY   LayerAttributesRestrictionsItemsEndDay = "tuesday"
	LAYERATTRIBUTESRESTRICTIONSITEMSENDDAY_WEDNESDAY LayerAttributesRestrictionsItemsEndDay = "wednesday"
	LAYERATTRIBUTESRESTRICTIONSITEMSENDDAY_THURSDAY  LayerAttributesRestrictionsItemsEndDay = "thursday"
	LAYERATTRIBUTESRESTRICTIONSITEMSENDDAY_FRIDAY    LayerAttributesRestrictionsItemsEndDay = "friday"
	LAYERATTRIBUTESRESTRICTIONSITEMSENDDAY_SATURDAY  LayerAttributesRestrictionsItemsEndDay = "saturday"
	LAYERATTRIBUTESRESTRICTIONSITEMSENDDAY_SUNDAY    LayerAttributesRestrictionsItemsEndDay = "sunday"
)

var allowedLayerAttributesRestrictionsItemsEndDayEnumValues = []LayerAttributesRestrictionsItemsEndDay{
	LAYERATTRIBUTESRESTRICTIONSITEMSENDDAY_MONDAY,
	LAYERATTRIBUTESRESTRICTIONSITEMSENDDAY_TUESDAY,
	LAYERATTRIBUTESRESTRICTIONSITEMSENDDAY_WEDNESDAY,
	LAYERATTRIBUTESRESTRICTIONSITEMSENDDAY_THURSDAY,
	LAYERATTRIBUTESRESTRICTIONSITEMSENDDAY_FRIDAY,
	LAYERATTRIBUTESRESTRICTIONSITEMSENDDAY_SATURDAY,
	LAYERATTRIBUTESRESTRICTIONSITEMSENDDAY_SUNDAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LayerAttributesRestrictionsItemsEndDay) GetAllowedValues() []LayerAttributesRestrictionsItemsEndDay {
	return allowedLayerAttributesRestrictionsItemsEndDayEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LayerAttributesRestrictionsItemsEndDay) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LayerAttributesRestrictionsItemsEndDay(value)
	return nil
}

// NewLayerAttributesRestrictionsItemsEndDayFromValue returns a pointer to a valid LayerAttributesRestrictionsItemsEndDay
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLayerAttributesRestrictionsItemsEndDayFromValue(v string) (*LayerAttributesRestrictionsItemsEndDay, error) {
	ev := LayerAttributesRestrictionsItemsEndDay(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LayerAttributesRestrictionsItemsEndDay: valid values are %v", v, allowedLayerAttributesRestrictionsItemsEndDayEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LayerAttributesRestrictionsItemsEndDay) IsValid() bool {
	for _, existing := range allowedLayerAttributesRestrictionsItemsEndDayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LayerAttributesRestrictionsItemsEndDay value.
func (v LayerAttributesRestrictionsItemsEndDay) Ptr() *LayerAttributesRestrictionsItemsEndDay {
	return &v
}
