// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimelineCellMarkdownContentType Type of the Markdown timeline cell.
type IncidentTimelineCellMarkdownContentType string

// List of IncidentTimelineCellMarkdownContentType.
const (
	INCIDENTTIMELINECELLMARKDOWNCONTENTTYPE_MARKDOWN IncidentTimelineCellMarkdownContentType = "markdown"
)

var allowedIncidentTimelineCellMarkdownContentTypeEnumValues = []IncidentTimelineCellMarkdownContentType{
	INCIDENTTIMELINECELLMARKDOWNCONTENTTYPE_MARKDOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentTimelineCellMarkdownContentType) GetAllowedValues() []IncidentTimelineCellMarkdownContentType {
	return allowedIncidentTimelineCellMarkdownContentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentTimelineCellMarkdownContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentTimelineCellMarkdownContentType(value)
	return nil
}

// NewIncidentTimelineCellMarkdownContentTypeFromValue returns a pointer to a valid IncidentTimelineCellMarkdownContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTimelineCellMarkdownContentTypeFromValue(v string) (*IncidentTimelineCellMarkdownContentType, error) {
	ev := IncidentTimelineCellMarkdownContentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentTimelineCellMarkdownContentType: valid values are %v", v, allowedIncidentTimelineCellMarkdownContentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentTimelineCellMarkdownContentType) IsValid() bool {
	for _, existing := range allowedIncidentTimelineCellMarkdownContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentTimelineCellMarkdownContentType value.
func (v IncidentTimelineCellMarkdownContentType) Ptr() *IncidentTimelineCellMarkdownContentType {
	return &v
}
