// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORADeploymentPatchRequestDataType JSON:API type for DORA deployment patch request.
type DORADeploymentPatchRequestDataType string

// List of DORADeploymentPatchRequestDataType.
const (
	DORADEPLOYMENTPATCHREQUESTDATATYPE_DORA_DEPLOYMENT_PATCH_REQUEST DORADeploymentPatchRequestDataType = "dora_deployment_patch_request"
)

var allowedDORADeploymentPatchRequestDataTypeEnumValues = []DORADeploymentPatchRequestDataType{
	DORADEPLOYMENTPATCHREQUESTDATATYPE_DORA_DEPLOYMENT_PATCH_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DORADeploymentPatchRequestDataType) GetAllowedValues() []DORADeploymentPatchRequestDataType {
	return allowedDORADeploymentPatchRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DORADeploymentPatchRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DORADeploymentPatchRequestDataType(value)
	return nil
}

// NewDORADeploymentPatchRequestDataTypeFromValue returns a pointer to a valid DORADeploymentPatchRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDORADeploymentPatchRequestDataTypeFromValue(v string) (*DORADeploymentPatchRequestDataType, error) {
	ev := DORADeploymentPatchRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DORADeploymentPatchRequestDataType: valid values are %v", v, allowedDORADeploymentPatchRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DORADeploymentPatchRequestDataType) IsValid() bool {
	for _, existing := range allowedDORADeploymentPatchRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DORADeploymentPatchRequestDataType value.
func (v DORADeploymentPatchRequestDataType) Ptr() *DORADeploymentPatchRequestDataType {
	return &v
}
