// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftTeamsWorkflowsWebhookHandleType Specifies the Workflows webhook handle resource type.
type MicrosoftTeamsWorkflowsWebhookHandleType string

// List of MicrosoftTeamsWorkflowsWebhookHandleType.
const (
	MICROSOFTTEAMSWORKFLOWSWEBHOOKHANDLETYPE_WORKFLOWS_WEBHOOK_HANDLE MicrosoftTeamsWorkflowsWebhookHandleType = "workflows-webhook-handle"
)

var allowedMicrosoftTeamsWorkflowsWebhookHandleTypeEnumValues = []MicrosoftTeamsWorkflowsWebhookHandleType{
	MICROSOFTTEAMSWORKFLOWSWEBHOOKHANDLETYPE_WORKFLOWS_WEBHOOK_HANDLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MicrosoftTeamsWorkflowsWebhookHandleType) GetAllowedValues() []MicrosoftTeamsWorkflowsWebhookHandleType {
	return allowedMicrosoftTeamsWorkflowsWebhookHandleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MicrosoftTeamsWorkflowsWebhookHandleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MicrosoftTeamsWorkflowsWebhookHandleType(value)
	return nil
}

// NewMicrosoftTeamsWorkflowsWebhookHandleTypeFromValue returns a pointer to a valid MicrosoftTeamsWorkflowsWebhookHandleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMicrosoftTeamsWorkflowsWebhookHandleTypeFromValue(v string) (*MicrosoftTeamsWorkflowsWebhookHandleType, error) {
	ev := MicrosoftTeamsWorkflowsWebhookHandleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MicrosoftTeamsWorkflowsWebhookHandleType: valid values are %v", v, allowedMicrosoftTeamsWorkflowsWebhookHandleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MicrosoftTeamsWorkflowsWebhookHandleType) IsValid() bool {
	for _, existing := range allowedMicrosoftTeamsWorkflowsWebhookHandleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MicrosoftTeamsWorkflowsWebhookHandleType value.
func (v MicrosoftTeamsWorkflowsWebhookHandleType) Ptr() *MicrosoftTeamsWorkflowsWebhookHandleType {
	return &v
}
