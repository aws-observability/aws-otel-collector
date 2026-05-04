// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecureEmbedUpdateResponseAttributes Attributes of an updated secure embed shared dashboard.
type SecureEmbedUpdateResponseAttributes struct {
	// Creation timestamp.
	CreatedAt *string `json:"created_at,omitempty"`
	// Last 4 characters of the credential. Defaults to `0000` if unavailable.
	CredentialSuffix *string `json:"credential_suffix,omitempty"`
	// The source dashboard ID.
	DashboardId *string `json:"dashboard_id,omitempty"`
	// Default time range configuration for the secure embed.
	GlobalTime *SecureEmbedGlobalTime `json:"global_time,omitempty"`
	// Whether time range is viewer-selectable.
	GlobalTimeSelectable *bool `json:"global_time_selectable,omitempty"`
	// Internal share ID.
	Id *string `json:"id,omitempty"`
	// Template variables with their configuration.
	SelectableTemplateVars []SecureEmbedSelectableTemplateVariable `json:"selectable_template_vars,omitempty"`
	// The type of share. Always `secure_embed`.
	ShareType *SecureEmbedShareType `json:"share_type,omitempty"`
	// The status of the secure embed share. Active means the shared dashboard is available. Paused means it is not.
	Status *SecureEmbedStatus `json:"status,omitempty"`
	// Display title.
	Title *string `json:"title,omitempty"`
	// Public share token.
	Token *string `json:"token,omitempty"`
	// CDN URL for the shared dashboard.
	Url *string `json:"url,omitempty"`
	// Display settings for the secure embed shared dashboard.
	ViewingPreferences *SecureEmbedViewingPreferences `json:"viewing_preferences,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecureEmbedUpdateResponseAttributes instantiates a new SecureEmbedUpdateResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecureEmbedUpdateResponseAttributes() *SecureEmbedUpdateResponseAttributes {
	this := SecureEmbedUpdateResponseAttributes{}
	return &this
}

// NewSecureEmbedUpdateResponseAttributesWithDefaults instantiates a new SecureEmbedUpdateResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecureEmbedUpdateResponseAttributesWithDefaults() *SecureEmbedUpdateResponseAttributes {
	this := SecureEmbedUpdateResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SecureEmbedUpdateResponseAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateResponseAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SecureEmbedUpdateResponseAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SecureEmbedUpdateResponseAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCredentialSuffix returns the CredentialSuffix field value if set, zero value otherwise.
func (o *SecureEmbedUpdateResponseAttributes) GetCredentialSuffix() string {
	if o == nil || o.CredentialSuffix == nil {
		var ret string
		return ret
	}
	return *o.CredentialSuffix
}

// GetCredentialSuffixOk returns a tuple with the CredentialSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateResponseAttributes) GetCredentialSuffixOk() (*string, bool) {
	if o == nil || o.CredentialSuffix == nil {
		return nil, false
	}
	return o.CredentialSuffix, true
}

// HasCredentialSuffix returns a boolean if a field has been set.
func (o *SecureEmbedUpdateResponseAttributes) HasCredentialSuffix() bool {
	return o != nil && o.CredentialSuffix != nil
}

// SetCredentialSuffix gets a reference to the given string and assigns it to the CredentialSuffix field.
func (o *SecureEmbedUpdateResponseAttributes) SetCredentialSuffix(v string) {
	o.CredentialSuffix = &v
}

// GetDashboardId returns the DashboardId field value if set, zero value otherwise.
func (o *SecureEmbedUpdateResponseAttributes) GetDashboardId() string {
	if o == nil || o.DashboardId == nil {
		var ret string
		return ret
	}
	return *o.DashboardId
}

// GetDashboardIdOk returns a tuple with the DashboardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateResponseAttributes) GetDashboardIdOk() (*string, bool) {
	if o == nil || o.DashboardId == nil {
		return nil, false
	}
	return o.DashboardId, true
}

// HasDashboardId returns a boolean if a field has been set.
func (o *SecureEmbedUpdateResponseAttributes) HasDashboardId() bool {
	return o != nil && o.DashboardId != nil
}

// SetDashboardId gets a reference to the given string and assigns it to the DashboardId field.
func (o *SecureEmbedUpdateResponseAttributes) SetDashboardId(v string) {
	o.DashboardId = &v
}

// GetGlobalTime returns the GlobalTime field value if set, zero value otherwise.
func (o *SecureEmbedUpdateResponseAttributes) GetGlobalTime() SecureEmbedGlobalTime {
	if o == nil || o.GlobalTime == nil {
		var ret SecureEmbedGlobalTime
		return ret
	}
	return *o.GlobalTime
}

// GetGlobalTimeOk returns a tuple with the GlobalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateResponseAttributes) GetGlobalTimeOk() (*SecureEmbedGlobalTime, bool) {
	if o == nil || o.GlobalTime == nil {
		return nil, false
	}
	return o.GlobalTime, true
}

// HasGlobalTime returns a boolean if a field has been set.
func (o *SecureEmbedUpdateResponseAttributes) HasGlobalTime() bool {
	return o != nil && o.GlobalTime != nil
}

// SetGlobalTime gets a reference to the given SecureEmbedGlobalTime and assigns it to the GlobalTime field.
func (o *SecureEmbedUpdateResponseAttributes) SetGlobalTime(v SecureEmbedGlobalTime) {
	o.GlobalTime = &v
}

// GetGlobalTimeSelectable returns the GlobalTimeSelectable field value if set, zero value otherwise.
func (o *SecureEmbedUpdateResponseAttributes) GetGlobalTimeSelectable() bool {
	if o == nil || o.GlobalTimeSelectable == nil {
		var ret bool
		return ret
	}
	return *o.GlobalTimeSelectable
}

// GetGlobalTimeSelectableOk returns a tuple with the GlobalTimeSelectable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateResponseAttributes) GetGlobalTimeSelectableOk() (*bool, bool) {
	if o == nil || o.GlobalTimeSelectable == nil {
		return nil, false
	}
	return o.GlobalTimeSelectable, true
}

// HasGlobalTimeSelectable returns a boolean if a field has been set.
func (o *SecureEmbedUpdateResponseAttributes) HasGlobalTimeSelectable() bool {
	return o != nil && o.GlobalTimeSelectable != nil
}

// SetGlobalTimeSelectable gets a reference to the given bool and assigns it to the GlobalTimeSelectable field.
func (o *SecureEmbedUpdateResponseAttributes) SetGlobalTimeSelectable(v bool) {
	o.GlobalTimeSelectable = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecureEmbedUpdateResponseAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateResponseAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecureEmbedUpdateResponseAttributes) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SecureEmbedUpdateResponseAttributes) SetId(v string) {
	o.Id = &v
}

// GetSelectableTemplateVars returns the SelectableTemplateVars field value if set, zero value otherwise.
func (o *SecureEmbedUpdateResponseAttributes) GetSelectableTemplateVars() []SecureEmbedSelectableTemplateVariable {
	if o == nil || o.SelectableTemplateVars == nil {
		var ret []SecureEmbedSelectableTemplateVariable
		return ret
	}
	return o.SelectableTemplateVars
}

// GetSelectableTemplateVarsOk returns a tuple with the SelectableTemplateVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateResponseAttributes) GetSelectableTemplateVarsOk() (*[]SecureEmbedSelectableTemplateVariable, bool) {
	if o == nil || o.SelectableTemplateVars == nil {
		return nil, false
	}
	return &o.SelectableTemplateVars, true
}

// HasSelectableTemplateVars returns a boolean if a field has been set.
func (o *SecureEmbedUpdateResponseAttributes) HasSelectableTemplateVars() bool {
	return o != nil && o.SelectableTemplateVars != nil
}

// SetSelectableTemplateVars gets a reference to the given []SecureEmbedSelectableTemplateVariable and assigns it to the SelectableTemplateVars field.
func (o *SecureEmbedUpdateResponseAttributes) SetSelectableTemplateVars(v []SecureEmbedSelectableTemplateVariable) {
	o.SelectableTemplateVars = v
}

// GetShareType returns the ShareType field value if set, zero value otherwise.
func (o *SecureEmbedUpdateResponseAttributes) GetShareType() SecureEmbedShareType {
	if o == nil || o.ShareType == nil {
		var ret SecureEmbedShareType
		return ret
	}
	return *o.ShareType
}

// GetShareTypeOk returns a tuple with the ShareType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateResponseAttributes) GetShareTypeOk() (*SecureEmbedShareType, bool) {
	if o == nil || o.ShareType == nil {
		return nil, false
	}
	return o.ShareType, true
}

// HasShareType returns a boolean if a field has been set.
func (o *SecureEmbedUpdateResponseAttributes) HasShareType() bool {
	return o != nil && o.ShareType != nil
}

// SetShareType gets a reference to the given SecureEmbedShareType and assigns it to the ShareType field.
func (o *SecureEmbedUpdateResponseAttributes) SetShareType(v SecureEmbedShareType) {
	o.ShareType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecureEmbedUpdateResponseAttributes) GetStatus() SecureEmbedStatus {
	if o == nil || o.Status == nil {
		var ret SecureEmbedStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateResponseAttributes) GetStatusOk() (*SecureEmbedStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SecureEmbedUpdateResponseAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given SecureEmbedStatus and assigns it to the Status field.
func (o *SecureEmbedUpdateResponseAttributes) SetStatus(v SecureEmbedStatus) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SecureEmbedUpdateResponseAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateResponseAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SecureEmbedUpdateResponseAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SecureEmbedUpdateResponseAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SecureEmbedUpdateResponseAttributes) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateResponseAttributes) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SecureEmbedUpdateResponseAttributes) HasToken() bool {
	return o != nil && o.Token != nil
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SecureEmbedUpdateResponseAttributes) SetToken(v string) {
	o.Token = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SecureEmbedUpdateResponseAttributes) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateResponseAttributes) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SecureEmbedUpdateResponseAttributes) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SecureEmbedUpdateResponseAttributes) SetUrl(v string) {
	o.Url = &v
}

// GetViewingPreferences returns the ViewingPreferences field value if set, zero value otherwise.
func (o *SecureEmbedUpdateResponseAttributes) GetViewingPreferences() SecureEmbedViewingPreferences {
	if o == nil || o.ViewingPreferences == nil {
		var ret SecureEmbedViewingPreferences
		return ret
	}
	return *o.ViewingPreferences
}

// GetViewingPreferencesOk returns a tuple with the ViewingPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateResponseAttributes) GetViewingPreferencesOk() (*SecureEmbedViewingPreferences, bool) {
	if o == nil || o.ViewingPreferences == nil {
		return nil, false
	}
	return o.ViewingPreferences, true
}

// HasViewingPreferences returns a boolean if a field has been set.
func (o *SecureEmbedUpdateResponseAttributes) HasViewingPreferences() bool {
	return o != nil && o.ViewingPreferences != nil
}

// SetViewingPreferences gets a reference to the given SecureEmbedViewingPreferences and assigns it to the ViewingPreferences field.
func (o *SecureEmbedUpdateResponseAttributes) SetViewingPreferences(v SecureEmbedViewingPreferences) {
	o.ViewingPreferences = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecureEmbedUpdateResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CredentialSuffix != nil {
		toSerialize["credential_suffix"] = o.CredentialSuffix
	}
	if o.DashboardId != nil {
		toSerialize["dashboard_id"] = o.DashboardId
	}
	if o.GlobalTime != nil {
		toSerialize["global_time"] = o.GlobalTime
	}
	if o.GlobalTimeSelectable != nil {
		toSerialize["global_time_selectable"] = o.GlobalTimeSelectable
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SelectableTemplateVars != nil {
		toSerialize["selectable_template_vars"] = o.SelectableTemplateVars
	}
	if o.ShareType != nil {
		toSerialize["share_type"] = o.ShareType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
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
func (o *SecureEmbedUpdateResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt              *string                                 `json:"created_at,omitempty"`
		CredentialSuffix       *string                                 `json:"credential_suffix,omitempty"`
		DashboardId            *string                                 `json:"dashboard_id,omitempty"`
		GlobalTime             *SecureEmbedGlobalTime                  `json:"global_time,omitempty"`
		GlobalTimeSelectable   *bool                                   `json:"global_time_selectable,omitempty"`
		Id                     *string                                 `json:"id,omitempty"`
		SelectableTemplateVars []SecureEmbedSelectableTemplateVariable `json:"selectable_template_vars,omitempty"`
		ShareType              *SecureEmbedShareType                   `json:"share_type,omitempty"`
		Status                 *SecureEmbedStatus                      `json:"status,omitempty"`
		Title                  *string                                 `json:"title,omitempty"`
		Token                  *string                                 `json:"token,omitempty"`
		Url                    *string                                 `json:"url,omitempty"`
		ViewingPreferences     *SecureEmbedViewingPreferences          `json:"viewing_preferences,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "credential_suffix", "dashboard_id", "global_time", "global_time_selectable", "id", "selectable_template_vars", "share_type", "status", "title", "token", "url", "viewing_preferences"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.CredentialSuffix = all.CredentialSuffix
	o.DashboardId = all.DashboardId
	if all.GlobalTime != nil && all.GlobalTime.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GlobalTime = all.GlobalTime
	o.GlobalTimeSelectable = all.GlobalTimeSelectable
	o.Id = all.Id
	o.SelectableTemplateVars = all.SelectableTemplateVars
	if all.ShareType != nil && !all.ShareType.IsValid() {
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
	o.Token = all.Token
	o.Url = all.Url
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
