// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmDependencyStatName The APM dependency statistic to query.
type ApmDependencyStatName string

// List of ApmDependencyStatName.
const (
	APMDEPENDENCYSTATNAME_AVG_DURATION        ApmDependencyStatName = "avg_duration"
	APMDEPENDENCYSTATNAME_AVG_ROOT_DURATION   ApmDependencyStatName = "avg_root_duration"
	APMDEPENDENCYSTATNAME_AVG_SPANS_PER_TRACE ApmDependencyStatName = "avg_spans_per_trace"
	APMDEPENDENCYSTATNAME_ERROR_RATE          ApmDependencyStatName = "error_rate"
	APMDEPENDENCYSTATNAME_PCT_EXEC_TIME       ApmDependencyStatName = "pct_exec_time"
	APMDEPENDENCYSTATNAME_PCT_OF_TRACES       ApmDependencyStatName = "pct_of_traces"
	APMDEPENDENCYSTATNAME_TOTAL_TRACES_COUNT  ApmDependencyStatName = "total_traces_count"
)

var allowedApmDependencyStatNameEnumValues = []ApmDependencyStatName{
	APMDEPENDENCYSTATNAME_AVG_DURATION,
	APMDEPENDENCYSTATNAME_AVG_ROOT_DURATION,
	APMDEPENDENCYSTATNAME_AVG_SPANS_PER_TRACE,
	APMDEPENDENCYSTATNAME_ERROR_RATE,
	APMDEPENDENCYSTATNAME_PCT_EXEC_TIME,
	APMDEPENDENCYSTATNAME_PCT_OF_TRACES,
	APMDEPENDENCYSTATNAME_TOTAL_TRACES_COUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmDependencyStatName) GetAllowedValues() []ApmDependencyStatName {
	return allowedApmDependencyStatNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmDependencyStatName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmDependencyStatName(value)
	return nil
}

// NewApmDependencyStatNameFromValue returns a pointer to a valid ApmDependencyStatName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmDependencyStatNameFromValue(v string) (*ApmDependencyStatName, error) {
	ev := ApmDependencyStatName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmDependencyStatName: valid values are %v", v, allowedApmDependencyStatNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmDependencyStatName) IsValid() bool {
	for _, existing := range allowedApmDependencyStatNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmDependencyStatName value.
func (v ApmDependencyStatName) Ptr() *ApmDependencyStatName {
	return &v
}
