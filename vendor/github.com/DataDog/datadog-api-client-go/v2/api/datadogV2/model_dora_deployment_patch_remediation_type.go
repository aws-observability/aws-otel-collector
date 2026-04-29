// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORADeploymentPatchRemediationType The type of remediation action taken. Required when the failed deployment must be linked to a remediation deployment.
type DORADeploymentPatchRemediationType string

// List of DORADeploymentPatchRemediationType.
const (
	DORADEPLOYMENTPATCHREMEDIATIONTYPE_ROLLBACK    DORADeploymentPatchRemediationType = "rollback"
	DORADEPLOYMENTPATCHREMEDIATIONTYPE_ROLLFORWARD DORADeploymentPatchRemediationType = "rollforward"
)

var allowedDORADeploymentPatchRemediationTypeEnumValues = []DORADeploymentPatchRemediationType{
	DORADEPLOYMENTPATCHREMEDIATIONTYPE_ROLLBACK,
	DORADEPLOYMENTPATCHREMEDIATIONTYPE_ROLLFORWARD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DORADeploymentPatchRemediationType) GetAllowedValues() []DORADeploymentPatchRemediationType {
	return allowedDORADeploymentPatchRemediationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DORADeploymentPatchRemediationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DORADeploymentPatchRemediationType(value)
	return nil
}

// NewDORADeploymentPatchRemediationTypeFromValue returns a pointer to a valid DORADeploymentPatchRemediationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDORADeploymentPatchRemediationTypeFromValue(v string) (*DORADeploymentPatchRemediationType, error) {
	ev := DORADeploymentPatchRemediationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DORADeploymentPatchRemediationType: valid values are %v", v, allowedDORADeploymentPatchRemediationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DORADeploymentPatchRemediationType) IsValid() bool {
	for _, existing := range allowedDORADeploymentPatchRemediationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DORADeploymentPatchRemediationType value.
func (v DORADeploymentPatchRemediationType) Ptr() *DORADeploymentPatchRemediationType {
	return &v
}
