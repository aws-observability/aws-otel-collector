// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatadogAPIKeyType The definition of the `DatadogAPIKey` object.
type DatadogAPIKeyType string

// List of DatadogAPIKeyType.
const (
	DATADOGAPIKEYTYPE_DATADOGAPIKEY DatadogAPIKeyType = "DatadogAPIKey"
)

var allowedDatadogAPIKeyTypeEnumValues = []DatadogAPIKeyType{
	DATADOGAPIKEYTYPE_DATADOGAPIKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatadogAPIKeyType) GetAllowedValues() []DatadogAPIKeyType {
	return allowedDatadogAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatadogAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatadogAPIKeyType(value)
	return nil
}

// NewDatadogAPIKeyTypeFromValue returns a pointer to a valid DatadogAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatadogAPIKeyTypeFromValue(v string) (*DatadogAPIKeyType, error) {
	ev := DatadogAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatadogAPIKeyType: valid values are %v", v, allowedDatadogAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatadogAPIKeyType) IsValid() bool {
	for _, existing := range allowedDatadogAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatadogAPIKeyType value.
func (v DatadogAPIKeyType) Ptr() *DatadogAPIKeyType {
	return &v
}
