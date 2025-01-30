// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MSTeamsIntegrationMetadataTeamsItem Item in the Microsoft Teams integration metadata teams array.
type MSTeamsIntegrationMetadataTeamsItem struct {
	// Microsoft Teams channel ID.
	MsChannelId string `json:"ms_channel_id"`
	// Microsoft Teams channel name.
	MsChannelName string `json:"ms_channel_name"`
	// Microsoft Teams tenant ID.
	MsTenantId string `json:"ms_tenant_id"`
	// URL redirecting to the Microsoft Teams channel.
	RedirectUrl string `json:"redirect_url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMSTeamsIntegrationMetadataTeamsItem instantiates a new MSTeamsIntegrationMetadataTeamsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMSTeamsIntegrationMetadataTeamsItem(msChannelId string, msChannelName string, msTenantId string, redirectUrl string) *MSTeamsIntegrationMetadataTeamsItem {
	this := MSTeamsIntegrationMetadataTeamsItem{}
	this.MsChannelId = msChannelId
	this.MsChannelName = msChannelName
	this.MsTenantId = msTenantId
	this.RedirectUrl = redirectUrl
	return &this
}

// NewMSTeamsIntegrationMetadataTeamsItemWithDefaults instantiates a new MSTeamsIntegrationMetadataTeamsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMSTeamsIntegrationMetadataTeamsItemWithDefaults() *MSTeamsIntegrationMetadataTeamsItem {
	this := MSTeamsIntegrationMetadataTeamsItem{}
	return &this
}

// GetMsChannelId returns the MsChannelId field value.
func (o *MSTeamsIntegrationMetadataTeamsItem) GetMsChannelId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MsChannelId
}

// GetMsChannelIdOk returns a tuple with the MsChannelId field value
// and a boolean to check if the value has been set.
func (o *MSTeamsIntegrationMetadataTeamsItem) GetMsChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsChannelId, true
}

// SetMsChannelId sets field value.
func (o *MSTeamsIntegrationMetadataTeamsItem) SetMsChannelId(v string) {
	o.MsChannelId = v
}

// GetMsChannelName returns the MsChannelName field value.
func (o *MSTeamsIntegrationMetadataTeamsItem) GetMsChannelName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MsChannelName
}

// GetMsChannelNameOk returns a tuple with the MsChannelName field value
// and a boolean to check if the value has been set.
func (o *MSTeamsIntegrationMetadataTeamsItem) GetMsChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsChannelName, true
}

// SetMsChannelName sets field value.
func (o *MSTeamsIntegrationMetadataTeamsItem) SetMsChannelName(v string) {
	o.MsChannelName = v
}

// GetMsTenantId returns the MsTenantId field value.
func (o *MSTeamsIntegrationMetadataTeamsItem) GetMsTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MsTenantId
}

// GetMsTenantIdOk returns a tuple with the MsTenantId field value
// and a boolean to check if the value has been set.
func (o *MSTeamsIntegrationMetadataTeamsItem) GetMsTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsTenantId, true
}

// SetMsTenantId sets field value.
func (o *MSTeamsIntegrationMetadataTeamsItem) SetMsTenantId(v string) {
	o.MsTenantId = v
}

// GetRedirectUrl returns the RedirectUrl field value.
func (o *MSTeamsIntegrationMetadataTeamsItem) GetRedirectUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value
// and a boolean to check if the value has been set.
func (o *MSTeamsIntegrationMetadataTeamsItem) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectUrl, true
}

// SetRedirectUrl sets field value.
func (o *MSTeamsIntegrationMetadataTeamsItem) SetRedirectUrl(v string) {
	o.RedirectUrl = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MSTeamsIntegrationMetadataTeamsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ms_channel_id"] = o.MsChannelId
	toSerialize["ms_channel_name"] = o.MsChannelName
	toSerialize["ms_tenant_id"] = o.MsTenantId
	toSerialize["redirect_url"] = o.RedirectUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MSTeamsIntegrationMetadataTeamsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MsChannelId   *string `json:"ms_channel_id"`
		MsChannelName *string `json:"ms_channel_name"`
		MsTenantId    *string `json:"ms_tenant_id"`
		RedirectUrl   *string `json:"redirect_url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MsChannelId == nil {
		return fmt.Errorf("required field ms_channel_id missing")
	}
	if all.MsChannelName == nil {
		return fmt.Errorf("required field ms_channel_name missing")
	}
	if all.MsTenantId == nil {
		return fmt.Errorf("required field ms_tenant_id missing")
	}
	if all.RedirectUrl == nil {
		return fmt.Errorf("required field redirect_url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ms_channel_id", "ms_channel_name", "ms_tenant_id", "redirect_url"})
	} else {
		return err
	}
	o.MsChannelId = *all.MsChannelId
	o.MsChannelName = *all.MsChannelName
	o.MsTenantId = *all.MsTenantId
	o.RedirectUrl = *all.RedirectUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
