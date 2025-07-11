// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PageUrgency On-Call Page urgency level.
type PageUrgency string

// List of PageUrgency.
const (
	PAGEURGENCY_LOW  PageUrgency = "low"
	PAGEURGENCY_HIGH PageUrgency = "high"
)

var allowedPageUrgencyEnumValues = []PageUrgency{
	PAGEURGENCY_LOW,
	PAGEURGENCY_HIGH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PageUrgency) GetAllowedValues() []PageUrgency {
	return allowedPageUrgencyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PageUrgency) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PageUrgency(value)
	return nil
}

// NewPageUrgencyFromValue returns a pointer to a valid PageUrgency
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPageUrgencyFromValue(v string) (*PageUrgency, error) {
	ev := PageUrgency(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PageUrgency: valid values are %v", v, allowedPageUrgencyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PageUrgency) IsValid() bool {
	for _, existing := range allowedPageUrgencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PageUrgency value.
func (v PageUrgency) Ptr() *PageUrgency {
	return &v
}
