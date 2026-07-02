// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// APMSpanErrorFlag Error flag for a span. `1` when the span is in error, `0` otherwise.
type APMSpanErrorFlag int32

// List of APMSpanErrorFlag.
const (
	APMSPANERRORFLAG_NO_ERROR APMSpanErrorFlag = 0
	APMSPANERRORFLAG_ERROR    APMSpanErrorFlag = 1
)

var allowedAPMSpanErrorFlagEnumValues = []APMSpanErrorFlag{
	APMSPANERRORFLAG_NO_ERROR,
	APMSPANERRORFLAG_ERROR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *APMSpanErrorFlag) GetAllowedValues() []APMSpanErrorFlag {
	return allowedAPMSpanErrorFlagEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *APMSpanErrorFlag) UnmarshalJSON(src []byte) error {
	var value int32
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = APMSpanErrorFlag(value)
	return nil
}

// NewAPMSpanErrorFlagFromValue returns a pointer to a valid APMSpanErrorFlag
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAPMSpanErrorFlagFromValue(v int32) (*APMSpanErrorFlag, error) {
	ev := APMSpanErrorFlag(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for APMSpanErrorFlag: valid values are %v", v, allowedAPMSpanErrorFlagEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v APMSpanErrorFlag) IsValid() bool {
	for _, existing := range allowedAPMSpanErrorFlagEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to APMSpanErrorFlag value.
func (v APMSpanErrorFlag) Ptr() *APMSpanErrorFlag {
	return &v
}
