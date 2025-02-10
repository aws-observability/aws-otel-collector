// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrderDirection The sort direction for results.
type OrderDirection string

// List of OrderDirection.
const (
	ORDERDIRECTION_ASC  OrderDirection = "asc"
	ORDERDIRECTION_DESC OrderDirection = "desc"
)

var allowedOrderDirectionEnumValues = []OrderDirection{
	ORDERDIRECTION_ASC,
	ORDERDIRECTION_DESC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrderDirection) GetAllowedValues() []OrderDirection {
	return allowedOrderDirectionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrderDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrderDirection(value)
	return nil
}

// NewOrderDirectionFromValue returns a pointer to a valid OrderDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrderDirectionFromValue(v string) (*OrderDirection, error) {
	ev := OrderDirection(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrderDirection: valid values are %v", v, allowedOrderDirectionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrderDirection) IsValid() bool {
	for _, existing := range allowedOrderDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderDirection value.
func (v OrderDirection) Ptr() *OrderDirection {
	return &v
}
