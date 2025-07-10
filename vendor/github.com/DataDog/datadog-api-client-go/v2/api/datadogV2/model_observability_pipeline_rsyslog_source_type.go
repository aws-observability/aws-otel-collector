// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineRsyslogSourceType The source type. The value should always be `rsyslog`.
type ObservabilityPipelineRsyslogSourceType string

// List of ObservabilityPipelineRsyslogSourceType.
const (
	OBSERVABILITYPIPELINERSYSLOGSOURCETYPE_RSYSLOG ObservabilityPipelineRsyslogSourceType = "rsyslog"
)

var allowedObservabilityPipelineRsyslogSourceTypeEnumValues = []ObservabilityPipelineRsyslogSourceType{
	OBSERVABILITYPIPELINERSYSLOGSOURCETYPE_RSYSLOG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineRsyslogSourceType) GetAllowedValues() []ObservabilityPipelineRsyslogSourceType {
	return allowedObservabilityPipelineRsyslogSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineRsyslogSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineRsyslogSourceType(value)
	return nil
}

// NewObservabilityPipelineRsyslogSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineRsyslogSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineRsyslogSourceTypeFromValue(v string) (*ObservabilityPipelineRsyslogSourceType, error) {
	ev := ObservabilityPipelineRsyslogSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineRsyslogSourceType: valid values are %v", v, allowedObservabilityPipelineRsyslogSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineRsyslogSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineRsyslogSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineRsyslogSourceType value.
func (v ObservabilityPipelineRsyslogSourceType) Ptr() *ObservabilityPipelineRsyslogSourceType {
	return &v
}
