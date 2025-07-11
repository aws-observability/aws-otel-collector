// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SendTeamsMessageAction Sends a message to a Microsoft Teams channel.
type SendTeamsMessageAction struct {
	// The channel ID.
	Channel string `json:"channel"`
	// The team ID.
	Team string `json:"team"`
	// The tenant ID.
	Tenant string `json:"tenant"`
	// Indicates that the action is a send Microsoft Teams message action.
	Type SendTeamsMessageActionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSendTeamsMessageAction instantiates a new SendTeamsMessageAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSendTeamsMessageAction(channel string, team string, tenant string, typeVar SendTeamsMessageActionType) *SendTeamsMessageAction {
	this := SendTeamsMessageAction{}
	this.Channel = channel
	this.Team = team
	this.Tenant = tenant
	this.Type = typeVar
	return &this
}

// NewSendTeamsMessageActionWithDefaults instantiates a new SendTeamsMessageAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSendTeamsMessageActionWithDefaults() *SendTeamsMessageAction {
	this := SendTeamsMessageAction{}
	var typeVar SendTeamsMessageActionType = SENDTEAMSMESSAGEACTIONTYPE_SEND_TEAMS_MESSAGE
	this.Type = typeVar
	return &this
}

// GetChannel returns the Channel field value.
func (o *SendTeamsMessageAction) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *SendTeamsMessageAction) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value.
func (o *SendTeamsMessageAction) SetChannel(v string) {
	o.Channel = v
}

// GetTeam returns the Team field value.
func (o *SendTeamsMessageAction) GetTeam() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Team
}

// GetTeamOk returns a tuple with the Team field value
// and a boolean to check if the value has been set.
func (o *SendTeamsMessageAction) GetTeamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Team, true
}

// SetTeam sets field value.
func (o *SendTeamsMessageAction) SetTeam(v string) {
	o.Team = v
}

// GetTenant returns the Tenant field value.
func (o *SendTeamsMessageAction) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *SendTeamsMessageAction) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value.
func (o *SendTeamsMessageAction) SetTenant(v string) {
	o.Tenant = v
}

// GetType returns the Type field value.
func (o *SendTeamsMessageAction) GetType() SendTeamsMessageActionType {
	if o == nil {
		var ret SendTeamsMessageActionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SendTeamsMessageAction) GetTypeOk() (*SendTeamsMessageActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SendTeamsMessageAction) SetType(v SendTeamsMessageActionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SendTeamsMessageAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["channel"] = o.Channel
	toSerialize["team"] = o.Team
	toSerialize["tenant"] = o.Tenant
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SendTeamsMessageAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Channel *string                     `json:"channel"`
		Team    *string                     `json:"team"`
		Tenant  *string                     `json:"tenant"`
		Type    *SendTeamsMessageActionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Channel == nil {
		return fmt.Errorf("required field channel missing")
	}
	if all.Team == nil {
		return fmt.Errorf("required field team missing")
	}
	if all.Tenant == nil {
		return fmt.Errorf("required field tenant missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"channel", "team", "tenant", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Channel = *all.Channel
	o.Team = *all.Team
	o.Tenant = *all.Tenant
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
