// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SalesforceIncidentsTemplateType Salesforce incident template resource type.
type SalesforceIncidentsTemplateType string

// List of SalesforceIncidentsTemplateType.
const (
	SALESFORCEINCIDENTSTEMPLATETYPE_SALESFORCE_INCIDENTS_INCIDENT_TEMPLATE SalesforceIncidentsTemplateType = "salesforce-incidents-incident-template"
)

var allowedSalesforceIncidentsTemplateTypeEnumValues = []SalesforceIncidentsTemplateType{
	SALESFORCEINCIDENTSTEMPLATETYPE_SALESFORCE_INCIDENTS_INCIDENT_TEMPLATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SalesforceIncidentsTemplateType) GetAllowedValues() []SalesforceIncidentsTemplateType {
	return allowedSalesforceIncidentsTemplateTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SalesforceIncidentsTemplateType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SalesforceIncidentsTemplateType(value)
	return nil
}

// NewSalesforceIncidentsTemplateTypeFromValue returns a pointer to a valid SalesforceIncidentsTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSalesforceIncidentsTemplateTypeFromValue(v string) (*SalesforceIncidentsTemplateType, error) {
	ev := SalesforceIncidentsTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SalesforceIncidentsTemplateType: valid values are %v", v, allowedSalesforceIncidentsTemplateTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SalesforceIncidentsTemplateType) IsValid() bool {
	for _, existing := range allowedSalesforceIncidentsTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SalesforceIncidentsTemplateType value.
func (v SalesforceIncidentsTemplateType) Ptr() *SalesforceIncidentsTemplateType {
	return &v
}
