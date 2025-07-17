// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplunkTcpSourceType The source type. Always `splunk_tcp`.
type ObservabilityPipelineSplunkTcpSourceType string

// List of ObservabilityPipelineSplunkTcpSourceType.
const (
	OBSERVABILITYPIPELINESPLUNKTCPSOURCETYPE_SPLUNK_TCP ObservabilityPipelineSplunkTcpSourceType = "splunk_tcp"
)

var allowedObservabilityPipelineSplunkTcpSourceTypeEnumValues = []ObservabilityPipelineSplunkTcpSourceType{
	OBSERVABILITYPIPELINESPLUNKTCPSOURCETYPE_SPLUNK_TCP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSplunkTcpSourceType) GetAllowedValues() []ObservabilityPipelineSplunkTcpSourceType {
	return allowedObservabilityPipelineSplunkTcpSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSplunkTcpSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSplunkTcpSourceType(value)
	return nil
}

// NewObservabilityPipelineSplunkTcpSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineSplunkTcpSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSplunkTcpSourceTypeFromValue(v string) (*ObservabilityPipelineSplunkTcpSourceType, error) {
	ev := ObservabilityPipelineSplunkTcpSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSplunkTcpSourceType: valid values are %v", v, allowedObservabilityPipelineSplunkTcpSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSplunkTcpSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSplunkTcpSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSplunkTcpSourceType value.
func (v ObservabilityPipelineSplunkTcpSourceType) Ptr() *ObservabilityPipelineSplunkTcpSourceType {
	return &v
}
