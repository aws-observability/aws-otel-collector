// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardResponseAttributes Attributes of a shared dashboard response.
type SharedDashboardResponseAttributes struct {
	// Time when the shared dashboard was created.
	CreatedAt time.Time `json:"created_at"`
	// Domains where embed-type shared dashboards can be embedded.
	EmbeddableDomains []string `json:"embeddable_domains"`
	// Time when the shared dashboard expires.
	Expiration datadog.NullableTime `json:"expiration"`
	// Default time range configuration for the shared dashboard.
	GlobalTime map[string]interface{} `json:"global_time"`
	// Whether viewers can select a different global time setting.
	GlobalTimeSelectable bool `json:"global_time_selectable"`
	// Invitees for invite-only shared dashboards.
	Invitees []SharedDashboardInvitee `json:"invitees"`
	// Time when the shared dashboard was last accessed.
	LastAccessed datadog.NullableTime `json:"last_accessed"`
	// Template variables that viewers can modify.
	SelectableTemplateVars []SharedDashboardSelectableTemplateVariable `json:"selectable_template_vars"`
	// Type of dashboard sharing.
	ShareType SharedDashboardShareType `json:"share_type"`
	// Whether the user who shared the dashboard is disabled.
	SharerDisabled bool `json:"sharer_disabled"`
	// Status of the shared dashboard.
	Status SharedDashboardStatus `json:"status"`
	// Display title for the shared dashboard.
	Title string `json:"title"`
	// Token assigned to the shared dashboard.
	Token string `json:"token"`
	// URL for the shared dashboard.
	Url string `json:"url"`
	// Display settings for the shared dashboard.
	ViewingPreferences SharedDashboardViewingPreferences `json:"viewing_preferences"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSharedDashboardResponseAttributes instantiates a new SharedDashboardResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSharedDashboardResponseAttributes(createdAt time.Time, embeddableDomains []string, expiration datadog.NullableTime, globalTime map[string]interface{}, globalTimeSelectable bool, invitees []SharedDashboardInvitee, lastAccessed datadog.NullableTime, selectableTemplateVars []SharedDashboardSelectableTemplateVariable, shareType SharedDashboardShareType, sharerDisabled bool, status SharedDashboardStatus, title string, token string, url string, viewingPreferences SharedDashboardViewingPreferences) *SharedDashboardResponseAttributes {
	this := SharedDashboardResponseAttributes{}
	this.CreatedAt = createdAt
	this.EmbeddableDomains = embeddableDomains
	this.Expiration = expiration
	this.GlobalTime = globalTime
	this.GlobalTimeSelectable = globalTimeSelectable
	this.Invitees = invitees
	this.LastAccessed = lastAccessed
	this.SelectableTemplateVars = selectableTemplateVars
	this.ShareType = shareType
	this.SharerDisabled = sharerDisabled
	this.Status = status
	this.Title = title
	this.Token = token
	this.Url = url
	this.ViewingPreferences = viewingPreferences
	return &this
}

// NewSharedDashboardResponseAttributesWithDefaults instantiates a new SharedDashboardResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSharedDashboardResponseAttributesWithDefaults() *SharedDashboardResponseAttributes {
	this := SharedDashboardResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *SharedDashboardResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *SharedDashboardResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEmbeddableDomains returns the EmbeddableDomains field value.
func (o *SharedDashboardResponseAttributes) GetEmbeddableDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EmbeddableDomains
}

// GetEmbeddableDomainsOk returns a tuple with the EmbeddableDomains field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardResponseAttributes) GetEmbeddableDomainsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmbeddableDomains, true
}

// SetEmbeddableDomains sets field value.
func (o *SharedDashboardResponseAttributes) SetEmbeddableDomains(v []string) {
	o.EmbeddableDomains = v
}

// GetExpiration returns the Expiration field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *SharedDashboardResponseAttributes) GetExpiration() time.Time {
	if o == nil || o.Expiration.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboardResponseAttributes) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// SetExpiration sets field value.
func (o *SharedDashboardResponseAttributes) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}

// GetGlobalTime returns the GlobalTime field value.
// If the value is explicit nil, the zero value for map[string]interface{} will be returned.
func (o *SharedDashboardResponseAttributes) GetGlobalTime() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.GlobalTime
}

// GetGlobalTimeOk returns a tuple with the GlobalTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboardResponseAttributes) GetGlobalTimeOk() (*map[string]interface{}, bool) {
	if o == nil || o.GlobalTime == nil {
		return nil, false
	}
	return &o.GlobalTime, true
}

// SetGlobalTime sets field value.
func (o *SharedDashboardResponseAttributes) SetGlobalTime(v map[string]interface{}) {
	o.GlobalTime = v
}

// GetGlobalTimeSelectable returns the GlobalTimeSelectable field value.
func (o *SharedDashboardResponseAttributes) GetGlobalTimeSelectable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.GlobalTimeSelectable
}

// GetGlobalTimeSelectableOk returns a tuple with the GlobalTimeSelectable field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardResponseAttributes) GetGlobalTimeSelectableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalTimeSelectable, true
}

// SetGlobalTimeSelectable sets field value.
func (o *SharedDashboardResponseAttributes) SetGlobalTimeSelectable(v bool) {
	o.GlobalTimeSelectable = v
}

// GetInvitees returns the Invitees field value.
func (o *SharedDashboardResponseAttributes) GetInvitees() []SharedDashboardInvitee {
	if o == nil {
		var ret []SharedDashboardInvitee
		return ret
	}
	return o.Invitees
}

// GetInviteesOk returns a tuple with the Invitees field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardResponseAttributes) GetInviteesOk() (*[]SharedDashboardInvitee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invitees, true
}

// SetInvitees sets field value.
func (o *SharedDashboardResponseAttributes) SetInvitees(v []SharedDashboardInvitee) {
	o.Invitees = v
}

// GetLastAccessed returns the LastAccessed field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *SharedDashboardResponseAttributes) GetLastAccessed() time.Time {
	if o == nil || o.LastAccessed.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessed.Get()
}

// GetLastAccessedOk returns a tuple with the LastAccessed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SharedDashboardResponseAttributes) GetLastAccessedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastAccessed.Get(), o.LastAccessed.IsSet()
}

// SetLastAccessed sets field value.
func (o *SharedDashboardResponseAttributes) SetLastAccessed(v time.Time) {
	o.LastAccessed.Set(&v)
}

// GetSelectableTemplateVars returns the SelectableTemplateVars field value.
func (o *SharedDashboardResponseAttributes) GetSelectableTemplateVars() []SharedDashboardSelectableTemplateVariable {
	if o == nil {
		var ret []SharedDashboardSelectableTemplateVariable
		return ret
	}
	return o.SelectableTemplateVars
}

// GetSelectableTemplateVarsOk returns a tuple with the SelectableTemplateVars field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardResponseAttributes) GetSelectableTemplateVarsOk() (*[]SharedDashboardSelectableTemplateVariable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectableTemplateVars, true
}

// SetSelectableTemplateVars sets field value.
func (o *SharedDashboardResponseAttributes) SetSelectableTemplateVars(v []SharedDashboardSelectableTemplateVariable) {
	o.SelectableTemplateVars = v
}

// GetShareType returns the ShareType field value.
func (o *SharedDashboardResponseAttributes) GetShareType() SharedDashboardShareType {
	if o == nil {
		var ret SharedDashboardShareType
		return ret
	}
	return o.ShareType
}

// GetShareTypeOk returns a tuple with the ShareType field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardResponseAttributes) GetShareTypeOk() (*SharedDashboardShareType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShareType, true
}

// SetShareType sets field value.
func (o *SharedDashboardResponseAttributes) SetShareType(v SharedDashboardShareType) {
	o.ShareType = v
}

// GetSharerDisabled returns the SharerDisabled field value.
func (o *SharedDashboardResponseAttributes) GetSharerDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SharerDisabled
}

// GetSharerDisabledOk returns a tuple with the SharerDisabled field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardResponseAttributes) GetSharerDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SharerDisabled, true
}

// SetSharerDisabled sets field value.
func (o *SharedDashboardResponseAttributes) SetSharerDisabled(v bool) {
	o.SharerDisabled = v
}

// GetStatus returns the Status field value.
func (o *SharedDashboardResponseAttributes) GetStatus() SharedDashboardStatus {
	if o == nil {
		var ret SharedDashboardStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardResponseAttributes) GetStatusOk() (*SharedDashboardStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *SharedDashboardResponseAttributes) SetStatus(v SharedDashboardStatus) {
	o.Status = v
}

// GetTitle returns the Title field value.
func (o *SharedDashboardResponseAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardResponseAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *SharedDashboardResponseAttributes) SetTitle(v string) {
	o.Title = v
}

// GetToken returns the Token field value.
func (o *SharedDashboardResponseAttributes) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardResponseAttributes) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value.
func (o *SharedDashboardResponseAttributes) SetToken(v string) {
	o.Token = v
}

// GetUrl returns the Url field value.
func (o *SharedDashboardResponseAttributes) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardResponseAttributes) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *SharedDashboardResponseAttributes) SetUrl(v string) {
	o.Url = v
}

// GetViewingPreferences returns the ViewingPreferences field value.
func (o *SharedDashboardResponseAttributes) GetViewingPreferences() SharedDashboardViewingPreferences {
	if o == nil {
		var ret SharedDashboardViewingPreferences
		return ret
	}
	return o.ViewingPreferences
}

// GetViewingPreferencesOk returns a tuple with the ViewingPreferences field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardResponseAttributes) GetViewingPreferencesOk() (*SharedDashboardViewingPreferences, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewingPreferences, true
}

// SetViewingPreferences sets field value.
func (o *SharedDashboardResponseAttributes) SetViewingPreferences(v SharedDashboardViewingPreferences) {
	o.ViewingPreferences = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SharedDashboardResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["embeddable_domains"] = o.EmbeddableDomains
	toSerialize["expiration"] = o.Expiration.Get()
	if o.GlobalTime != nil {
		toSerialize["global_time"] = o.GlobalTime
	}
	toSerialize["global_time_selectable"] = o.GlobalTimeSelectable
	toSerialize["invitees"] = o.Invitees
	toSerialize["last_accessed"] = o.LastAccessed.Get()
	toSerialize["selectable_template_vars"] = o.SelectableTemplateVars
	toSerialize["share_type"] = o.ShareType
	toSerialize["sharer_disabled"] = o.SharerDisabled
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title
	toSerialize["token"] = o.Token
	toSerialize["url"] = o.Url
	toSerialize["viewing_preferences"] = o.ViewingPreferences

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SharedDashboardResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt              *time.Time                                   `json:"created_at"`
		EmbeddableDomains      *[]string                                    `json:"embeddable_domains"`
		Expiration             datadog.NullableTime                         `json:"expiration"`
		GlobalTime             map[string]interface{}                       `json:"global_time"`
		GlobalTimeSelectable   *bool                                        `json:"global_time_selectable"`
		Invitees               *[]SharedDashboardInvitee                    `json:"invitees"`
		LastAccessed           datadog.NullableTime                         `json:"last_accessed"`
		SelectableTemplateVars *[]SharedDashboardSelectableTemplateVariable `json:"selectable_template_vars"`
		ShareType              *SharedDashboardShareType                    `json:"share_type"`
		SharerDisabled         *bool                                        `json:"sharer_disabled"`
		Status                 *SharedDashboardStatus                       `json:"status"`
		Title                  *string                                      `json:"title"`
		Token                  *string                                      `json:"token"`
		Url                    *string                                      `json:"url"`
		ViewingPreferences     *SharedDashboardViewingPreferences           `json:"viewing_preferences"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.EmbeddableDomains == nil {
		return fmt.Errorf("required field embeddable_domains missing")
	}
	if !all.Expiration.IsSet() {
		return fmt.Errorf("required field expiration missing")
	}
	if all.GlobalTime == nil {
		return fmt.Errorf("required field global_time missing")
	}
	if all.GlobalTimeSelectable == nil {
		return fmt.Errorf("required field global_time_selectable missing")
	}
	if all.Invitees == nil {
		return fmt.Errorf("required field invitees missing")
	}
	if !all.LastAccessed.IsSet() {
		return fmt.Errorf("required field last_accessed missing")
	}
	if all.SelectableTemplateVars == nil {
		return fmt.Errorf("required field selectable_template_vars missing")
	}
	if all.ShareType == nil {
		return fmt.Errorf("required field share_type missing")
	}
	if all.SharerDisabled == nil {
		return fmt.Errorf("required field sharer_disabled missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Token == nil {
		return fmt.Errorf("required field token missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	if all.ViewingPreferences == nil {
		return fmt.Errorf("required field viewing_preferences missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "embeddable_domains", "expiration", "global_time", "global_time_selectable", "invitees", "last_accessed", "selectable_template_vars", "share_type", "sharer_disabled", "status", "title", "token", "url", "viewing_preferences"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.EmbeddableDomains = *all.EmbeddableDomains
	o.Expiration = all.Expiration
	o.GlobalTime = all.GlobalTime
	o.GlobalTimeSelectable = *all.GlobalTimeSelectable
	o.Invitees = *all.Invitees
	o.LastAccessed = all.LastAccessed
	o.SelectableTemplateVars = *all.SelectableTemplateVars
	if !all.ShareType.IsValid() {
		hasInvalidField = true
	} else {
		o.ShareType = *all.ShareType
	}
	o.SharerDisabled = *all.SharerDisabled
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Title = *all.Title
	o.Token = *all.Token
	o.Url = *all.Url
	if all.ViewingPreferences.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ViewingPreferences = *all.ViewingPreferences

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
