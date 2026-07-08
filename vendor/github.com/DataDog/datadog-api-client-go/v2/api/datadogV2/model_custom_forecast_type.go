// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomForecastType The type of the custom forecast resource. Must be `custom_forecast`.
type CustomForecastType string

// List of CustomForecastType.
const (
	CUSTOMFORECASTTYPE_CUSTOM_FORECAST CustomForecastType = "custom_forecast"
)

var allowedCustomForecastTypeEnumValues = []CustomForecastType{
	CUSTOMFORECASTTYPE_CUSTOM_FORECAST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomForecastType) GetAllowedValues() []CustomForecastType {
	return allowedCustomForecastTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomForecastType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomForecastType(value)
	return nil
}

// NewCustomForecastTypeFromValue returns a pointer to a valid CustomForecastType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomForecastTypeFromValue(v string) (*CustomForecastType, error) {
	ev := CustomForecastType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomForecastType: valid values are %v", v, allowedCustomForecastTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomForecastType) IsValid() bool {
	for _, existing := range allowedCustomForecastTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomForecastType value.
func (v CustomForecastType) Ptr() *CustomForecastType {
	return &v
}
