// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackTimestampBucket Timestamp bucket indicating when logs were last collected.
type SecurityMonitoringContentPackTimestampBucket string

// List of SecurityMonitoringContentPackTimestampBucket.
const (
	SECURITYMONITORINGCONTENTPACKTIMESTAMPBUCKET_NOT_SEEN              SecurityMonitoringContentPackTimestampBucket = "not_seen"
	SECURITYMONITORINGCONTENTPACKTIMESTAMPBUCKET_WITHIN_24_HOURS       SecurityMonitoringContentPackTimestampBucket = "within_24_hours"
	SECURITYMONITORINGCONTENTPACKTIMESTAMPBUCKET_WITHIN_24_TO_72_HOURS SecurityMonitoringContentPackTimestampBucket = "within_24_to_72_hours"
	SECURITYMONITORINGCONTENTPACKTIMESTAMPBUCKET_OVER_72H_TO_30D       SecurityMonitoringContentPackTimestampBucket = "over_72h_to_30d"
	SECURITYMONITORINGCONTENTPACKTIMESTAMPBUCKET_OVER_30D              SecurityMonitoringContentPackTimestampBucket = "over_30d"
)

var allowedSecurityMonitoringContentPackTimestampBucketEnumValues = []SecurityMonitoringContentPackTimestampBucket{
	SECURITYMONITORINGCONTENTPACKTIMESTAMPBUCKET_NOT_SEEN,
	SECURITYMONITORINGCONTENTPACKTIMESTAMPBUCKET_WITHIN_24_HOURS,
	SECURITYMONITORINGCONTENTPACKTIMESTAMPBUCKET_WITHIN_24_TO_72_HOURS,
	SECURITYMONITORINGCONTENTPACKTIMESTAMPBUCKET_OVER_72H_TO_30D,
	SECURITYMONITORINGCONTENTPACKTIMESTAMPBUCKET_OVER_30D,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringContentPackTimestampBucket) GetAllowedValues() []SecurityMonitoringContentPackTimestampBucket {
	return allowedSecurityMonitoringContentPackTimestampBucketEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringContentPackTimestampBucket) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringContentPackTimestampBucket(value)
	return nil
}

// NewSecurityMonitoringContentPackTimestampBucketFromValue returns a pointer to a valid SecurityMonitoringContentPackTimestampBucket
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringContentPackTimestampBucketFromValue(v string) (*SecurityMonitoringContentPackTimestampBucket, error) {
	ev := SecurityMonitoringContentPackTimestampBucket(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringContentPackTimestampBucket: valid values are %v", v, allowedSecurityMonitoringContentPackTimestampBucketEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringContentPackTimestampBucket) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringContentPackTimestampBucketEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringContentPackTimestampBucket value.
func (v SecurityMonitoringContentPackTimestampBucket) Ptr() *SecurityMonitoringContentPackTimestampBucket {
	return &v
}
