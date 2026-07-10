// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDataDeletionResponseType Resource type for an LLM Observability data deletion response.
type LLMObsDataDeletionResponseType string

// List of LLMObsDataDeletionResponseType.
const (
	LLMOBSDATADELETIONRESPONSETYPE_DELETION_REQUEST LLMObsDataDeletionResponseType = "deletion_request"
)

var allowedLLMObsDataDeletionResponseTypeEnumValues = []LLMObsDataDeletionResponseType{
	LLMOBSDATADELETIONRESPONSETYPE_DELETION_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsDataDeletionResponseType) GetAllowedValues() []LLMObsDataDeletionResponseType {
	return allowedLLMObsDataDeletionResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsDataDeletionResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsDataDeletionResponseType(value)
	return nil
}

// NewLLMObsDataDeletionResponseTypeFromValue returns a pointer to a valid LLMObsDataDeletionResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsDataDeletionResponseTypeFromValue(v string) (*LLMObsDataDeletionResponseType, error) {
	ev := LLMObsDataDeletionResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsDataDeletionResponseType: valid values are %v", v, allowedLLMObsDataDeletionResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsDataDeletionResponseType) IsValid() bool {
	for _, existing := range allowedLLMObsDataDeletionResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsDataDeletionResponseType value.
func (v LLMObsDataDeletionResponseType) Ptr() *LLMObsDataDeletionResponseType {
	return &v
}
