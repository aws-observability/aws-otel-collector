// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertion - Object describing the assertions type, their associated operator,
// which property they apply, and upon which target.
type SyntheticsAssertion struct {
	SyntheticsAssertionTarget           *SyntheticsAssertionTarget
	SyntheticsAssertionBodyHashTarget   *SyntheticsAssertionBodyHashTarget
	SyntheticsAssertionJSONPathTarget   *SyntheticsAssertionJSONPathTarget
	SyntheticsAssertionJSONSchemaTarget *SyntheticsAssertionJSONSchemaTarget
	SyntheticsAssertionXPathTarget      *SyntheticsAssertionXPathTarget
	SyntheticsAssertionJavascript       *SyntheticsAssertionJavascript

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SyntheticsAssertionTargetAsSyntheticsAssertion is a convenience function that returns SyntheticsAssertionTarget wrapped in SyntheticsAssertion.
func SyntheticsAssertionTargetAsSyntheticsAssertion(v *SyntheticsAssertionTarget) SyntheticsAssertion {
	return SyntheticsAssertion{SyntheticsAssertionTarget: v}
}

// SyntheticsAssertionBodyHashTargetAsSyntheticsAssertion is a convenience function that returns SyntheticsAssertionBodyHashTarget wrapped in SyntheticsAssertion.
func SyntheticsAssertionBodyHashTargetAsSyntheticsAssertion(v *SyntheticsAssertionBodyHashTarget) SyntheticsAssertion {
	return SyntheticsAssertion{SyntheticsAssertionBodyHashTarget: v}
}

// SyntheticsAssertionJSONPathTargetAsSyntheticsAssertion is a convenience function that returns SyntheticsAssertionJSONPathTarget wrapped in SyntheticsAssertion.
func SyntheticsAssertionJSONPathTargetAsSyntheticsAssertion(v *SyntheticsAssertionJSONPathTarget) SyntheticsAssertion {
	return SyntheticsAssertion{SyntheticsAssertionJSONPathTarget: v}
}

// SyntheticsAssertionJSONSchemaTargetAsSyntheticsAssertion is a convenience function that returns SyntheticsAssertionJSONSchemaTarget wrapped in SyntheticsAssertion.
func SyntheticsAssertionJSONSchemaTargetAsSyntheticsAssertion(v *SyntheticsAssertionJSONSchemaTarget) SyntheticsAssertion {
	return SyntheticsAssertion{SyntheticsAssertionJSONSchemaTarget: v}
}

// SyntheticsAssertionXPathTargetAsSyntheticsAssertion is a convenience function that returns SyntheticsAssertionXPathTarget wrapped in SyntheticsAssertion.
func SyntheticsAssertionXPathTargetAsSyntheticsAssertion(v *SyntheticsAssertionXPathTarget) SyntheticsAssertion {
	return SyntheticsAssertion{SyntheticsAssertionXPathTarget: v}
}

// SyntheticsAssertionJavascriptAsSyntheticsAssertion is a convenience function that returns SyntheticsAssertionJavascript wrapped in SyntheticsAssertion.
func SyntheticsAssertionJavascriptAsSyntheticsAssertion(v *SyntheticsAssertionJavascript) SyntheticsAssertion {
	return SyntheticsAssertion{SyntheticsAssertionJavascript: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SyntheticsAssertion) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SyntheticsAssertionTarget
	err = datadog.Unmarshal(data, &obj.SyntheticsAssertionTarget)
	if err == nil {
		if obj.SyntheticsAssertionTarget != nil && obj.SyntheticsAssertionTarget.UnparsedObject == nil {
			jsonSyntheticsAssertionTarget, _ := datadog.Marshal(obj.SyntheticsAssertionTarget)
			if string(jsonSyntheticsAssertionTarget) == "{}" { // empty struct
				obj.SyntheticsAssertionTarget = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsAssertionTarget = nil
		}
	} else {
		obj.SyntheticsAssertionTarget = nil
	}

	// try to unmarshal data into SyntheticsAssertionBodyHashTarget
	err = datadog.Unmarshal(data, &obj.SyntheticsAssertionBodyHashTarget)
	if err == nil {
		if obj.SyntheticsAssertionBodyHashTarget != nil && obj.SyntheticsAssertionBodyHashTarget.UnparsedObject == nil {
			jsonSyntheticsAssertionBodyHashTarget, _ := datadog.Marshal(obj.SyntheticsAssertionBodyHashTarget)
			if string(jsonSyntheticsAssertionBodyHashTarget) == "{}" { // empty struct
				obj.SyntheticsAssertionBodyHashTarget = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsAssertionBodyHashTarget = nil
		}
	} else {
		obj.SyntheticsAssertionBodyHashTarget = nil
	}

	// try to unmarshal data into SyntheticsAssertionJSONPathTarget
	err = datadog.Unmarshal(data, &obj.SyntheticsAssertionJSONPathTarget)
	if err == nil {
		if obj.SyntheticsAssertionJSONPathTarget != nil && obj.SyntheticsAssertionJSONPathTarget.UnparsedObject == nil {
			jsonSyntheticsAssertionJSONPathTarget, _ := datadog.Marshal(obj.SyntheticsAssertionJSONPathTarget)
			if string(jsonSyntheticsAssertionJSONPathTarget) == "{}" { // empty struct
				obj.SyntheticsAssertionJSONPathTarget = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsAssertionJSONPathTarget = nil
		}
	} else {
		obj.SyntheticsAssertionJSONPathTarget = nil
	}

	// try to unmarshal data into SyntheticsAssertionJSONSchemaTarget
	err = datadog.Unmarshal(data, &obj.SyntheticsAssertionJSONSchemaTarget)
	if err == nil {
		if obj.SyntheticsAssertionJSONSchemaTarget != nil && obj.SyntheticsAssertionJSONSchemaTarget.UnparsedObject == nil {
			jsonSyntheticsAssertionJSONSchemaTarget, _ := datadog.Marshal(obj.SyntheticsAssertionJSONSchemaTarget)
			if string(jsonSyntheticsAssertionJSONSchemaTarget) == "{}" { // empty struct
				obj.SyntheticsAssertionJSONSchemaTarget = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsAssertionJSONSchemaTarget = nil
		}
	} else {
		obj.SyntheticsAssertionJSONSchemaTarget = nil
	}

	// try to unmarshal data into SyntheticsAssertionXPathTarget
	err = datadog.Unmarshal(data, &obj.SyntheticsAssertionXPathTarget)
	if err == nil {
		if obj.SyntheticsAssertionXPathTarget != nil && obj.SyntheticsAssertionXPathTarget.UnparsedObject == nil {
			jsonSyntheticsAssertionXPathTarget, _ := datadog.Marshal(obj.SyntheticsAssertionXPathTarget)
			if string(jsonSyntheticsAssertionXPathTarget) == "{}" { // empty struct
				obj.SyntheticsAssertionXPathTarget = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsAssertionXPathTarget = nil
		}
	} else {
		obj.SyntheticsAssertionXPathTarget = nil
	}

	// try to unmarshal data into SyntheticsAssertionJavascript
	err = datadog.Unmarshal(data, &obj.SyntheticsAssertionJavascript)
	if err == nil {
		if obj.SyntheticsAssertionJavascript != nil && obj.SyntheticsAssertionJavascript.UnparsedObject == nil {
			jsonSyntheticsAssertionJavascript, _ := datadog.Marshal(obj.SyntheticsAssertionJavascript)
			if string(jsonSyntheticsAssertionJavascript) == "{}" { // empty struct
				obj.SyntheticsAssertionJavascript = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsAssertionJavascript = nil
		}
	} else {
		obj.SyntheticsAssertionJavascript = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SyntheticsAssertionTarget = nil
		obj.SyntheticsAssertionBodyHashTarget = nil
		obj.SyntheticsAssertionJSONPathTarget = nil
		obj.SyntheticsAssertionJSONSchemaTarget = nil
		obj.SyntheticsAssertionXPathTarget = nil
		obj.SyntheticsAssertionJavascript = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SyntheticsAssertion) MarshalJSON() ([]byte, error) {
	if obj.SyntheticsAssertionTarget != nil {
		return datadog.Marshal(&obj.SyntheticsAssertionTarget)
	}

	if obj.SyntheticsAssertionBodyHashTarget != nil {
		return datadog.Marshal(&obj.SyntheticsAssertionBodyHashTarget)
	}

	if obj.SyntheticsAssertionJSONPathTarget != nil {
		return datadog.Marshal(&obj.SyntheticsAssertionJSONPathTarget)
	}

	if obj.SyntheticsAssertionJSONSchemaTarget != nil {
		return datadog.Marshal(&obj.SyntheticsAssertionJSONSchemaTarget)
	}

	if obj.SyntheticsAssertionXPathTarget != nil {
		return datadog.Marshal(&obj.SyntheticsAssertionXPathTarget)
	}

	if obj.SyntheticsAssertionJavascript != nil {
		return datadog.Marshal(&obj.SyntheticsAssertionJavascript)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SyntheticsAssertion) GetActualInstance() interface{} {
	if obj.SyntheticsAssertionTarget != nil {
		return obj.SyntheticsAssertionTarget
	}

	if obj.SyntheticsAssertionBodyHashTarget != nil {
		return obj.SyntheticsAssertionBodyHashTarget
	}

	if obj.SyntheticsAssertionJSONPathTarget != nil {
		return obj.SyntheticsAssertionJSONPathTarget
	}

	if obj.SyntheticsAssertionJSONSchemaTarget != nil {
		return obj.SyntheticsAssertionJSONSchemaTarget
	}

	if obj.SyntheticsAssertionXPathTarget != nil {
		return obj.SyntheticsAssertionXPathTarget
	}

	if obj.SyntheticsAssertionJavascript != nil {
		return obj.SyntheticsAssertionJavascript
	}

	// all schemas are nil
	return nil
}
