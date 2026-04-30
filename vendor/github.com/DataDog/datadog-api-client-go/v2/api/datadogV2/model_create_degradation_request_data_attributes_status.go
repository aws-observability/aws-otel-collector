// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateDegradationRequestDataAttributesStatus The status of the degradation.
type CreateDegradationRequestDataAttributesStatus string

// List of CreateDegradationRequestDataAttributesStatus.
const (
	CREATEDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_INVESTIGATING CreateDegradationRequestDataAttributesStatus = "investigating"
	CREATEDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_IDENTIFIED    CreateDegradationRequestDataAttributesStatus = "identified"
	CREATEDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_MONITORING    CreateDegradationRequestDataAttributesStatus = "monitoring"
	CREATEDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_RESOLVED      CreateDegradationRequestDataAttributesStatus = "resolved"
)

var allowedCreateDegradationRequestDataAttributesStatusEnumValues = []CreateDegradationRequestDataAttributesStatus{
	CREATEDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_INVESTIGATING,
	CREATEDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_IDENTIFIED,
	CREATEDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_MONITORING,
	CREATEDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_RESOLVED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateDegradationRequestDataAttributesStatus) GetAllowedValues() []CreateDegradationRequestDataAttributesStatus {
	return allowedCreateDegradationRequestDataAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateDegradationRequestDataAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateDegradationRequestDataAttributesStatus(value)
	return nil
}

// NewCreateDegradationRequestDataAttributesStatusFromValue returns a pointer to a valid CreateDegradationRequestDataAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateDegradationRequestDataAttributesStatusFromValue(v string) (*CreateDegradationRequestDataAttributesStatus, error) {
	ev := CreateDegradationRequestDataAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateDegradationRequestDataAttributesStatus: valid values are %v", v, allowedCreateDegradationRequestDataAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateDegradationRequestDataAttributesStatus) IsValid() bool {
	for _, existing := range allowedCreateDegradationRequestDataAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateDegradationRequestDataAttributesStatus value.
func (v CreateDegradationRequestDataAttributesStatus) Ptr() *CreateDegradationRequestDataAttributesStatus {
	return &v
}
