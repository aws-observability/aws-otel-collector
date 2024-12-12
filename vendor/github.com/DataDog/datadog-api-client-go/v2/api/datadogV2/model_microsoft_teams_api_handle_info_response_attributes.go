// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftTeamsApiHandleInfoResponseAttributes Handle attributes.
type MicrosoftTeamsApiHandleInfoResponseAttributes struct {
	// Channel id.
	ChannelId *string `json:"channel_id,omitempty"`
	// Channel name.
	ChannelName *string `json:"channel_name,omitempty"`
	// Handle name.
	Name *string `json:"name,omitempty"`
	// Team id.
	TeamId *string `json:"team_id,omitempty"`
	// Team name.
	TeamName *string `json:"team_name,omitempty"`
	// Tenant id.
	TenantId *string `json:"tenant_id,omitempty"`
	// Tenant name.
	TenantName *string `json:"tenant_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMicrosoftTeamsApiHandleInfoResponseAttributes instantiates a new MicrosoftTeamsApiHandleInfoResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMicrosoftTeamsApiHandleInfoResponseAttributes() *MicrosoftTeamsApiHandleInfoResponseAttributes {
	this := MicrosoftTeamsApiHandleInfoResponseAttributes{}
	return &this
}

// NewMicrosoftTeamsApiHandleInfoResponseAttributesWithDefaults instantiates a new MicrosoftTeamsApiHandleInfoResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMicrosoftTeamsApiHandleInfoResponseAttributesWithDefaults() *MicrosoftTeamsApiHandleInfoResponseAttributes {
	this := MicrosoftTeamsApiHandleInfoResponseAttributes{}
	return &this
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) GetChannelId() string {
	if o == nil || o.ChannelId == nil {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) GetChannelIdOk() (*string, bool) {
	if o == nil || o.ChannelId == nil {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) HasChannelId() bool {
	return o != nil && o.ChannelId != nil
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) GetChannelName() string {
	if o == nil || o.ChannelName == nil {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) GetChannelNameOk() (*string, bool) {
	if o == nil || o.ChannelName == nil {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) HasChannelName() bool {
	return o != nil && o.ChannelName != nil
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) SetName(v string) {
	o.Name = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) HasTeamId() bool {
	return o != nil && o.TeamId != nil
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) SetTeamId(v string) {
	o.TeamId = &v
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) GetTeamName() string {
	if o == nil || o.TeamName == nil {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) GetTeamNameOk() (*string, bool) {
	if o == nil || o.TeamName == nil {
		return nil, false
	}
	return o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) HasTeamName() bool {
	return o != nil && o.TeamName != nil
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) SetTeamName(v string) {
	o.TeamName = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) HasTenantId() bool {
	return o != nil && o.TenantId != nil
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) SetTenantId(v string) {
	o.TenantId = &v
}

// GetTenantName returns the TenantName field value if set, zero value otherwise.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) GetTenantName() string {
	if o == nil || o.TenantName == nil {
		var ret string
		return ret
	}
	return *o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) GetTenantNameOk() (*string, bool) {
	if o == nil || o.TenantName == nil {
		return nil, false
	}
	return o.TenantName, true
}

// HasTenantName returns a boolean if a field has been set.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) HasTenantName() bool {
	return o != nil && o.TenantName != nil
}

// SetTenantName gets a reference to the given string and assigns it to the TenantName field.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) SetTenantName(v string) {
	o.TenantName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MicrosoftTeamsApiHandleInfoResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChannelId != nil {
		toSerialize["channel_id"] = o.ChannelId
	}
	if o.ChannelName != nil {
		toSerialize["channel_name"] = o.ChannelName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}
	if o.TeamName != nil {
		toSerialize["team_name"] = o.TeamName
	}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
	}
	if o.TenantName != nil {
		toSerialize["tenant_name"] = o.TenantName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MicrosoftTeamsApiHandleInfoResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChannelId   *string `json:"channel_id,omitempty"`
		ChannelName *string `json:"channel_name,omitempty"`
		Name        *string `json:"name,omitempty"`
		TeamId      *string `json:"team_id,omitempty"`
		TeamName    *string `json:"team_name,omitempty"`
		TenantId    *string `json:"tenant_id,omitempty"`
		TenantName  *string `json:"tenant_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"channel_id", "channel_name", "name", "team_id", "team_name", "tenant_id", "tenant_name"})
	} else {
		return err
	}
	o.ChannelId = all.ChannelId
	o.ChannelName = all.ChannelName
	o.Name = all.Name
	o.TeamId = all.TeamId
	o.TeamName = all.TeamName
	o.TenantId = all.TenantId
	o.TenantName = all.TenantName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
