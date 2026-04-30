// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyNetworkQueryMode Sankey mode for network queries.
type SankeyNetworkQueryMode string

// List of SankeyNetworkQueryMode.
const (
	SANKEYNETWORKQUERYMODE_TARGET SankeyNetworkQueryMode = "target"
)

var allowedSankeyNetworkQueryModeEnumValues = []SankeyNetworkQueryMode{
	SANKEYNETWORKQUERYMODE_TARGET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SankeyNetworkQueryMode) GetAllowedValues() []SankeyNetworkQueryMode {
	return allowedSankeyNetworkQueryModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SankeyNetworkQueryMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SankeyNetworkQueryMode(value)
	return nil
}

// NewSankeyNetworkQueryModeFromValue returns a pointer to a valid SankeyNetworkQueryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSankeyNetworkQueryModeFromValue(v string) (*SankeyNetworkQueryMode, error) {
	ev := SankeyNetworkQueryMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SankeyNetworkQueryMode: valid values are %v", v, allowedSankeyNetworkQueryModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SankeyNetworkQueryMode) IsValid() bool {
	for _, existing := range allowedSankeyNetworkQueryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SankeyNetworkQueryMode value.
func (v SankeyNetworkQueryMode) Ptr() *SankeyNetworkQueryMode {
	return &v
}
