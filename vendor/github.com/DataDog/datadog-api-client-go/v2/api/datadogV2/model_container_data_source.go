// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerDataSource A data source for container-level infrastructure metrics.
type ContainerDataSource string

// List of ContainerDataSource.
const (
	CONTAINERDATASOURCE_CONTAINER ContainerDataSource = "container"
)

var allowedContainerDataSourceEnumValues = []ContainerDataSource{
	CONTAINERDATASOURCE_CONTAINER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ContainerDataSource) GetAllowedValues() []ContainerDataSource {
	return allowedContainerDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ContainerDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ContainerDataSource(value)
	return nil
}

// NewContainerDataSourceFromValue returns a pointer to a valid ContainerDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewContainerDataSourceFromValue(v string) (*ContainerDataSource, error) {
	ev := ContainerDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ContainerDataSource: valid values are %v", v, allowedContainerDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ContainerDataSource) IsValid() bool {
	for _, existing := range allowedContainerDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContainerDataSource value.
func (v ContainerDataSource) Ptr() *ContainerDataSource {
	return &v
}
