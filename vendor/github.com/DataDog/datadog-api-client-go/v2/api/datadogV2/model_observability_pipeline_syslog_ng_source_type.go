// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSyslogNgSourceType The source type. The value should always be `syslog_ng`.
type ObservabilityPipelineSyslogNgSourceType string

// List of ObservabilityPipelineSyslogNgSourceType.
const (
	OBSERVABILITYPIPELINESYSLOGNGSOURCETYPE_SYSLOG_NG ObservabilityPipelineSyslogNgSourceType = "syslog_ng"
)

var allowedObservabilityPipelineSyslogNgSourceTypeEnumValues = []ObservabilityPipelineSyslogNgSourceType{
	OBSERVABILITYPIPELINESYSLOGNGSOURCETYPE_SYSLOG_NG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSyslogNgSourceType) GetAllowedValues() []ObservabilityPipelineSyslogNgSourceType {
	return allowedObservabilityPipelineSyslogNgSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSyslogNgSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSyslogNgSourceType(value)
	return nil
}

// NewObservabilityPipelineSyslogNgSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineSyslogNgSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSyslogNgSourceTypeFromValue(v string) (*ObservabilityPipelineSyslogNgSourceType, error) {
	ev := ObservabilityPipelineSyslogNgSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSyslogNgSourceType: valid values are %v", v, allowedObservabilityPipelineSyslogNgSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSyslogNgSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSyslogNgSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSyslogNgSourceType value.
func (v ObservabilityPipelineSyslogNgSourceType) Ptr() *ObservabilityPipelineSyslogNgSourceType {
	return &v
}
