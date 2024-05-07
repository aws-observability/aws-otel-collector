// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConfluentResourceType The JSON:API type for this request.
type ConfluentResourceType string

// List of ConfluentResourceType.
const (
	CONFLUENTRESOURCETYPE_CONFLUENT_CLOUD_RESOURCES ConfluentResourceType = "confluent-cloud-resources"
)

var allowedConfluentResourceTypeEnumValues = []ConfluentResourceType{
	CONFLUENTRESOURCETYPE_CONFLUENT_CLOUD_RESOURCES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ConfluentResourceType) GetAllowedValues() []ConfluentResourceType {
	return allowedConfluentResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ConfluentResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ConfluentResourceType(value)
	return nil
}

// NewConfluentResourceTypeFromValue returns a pointer to a valid ConfluentResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewConfluentResourceTypeFromValue(v string) (*ConfluentResourceType, error) {
	ev := ConfluentResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ConfluentResourceType: valid values are %v", v, allowedConfluentResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ConfluentResourceType) IsValid() bool {
	for _, existing := range allowedConfluentResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConfluentResourceType value.
func (v ConfluentResourceType) Ptr() *ConfluentResourceType {
	return &v
}
