// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PlaylistsSessionDataAttributes Attributes of a session within a playlist, including the session event data and its replay track.
type PlaylistsSessionDataAttributes struct {
	// Raw event data associated with the replay session.
	SessionEvent map[string]interface{} `json:"session_event,omitempty"`
	// Replay track identifier indicating which recording track the session belongs to.
	Track *string `json:"track,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPlaylistsSessionDataAttributes instantiates a new PlaylistsSessionDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPlaylistsSessionDataAttributes() *PlaylistsSessionDataAttributes {
	this := PlaylistsSessionDataAttributes{}
	return &this
}

// NewPlaylistsSessionDataAttributesWithDefaults instantiates a new PlaylistsSessionDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPlaylistsSessionDataAttributesWithDefaults() *PlaylistsSessionDataAttributes {
	this := PlaylistsSessionDataAttributes{}
	return &this
}

// GetSessionEvent returns the SessionEvent field value if set, zero value otherwise.
func (o *PlaylistsSessionDataAttributes) GetSessionEvent() map[string]interface{} {
	if o == nil || o.SessionEvent == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SessionEvent
}

// GetSessionEventOk returns a tuple with the SessionEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistsSessionDataAttributes) GetSessionEventOk() (*map[string]interface{}, bool) {
	if o == nil || o.SessionEvent == nil {
		return nil, false
	}
	return &o.SessionEvent, true
}

// HasSessionEvent returns a boolean if a field has been set.
func (o *PlaylistsSessionDataAttributes) HasSessionEvent() bool {
	return o != nil && o.SessionEvent != nil
}

// SetSessionEvent gets a reference to the given map[string]interface{} and assigns it to the SessionEvent field.
func (o *PlaylistsSessionDataAttributes) SetSessionEvent(v map[string]interface{}) {
	o.SessionEvent = v
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *PlaylistsSessionDataAttributes) GetTrack() string {
	if o == nil || o.Track == nil {
		var ret string
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistsSessionDataAttributes) GetTrackOk() (*string, bool) {
	if o == nil || o.Track == nil {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *PlaylistsSessionDataAttributes) HasTrack() bool {
	return o != nil && o.Track != nil
}

// SetTrack gets a reference to the given string and assigns it to the Track field.
func (o *PlaylistsSessionDataAttributes) SetTrack(v string) {
	o.Track = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PlaylistsSessionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.SessionEvent != nil {
		toSerialize["session_event"] = o.SessionEvent
	}
	if o.Track != nil {
		toSerialize["track"] = o.Track
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PlaylistsSessionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SessionEvent map[string]interface{} `json:"session_event,omitempty"`
		Track        *string                `json:"track,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"session_event", "track"})
	} else {
		return err
	}
	o.SessionEvent = all.SessionEvent
	o.Track = all.Track

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
