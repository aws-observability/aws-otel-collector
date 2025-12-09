// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssuesSearchRequestDataAttributesTrack Track of the events to query. Either track(s) or persona(s) must be specified.
type IssuesSearchRequestDataAttributesTrack string

// List of IssuesSearchRequestDataAttributesTrack.
const (
	ISSUESSEARCHREQUESTDATAATTRIBUTESTRACK_TRACE IssuesSearchRequestDataAttributesTrack = "trace"
	ISSUESSEARCHREQUESTDATAATTRIBUTESTRACK_LOGS  IssuesSearchRequestDataAttributesTrack = "logs"
	ISSUESSEARCHREQUESTDATAATTRIBUTESTRACK_RUM   IssuesSearchRequestDataAttributesTrack = "rum"
)

var allowedIssuesSearchRequestDataAttributesTrackEnumValues = []IssuesSearchRequestDataAttributesTrack{
	ISSUESSEARCHREQUESTDATAATTRIBUTESTRACK_TRACE,
	ISSUESSEARCHREQUESTDATAATTRIBUTESTRACK_LOGS,
	ISSUESSEARCHREQUESTDATAATTRIBUTESTRACK_RUM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssuesSearchRequestDataAttributesTrack) GetAllowedValues() []IssuesSearchRequestDataAttributesTrack {
	return allowedIssuesSearchRequestDataAttributesTrackEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssuesSearchRequestDataAttributesTrack) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssuesSearchRequestDataAttributesTrack(value)
	return nil
}

// NewIssuesSearchRequestDataAttributesTrackFromValue returns a pointer to a valid IssuesSearchRequestDataAttributesTrack
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssuesSearchRequestDataAttributesTrackFromValue(v string) (*IssuesSearchRequestDataAttributesTrack, error) {
	ev := IssuesSearchRequestDataAttributesTrack(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssuesSearchRequestDataAttributesTrack: valid values are %v", v, allowedIssuesSearchRequestDataAttributesTrackEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssuesSearchRequestDataAttributesTrack) IsValid() bool {
	for _, existing := range allowedIssuesSearchRequestDataAttributesTrackEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssuesSearchRequestDataAttributesTrack value.
func (v IssuesSearchRequestDataAttributesTrack) Ptr() *IssuesSearchRequestDataAttributesTrack {
	return &v
}
