// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatadogIntegrationType The definition of the `DatadogIntegrationType` object.
type DatadogIntegrationType string

// List of DatadogIntegrationType.
const (
	DATADOGINTEGRATIONTYPE_DATADOG DatadogIntegrationType = "Datadog"
)

var allowedDatadogIntegrationTypeEnumValues = []DatadogIntegrationType{
	DATADOGINTEGRATIONTYPE_DATADOG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatadogIntegrationType) GetAllowedValues() []DatadogIntegrationType {
	return allowedDatadogIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatadogIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatadogIntegrationType(value)
	return nil
}

// NewDatadogIntegrationTypeFromValue returns a pointer to a valid DatadogIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatadogIntegrationTypeFromValue(v string) (*DatadogIntegrationType, error) {
	ev := DatadogIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatadogIntegrationType: valid values are %v", v, allowedDatadogIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatadogIntegrationType) IsValid() bool {
	for _, existing := range allowedDatadogIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatadogIntegrationType value.
func (v DatadogIntegrationType) Ptr() *DatadogIntegrationType {
	return &v
}
