// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardUpdateRequest Update a shared dashboard's settings.
type SharedDashboardUpdateRequest struct {
	// Timeframe setting for the shared dashboard.
	GlobalTime NullableSharedDashboardUpdateRequestGlobalTime `json:"global_time"`
	// Whether to allow viewers to select a different global time setting for the shared dashboard.
	GlobalTimeSelectableEnabled datadog.NullableBool `json:"global_time_selectable_enabled,omitempty"`
	// List of objects representing template variables on the shared dashboard which can have selectable values.
	SelectableTemplateVars []SelectableTemplateVariableItems `json:"selectable_template_vars,omitempty"`
	// List of email addresses that can be given access to the shared dashboard.
	ShareList datadog.NullableList[string] `json:"share_list,omitempty"`
	// Type of sharing access (either open to anyone who has the public URL or invite-only).
	ShareType NullableDashboardShareType `json:"share_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSharedDashboardUpdateRequest instantiates a new SharedDashboardUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSharedDashboardUpdateRequest(globalTime NullableSharedDashboardUpdateRequestGlobalTime) *SharedDashboardUpdateRequest {
	this := SharedDashboardUpdateRequest{}
	this.GlobalTime = globalTime
	return &this
}

// NewSharedDashboardUpdateRequestWithDefaults instantiates a new SharedDashboardUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSharedDashboardUpdateRequestWithDefaults() *SharedDashboardUpdateRequest {
	this := SharedDashboardUpdateRequest{}
	return &this
}

// GetGlobalTime returns the GlobalTime field value.
// If the value is explicit nil, the zero value for SharedDashboardUpdateRequestGlobalTime will be returned.
func (o *SharedDashboardUpdateRequest) GetGlobalTime() SharedDashboardUpdateRequestGlobalTime {
	if o == nil || o.GlobalTime.Get() == nil {
		var ret SharedDashboardUpdateRequestGlobalTime
		return ret
	}
	return *o.GlobalTime.Get()
}

// GetGlobalTimeOk returns a tuple with the GlobalTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboardUpdateRequest) GetGlobalTimeOk() (*SharedDashboardUpdateRequestGlobalTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.GlobalTime.Get(), o.GlobalTime.IsSet()
}

// SetGlobalTime sets field value.
func (o *SharedDashboardUpdateRequest) SetGlobalTime(v SharedDashboardUpdateRequestGlobalTime) {
	o.GlobalTime.Set(&v)
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

// MarshalJSON serializes the struct using spec logic.
func (o SharedDashboardUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["global_time"] = o.GlobalTime.Get()
	if o.GlobalTimeSelectableEnabled.IsSet() {
		toSerialize["global_time_selectable_enabled"] = o.GlobalTimeSelectableEnabled.Get()
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SharedDashboardUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GlobalTime                  NullableSharedDashboardUpdateRequestGlobalTime `json:"global_time"`
		GlobalTimeSelectableEnabled datadog.NullableBool                           `json:"global_time_selectable_enabled,omitempty"`
		SelectableTemplateVars      []SelectableTemplateVariableItems              `json:"selectable_template_vars,omitempty"`
		ShareList                   datadog.NullableList[string]                   `json:"share_list,omitempty"`
		ShareType                   NullableDashboardShareType                     `json:"share_type,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.GlobalTime.IsSet() {
		return fmt.Errorf("required field global_time missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"global_time", "global_time_selectable_enabled", "selectable_template_vars", "share_list", "share_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.GlobalTime = all.GlobalTime
	o.GlobalTimeSelectableEnabled = all.GlobalTimeSelectableEnabled
	o.SelectableTemplateVars = all.SelectableTemplateVars
	o.ShareList = all.ShareList
	if all.ShareType.Get() != nil && !all.ShareType.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.ShareType = all.ShareType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
