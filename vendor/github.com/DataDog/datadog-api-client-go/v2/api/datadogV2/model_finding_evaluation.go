// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
=======
	"fmt"

	"github.com/goccy/go-json"
>>>>>>> main
)

// FindingEvaluation The evaluation of the finding.
type FindingEvaluation string

// List of FindingEvaluation.
const (
	FINDINGEVALUATION_PASS FindingEvaluation = "pass"
	FINDINGEVALUATION_FAIL FindingEvaluation = "fail"
)

var allowedFindingEvaluationEnumValues = []FindingEvaluation{
	FINDINGEVALUATION_PASS,
	FINDINGEVALUATION_FAIL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FindingEvaluation) GetAllowedValues() []FindingEvaluation {
	return allowedFindingEvaluationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FindingEvaluation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FindingEvaluation(value)
	return nil
}

// NewFindingEvaluationFromValue returns a pointer to a valid FindingEvaluation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFindingEvaluationFromValue(v string) (*FindingEvaluation, error) {
	ev := FindingEvaluation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FindingEvaluation: valid values are %v", v, allowedFindingEvaluationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FindingEvaluation) IsValid() bool {
	for _, existing := range allowedFindingEvaluationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FindingEvaluation value.
func (v FindingEvaluation) Ptr() *FindingEvaluation {
	return &v
}
<<<<<<< HEAD

// NullableFindingEvaluation handles when a null is used for FindingEvaluation.
type NullableFindingEvaluation struct {
	value *FindingEvaluation
	isSet bool
}

// Get returns the associated value.
func (v NullableFindingEvaluation) Get() *FindingEvaluation {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableFindingEvaluation) Set(val *FindingEvaluation) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableFindingEvaluation) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableFindingEvaluation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableFindingEvaluation initializes the struct as if Set has been called.
func NewNullableFindingEvaluation(val *FindingEvaluation) *NullableFindingEvaluation {
	return &NullableFindingEvaluation{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableFindingEvaluation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableFindingEvaluation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
