// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueUpdateStateRequestDataType Type of the object.
type IssueUpdateStateRequestDataType string

// List of IssueUpdateStateRequestDataType.
const (
	ISSUEUPDATESTATEREQUESTDATATYPE_ERROR_TRACKING_ISSUE IssueUpdateStateRequestDataType = "error_tracking_issue"
)

var allowedIssueUpdateStateRequestDataTypeEnumValues = []IssueUpdateStateRequestDataType{
	ISSUEUPDATESTATEREQUESTDATATYPE_ERROR_TRACKING_ISSUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueUpdateStateRequestDataType) GetAllowedValues() []IssueUpdateStateRequestDataType {
	return allowedIssueUpdateStateRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueUpdateStateRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueUpdateStateRequestDataType(value)
	return nil
}

// NewIssueUpdateStateRequestDataTypeFromValue returns a pointer to a valid IssueUpdateStateRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueUpdateStateRequestDataTypeFromValue(v string) (*IssueUpdateStateRequestDataType, error) {
	ev := IssueUpdateStateRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueUpdateStateRequestDataType: valid values are %v", v, allowedIssueUpdateStateRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueUpdateStateRequestDataType) IsValid() bool {
	for _, existing := range allowedIssueUpdateStateRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueUpdateStateRequestDataType value.
func (v IssueUpdateStateRequestDataType) Ptr() *IssueUpdateStateRequestDataType {
	return &v
}
