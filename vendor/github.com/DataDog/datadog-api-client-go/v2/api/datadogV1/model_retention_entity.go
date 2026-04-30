// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionEntity Entity to track for retention.
type RetentionEntity string

// List of RetentionEntity.
const (
	RETENTIONENTITY_USER_ID    RetentionEntity = "@usr.id"
	RETENTIONENTITY_ACCOUNT_ID RetentionEntity = "@account.id"
)

var allowedRetentionEntityEnumValues = []RetentionEntity{
	RETENTIONENTITY_USER_ID,
	RETENTIONENTITY_ACCOUNT_ID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionEntity) GetAllowedValues() []RetentionEntity {
	return allowedRetentionEntityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionEntity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionEntity(value)
	return nil
}

// NewRetentionEntityFromValue returns a pointer to a valid RetentionEntity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionEntityFromValue(v string) (*RetentionEntity, error) {
	ev := RetentionEntity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionEntity: valid values are %v", v, allowedRetentionEntityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionEntity) IsValid() bool {
	for _, existing := range allowedRetentionEntityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionEntity value.
func (v RetentionEntity) Ptr() *RetentionEntity {
	return &v
}
