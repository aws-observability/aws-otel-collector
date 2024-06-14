// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestRequestBodyType Type of the request body.
type SyntheticsTestRequestBodyType string

// List of SyntheticsTestRequestBodyType.
const (
	SYNTHETICSTESTREQUESTBODYTYPE_TEXT_PLAIN                        SyntheticsTestRequestBodyType = "text/plain"
	SYNTHETICSTESTREQUESTBODYTYPE_APPLICATION_JSON                  SyntheticsTestRequestBodyType = "application/json"
	SYNTHETICSTESTREQUESTBODYTYPE_TEXT_XML                          SyntheticsTestRequestBodyType = "text/xml"
	SYNTHETICSTESTREQUESTBODYTYPE_TEXT_HTML                         SyntheticsTestRequestBodyType = "text/html"
	SYNTHETICSTESTREQUESTBODYTYPE_APPLICATION_X_WWW_FORM_URLENCODED SyntheticsTestRequestBodyType = "application/x-www-form-urlencoded"
	SYNTHETICSTESTREQUESTBODYTYPE_GRAPHQL                           SyntheticsTestRequestBodyType = "graphql"
	SYNTHETICSTESTREQUESTBODYTYPE_APPLICATION_OCTET_STREAM          SyntheticsTestRequestBodyType = "application/octet-stream"
	SYNTHETICSTESTREQUESTBODYTYPE_MULTIPART_FORM_DATA               SyntheticsTestRequestBodyType = "multipart/form-data"
)

var allowedSyntheticsTestRequestBodyTypeEnumValues = []SyntheticsTestRequestBodyType{
	SYNTHETICSTESTREQUESTBODYTYPE_TEXT_PLAIN,
	SYNTHETICSTESTREQUESTBODYTYPE_APPLICATION_JSON,
	SYNTHETICSTESTREQUESTBODYTYPE_TEXT_XML,
	SYNTHETICSTESTREQUESTBODYTYPE_TEXT_HTML,
	SYNTHETICSTESTREQUESTBODYTYPE_APPLICATION_X_WWW_FORM_URLENCODED,
	SYNTHETICSTESTREQUESTBODYTYPE_GRAPHQL,
	SYNTHETICSTESTREQUESTBODYTYPE_APPLICATION_OCTET_STREAM,
	SYNTHETICSTESTREQUESTBODYTYPE_MULTIPART_FORM_DATA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestRequestBodyType) GetAllowedValues() []SyntheticsTestRequestBodyType {
	return allowedSyntheticsTestRequestBodyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestRequestBodyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestRequestBodyType(value)
	return nil
}

// NewSyntheticsTestRequestBodyTypeFromValue returns a pointer to a valid SyntheticsTestRequestBodyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestRequestBodyTypeFromValue(v string) (*SyntheticsTestRequestBodyType, error) {
	ev := SyntheticsTestRequestBodyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestRequestBodyType: valid values are %v", v, allowedSyntheticsTestRequestBodyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestRequestBodyType) IsValid() bool {
	for _, existing := range allowedSyntheticsTestRequestBodyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestRequestBodyType value.
func (v SyntheticsTestRequestBodyType) Ptr() *SyntheticsTestRequestBodyType {
	return &v
}
