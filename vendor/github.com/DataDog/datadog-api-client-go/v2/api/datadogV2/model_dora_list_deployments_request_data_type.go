// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORAListDeploymentsRequestDataType The definition of `DORAListDeploymentsRequestDataType` object.
type DORAListDeploymentsRequestDataType string

// List of DORAListDeploymentsRequestDataType.
const (
	DORALISTDEPLOYMENTSREQUESTDATATYPE_DORA_DEPLOYMENTS_LIST_REQUEST DORAListDeploymentsRequestDataType = "dora_deployments_list_request"
)

var allowedDORAListDeploymentsRequestDataTypeEnumValues = []DORAListDeploymentsRequestDataType{
	DORALISTDEPLOYMENTSREQUESTDATATYPE_DORA_DEPLOYMENTS_LIST_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DORAListDeploymentsRequestDataType) GetAllowedValues() []DORAListDeploymentsRequestDataType {
	return allowedDORAListDeploymentsRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DORAListDeploymentsRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DORAListDeploymentsRequestDataType(value)
	return nil
}

// NewDORAListDeploymentsRequestDataTypeFromValue returns a pointer to a valid DORAListDeploymentsRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDORAListDeploymentsRequestDataTypeFromValue(v string) (*DORAListDeploymentsRequestDataType, error) {
	ev := DORAListDeploymentsRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DORAListDeploymentsRequestDataType: valid values are %v", v, allowedDORAListDeploymentsRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DORAListDeploymentsRequestDataType) IsValid() bool {
	for _, existing := range allowedDORAListDeploymentsRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DORAListDeploymentsRequestDataType value.
func (v DORAListDeploymentsRequestDataType) Ptr() *DORAListDeploymentsRequestDataType {
	return &v
}
