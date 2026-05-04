// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnapshotDataAttributes Attributes of a heatmap snapshot, including view context, device information, and audit metadata.
type SnapshotDataAttributes struct {
	// Unique identifier of the RUM application.
	ApplicationId *string `json:"application_id,omitempty"`
	// Timestamp when the snapshot was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Display name of the user who created the snapshot.
	CreatedBy *string `json:"created_by,omitempty"`
	// Email handle of the user who created the snapshot.
	CreatedByHandle *string `json:"created_by_handle,omitempty"`
	// Numeric identifier of the user who created the snapshot.
	CreatedByUserId *int64 `json:"created_by_user_id,omitempty"`
	// Device type used when capturing the snapshot (e.g., desktop, mobile, tablet).
	DeviceType *string `json:"device_type,omitempty"`
	// Unique identifier of the RUM event associated with the snapshot.
	EventId *string `json:"event_id,omitempty"`
	// Indicates whether the device type was explicitly selected by the user rather than auto-detected.
	IsDeviceTypeSelectedByUser *bool `json:"is_device_type_selected_by_user,omitempty"`
	// Timestamp when the snapshot was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Numeric identifier of the organization that owns the snapshot.
	OrgId *int64 `json:"org_id,omitempty"`
	// Unique identifier of the RUM session associated with the snapshot.
	SessionId *string `json:"session_id,omitempty"`
	// Human-readable name for the snapshot.
	SnapshotName *string `json:"snapshot_name,omitempty"`
	// Offset in milliseconds from the start of the session at which the snapshot was captured.
	Start *int64 `json:"start,omitempty"`
	// Unique identifier of the RUM view associated with the snapshot.
	ViewId *string `json:"view_id,omitempty"`
	// URL path or name of the view where the snapshot was captured.
	ViewName *string `json:"view_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSnapshotDataAttributes instantiates a new SnapshotDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnapshotDataAttributes() *SnapshotDataAttributes {
	this := SnapshotDataAttributes{}
	return &this
}

// NewSnapshotDataAttributesWithDefaults instantiates a new SnapshotDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnapshotDataAttributesWithDefaults() *SnapshotDataAttributes {
	this := SnapshotDataAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasApplicationId() bool {
	return o != nil && o.ApplicationId != nil
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *SnapshotDataAttributes) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SnapshotDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SnapshotDataAttributes) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedByHandle returns the CreatedByHandle field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetCreatedByHandle() string {
	if o == nil || o.CreatedByHandle == nil {
		var ret string
		return ret
	}
	return *o.CreatedByHandle
}

// GetCreatedByHandleOk returns a tuple with the CreatedByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetCreatedByHandleOk() (*string, bool) {
	if o == nil || o.CreatedByHandle == nil {
		return nil, false
	}
	return o.CreatedByHandle, true
}

// HasCreatedByHandle returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasCreatedByHandle() bool {
	return o != nil && o.CreatedByHandle != nil
}

// SetCreatedByHandle gets a reference to the given string and assigns it to the CreatedByHandle field.
func (o *SnapshotDataAttributes) SetCreatedByHandle(v string) {
	o.CreatedByHandle = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetCreatedByUserId() int64 {
	if o == nil || o.CreatedByUserId == nil {
		var ret int64
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetCreatedByUserIdOk() (*int64, bool) {
	if o == nil || o.CreatedByUserId == nil {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasCreatedByUserId() bool {
	return o != nil && o.CreatedByUserId != nil
}

// SetCreatedByUserId gets a reference to the given int64 and assigns it to the CreatedByUserId field.
func (o *SnapshotDataAttributes) SetCreatedByUserId(v int64) {
	o.CreatedByUserId = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasDeviceType() bool {
	return o != nil && o.DeviceType != nil
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *SnapshotDataAttributes) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetEventId() string {
	if o == nil || o.EventId == nil {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetEventIdOk() (*string, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasEventId() bool {
	return o != nil && o.EventId != nil
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *SnapshotDataAttributes) SetEventId(v string) {
	o.EventId = &v
}

// GetIsDeviceTypeSelectedByUser returns the IsDeviceTypeSelectedByUser field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetIsDeviceTypeSelectedByUser() bool {
	if o == nil || o.IsDeviceTypeSelectedByUser == nil {
		var ret bool
		return ret
	}
	return *o.IsDeviceTypeSelectedByUser
}

// GetIsDeviceTypeSelectedByUserOk returns a tuple with the IsDeviceTypeSelectedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetIsDeviceTypeSelectedByUserOk() (*bool, bool) {
	if o == nil || o.IsDeviceTypeSelectedByUser == nil {
		return nil, false
	}
	return o.IsDeviceTypeSelectedByUser, true
}

// HasIsDeviceTypeSelectedByUser returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasIsDeviceTypeSelectedByUser() bool {
	return o != nil && o.IsDeviceTypeSelectedByUser != nil
}

// SetIsDeviceTypeSelectedByUser gets a reference to the given bool and assigns it to the IsDeviceTypeSelectedByUser field.
func (o *SnapshotDataAttributes) SetIsDeviceTypeSelectedByUser(v bool) {
	o.IsDeviceTypeSelectedByUser = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *SnapshotDataAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *SnapshotDataAttributes) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasSessionId() bool {
	return o != nil && o.SessionId != nil
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *SnapshotDataAttributes) SetSessionId(v string) {
	o.SessionId = &v
}

// GetSnapshotName returns the SnapshotName field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetSnapshotName() string {
	if o == nil || o.SnapshotName == nil {
		var ret string
		return ret
	}
	return *o.SnapshotName
}

// GetSnapshotNameOk returns a tuple with the SnapshotName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetSnapshotNameOk() (*string, bool) {
	if o == nil || o.SnapshotName == nil {
		return nil, false
	}
	return o.SnapshotName, true
}

// HasSnapshotName returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasSnapshotName() bool {
	return o != nil && o.SnapshotName != nil
}

// SetSnapshotName gets a reference to the given string and assigns it to the SnapshotName field.
func (o *SnapshotDataAttributes) SetSnapshotName(v string) {
	o.SnapshotName = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetStart() int64 {
	if o == nil || o.Start == nil {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetStartOk() (*int64, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *SnapshotDataAttributes) SetStart(v int64) {
	o.Start = &v
}

// GetViewId returns the ViewId field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetViewId() string {
	if o == nil || o.ViewId == nil {
		var ret string
		return ret
	}
	return *o.ViewId
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetViewIdOk() (*string, bool) {
	if o == nil || o.ViewId == nil {
		return nil, false
	}
	return o.ViewId, true
}

// HasViewId returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasViewId() bool {
	return o != nil && o.ViewId != nil
}

// SetViewId gets a reference to the given string and assigns it to the ViewId field.
func (o *SnapshotDataAttributes) SetViewId(v string) {
	o.ViewId = &v
}

// GetViewName returns the ViewName field value if set, zero value otherwise.
func (o *SnapshotDataAttributes) GetViewName() string {
	if o == nil || o.ViewName == nil {
		var ret string
		return ret
	}
	return *o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDataAttributes) GetViewNameOk() (*string, bool) {
	if o == nil || o.ViewName == nil {
		return nil, false
	}
	return o.ViewName, true
}

// HasViewName returns a boolean if a field has been set.
func (o *SnapshotDataAttributes) HasViewName() bool {
	return o != nil && o.ViewName != nil
}

// SetViewName gets a reference to the given string and assigns it to the ViewName field.
func (o *SnapshotDataAttributes) SetViewName(v string) {
	o.ViewName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnapshotDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplicationId != nil {
		toSerialize["application_id"] = o.ApplicationId
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.CreatedByHandle != nil {
		toSerialize["created_by_handle"] = o.CreatedByHandle
	}
	if o.CreatedByUserId != nil {
		toSerialize["created_by_user_id"] = o.CreatedByUserId
	}
	if o.DeviceType != nil {
		toSerialize["device_type"] = o.DeviceType
	}
	if o.EventId != nil {
		toSerialize["event_id"] = o.EventId
	}
	if o.IsDeviceTypeSelectedByUser != nil {
		toSerialize["is_device_type_selected_by_user"] = o.IsDeviceTypeSelectedByUser
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.SessionId != nil {
		toSerialize["session_id"] = o.SessionId
	}
	if o.SnapshotName != nil {
		toSerialize["snapshot_name"] = o.SnapshotName
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.ViewId != nil {
		toSerialize["view_id"] = o.ViewId
	}
	if o.ViewName != nil {
		toSerialize["view_name"] = o.ViewName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnapshotDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId              *string    `json:"application_id,omitempty"`
		CreatedAt                  *time.Time `json:"created_at,omitempty"`
		CreatedBy                  *string    `json:"created_by,omitempty"`
		CreatedByHandle            *string    `json:"created_by_handle,omitempty"`
		CreatedByUserId            *int64     `json:"created_by_user_id,omitempty"`
		DeviceType                 *string    `json:"device_type,omitempty"`
		EventId                    *string    `json:"event_id,omitempty"`
		IsDeviceTypeSelectedByUser *bool      `json:"is_device_type_selected_by_user,omitempty"`
		ModifiedAt                 *time.Time `json:"modified_at,omitempty"`
		OrgId                      *int64     `json:"org_id,omitempty"`
		SessionId                  *string    `json:"session_id,omitempty"`
		SnapshotName               *string    `json:"snapshot_name,omitempty"`
		Start                      *int64     `json:"start,omitempty"`
		ViewId                     *string    `json:"view_id,omitempty"`
		ViewName                   *string    `json:"view_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "created_at", "created_by", "created_by_handle", "created_by_user_id", "device_type", "event_id", "is_device_type_selected_by_user", "modified_at", "org_id", "session_id", "snapshot_name", "start", "view_id", "view_name"})
	} else {
		return err
	}
	o.ApplicationId = all.ApplicationId
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.CreatedByHandle = all.CreatedByHandle
	o.CreatedByUserId = all.CreatedByUserId
	o.DeviceType = all.DeviceType
	o.EventId = all.EventId
	o.IsDeviceTypeSelectedByUser = all.IsDeviceTypeSelectedByUser
	o.ModifiedAt = all.ModifiedAt
	o.OrgId = all.OrgId
	o.SessionId = all.SessionId
	o.SnapshotName = all.SnapshotName
	o.Start = all.Start
	o.ViewId = all.ViewId
	o.ViewName = all.ViewName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
