// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SlackIntegrationMetadata Incident integration metadata for the Slack integration.
type SlackIntegrationMetadata struct {
	// Array of Slack channels in this integration metadata.
	Channels []SlackIntegrationMetadataChannelItem `json:"channels"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSlackIntegrationMetadata instantiates a new SlackIntegrationMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSlackIntegrationMetadata(channels []SlackIntegrationMetadataChannelItem) *SlackIntegrationMetadata {
	this := SlackIntegrationMetadata{}
	this.Channels = channels
	return &this
}

// NewSlackIntegrationMetadataWithDefaults instantiates a new SlackIntegrationMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSlackIntegrationMetadataWithDefaults() *SlackIntegrationMetadata {
	this := SlackIntegrationMetadata{}
	return &this
}

// GetChannels returns the Channels field value.
func (o *SlackIntegrationMetadata) GetChannels() []SlackIntegrationMetadataChannelItem {
	if o == nil {
		var ret []SlackIntegrationMetadataChannelItem
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *SlackIntegrationMetadata) GetChannelsOk() (*[]SlackIntegrationMetadataChannelItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channels, true
}

// SetChannels sets field value.
func (o *SlackIntegrationMetadata) SetChannels(v []SlackIntegrationMetadataChannelItem) {
	o.Channels = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SlackIntegrationMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["channels"] = o.Channels

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SlackIntegrationMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Channels *[]SlackIntegrationMetadataChannelItem `json:"channels"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Channels == nil {
		return fmt.Errorf("required field channels missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"channels"})
	} else {
		return err
	}
	o.Channels = *all.Channels

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
