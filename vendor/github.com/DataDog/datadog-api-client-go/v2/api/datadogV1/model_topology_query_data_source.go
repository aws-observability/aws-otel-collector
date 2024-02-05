// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyQueryDataSource Name of the data source
type TopologyQueryDataSource string

// List of TopologyQueryDataSource.
const (
	TOPOLOGYQUERYDATASOURCE_DATA_STREAMS TopologyQueryDataSource = "data_streams"
	TOPOLOGYQUERYDATASOURCE_SERVICE_MAP  TopologyQueryDataSource = "service_map"
)

var allowedTopologyQueryDataSourceEnumValues = []TopologyQueryDataSource{
	TOPOLOGYQUERYDATASOURCE_DATA_STREAMS,
	TOPOLOGYQUERYDATASOURCE_SERVICE_MAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyQueryDataSource) GetAllowedValues() []TopologyQueryDataSource {
	return allowedTopologyQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyQueryDataSource(value)
	return nil
}

// NewTopologyQueryDataSourceFromValue returns a pointer to a valid TopologyQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyQueryDataSourceFromValue(v string) (*TopologyQueryDataSource, error) {
	ev := TopologyQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyQueryDataSource: valid values are %v", v, allowedTopologyQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyQueryDataSource) IsValid() bool {
	for _, existing := range allowedTopologyQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyQueryDataSource value.
func (v TopologyQueryDataSource) Ptr() *TopologyQueryDataSource {
	return &v
}
