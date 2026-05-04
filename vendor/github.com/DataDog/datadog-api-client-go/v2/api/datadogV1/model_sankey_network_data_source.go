// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyNetworkDataSource Network data source type.
type SankeyNetworkDataSource string

// List of SankeyNetworkDataSource.
const (
	SANKEYNETWORKDATASOURCE_NETWORK_DEVICE_FLOWS SankeyNetworkDataSource = "network_device_flows"
	SANKEYNETWORKDATASOURCE_NETWORK              SankeyNetworkDataSource = "network"
)

var allowedSankeyNetworkDataSourceEnumValues = []SankeyNetworkDataSource{
	SANKEYNETWORKDATASOURCE_NETWORK_DEVICE_FLOWS,
	SANKEYNETWORKDATASOURCE_NETWORK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SankeyNetworkDataSource) GetAllowedValues() []SankeyNetworkDataSource {
	return allowedSankeyNetworkDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SankeyNetworkDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SankeyNetworkDataSource(value)
	return nil
}

// NewSankeyNetworkDataSourceFromValue returns a pointer to a valid SankeyNetworkDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSankeyNetworkDataSourceFromValue(v string) (*SankeyNetworkDataSource, error) {
	ev := SankeyNetworkDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SankeyNetworkDataSource: valid values are %v", v, allowedSankeyNetworkDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SankeyNetworkDataSource) IsValid() bool {
	for _, existing := range allowedSankeyNetworkDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SankeyNetworkDataSource value.
func (v SankeyNetworkDataSource) Ptr() *SankeyNetworkDataSource {
	return &v
}
