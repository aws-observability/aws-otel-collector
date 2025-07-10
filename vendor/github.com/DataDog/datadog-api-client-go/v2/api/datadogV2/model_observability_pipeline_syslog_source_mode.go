// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSyslogSourceMode Protocol used by the syslog source to receive messages.
type ObservabilityPipelineSyslogSourceMode string

// List of ObservabilityPipelineSyslogSourceMode.
const (
	OBSERVABILITYPIPELINESYSLOGSOURCEMODE_TCP ObservabilityPipelineSyslogSourceMode = "tcp"
	OBSERVABILITYPIPELINESYSLOGSOURCEMODE_UDP ObservabilityPipelineSyslogSourceMode = "udp"
)

var allowedObservabilityPipelineSyslogSourceModeEnumValues = []ObservabilityPipelineSyslogSourceMode{
	OBSERVABILITYPIPELINESYSLOGSOURCEMODE_TCP,
	OBSERVABILITYPIPELINESYSLOGSOURCEMODE_UDP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSyslogSourceMode) GetAllowedValues() []ObservabilityPipelineSyslogSourceMode {
	return allowedObservabilityPipelineSyslogSourceModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSyslogSourceMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSyslogSourceMode(value)
	return nil
}

// NewObservabilityPipelineSyslogSourceModeFromValue returns a pointer to a valid ObservabilityPipelineSyslogSourceMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSyslogSourceModeFromValue(v string) (*ObservabilityPipelineSyslogSourceMode, error) {
	ev := ObservabilityPipelineSyslogSourceMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSyslogSourceMode: valid values are %v", v, allowedObservabilityPipelineSyslogSourceModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSyslogSourceMode) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSyslogSourceModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSyslogSourceMode value.
func (v ObservabilityPipelineSyslogSourceMode) Ptr() *ObservabilityPipelineSyslogSourceMode {
	return &v
}
