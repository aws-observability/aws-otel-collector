// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsContentBlockType Discriminator for a single `display_block` content block. Adding a
// variant requires coordinated changes in the frontend renderer.
type LLMObsContentBlockType string

// List of LLMObsContentBlockType.
const (
	LLMOBSCONTENTBLOCKTYPE_MARKDOWN     LLMObsContentBlockType = "markdown"
	LLMOBSCONTENTBLOCKTYPE_HEADER       LLMObsContentBlockType = "header"
	LLMOBSCONTENTBLOCKTYPE_TEXT         LLMObsContentBlockType = "text"
	LLMOBSCONTENTBLOCKTYPE_JSON         LLMObsContentBlockType = "json"
	LLMOBSCONTENTBLOCKTYPE_IMAGE        LLMObsContentBlockType = "image"
	LLMOBSCONTENTBLOCKTYPE_WIDGET       LLMObsContentBlockType = "widget"
	LLMOBSCONTENTBLOCKTYPE_LLMOBS_TRACE LLMObsContentBlockType = "llmobs_trace"
)

var allowedLLMObsContentBlockTypeEnumValues = []LLMObsContentBlockType{
	LLMOBSCONTENTBLOCKTYPE_MARKDOWN,
	LLMOBSCONTENTBLOCKTYPE_HEADER,
	LLMOBSCONTENTBLOCKTYPE_TEXT,
	LLMOBSCONTENTBLOCKTYPE_JSON,
	LLMOBSCONTENTBLOCKTYPE_IMAGE,
	LLMOBSCONTENTBLOCKTYPE_WIDGET,
	LLMOBSCONTENTBLOCKTYPE_LLMOBS_TRACE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsContentBlockType) GetAllowedValues() []LLMObsContentBlockType {
	return allowedLLMObsContentBlockTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsContentBlockType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsContentBlockType(value)
	return nil
}

// NewLLMObsContentBlockTypeFromValue returns a pointer to a valid LLMObsContentBlockType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsContentBlockTypeFromValue(v string) (*LLMObsContentBlockType, error) {
	ev := LLMObsContentBlockType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsContentBlockType: valid values are %v", v, allowedLLMObsContentBlockTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsContentBlockType) IsValid() bool {
	for _, existing := range allowedLLMObsContentBlockTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsContentBlockType value.
func (v LLMObsContentBlockType) Ptr() *LLMObsContentBlockType {
	return &v
}
