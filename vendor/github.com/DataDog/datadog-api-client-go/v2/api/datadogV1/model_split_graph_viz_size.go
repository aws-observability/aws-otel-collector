// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitGraphVizSize Size of the individual graphs in the split.
type SplitGraphVizSize string

// List of SplitGraphVizSize.
const (
	SPLITGRAPHVIZSIZE_XS SplitGraphVizSize = "xs"
	SPLITGRAPHVIZSIZE_SM SplitGraphVizSize = "sm"
	SPLITGRAPHVIZSIZE_MD SplitGraphVizSize = "md"
	SPLITGRAPHVIZSIZE_LG SplitGraphVizSize = "lg"
)

var allowedSplitGraphVizSizeEnumValues = []SplitGraphVizSize{
	SPLITGRAPHVIZSIZE_XS,
	SPLITGRAPHVIZSIZE_SM,
	SPLITGRAPHVIZSIZE_MD,
	SPLITGRAPHVIZSIZE_LG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SplitGraphVizSize) GetAllowedValues() []SplitGraphVizSize {
	return allowedSplitGraphVizSizeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SplitGraphVizSize) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SplitGraphVizSize(value)
	return nil
}

// NewSplitGraphVizSizeFromValue returns a pointer to a valid SplitGraphVizSize
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSplitGraphVizSizeFromValue(v string) (*SplitGraphVizSize, error) {
	ev := SplitGraphVizSize(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SplitGraphVizSize: valid values are %v", v, allowedSplitGraphVizSizeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SplitGraphVizSize) IsValid() bool {
	for _, existing := range allowedSplitGraphVizSizeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SplitGraphVizSize value.
func (v SplitGraphVizSize) Ptr() *SplitGraphVizSize {
	return &v
}
