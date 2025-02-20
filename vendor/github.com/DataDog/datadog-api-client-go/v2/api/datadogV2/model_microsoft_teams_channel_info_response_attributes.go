// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftTeamsChannelInfoResponseAttributes Channel attributes.
type MicrosoftTeamsChannelInfoResponseAttributes struct {
	// Indicates if this is the primary channel.
	IsPrimary *bool `json:"is_primary,omitempty"`
	// Team id.
	TeamId *string `json:"team_id,omitempty"`
	// Tenant id.
	TenantId *string `json:"tenant_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMicrosoftTeamsChannelInfoResponseAttributes instantiates a new MicrosoftTeamsChannelInfoResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMicrosoftTeamsChannelInfoResponseAttributes() *MicrosoftTeamsChannelInfoResponseAttributes {
	this := MicrosoftTeamsChannelInfoResponseAttributes{}
	return &this
}

// NewMicrosoftTeamsChannelInfoResponseAttributesWithDefaults instantiates a new MicrosoftTeamsChannelInfoResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMicrosoftTeamsChannelInfoResponseAttributesWithDefaults() *MicrosoftTeamsChannelInfoResponseAttributes {
	this := MicrosoftTeamsChannelInfoResponseAttributes{}
	return &this
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *MicrosoftTeamsChannelInfoResponseAttributes) GetIsPrimary() bool {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsChannelInfoResponseAttributes) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || o.IsPrimary == nil {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *MicrosoftTeamsChannelInfoResponseAttributes) HasIsPrimary() bool {
	return o != nil && o.IsPrimary != nil
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *MicrosoftTeamsChannelInfoResponseAttributes) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *MicrosoftTeamsChannelInfoResponseAttributes) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsChannelInfoResponseAttributes) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *MicrosoftTeamsChannelInfoResponseAttributes) HasTeamId() bool {
	return o != nil && o.TeamId != nil
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *MicrosoftTeamsChannelInfoResponseAttributes) SetTeamId(v string) {
	o.TeamId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *MicrosoftTeamsChannelInfoResponseAttributes) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsChannelInfoResponseAttributes) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *MicrosoftTeamsChannelInfoResponseAttributes) HasTenantId() bool {
	return o != nil && o.TenantId != nil
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *MicrosoftTeamsChannelInfoResponseAttributes) SetTenantId(v string) {
	o.TenantId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MicrosoftTeamsChannelInfoResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IsPrimary != nil {
		toSerialize["is_primary"] = o.IsPrimary
	}
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MicrosoftTeamsChannelInfoResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsPrimary *bool   `json:"is_primary,omitempty"`
		TeamId    *string `json:"team_id,omitempty"`
		TenantId  *string `json:"tenant_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_primary", "team_id", "tenant_id"})
	} else {
		return err
	}
	o.IsPrimary = all.IsPrimary
	o.TeamId = all.TeamId
	o.TenantId = all.TenantId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
