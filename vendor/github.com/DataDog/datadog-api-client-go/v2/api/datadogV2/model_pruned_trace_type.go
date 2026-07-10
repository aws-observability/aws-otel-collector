// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PrunedTraceType The type of the pruned trace resource. The value is always `pruned_trace`.
type PrunedTraceType string

// List of PrunedTraceType.
const (
	PRUNEDTRACETYPE_PRUNED_TRACE PrunedTraceType = "pruned_trace"
)

var allowedPrunedTraceTypeEnumValues = []PrunedTraceType{
	PRUNEDTRACETYPE_PRUNED_TRACE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PrunedTraceType) GetAllowedValues() []PrunedTraceType {
	return allowedPrunedTraceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PrunedTraceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PrunedTraceType(value)
	return nil
}

// NewPrunedTraceTypeFromValue returns a pointer to a valid PrunedTraceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPrunedTraceTypeFromValue(v string) (*PrunedTraceType, error) {
	ev := PrunedTraceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PrunedTraceType: valid values are %v", v, allowedPrunedTraceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PrunedTraceType) IsValid() bool {
	for _, existing := range allowedPrunedTraceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PrunedTraceType value.
func (v PrunedTraceType) Ptr() *PrunedTraceType {
	return &v
}
