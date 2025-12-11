// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentFileOp Type of file operation to perform on the target configuration file.
//   - `merge-patch`: Merges the provided patch data with the existing configuration file.
//     Creates the file if it doesn't exist.
//   - `delete`: Removes the specified configuration file from the target hosts.
type FleetDeploymentFileOp string

// List of FleetDeploymentFileOp.
const (
	FLEETDEPLOYMENTFILEOP_MERGE_PATCH FleetDeploymentFileOp = "merge-patch"
	FLEETDEPLOYMENTFILEOP_DELETE      FleetDeploymentFileOp = "delete"
)

var allowedFleetDeploymentFileOpEnumValues = []FleetDeploymentFileOp{
	FLEETDEPLOYMENTFILEOP_MERGE_PATCH,
	FLEETDEPLOYMENTFILEOP_DELETE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FleetDeploymentFileOp) GetAllowedValues() []FleetDeploymentFileOp {
	return allowedFleetDeploymentFileOpEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FleetDeploymentFileOp) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FleetDeploymentFileOp(value)
	return nil
}

// NewFleetDeploymentFileOpFromValue returns a pointer to a valid FleetDeploymentFileOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFleetDeploymentFileOpFromValue(v string) (*FleetDeploymentFileOp, error) {
	ev := FleetDeploymentFileOp(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FleetDeploymentFileOp: valid values are %v", v, allowedFleetDeploymentFileOpEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FleetDeploymentFileOp) IsValid() bool {
	for _, existing := range allowedFleetDeploymentFileOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FleetDeploymentFileOp value.
func (v FleetDeploymentFileOp) Ptr() *FleetDeploymentFileOp {
	return &v
}
