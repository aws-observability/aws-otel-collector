// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationForwardDestinationElasticsearchType Type of the Elasticsearch destination.
type CustomDestinationForwardDestinationElasticsearchType string

// List of CustomDestinationForwardDestinationElasticsearchType.
const (
	CUSTOMDESTINATIONFORWARDDESTINATIONELASTICSEARCHTYPE_ELASTICSEARCH CustomDestinationForwardDestinationElasticsearchType = "elasticsearch"
)

var allowedCustomDestinationForwardDestinationElasticsearchTypeEnumValues = []CustomDestinationForwardDestinationElasticsearchType{
	CUSTOMDESTINATIONFORWARDDESTINATIONELASTICSEARCHTYPE_ELASTICSEARCH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationForwardDestinationElasticsearchType) GetAllowedValues() []CustomDestinationForwardDestinationElasticsearchType {
	return allowedCustomDestinationForwardDestinationElasticsearchTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationForwardDestinationElasticsearchType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationForwardDestinationElasticsearchType(value)
	return nil
}

// NewCustomDestinationForwardDestinationElasticsearchTypeFromValue returns a pointer to a valid CustomDestinationForwardDestinationElasticsearchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationForwardDestinationElasticsearchTypeFromValue(v string) (*CustomDestinationForwardDestinationElasticsearchType, error) {
	ev := CustomDestinationForwardDestinationElasticsearchType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationForwardDestinationElasticsearchType: valid values are %v", v, allowedCustomDestinationForwardDestinationElasticsearchTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationForwardDestinationElasticsearchType) IsValid() bool {
	for _, existing := range allowedCustomDestinationForwardDestinationElasticsearchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationForwardDestinationElasticsearchType value.
func (v CustomDestinationForwardDestinationElasticsearchType) Ptr() *CustomDestinationForwardDestinationElasticsearchType {
	return &v
}
