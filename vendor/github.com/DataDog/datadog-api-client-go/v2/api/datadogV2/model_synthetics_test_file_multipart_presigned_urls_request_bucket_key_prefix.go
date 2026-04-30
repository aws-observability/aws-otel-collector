// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix The bucket key prefix indicating the type of file upload.
type SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix string

// List of SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix.
const (
	SYNTHETICSTESTFILEMULTIPARTPRESIGNEDURLSREQUESTBUCKETKEYPREFIX_API_UPLOAD_FILE          SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix = "api-upload-file"
	SYNTHETICSTESTFILEMULTIPARTPRESIGNEDURLSREQUESTBUCKETKEYPREFIX_BROWSER_UPLOAD_FILE_STEP SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix = "browser-upload-file-step"
)

var allowedSyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefixEnumValues = []SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix{
	SYNTHETICSTESTFILEMULTIPARTPRESIGNEDURLSREQUESTBUCKETKEYPREFIX_API_UPLOAD_FILE,
	SYNTHETICSTESTFILEMULTIPARTPRESIGNEDURLSREQUESTBUCKETKEYPREFIX_BROWSER_UPLOAD_FILE_STEP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix) GetAllowedValues() []SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix {
	return allowedSyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefixEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix(value)
	return nil
}

// NewSyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefixFromValue returns a pointer to a valid SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefixFromValue(v string) (*SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix, error) {
	ev := SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix: valid values are %v", v, allowedSyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefixEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix) IsValid() bool {
	for _, existing := range allowedSyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefixEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix value.
func (v SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix) Ptr() *SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix {
	return &v
}
