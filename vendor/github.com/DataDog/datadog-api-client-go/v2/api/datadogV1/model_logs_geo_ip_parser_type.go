// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsGeoIPParserType Type of GeoIP parser.
type LogsGeoIPParserType string

// List of LogsGeoIPParserType.
const (
	LOGSGEOIPPARSERTYPE_GEO_IP_PARSER LogsGeoIPParserType = "geo-ip-parser"
)

var allowedLogsGeoIPParserTypeEnumValues = []LogsGeoIPParserType{
	LOGSGEOIPPARSERTYPE_GEO_IP_PARSER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsGeoIPParserType) GetAllowedValues() []LogsGeoIPParserType {
	return allowedLogsGeoIPParserTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsGeoIPParserType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsGeoIPParserType(value)
	return nil
}

// NewLogsGeoIPParserTypeFromValue returns a pointer to a valid LogsGeoIPParserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsGeoIPParserTypeFromValue(v string) (*LogsGeoIPParserType, error) {
	ev := LogsGeoIPParserType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsGeoIPParserType: valid values are %v", v, allowedLogsGeoIPParserTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsGeoIPParserType) IsValid() bool {
	for _, existing := range allowedLogsGeoIPParserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsGeoIPParserType value.
func (v LogsGeoIPParserType) Ptr() *LogsGeoIPParserType {
	return &v
}
