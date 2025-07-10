// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SendSlackMessageAction Sends a message to a Slack channel.
type SendSlackMessageAction struct {
	// The channel ID.
	Channel string `json:"channel"`
	// Indicates that the action is a send Slack message action.
	Type SendSlackMessageActionType `json:"type"`
	// The workspace ID.
	Workspace string `json:"workspace"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSendSlackMessageAction instantiates a new SendSlackMessageAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSendSlackMessageAction(channel string, typeVar SendSlackMessageActionType, workspace string) *SendSlackMessageAction {
	this := SendSlackMessageAction{}
	this.Channel = channel
	this.Type = typeVar
	this.Workspace = workspace
	return &this
}

// NewSendSlackMessageActionWithDefaults instantiates a new SendSlackMessageAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSendSlackMessageActionWithDefaults() *SendSlackMessageAction {
	this := SendSlackMessageAction{}
	var typeVar SendSlackMessageActionType = SENDSLACKMESSAGEACTIONTYPE_SEND_SLACK_MESSAGE
	this.Type = typeVar
	return &this
}

// GetChannel returns the Channel field value.
func (o *SendSlackMessageAction) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *SendSlackMessageAction) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value.
func (o *SendSlackMessageAction) SetChannel(v string) {
	o.Channel = v
}

// GetType returns the Type field value.
func (o *SendSlackMessageAction) GetType() SendSlackMessageActionType {
	if o == nil {
		var ret SendSlackMessageActionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SendSlackMessageAction) GetTypeOk() (*SendSlackMessageActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SendSlackMessageAction) SetType(v SendSlackMessageActionType) {
	o.Type = v
}

// GetWorkspace returns the Workspace field value.
func (o *SendSlackMessageAction) GetWorkspace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value
// and a boolean to check if the value has been set.
func (o *SendSlackMessageAction) GetWorkspaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workspace, true
}

// SetWorkspace sets field value.
func (o *SendSlackMessageAction) SetWorkspace(v string) {
	o.Workspace = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SendSlackMessageAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["channel"] = o.Channel
	toSerialize["type"] = o.Type
	toSerialize["workspace"] = o.Workspace

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SendSlackMessageAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Channel   *string                     `json:"channel"`
		Type      *SendSlackMessageActionType `json:"type"`
		Workspace *string                     `json:"workspace"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Channel == nil {
		return fmt.Errorf("required field channel missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Workspace == nil {
		return fmt.Errorf("required field workspace missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"channel", "type", "workspace"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Channel = *all.Channel
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Workspace = *all.Workspace

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
