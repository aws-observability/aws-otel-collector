// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotionIntegrationType The definition of the `NotionIntegrationType` object.
type NotionIntegrationType string

// List of NotionIntegrationType.
const (
	NOTIONINTEGRATIONTYPE_NOTION NotionIntegrationType = "Notion"
)

var allowedNotionIntegrationTypeEnumValues = []NotionIntegrationType{
	NOTIONINTEGRATIONTYPE_NOTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotionIntegrationType) GetAllowedValues() []NotionIntegrationType {
	return allowedNotionIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotionIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotionIntegrationType(value)
	return nil
}

// NewNotionIntegrationTypeFromValue returns a pointer to a valid NotionIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotionIntegrationTypeFromValue(v string) (*NotionIntegrationType, error) {
	ev := NotionIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotionIntegrationType: valid values are %v", v, allowedNotionIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotionIntegrationType) IsValid() bool {
	for _, existing := range allowedNotionIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotionIntegrationType value.
func (v NotionIntegrationType) Ptr() *NotionIntegrationType {
	return &v
}
