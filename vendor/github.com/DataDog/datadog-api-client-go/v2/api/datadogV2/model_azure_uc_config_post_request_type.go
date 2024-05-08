// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureUCConfigPostRequestType Type of Azure config Post Request.
type AzureUCConfigPostRequestType string

// List of AzureUCConfigPostRequestType.
const (
	AZUREUCCONFIGPOSTREQUESTTYPE_AZURE_UC_CONFIG_POST_REQUEST AzureUCConfigPostRequestType = "azure_uc_config_post_request"
)

var allowedAzureUCConfigPostRequestTypeEnumValues = []AzureUCConfigPostRequestType{
	AZUREUCCONFIGPOSTREQUESTTYPE_AZURE_UC_CONFIG_POST_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AzureUCConfigPostRequestType) GetAllowedValues() []AzureUCConfigPostRequestType {
	return allowedAzureUCConfigPostRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AzureUCConfigPostRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AzureUCConfigPostRequestType(value)
	return nil
}

// NewAzureUCConfigPostRequestTypeFromValue returns a pointer to a valid AzureUCConfigPostRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAzureUCConfigPostRequestTypeFromValue(v string) (*AzureUCConfigPostRequestType, error) {
	ev := AzureUCConfigPostRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AzureUCConfigPostRequestType: valid values are %v", v, allowedAzureUCConfigPostRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AzureUCConfigPostRequestType) IsValid() bool {
	for _, existing := range allowedAzureUCConfigPostRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AzureUCConfigPostRequestType value.
func (v AzureUCConfigPostRequestType) Ptr() *AzureUCConfigPostRequestType {
	return &v
}
