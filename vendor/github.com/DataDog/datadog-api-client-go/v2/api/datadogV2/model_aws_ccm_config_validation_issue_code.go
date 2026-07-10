// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCcmConfigValidationIssueCode Identifies the specific reason a Cost and Usage Report (CUR) 2.0 configuration failed validation.
type AWSCcmConfigValidationIssueCode string

// List of AWSCcmConfigValidationIssueCode.
const (
	AWSCCMCONFIGVALIDATIONISSUECODE_ISSUE_CODE_UNSPECIFIED         AWSCcmConfigValidationIssueCode = "ISSUE_CODE_UNSPECIFIED"
	AWSCCMCONFIGVALIDATIONISSUECODE_CREDENTIAL_ERROR               AWSCcmConfigValidationIssueCode = "CREDENTIAL_ERROR"
	AWSCCMCONFIGVALIDATIONISSUECODE_BUCKET_NAME_INVALID_GOVCLOUD   AWSCcmConfigValidationIssueCode = "BUCKET_NAME_INVALID_GOVCLOUD"
	AWSCCMCONFIGVALIDATIONISSUECODE_S3_LIST_PERMISSION_MISSING     AWSCcmConfigValidationIssueCode = "S3_LIST_PERMISSION_MISSING"
	AWSCCMCONFIGVALIDATIONISSUECODE_S3_GET_PERMISSION_MISSING      AWSCcmConfigValidationIssueCode = "S3_GET_PERMISSION_MISSING"
	AWSCCMCONFIGVALIDATIONISSUECODE_S3_BUCKET_REGION_MISMATCH      AWSCcmConfigValidationIssueCode = "S3_BUCKET_REGION_MISMATCH"
	AWSCCMCONFIGVALIDATIONISSUECODE_S3_BUCKET_NOT_ACCESSIBLE       AWSCcmConfigValidationIssueCode = "S3_BUCKET_NOT_ACCESSIBLE"
	AWSCCMCONFIGVALIDATIONISSUECODE_EXPORT_LIST_PERMISSION_MISSING AWSCcmConfigValidationIssueCode = "EXPORT_LIST_PERMISSION_MISSING"
	AWSCCMCONFIGVALIDATIONISSUECODE_EXPORT_GET_PERMISSION_MISSING  AWSCcmConfigValidationIssueCode = "EXPORT_GET_PERMISSION_MISSING"
	AWSCCMCONFIGVALIDATIONISSUECODE_EXPORT_NOT_FOUND               AWSCcmConfigValidationIssueCode = "EXPORT_NOT_FOUND"
	AWSCCMCONFIGVALIDATIONISSUECODE_EXPORT_STATUS_UNHEALTHY        AWSCcmConfigValidationIssueCode = "EXPORT_STATUS_UNHEALTHY"
	AWSCCMCONFIGVALIDATIONISSUECODE_TIME_GRANULARITY_INVALID       AWSCcmConfigValidationIssueCode = "TIME_GRANULARITY_INVALID"
	AWSCCMCONFIGVALIDATIONISSUECODE_FILE_FORMAT_INVALID            AWSCcmConfigValidationIssueCode = "FILE_FORMAT_INVALID"
	AWSCCMCONFIGVALIDATIONISSUECODE_INCLUDE_RESOURCES_DISABLED     AWSCcmConfigValidationIssueCode = "INCLUDE_RESOURCES_DISABLED"
	AWSCCMCONFIGVALIDATIONISSUECODE_REFRESH_CADENCE_INVALID        AWSCcmConfigValidationIssueCode = "REFRESH_CADENCE_INVALID"
	AWSCCMCONFIGVALIDATIONISSUECODE_OVERWRITE_MODE_INVALID         AWSCcmConfigValidationIssueCode = "OVERWRITE_MODE_INVALID"
	AWSCCMCONFIGVALIDATIONISSUECODE_QUERY_STATEMENT_INVALID        AWSCcmConfigValidationIssueCode = "QUERY_STATEMENT_INVALID"
)

var allowedAWSCcmConfigValidationIssueCodeEnumValues = []AWSCcmConfigValidationIssueCode{
	AWSCCMCONFIGVALIDATIONISSUECODE_ISSUE_CODE_UNSPECIFIED,
	AWSCCMCONFIGVALIDATIONISSUECODE_CREDENTIAL_ERROR,
	AWSCCMCONFIGVALIDATIONISSUECODE_BUCKET_NAME_INVALID_GOVCLOUD,
	AWSCCMCONFIGVALIDATIONISSUECODE_S3_LIST_PERMISSION_MISSING,
	AWSCCMCONFIGVALIDATIONISSUECODE_S3_GET_PERMISSION_MISSING,
	AWSCCMCONFIGVALIDATIONISSUECODE_S3_BUCKET_REGION_MISMATCH,
	AWSCCMCONFIGVALIDATIONISSUECODE_S3_BUCKET_NOT_ACCESSIBLE,
	AWSCCMCONFIGVALIDATIONISSUECODE_EXPORT_LIST_PERMISSION_MISSING,
	AWSCCMCONFIGVALIDATIONISSUECODE_EXPORT_GET_PERMISSION_MISSING,
	AWSCCMCONFIGVALIDATIONISSUECODE_EXPORT_NOT_FOUND,
	AWSCCMCONFIGVALIDATIONISSUECODE_EXPORT_STATUS_UNHEALTHY,
	AWSCCMCONFIGVALIDATIONISSUECODE_TIME_GRANULARITY_INVALID,
	AWSCCMCONFIGVALIDATIONISSUECODE_FILE_FORMAT_INVALID,
	AWSCCMCONFIGVALIDATIONISSUECODE_INCLUDE_RESOURCES_DISABLED,
	AWSCCMCONFIGVALIDATIONISSUECODE_REFRESH_CADENCE_INVALID,
	AWSCCMCONFIGVALIDATIONISSUECODE_OVERWRITE_MODE_INVALID,
	AWSCCMCONFIGVALIDATIONISSUECODE_QUERY_STATEMENT_INVALID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSCcmConfigValidationIssueCode) GetAllowedValues() []AWSCcmConfigValidationIssueCode {
	return allowedAWSCcmConfigValidationIssueCodeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSCcmConfigValidationIssueCode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSCcmConfigValidationIssueCode(value)
	return nil
}

// NewAWSCcmConfigValidationIssueCodeFromValue returns a pointer to a valid AWSCcmConfigValidationIssueCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSCcmConfigValidationIssueCodeFromValue(v string) (*AWSCcmConfigValidationIssueCode, error) {
	ev := AWSCcmConfigValidationIssueCode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSCcmConfigValidationIssueCode: valid values are %v", v, allowedAWSCcmConfigValidationIssueCodeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSCcmConfigValidationIssueCode) IsValid() bool {
	for _, existing := range allowedAWSCcmConfigValidationIssueCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSCcmConfigValidationIssueCode value.
func (v AWSCcmConfigValidationIssueCode) Ptr() *AWSCcmConfigValidationIssueCode {
	return &v
}
