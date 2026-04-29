// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ViewershipHistorySessionDataAttributes Attributes of a viewership history session entry, capturing when it was last watched and the associated event data.
type ViewershipHistorySessionDataAttributes struct {
	// Unique identifier of the RUM event associated with the watched session.
	EventId *string `json:"event_id,omitempty"`
	// Timestamp when the session was last watched by the user.
	LastWatchedAt time.Time `json:"last_watched_at"`
	// Raw event data associated with the replay session.
	SessionEvent map[string]interface{} `json:"session_event,omitempty"`
	// Replay track identifier indicating which recording track the session belongs to.
	Track *string `json:"track,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewViewershipHistorySessionDataAttributes instantiates a new ViewershipHistorySessionDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewViewershipHistorySessionDataAttributes(lastWatchedAt time.Time) *ViewershipHistorySessionDataAttributes {
	this := ViewershipHistorySessionDataAttributes{}
	this.LastWatchedAt = lastWatchedAt
	return &this
}

// NewViewershipHistorySessionDataAttributesWithDefaults instantiates a new ViewershipHistorySessionDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewViewershipHistorySessionDataAttributesWithDefaults() *ViewershipHistorySessionDataAttributes {
	this := ViewershipHistorySessionDataAttributes{}
	return &this
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *ViewershipHistorySessionDataAttributes) GetEventId() string {
	if o == nil || o.EventId == nil {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewershipHistorySessionDataAttributes) GetEventIdOk() (*string, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *ViewershipHistorySessionDataAttributes) HasEventId() bool {
	return o != nil && o.EventId != nil
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *ViewershipHistorySessionDataAttributes) SetEventId(v string) {
	o.EventId = &v
}

// GetLastWatchedAt returns the LastWatchedAt field value.
func (o *ViewershipHistorySessionDataAttributes) GetLastWatchedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.LastWatchedAt
}

// GetLastWatchedAtOk returns a tuple with the LastWatchedAt field value
// and a boolean to check if the value has been set.
func (o *ViewershipHistorySessionDataAttributes) GetLastWatchedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastWatchedAt, true
}

// SetLastWatchedAt sets field value.
func (o *ViewershipHistorySessionDataAttributes) SetLastWatchedAt(v time.Time) {
	o.LastWatchedAt = v
}

// GetSessionEvent returns the SessionEvent field value if set, zero value otherwise.
func (o *ViewershipHistorySessionDataAttributes) GetSessionEvent() map[string]interface{} {
	if o == nil || o.SessionEvent == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SessionEvent
}

// GetSessionEventOk returns a tuple with the SessionEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewershipHistorySessionDataAttributes) GetSessionEventOk() (*map[string]interface{}, bool) {
	if o == nil || o.SessionEvent == nil {
		return nil, false
	}
	return &o.SessionEvent, true
}

// HasSessionEvent returns a boolean if a field has been set.
func (o *ViewershipHistorySessionDataAttributes) HasSessionEvent() bool {
	return o != nil && o.SessionEvent != nil
}

// SetSessionEvent gets a reference to the given map[string]interface{} and assigns it to the SessionEvent field.
func (o *ViewershipHistorySessionDataAttributes) SetSessionEvent(v map[string]interface{}) {
	o.SessionEvent = v
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *ViewershipHistorySessionDataAttributes) GetTrack() string {
	if o == nil || o.Track == nil {
		var ret string
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewershipHistorySessionDataAttributes) GetTrackOk() (*string, bool) {
	if o == nil || o.Track == nil {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *ViewershipHistorySessionDataAttributes) HasTrack() bool {
	return o != nil && o.Track != nil
}

// SetTrack gets a reference to the given string and assigns it to the Track field.
func (o *ViewershipHistorySessionDataAttributes) SetTrack(v string) {
	o.Track = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ViewershipHistorySessionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EventId != nil {
		toSerialize["event_id"] = o.EventId
	}
	if o.LastWatchedAt.Nanosecond() == 0 {
		toSerialize["last_watched_at"] = o.LastWatchedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["last_watched_at"] = o.LastWatchedAt.Format("2006-01-02T15:04:05.000Z07:00")
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
func (o *ViewershipHistorySessionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventId       *string                `json:"event_id,omitempty"`
		LastWatchedAt *time.Time             `json:"last_watched_at"`
		SessionEvent  map[string]interface{} `json:"session_event,omitempty"`
		Track         *string                `json:"track,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LastWatchedAt == nil {
		return fmt.Errorf("required field last_watched_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event_id", "last_watched_at", "session_event", "track"})
	} else {
		return err
	}
	o.EventId = all.EventId
	o.LastWatchedAt = *all.LastWatchedAt
	o.SessionEvent = all.SessionEvent
	o.Track = all.Track

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
