// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatastorePrimaryKeyGenerationStrategy Can be set to `uuid` to automatically generate primary keys when new items are added. Default value is `none`, which requires you to supply a primary key for each new item.
type DatastorePrimaryKeyGenerationStrategy string

// List of DatastorePrimaryKeyGenerationStrategy.
const (
	DATASTOREPRIMARYKEYGENERATIONSTRATEGY_NONE DatastorePrimaryKeyGenerationStrategy = "none"
	DATASTOREPRIMARYKEYGENERATIONSTRATEGY_UUID DatastorePrimaryKeyGenerationStrategy = "uuid"
)

var allowedDatastorePrimaryKeyGenerationStrategyEnumValues = []DatastorePrimaryKeyGenerationStrategy{
	DATASTOREPRIMARYKEYGENERATIONSTRATEGY_NONE,
	DATASTOREPRIMARYKEYGENERATIONSTRATEGY_UUID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatastorePrimaryKeyGenerationStrategy) GetAllowedValues() []DatastorePrimaryKeyGenerationStrategy {
	return allowedDatastorePrimaryKeyGenerationStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatastorePrimaryKeyGenerationStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatastorePrimaryKeyGenerationStrategy(value)
	return nil
}

// NewDatastorePrimaryKeyGenerationStrategyFromValue returns a pointer to a valid DatastorePrimaryKeyGenerationStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatastorePrimaryKeyGenerationStrategyFromValue(v string) (*DatastorePrimaryKeyGenerationStrategy, error) {
	ev := DatastorePrimaryKeyGenerationStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatastorePrimaryKeyGenerationStrategy: valid values are %v", v, allowedDatastorePrimaryKeyGenerationStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatastorePrimaryKeyGenerationStrategy) IsValid() bool {
	for _, existing := range allowedDatastorePrimaryKeyGenerationStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatastorePrimaryKeyGenerationStrategy value.
func (v DatastorePrimaryKeyGenerationStrategy) Ptr() *DatastorePrimaryKeyGenerationStrategy {
	return &v
}
