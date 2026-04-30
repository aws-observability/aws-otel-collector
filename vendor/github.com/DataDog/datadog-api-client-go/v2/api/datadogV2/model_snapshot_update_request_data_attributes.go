// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnapshotUpdateRequestDataAttributes Attributes for updating a heatmap snapshot, including event, session, and view context.
type SnapshotUpdateRequestDataAttributes struct {
	// Unique identifier of the RUM event associated with the snapshot.
	EventId string `json:"event_id"`
	// Indicates whether the device type was explicitly selected by the user rather than auto-detected.
	IsDeviceTypeSelectedByUser bool `json:"is_device_type_selected_by_user"`
	// Unique identifier of the RUM session associated with the snapshot.
	SessionId *string `json:"session_id,omitempty"`
	// Offset in milliseconds from the start of the session at which the snapshot was captured.
	Start int64 `json:"start"`
	// Unique identifier of the RUM view associated with the snapshot.
	ViewId *string `json:"view_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSnapshotUpdateRequestDataAttributes instantiates a new SnapshotUpdateRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnapshotUpdateRequestDataAttributes(eventId string, isDeviceTypeSelectedByUser bool, start int64) *SnapshotUpdateRequestDataAttributes {
	this := SnapshotUpdateRequestDataAttributes{}
	this.EventId = eventId
	this.IsDeviceTypeSelectedByUser = isDeviceTypeSelectedByUser
	this.Start = start
	return &this
}

// NewSnapshotUpdateRequestDataAttributesWithDefaults instantiates a new SnapshotUpdateRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnapshotUpdateRequestDataAttributesWithDefaults() *SnapshotUpdateRequestDataAttributes {
	this := SnapshotUpdateRequestDataAttributes{}
	return &this
}

// GetEventId returns the EventId field value.
func (o *SnapshotUpdateRequestDataAttributes) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *SnapshotUpdateRequestDataAttributes) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value.
func (o *SnapshotUpdateRequestDataAttributes) SetEventId(v string) {
	o.EventId = v
}

// GetIsDeviceTypeSelectedByUser returns the IsDeviceTypeSelectedByUser field value.
func (o *SnapshotUpdateRequestDataAttributes) GetIsDeviceTypeSelectedByUser() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDeviceTypeSelectedByUser
}

// GetIsDeviceTypeSelectedByUserOk returns a tuple with the IsDeviceTypeSelectedByUser field value
// and a boolean to check if the value has been set.
func (o *SnapshotUpdateRequestDataAttributes) GetIsDeviceTypeSelectedByUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeviceTypeSelectedByUser, true
}

// SetIsDeviceTypeSelectedByUser sets field value.
func (o *SnapshotUpdateRequestDataAttributes) SetIsDeviceTypeSelectedByUser(v bool) {
	o.IsDeviceTypeSelectedByUser = v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *SnapshotUpdateRequestDataAttributes) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotUpdateRequestDataAttributes) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *SnapshotUpdateRequestDataAttributes) HasSessionId() bool {
	return o != nil && o.SessionId != nil
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *SnapshotUpdateRequestDataAttributes) SetSessionId(v string) {
	o.SessionId = &v
}

// GetStart returns the Start field value.
func (o *SnapshotUpdateRequestDataAttributes) GetStart() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *SnapshotUpdateRequestDataAttributes) GetStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *SnapshotUpdateRequestDataAttributes) SetStart(v int64) {
	o.Start = v
}

// GetViewId returns the ViewId field value if set, zero value otherwise.
func (o *SnapshotUpdateRequestDataAttributes) GetViewId() string {
	if o == nil || o.ViewId == nil {
		var ret string
		return ret
	}
	return *o.ViewId
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotUpdateRequestDataAttributes) GetViewIdOk() (*string, bool) {
	if o == nil || o.ViewId == nil {
		return nil, false
	}
	return o.ViewId, true
}

// HasViewId returns a boolean if a field has been set.
func (o *SnapshotUpdateRequestDataAttributes) HasViewId() bool {
	return o != nil && o.ViewId != nil
}

// SetViewId gets a reference to the given string and assigns it to the ViewId field.
func (o *SnapshotUpdateRequestDataAttributes) SetViewId(v string) {
	o.ViewId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnapshotUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["event_id"] = o.EventId
	toSerialize["is_device_type_selected_by_user"] = o.IsDeviceTypeSelectedByUser
	if o.SessionId != nil {
		toSerialize["session_id"] = o.SessionId
	}
	toSerialize["start"] = o.Start
	if o.ViewId != nil {
		toSerialize["view_id"] = o.ViewId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnapshotUpdateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventId                    *string `json:"event_id"`
		IsDeviceTypeSelectedByUser *bool   `json:"is_device_type_selected_by_user"`
		SessionId                  *string `json:"session_id,omitempty"`
		Start                      *int64  `json:"start"`
		ViewId                     *string `json:"view_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EventId == nil {
		return fmt.Errorf("required field event_id missing")
	}
	if all.IsDeviceTypeSelectedByUser == nil {
		return fmt.Errorf("required field is_device_type_selected_by_user missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event_id", "is_device_type_selected_by_user", "session_id", "start", "view_id"})
	} else {
		return err
	}
	o.EventId = *all.EventId
	o.IsDeviceTypeSelectedByUser = *all.IsDeviceTypeSelectedByUser
	o.SessionId = all.SessionId
	o.Start = *all.Start
	o.ViewId = all.ViewId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
