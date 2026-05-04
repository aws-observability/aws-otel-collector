// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WatcherDataAttributes Attributes of a user who has watched a RUM replay session, including contact information and watch statistics.
type WatcherDataAttributes struct {
	// Email handle of the user who watched the session.
	Handle string `json:"handle"`
	// URL or identifier of the watcher's avatar icon.
	Icon *string `json:"icon,omitempty"`
	// Timestamp when the watcher last viewed the session.
	LastWatchedAt time.Time `json:"last_watched_at"`
	// Display name of the user who watched the session.
	Name *string `json:"name,omitempty"`
	// Total number of times the user has watched the session.
	WatchCount int32 `json:"watch_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWatcherDataAttributes instantiates a new WatcherDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWatcherDataAttributes(handle string, lastWatchedAt time.Time, watchCount int32) *WatcherDataAttributes {
	this := WatcherDataAttributes{}
	this.Handle = handle
	this.LastWatchedAt = lastWatchedAt
	this.WatchCount = watchCount
	return &this
}

// NewWatcherDataAttributesWithDefaults instantiates a new WatcherDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWatcherDataAttributesWithDefaults() *WatcherDataAttributes {
	this := WatcherDataAttributes{}
	return &this
}

// GetHandle returns the Handle field value.
func (o *WatcherDataAttributes) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *WatcherDataAttributes) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *WatcherDataAttributes) SetHandle(v string) {
	o.Handle = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *WatcherDataAttributes) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatcherDataAttributes) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *WatcherDataAttributes) HasIcon() bool {
	return o != nil && o.Icon != nil
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *WatcherDataAttributes) SetIcon(v string) {
	o.Icon = &v
}

// GetLastWatchedAt returns the LastWatchedAt field value.
func (o *WatcherDataAttributes) GetLastWatchedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.LastWatchedAt
}

// GetLastWatchedAtOk returns a tuple with the LastWatchedAt field value
// and a boolean to check if the value has been set.
func (o *WatcherDataAttributes) GetLastWatchedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastWatchedAt, true
}

// SetLastWatchedAt sets field value.
func (o *WatcherDataAttributes) SetLastWatchedAt(v time.Time) {
	o.LastWatchedAt = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WatcherDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatcherDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WatcherDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WatcherDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetWatchCount returns the WatchCount field value.
func (o *WatcherDataAttributes) GetWatchCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.WatchCount
}

// GetWatchCountOk returns a tuple with the WatchCount field value
// and a boolean to check if the value has been set.
func (o *WatcherDataAttributes) GetWatchCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WatchCount, true
}

// SetWatchCount sets field value.
func (o *WatcherDataAttributes) SetWatchCount(v int32) {
	o.WatchCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WatcherDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["handle"] = o.Handle
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.LastWatchedAt.Nanosecond() == 0 {
		toSerialize["last_watched_at"] = o.LastWatchedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["last_watched_at"] = o.LastWatchedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["watch_count"] = o.WatchCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WatcherDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Handle        *string    `json:"handle"`
		Icon          *string    `json:"icon,omitempty"`
		LastWatchedAt *time.Time `json:"last_watched_at"`
		Name          *string    `json:"name,omitempty"`
		WatchCount    *int32     `json:"watch_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.LastWatchedAt == nil {
		return fmt.Errorf("required field last_watched_at missing")
	}
	if all.WatchCount == nil {
		return fmt.Errorf("required field watch_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"handle", "icon", "last_watched_at", "name", "watch_count"})
	} else {
		return err
	}
	o.Handle = *all.Handle
	o.Icon = all.Icon
	o.LastWatchedAt = *all.LastWatchedAt
	o.Name = all.Name
	o.WatchCount = *all.WatchCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
