// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafCustomRuleConditionInputAddress Input from the request on which the condition should apply.
type ApplicationSecurityWafCustomRuleConditionInputAddress string

// List of ApplicationSecurityWafCustomRuleConditionInputAddress.
const (
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_DB_STATEMENT                ApplicationSecurityWafCustomRuleConditionInputAddress = "server.db.statement"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_IO_FS_FILE                  ApplicationSecurityWafCustomRuleConditionInputAddress = "server.io.fs.file"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_IO_NET_URL                  ApplicationSecurityWafCustomRuleConditionInputAddress = "server.io.net.url"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_SYS_SHELL_CMD               ApplicationSecurityWafCustomRuleConditionInputAddress = "server.sys.shell.cmd"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_METHOD              ApplicationSecurityWafCustomRuleConditionInputAddress = "server.request.method"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_URI_RAW             ApplicationSecurityWafCustomRuleConditionInputAddress = "server.request.uri.raw"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_PATH_PARAMS         ApplicationSecurityWafCustomRuleConditionInputAddress = "server.request.path_params"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_QUERY               ApplicationSecurityWafCustomRuleConditionInputAddress = "server.request.query"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_HEADERS_NO_COOKIES  ApplicationSecurityWafCustomRuleConditionInputAddress = "server.request.headers.no_cookies"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_COOKIES             ApplicationSecurityWafCustomRuleConditionInputAddress = "server.request.cookies"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_TRAILERS            ApplicationSecurityWafCustomRuleConditionInputAddress = "server.request.trailers"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_BODY                ApplicationSecurityWafCustomRuleConditionInputAddress = "server.request.body"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_RESPONSE_STATUS             ApplicationSecurityWafCustomRuleConditionInputAddress = "server.response.status"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_RESPONSE_HEADERS_NO_COOKIES ApplicationSecurityWafCustomRuleConditionInputAddress = "server.response.headers.no_cookies"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_RESPONSE_TRAILERS           ApplicationSecurityWafCustomRuleConditionInputAddress = "server.response.trailers"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_GRPC_SERVER_REQUEST_METADATA       ApplicationSecurityWafCustomRuleConditionInputAddress = "grpc.server.request.metadata"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_GRPC_SERVER_REQUEST_MESSAGE        ApplicationSecurityWafCustomRuleConditionInputAddress = "grpc.server.request.message"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_GRPC_SERVER_METHOD                 ApplicationSecurityWafCustomRuleConditionInputAddress = "grpc.server.method"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_GRAPHQL_SERVER_ALL_RESOLVERS       ApplicationSecurityWafCustomRuleConditionInputAddress = "graphql.server.all_resolvers"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_USR_ID                             ApplicationSecurityWafCustomRuleConditionInputAddress = "usr.id"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_HTTP_CLIENT_IP                     ApplicationSecurityWafCustomRuleConditionInputAddress = "http.client_ip"
)

var allowedApplicationSecurityWafCustomRuleConditionInputAddressEnumValues = []ApplicationSecurityWafCustomRuleConditionInputAddress{
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_DB_STATEMENT,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_IO_FS_FILE,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_IO_NET_URL,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_SYS_SHELL_CMD,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_METHOD,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_URI_RAW,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_PATH_PARAMS,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_QUERY,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_HEADERS_NO_COOKIES,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_COOKIES,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_TRAILERS,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_REQUEST_BODY,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_RESPONSE_STATUS,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_RESPONSE_HEADERS_NO_COOKIES,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_SERVER_RESPONSE_TRAILERS,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_GRPC_SERVER_REQUEST_METADATA,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_GRPC_SERVER_REQUEST_MESSAGE,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_GRPC_SERVER_METHOD,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_GRAPHQL_SERVER_ALL_RESOLVERS,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_USR_ID,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONINPUTADDRESS_HTTP_CLIENT_IP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApplicationSecurityWafCustomRuleConditionInputAddress) GetAllowedValues() []ApplicationSecurityWafCustomRuleConditionInputAddress {
	return allowedApplicationSecurityWafCustomRuleConditionInputAddressEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApplicationSecurityWafCustomRuleConditionInputAddress) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApplicationSecurityWafCustomRuleConditionInputAddress(value)
	return nil
}

// NewApplicationSecurityWafCustomRuleConditionInputAddressFromValue returns a pointer to a valid ApplicationSecurityWafCustomRuleConditionInputAddress
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApplicationSecurityWafCustomRuleConditionInputAddressFromValue(v string) (*ApplicationSecurityWafCustomRuleConditionInputAddress, error) {
	ev := ApplicationSecurityWafCustomRuleConditionInputAddress(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApplicationSecurityWafCustomRuleConditionInputAddress: valid values are %v", v, allowedApplicationSecurityWafCustomRuleConditionInputAddressEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApplicationSecurityWafCustomRuleConditionInputAddress) IsValid() bool {
	for _, existing := range allowedApplicationSecurityWafCustomRuleConditionInputAddressEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationSecurityWafCustomRuleConditionInputAddress value.
func (v ApplicationSecurityWafCustomRuleConditionInputAddress) Ptr() *ApplicationSecurityWafCustomRuleConditionInputAddress {
	return &v
}
