// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData Workflows Webhook handle data from a response.
type MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData struct {
	// Workflows Webhook handle attributes.
	Attributes MicrosoftTeamsWorkflowsWebhookHandleAttributes `json:"attributes"`
	// Specifies the Workflows webhook handle resource type.
	Type MicrosoftTeamsWorkflowsWebhookHandleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData instantiates a new MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData(attributes MicrosoftTeamsWorkflowsWebhookHandleAttributes, typeVar MicrosoftTeamsWorkflowsWebhookHandleType) *MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData {
	this := MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewMicrosoftTeamsUpdateWorkflowsWebhookHandleRequestDataWithDefaults instantiates a new MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMicrosoftTeamsUpdateWorkflowsWebhookHandleRequestDataWithDefaults() *MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData {
	this := MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData{}
	var typeVar MicrosoftTeamsWorkflowsWebhookHandleType = MICROSOFTTEAMSWORKFLOWSWEBHOOKHANDLETYPE_WORKFLOWS_WEBHOOK_HANDLE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData) GetAttributes() MicrosoftTeamsWorkflowsWebhookHandleAttributes {
	if o == nil {
		var ret MicrosoftTeamsWorkflowsWebhookHandleAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData) GetAttributesOk() (*MicrosoftTeamsWorkflowsWebhookHandleAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData) SetAttributes(v MicrosoftTeamsWorkflowsWebhookHandleAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData) GetType() MicrosoftTeamsWorkflowsWebhookHandleType {
	if o == nil {
		var ret MicrosoftTeamsWorkflowsWebhookHandleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData) GetTypeOk() (*MicrosoftTeamsWorkflowsWebhookHandleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData) SetType(v MicrosoftTeamsWorkflowsWebhookHandleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MicrosoftTeamsUpdateWorkflowsWebhookHandleRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *MicrosoftTeamsWorkflowsWebhookHandleAttributes `json:"attributes"`
		Type       *MicrosoftTeamsWorkflowsWebhookHandleType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
