// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSyslogNgDestinationType The destination type. The value should always be `syslog_ng`.
type ObservabilityPipelineSyslogNgDestinationType string

// List of ObservabilityPipelineSyslogNgDestinationType.
const (
	OBSERVABILITYPIPELINESYSLOGNGDESTINATIONTYPE_SYSLOG_NG ObservabilityPipelineSyslogNgDestinationType = "syslog_ng"
)

var allowedObservabilityPipelineSyslogNgDestinationTypeEnumValues = []ObservabilityPipelineSyslogNgDestinationType{
	OBSERVABILITYPIPELINESYSLOGNGDESTINATIONTYPE_SYSLOG_NG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSyslogNgDestinationType) GetAllowedValues() []ObservabilityPipelineSyslogNgDestinationType {
	return allowedObservabilityPipelineSyslogNgDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSyslogNgDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSyslogNgDestinationType(value)
	return nil
}

// NewObservabilityPipelineSyslogNgDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineSyslogNgDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSyslogNgDestinationTypeFromValue(v string) (*ObservabilityPipelineSyslogNgDestinationType, error) {
	ev := ObservabilityPipelineSyslogNgDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSyslogNgDestinationType: valid values are %v", v, allowedObservabilityPipelineSyslogNgDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSyslogNgDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSyslogNgDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSyslogNgDestinationType value.
func (v ObservabilityPipelineSyslogNgDestinationType) Ptr() *ObservabilityPipelineSyslogNgDestinationType {
	return &v
}
