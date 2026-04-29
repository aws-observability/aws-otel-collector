// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsFastTestResultType JSON:API type for a fast test result.
type SyntheticsFastTestResultType string

// List of SyntheticsFastTestResultType.
const (
	SYNTHETICSFASTTESTRESULTTYPE_RESULT SyntheticsFastTestResultType = "result"
)

var allowedSyntheticsFastTestResultTypeEnumValues = []SyntheticsFastTestResultType{
	SYNTHETICSFASTTESTRESULTTYPE_RESULT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsFastTestResultType) GetAllowedValues() []SyntheticsFastTestResultType {
	return allowedSyntheticsFastTestResultTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsFastTestResultType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsFastTestResultType(value)
	return nil
}

// NewSyntheticsFastTestResultTypeFromValue returns a pointer to a valid SyntheticsFastTestResultType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsFastTestResultTypeFromValue(v string) (*SyntheticsFastTestResultType, error) {
	ev := SyntheticsFastTestResultType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsFastTestResultType: valid values are %v", v, allowedSyntheticsFastTestResultTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsFastTestResultType) IsValid() bool {
	for _, existing := range allowedSyntheticsFastTestResultTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsFastTestResultType value.
func (v SyntheticsFastTestResultType) Ptr() *SyntheticsFastTestResultType {
	return &v
}
