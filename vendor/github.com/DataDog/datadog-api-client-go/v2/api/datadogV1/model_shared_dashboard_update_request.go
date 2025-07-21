// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardUpdateRequest Update a shared dashboard's settings.
type SharedDashboardUpdateRequest struct {
	// The `SharedDashboard` `embeddable_domains`.
	EmbeddableDomains []string `json:"embeddable_domains,omitempty"`
	// The time when an OPEN shared dashboard becomes publicly unavailable.
	Expiration datadog.NullableTime `json:"expiration,omitempty"`
	// Timeframe setting for the shared dashboard.
	GlobalTime NullableSharedDashboardUpdateRequestGlobalTime `json:"global_time,omitempty"`
	// Whether to allow viewers to select a different global time setting for the shared dashboard.
	GlobalTimeSelectableEnabled datadog.NullableBool `json:"global_time_selectable_enabled,omitempty"`
	// The `SharedDashboard` `invitees`.
	Invitees []SharedDashboardInviteesItems `json:"invitees,omitempty"`
	// List of objects representing template variables on the shared dashboard which can have selectable values.
	SelectableTemplateVars []SelectableTemplateVariableItems `json:"selectable_template_vars,omitempty"`
	// List of email addresses that can be given access to the shared dashboard.
	// Deprecated
	ShareList datadog.NullableList[string] `json:"share_list,omitempty"`
	// Type of sharing access (either open to anyone who has the public URL or invite-only).
	ShareType NullableDashboardShareType `json:"share_type,omitempty"`
	// Active means the dashboard is publicly available. Paused means the dashboard is not publicly available.
	Status *SharedDashboardStatus `json:"status,omitempty"`
	// Title of the shared dashboard.
	Title *string `json:"title,omitempty"`
	// The viewing preferences for a shared dashboard.
	ViewingPreferences *ViewingPreferences `json:"viewing_preferences,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSharedDashboardUpdateRequest instantiates a new SharedDashboardUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSharedDashboardUpdateRequest() *SharedDashboardUpdateRequest {
	this := SharedDashboardUpdateRequest{}
	return &this
}

// NewSharedDashboardUpdateRequestWithDefaults instantiates a new SharedDashboardUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSharedDashboardUpdateRequestWithDefaults() *SharedDashboardUpdateRequest {
	this := SharedDashboardUpdateRequest{}
	return &this
}

// GetEmbeddableDomains returns the EmbeddableDomains field value if set, zero value otherwise.
func (o *SharedDashboardUpdateRequest) GetEmbeddableDomains() []string {
	if o == nil || o.EmbeddableDomains == nil {
		var ret []string
		return ret
	}
	return o.EmbeddableDomains
}

// GetEmbeddableDomainsOk returns a tuple with the EmbeddableDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboardUpdateRequest) GetEmbeddableDomainsOk() (*[]string, bool) {
	if o == nil || o.EmbeddableDomains == nil {
		return nil, false
	}
	return &o.EmbeddableDomains, true
}

// HasEmbeddableDomains returns a boolean if a field has been set.
func (o *SharedDashboardUpdateRequest) HasEmbeddableDomains() bool {
	return o != nil && o.EmbeddableDomains != nil
}

// SetEmbeddableDomains gets a reference to the given []string and assigns it to the EmbeddableDomains field.
func (o *SharedDashboardUpdateRequest) SetEmbeddableDomains(v []string) {
	o.EmbeddableDomains = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedDashboardUpdateRequest) GetExpiration() time.Time {
	if o == nil || o.Expiration.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboardUpdateRequest) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// HasExpiration returns a boolean if a field has been set.
func (o *SharedDashboardUpdateRequest) HasExpiration() bool {
	return o != nil && o.Expiration.IsSet()
}

// SetExpiration gets a reference to the given datadog.NullableTime and assigns it to the Expiration field.
func (o *SharedDashboardUpdateRequest) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}

// SetExpirationNil sets the value for Expiration to be an explicit nil.
func (o *SharedDashboardUpdateRequest) SetExpirationNil() {
	o.Expiration.Set(nil)
}

// UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil.
func (o *SharedDashboardUpdateRequest) UnsetExpiration() {
	o.Expiration.Unset()
}

// GetGlobalTime returns the GlobalTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedDashboardUpdateRequest) GetGlobalTime() SharedDashboardUpdateRequestGlobalTime {
	if o == nil || o.GlobalTime.Get() == nil {
		var ret SharedDashboardUpdateRequestGlobalTime
		return ret
	}
	return *o.GlobalTime.Get()
}

// GetGlobalTimeOk returns a tuple with the GlobalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboardUpdateRequest) GetGlobalTimeOk() (*SharedDashboardUpdateRequestGlobalTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.GlobalTime.Get(), o.GlobalTime.IsSet()
}

// HasGlobalTime returns a boolean if a field has been set.
func (o *SharedDashboardUpdateRequest) HasGlobalTime() bool {
	return o != nil && o.GlobalTime.IsSet()
}

// SetGlobalTime gets a reference to the given NullableSharedDashboardUpdateRequestGlobalTime and assigns it to the GlobalTime field.
func (o *SharedDashboardUpdateRequest) SetGlobalTime(v SharedDashboardUpdateRequestGlobalTime) {
	o.GlobalTime.Set(&v)
}

// SetGlobalTimeNil sets the value for GlobalTime to be an explicit nil.
func (o *SharedDashboardUpdateRequest) SetGlobalTimeNil() {
	o.GlobalTime.Set(nil)
}

// UnsetGlobalTime ensures that no value is present for GlobalTime, not even an explicit nil.
func (o *SharedDashboardUpdateRequest) UnsetGlobalTime() {
	o.GlobalTime.Unset()
}

// GetGlobalTimeSelectableEnabled returns the GlobalTimeSelectableEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedDashboardUpdateRequest) GetGlobalTimeSelectableEnabled() bool {
	if o == nil || o.GlobalTimeSelectableEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.GlobalTimeSelectableEnabled.Get()
}

// GetGlobalTimeSelectableEnabledOk returns a tuple with the GlobalTimeSelectableEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboardUpdateRequest) GetGlobalTimeSelectableEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.GlobalTimeSelectableEnabled.Get(), o.GlobalTimeSelectableEnabled.IsSet()
}

// HasGlobalTimeSelectableEnabled returns a boolean if a field has been set.
func (o *SharedDashboardUpdateRequest) HasGlobalTimeSelectableEnabled() bool {
	return o != nil && o.GlobalTimeSelectableEnabled.IsSet()
}

// SetGlobalTimeSelectableEnabled gets a reference to the given datadog.NullableBool and assigns it to the GlobalTimeSelectableEnabled field.
func (o *SharedDashboardUpdateRequest) SetGlobalTimeSelectableEnabled(v bool) {
	o.GlobalTimeSelectableEnabled.Set(&v)
}

// SetGlobalTimeSelectableEnabledNil sets the value for GlobalTimeSelectableEnabled to be an explicit nil.
func (o *SharedDashboardUpdateRequest) SetGlobalTimeSelectableEnabledNil() {
	o.GlobalTimeSelectableEnabled.Set(nil)
}

// UnsetGlobalTimeSelectableEnabled ensures that no value is present for GlobalTimeSelectableEnabled, not even an explicit nil.
func (o *SharedDashboardUpdateRequest) UnsetGlobalTimeSelectableEnabled() {
	o.GlobalTimeSelectableEnabled.Unset()
}

// GetInvitees returns the Invitees field value if set, zero value otherwise.
func (o *SharedDashboardUpdateRequest) GetInvitees() []SharedDashboardInviteesItems {
	if o == nil || o.Invitees == nil {
		var ret []SharedDashboardInviteesItems
		return ret
	}
	return o.Invitees
}

// GetInviteesOk returns a tuple with the Invitees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboardUpdateRequest) GetInviteesOk() (*[]SharedDashboardInviteesItems, bool) {
	if o == nil || o.Invitees == nil {
		return nil, false
	}
	return &o.Invitees, true
}

// HasInvitees returns a boolean if a field has been set.
func (o *SharedDashboardUpdateRequest) HasInvitees() bool {
	return o != nil && o.Invitees != nil
}

// SetInvitees gets a reference to the given []SharedDashboardInviteesItems and assigns it to the Invitees field.
func (o *SharedDashboardUpdateRequest) SetInvitees(v []SharedDashboardInviteesItems) {
	o.Invitees = v
}

// GetSelectableTemplateVars returns the SelectableTemplateVars field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedDashboardUpdateRequest) GetSelectableTemplateVars() []SelectableTemplateVariableItems {
	if o == nil {
		var ret []SelectableTemplateVariableItems
		return ret
	}
	return o.SelectableTemplateVars
}

// GetSelectableTemplateVarsOk returns a tuple with the SelectableTemplateVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboardUpdateRequest) GetSelectableTemplateVarsOk() (*[]SelectableTemplateVariableItems, bool) {
	if o == nil || o.SelectableTemplateVars == nil {
		return nil, false
	}
	return &o.SelectableTemplateVars, true
}

// HasSelectableTemplateVars returns a boolean if a field has been set.
func (o *SharedDashboardUpdateRequest) HasSelectableTemplateVars() bool {
	return o != nil && o.SelectableTemplateVars != nil
}

// SetSelectableTemplateVars gets a reference to the given []SelectableTemplateVariableItems and assigns it to the SelectableTemplateVars field.
func (o *SharedDashboardUpdateRequest) SetSelectableTemplateVars(v []SelectableTemplateVariableItems) {
	o.SelectableTemplateVars = v
}

// GetShareList returns the ShareList field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *SharedDashboardUpdateRequest) GetShareList() []string {
	if o == nil || o.ShareList.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ShareList.Get()
}

// GetShareListOk returns a tuple with the ShareList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
// Deprecated
func (o *SharedDashboardUpdateRequest) GetShareListOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareList.Get(), o.ShareList.IsSet()
}

// HasShareList returns a boolean if a field has been set.
func (o *SharedDashboardUpdateRequest) HasShareList() bool {
	return o != nil && o.ShareList.IsSet()
}

// SetShareList gets a reference to the given datadog.NullableList[string] and assigns it to the ShareList field.
// Deprecated
func (o *SharedDashboardUpdateRequest) SetShareList(v []string) {
	o.ShareList.Set(&v)
}

// SetShareListNil sets the value for ShareList to be an explicit nil.
func (o *SharedDashboardUpdateRequest) SetShareListNil() {
	o.ShareList.Set(nil)
}

// UnsetShareList ensures that no value is present for ShareList, not even an explicit nil.
func (o *SharedDashboardUpdateRequest) UnsetShareList() {
	o.ShareList.Unset()
}

// GetShareType returns the ShareType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedDashboardUpdateRequest) GetShareType() DashboardShareType {
	if o == nil || o.ShareType.Get() == nil {
		var ret DashboardShareType
		return ret
	}
	return *o.ShareType.Get()
}

// GetShareTypeOk returns a tuple with the ShareType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboardUpdateRequest) GetShareTypeOk() (*DashboardShareType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareType.Get(), o.ShareType.IsSet()
}

// HasShareType returns a boolean if a field has been set.
func (o *SharedDashboardUpdateRequest) HasShareType() bool {
	return o != nil && o.ShareType.IsSet()
}

// SetShareType gets a reference to the given NullableDashboardShareType and assigns it to the ShareType field.
func (o *SharedDashboardUpdateRequest) SetShareType(v DashboardShareType) {
	o.ShareType.Set(&v)
}

// SetShareTypeNil sets the value for ShareType to be an explicit nil.
func (o *SharedDashboardUpdateRequest) SetShareTypeNil() {
	o.ShareType.Set(nil)
}

// UnsetShareType ensures that no value is present for ShareType, not even an explicit nil.
func (o *SharedDashboardUpdateRequest) UnsetShareType() {
	o.ShareType.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SharedDashboardUpdateRequest) GetStatus() SharedDashboardStatus {
	if o == nil || o.Status == nil {
		var ret SharedDashboardStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboardUpdateRequest) GetStatusOk() (*SharedDashboardStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SharedDashboardUpdateRequest) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given SharedDashboardStatus and assigns it to the Status field.
func (o *SharedDashboardUpdateRequest) SetStatus(v SharedDashboardStatus) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SharedDashboardUpdateRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboardUpdateRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SharedDashboardUpdateRequest) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SharedDashboardUpdateRequest) SetTitle(v string) {
	o.Title = &v
}

// GetViewingPreferences returns the ViewingPreferences field value if set, zero value otherwise.
func (o *SharedDashboardUpdateRequest) GetViewingPreferences() ViewingPreferences {
	if o == nil || o.ViewingPreferences == nil {
		var ret ViewingPreferences
		return ret
	}
	return *o.ViewingPreferences
}

// GetViewingPreferencesOk returns a tuple with the ViewingPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboardUpdateRequest) GetViewingPreferencesOk() (*ViewingPreferences, bool) {
	if o == nil || o.ViewingPreferences == nil {
		return nil, false
	}
	return o.ViewingPreferences, true
}

// HasViewingPreferences returns a boolean if a field has been set.
func (o *SharedDashboardUpdateRequest) HasViewingPreferences() bool {
	return o != nil && o.ViewingPreferences != nil
}

// SetViewingPreferences gets a reference to the given ViewingPreferences and assigns it to the ViewingPreferences field.
func (o *SharedDashboardUpdateRequest) SetViewingPreferences(v ViewingPreferences) {
	o.ViewingPreferences = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SharedDashboardUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EmbeddableDomains != nil {
		toSerialize["embeddable_domains"] = o.EmbeddableDomains
	}
	if o.Expiration.IsSet() {
		toSerialize["expiration"] = o.Expiration.Get()
	}
	if o.GlobalTime.IsSet() {
		toSerialize["global_time"] = o.GlobalTime.Get()
	}
	if o.GlobalTimeSelectableEnabled.IsSet() {
		toSerialize["global_time_selectable_enabled"] = o.GlobalTimeSelectableEnabled.Get()
	}
	if o.Invitees != nil {
		toSerialize["invitees"] = o.Invitees
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
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.ViewingPreferences != nil {
		toSerialize["viewing_preferences"] = o.ViewingPreferences
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SharedDashboardUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EmbeddableDomains           []string                                       `json:"embeddable_domains,omitempty"`
		Expiration                  datadog.NullableTime                           `json:"expiration,omitempty"`
		GlobalTime                  NullableSharedDashboardUpdateRequestGlobalTime `json:"global_time,omitempty"`
		GlobalTimeSelectableEnabled datadog.NullableBool                           `json:"global_time_selectable_enabled,omitempty"`
		Invitees                    []SharedDashboardInviteesItems                 `json:"invitees,omitempty"`
		SelectableTemplateVars      []SelectableTemplateVariableItems              `json:"selectable_template_vars,omitempty"`
		ShareList                   datadog.NullableList[string]                   `json:"share_list,omitempty"`
		ShareType                   NullableDashboardShareType                     `json:"share_type,omitempty"`
		Status                      *SharedDashboardStatus                         `json:"status,omitempty"`
		Title                       *string                                        `json:"title,omitempty"`
		ViewingPreferences          *ViewingPreferences                            `json:"viewing_preferences,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"embeddable_domains", "expiration", "global_time", "global_time_selectable_enabled", "invitees", "selectable_template_vars", "share_list", "share_type", "status", "title", "viewing_preferences"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EmbeddableDomains = all.EmbeddableDomains
	o.Expiration = all.Expiration
	o.GlobalTime = all.GlobalTime
	o.GlobalTimeSelectableEnabled = all.GlobalTimeSelectableEnabled
	o.Invitees = all.Invitees
	o.SelectableTemplateVars = all.SelectableTemplateVars
	o.ShareList = all.ShareList
	if all.ShareType.Get() != nil && !all.ShareType.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.ShareType = all.ShareType
	}
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Title = all.Title
	if all.ViewingPreferences != nil && all.ViewingPreferences.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ViewingPreferences = all.ViewingPreferences

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
