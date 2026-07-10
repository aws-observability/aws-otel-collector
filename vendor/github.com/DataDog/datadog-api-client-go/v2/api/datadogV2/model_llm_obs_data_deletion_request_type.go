// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDataDeletionRequestType Resource type for an LLM Observability data deletion request.
type LLMObsDataDeletionRequestType string

// List of LLMObsDataDeletionRequestType.
const (
	LLMOBSDATADELETIONREQUESTTYPE_CREATE_DELETION_REQ LLMObsDataDeletionRequestType = "create_deletion_req"
)

var allowedLLMObsDataDeletionRequestTypeEnumValues = []LLMObsDataDeletionRequestType{
	LLMOBSDATADELETIONREQUESTTYPE_CREATE_DELETION_REQ,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsDataDeletionRequestType) GetAllowedValues() []LLMObsDataDeletionRequestType {
	return allowedLLMObsDataDeletionRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsDataDeletionRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsDataDeletionRequestType(value)
	return nil
}

// NewLLMObsDataDeletionRequestTypeFromValue returns a pointer to a valid LLMObsDataDeletionRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsDataDeletionRequestTypeFromValue(v string) (*LLMObsDataDeletionRequestType, error) {
	ev := LLMObsDataDeletionRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsDataDeletionRequestType: valid values are %v", v, allowedLLMObsDataDeletionRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsDataDeletionRequestType) IsValid() bool {
	for _, existing := range allowedLLMObsDataDeletionRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsDataDeletionRequestType value.
func (v LLMObsDataDeletionRequestType) Ptr() *LLMObsDataDeletionRequestType {
	return &v
}
