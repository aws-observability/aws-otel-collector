// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataQualityModelTypeOverride Override for the model type used in anomaly detection.
type MonitorFormulaAndFunctionDataQualityModelTypeOverride string

// List of MonitorFormulaAndFunctionDataQualityModelTypeOverride.
const (
	MONITORFORMULAANDFUNCTIONDATAQUALITYMODELTYPEOVERRIDE_FRESHNESS  MonitorFormulaAndFunctionDataQualityModelTypeOverride = "freshness"
	MONITORFORMULAANDFUNCTIONDATAQUALITYMODELTYPEOVERRIDE_PERCENTAGE MonitorFormulaAndFunctionDataQualityModelTypeOverride = "percentage"
	MONITORFORMULAANDFUNCTIONDATAQUALITYMODELTYPEOVERRIDE_ANY        MonitorFormulaAndFunctionDataQualityModelTypeOverride = "any"
)

var allowedMonitorFormulaAndFunctionDataQualityModelTypeOverrideEnumValues = []MonitorFormulaAndFunctionDataQualityModelTypeOverride{
	MONITORFORMULAANDFUNCTIONDATAQUALITYMODELTYPEOVERRIDE_FRESHNESS,
	MONITORFORMULAANDFUNCTIONDATAQUALITYMODELTYPEOVERRIDE_PERCENTAGE,
	MONITORFORMULAANDFUNCTIONDATAQUALITYMODELTYPEOVERRIDE_ANY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionDataQualityModelTypeOverride) GetAllowedValues() []MonitorFormulaAndFunctionDataQualityModelTypeOverride {
	return allowedMonitorFormulaAndFunctionDataQualityModelTypeOverrideEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionDataQualityModelTypeOverride) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionDataQualityModelTypeOverride(value)
	return nil
}

// NewMonitorFormulaAndFunctionDataQualityModelTypeOverrideFromValue returns a pointer to a valid MonitorFormulaAndFunctionDataQualityModelTypeOverride
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionDataQualityModelTypeOverrideFromValue(v string) (*MonitorFormulaAndFunctionDataQualityModelTypeOverride, error) {
	ev := MonitorFormulaAndFunctionDataQualityModelTypeOverride(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionDataQualityModelTypeOverride: valid values are %v", v, allowedMonitorFormulaAndFunctionDataQualityModelTypeOverrideEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionDataQualityModelTypeOverride) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionDataQualityModelTypeOverrideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionDataQualityModelTypeOverride value.
func (v MonitorFormulaAndFunctionDataQualityModelTypeOverride) Ptr() *MonitorFormulaAndFunctionDataQualityModelTypeOverride {
	return &v
}
