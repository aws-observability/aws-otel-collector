// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyLinkTeamAttributes Team hierarchy links connect different teams. This represents attributes from teams that are connected by the team hierarchy link.
type TeamHierarchyLinkTeamAttributes struct {
	// The team's avatar
	Avatar datadog.NullableString `json:"avatar,omitempty"`
	// The team's banner
	Banner *int64 `json:"banner,omitempty"`
	// The team's handle
	Handle string `json:"handle"`
	// Whether the team is managed
	IsManaged *bool `json:"is_managed,omitempty"`
	// Whether the team has open membership
	IsOpenMembership *bool `json:"is_open_membership,omitempty"`
	// The number of links for the team
	LinkCount *int64 `json:"link_count,omitempty"`
	// The team's name
	Name string `json:"name"`
	// The team's summary
	Summary datadog.NullableString `json:"summary,omitempty"`
	// The number of users in the team
	UserCount *int64 `json:"user_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyLinkTeamAttributes instantiates a new TeamHierarchyLinkTeamAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyLinkTeamAttributes(handle string, name string) *TeamHierarchyLinkTeamAttributes {
	this := TeamHierarchyLinkTeamAttributes{}
	this.Handle = handle
	this.Name = name
	return &this
}

// NewTeamHierarchyLinkTeamAttributesWithDefaults instantiates a new TeamHierarchyLinkTeamAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyLinkTeamAttributesWithDefaults() *TeamHierarchyLinkTeamAttributes {
	this := TeamHierarchyLinkTeamAttributes{}
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamHierarchyLinkTeamAttributes) GetAvatar() string {
	if o == nil || o.Avatar.Get() == nil {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamHierarchyLinkTeamAttributes) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *TeamHierarchyLinkTeamAttributes) HasAvatar() bool {
	return o != nil && o.Avatar.IsSet()
}

// SetAvatar gets a reference to the given datadog.NullableString and assigns it to the Avatar field.
func (o *TeamHierarchyLinkTeamAttributes) SetAvatar(v string) {
	o.Avatar.Set(&v)
}

// SetAvatarNil sets the value for Avatar to be an explicit nil.
func (o *TeamHierarchyLinkTeamAttributes) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil.
func (o *TeamHierarchyLinkTeamAttributes) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *TeamHierarchyLinkTeamAttributes) GetBanner() int64 {
	if o == nil || o.Banner == nil {
		var ret int64
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkTeamAttributes) GetBannerOk() (*int64, bool) {
	if o == nil || o.Banner == nil {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *TeamHierarchyLinkTeamAttributes) HasBanner() bool {
	return o != nil && o.Banner != nil
}

// SetBanner gets a reference to the given int64 and assigns it to the Banner field.
func (o *TeamHierarchyLinkTeamAttributes) SetBanner(v int64) {
	o.Banner = &v
}

// GetHandle returns the Handle field value.
func (o *TeamHierarchyLinkTeamAttributes) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkTeamAttributes) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *TeamHierarchyLinkTeamAttributes) SetHandle(v string) {
	o.Handle = v
}

// GetIsManaged returns the IsManaged field value if set, zero value otherwise.
func (o *TeamHierarchyLinkTeamAttributes) GetIsManaged() bool {
	if o == nil || o.IsManaged == nil {
		var ret bool
		return ret
	}
	return *o.IsManaged
}

// GetIsManagedOk returns a tuple with the IsManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkTeamAttributes) GetIsManagedOk() (*bool, bool) {
	if o == nil || o.IsManaged == nil {
		return nil, false
	}
	return o.IsManaged, true
}

// HasIsManaged returns a boolean if a field has been set.
func (o *TeamHierarchyLinkTeamAttributes) HasIsManaged() bool {
	return o != nil && o.IsManaged != nil
}

// SetIsManaged gets a reference to the given bool and assigns it to the IsManaged field.
func (o *TeamHierarchyLinkTeamAttributes) SetIsManaged(v bool) {
	o.IsManaged = &v
}

// GetIsOpenMembership returns the IsOpenMembership field value if set, zero value otherwise.
func (o *TeamHierarchyLinkTeamAttributes) GetIsOpenMembership() bool {
	if o == nil || o.IsOpenMembership == nil {
		var ret bool
		return ret
	}
	return *o.IsOpenMembership
}

// GetIsOpenMembershipOk returns a tuple with the IsOpenMembership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkTeamAttributes) GetIsOpenMembershipOk() (*bool, bool) {
	if o == nil || o.IsOpenMembership == nil {
		return nil, false
	}
	return o.IsOpenMembership, true
}

// HasIsOpenMembership returns a boolean if a field has been set.
func (o *TeamHierarchyLinkTeamAttributes) HasIsOpenMembership() bool {
	return o != nil && o.IsOpenMembership != nil
}

// SetIsOpenMembership gets a reference to the given bool and assigns it to the IsOpenMembership field.
func (o *TeamHierarchyLinkTeamAttributes) SetIsOpenMembership(v bool) {
	o.IsOpenMembership = &v
}

// GetLinkCount returns the LinkCount field value if set, zero value otherwise.
func (o *TeamHierarchyLinkTeamAttributes) GetLinkCount() int64 {
	if o == nil || o.LinkCount == nil {
		var ret int64
		return ret
	}
	return *o.LinkCount
}

// GetLinkCountOk returns a tuple with the LinkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkTeamAttributes) GetLinkCountOk() (*int64, bool) {
	if o == nil || o.LinkCount == nil {
		return nil, false
	}
	return o.LinkCount, true
}

// HasLinkCount returns a boolean if a field has been set.
func (o *TeamHierarchyLinkTeamAttributes) HasLinkCount() bool {
	return o != nil && o.LinkCount != nil
}

// SetLinkCount gets a reference to the given int64 and assigns it to the LinkCount field.
func (o *TeamHierarchyLinkTeamAttributes) SetLinkCount(v int64) {
	o.LinkCount = &v
}

// GetName returns the Name field value.
func (o *TeamHierarchyLinkTeamAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkTeamAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TeamHierarchyLinkTeamAttributes) SetName(v string) {
	o.Name = v
}

// GetSummary returns the Summary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamHierarchyLinkTeamAttributes) GetSummary() string {
	if o == nil || o.Summary.Get() == nil {
		var ret string
		return ret
	}
	return *o.Summary.Get()
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamHierarchyLinkTeamAttributes) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Summary.Get(), o.Summary.IsSet()
}

// HasSummary returns a boolean if a field has been set.
func (o *TeamHierarchyLinkTeamAttributes) HasSummary() bool {
	return o != nil && o.Summary.IsSet()
}

// SetSummary gets a reference to the given datadog.NullableString and assigns it to the Summary field.
func (o *TeamHierarchyLinkTeamAttributes) SetSummary(v string) {
	o.Summary.Set(&v)
}

// SetSummaryNil sets the value for Summary to be an explicit nil.
func (o *TeamHierarchyLinkTeamAttributes) SetSummaryNil() {
	o.Summary.Set(nil)
}

// UnsetSummary ensures that no value is present for Summary, not even an explicit nil.
func (o *TeamHierarchyLinkTeamAttributes) UnsetSummary() {
	o.Summary.Unset()
}

// GetUserCount returns the UserCount field value if set, zero value otherwise.
func (o *TeamHierarchyLinkTeamAttributes) GetUserCount() int64 {
	if o == nil || o.UserCount == nil {
		var ret int64
		return ret
	}
	return *o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkTeamAttributes) GetUserCountOk() (*int64, bool) {
	if o == nil || o.UserCount == nil {
		return nil, false
	}
	return o.UserCount, true
}

// HasUserCount returns a boolean if a field has been set.
func (o *TeamHierarchyLinkTeamAttributes) HasUserCount() bool {
	return o != nil && o.UserCount != nil
}

// SetUserCount gets a reference to the given int64 and assigns it to the UserCount field.
func (o *TeamHierarchyLinkTeamAttributes) SetUserCount(v int64) {
	o.UserCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyLinkTeamAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	if o.Banner != nil {
		toSerialize["banner"] = o.Banner
	}
	toSerialize["handle"] = o.Handle
	if o.IsManaged != nil {
		toSerialize["is_managed"] = o.IsManaged
	}
	if o.IsOpenMembership != nil {
		toSerialize["is_open_membership"] = o.IsOpenMembership
	}
	if o.LinkCount != nil {
		toSerialize["link_count"] = o.LinkCount
	}
	toSerialize["name"] = o.Name
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
func (o *TeamHierarchyLinkTeamAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Avatar           datadog.NullableString `json:"avatar,omitempty"`
		Banner           *int64                 `json:"banner,omitempty"`
		Handle           *string                `json:"handle"`
		IsManaged        *bool                  `json:"is_managed,omitempty"`
		IsOpenMembership *bool                  `json:"is_open_membership,omitempty"`
		LinkCount        *int64                 `json:"link_count,omitempty"`
		Name             *string                `json:"name"`
		Summary          datadog.NullableString `json:"summary,omitempty"`
		UserCount        *int64                 `json:"user_count,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"avatar", "banner", "handle", "is_managed", "is_open_membership", "link_count", "name", "summary", "user_count"})
	} else {
		return err
	}
	o.Avatar = all.Avatar
	o.Banner = all.Banner
	o.Handle = *all.Handle
	o.IsManaged = all.IsManaged
	o.IsOpenMembership = all.IsOpenMembership
	o.LinkCount = all.LinkCount
	o.Name = *all.Name
	o.Summary = all.Summary
	o.UserCount = all.UserCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
