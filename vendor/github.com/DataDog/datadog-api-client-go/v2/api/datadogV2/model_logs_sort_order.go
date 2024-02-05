// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsSortOrder The order to use, ascending or descending
type LogsSortOrder string

// List of LogsSortOrder.
const (
	LOGSSORTORDER_ASCENDING  LogsSortOrder = "asc"
	LOGSSORTORDER_DESCENDING LogsSortOrder = "desc"
)

var allowedLogsSortOrderEnumValues = []LogsSortOrder{
	LOGSSORTORDER_ASCENDING,
	LOGSSORTORDER_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsSortOrder) GetAllowedValues() []LogsSortOrder {
	return allowedLogsSortOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsSortOrder(value)
	return nil
}

// NewLogsSortOrderFromValue returns a pointer to a valid LogsSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsSortOrderFromValue(v string) (*LogsSortOrder, error) {
	ev := LogsSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsSortOrder: valid values are %v", v, allowedLogsSortOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsSortOrder) IsValid() bool {
	for _, existing := range allowedLogsSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsSortOrder value.
func (v LogsSortOrder) Ptr() *LogsSortOrder {
	return &v
}
