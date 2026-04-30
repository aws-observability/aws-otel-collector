// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPagesComponentGroupAttributesComponentsItemsStatus The status of the component.
type StatusPagesComponentGroupAttributesComponentsItemsStatus string

// List of StatusPagesComponentGroupAttributesComponentsItemsStatus.
const (
	STATUSPAGESCOMPONENTGROUPATTRIBUTESCOMPONENTSITEMSSTATUS_OPERATIONAL    StatusPagesComponentGroupAttributesComponentsItemsStatus = "operational"
	STATUSPAGESCOMPONENTGROUPATTRIBUTESCOMPONENTSITEMSSTATUS_DEGRADED       StatusPagesComponentGroupAttributesComponentsItemsStatus = "degraded"
	STATUSPAGESCOMPONENTGROUPATTRIBUTESCOMPONENTSITEMSSTATUS_PARTIAL_OUTAGE StatusPagesComponentGroupAttributesComponentsItemsStatus = "partial_outage"
	STATUSPAGESCOMPONENTGROUPATTRIBUTESCOMPONENTSITEMSSTATUS_MAJOR_OUTAGE   StatusPagesComponentGroupAttributesComponentsItemsStatus = "major_outage"
	STATUSPAGESCOMPONENTGROUPATTRIBUTESCOMPONENTSITEMSSTATUS_MAINTENANCE    StatusPagesComponentGroupAttributesComponentsItemsStatus = "maintenance"
)

var allowedStatusPagesComponentGroupAttributesComponentsItemsStatusEnumValues = []StatusPagesComponentGroupAttributesComponentsItemsStatus{
	STATUSPAGESCOMPONENTGROUPATTRIBUTESCOMPONENTSITEMSSTATUS_OPERATIONAL,
	STATUSPAGESCOMPONENTGROUPATTRIBUTESCOMPONENTSITEMSSTATUS_DEGRADED,
	STATUSPAGESCOMPONENTGROUPATTRIBUTESCOMPONENTSITEMSSTATUS_PARTIAL_OUTAGE,
	STATUSPAGESCOMPONENTGROUPATTRIBUTESCOMPONENTSITEMSSTATUS_MAJOR_OUTAGE,
	STATUSPAGESCOMPONENTGROUPATTRIBUTESCOMPONENTSITEMSSTATUS_MAINTENANCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StatusPagesComponentGroupAttributesComponentsItemsStatus) GetAllowedValues() []StatusPagesComponentGroupAttributesComponentsItemsStatus {
	return allowedStatusPagesComponentGroupAttributesComponentsItemsStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StatusPagesComponentGroupAttributesComponentsItemsStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StatusPagesComponentGroupAttributesComponentsItemsStatus(value)
	return nil
}

// NewStatusPagesComponentGroupAttributesComponentsItemsStatusFromValue returns a pointer to a valid StatusPagesComponentGroupAttributesComponentsItemsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStatusPagesComponentGroupAttributesComponentsItemsStatusFromValue(v string) (*StatusPagesComponentGroupAttributesComponentsItemsStatus, error) {
	ev := StatusPagesComponentGroupAttributesComponentsItemsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StatusPagesComponentGroupAttributesComponentsItemsStatus: valid values are %v", v, allowedStatusPagesComponentGroupAttributesComponentsItemsStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StatusPagesComponentGroupAttributesComponentsItemsStatus) IsValid() bool {
	for _, existing := range allowedStatusPagesComponentGroupAttributesComponentsItemsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusPagesComponentGroupAttributesComponentsItemsStatus value.
func (v StatusPagesComponentGroupAttributesComponentsItemsStatus) Ptr() *StatusPagesComponentGroupAttributesComponentsItemsStatus {
	return &v
}
