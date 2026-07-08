// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatuspageUrlSettingType Statuspage URL setting resource type.
type StatuspageUrlSettingType string

// List of StatuspageUrlSettingType.
const (
	STATUSPAGEURLSETTINGTYPE_STATUSPAGE_URL_SETTING StatuspageUrlSettingType = "statuspage-url-setting"
)

var allowedStatuspageUrlSettingTypeEnumValues = []StatuspageUrlSettingType{
	STATUSPAGEURLSETTINGTYPE_STATUSPAGE_URL_SETTING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StatuspageUrlSettingType) GetAllowedValues() []StatuspageUrlSettingType {
	return allowedStatuspageUrlSettingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StatuspageUrlSettingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StatuspageUrlSettingType(value)
	return nil
}

// NewStatuspageUrlSettingTypeFromValue returns a pointer to a valid StatuspageUrlSettingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStatuspageUrlSettingTypeFromValue(v string) (*StatuspageUrlSettingType, error) {
	ev := StatuspageUrlSettingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StatuspageUrlSettingType: valid values are %v", v, allowedStatuspageUrlSettingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StatuspageUrlSettingType) IsValid() bool {
	for _, existing := range allowedStatuspageUrlSettingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatuspageUrlSettingType value.
func (v StatuspageUrlSettingType) Ptr() *StatuspageUrlSettingType {
	return &v
}
