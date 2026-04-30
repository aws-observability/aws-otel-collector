// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmMetricsSpanKind Describes the relationship between the span, its parents, and its children in a trace.
type ApmMetricsSpanKind string

// List of ApmMetricsSpanKind.
const (
	APMMETRICSSPANKIND_CONSUMER ApmMetricsSpanKind = "consumer"
	APMMETRICSSPANKIND_SERVER   ApmMetricsSpanKind = "server"
	APMMETRICSSPANKIND_CLIENT   ApmMetricsSpanKind = "client"
	APMMETRICSSPANKIND_PRODUCER ApmMetricsSpanKind = "producer"
	APMMETRICSSPANKIND_INTERNAL ApmMetricsSpanKind = "internal"
)

var allowedApmMetricsSpanKindEnumValues = []ApmMetricsSpanKind{
	APMMETRICSSPANKIND_CONSUMER,
	APMMETRICSSPANKIND_SERVER,
	APMMETRICSSPANKIND_CLIENT,
	APMMETRICSSPANKIND_PRODUCER,
	APMMETRICSSPANKIND_INTERNAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmMetricsSpanKind) GetAllowedValues() []ApmMetricsSpanKind {
	return allowedApmMetricsSpanKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmMetricsSpanKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmMetricsSpanKind(value)
	return nil
}

// NewApmMetricsSpanKindFromValue returns a pointer to a valid ApmMetricsSpanKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmMetricsSpanKindFromValue(v string) (*ApmMetricsSpanKind, error) {
	ev := ApmMetricsSpanKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmMetricsSpanKind: valid values are %v", v, allowedApmMetricsSpanKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmMetricsSpanKind) IsValid() bool {
	for _, existing := range allowedApmMetricsSpanKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmMetricsSpanKind value.
func (v ApmMetricsSpanKind) Ptr() *ApmMetricsSpanKind {
	return &v
}
