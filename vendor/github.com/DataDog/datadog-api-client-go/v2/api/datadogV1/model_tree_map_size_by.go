// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreeMapSizeBy (deprecated) The attribute formerly used to determine size in the widget.
type TreeMapSizeBy string

// List of TreeMapSizeBy.
const (
	TREEMAPSIZEBY_PCT_CPU TreeMapSizeBy = "pct_cpu"
	TREEMAPSIZEBY_PCT_MEM TreeMapSizeBy = "pct_mem"
)

var allowedTreeMapSizeByEnumValues = []TreeMapSizeBy{
	TREEMAPSIZEBY_PCT_CPU,
	TREEMAPSIZEBY_PCT_MEM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreeMapSizeBy) GetAllowedValues() []TreeMapSizeBy {
	return allowedTreeMapSizeByEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreeMapSizeBy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreeMapSizeBy(value)
	return nil
}

// NewTreeMapSizeByFromValue returns a pointer to a valid TreeMapSizeBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreeMapSizeByFromValue(v string) (*TreeMapSizeBy, error) {
	ev := TreeMapSizeBy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreeMapSizeBy: valid values are %v", v, allowedTreeMapSizeByEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreeMapSizeBy) IsValid() bool {
	for _, existing := range allowedTreeMapSizeByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreeMapSizeBy value.
func (v TreeMapSizeBy) Ptr() *TreeMapSizeBy {
	return &v
}
