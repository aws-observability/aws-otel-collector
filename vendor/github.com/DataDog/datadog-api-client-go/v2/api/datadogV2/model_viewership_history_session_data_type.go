// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ViewershipHistorySessionDataType Rum replay session resource type.
type ViewershipHistorySessionDataType string

// List of ViewershipHistorySessionDataType.
const (
	VIEWERSHIPHISTORYSESSIONDATATYPE_RUM_REPLAY_SESSION ViewershipHistorySessionDataType = "rum_replay_session"
)

var allowedViewershipHistorySessionDataTypeEnumValues = []ViewershipHistorySessionDataType{
	VIEWERSHIPHISTORYSESSIONDATATYPE_RUM_REPLAY_SESSION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ViewershipHistorySessionDataType) GetAllowedValues() []ViewershipHistorySessionDataType {
	return allowedViewershipHistorySessionDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ViewershipHistorySessionDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ViewershipHistorySessionDataType(value)
	return nil
}

// NewViewershipHistorySessionDataTypeFromValue returns a pointer to a valid ViewershipHistorySessionDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewViewershipHistorySessionDataTypeFromValue(v string) (*ViewershipHistorySessionDataType, error) {
	ev := ViewershipHistorySessionDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ViewershipHistorySessionDataType: valid values are %v", v, allowedViewershipHistorySessionDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ViewershipHistorySessionDataType) IsValid() bool {
	for _, existing := range allowedViewershipHistorySessionDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ViewershipHistorySessionDataType value.
func (v ViewershipHistorySessionDataType) Ptr() *ViewershipHistorySessionDataType {
	return &v
}
