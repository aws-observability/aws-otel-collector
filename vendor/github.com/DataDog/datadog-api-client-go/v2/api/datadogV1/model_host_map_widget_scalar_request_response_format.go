// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetScalarRequestResponseFormat Response format for the scalar formula request. Only `scalar` is supported.
type HostMapWidgetScalarRequestResponseFormat string

// List of HostMapWidgetScalarRequestResponseFormat.
const (
	HOSTMAPWIDGETSCALARREQUESTRESPONSEFORMAT_SCALAR HostMapWidgetScalarRequestResponseFormat = "scalar"
)

var allowedHostMapWidgetScalarRequestResponseFormatEnumValues = []HostMapWidgetScalarRequestResponseFormat{
	HOSTMAPWIDGETSCALARREQUESTRESPONSEFORMAT_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostMapWidgetScalarRequestResponseFormat) GetAllowedValues() []HostMapWidgetScalarRequestResponseFormat {
	return allowedHostMapWidgetScalarRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostMapWidgetScalarRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostMapWidgetScalarRequestResponseFormat(value)
	return nil
}

// NewHostMapWidgetScalarRequestResponseFormatFromValue returns a pointer to a valid HostMapWidgetScalarRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostMapWidgetScalarRequestResponseFormatFromValue(v string) (*HostMapWidgetScalarRequestResponseFormat, error) {
	ev := HostMapWidgetScalarRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostMapWidgetScalarRequestResponseFormat: valid values are %v", v, allowedHostMapWidgetScalarRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostMapWidgetScalarRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedHostMapWidgetScalarRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostMapWidgetScalarRequestResponseFormat value.
func (v HostMapWidgetScalarRequestResponseFormat) Ptr() *HostMapWidgetScalarRequestResponseFormat {
	return &v
}
