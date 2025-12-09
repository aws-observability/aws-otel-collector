// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureUCConfigPatchRequestType Type of Azure config Patch Request.
type AzureUCConfigPatchRequestType string

// List of AzureUCConfigPatchRequestType.
const (
	AZUREUCCONFIGPATCHREQUESTTYPE_AZURE_UC_CONFIG_PATCH_REQUEST AzureUCConfigPatchRequestType = "azure_uc_config_patch_request"
)

var allowedAzureUCConfigPatchRequestTypeEnumValues = []AzureUCConfigPatchRequestType{
	AZUREUCCONFIGPATCHREQUESTTYPE_AZURE_UC_CONFIG_PATCH_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AzureUCConfigPatchRequestType) GetAllowedValues() []AzureUCConfigPatchRequestType {
	return allowedAzureUCConfigPatchRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AzureUCConfigPatchRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AzureUCConfigPatchRequestType(value)
	return nil
}

// NewAzureUCConfigPatchRequestTypeFromValue returns a pointer to a valid AzureUCConfigPatchRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAzureUCConfigPatchRequestTypeFromValue(v string) (*AzureUCConfigPatchRequestType, error) {
	ev := AzureUCConfigPatchRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AzureUCConfigPatchRequestType: valid values are %v", v, allowedAzureUCConfigPatchRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AzureUCConfigPatchRequestType) IsValid() bool {
	for _, existing := range allowedAzureUCConfigPatchRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AzureUCConfigPatchRequestType value.
func (v AzureUCConfigPatchRequestType) Ptr() *AzureUCConfigPatchRequestType {
	return &v
}
