// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuthNMappingTeamAttributes Team attributes.
type AuthNMappingTeamAttributes struct {
	// Unicode representation of the avatar for the team, limited to a single grapheme
	Avatar datadog.NullableString `json:"avatar,omitempty"`
	// Banner selection for the team
	Banner datadog.NullableInt64 `json:"banner,omitempty"`
	// The team's identifier
	Handle *string `json:"handle,omitempty"`
	// The number of links belonging to the team
	LinkCount *int32 `json:"link_count,omitempty"`
	// The name of the team
	Name *string `json:"name,omitempty"`
	// A brief summary of the team, derived from the `description`
	Summary datadog.NullableString `json:"summary,omitempty"`
	// The number of users belonging to the team
	UserCount *int32 `json:"user_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAuthNMappingTeamAttributes instantiates a new AuthNMappingTeamAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthNMappingTeamAttributes() *AuthNMappingTeamAttributes {
	this := AuthNMappingTeamAttributes{}
	return &this
}

// NewAuthNMappingTeamAttributesWithDefaults instantiates a new AuthNMappingTeamAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAuthNMappingTeamAttributesWithDefaults() *AuthNMappingTeamAttributes {
	this := AuthNMappingTeamAttributes{}
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthNMappingTeamAttributes) GetAvatar() string {
	if o == nil || o.Avatar.Get() == nil {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AuthNMappingTeamAttributes) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *AuthNMappingTeamAttributes) HasAvatar() bool {
	return o != nil && o.Avatar.IsSet()
}

// SetAvatar gets a reference to the given datadog.NullableString and assigns it to the Avatar field.
func (o *AuthNMappingTeamAttributes) SetAvatar(v string) {
	o.Avatar.Set(&v)
}

// SetAvatarNil sets the value for Avatar to be an explicit nil.
func (o *AuthNMappingTeamAttributes) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil.
func (o *AuthNMappingTeamAttributes) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetBanner returns the Banner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthNMappingTeamAttributes) GetBanner() int64 {
	if o == nil || o.Banner.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Banner.Get()
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AuthNMappingTeamAttributes) GetBannerOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Banner.Get(), o.Banner.IsSet()
}

// HasBanner returns a boolean if a field has been set.
func (o *AuthNMappingTeamAttributes) HasBanner() bool {
	return o != nil && o.Banner.IsSet()
}

// SetBanner gets a reference to the given datadog.NullableInt64 and assigns it to the Banner field.
func (o *AuthNMappingTeamAttributes) SetBanner(v int64) {
	o.Banner.Set(&v)
}

// SetBannerNil sets the value for Banner to be an explicit nil.
func (o *AuthNMappingTeamAttributes) SetBannerNil() {
	o.Banner.Set(nil)
}

// UnsetBanner ensures that no value is present for Banner, not even an explicit nil.
func (o *AuthNMappingTeamAttributes) UnsetBanner() {
	o.Banner.Unset()
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *AuthNMappingTeamAttributes) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthNMappingTeamAttributes) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *AuthNMappingTeamAttributes) HasHandle() bool {
	return o != nil && o.Handle != nil
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *AuthNMappingTeamAttributes) SetHandle(v string) {
	o.Handle = &v
}

// GetLinkCount returns the LinkCount field value if set, zero value otherwise.
func (o *AuthNMappingTeamAttributes) GetLinkCount() int32 {
	if o == nil || o.LinkCount == nil {
		var ret int32
		return ret
	}
	return *o.LinkCount
}

// GetLinkCountOk returns a tuple with the LinkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthNMappingTeamAttributes) GetLinkCountOk() (*int32, bool) {
	if o == nil || o.LinkCount == nil {
		return nil, false
	}
	return o.LinkCount, true
}

// HasLinkCount returns a boolean if a field has been set.
func (o *AuthNMappingTeamAttributes) HasLinkCount() bool {
	return o != nil && o.LinkCount != nil
}

// SetLinkCount gets a reference to the given int32 and assigns it to the LinkCount field.
func (o *AuthNMappingTeamAttributes) SetLinkCount(v int32) {
	o.LinkCount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthNMappingTeamAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthNMappingTeamAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthNMappingTeamAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthNMappingTeamAttributes) SetName(v string) {
	o.Name = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthNMappingTeamAttributes) GetSummary() string {
	if o == nil || o.Summary.Get() == nil {
		var ret string
		return ret
	}
	return *o.Summary.Get()
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AuthNMappingTeamAttributes) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Summary.Get(), o.Summary.IsSet()
}

// HasSummary returns a boolean if a field has been set.
func (o *AuthNMappingTeamAttributes) HasSummary() bool {
	return o != nil && o.Summary.IsSet()
}

// SetSummary gets a reference to the given datadog.NullableString and assigns it to the Summary field.
func (o *AuthNMappingTeamAttributes) SetSummary(v string) {
	o.Summary.Set(&v)
}

// SetSummaryNil sets the value for Summary to be an explicit nil.
func (o *AuthNMappingTeamAttributes) SetSummaryNil() {
	o.Summary.Set(nil)
}

// UnsetSummary ensures that no value is present for Summary, not even an explicit nil.
func (o *AuthNMappingTeamAttributes) UnsetSummary() {
	o.Summary.Unset()
}

// GetUserCount returns the UserCount field value if set, zero value otherwise.
func (o *AuthNMappingTeamAttributes) GetUserCount() int32 {
	if o == nil || o.UserCount == nil {
		var ret int32
		return ret
	}
	return *o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthNMappingTeamAttributes) GetUserCountOk() (*int32, bool) {
	if o == nil || o.UserCount == nil {
		return nil, false
	}
	return o.UserCount, true
}

// HasUserCount returns a boolean if a field has been set.
func (o *AuthNMappingTeamAttributes) HasUserCount() bool {
	return o != nil && o.UserCount != nil
}

// SetUserCount gets a reference to the given int32 and assigns it to the UserCount field.
func (o *AuthNMappingTeamAttributes) SetUserCount(v int32) {
	o.UserCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AuthNMappingTeamAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.LinkCount != nil {
		toSerialize["link_count"] = o.LinkCount
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Summary.IsSet() {
		toSerialize["summary"] = o.Summary.Get()
	}
	if o.UserCount != nil {
		toSerialize["user_count"] = o.UserCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AuthNMappingTeamAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Avatar    datadog.NullableString `json:"avatar,omitempty"`
		Banner    datadog.NullableInt64  `json:"banner,omitempty"`
		Handle    *string                `json:"handle,omitempty"`
		LinkCount *int32                 `json:"link_count,omitempty"`
		Name      *string                `json:"name,omitempty"`
		Summary   datadog.NullableString `json:"summary,omitempty"`
		UserCount *int32                 `json:"user_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"avatar", "banner", "handle", "link_count", "name", "summary", "user_count"})
	} else {
		return err
	}
	o.Avatar = all.Avatar
	o.Banner = all.Banner
	o.Handle = all.Handle
	o.LinkCount = all.LinkCount
	o.Name = all.Name
	o.Summary = all.Summary
	o.UserCount = all.UserCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
