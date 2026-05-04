// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyNetworkRequestType Type of request for network Sankey widget.
type SankeyNetworkRequestType string

// List of SankeyNetworkRequestType.
const (
	SANKEYNETWORKREQUESTTYPE_NETFLOW_SANKEY SankeyNetworkRequestType = "netflow_sankey"
)

var allowedSankeyNetworkRequestTypeEnumValues = []SankeyNetworkRequestType{
	SANKEYNETWORKREQUESTTYPE_NETFLOW_SANKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SankeyNetworkRequestType) GetAllowedValues() []SankeyNetworkRequestType {
	return allowedSankeyNetworkRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SankeyNetworkRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SankeyNetworkRequestType(value)
	return nil
}

// NewSankeyNetworkRequestTypeFromValue returns a pointer to a valid SankeyNetworkRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSankeyNetworkRequestTypeFromValue(v string) (*SankeyNetworkRequestType, error) {
	ev := SankeyNetworkRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SankeyNetworkRequestType: valid values are %v", v, allowedSankeyNetworkRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SankeyNetworkRequestType) IsValid() bool {
	for _, existing := range allowedSankeyNetworkRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SankeyNetworkRequestType value.
func (v SankeyNetworkRequestType) Ptr() *SankeyNetworkRequestType {
	return &v
}
