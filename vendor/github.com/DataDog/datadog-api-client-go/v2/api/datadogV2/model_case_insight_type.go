// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseInsightType The type of Datadog resource linked to the case as contextual evidence. Each type corresponds to a different Datadog product signal (for example, a security finding, a monitor alert, or an incident).
type CaseInsightType string

// List of CaseInsightType.
const (
	CASEINSIGHTTYPE_SECURITY_SIGNAL                 CaseInsightType = "SECURITY_SIGNAL"
	CASEINSIGHTTYPE_MONITOR                         CaseInsightType = "MONITOR"
	CASEINSIGHTTYPE_EVENT_CORRELATION               CaseInsightType = "EVENT_CORRELATION"
	CASEINSIGHTTYPE_ERROR_TRACKING                  CaseInsightType = "ERROR_TRACKING"
	CASEINSIGHTTYPE_CLOUD_COST_RECOMMENDATION       CaseInsightType = "CLOUD_COST_RECOMMENDATION"
	CASEINSIGHTTYPE_INCIDENT                        CaseInsightType = "INCIDENT"
	CASEINSIGHTTYPE_SENSITIVE_DATA_SCANNER_ISSUE    CaseInsightType = "SENSITIVE_DATA_SCANNER_ISSUE"
	CASEINSIGHTTYPE_EVENT                           CaseInsightType = "EVENT"
	CASEINSIGHTTYPE_WATCHDOG_STORY                  CaseInsightType = "WATCHDOG_STORY"
	CASEINSIGHTTYPE_WIDGET                          CaseInsightType = "WIDGET"
	CASEINSIGHTTYPE_SECURITY_FINDING                CaseInsightType = "SECURITY_FINDING"
	CASEINSIGHTTYPE_INSIGHT_SCORECARD_CAMPAIGN      CaseInsightType = "INSIGHT_SCORECARD_CAMPAIGN"
	CASEINSIGHTTYPE_RESOURCE_POLICY                 CaseInsightType = "RESOURCE_POLICY"
	CASEINSIGHTTYPE_APM_RECOMMENDATION              CaseInsightType = "APM_RECOMMENDATION"
	CASEINSIGHTTYPE_SCM_URL                         CaseInsightType = "SCM_URL"
	CASEINSIGHTTYPE_PROFILING_DOWNSIZING_EXPERIMENT CaseInsightType = "PROFILING_DOWNSIZING_EXPERIMENT"
)

var allowedCaseInsightTypeEnumValues = []CaseInsightType{
	CASEINSIGHTTYPE_SECURITY_SIGNAL,
	CASEINSIGHTTYPE_MONITOR,
	CASEINSIGHTTYPE_EVENT_CORRELATION,
	CASEINSIGHTTYPE_ERROR_TRACKING,
	CASEINSIGHTTYPE_CLOUD_COST_RECOMMENDATION,
	CASEINSIGHTTYPE_INCIDENT,
	CASEINSIGHTTYPE_SENSITIVE_DATA_SCANNER_ISSUE,
	CASEINSIGHTTYPE_EVENT,
	CASEINSIGHTTYPE_WATCHDOG_STORY,
	CASEINSIGHTTYPE_WIDGET,
	CASEINSIGHTTYPE_SECURITY_FINDING,
	CASEINSIGHTTYPE_INSIGHT_SCORECARD_CAMPAIGN,
	CASEINSIGHTTYPE_RESOURCE_POLICY,
	CASEINSIGHTTYPE_APM_RECOMMENDATION,
	CASEINSIGHTTYPE_SCM_URL,
	CASEINSIGHTTYPE_PROFILING_DOWNSIZING_EXPERIMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseInsightType) GetAllowedValues() []CaseInsightType {
	return allowedCaseInsightTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseInsightType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseInsightType(value)
	return nil
}

// NewCaseInsightTypeFromValue returns a pointer to a valid CaseInsightType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseInsightTypeFromValue(v string) (*CaseInsightType, error) {
	ev := CaseInsightType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseInsightType: valid values are %v", v, allowedCaseInsightTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseInsightType) IsValid() bool {
	for _, existing := range allowedCaseInsightTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseInsightType value.
func (v CaseInsightType) Ptr() *CaseInsightType {
	return &v
}
