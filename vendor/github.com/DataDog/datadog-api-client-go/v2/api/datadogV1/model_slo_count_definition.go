// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOCountDefinition - A count-based (metric) SLI specification, composed of three parts: the good events formula,
// the bad or total events formula, and the underlying queries.
// Exactly one of `total_events_formula` or `bad_events_formula` must be provided.
type SLOCountDefinition struct {
	SLOCountDefinitionWithTotalEventsFormula *SLOCountDefinitionWithTotalEventsFormula
	SLOCountDefinitionWithBadEventsFormula   *SLOCountDefinitionWithBadEventsFormula

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SLOCountDefinitionWithTotalEventsFormulaAsSLOCountDefinition is a convenience function that returns SLOCountDefinitionWithTotalEventsFormula wrapped in SLOCountDefinition.
func SLOCountDefinitionWithTotalEventsFormulaAsSLOCountDefinition(v *SLOCountDefinitionWithTotalEventsFormula) SLOCountDefinition {
	return SLOCountDefinition{SLOCountDefinitionWithTotalEventsFormula: v}
}

// SLOCountDefinitionWithBadEventsFormulaAsSLOCountDefinition is a convenience function that returns SLOCountDefinitionWithBadEventsFormula wrapped in SLOCountDefinition.
func SLOCountDefinitionWithBadEventsFormulaAsSLOCountDefinition(v *SLOCountDefinitionWithBadEventsFormula) SLOCountDefinition {
	return SLOCountDefinition{SLOCountDefinitionWithBadEventsFormula: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SLOCountDefinition) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SLOCountDefinitionWithTotalEventsFormula
	err = datadog.Unmarshal(data, &obj.SLOCountDefinitionWithTotalEventsFormula)
	if err == nil {
		if obj.SLOCountDefinitionWithTotalEventsFormula != nil && obj.SLOCountDefinitionWithTotalEventsFormula.UnparsedObject == nil {
			jsonSLOCountDefinitionWithTotalEventsFormula, _ := datadog.Marshal(obj.SLOCountDefinitionWithTotalEventsFormula)
			if string(jsonSLOCountDefinitionWithTotalEventsFormula) == "{}" { // empty struct
				obj.SLOCountDefinitionWithTotalEventsFormula = nil
			} else {
				match++
			}
		} else {
			obj.SLOCountDefinitionWithTotalEventsFormula = nil
		}
	} else {
		obj.SLOCountDefinitionWithTotalEventsFormula = nil
	}

	// try to unmarshal data into SLOCountDefinitionWithBadEventsFormula
	err = datadog.Unmarshal(data, &obj.SLOCountDefinitionWithBadEventsFormula)
	if err == nil {
		if obj.SLOCountDefinitionWithBadEventsFormula != nil && obj.SLOCountDefinitionWithBadEventsFormula.UnparsedObject == nil {
			jsonSLOCountDefinitionWithBadEventsFormula, _ := datadog.Marshal(obj.SLOCountDefinitionWithBadEventsFormula)
			if string(jsonSLOCountDefinitionWithBadEventsFormula) == "{}" { // empty struct
				obj.SLOCountDefinitionWithBadEventsFormula = nil
			} else {
				match++
			}
		} else {
			obj.SLOCountDefinitionWithBadEventsFormula = nil
		}
	} else {
		obj.SLOCountDefinitionWithBadEventsFormula = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SLOCountDefinitionWithTotalEventsFormula = nil
		obj.SLOCountDefinitionWithBadEventsFormula = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SLOCountDefinition) MarshalJSON() ([]byte, error) {
	if obj.SLOCountDefinitionWithTotalEventsFormula != nil {
		return datadog.Marshal(&obj.SLOCountDefinitionWithTotalEventsFormula)
	}

	if obj.SLOCountDefinitionWithBadEventsFormula != nil {
		return datadog.Marshal(&obj.SLOCountDefinitionWithBadEventsFormula)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SLOCountDefinition) GetActualInstance() interface{} {
	if obj.SLOCountDefinitionWithTotalEventsFormula != nil {
		return obj.SLOCountDefinitionWithTotalEventsFormula
	}

	if obj.SLOCountDefinitionWithBadEventsFormula != nil {
		return obj.SLOCountDefinitionWithBadEventsFormula
	}

	// all schemas are nil
	return nil
}
