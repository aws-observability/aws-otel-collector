// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftTeamsTenantBasedHandleRequestAttributes Tenant-based handle attributes.
type MicrosoftTeamsTenantBasedHandleRequestAttributes struct {
	// Channel id.
	ChannelId string `json:"channel_id"`
	// Tenant-based handle name.
	Name string `json:"name"`
	// Team id.
	TeamId string `json:"team_id"`
	// Tenant id.
	TenantId string `json:"tenant_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMicrosoftTeamsTenantBasedHandleRequestAttributes instantiates a new MicrosoftTeamsTenantBasedHandleRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMicrosoftTeamsTenantBasedHandleRequestAttributes(channelId string, name string, teamId string, tenantId string) *MicrosoftTeamsTenantBasedHandleRequestAttributes {
	this := MicrosoftTeamsTenantBasedHandleRequestAttributes{}
	this.ChannelId = channelId
	this.Name = name
	this.TeamId = teamId
	this.TenantId = tenantId
	return &this
}

// NewMicrosoftTeamsTenantBasedHandleRequestAttributesWithDefaults instantiates a new MicrosoftTeamsTenantBasedHandleRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMicrosoftTeamsTenantBasedHandleRequestAttributesWithDefaults() *MicrosoftTeamsTenantBasedHandleRequestAttributes {
	this := MicrosoftTeamsTenantBasedHandleRequestAttributes{}
	return &this
}

// GetChannelId returns the ChannelId field value.
func (o *MicrosoftTeamsTenantBasedHandleRequestAttributes) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsTenantBasedHandleRequestAttributes) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value.
func (o *MicrosoftTeamsTenantBasedHandleRequestAttributes) SetChannelId(v string) {
	o.ChannelId = v
}

// GetName returns the Name field value.
func (o *MicrosoftTeamsTenantBasedHandleRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsTenantBasedHandleRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MicrosoftTeamsTenantBasedHandleRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetTeamId returns the TeamId field value.
func (o *MicrosoftTeamsTenantBasedHandleRequestAttributes) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsTenantBasedHandleRequestAttributes) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value.
func (o *MicrosoftTeamsTenantBasedHandleRequestAttributes) SetTeamId(v string) {
	o.TeamId = v
}

// GetTenantId returns the TenantId field value.
func (o *MicrosoftTeamsTenantBasedHandleRequestAttributes) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsTenantBasedHandleRequestAttributes) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value.
func (o *MicrosoftTeamsTenantBasedHandleRequestAttributes) SetTenantId(v string) {
	o.TenantId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MicrosoftTeamsTenantBasedHandleRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["channel_id"] = o.ChannelId
	toSerialize["name"] = o.Name
	toSerialize["team_id"] = o.TeamId
	toSerialize["tenant_id"] = o.TenantId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MicrosoftTeamsTenantBasedHandleRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChannelId *string `json:"channel_id"`
		Name      *string `json:"name"`
		TeamId    *string `json:"team_id"`
		TenantId  *string `json:"tenant_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ChannelId == nil {
		return fmt.Errorf("required field channel_id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.TeamId == nil {
		return fmt.Errorf("required field team_id missing")
	}
	if all.TenantId == nil {
		return fmt.Errorf("required field tenant_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"channel_id", "name", "team_id", "tenant_id"})
	} else {
		return err
	}
	o.ChannelId = *all.ChannelId
	o.Name = *all.Name
	o.TeamId = *all.TeamId
	o.TenantId = *all.TenantId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
