// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsApiTestFailureCode Error code that can be returned by a Synthetic test.
type SyntheticsApiTestFailureCode string

// List of SyntheticsApiTestFailureCode.
const (
	SYNTHETICSAPITESTFAILURECODE_BODY_TOO_LARGE                       SyntheticsApiTestFailureCode = "BODY_TOO_LARGE"
	SYNTHETICSAPITESTFAILURECODE_DENIED                               SyntheticsApiTestFailureCode = "DENIED"
	SYNTHETICSAPITESTFAILURECODE_TOO_MANY_REDIRECTS                   SyntheticsApiTestFailureCode = "TOO_MANY_REDIRECTS"
	SYNTHETICSAPITESTFAILURECODE_AUTHENTICATION_ERROR                 SyntheticsApiTestFailureCode = "AUTHENTICATION_ERROR"
	SYNTHETICSAPITESTFAILURECODE_DECRYPTION                           SyntheticsApiTestFailureCode = "DECRYPTION"
	SYNTHETICSAPITESTFAILURECODE_INVALID_CHAR_IN_HEADER               SyntheticsApiTestFailureCode = "INVALID_CHAR_IN_HEADER"
	SYNTHETICSAPITESTFAILURECODE_HEADER_TOO_LARGE                     SyntheticsApiTestFailureCode = "HEADER_TOO_LARGE"
	SYNTHETICSAPITESTFAILURECODE_HEADERS_INCOMPATIBLE_CONTENT_LENGTH  SyntheticsApiTestFailureCode = "HEADERS_INCOMPATIBLE_CONTENT_LENGTH"
	SYNTHETICSAPITESTFAILURECODE_INVALID_REQUEST                      SyntheticsApiTestFailureCode = "INVALID_REQUEST"
	SYNTHETICSAPITESTFAILURECODE_REQUIRES_UPDATE                      SyntheticsApiTestFailureCode = "REQUIRES_UPDATE"
	SYNTHETICSAPITESTFAILURECODE_UNESCAPED_CHARACTERS_IN_REQUEST_PATH SyntheticsApiTestFailureCode = "UNESCAPED_CHARACTERS_IN_REQUEST_PATH"
	SYNTHETICSAPITESTFAILURECODE_MALFORMED_RESPONSE                   SyntheticsApiTestFailureCode = "MALFORMED_RESPONSE"
	SYNTHETICSAPITESTFAILURECODE_INCORRECT_ASSERTION                  SyntheticsApiTestFailureCode = "INCORRECT_ASSERTION"
	SYNTHETICSAPITESTFAILURECODE_CONNREFUSED                          SyntheticsApiTestFailureCode = "CONNREFUSED"
	SYNTHETICSAPITESTFAILURECODE_CONNRESET                            SyntheticsApiTestFailureCode = "CONNRESET"
	SYNTHETICSAPITESTFAILURECODE_DNS                                  SyntheticsApiTestFailureCode = "DNS"
	SYNTHETICSAPITESTFAILURECODE_HOSTUNREACH                          SyntheticsApiTestFailureCode = "HOSTUNREACH"
	SYNTHETICSAPITESTFAILURECODE_NETUNREACH                           SyntheticsApiTestFailureCode = "NETUNREACH"
	SYNTHETICSAPITESTFAILURECODE_TIMEOUT                              SyntheticsApiTestFailureCode = "TIMEOUT"
	SYNTHETICSAPITESTFAILURECODE_SSL                                  SyntheticsApiTestFailureCode = "SSL"
	SYNTHETICSAPITESTFAILURECODE_OCSP                                 SyntheticsApiTestFailureCode = "OCSP"
	SYNTHETICSAPITESTFAILURECODE_INVALID_TEST                         SyntheticsApiTestFailureCode = "INVALID_TEST"
	SYNTHETICSAPITESTFAILURECODE_TUNNEL                               SyntheticsApiTestFailureCode = "TUNNEL"
	SYNTHETICSAPITESTFAILURECODE_WEBSOCKET                            SyntheticsApiTestFailureCode = "WEBSOCKET"
	SYNTHETICSAPITESTFAILURECODE_UNKNOWN                              SyntheticsApiTestFailureCode = "UNKNOWN"
	SYNTHETICSAPITESTFAILURECODE_INTERNAL_ERROR                       SyntheticsApiTestFailureCode = "INTERNAL_ERROR"
)

var allowedSyntheticsApiTestFailureCodeEnumValues = []SyntheticsApiTestFailureCode{
	SYNTHETICSAPITESTFAILURECODE_BODY_TOO_LARGE,
	SYNTHETICSAPITESTFAILURECODE_DENIED,
	SYNTHETICSAPITESTFAILURECODE_TOO_MANY_REDIRECTS,
	SYNTHETICSAPITESTFAILURECODE_AUTHENTICATION_ERROR,
	SYNTHETICSAPITESTFAILURECODE_DECRYPTION,
	SYNTHETICSAPITESTFAILURECODE_INVALID_CHAR_IN_HEADER,
	SYNTHETICSAPITESTFAILURECODE_HEADER_TOO_LARGE,
	SYNTHETICSAPITESTFAILURECODE_HEADERS_INCOMPATIBLE_CONTENT_LENGTH,
	SYNTHETICSAPITESTFAILURECODE_INVALID_REQUEST,
	SYNTHETICSAPITESTFAILURECODE_REQUIRES_UPDATE,
	SYNTHETICSAPITESTFAILURECODE_UNESCAPED_CHARACTERS_IN_REQUEST_PATH,
	SYNTHETICSAPITESTFAILURECODE_MALFORMED_RESPONSE,
	SYNTHETICSAPITESTFAILURECODE_INCORRECT_ASSERTION,
	SYNTHETICSAPITESTFAILURECODE_CONNREFUSED,
	SYNTHETICSAPITESTFAILURECODE_CONNRESET,
	SYNTHETICSAPITESTFAILURECODE_DNS,
	SYNTHETICSAPITESTFAILURECODE_HOSTUNREACH,
	SYNTHETICSAPITESTFAILURECODE_NETUNREACH,
	SYNTHETICSAPITESTFAILURECODE_TIMEOUT,
	SYNTHETICSAPITESTFAILURECODE_SSL,
	SYNTHETICSAPITESTFAILURECODE_OCSP,
	SYNTHETICSAPITESTFAILURECODE_INVALID_TEST,
	SYNTHETICSAPITESTFAILURECODE_TUNNEL,
	SYNTHETICSAPITESTFAILURECODE_WEBSOCKET,
	SYNTHETICSAPITESTFAILURECODE_UNKNOWN,
	SYNTHETICSAPITESTFAILURECODE_INTERNAL_ERROR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsApiTestFailureCode) GetAllowedValues() []SyntheticsApiTestFailureCode {
	return allowedSyntheticsApiTestFailureCodeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsApiTestFailureCode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsApiTestFailureCode(value)
	return nil
}

// NewSyntheticsApiTestFailureCodeFromValue returns a pointer to a valid SyntheticsApiTestFailureCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsApiTestFailureCodeFromValue(v string) (*SyntheticsApiTestFailureCode, error) {
	ev := SyntheticsApiTestFailureCode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsApiTestFailureCode: valid values are %v", v, allowedSyntheticsApiTestFailureCodeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsApiTestFailureCode) IsValid() bool {
	for _, existing := range allowedSyntheticsApiTestFailureCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsApiTestFailureCode value.
func (v SyntheticsApiTestFailureCode) Ptr() *SyntheticsApiTestFailureCode {
	return &v
}
