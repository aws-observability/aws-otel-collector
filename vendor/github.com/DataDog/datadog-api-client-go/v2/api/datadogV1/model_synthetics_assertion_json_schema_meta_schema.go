// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionJSONSchemaMetaSchema The JSON Schema meta-schema version used in the assertion.
type SyntheticsAssertionJSONSchemaMetaSchema string

// List of SyntheticsAssertionJSONSchemaMetaSchema.
const (
	SYNTHETICSASSERTIONJSONSCHEMAMETASCHEMA_DRAFT_07 SyntheticsAssertionJSONSchemaMetaSchema = "draft-07"
	SYNTHETICSASSERTIONJSONSCHEMAMETASCHEMA_DRAFT_06 SyntheticsAssertionJSONSchemaMetaSchema = "draft-06"
)

var allowedSyntheticsAssertionJSONSchemaMetaSchemaEnumValues = []SyntheticsAssertionJSONSchemaMetaSchema{
	SYNTHETICSASSERTIONJSONSCHEMAMETASCHEMA_DRAFT_07,
	SYNTHETICSASSERTIONJSONSCHEMAMETASCHEMA_DRAFT_06,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsAssertionJSONSchemaMetaSchema) GetAllowedValues() []SyntheticsAssertionJSONSchemaMetaSchema {
	return allowedSyntheticsAssertionJSONSchemaMetaSchemaEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsAssertionJSONSchemaMetaSchema) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAssertionJSONSchemaMetaSchema(value)
	return nil
}

// NewSyntheticsAssertionJSONSchemaMetaSchemaFromValue returns a pointer to a valid SyntheticsAssertionJSONSchemaMetaSchema
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsAssertionJSONSchemaMetaSchemaFromValue(v string) (*SyntheticsAssertionJSONSchemaMetaSchema, error) {
	ev := SyntheticsAssertionJSONSchemaMetaSchema(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsAssertionJSONSchemaMetaSchema: valid values are %v", v, allowedSyntheticsAssertionJSONSchemaMetaSchemaEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsAssertionJSONSchemaMetaSchema) IsValid() bool {
	for _, existing := range allowedSyntheticsAssertionJSONSchemaMetaSchemaEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAssertionJSONSchemaMetaSchema value.
func (v SyntheticsAssertionJSONSchemaMetaSchema) Ptr() *SyntheticsAssertionJSONSchemaMetaSchema {
	return &v
}
