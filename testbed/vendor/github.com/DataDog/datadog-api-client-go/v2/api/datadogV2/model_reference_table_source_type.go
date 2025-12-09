// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReferenceTableSourceType The source type for reference table data. Includes all possible source types that can appear in responses.
type ReferenceTableSourceType string

// List of ReferenceTableSourceType.
const (
	REFERENCETABLESOURCETYPE_LOCAL_FILE ReferenceTableSourceType = "LOCAL_FILE"
	REFERENCETABLESOURCETYPE_S3         ReferenceTableSourceType = "S3"
	REFERENCETABLESOURCETYPE_GCS        ReferenceTableSourceType = "GCS"
	REFERENCETABLESOURCETYPE_AZURE      ReferenceTableSourceType = "AZURE"
	REFERENCETABLESOURCETYPE_SERVICENOW ReferenceTableSourceType = "SERVICENOW"
	REFERENCETABLESOURCETYPE_SALESFORCE ReferenceTableSourceType = "SALESFORCE"
	REFERENCETABLESOURCETYPE_DATABRICKS ReferenceTableSourceType = "DATABRICKS"
	REFERENCETABLESOURCETYPE_SNOWFLAKE  ReferenceTableSourceType = "SNOWFLAKE"
)

var allowedReferenceTableSourceTypeEnumValues = []ReferenceTableSourceType{
	REFERENCETABLESOURCETYPE_LOCAL_FILE,
	REFERENCETABLESOURCETYPE_S3,
	REFERENCETABLESOURCETYPE_GCS,
	REFERENCETABLESOURCETYPE_AZURE,
	REFERENCETABLESOURCETYPE_SERVICENOW,
	REFERENCETABLESOURCETYPE_SALESFORCE,
	REFERENCETABLESOURCETYPE_DATABRICKS,
	REFERENCETABLESOURCETYPE_SNOWFLAKE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReferenceTableSourceType) GetAllowedValues() []ReferenceTableSourceType {
	return allowedReferenceTableSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReferenceTableSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReferenceTableSourceType(value)
	return nil
}

// NewReferenceTableSourceTypeFromValue returns a pointer to a valid ReferenceTableSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReferenceTableSourceTypeFromValue(v string) (*ReferenceTableSourceType, error) {
	ev := ReferenceTableSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReferenceTableSourceType: valid values are %v", v, allowedReferenceTableSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReferenceTableSourceType) IsValid() bool {
	for _, existing := range allowedReferenceTableSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReferenceTableSourceType value.
func (v ReferenceTableSourceType) Ptr() *ReferenceTableSourceType {
	return &v
}
