// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SalesforceIncidentsTemplatePriority Priority of the Salesforce incident created from this template.
type SalesforceIncidentsTemplatePriority string

// List of SalesforceIncidentsTemplatePriority.
const (
	SALESFORCEINCIDENTSTEMPLATEPRIORITY_CRITICAL SalesforceIncidentsTemplatePriority = "Critical"
	SALESFORCEINCIDENTSTEMPLATEPRIORITY_HIGH     SalesforceIncidentsTemplatePriority = "High"
	SALESFORCEINCIDENTSTEMPLATEPRIORITY_MODERATE SalesforceIncidentsTemplatePriority = "Moderate"
	SALESFORCEINCIDENTSTEMPLATEPRIORITY_LOW      SalesforceIncidentsTemplatePriority = "Low"
)

var allowedSalesforceIncidentsTemplatePriorityEnumValues = []SalesforceIncidentsTemplatePriority{
	SALESFORCEINCIDENTSTEMPLATEPRIORITY_CRITICAL,
	SALESFORCEINCIDENTSTEMPLATEPRIORITY_HIGH,
	SALESFORCEINCIDENTSTEMPLATEPRIORITY_MODERATE,
	SALESFORCEINCIDENTSTEMPLATEPRIORITY_LOW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SalesforceIncidentsTemplatePriority) GetAllowedValues() []SalesforceIncidentsTemplatePriority {
	return allowedSalesforceIncidentsTemplatePriorityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SalesforceIncidentsTemplatePriority) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SalesforceIncidentsTemplatePriority(value)
	return nil
}

// NewSalesforceIncidentsTemplatePriorityFromValue returns a pointer to a valid SalesforceIncidentsTemplatePriority
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSalesforceIncidentsTemplatePriorityFromValue(v string) (*SalesforceIncidentsTemplatePriority, error) {
	ev := SalesforceIncidentsTemplatePriority(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SalesforceIncidentsTemplatePriority: valid values are %v", v, allowedSalesforceIncidentsTemplatePriorityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SalesforceIncidentsTemplatePriority) IsValid() bool {
	for _, existing := range allowedSalesforceIncidentsTemplatePriorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SalesforceIncidentsTemplatePriority value.
func (v SalesforceIncidentsTemplatePriority) Ptr() *SalesforceIncidentsTemplatePriority {
	return &v
}
