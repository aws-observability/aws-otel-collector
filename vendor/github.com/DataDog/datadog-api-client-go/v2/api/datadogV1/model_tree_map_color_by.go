// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreeMapColorBy (deprecated) The attribute formerly used to determine color in the widget.
type TreeMapColorBy string

// List of TreeMapColorBy.
const (
	TREEMAPCOLORBY_USER TreeMapColorBy = "user"
)

var allowedTreeMapColorByEnumValues = []TreeMapColorBy{
	TREEMAPCOLORBY_USER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreeMapColorBy) GetAllowedValues() []TreeMapColorBy {
	return allowedTreeMapColorByEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreeMapColorBy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreeMapColorBy(value)
	return nil
}

// NewTreeMapColorByFromValue returns a pointer to a valid TreeMapColorBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreeMapColorByFromValue(v string) (*TreeMapColorBy, error) {
	ev := TreeMapColorBy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreeMapColorBy: valid values are %v", v, allowedTreeMapColorByEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreeMapColorBy) IsValid() bool {
	for _, existing := range allowedTreeMapColorByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreeMapColorBy value.
func (v TreeMapColorBy) Ptr() *TreeMapColorBy {
	return &v
}
