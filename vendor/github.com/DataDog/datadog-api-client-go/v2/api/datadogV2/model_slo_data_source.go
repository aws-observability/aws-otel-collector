// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloDataSource A data source for SLO queries.
type SloDataSource string

// List of SloDataSource.
const (
	SLODATASOURCE_SLO SloDataSource = "slo"
)

var allowedSloDataSourceEnumValues = []SloDataSource{
	SLODATASOURCE_SLO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SloDataSource) GetAllowedValues() []SloDataSource {
	return allowedSloDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SloDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SloDataSource(value)
	return nil
}

// NewSloDataSourceFromValue returns a pointer to a valid SloDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSloDataSourceFromValue(v string) (*SloDataSource, error) {
	ev := SloDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SloDataSource: valid values are %v", v, allowedSloDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SloDataSource) IsValid() bool {
	for _, existing := range allowedSloDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SloDataSource value.
func (v SloDataSource) Ptr() *SloDataSource {
	return &v
}
