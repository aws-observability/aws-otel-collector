// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyRumDataSource Sankey widget with RUM data source.
type SankeyRumDataSource string

// List of SankeyRumDataSource.
const (
	SANKEYRUMDATASOURCE_RUM               SankeyRumDataSource = "rum"
	SANKEYRUMDATASOURCE_PRODUCT_ANALYTICS SankeyRumDataSource = "product_analytics"
)

var allowedSankeyRumDataSourceEnumValues = []SankeyRumDataSource{
	SANKEYRUMDATASOURCE_RUM,
	SANKEYRUMDATASOURCE_PRODUCT_ANALYTICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SankeyRumDataSource) GetAllowedValues() []SankeyRumDataSource {
	return allowedSankeyRumDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SankeyRumDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SankeyRumDataSource(value)
	return nil
}

// NewSankeyRumDataSourceFromValue returns a pointer to a valid SankeyRumDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSankeyRumDataSourceFromValue(v string) (*SankeyRumDataSource, error) {
	ev := SankeyRumDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SankeyRumDataSource: valid values are %v", v, allowedSankeyRumDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SankeyRumDataSource) IsValid() bool {
	for _, existing := range allowedSankeyRumDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SankeyRumDataSource value.
func (v SankeyRumDataSource) Ptr() *SankeyRumDataSource {
	return &v
}
