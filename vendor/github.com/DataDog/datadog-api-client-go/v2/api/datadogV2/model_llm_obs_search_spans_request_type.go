// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsSearchSpansRequestType Resource type for an LLM Observability spans search request.
type LLMObsSearchSpansRequestType string

// List of LLMObsSearchSpansRequestType.
const (
	LLMOBSSEARCHSPANSREQUESTTYPE_SPANS LLMObsSearchSpansRequestType = "spans"
)

var allowedLLMObsSearchSpansRequestTypeEnumValues = []LLMObsSearchSpansRequestType{
	LLMOBSSEARCHSPANSREQUESTTYPE_SPANS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsSearchSpansRequestType) GetAllowedValues() []LLMObsSearchSpansRequestType {
	return allowedLLMObsSearchSpansRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsSearchSpansRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsSearchSpansRequestType(value)
	return nil
}

// NewLLMObsSearchSpansRequestTypeFromValue returns a pointer to a valid LLMObsSearchSpansRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsSearchSpansRequestTypeFromValue(v string) (*LLMObsSearchSpansRequestType, error) {
	ev := LLMObsSearchSpansRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsSearchSpansRequestType: valid values are %v", v, allowedLLMObsSearchSpansRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsSearchSpansRequestType) IsValid() bool {
	for _, existing := range allowedLLMObsSearchSpansRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsSearchSpansRequestType value.
func (v LLMObsSearchSpansRequestType) Ptr() *LLMObsSearchSpansRequestType {
	return &v
}
