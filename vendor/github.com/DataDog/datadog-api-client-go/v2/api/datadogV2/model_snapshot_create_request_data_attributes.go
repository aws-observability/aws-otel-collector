// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnapshotCreateRequestDataAttributes Attributes for creating a heatmap snapshot, including the view, session, event, and device context.
type SnapshotCreateRequestDataAttributes struct {
	// Unique identifier of the RUM application.
	ApplicationId string `json:"application_id"`
	// Device type used when capturing the snapshot (e.g., desktop, mobile, tablet).
	DeviceType string `json:"device_type"`
	// Unique identifier of the RUM event associated with the snapshot.
	EventId string `json:"event_id"`
	// Indicates whether the device type was explicitly selected by the user rather than auto-detected.
	IsDeviceTypeSelectedByUser bool `json:"is_device_type_selected_by_user"`
	// Unique identifier of the RUM session associated with the snapshot.
	SessionId *string `json:"session_id,omitempty"`
	// Human-readable name for the snapshot.
	SnapshotName string `json:"snapshot_name"`
	// Offset in milliseconds from the start of the session at which the snapshot was captured.
	Start int64 `json:"start"`
	// Unique identifier of the RUM view associated with the snapshot.
	ViewId *string `json:"view_id,omitempty"`
	// URL path or name of the view where the snapshot was captured.
	ViewName string `json:"view_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSnapshotCreateRequestDataAttributes instantiates a new SnapshotCreateRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnapshotCreateRequestDataAttributes(applicationId string, deviceType string, eventId string, isDeviceTypeSelectedByUser bool, snapshotName string, start int64, viewName string) *SnapshotCreateRequestDataAttributes {
	this := SnapshotCreateRequestDataAttributes{}
	this.ApplicationId = applicationId
	this.DeviceType = deviceType
	this.EventId = eventId
	this.IsDeviceTypeSelectedByUser = isDeviceTypeSelectedByUser
	this.SnapshotName = snapshotName
	this.Start = start
	this.ViewName = viewName
	return &this
}

// NewSnapshotCreateRequestDataAttributesWithDefaults instantiates a new SnapshotCreateRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnapshotCreateRequestDataAttributesWithDefaults() *SnapshotCreateRequestDataAttributes {
	this := SnapshotCreateRequestDataAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value.
func (o *SnapshotCreateRequestDataAttributes) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *SnapshotCreateRequestDataAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *SnapshotCreateRequestDataAttributes) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetDeviceType returns the DeviceType field value.
func (o *SnapshotCreateRequestDataAttributes) GetDeviceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *SnapshotCreateRequestDataAttributes) GetDeviceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value.
func (o *SnapshotCreateRequestDataAttributes) SetDeviceType(v string) {
	o.DeviceType = v
}

// GetEventId returns the EventId field value.
func (o *SnapshotCreateRequestDataAttributes) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *SnapshotCreateRequestDataAttributes) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value.
func (o *SnapshotCreateRequestDataAttributes) SetEventId(v string) {
	o.EventId = v
}

// GetIsDeviceTypeSelectedByUser returns the IsDeviceTypeSelectedByUser field value.
func (o *SnapshotCreateRequestDataAttributes) GetIsDeviceTypeSelectedByUser() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDeviceTypeSelectedByUser
}

// GetIsDeviceTypeSelectedByUserOk returns a tuple with the IsDeviceTypeSelectedByUser field value
// and a boolean to check if the value has been set.
func (o *SnapshotCreateRequestDataAttributes) GetIsDeviceTypeSelectedByUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeviceTypeSelectedByUser, true
}

// SetIsDeviceTypeSelectedByUser sets field value.
func (o *SnapshotCreateRequestDataAttributes) SetIsDeviceTypeSelectedByUser(v bool) {
	o.IsDeviceTypeSelectedByUser = v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *SnapshotCreateRequestDataAttributes) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotCreateRequestDataAttributes) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *SnapshotCreateRequestDataAttributes) HasSessionId() bool {
	return o != nil && o.SessionId != nil
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *SnapshotCreateRequestDataAttributes) SetSessionId(v string) {
	o.SessionId = &v
}

// GetSnapshotName returns the SnapshotName field value.
func (o *SnapshotCreateRequestDataAttributes) GetSnapshotName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SnapshotName
}

// GetSnapshotNameOk returns a tuple with the SnapshotName field value
// and a boolean to check if the value has been set.
func (o *SnapshotCreateRequestDataAttributes) GetSnapshotNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotName, true
}

// SetSnapshotName sets field value.
func (o *SnapshotCreateRequestDataAttributes) SetSnapshotName(v string) {
	o.SnapshotName = v
}

// GetStart returns the Start field value.
func (o *SnapshotCreateRequestDataAttributes) GetStart() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *SnapshotCreateRequestDataAttributes) GetStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *SnapshotCreateRequestDataAttributes) SetStart(v int64) {
	o.Start = v
}

// GetViewId returns the ViewId field value if set, zero value otherwise.
func (o *SnapshotCreateRequestDataAttributes) GetViewId() string {
	if o == nil || o.ViewId == nil {
		var ret string
		return ret
	}
	return *o.ViewId
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotCreateRequestDataAttributes) GetViewIdOk() (*string, bool) {
	if o == nil || o.ViewId == nil {
		return nil, false
	}
	return o.ViewId, true
}

// HasViewId returns a boolean if a field has been set.
func (o *SnapshotCreateRequestDataAttributes) HasViewId() bool {
	return o != nil && o.ViewId != nil
}

// SetViewId gets a reference to the given string and assigns it to the ViewId field.
func (o *SnapshotCreateRequestDataAttributes) SetViewId(v string) {
	o.ViewId = &v
}

// GetViewName returns the ViewName field value.
func (o *SnapshotCreateRequestDataAttributes) GetViewName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
func (o *SnapshotCreateRequestDataAttributes) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewName, true
}

// SetViewName sets field value.
func (o *SnapshotCreateRequestDataAttributes) SetViewName(v string) {
	o.ViewName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnapshotCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["application_id"] = o.ApplicationId
	toSerialize["device_type"] = o.DeviceType
	toSerialize["event_id"] = o.EventId
	toSerialize["is_device_type_selected_by_user"] = o.IsDeviceTypeSelectedByUser
	if o.SessionId != nil {
		toSerialize["session_id"] = o.SessionId
	}
	toSerialize["snapshot_name"] = o.SnapshotName
	toSerialize["start"] = o.Start
	if o.ViewId != nil {
		toSerialize["view_id"] = o.ViewId
	}
	toSerialize["view_name"] = o.ViewName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnapshotCreateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId              *string `json:"application_id"`
		DeviceType                 *string `json:"device_type"`
		EventId                    *string `json:"event_id"`
		IsDeviceTypeSelectedByUser *bool   `json:"is_device_type_selected_by_user"`
		SessionId                  *string `json:"session_id,omitempty"`
		SnapshotName               *string `json:"snapshot_name"`
		Start                      *int64  `json:"start"`
		ViewId                     *string `json:"view_id,omitempty"`
		ViewName                   *string `json:"view_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApplicationId == nil {
		return fmt.Errorf("required field application_id missing")
	}
	if all.DeviceType == nil {
		return fmt.Errorf("required field device_type missing")
	}
	if all.EventId == nil {
		return fmt.Errorf("required field event_id missing")
	}
	if all.IsDeviceTypeSelectedByUser == nil {
		return fmt.Errorf("required field is_device_type_selected_by_user missing")
	}
	if all.SnapshotName == nil {
		return fmt.Errorf("required field snapshot_name missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	if all.ViewName == nil {
		return fmt.Errorf("required field view_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "device_type", "event_id", "is_device_type_selected_by_user", "session_id", "snapshot_name", "start", "view_id", "view_name"})
	} else {
		return err
	}
	o.ApplicationId = *all.ApplicationId
	o.DeviceType = *all.DeviceType
	o.EventId = *all.EventId
	o.IsDeviceTypeSelectedByUser = *all.IsDeviceTypeSelectedByUser
	o.SessionId = all.SessionId
	o.SnapshotName = *all.SnapshotName
	o.Start = *all.Start
	o.ViewId = all.ViewId
	o.ViewName = *all.ViewName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
