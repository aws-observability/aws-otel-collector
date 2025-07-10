// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineRsyslogDestinationType The destination type. The value should always be `rsyslog`.
type ObservabilityPipelineRsyslogDestinationType string

// List of ObservabilityPipelineRsyslogDestinationType.
const (
	OBSERVABILITYPIPELINERSYSLOGDESTINATIONTYPE_RSYSLOG ObservabilityPipelineRsyslogDestinationType = "rsyslog"
)

var allowedObservabilityPipelineRsyslogDestinationTypeEnumValues = []ObservabilityPipelineRsyslogDestinationType{
	OBSERVABILITYPIPELINERSYSLOGDESTINATIONTYPE_RSYSLOG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineRsyslogDestinationType) GetAllowedValues() []ObservabilityPipelineRsyslogDestinationType {
	return allowedObservabilityPipelineRsyslogDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineRsyslogDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineRsyslogDestinationType(value)
	return nil
}

// NewObservabilityPipelineRsyslogDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineRsyslogDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineRsyslogDestinationTypeFromValue(v string) (*ObservabilityPipelineRsyslogDestinationType, error) {
	ev := ObservabilityPipelineRsyslogDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineRsyslogDestinationType: valid values are %v", v, allowedObservabilityPipelineRsyslogDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineRsyslogDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineRsyslogDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineRsyslogDestinationType value.
func (v ObservabilityPipelineRsyslogDestinationType) Ptr() *ObservabilityPipelineRsyslogDestinationType {
	return &v
}
