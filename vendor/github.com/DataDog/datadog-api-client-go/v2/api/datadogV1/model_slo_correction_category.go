// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOCorrectionCategory Category the SLO correction belongs to.
type SLOCorrectionCategory string

// List of SLOCorrectionCategory.
const (
	SLOCORRECTIONCATEGORY_SCHEDULED_MAINTENANCE  SLOCorrectionCategory = "Scheduled Maintenance"
	SLOCORRECTIONCATEGORY_OUTSIDE_BUSINESS_HOURS SLOCorrectionCategory = "Outside Business Hours"
	SLOCORRECTIONCATEGORY_DEPLOYMENT             SLOCorrectionCategory = "Deployment"
	SLOCORRECTIONCATEGORY_OTHER                  SLOCorrectionCategory = "Other"
)

var allowedSLOCorrectionCategoryEnumValues = []SLOCorrectionCategory{
	SLOCORRECTIONCATEGORY_SCHEDULED_MAINTENANCE,
	SLOCORRECTIONCATEGORY_OUTSIDE_BUSINESS_HOURS,
	SLOCORRECTIONCATEGORY_DEPLOYMENT,
	SLOCORRECTIONCATEGORY_OTHER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SLOCorrectionCategory) GetAllowedValues() []SLOCorrectionCategory {
	return allowedSLOCorrectionCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SLOCorrectionCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SLOCorrectionCategory(value)
	return nil
}

// NewSLOCorrectionCategoryFromValue returns a pointer to a valid SLOCorrectionCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSLOCorrectionCategoryFromValue(v string) (*SLOCorrectionCategory, error) {
	ev := SLOCorrectionCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SLOCorrectionCategory: valid values are %v", v, allowedSLOCorrectionCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SLOCorrectionCategory) IsValid() bool {
	for _, existing := range allowedSLOCorrectionCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SLOCorrectionCategory value.
func (v SLOCorrectionCategory) Ptr() *SLOCorrectionCategory {
	return &v
}
