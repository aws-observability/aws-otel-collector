// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboard The metadata object associated with how a dashboard has been/will be shared.
type SharedDashboard struct {
	// User who shared the dashboard.
	Author *SharedDashboardAuthor `json:"author,omitempty"`
	// Date the dashboard was shared.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ID of the dashboard to share.
	DashboardId string `json:"dashboard_id"`
	// The type of the associated private dashboard.
	DashboardType DashboardType `json:"dashboard_type"`
	// Object containing the live span selection for the dashboard.
	GlobalTime *DashboardGlobalTime `json:"global_time,omitempty"`
	// Whether to allow viewers to select a different global time setting for the shared dashboard.
	GlobalTimeSelectableEnabled datadog.NullableBool `json:"global_time_selectable_enabled,omitempty"`
	// URL of the shared dashboard.
	PublicUrl *string `json:"public_url,omitempty"`
	// List of objects representing template variables on the shared dashboard which can have selectable values.
	SelectableTemplateVars []SelectableTemplateVariableItems `json:"selectable_template_vars,omitempty"`
	// List of email addresses that can receive an invitation to access to the shared dashboard.
	ShareList datadog.NullableList[string] `json:"share_list,omitempty"`
	// Type of sharing access (either open to anyone who has the public URL or invite-only).
	ShareType NullableDashboardShareType `json:"share_type,omitempty"`
	// A unique token assigned to the shared dashboard.
	Token *string `json:"token,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSharedDashboard instantiates a new SharedDashboard object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSharedDashboard(dashboardId string, dashboardType DashboardType) *SharedDashboard {
	this := SharedDashboard{}
	this.DashboardId = dashboardId
	this.DashboardType = dashboardType
	return &this
}

// NewSharedDashboardWithDefaults instantiates a new SharedDashboard object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSharedDashboardWithDefaults() *SharedDashboard {
	this := SharedDashboard{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *SharedDashboard) GetAuthor() SharedDashboardAuthor {
	if o == nil || o.Author == nil {
		var ret SharedDashboardAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboard) GetAuthorOk() (*SharedDashboardAuthor, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *SharedDashboard) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given SharedDashboardAuthor and assigns it to the Author field.
func (o *SharedDashboard) SetAuthor(v SharedDashboardAuthor) {
	o.Author = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SharedDashboard) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboard) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SharedDashboard) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SharedDashboard) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDashboardId returns the DashboardId field value.
func (o *SharedDashboard) GetDashboardId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DashboardId
}

// GetDashboardIdOk returns a tuple with the DashboardId field value
// and a boolean to check if the value has been set.
func (o *SharedDashboard) GetDashboardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DashboardId, true
}

// SetDashboardId sets field value.
func (o *SharedDashboard) SetDashboardId(v string) {
	o.DashboardId = v
}

// GetDashboardType returns the DashboardType field value.
func (o *SharedDashboard) GetDashboardType() DashboardType {
	if o == nil {
		var ret DashboardType
		return ret
	}
	return o.DashboardType
}

// GetDashboardTypeOk returns a tuple with the DashboardType field value
// and a boolean to check if the value has been set.
func (o *SharedDashboard) GetDashboardTypeOk() (*DashboardType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DashboardType, true
}

// SetDashboardType sets field value.
func (o *SharedDashboard) SetDashboardType(v DashboardType) {
	o.DashboardType = v
}

// GetGlobalTime returns the GlobalTime field value if set, zero value otherwise.
func (o *SharedDashboard) GetGlobalTime() DashboardGlobalTime {
	if o == nil || o.GlobalTime == nil {
		var ret DashboardGlobalTime
		return ret
	}
	return *o.GlobalTime
}

// GetGlobalTimeOk returns a tuple with the GlobalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboard) GetGlobalTimeOk() (*DashboardGlobalTime, bool) {
	if o == nil || o.GlobalTime == nil {
		return nil, false
	}
	return o.GlobalTime, true
}

// HasGlobalTime returns a boolean if a field has been set.
func (o *SharedDashboard) HasGlobalTime() bool {
	return o != nil && o.GlobalTime != nil
}

// SetGlobalTime gets a reference to the given DashboardGlobalTime and assigns it to the GlobalTime field.
func (o *SharedDashboard) SetGlobalTime(v DashboardGlobalTime) {
	o.GlobalTime = &v
}

// GetGlobalTimeSelectableEnabled returns the GlobalTimeSelectableEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedDashboard) GetGlobalTimeSelectableEnabled() bool {
	if o == nil || o.GlobalTimeSelectableEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.GlobalTimeSelectableEnabled.Get()
}

// GetGlobalTimeSelectableEnabledOk returns a tuple with the GlobalTimeSelectableEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboard) GetGlobalTimeSelectableEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.GlobalTimeSelectableEnabled.Get(), o.GlobalTimeSelectableEnabled.IsSet()
}

// HasGlobalTimeSelectableEnabled returns a boolean if a field has been set.
func (o *SharedDashboard) HasGlobalTimeSelectableEnabled() bool {
	return o != nil && o.GlobalTimeSelectableEnabled.IsSet()
}

// SetGlobalTimeSelectableEnabled gets a reference to the given datadog.NullableBool and assigns it to the GlobalTimeSelectableEnabled field.
func (o *SharedDashboard) SetGlobalTimeSelectableEnabled(v bool) {
	o.GlobalTimeSelectableEnabled.Set(&v)
}

// SetGlobalTimeSelectableEnabledNil sets the value for GlobalTimeSelectableEnabled to be an explicit nil.
func (o *SharedDashboard) SetGlobalTimeSelectableEnabledNil() {
	o.GlobalTimeSelectableEnabled.Set(nil)
}

// UnsetGlobalTimeSelectableEnabled ensures that no value is present for GlobalTimeSelectableEnabled, not even an explicit nil.
func (o *SharedDashboard) UnsetGlobalTimeSelectableEnabled() {
	o.GlobalTimeSelectableEnabled.Unset()
}

// GetPublicUrl returns the PublicUrl field value if set, zero value otherwise.
func (o *SharedDashboard) GetPublicUrl() string {
	if o == nil || o.PublicUrl == nil {
		var ret string
		return ret
	}
	return *o.PublicUrl
}

// GetPublicUrlOk returns a tuple with the PublicUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboard) GetPublicUrlOk() (*string, bool) {
	if o == nil || o.PublicUrl == nil {
		return nil, false
	}
	return o.PublicUrl, true
}

// HasPublicUrl returns a boolean if a field has been set.
func (o *SharedDashboard) HasPublicUrl() bool {
	return o != nil && o.PublicUrl != nil
}

// SetPublicUrl gets a reference to the given string and assigns it to the PublicUrl field.
func (o *SharedDashboard) SetPublicUrl(v string) {
	o.PublicUrl = &v
}

// GetSelectableTemplateVars returns the SelectableTemplateVars field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedDashboard) GetSelectableTemplateVars() []SelectableTemplateVariableItems {
	if o == nil {
		var ret []SelectableTemplateVariableItems
		return ret
	}
	return o.SelectableTemplateVars
}

// GetSelectableTemplateVarsOk returns a tuple with the SelectableTemplateVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboard) GetSelectableTemplateVarsOk() (*[]SelectableTemplateVariableItems, bool) {
	if o == nil || o.SelectableTemplateVars == nil {
		return nil, false
	}
	return &o.SelectableTemplateVars, true
}

// HasSelectableTemplateVars returns a boolean if a field has been set.
func (o *SharedDashboard) HasSelectableTemplateVars() bool {
	return o != nil && o.SelectableTemplateVars != nil
}

// SetSelectableTemplateVars gets a reference to the given []SelectableTemplateVariableItems and assigns it to the SelectableTemplateVars field.
func (o *SharedDashboard) SetSelectableTemplateVars(v []SelectableTemplateVariableItems) {
	o.SelectableTemplateVars = v
}

// GetShareList returns the ShareList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedDashboard) GetShareList() []string {
	if o == nil || o.ShareList.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ShareList.Get()
}

// GetShareListOk returns a tuple with the ShareList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboard) GetShareListOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareList.Get(), o.ShareList.IsSet()
}

// HasShareList returns a boolean if a field has been set.
func (o *SharedDashboard) HasShareList() bool {
	return o != nil && o.ShareList.IsSet()
}

// SetShareList gets a reference to the given datadog.NullableList[string] and assigns it to the ShareList field.
func (o *SharedDashboard) SetShareList(v []string) {
	o.ShareList.Set(&v)
}

// SetShareListNil sets the value for ShareList to be an explicit nil.
func (o *SharedDashboard) SetShareListNil() {
	o.ShareList.Set(nil)
}

// UnsetShareList ensures that no value is present for ShareList, not even an explicit nil.
func (o *SharedDashboard) UnsetShareList() {
	o.ShareList.Unset()
}

// GetShareType returns the ShareType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedDashboard) GetShareType() DashboardShareType {
	if o == nil || o.ShareType.Get() == nil {
		var ret DashboardShareType
		return ret
	}
	return *o.ShareType.Get()
}

// GetShareTypeOk returns a tuple with the ShareType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboard) GetShareTypeOk() (*DashboardShareType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareType.Get(), o.ShareType.IsSet()
}

// HasShareType returns a boolean if a field has been set.
func (o *SharedDashboard) HasShareType() bool {
	return o != nil && o.ShareType.IsSet()
}

// SetShareType gets a reference to the given NullableDashboardShareType and assigns it to the ShareType field.
func (o *SharedDashboard) SetShareType(v DashboardShareType) {
	o.ShareType.Set(&v)
}

// SetShareTypeNil sets the value for ShareType to be an explicit nil.
func (o *SharedDashboard) SetShareTypeNil() {
	o.ShareType.Set(nil)
}

// UnsetShareType ensures that no value is present for ShareType, not even an explicit nil.
func (o *SharedDashboard) UnsetShareType() {
	o.ShareType.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SharedDashboard) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboard) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SharedDashboard) HasToken() bool {
	return o != nil && o.Token != nil
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SharedDashboard) SetToken(v string) {
	o.Token = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SharedDashboard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["dashboard_id"] = o.DashboardId
	toSerialize["dashboard_type"] = o.DashboardType
	if o.GlobalTime != nil {
		toSerialize["global_time"] = o.GlobalTime
	}
	if o.GlobalTimeSelectableEnabled.IsSet() {
		toSerialize["global_time_selectable_enabled"] = o.GlobalTimeSelectableEnabled.Get()
	}
	if o.PublicUrl != nil {
		toSerialize["public_url"] = o.PublicUrl
	}
	if o.SelectableTemplateVars != nil {
		toSerialize["selectable_template_vars"] = o.SelectableTemplateVars
	}
	if o.ShareList.IsSet() {
		toSerialize["share_list"] = o.ShareList.Get()
	}
	if o.ShareType.IsSet() {
		toSerialize["share_type"] = o.ShareType.Get()
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SharedDashboard) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author                      *SharedDashboardAuthor            `json:"author,omitempty"`
		CreatedAt                   *time.Time                        `json:"created_at,omitempty"`
		DashboardId                 *string                           `json:"dashboard_id"`
		DashboardType               *DashboardType                    `json:"dashboard_type"`
		GlobalTime                  *DashboardGlobalTime              `json:"global_time,omitempty"`
		GlobalTimeSelectableEnabled datadog.NullableBool              `json:"global_time_selectable_enabled,omitempty"`
		PublicUrl                   *string                           `json:"public_url,omitempty"`
		SelectableTemplateVars      []SelectableTemplateVariableItems `json:"selectable_template_vars,omitempty"`
		ShareList                   datadog.NullableList[string]      `json:"share_list,omitempty"`
		ShareType                   NullableDashboardShareType        `json:"share_type,omitempty"`
		Token                       *string                           `json:"token,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DashboardId == nil {
		return fmt.Errorf("required field dashboard_id missing")
	}
	if all.DashboardType == nil {
		return fmt.Errorf("required field dashboard_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "created_at", "dashboard_id", "dashboard_type", "global_time", "global_time_selectable_enabled", "public_url", "selectable_template_vars", "share_list", "share_type", "token"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Author != nil && all.Author.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Author = all.Author
	o.CreatedAt = all.CreatedAt
	o.DashboardId = *all.DashboardId
	if !all.DashboardType.IsValid() {
		hasInvalidField = true
	} else {
		o.DashboardType = *all.DashboardType
	}
	if all.GlobalTime != nil && all.GlobalTime.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GlobalTime = all.GlobalTime
	o.GlobalTimeSelectableEnabled = all.GlobalTimeSelectableEnabled
	o.PublicUrl = all.PublicUrl
	o.SelectableTemplateVars = all.SelectableTemplateVars
	o.ShareList = all.ShareList
	if all.ShareType.Get() != nil && !all.ShareType.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.ShareType = all.ShareType
	}
	o.Token = all.Token

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
