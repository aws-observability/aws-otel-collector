// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamAttributes Team attributes
type TeamAttributes struct {
	// Unicode representation of the avatar for the team, limited to a single grapheme
	Avatar datadog.NullableString `json:"avatar,omitempty"`
	// Banner selection for the team
	Banner datadog.NullableInt64 `json:"banner,omitempty"`
	// Creation date of the team
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Free-form markdown description/content for the team's homepage
	Description datadog.NullableString `json:"description,omitempty"`
	// The team's identifier
	Handle string `json:"handle"`
	// Collection of hidden modules for the team
	HiddenModules []string `json:"hidden_modules,omitempty"`
	// The number of links belonging to the team
	LinkCount *int32 `json:"link_count,omitempty"`
	// Modification date of the team
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The name of the team
	Name string `json:"name"`
	// A brief summary of the team, derived from the `description`
	Summary datadog.NullableString `json:"summary,omitempty"`
	// The number of users belonging to the team
	UserCount *int32 `json:"user_count,omitempty"`
	// Collection of visible modules for the team
	VisibleModules []string `json:"visible_modules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamAttributes instantiates a new TeamAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamAttributes(handle string, name string) *TeamAttributes {
	this := TeamAttributes{}
	this.Handle = handle
	this.Name = name
	return &this
}

// NewTeamAttributesWithDefaults instantiates a new TeamAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamAttributesWithDefaults() *TeamAttributes {
	this := TeamAttributes{}
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamAttributes) GetAvatar() string {
	if o == nil || o.Avatar.Get() == nil {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamAttributes) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *TeamAttributes) HasAvatar() bool {
	return o != nil && o.Avatar.IsSet()
}

// SetAvatar gets a reference to the given datadog.NullableString and assigns it to the Avatar field.
func (o *TeamAttributes) SetAvatar(v string) {
	o.Avatar.Set(&v)
}

// SetAvatarNil sets the value for Avatar to be an explicit nil.
func (o *TeamAttributes) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil.
func (o *TeamAttributes) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetBanner returns the Banner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamAttributes) GetBanner() int64 {
	if o == nil || o.Banner.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Banner.Get()
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamAttributes) GetBannerOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Banner.Get(), o.Banner.IsSet()
}

// HasBanner returns a boolean if a field has been set.
func (o *TeamAttributes) HasBanner() bool {
	return o != nil && o.Banner.IsSet()
}

// SetBanner gets a reference to the given datadog.NullableInt64 and assigns it to the Banner field.
func (o *TeamAttributes) SetBanner(v int64) {
	o.Banner.Set(&v)
}

// SetBannerNil sets the value for Banner to be an explicit nil.
func (o *TeamAttributes) SetBannerNil() {
	o.Banner.Set(nil)
}

// UnsetBanner ensures that no value is present for Banner, not even an explicit nil.
func (o *TeamAttributes) UnsetBanner() {
	o.Banner.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TeamAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TeamAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TeamAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TeamAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *TeamAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *TeamAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *TeamAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetHandle returns the Handle field value.
func (o *TeamAttributes) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *TeamAttributes) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *TeamAttributes) SetHandle(v string) {
	o.Handle = v
}

// GetHiddenModules returns the HiddenModules field value if set, zero value otherwise.
func (o *TeamAttributes) GetHiddenModules() []string {
	if o == nil || o.HiddenModules == nil {
		var ret []string
		return ret
	}
	return o.HiddenModules
}

// GetHiddenModulesOk returns a tuple with the HiddenModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAttributes) GetHiddenModulesOk() (*[]string, bool) {
	if o == nil || o.HiddenModules == nil {
		return nil, false
	}
	return &o.HiddenModules, true
}

// HasHiddenModules returns a boolean if a field has been set.
func (o *TeamAttributes) HasHiddenModules() bool {
	return o != nil && o.HiddenModules != nil
}

// SetHiddenModules gets a reference to the given []string and assigns it to the HiddenModules field.
func (o *TeamAttributes) SetHiddenModules(v []string) {
	o.HiddenModules = v
}

// GetLinkCount returns the LinkCount field value if set, zero value otherwise.
func (o *TeamAttributes) GetLinkCount() int32 {
	if o == nil || o.LinkCount == nil {
		var ret int32
		return ret
	}
	return *o.LinkCount
}

// GetLinkCountOk returns a tuple with the LinkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAttributes) GetLinkCountOk() (*int32, bool) {
	if o == nil || o.LinkCount == nil {
		return nil, false
	}
	return o.LinkCount, true
}

// HasLinkCount returns a boolean if a field has been set.
func (o *TeamAttributes) HasLinkCount() bool {
	return o != nil && o.LinkCount != nil
}

// SetLinkCount gets a reference to the given int32 and assigns it to the LinkCount field.
func (o *TeamAttributes) SetLinkCount(v int32) {
	o.LinkCount = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *TeamAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *TeamAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *TeamAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value.
func (o *TeamAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TeamAttributes) SetName(v string) {
	o.Name = v
}

// GetSummary returns the Summary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamAttributes) GetSummary() string {
	if o == nil || o.Summary.Get() == nil {
		var ret string
		return ret
	}
	return *o.Summary.Get()
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamAttributes) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Summary.Get(), o.Summary.IsSet()
}

// HasSummary returns a boolean if a field has been set.
func (o *TeamAttributes) HasSummary() bool {
	return o != nil && o.Summary.IsSet()
}

// SetSummary gets a reference to the given datadog.NullableString and assigns it to the Summary field.
func (o *TeamAttributes) SetSummary(v string) {
	o.Summary.Set(&v)
}

// SetSummaryNil sets the value for Summary to be an explicit nil.
func (o *TeamAttributes) SetSummaryNil() {
	o.Summary.Set(nil)
}

// UnsetSummary ensures that no value is present for Summary, not even an explicit nil.
func (o *TeamAttributes) UnsetSummary() {
	o.Summary.Unset()
}

// GetUserCount returns the UserCount field value if set, zero value otherwise.
func (o *TeamAttributes) GetUserCount() int32 {
	if o == nil || o.UserCount == nil {
		var ret int32
		return ret
	}
	return *o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAttributes) GetUserCountOk() (*int32, bool) {
	if o == nil || o.UserCount == nil {
		return nil, false
	}
	return o.UserCount, true
}

// HasUserCount returns a boolean if a field has been set.
func (o *TeamAttributes) HasUserCount() bool {
	return o != nil && o.UserCount != nil
}

// SetUserCount gets a reference to the given int32 and assigns it to the UserCount field.
func (o *TeamAttributes) SetUserCount(v int32) {
	o.UserCount = &v
}

// GetVisibleModules returns the VisibleModules field value if set, zero value otherwise.
func (o *TeamAttributes) GetVisibleModules() []string {
	if o == nil || o.VisibleModules == nil {
		var ret []string
		return ret
	}
	return o.VisibleModules
}

// GetVisibleModulesOk returns a tuple with the VisibleModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAttributes) GetVisibleModulesOk() (*[]string, bool) {
	if o == nil || o.VisibleModules == nil {
		return nil, false
	}
	return &o.VisibleModules, true
}

// HasVisibleModules returns a boolean if a field has been set.
func (o *TeamAttributes) HasVisibleModules() bool {
	return o != nil && o.VisibleModules != nil
}

// SetVisibleModules gets a reference to the given []string and assigns it to the VisibleModules field.
func (o *TeamAttributes) SetVisibleModules(v []string) {
	o.VisibleModules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	if o.Banner.IsSet() {
		toSerialize["banner"] = o.Banner.Get()
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["handle"] = o.Handle
	if o.HiddenModules != nil {
		toSerialize["hidden_modules"] = o.HiddenModules
	}
	if o.LinkCount != nil {
		toSerialize["link_count"] = o.LinkCount
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["name"] = o.Name
	if o.Summary.IsSet() {
		toSerialize["summary"] = o.Summary.Get()
	}
	if o.UserCount != nil {
		toSerialize["user_count"] = o.UserCount
	}
	if o.VisibleModules != nil {
		toSerialize["visible_modules"] = o.VisibleModules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Avatar         datadog.NullableString `json:"avatar,omitempty"`
		Banner         datadog.NullableInt64  `json:"banner,omitempty"`
		CreatedAt      *time.Time             `json:"created_at,omitempty"`
		Description    datadog.NullableString `json:"description,omitempty"`
		Handle         *string                `json:"handle"`
		HiddenModules  []string               `json:"hidden_modules,omitempty"`
		LinkCount      *int32                 `json:"link_count,omitempty"`
		ModifiedAt     *time.Time             `json:"modified_at,omitempty"`
		Name           *string                `json:"name"`
		Summary        datadog.NullableString `json:"summary,omitempty"`
		UserCount      *int32                 `json:"user_count,omitempty"`
		VisibleModules []string               `json:"visible_modules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"avatar", "banner", "created_at", "description", "handle", "hidden_modules", "link_count", "modified_at", "name", "summary", "user_count", "visible_modules"})
	} else {
		return err
	}
	o.Avatar = all.Avatar
	o.Banner = all.Banner
	o.CreatedAt = all.CreatedAt
	o.Description = all.Description
	o.Handle = *all.Handle
	o.HiddenModules = all.HiddenModules
	o.LinkCount = all.LinkCount
	o.ModifiedAt = all.ModifiedAt
	o.Name = *all.Name
	o.Summary = all.Summary
	o.UserCount = all.UserCount
	o.VisibleModules = all.VisibleModules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
