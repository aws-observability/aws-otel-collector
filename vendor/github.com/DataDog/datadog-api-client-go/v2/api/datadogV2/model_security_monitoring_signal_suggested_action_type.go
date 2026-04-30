// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalSuggestedActionType The type of the suggested action resource.
type SecurityMonitoringSignalSuggestedActionType string

// List of SecurityMonitoringSignalSuggestedActionType.
const (
	SECURITYMONITORINGSIGNALSUGGESTEDACTIONTYPE_INVESTIGATION_LOG_QUERIES SecurityMonitoringSignalSuggestedActionType = "investigation_log_queries"
	SECURITYMONITORINGSIGNALSUGGESTEDACTIONTYPE_RECOMMENDED_BLOG_POSTS    SecurityMonitoringSignalSuggestedActionType = "recommended_blog_posts"
)

var allowedSecurityMonitoringSignalSuggestedActionTypeEnumValues = []SecurityMonitoringSignalSuggestedActionType{
	SECURITYMONITORINGSIGNALSUGGESTEDACTIONTYPE_INVESTIGATION_LOG_QUERIES,
	SECURITYMONITORINGSIGNALSUGGESTEDACTIONTYPE_RECOMMENDED_BLOG_POSTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringSignalSuggestedActionType) GetAllowedValues() []SecurityMonitoringSignalSuggestedActionType {
	return allowedSecurityMonitoringSignalSuggestedActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringSignalSuggestedActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringSignalSuggestedActionType(value)
	return nil
}

// NewSecurityMonitoringSignalSuggestedActionTypeFromValue returns a pointer to a valid SecurityMonitoringSignalSuggestedActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringSignalSuggestedActionTypeFromValue(v string) (*SecurityMonitoringSignalSuggestedActionType, error) {
	ev := SecurityMonitoringSignalSuggestedActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringSignalSuggestedActionType: valid values are %v", v, allowedSecurityMonitoringSignalSuggestedActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringSignalSuggestedActionType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringSignalSuggestedActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringSignalSuggestedActionType value.
func (v SecurityMonitoringSignalSuggestedActionType) Ptr() *SecurityMonitoringSignalSuggestedActionType {
	return &v
}
