// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPagesComponentDataAttributesStatus The status of the component.
type StatusPagesComponentDataAttributesStatus string

// List of StatusPagesComponentDataAttributesStatus.
const (
	STATUSPAGESCOMPONENTDATAATTRIBUTESSTATUS_OPERATIONAL    StatusPagesComponentDataAttributesStatus = "operational"
	STATUSPAGESCOMPONENTDATAATTRIBUTESSTATUS_DEGRADED       StatusPagesComponentDataAttributesStatus = "degraded"
	STATUSPAGESCOMPONENTDATAATTRIBUTESSTATUS_PARTIAL_OUTAGE StatusPagesComponentDataAttributesStatus = "partial_outage"
	STATUSPAGESCOMPONENTDATAATTRIBUTESSTATUS_MAJOR_OUTAGE   StatusPagesComponentDataAttributesStatus = "major_outage"
	STATUSPAGESCOMPONENTDATAATTRIBUTESSTATUS_MAINTENANCE    StatusPagesComponentDataAttributesStatus = "maintenance"
)

var allowedStatusPagesComponentDataAttributesStatusEnumValues = []StatusPagesComponentDataAttributesStatus{
	STATUSPAGESCOMPONENTDATAATTRIBUTESSTATUS_OPERATIONAL,
	STATUSPAGESCOMPONENTDATAATTRIBUTESSTATUS_DEGRADED,
	STATUSPAGESCOMPONENTDATAATTRIBUTESSTATUS_PARTIAL_OUTAGE,
	STATUSPAGESCOMPONENTDATAATTRIBUTESSTATUS_MAJOR_OUTAGE,
	STATUSPAGESCOMPONENTDATAATTRIBUTESSTATUS_MAINTENANCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StatusPagesComponentDataAttributesStatus) GetAllowedValues() []StatusPagesComponentDataAttributesStatus {
	return allowedStatusPagesComponentDataAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StatusPagesComponentDataAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StatusPagesComponentDataAttributesStatus(value)
	return nil
}

// NewStatusPagesComponentDataAttributesStatusFromValue returns a pointer to a valid StatusPagesComponentDataAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStatusPagesComponentDataAttributesStatusFromValue(v string) (*StatusPagesComponentDataAttributesStatus, error) {
	ev := StatusPagesComponentDataAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StatusPagesComponentDataAttributesStatus: valid values are %v", v, allowedStatusPagesComponentDataAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StatusPagesComponentDataAttributesStatus) IsValid() bool {
	for _, existing := range allowedStatusPagesComponentDataAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusPagesComponentDataAttributesStatus value.
func (v StatusPagesComponentDataAttributesStatus) Ptr() *StatusPagesComponentDataAttributesStatus {
	return &v
}
