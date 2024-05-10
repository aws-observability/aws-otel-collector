// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseForwardDestinationElasticsearchType Type of the Elasticsearch destination.
type CustomDestinationResponseForwardDestinationElasticsearchType string

// List of CustomDestinationResponseForwardDestinationElasticsearchType.
const (
	CUSTOMDESTINATIONRESPONSEFORWARDDESTINATIONELASTICSEARCHTYPE_ELASTICSEARCH CustomDestinationResponseForwardDestinationElasticsearchType = "elasticsearch"
)

var allowedCustomDestinationResponseForwardDestinationElasticsearchTypeEnumValues = []CustomDestinationResponseForwardDestinationElasticsearchType{
	CUSTOMDESTINATIONRESPONSEFORWARDDESTINATIONELASTICSEARCHTYPE_ELASTICSEARCH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationResponseForwardDestinationElasticsearchType) GetAllowedValues() []CustomDestinationResponseForwardDestinationElasticsearchType {
	return allowedCustomDestinationResponseForwardDestinationElasticsearchTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationResponseForwardDestinationElasticsearchType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationResponseForwardDestinationElasticsearchType(value)
	return nil
}

// NewCustomDestinationResponseForwardDestinationElasticsearchTypeFromValue returns a pointer to a valid CustomDestinationResponseForwardDestinationElasticsearchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationResponseForwardDestinationElasticsearchTypeFromValue(v string) (*CustomDestinationResponseForwardDestinationElasticsearchType, error) {
	ev := CustomDestinationResponseForwardDestinationElasticsearchType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationResponseForwardDestinationElasticsearchType: valid values are %v", v, allowedCustomDestinationResponseForwardDestinationElasticsearchTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationResponseForwardDestinationElasticsearchType) IsValid() bool {
	for _, existing := range allowedCustomDestinationResponseForwardDestinationElasticsearchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationResponseForwardDestinationElasticsearchType value.
func (v CustomDestinationResponseForwardDestinationElasticsearchType) Ptr() *CustomDestinationResponseForwardDestinationElasticsearchType {
	return &v
}
