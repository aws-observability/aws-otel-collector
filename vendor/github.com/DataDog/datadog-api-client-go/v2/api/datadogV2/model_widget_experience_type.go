// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetExperienceType Widget experience types that differentiate between the products using the specific widget.
type WidgetExperienceType string

// List of WidgetExperienceType.
const (
	WIDGETEXPERIENCETYPE_CCM_REPORTS       WidgetExperienceType = "ccm_reports"
	WIDGETEXPERIENCETYPE_LOGS_REPORTS      WidgetExperienceType = "logs_reports"
	WIDGETEXPERIENCETYPE_CSV_REPORTS       WidgetExperienceType = "csv_reports"
	WIDGETEXPERIENCETYPE_PRODUCT_ANALYTICS WidgetExperienceType = "product_analytics"
)

var allowedWidgetExperienceTypeEnumValues = []WidgetExperienceType{
	WIDGETEXPERIENCETYPE_CCM_REPORTS,
	WIDGETEXPERIENCETYPE_LOGS_REPORTS,
	WIDGETEXPERIENCETYPE_CSV_REPORTS,
	WIDGETEXPERIENCETYPE_PRODUCT_ANALYTICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetExperienceType) GetAllowedValues() []WidgetExperienceType {
	return allowedWidgetExperienceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetExperienceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetExperienceType(value)
	return nil
}

// NewWidgetExperienceTypeFromValue returns a pointer to a valid WidgetExperienceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetExperienceTypeFromValue(v string) (*WidgetExperienceType, error) {
	ev := WidgetExperienceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetExperienceType: valid values are %v", v, allowedWidgetExperienceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetExperienceType) IsValid() bool {
	for _, existing := range allowedWidgetExperienceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetExperienceType value.
func (v WidgetExperienceType) Ptr() *WidgetExperienceType {
	return &v
}
