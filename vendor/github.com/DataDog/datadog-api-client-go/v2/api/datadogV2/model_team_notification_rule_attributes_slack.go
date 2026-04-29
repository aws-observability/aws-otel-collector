// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamNotificationRuleAttributesSlack Slack notification settings for the team
type TeamNotificationRuleAttributesSlack struct {
	// Channel for Slack notification
	Channel *string `json:"channel,omitempty"`
	// Workspace for Slack notification
	Workspace *string `json:"workspace,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamNotificationRuleAttributesSlack instantiates a new TeamNotificationRuleAttributesSlack object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamNotificationRuleAttributesSlack() *TeamNotificationRuleAttributesSlack {
	this := TeamNotificationRuleAttributesSlack{}
	return &this
}

// NewTeamNotificationRuleAttributesSlackWithDefaults instantiates a new TeamNotificationRuleAttributesSlack object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamNotificationRuleAttributesSlackWithDefaults() *TeamNotificationRuleAttributesSlack {
	this := TeamNotificationRuleAttributesSlack{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *TeamNotificationRuleAttributesSlack) GetChannel() string {
	if o == nil || o.Channel == nil {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRuleAttributesSlack) GetChannelOk() (*string, bool) {
	if o == nil || o.Channel == nil {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *TeamNotificationRuleAttributesSlack) HasChannel() bool {
	return o != nil && o.Channel != nil
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *TeamNotificationRuleAttributesSlack) SetChannel(v string) {
	o.Channel = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *TeamNotificationRuleAttributesSlack) GetWorkspace() string {
	if o == nil || o.Workspace == nil {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRuleAttributesSlack) GetWorkspaceOk() (*string, bool) {
	if o == nil || o.Workspace == nil {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *TeamNotificationRuleAttributesSlack) HasWorkspace() bool {
	return o != nil && o.Workspace != nil
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *TeamNotificationRuleAttributesSlack) SetWorkspace(v string) {
	o.Workspace = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamNotificationRuleAttributesSlack) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Channel != nil {
		toSerialize["channel"] = o.Channel
	}
	if o.Workspace != nil {
		toSerialize["workspace"] = o.Workspace
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamNotificationRuleAttributesSlack) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Channel   *string `json:"channel,omitempty"`
		Workspace *string `json:"workspace,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"channel", "workspace"})
	} else {
		return err
	}
	o.Channel = all.Channel
	o.Workspace = all.Workspace

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
